package handlers

import (
	"fmt"
	"net/http"

	"github.com/chrslex/bookings-mini-project/pkg/config"
	"github.com/chrslex/bookings-mini-project/pkg/models"
	"github.com/chrslex/bookings-mini-project/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(cfg *config.AppConfig) *Repository {
	return &Repository{
		App: cfg,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again !"

	remoteIP := repo.App.Session.Get(r.Context(), "remote_ip")
	stringMap["remote_ip"] = fmt.Sprint(remoteIP)
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
