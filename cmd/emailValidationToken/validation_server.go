package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/enum"
	"warrant-api/pkg/repo"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormDB *gorm.DB

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	user := getenv("POSTGRES_USER", "admin")
	password := getenv("POSTGRES_PASSWORD", "xuuoH8FXDSTQkWFA7QRg")
	host := getenv("POSTGRES_HOST", "postgres")
	port := getenv("POSTGRES_PORT", "5432")
	dbname := getenv("POSTGRES_DB", "dev")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user, password, host, port, dbname)
	sqlDB, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	gormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{Logger: newLogger})
	newLogger.Info(context.Background(), "connectino string  --- ", connectionString)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/confirmations/:email", confirmationsHandler)
	r.GET("/qr", qrCodeHandler)
	r.Run(":9877")

}

func qrCodeHandler(g *gin.Context) {
	uuid := repo.GenUUID.Next()
	image, err := qrcode.Encode("shiggy"+uuid.String(), qrcode.Medium, 250)
	if err != nil {
		g.JSON(500, gin.H{"error": "server error"})
	} else {
		g.Data(http.StatusOK, "image/png", image)
	}
}
func confirmationsHandler(g *gin.Context) {
	email, ok := g.Params.Get("email")
	if !ok {
		g.JSON(404, gin.H{"message": "please provide email"})
		return
	}

	confirmation := model.Confirmation{}

	dbRes := gormDB.Model(&confirmation).Joins("left join users on users.id = confirmations.user_id").
		Where("users.email", email).Where("confirmations.status", enum.Unconfirmed).First(&confirmation)
	if dbRes.Error != nil {
		g.JSON(404, gin.H{"message": "confirmation not found"})
		g.Abort()
		return
	}
	confirmation.Status = enum.Confirmed
	g.JSON(200, gin.H{"message": confirmation.Code})
}
