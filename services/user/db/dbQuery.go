package db

func CreateUserQuery() string {

	return `INSERT INTO users_models( name , phone_number )
	VALUES( $1 , $2 )`
}

func CreateChatQuery() string {

	return `INSERT INTO chats_models( name , type , members )
	VALUES( $1 , $2 ,  $3)`
}
