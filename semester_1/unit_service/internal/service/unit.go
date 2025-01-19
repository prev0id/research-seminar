package service

import (
	"fmt"
	"log/slog"

	"unit_service/internal/domain"
	"unit_service/internal/metrics"
	"unit_service/internal/repository"

	"github.com/google/uuid"
)

type UnitService struct {
	unitRepo *repository.UnitRepository
}

func NewUnitService(unitRepo *repository.UnitRepository, logger *slog.Logger) *UnitService {
	return &UnitService{
		unitRepo: unitRepo,
	}
}

func (s *UnitService) AddUnit(address, description string, available bool) error {
	unit := domain.Unit{
		ID:          uuid.NewString(),
		Address:     address,
		Description: description,
		Available:   available,
	}

	if err := s.unitRepo.AddUnit(unit); err != nil {
		slog.Error("failed to add unit", "error", err)
		return fmt.Errorf("failed to add unit: %w", err)
	}

	metrics.TotalUnits.Inc()

	return nil
}

func (s *UnitService) GetAvailableUnits() []domain.Unit {
	return s.unitRepo.GetAvailableUnits()
}
