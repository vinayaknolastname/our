package user

type User interface {
	createNewUser()
}

func createNewUser(user User) {
	user.createNewUser()
}
