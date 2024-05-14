package types

type ReactionOnMessageStruct struct {
	MessageId int32  `json:"message_id"`
	Reaction  string `json:"reaction"`
	ReactorId int32  `json:"reactor_id"`
	ChatId    int32  `json:"chat_id"`
}
