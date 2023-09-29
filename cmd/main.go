package main

import (
	"database/sql"
	"net/http"

	"github.com/codedbyshoe/go-blog/pkg/comment"
	"github.com/codedbyshoe/go-blog/pkg/post"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// init db
	db, err := sql.Open("sqlite3", "../blog.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// init post func
	postRepo := post.NewRepository(db)
	postService := post.NewService(postRepo)
	postHandler := post.NewHandler(postService)
	// init comment func
	commentRepo := comment.NewRepository(db)
	commentService := comment.NewService(commentRepo)
	commentHandler := comment.NewHandler(commentService)

	//routes
	http.HandleFunc("/", postHandler.GetAllPosts)
	http.HandleFunc("/posts/", postHandler.GetSinglePost)
	http.HandleFunc("/create_post", postHandler.CreatePost)
	http.HandleFunc("/update_post", postHandler.UpdatePost)
	http.HandleFunc("/delete_post", postHandler.DeletePost)
	http.HandleFunc("/add_comment", commentHandler.AddComment)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
