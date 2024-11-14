package entity

import "gorm.io/gorm"

type RoomTypes struct {
	gorm.Model
	TypeName	string

	BookingRooms	[]BookingRooms	`gorm:"foreignKey:RoomTypeID"`
}