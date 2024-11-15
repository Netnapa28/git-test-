package entity

import "gorm.io/gorm"

type VehicleTypes struct {
	gorm.Model
	TypeName	string

	Vehicles	[]Vehicles `gorm:"foreignKey:VehicleTypeID"`
}