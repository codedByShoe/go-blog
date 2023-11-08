package main

import (
	"database/sql"
	"net/http"

	"github.com/codedbyshoe/go-blog/pkg/comment"
	"github.com/codedbyshoe/go-blog/pkg/post"
	"github.com/codedbyshoe/go-blog/pkg/templates"
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

	// init template service
	templateService, err := templates.NewTemplateService()
	if err != nil {
		panic(err)
	}

	// init post func
	postRepo := post.NewPostRepository(db)
	postService := post.NewPostService(postRepo)
	postHandler := post.NewPostHandler(postService, templateService)
	// init comment func
	commentRepo := comment.NewCommentRepository(db)
	commentService := comment.NewCommentService(commentRepo)
	commentHandler := comment.NewCommentHandler(commentService, templateService)
	// init user func
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService, templateService)

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
