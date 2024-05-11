package db

func CreateUserQuery() string {

	return `INSERT INTO users_models( name , phone_number )
	VALUES( $1 , $2 )`
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
