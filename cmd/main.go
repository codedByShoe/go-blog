package main

import (
	"database/sql"
	"net/http"

	"github.com/codedbyshoe/go-blog/pkg/comment"
	"github.com/codedbyshoe/go-blog/pkg/post"
	"github.com/codedbyshoe/go-blog/pkg/user"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// init db
	db, err := sql.Open("sqlite3", "blog.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//init sessions cookie store
	// init post func
	postRepo := post.NewRepository(db)
	postService := post.NewService(postRepo)
	postHandler := post.NewHandler(postService)
	// init comment func
	commentRepo := comment.NewRepository(db)
	commentService := comment.NewService(commentRepo)
	commentHandler := comment.NewHandler(commentService)
	// init user func
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	//init routes
	http.HandleFunc("/", postHandler.GetAllPosts)
	http.HandleFunc("/posts/", postHandler.GetSinglePost)
	http.HandleFunc("/create_post", postHandler.CreatePost)
	http.HandleFunc("/update_post", postHandler.UpdatePost)
	http.HandleFunc("/delete_post", postHandler.DeletePost)
	http.HandleFunc("/add_comment", commentHandler.AddComment)
	http.HandleFunc("/login", userHandler.Login)
	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/logout", userHandler.Logout)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
