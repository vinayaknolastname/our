package models

import (
	"time"

	"github.com/lib/pq"
)

type MessageModel struct {
	ID           int32         `gorm:"PRIMARY_KEY;auto_increment;unique"`
	Content      string        `gorm:""`
	ChatId       int32         `gorm:"NOT NULL;index"`
	SenderId     int32         `gorm:"NOT NULL;index"`
	DateTime     time.Time     `gorm:"default:CURRENT_TIMESTAMP"`
	DeliveredToo pq.Int32Array `gorm:"type:int[];index"`
	ReadedBy     pq.Int32Array `gorm:"type:int[];index"`
	IsDeleted    bool          `gorm:"default:FALSE;index"`
	Seq          int32         `gorm:"NOT NULL;index"`
}
