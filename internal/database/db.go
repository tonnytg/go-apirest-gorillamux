package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	DB *gorm.DB
}

func (db *Db) Connector(dbName string) *gorm.DB {
	
	// dns postgres
	host := "localhost"
	user := "root"
	password := "root"
	port := 5432
	sslmode := "disable"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, user, password, dbName, port, sslmode)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func (db *Db) Save(items []string) {

	tx := db.DB.Create(items)
	if tx != nil {
		fmt.Println("Error:", tx)
	}
}
