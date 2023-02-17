package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	CreatedBy   uint      `gorm:"notnull:true"`
	Title       string    `gorm:"notnull:true"`
	Description string    `gorm:"notnull:true"`
	Comment     []Comment `gorm:"foreignKey:TicketId;references:ID"`
	AssignedTo  uint
	Status      string `gorm:"notnull:true"`
}
