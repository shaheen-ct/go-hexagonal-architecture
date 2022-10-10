package pgx

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Option struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

// Setup closure to set up and connect to postgres database
func Setup(ot Option) func() *gorm.DB {
	return func() *gorm.DB {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", ot.Host, ot.User, ot.Password, ot.DBName, ot.Port, ot.SSLMode)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalln("failed to connect to db")
		}
		return db
	}
}
