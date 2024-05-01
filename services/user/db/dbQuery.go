package db

func CreateUserQuery() string {

	return `INSERT INTO users_models( name , user_name , phone_number , email , password , auth_token)
	VALUES( $1 , $2 , $3 , $4 , $5 , $6 )`
}
