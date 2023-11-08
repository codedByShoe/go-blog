package user

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

var loggedInUser *User

func SetLoggedInUser(user *User) {
	loggedInUser = user
}

func GetLoggedInUser() *User {
	return loggedInUser
}
