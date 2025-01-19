package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	"calendar_app/internal/domain"
)

var emptyDate = time.Time{}

func (s *Server) ListEvents(w http.ResponseWriter, r *http.Request) {
	events, err := s.db.List(r.Context())
	if err != nil {
		http.Error(w, "Failed to list events", http.StatusInternalServerError)
		log.Printf("s.db.List: %s", err.Error())
		return
	}

	writeJSON(w, http.StatusOK, events)
}

func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event domain.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("json.Decode: %s", err.Error())
		return
	}

	if event.Name == "" || event.CreatedBy == "" || event.Date.Equal(emptyDate) {
		http.Error(w, "Name, CreatedBy and date are required", http.StatusBadRequest)
		return
	}

	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()

	id, err := s.db.Insert(r.Context(), event)
	if err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		log.Printf("s.db.Insert: %s", err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	eventIDStr := chi.URLParam(r, eventVar)
	eventID, err := strconv.ParseInt(eventIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	var event domain.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("json.Decode: %s", err.Error())
		return
	}

	if event.Name == "" || event.CreatedBy == "" {
		http.Error(w, "Name and CreatedBy are required", http.StatusBadRequest)
		return
	}

	event.ID = uint64(eventID)
	event.UpdatedAt = time.Now()

	if err := s.db.Update(r.Context(), event); err != nil {
		http.Error(w, "Failed to update event", http.StatusInternalServerError)
		log.Printf("s.db.Update: %s", err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "Event updated successfully",
	})
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventIDStr := chi.URLParam(r, eventVar)
	eventID, err := strconv.ParseInt(eventIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	if err := s.db.Delete(r.Context(), eventID); err != nil {
		http.Error(w, "Failed to delete event", http.StatusInternalServerError)
		log.Printf("s.db.Delete: %s", err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "Event deleted successfully",
	})
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to write JSON response: %s", err.Error())
	}
}
