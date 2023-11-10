package comment

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/codedbyshoe/go-blog/pkg/templates"
)

type Handler interface {
	AddComment(w http.ResponseWriter, r *http.Request)
}

type commentHandler struct {
	repo            Repository
	templateService *templates.TemplateService
}

func NewCommentHandler(r Repository, templateService *templates.TemplateService) Handler {
	return &commentHandler{
		repo:            r,
		templateService: templateService,
	}
}

func (h *commentHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		author := r.FormValue("author")
		content := r.FormValue("content")
		postIdStr := r.FormValue("post_id")
		postId, err := strconv.Atoi(postIdStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := h.repo.AddComment(postId, author, content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Redirect back to the respective post
		redirectURL := fmt.Sprintf("/posts/%d", postId)
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

}
