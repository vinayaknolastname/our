package models

import "github.com/lib/pq"

type ChatsModel struct {
	ID      int32         `gorm:"PRIMARY_KEY;auto_increment;unique"`
	Name    string        `gorm:"NOT NULL"`
	Type    int32         `gorm:"NOT NULL"`
	Members pq.Int32Array `gorm:"NOT NULL;type:int[];index;fk:users_models(id)"`
}
