package database

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var instance *gorm.DB
var once sync.Once

// Instance returns the singleton gorm.DB ptr with a connection to the db
func Instance() *gorm.DB {
	once.Do(func() {
		instance = getDatabaseConnection()
	})

	return instance
}

// Close closes the gorm.DB connection
func Close() {
	fmt.Println("Closing db connection...")
	instance.Close()
}

func getDatabaseConnection() *gorm.DB {
	db, err := gorm.Open("mysql", viper.GetString("database.connection_string"))
	if err != nil {
		log.Fatalf("failed to connect to db: %v\n", err)
	}

	return db
}
