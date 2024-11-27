package entity

import (
	"gorm.io/gorm"
)

type TourPackages struct {
	gorm.Model
	TourName	string `json:"tour_name"`
	PackageCode	string
	Duration	string

	ProvinceID	uint
	Province	Provinces	`gorm:"foreignKey:ProvinceID"`

	TourPrices	[]TourPrices	`gorm:"foreignKey:TourPackageID"`

	TourSchedules	[]TourSchedules	`gorm:"foreignKey:TourPackageID"`

	TourImages	[]TourImages	`gorm:"foreignKey:TourPackageID"`

	Transportations	[]Transportations	`gorm:"foreignKey:TourPackageID"`

	Meals		[]Meals			`gorm:"foreignKey:TourPackageID"`

	Accommodations	[]Accommodations	`gorm:"foreignKey:TourPackageID"`

	Activities	[]Activities	`gorm:"foreignKey:TourPackageID"`

	TourDescriptions	*TourDescriptions	`gorm:"foreignKey:TourPackageID"`
}