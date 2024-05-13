package db

func CreateUserQuery() string {

	return `INSERT INTO users_models( name , phone_number )
	VALUES( $1 , $2 )`
}

func CreateMessageQuery() string {

	return `INSERT INTO message_models(  content , chat_id , sender_id , date_time , is_deleted , seq )
	VALUES( $1 , $2 ,  $3 , $4 , $5 , $6) RETURNING id`
}

func AddDeliveredToInMessageProper() string {

	return `UPDATE message_models
	SET delivered_too = ARRAY_APPEND(delivered_too, $2)
	WHERE id = $1`
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
