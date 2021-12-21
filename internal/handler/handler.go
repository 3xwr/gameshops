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
	GetAppPriceByID(ID int) (string, error)
	GetAppIDByName(name string) (int, error)
	GetSteamPriceByName(name string) (model.GamePriceResponse, error)
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
