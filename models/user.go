package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"notnull:true"`
	Email    string    `gorm:"notnull:true;unique:true"`
	Password string    `gorm:"notnull:true"`
	Tickets  []Ticket  `gorm:"foreignKey:CreatedBy;references:ID"`
	Comment  []Comment `gorm:"foreignKey:CreatedBy;references:ID"`
	UserType string
	IsActive bool `gorm:"default:true"`
}
