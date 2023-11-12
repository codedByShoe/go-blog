package post

import (
	"strconv"

	// "github.com/codedbyshoe/go-blog/internal/user"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetAllPosts(c *fiber.Ctx) error
	GetSinglePost(c *fiber.Ctx) error
	ShowCreatePost(c *fiber.Ctx) error
	CreatePost(c *fiber.Ctx) error
	ShowUpdatePost(c *fiber.Ctx) error
	UpdatePost(c *fiber.Ctx) error
	DeletePost(c *fiber.Ctx) error
}

type postHandler struct {
	model Model
}

func NewPostHandler(m Model) Handler {
	return &postHandler{
		model: m,
	}
}

func (h *postHandler) GetAllPosts(c *fiber.Ctx) error {
	posts, err := h.model.GetAllPosts()
	if err != nil {
		panic(err)
	}

	return c.Render("index", fiber.Map{
		"Posts": posts,
		// "User":  loggedInUser,
	})
}

func (h *postHandler) GetSinglePost(c *fiber.Ctx) error {
	// Exctract the post ID from the url
	urlParam := c.Params("id")
	id, err := strconv.ParseUint(urlParam, 10, 64)
	if err != nil {
		panic(err)
	}
	post, err := h.model.GetPostByID(id)
	if err != nil {
		panic(err)
	}

	return c.Render("single_post", fiber.Map{
		"Post": post,
	})
}

func (h *postHandler) ShowCreatePost(c *fiber.Ctx) error {
	return c.Render("create_post", nil)
}

func (h *postHandler) CreatePost(c *fiber.Ctx) error {
	var post Post
	post.Title = c.FormValue("title")
	post.Content = c.FormValue("content")
	if err := h.model.CreatePost(post); err != nil {
		// TODO: Place in better error handling
		panic(err)
	}
	return c.RedirectBack("/")

}

func (h *postHandler) ShowUpdatePost(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	post, err := h.model.GetPostByID(id)
	if err != nil {
		panic(err)
	}
	return c.Render("update_post", fiber.Map{
		"Post": post,
	})
}

func (h *postHandler) UpdatePost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		panic(err)
	}
	var post Post
	post.ID = id
	post.Title = c.FormValue("title")
	post.Content = c.FormValue("content")

	if err := h.model.UpdatePost(post); err != nil {
		panic(err)
	}
	return c.RedirectBack("/posts")

}

func (h *postHandler) DeletePost(c *fiber.Ctx) error {

	id, err := strconv.ParseUint(c.FormValue("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	if err := h.model.DeletePost(id); err != nil {
		panic(err)
	}
	return c.RedirectBack("/posts")

}
