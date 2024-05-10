package models

type UsersModel struct {
	ID          uint   `gorm:"PRIMARY_KEY;auto_increment;unique"`
	Name        string `gorm:"NOT NULL"`
	PhoneNumber int32  `gorm:"NOT NULL;index"`
}

// type PgAddressModel struct {
// 	ID uint `gorm:"PRIMARY_KEY;auto_increment;unique"`

// 	Pincode   string `gorm:"DEFAULT:NULL"`
// 	Road      string `gorm:"DEFAULT:NULL"`
// 	City      string `gorm:"DEFAULT:NULL"`
// 	Area      string `gorm:"DEFAULT:NULL"`
// 	Landmarks string `gorm:"DEFFAULT:NULL"`
// }

// type PgFeaturesModel struct {
// 	ID uint `gorm:"PRIMARY_KEY;auto_increment;unique"`

// 	Ac            bool `gorm:"DEFAULT:False"`
// 	Cooler        bool `gorm:"DEFAULT:False"`
// 	WaterPurifier bool `gorm:"DEFAULT:False"`
// 	Tv            bool `gorm:"DEFAULT:False"`
// 	Fridge        bool `gorm:"DEFAULT:False"`
// }
