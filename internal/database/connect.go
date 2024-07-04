package database

import (
	"log"
	"os"
	"sample/api/internal/database/model"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB(dsn string, isDebug bool, isMigration bool) {
	logLevel := logger.Error
	if isDebug {
		logLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond * 500, // Slow SQL threshold
			LogLevel:                  logLevel,               // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,                   // Don't include params in the SQL log
			Colorful:                  false,                  // Disable color
		},
	)
	// Connect to the DB and initialize the DB variable
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connection Opened to Database")

	db, err := DB.DB()
	if err != nil {
		panic("failed to get db")
	}
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(30)
	db.SetConnMaxLifetime(time.Hour)

	if isMigration {
		// Migrate the database
		if err := DB.AutoMigrate(&model.User{}); err != nil {
			panic("failed to migrate users")
		}
		if err := DB.AutoMigrate(&model.Telegram{}); err != nil {
			panic("failed to migrate telegram")
		}
		log.Println("Database Migrated")
	}
}
