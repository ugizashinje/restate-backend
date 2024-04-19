package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	"warrant-api/pkg/config"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(cfg config.DBConfig) (*gorm.DB, error) {
	// connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable ssl=disable",
	// 	cfg.Host, cfg.Port, cfg.User, cfg.Password)
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
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
	sqlDB, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	gormConfig := &gorm.Config{}
	if cfg.DbLogging {
		gormConfig.Logger = newLogger
	} else {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), gormConfig)
	newLogger.Info(context.Background(), "connectino string  ", connectionString)

	// tables := []interface{}{
	// 	&model.Address{},
	// 	&model.Company{},
	// 	&model.Change{},
	// 	&model.Vehicle{},
	// 	&model.User{},
	// 	&model.DriverCategory{},
	// 	&model.Confirmation{},
	// 	&model.Login{},
	// 	&model.Route{},
	// 	&model.Repair{},
	// 	&model.ShippingInvoice{},
	// 	&model.TransportCost{},
	// 	&model.User{},
	// 	&model.UserCompany{},
	// 	&model.Vehicle{},
	// 	&model.Warrant{},
	// 	&model.WarrantEvent{},
	// }

	// gormDB.Statement.Migrator().DropTable(tables...)
	// if err := gormDB.AutoMigrate(tables...); err != nil {
	// 	log.Fatal("Failed to do migration ", err.Error())
	// }

	// err = gormDB.SetupJoinTable(&model.User{}, "Companies", &model.UserCompany{})
	// if err != nil {
	// 	return nil, err
	// }
	return gormDB, err
}
