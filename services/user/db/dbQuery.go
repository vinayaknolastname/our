package db

func CreateUserQuery() string {

	return `INSERT INTO users_models( name , phone_number )
	VALUES( $1 , $2 )`
}

func CreateMessageQuery() string {

	return `INSERT INTO message_models(  content , chat_id , sender_id , date_time , delivered_too , readed_by , is_deleted , seq )
	VALUES( $1 , $2 ,  $3 , $4 , $5 , $6 , $7 , $8) RETURNING id`
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
	SET chats = ARRAY_APPEND(chats, $2)
	WHERE id = $1`
}

func CreateChatQuery() string {

	return `INSERT INTO chats_models( name , type , members )
	VALUES( $1 , $2 ,  $3) RETURNING id`
}
