package handlers

import (
	"fmt"
	"github.com/aville22/greeneats/pkg/calculations"
	"net/http"
	"strconv"

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
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{StringMap: stringMap})
	fmt.Println(stringMap)
}
func (m *Repository) Profile(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "profile.page.html", &models.TemplateData{})
}
func (m *Repository) PostProfile(w http.ResponseWriter, r *http.Request) {
	weightStr := r.Form.Get("weight")
	heightStr := r.Form.Get("height")
	ageStr := r.Form.Get("age")
	gender := r.Form.Get("gender")
	activityStr := r.Form.Get("activity")
	goal := r.Form.Get("goal")

	// Преобразование строковых значений в числовой формат
	weight, err := strconv.ParseFloat(weightStr, 64)
	if err != nil {
		http.Error(w, "Invalid weight", http.StatusBadRequest)
		return
	}

	height, err := strconv.ParseFloat(heightStr, 64)
	if err != nil {
		http.Error(w, "Invalid height", http.StatusBadRequest)
		return
	}

	age, err := strconv.ParseFloat(ageStr, 64)
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}

	activity, err := strconv.ParseFloat(activityStr, 64)
	if err != nil {
		http.Error(w, "Invalid activity level", http.StatusBadRequest)
		return
	}
	json, err := calculations.CalculateCalories(weight, height, age, gender, goal, activity)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(json))
}
