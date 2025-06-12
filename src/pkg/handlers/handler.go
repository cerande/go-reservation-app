package handlers

import (
	"net/http"

	"github.com/cerande/go-reservation-app/src/models"
	"github.com/cerande/go-reservation-app/src/pkg/config"
	"github.com/cerande/go-reservation-app/src/pkg/renders"
)

var Repo *Handler

type Handler struct {
	config *config.AppConfig
}

func NewHandler(ac *config.AppConfig) *Handler {
	return &Handler{
		config: ac,
	}
}

func NewRepository(h *Handler) {
	Repo = h
}

func (h *Handler) Home(wr http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	h.config.Session.Put(r.Context(), "remote_ip", remoteIP)

	renders.RenderTemplate(wr, "home.page.tmpl", &models.TemplateData{})
}

func (h *Handler) About(wr http.ResponseWriter, r *http.Request) {

	remoteIP := h.config.Session.GetString(r.Context(), "remote_ip")

	stm := map[string]string{"remote_ip": remoteIP}

	renders.RenderTemplate(wr, "about.page.tmpl", &models.TemplateData{
		StringMap: stm,
	})
}
