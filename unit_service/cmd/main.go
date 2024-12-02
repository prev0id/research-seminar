package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"unit_service/internal/handlers"
	"unit_service/internal/metrics"
	"unit_service/internal/repository"
	"unit_service/internal/service"
)

func main() {
	logFile, err := os.OpenFile("service.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatalf("could not open log file: %v", err)
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(logFile, nil))
	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(slog.LevelWarn)

	unitRepo := repository.NewUnitRepository()
	unitService := service.NewUnitService(unitRepo, logger)
	unitHandler := handlers.NewUnitHandler(unitService)

	metrics.InitPrometheusMetrics()

	http.HandleFunc("GET /units", unitHandler.GetAvailable)
	http.HandleFunc("POST /units", unitHandler.AddUnit)
	http.HandleFunc("GET /metrics", metrics.ServeMetrics)

	logger.Info("Starting service")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Error("Server stopped", "error", err)
	}
}
