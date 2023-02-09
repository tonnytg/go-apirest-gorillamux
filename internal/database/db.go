package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	User     string
	Password string
	Port     int
	DbName   string
	SSLMode  string
}

type Db struct {
	DB *gorm.DB
}

func NewDB() *Db {
	return &Db{}
}

func (db *Db) Connector(c Config) *gorm.DB {

	// dns postgres
	c.Host = "localhost"
	c.User = "root"
	c.Password = "root"
	c.Port = 5432
	c.SSLMode = "disable"

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.Host, c.User, c.Password, c.DbName, c.Port, c.SSLMode)
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
