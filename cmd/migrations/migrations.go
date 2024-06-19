package main

import (
	"log"
	"warrant-api/pkg/config"
	"warrant-api/pkg/db"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/wire"
)

func main() {
	config.Init("dev")
	wire.Svc = wire.Init("dev")
	gormDB, _ := db.Init(config.Db)

	// 	cfg.Host, cfg.Port, cfg.User, cfg.Password)
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // Slow SQL threshold
	// 		LogLevel:                  logger.Info, // Log level
	// 		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	// 		ParameterizedQueries:      true,        // Don't include params in the SQL log
	// 		Colorful:                  true,        // Disable color
	// 	},
	// )
	// connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
	// 	config.Db.User, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Dbname)
	// sqlDB, err := sql.Open("postgres", connectionString)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// gormDB, err = gorm.Open(postgres.New(postgres.Config{
	// 	Conn: sqlDB,
	// }), &gorm.Config{Logger: newLogger})
	// newLogger.Info(context.Background(), "connectino string  ", connectionString)

	tables := []interface{}{
		&model.Ad{},
		&model.Address{},
		&model.Apartment{},
		&model.Change{},
		&model.CollectionEntry{},
		&model.Collection{},
		&model.Company{},
		&model.Confirmation{},
		&model.Location{},
		&model.Like{},
		&model.Login{},
		&model.UserCollection{},
		// &model.UserCompany{},
		&model.User{},
		&model.Warrant{},
	}

	gormDB.Statement.Migrator().DropTable(tables...)
	if err := gormDB.AutoMigrate(tables...); err != nil {
		log.Fatal("Failed to do migration ", err.Error())
	}
	// gormDB.SetupJoinTable(&model.User{}, "Companies", &model.UserCompany{})
	gormDB.SetupJoinTable(&model.User{}, "Collections", &model.UserCollection{})
	gormDB.SetupJoinTable(&model.User{}, "Likes", &model.Like{})

}
