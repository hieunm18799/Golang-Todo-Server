package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db_todos *gorm.DB

func ConnectDatabase(db_config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_config.Host, db_config.User, db_config.Password, db_config.DB_Name, db_config.PSQL_Port)

	Db_todos, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("? Connected Successfully to the Database")

	return
}

func CloseDatabaseConnection(db *gorm.DB) {
	temp, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	temp.Close()
}
