package blog

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"hochbaum.dev/blog/blog/markdown"
	"html/template"
	"net/http"
)

// Blog defines the application itself and allows for control over posts and users.
type Blog interface {
	// Publish effectively starts the blog app.
	Publish() error

	// Storage returns the Storage used by the app, used for managing the database.
	Storage() Storage
}

// Post is a single published post on the blog.
type Post struct {
	gorm.Model // gorm.Model contains fields such as ID and CreatedAt.

	Title string `gorm:"type:varchar(120);not null"`

	// This is used for an ugly workaround. The htmlContent is actually using Content for formatting,
	// which the blog later translates to HTML. For accessing the actual htmlContent, the HTML()
	// function should be used, as it uses the fields for lazy loading.
	htmlContent template.HTML `gorm:"-"`
	Content     string        `gorm:"size:max;not null"`
}

// HTML returns the htmlContent of the post. The htmlContent is lazily translated from Content to HTML. See
// Post documentation.
func (post *Post) HTML() template.HTML {
	if post.htmlContent == "" {
		post.htmlContent = markdown.ToHtml(post.Content)
	}
	return post.htmlContent
}

// server is the default Blog implementation, which spins up a web server using the mux library.
type server struct {
	router  *mux.Router
	storage Storage
}

// New constructs an instance of a blog server and returns a pointer to it.
func New(router *mux.Router, storage Storage) Blog {
	return &server{
		router:  router,
		storage: storage,
	}
}

// ...
func (server *server) Publish() error {
	router := server.router
	router.HandleFunc("/", server.handleHome)
	router.NotFoundHandler = http.HandlerFunc(server.handle404)

	// Set up a static file route.
	path := "/static/"
	rel := "." + path
	router.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(http.Dir(rel))))

	return http.ListenAndServe(":80", router)
}

// ...
func (server *server) Storage() Storage {
	return server.storage
}

// handleHome is handling requests to the homepage.
func (server *server) handleHome(w http.ResponseWriter, req *http.Request) {
	// TODO: This should be cached.
	posts, err := server.Storage().AllPosts()
	if err != nil {
		panic(err)
	}

	args := map[string]interface{}{
		"Posts": posts,
	}
	tmpl := templates["index.html"]
	_ = tmpl.Execute(w, args)
}

// handle404 deals with 404 cases, that being requests to non-existent routes.
func (server *server) handle404(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("Drake? Where are the posts?"))
}
