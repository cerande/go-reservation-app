package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cerande/go-reservation-app/src/internal/config"
	"github.com/cerande/go-reservation-app/src/internal/renders"
	"github.com/cerande/go-reservation-app/src/models"
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

	renders.RenderTemplate(r, wr, "home.page.tmpl", &models.TemplateData{})
}

func (h *Handler) About(wr http.ResponseWriter, r *http.Request) {

	remoteIP := h.config.Session.GetString(r.Context(), "remote_ip")

	stm := map[string]string{"remote_ip": remoteIP}

	renders.RenderTemplate(r, wr, "about.page.tmpl", &models.TemplateData{
		StringMap: stm,
	})
}

func (h *Handler) Generals(wr http.ResponseWriter, r *http.Request) {

	renders.RenderTemplate(r, wr, "generals.page.tmpl", &models.TemplateData{})
}

func (h *Handler) Majors(wr http.ResponseWriter, r *http.Request) {

	renders.RenderTemplate(r, wr, "majors.page.tmpl", &models.TemplateData{})
}

func (h *Handler) Availability(wr http.ResponseWriter, r *http.Request) {

	renders.RenderTemplate(r, wr, "search-availability.page.tmpl", &models.TemplateData{})
}

func (h *Handler) PostAvailability(wr http.ResponseWriter, r *http.Request) {

	wr.Write([]byte("Post search availability"))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (h *Handler) AvailabilityJSON(wr http.ResponseWriter, r *http.Request) {

	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	wr.Header().Set("Content-Type", "application/json")
	_, err = wr.Write(out)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) Reservation(wr http.ResponseWriter, r *http.Request) {

	renders.RenderTemplate(r, wr, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (h *Handler) Contact(wr http.ResponseWriter, r *http.Request) {

	renders.RenderTemplate(r, wr, "contact.page.tmpl", &models.TemplateData{})
}
