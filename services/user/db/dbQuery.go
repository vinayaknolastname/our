package db

func CreateUserQuery() string {

	return `INSERT INTO users_models( name , phone_number )
	VALUES( $1 , $2 )`
}

/////message model

func CreateMessageQuery() string {

	return `INSERT INTO message_models(  content , chat_id , sender_id , date_time , is_deleted , seq , media_link )
	VALUES( $1 , $2 ,  $3 , $4 , $5 , $6) RETURNING id`
}

func AddDeliveredToInMessageProper() string {

	return `UPDATE message_models
	SET delivered_too = ARRAY_APPEND(delivered_too, $2)
	WHERE id = $1`
}

func GetReactionQuery() string {
	return `SELECT * FROM reaction_on_chat_models WHERE msg_id=$1`
}

func GetMessageQuery() string {
	return `SELECT * FROM message_models WHERE chat_id=$1 AND seq=$2`
}

func GetUserQuery() string {

	return `SELECT * FROM users_models WHERE id=$1`
}

func GetChatRowQuery() string {

	return `SELECT * FROM chats_models WHERE id=$1`
}

func AddChatInUser() string {

	return `SELECT chats FROM users_models WHERE id=$1`
}

func UpdateSeqInChat() string {

	return `UPDATE chats_models
	 SET last_seq = $1
	 WHERE id = $2`
}

func AddChatInUserProper() string {

	return `UPDATE users_models
	SET chats = CASE
	          WHEN $2 = ANY(chats) THEN chats
			  ELSE ARRAY_APPEND(chats, $2)
			END  
	WHERE id = $1`
}

func CreateChatQuery() string {

	return `INSERT INTO chats_models( name , type , members )
	VALUES( $1 , $2 ,  $3) RETURNING id`
}

// //reaction
func CreateReactionQuery() string {

	return `INSERT INTO reaction_on_chat_models( reaction , msg_id , reactor_id , chat_id )
	VALUES( $1 , $2 ,  $3, $4) RETURNING id`
}
