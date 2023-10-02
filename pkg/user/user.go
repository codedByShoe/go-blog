package user

var loggedInUser *User

func SetLoggedInUser(user *User) {
	loggedInUser = user
}

func GetLoggedInUser() *User {
	return loggedInUser
}
