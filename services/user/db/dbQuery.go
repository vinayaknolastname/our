package db

func CreateUserQuery() string {

	return `INSERT INTO users_models( name , phone_number )
	VALUES( $1 , $2 )`
}
