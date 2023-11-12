package comment

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	AddComment(c *fiber.Ctx) error
}

type commentHandler struct {
	model Model
}

func NewCommentHandler(m Model) Handler {
	return &commentHandler{
		model: m,
	}
}

func (h *commentHandler) AddComment(c *fiber.Ctx) error {
	idStr := c.FormValue("post_id")
	postId, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		// TODO: fix error handling in comments
		panic(err)
	}

	var comment Comment
	comment.Author = c.FormValue("author")
	comment.Content = c.FormValue("content")
	comment.ID = postId

	if err := h.model.AddComment(comment); err != nil {
		panic(err)
	}
	// Redirect back to the respective post
	redirectURL := fmt.Sprintf("/posts/%d", postId)
	return c.Redirect(redirectURL)
}
