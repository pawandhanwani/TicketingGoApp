package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	TicketId  uint
	CreatedBy uint   `gorm:"notnull:true"`
	Comment   string `gorm:"notnull:true"`
}
