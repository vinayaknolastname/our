package models

import (
	"github.com/lib/pq"
)

type PgBasicModel struct {
	ID              uint           `gorm:"PRIMARY_KEY;auto_increment;unique"`
	Name            string         `gorm:"NOT NULL"`
	UserName        string         `gorm:"NOT NULL"`
	Lat             string         `gorm:"index"`
	Log             string         `gorm:"index"`
	PhoneNumber     int32          `gorm:"NOT NULL;index"`
	Email           string         `gorm:"DEFAULT:NULL"`
	Rent            int32          `gorm:"DEFFAULT:NULL"`
	IsDeposite      bool           `gorm:"DEFAULT:False"`
	VerifiedPartner bool           `gorm:"DEFAULT:False"`
	Address         int32          `gorm:"DEFAULT:NULL`
	Images          pq.StringArray `gorm:"type:text[]"`
	Features        int32          `gorm:"foreignKey:ID;DEFAULT:NULL`
	Gender          int16          `gorm:"NOT NULL"`
	PreferedBy      int16          `gorm:"DEFAULT:NULL"`
	OwnerName       string         `gorm:"NOT NULL"`
	Rating          int16          `gorm:"DEFFAULT:NULL"`
	TotalBeds       int16          `gorm:"DEFFAULT:NULL"`
	EmptyBeds       int16          `gorm:"DEFFAULT:NULL"`
}

type PgAddressModel struct {
	ID uint `gorm:"PRIMARY_KEY;auto_increment;unique"`

	Pincode   string `gorm:"DEFAULT:NULL"`
	Road      string `gorm:"DEFAULT:NULL"`
	City      string `gorm:"DEFAULT:NULL"`
	Area      string `gorm:"DEFAULT:NULL"`
	Landmarks string `gorm:"DEFFAULT:NULL"`
}

type PgFeaturesModel struct {
	ID uint `gorm:"PRIMARY_KEY;auto_increment;unique"`

	Ac            bool `gorm:"DEFAULT:False"`
	Cooler        bool `gorm:"DEFAULT:False"`
	WaterPurifier bool `gorm:"DEFAULT:False"`
	Tv            bool `gorm:"DEFAULT:False"`
	Fridge        bool `gorm:"DEFAULT:False"`
}
