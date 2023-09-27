package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", postsHandler)
	http.HandleFunc("/posts/", singlePostHandler)
	http.HandleFunc("/create_post", createPostHandler)
	http.HandleFunc("/update_post", updatePostHandler)
	http.HandleFunc("/delete_post", deletePostHandler)
	http.HandleFunc("/add_comment", addCommentHandler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
