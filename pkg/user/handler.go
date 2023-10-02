package user

import (
	"fmt"
	"net/http"

	"github.com/codedbyshoe/go-blog/utils"
)

type Handler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service Service
}

func NewHandler(s Service) Handler {
	return &handler{service: s}
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// render the registration form
		utils.RenderTemplate(w, "auth/register.html", nil)
	case "POST":
		// process the form submission
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")

		// Create a new user instance
		user := &User{
			Username: username,
			Password: password,
			Email:    email,
		}
		// User the service to reqister the user
		err := h.service.Register(user)

		if err != nil {
			http.Error(w, "Failed to register", http.StatusInternalServerError)
			return
		}
		// on success redirect to login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.RenderTemplate(w, "auth/login.html", nil)
	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, err := h.service.Login(username, password)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid Login Credentials", http.StatusUnauthorized)
			return
		}

		// TODO: set up session for logged in user
		session, _ := utils.Store.Get(r, "user-session")
		session.Values["user-id"] = user.ID
		session.Save(r, w)

		SetLoggedInUser(user)

		// Redirect to a protected page or the home page idk yet
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *handler) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "user-session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
