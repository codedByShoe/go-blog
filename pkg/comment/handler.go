package comment

import (
	"fmt"
	"net/http"
	"strconv"
)

type Handler interface {
	AddComment(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service Service
}

func NewHandler(s Service) Handler {
	return &handler{service: s}
}

func (h *handler) AddComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	author := r.FormValue("author")
	content := r.FormValue("content")
	postIdStr := r.FormValue("post_id")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.service.AddCommentToPost(postId, author, content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Redirect back to the respective post
	redirectURL := fmt.Sprintf("/posts/%d", postId)
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
