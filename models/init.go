package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB : Centralize database instance
var DB *gorm.DB

// InitialDatabase : Initial database function to connect with DBMS
func InitialDatabase() error {
	var connectErr error

	connectionString := fmt.Sprintf("host=35.240.191.214 port=5432 user=postgres dbname=fillgoods-lab password=IkDD43cIamwavO7s")
	DB, connectErr = gorm.Open("postgres", connectionString)
	if connectErr != nil {
		fmt.Println(connectErr)
		return connectErr
	}

	autoMigrateDatabase()

	fmt.Println("Database connection successfully...")
	return nil
}

func autoMigrateDatabase() {
	DB.AutoMigrate(&StudentScoreName{})
}
