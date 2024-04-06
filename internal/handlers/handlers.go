package handlers

import (
	"fmt"
	"github.com/aville22/greeneats/internal/calculations"
	"github.com/aville22/greeneats/internal/config"
	"github.com/aville22/greeneats/internal/forms"
	"github.com/aville22/greeneats/internal/models"
	"github.com/aville22/greeneats/internal/render"
	"net/http"
	"strconv"
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
	var emptyProfile models.ProfileForm
	data := make(map[string]interface{})
	data["profileform"] = emptyProfile
	render.RenderTemplate(w, r, "profile.page.html", &models.TemplateData{Form: forms.New(nil),
		Data: data})
}
func (m *Repository) PostProfile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)

	profile := models.ProfileForm{}
	profile.Weight, err = strconv.ParseFloat(r.Form.Get("weight"), 64)

	profile.Height, err = strconv.ParseFloat(r.Form.Get("height"), 64)

	profile.Age, err = strconv.ParseFloat(r.Form.Get("age"), 64)

	profile.Gender = r.Form.Get("gender")

	activityStr := r.Form.Get("activity")
	profile.Activity, err = strconv.ParseFloat(activityStr, 64)
	form.Has("weight", r)
	if !form.Valid() {
		data := make(map[string]interface{})
		data["profileform"] = profile
		render.RenderTemplate(w, r, "profile.page.html", &models.TemplateData{Form: form, Data: data})
		return
	}
	profile.Goal = r.Form.Get("goal")

	jsonResult, err := calculations.CalculateCalories(profile)
	if err != nil {
		http.Error(w, "Failed to calculate calories", http.StatusInternalServerError)
		return
	}

	// Отправляем результат в формате JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResult)
}
