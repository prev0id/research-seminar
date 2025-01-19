package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"calendar_app/internal/postgres"
)

const (
	pingPath  = "/ping"
	eventPath = "/event"

	eventVar = "event_id"
)

type Server struct {
	db     *postgres.Adapter
	router chi.Router
}

func New(db *postgres.Adapter) *Server {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Heartbeat(pingPath))

	s := &Server{
		db:     db,
		router: r,
	}

	s.routes()

	return s
}

func (s *Server) Run(addr string) error {
	log.Printf("starting server at %s", addr)
	defer log.Print("shut down server")

	return http.ListenAndServe(addr, s.router)
}

func (s *Server) routes() {
	s.router.Route(eventPath, func(r chi.Router) {
		r.Get("/", s.ListEvents)
		r.Post("/", s.CreateEvent)
		r.Put("/{event_id}", s.UpdateEvent)
		r.Delete("/{event_id}", s.DeleteEvent)
	})
}
