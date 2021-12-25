package handler

import (
	"gameservice/internal/model"
	"net/http"

	"github.com/rs/zerolog"
)

type Handler struct {
	logger  *zerolog.Logger
	service Service
}

type Service interface {
	GetSteamPriceByName(name string) (model.GamePriceResponse, error)
	GetGOGPriceByName(name string) (model.GamePriceResponse, error)
	GetSteamPayPriceByName(name string) (model.GamePriceResponse, error)
	GetPlatiruPriceByName(name string) (model.GamePriceResponse, error)
	GetAllPricesByName(name string) ([]model.GamePriceResponse, error)
}

func New(logger *zerolog.Logger, service Service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

func (h *Handler) SteamPriceHandler(w http.ResponseWriter, r *http.Request) {
	appName := r.URL.Query().Get("name")

	resp, err := h.service.GetSteamPriceByName(appName)
	if err != nil {
		h.logger.Info().Err(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Internal Server Error"}`))
		return
	}
	writeResponse(w, http.StatusOK, resp)
}

func (h *Handler) GOGPriceHandler(w http.ResponseWriter, r *http.Request) {
	appName := r.URL.Query().Get("name")

	resp, err := h.service.GetGOGPriceByName(appName)
	if err != nil {
		h.logger.Info().Err(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Internal Server Error"}`))
		return
	}
	writeResponse(w, http.StatusOK, resp)
}

func (h *Handler) SteamPayPriceHandler(w http.ResponseWriter, r *http.Request) {
	appName := r.URL.Query().Get("name")
	resp, err := h.service.GetSteamPayPriceByName(appName)
	if err != nil {
		h.logger.Info().Err(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Internal Server Error"}`))
		return
	}
	writeResponse(w, http.StatusOK, resp)
}

func (h *Handler) PlatiruHandler(w http.ResponseWriter, r *http.Request) {
	appName := r.URL.Query().Get("name")
	resp, err := h.service.GetPlatiruPriceByName(appName)
	if err != nil {
		h.logger.Info().Err(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Internal Server Error"}`))
		return
	}
	writeResponse(w, http.StatusOK, resp)
}

func (h *Handler) CompareHandler(w http.ResponseWriter, r *http.Request) {
	appName := r.URL.Query().Get("name")
	resp, err := h.service.GetAllPricesByName(appName)
	if err != nil {
		h.logger.Info().Err(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Internal Server Error"}`))
		return
	}
	writeResponse(w, http.StatusOK, resp)
}
