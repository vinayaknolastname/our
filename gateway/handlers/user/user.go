package user

import "database/sql"

type User struct {
	db *sql.DB
}

type UserErrorStruct struct {
	Status  int
	Message string
}

type UserSuccessStruct struct {
	Status   int
	Message  int
	response string
}

func (u User) createNewUser() {

}
