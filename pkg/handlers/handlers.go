package handlers

import (
	"fmt"
	"net/http"

	"github.com/siwodevilheart/bookings/pkg/config"
	"github.com/siwodevilheart/bookings/pkg/models"
	"github.com/siwodevilheart/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	fmt.Println(remoteIP)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TmplData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["test"] = "A test string"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	strMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TmplData{
		StringMap: strMap,
	})
}
