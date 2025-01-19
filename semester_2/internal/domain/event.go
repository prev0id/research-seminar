package domain

import (
	"time"
)

type Event struct {
	ID          uint64    `db:"id" fieldopt:"omitempty" json:"id"`
	Name        string    `db:"name" fieldopt:"omitempty" json:"name"`
	Description string    `db:"description" fieldopt:"omitempty" json:"description"`
	CreatedBy   string    `db:"created_by" fieldopt:"omitempty" json:"created_by"`
	Date        time.Time `db:"date" fieldopt:"omitempty" json:"date"`
	UpdatedAt   time.Time `db:"updated_at" fieldopt:"omitempty" json:"updated_at"`
	CreatedAt   time.Time `db:"created_at" fieldopt:"omitempty" json:"created_at"`
}
