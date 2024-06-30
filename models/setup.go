package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

type Database struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

type Model struct {
	Model interface{}
}

func (db *Database) DbInit() {
	db.Name = "belajar_go_restapi"
	db.User = "root"
	db.Password = ""
	db.Host = "localhost"
	db.Port = ":3306"
}

func (config *Config) ConnectDB() {
	var err error
	db := Database{}
	db.DbInit()
	dsn := fmt.Sprintf("%s:%s@tcp(%s%s)/%s", db.User, db.Password, db.Host, db.Port, db.Name) + "?charset=utf8mb4&parseTime=True&loc=Local"
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	config.dbMigrate()
}

func (config *Config) dbMigrate() {
	for _, model := range registerModels() {
		err := config.DB.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func registerModels() []Model {
	return []Model{
		{Model: Product{}},
	}
}
