package models

type ChatsModel struct {
	ID          uint   `gorm:"PRIMARY_KEY;auto_increment;unique"`
	Name        string `gorm:"NOT NULL"`
	PhoneNumber int32  `gorm:"NOT NULL;index"`
}
