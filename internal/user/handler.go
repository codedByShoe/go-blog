package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	"golang.org/x/crypto/bcrypt"
)

type Handler interface {
	ShowRegister(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	ShowLogin(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type userHandler struct {
	model   Model
	session session.Session
}

func NewUserHandler(r Model, sess session.Session) Handler {
	return &userHandler{
		model:   r,
		session: sess,
	}
}

func (h *userHandler) ShowRegister(c *fiber.Ctx) error {
	return c.Render("auth/register", nil)
}

func (h *userHandler) Register(c *fiber.Ctx) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Create a new user instance from form values
	user := User{
		Username: c.FormValue("username"),
		Password: string(hashedPassword),
		Email:    c.FormValue("email"),
	}
	// User the service to reqister the user
	existingUser, err := h.model.GetUserByUsername(user.Username)
	if err == nil && existingUser == user {
		// TODO: this most likely needs to be changed to a redirect with flash messages
		panic(err)
	}
	// on success redirect to login page
	return c.Redirect("/login")
}

func (h *userHandler) ShowLogin(c *fiber.Ctx) error {
	return c.Render("auth/login", nil)
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := h.model.GetUserByUsername(username)

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// http.Error(w, "Invalid Login Credentials", http.StatusUnauthorized)
		// TODO: handle failed login
		panic(err)
	}

	session := h.session.Get(c)
	session.Set("user_id", user.ID)
	err = session.Save()
	if err != nil {
		return err
	}
	// TODO: Redirect to a protected page or the home page idk yet
	return c.Redirect("/")
}

func (h *userHandler) Logout(c *fiber.Ctx) error {
	session := h.session.Get(c)
	err := session.Destroy()
	if err != nil {
		return err
	}
	return c.Redirect("/")
}
