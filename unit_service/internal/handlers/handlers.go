package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"unit_service/internal/metrics"
	"unit_service/internal/service"
)

type AddUnitRequest struct {
	Address     string `json:"address"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
}

type UnitResponse struct {
	ID          string `json:"id"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
}

type UnitHandler struct {
	service *service.UnitService
}

func NewUnitHandler(service *service.UnitService) *UnitHandler {
	return &UnitHandler{service: service}
}

func (h *UnitHandler) AddUnit(w http.ResponseWriter, r *http.Request) {
	metrics.Requests.WithLabelValues("add unit").Inc()

	var req AddUnitRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.service.AddUnit(req.Address, req.Description, req.Available); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UnitHandler) GetAvailable(w http.ResponseWriter, _ *http.Request) {
	metrics.Requests.WithLabelValues("get_available").Inc()

	units := h.service.GetAvailableUnits()

	resp := make([]UnitResponse, 0, len(units))
	for _, unit := range units {
		resp = append(resp, UnitResponse{
			ID:          unit.ID,
			Description: unit.Description,
			Address:     unit.Address,
			Available:   unit.Available,
		})
	}

	body, err := json.Marshal(units)
	if err != nil {
		slog.Error("marshalling error", slog.String("error", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(body)
}
