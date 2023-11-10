package post

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/codedbyshoe/go-blog/pkg/templates"
	"github.com/codedbyshoe/go-blog/pkg/user"
)

type Handler interface {
	GetAllPosts(w http.ResponseWriter, r *http.Request)
	GetSinglePost(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
}

type postHandler struct {
	repo            Repository
	templateService *templates.TemplateService
}

func NewPostHandler(repository Repository, templateService *templates.TemplateService) Handler {
	return &postHandler{
		repo:            repository,
		templateService: templateService,
	}
}

func (h *postHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.repo.GetAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	loggedInUser := user.GetLoggedInUser()

	var data struct {
		Title   string
		Content string
		Posts   []Post
		User    *user.User
	}

	data.Posts = posts
	data.User = loggedInUser

	h.templateService.Render(w, "index.html", "layout.html", data)
}

func (h *postHandler) GetSinglePost(w http.ResponseWriter, r *http.Request) {
	// Exctract the post ID from the url
	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/posts/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	post, err := h.repo.GetPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.templateService.Render(w, "single_post.html", "layout.html", post)
}

func (h *postHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl := template.Must(template.ParseFiles("templates/create_post.html"))
		tmpl.Execute(w, nil)
		h.templateService.Render(w, "create_post.html", "layout.html", nil)
	case "POST":
		title := r.FormValue("title")
		content := r.FormValue("content")

		if err := h.repo.CreatePost(title, content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

}

func (h *postHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		post, err := h.repo.GetPostByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.templateService.Render(w, "update_post.html", "layout.html", post)
	case "POST":
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		title := r.FormValue("title")
		content := r.FormValue("content")

		if err := h.repo.UpdatePost(id, title, content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/posts", http.StatusSeeOther)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *postHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.repo.DeletePost(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/posts", http.StatusSeeOther)
}
