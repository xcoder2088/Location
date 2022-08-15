package handlers

import (
	"net/http"

	"github.com/Xcoder2088/Location/pkg/handlers/config"
	"github.com/Xcoder2088/Location/pkg/handlers/render"
	"github.com/Xcoder2088/Location/pkg/models"
)

// the repository used by the handlers
var Repo *Repositary

// Repository is the resopitory type
type Repositary struct {
	App *config.AppConfig
}

// NewRepo create a neew repository
func NewRepo(a *config.AppConfig) *Repositary {
	return &Repositary{
		App: a,
	}
}

// NewHandlers sets the repository for handler
func NewHandlers(r *Repositary) {
	Repo = r
}

// Home is the home page handler
func (m *Repositary) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

// About is the about us page handler
func (m *Repositary) About(w http.ResponseWriter, r *http.Request) {
	//perfor some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again."

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	//send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
