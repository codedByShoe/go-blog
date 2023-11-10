package user

import (
	"net/http"

	"github.com/codedbyshoe/go-blog/pkg/templates"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type Handler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	repo            Repository
	templateService *templates.TemplateService
	store           sessions.CookieStore
}

func NewUserHandler(r Repository, templateService *templates.TemplateService, store sessions.CookieStore) Handler {
	return &userHandler{
		repo:            r,
		templateService: templateService,
		store:           store,
	}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// render the registration form
		h.templateService.Render(w, "auth/register.html", "layout.html", nil)
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
		existingUser, err := h.repo.GetUserByUsername(user.Username)
		if err == nil && existingUser != nil {
			// TODO: this most likely needs to be changed to a redirect with flash messages
			http.Error(w, "username already taken", http.StatusUnauthorized)
			return
		}
		// on success redirect to login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.templateService.Render(w, "auth/login.html", "layout.html", nil)
	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, err := h.repo.GetUserByUsername(username)

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			http.Error(w, "Invalid Login Credentials", http.StatusUnauthorized)
			return
		}

		// TODO: set up session for logged in user
		session, _ := h.store.Get(r, "user-session")
		session.Values["user-id"] = user.ID
		session.Save(r, w)

		SetLoggedInUser(user)

		// Redirect to a protected page or the home page idk yet
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *userHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := h.store.Get(r, "user-session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
