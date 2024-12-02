package repository

import (
	"errors"
	"sync"

	"unit_service/internal/domain"
)

type UnitRepository struct {
	m     sync.RWMutex
	units map[string]domain.Unit
}

func NewUnitRepository() *UnitRepository {
	return &UnitRepository{
		units: make(map[string]domain.Unit),
	}
}

func (r *UnitRepository) AddUnit(unit domain.Unit) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, exists := r.units[unit.ID]; exists {
		return errors.New("unit already exists")
	}

	r.units[unit.ID] = unit
	return nil
}

func (r *UnitRepository) GetAvailableUnits() []domain.Unit {
	r.m.RLock()
	defer r.m.RUnlock()

	availableUnits := make([]domain.Unit, 0, len(r.units))
	for _, unit := range r.units {
		if unit.Available {
			availableUnits = append(availableUnits, unit)
		}
	}

	return availableUnits
}
