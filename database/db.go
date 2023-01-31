package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	DB *gorm.DB
}

func (db *Db) Connector() *gorm.DB {
	// dns postgres
	dsn := "host=localhost user=root password=root dbname=Names port=5432 sslmode=disable"
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
