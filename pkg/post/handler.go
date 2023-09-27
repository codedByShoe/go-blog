package post

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/codedbyshoe/go-blog/utils"
)

type Handler interface {
	GetAllPosts(w http.ResponseWriter, r *http.Request)
	GetPostByID(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service Service
}

func NewHandler(s Service) Handler {
	return &handler{service: s}
}

func (h *handler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := getAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RenderTemplate(w, "index.html", posts)
}

func (h *handler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	// Exctract the post ID from the url
	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/posts/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	post, err := getPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RenderTemplate(w, "single_post.html", post)
}

func (h *handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/create_post.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")

		if err := createPost(title, content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	}
}

func (h *handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		post, err := getPostByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.RenderTemplate(w, "update_post.html", post)
	} else if r.Method == "POST" {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		title := r.FormValue("title")
		content := r.FormValue("content")

		if err := updatePost(id, title, content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	}
}

func (h *handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := deletePost(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/posts", http.StatusSeeOther)
}
