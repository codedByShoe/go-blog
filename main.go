package main

import (
	"log"

	"github.com/codedbyshoe/go-blog/internal/comment"
	"github.com/codedbyshoe/go-blog/internal/post"
	"github.com/codedbyshoe/go-blog/internal/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// init db
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sess := session.New()

	// init post func
	postModel := post.NewPostModel(db)
	postHandler := post.NewPostHandler(postModel)
	// init comment func
	commentModel := comment.NewCommentModel(db)
	commentHandler := comment.NewCommentHandler(commentModel)
	// init user func
	userModel := user.NewUserModel(db)
	userHandler := user.NewUserHandler(userModel, *sess)

	engine := html.New("web/templates", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layout/main",
	})

	//init routes
	app.Get("/", postHandler.GetAllPosts)
	app.Get("/posts/:id", postHandler.GetSinglePost)
	app.Get("/post/create", postHandler.ShowCreatePost)
	app.Post("/post/create", postHandler.CreatePost)
	app.Get("/post/update/:id", postHandler.ShowUpdatePost)
	app.Put("/post/update/:id", postHandler.UpdatePost)
	app.Delete("/post/delete", postHandler.DeletePost)
	app.Post("/comment/create", commentHandler.AddComment)
	app.Get("/login", userHandler.ShowLogin)
	app.Post("/login", userHandler.Login)
	app.Get("/register", userHandler.ShowRegister)
	app.Post("/register", userHandler.Register)
	app.Get("/logout", userHandler.Logout)

	log.Fatal(app.Listen(":8000"))
}
