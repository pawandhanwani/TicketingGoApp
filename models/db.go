package models

import (
	"fmt"

	"../configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Kolkata", configs.EnvDbHost(), configs.EnvDbUser(), configs.EnvDbPassword(), configs.EnvDbName(), configs.EnvDbPort())
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Ticket{})
	db.AutoMigrate(&Comment{})
	DB = db
}

//   dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
//   db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
