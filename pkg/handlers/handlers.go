package handlers

import (
	"fmt"
	"net/http"

	"github.com/aville22/greeneats/pkg/config"
	"github.com/aville22/greeneats/pkg/models"
	"github.com/aville22/greeneats/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}
func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIp)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
	fmt.Println(stringMap)
}
