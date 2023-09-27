package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func addCommentHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := addComment(postId, author, content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Redirect back to the respective post
	redirectURL := fmt.Sprintf("/posts/%d", postId)
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
