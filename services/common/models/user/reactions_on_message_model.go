package models

type ReactionOnChatModel struct {
	ID        int32  `gorm:"PRIMARY_KEY;auto_increment;unique"`
	Reaction  string `gorm:"NOT NULL"`
	MsgId     int32  `gorm:"NOT NULL"`
	ReactorId int32  `gorm:"NOT NULL"`
	ChatId    int32  `gorm:"NOT NULL"`
}
