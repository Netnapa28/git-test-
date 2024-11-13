package entity

import "gorm.io/gorm"

type TourPrices struct {
	gorm.Model
	Price 	float32

	TourPackageID	uint
	TourPackage		TourPackages	`gorm:"foreignKey:TourPackageID"`

	PersonTypeID	uint
	PersonType		PersonTypes		`gorm:"foreignKey:PersonTypeID"`
}