package entity

import "gorm.io/gorm"

type PaymentStatus struct {
	gorm.Model
	StatusName	string

	Payments	[]Payments	`gorm:"foreignKey:PaymentStatusID"`
}