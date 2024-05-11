package models

type ChatsModel struct {
	ID      uint     `gorm:"PRIMARY_KEY;auto_increment;unique"`
	Name    string   `gorm:"NOT NULL"`
	Type    int32    `gorm:"NOT NULL"`
	Members []string `gorm:"NOT NULL;type:text[];index;fk:users_models(id)"`
}
