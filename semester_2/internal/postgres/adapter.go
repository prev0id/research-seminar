package postgres

import (
	sql "github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5/pgxpool"

	"calendar_app/internal/domain"
)

const (
	eventTableName = "events"
	columnID       = "id"
)

type Adapter struct {
	conn       *pgxpool.Pool
	table      string
	eventTable *sql.Struct
}

func New(conn *pgxpool.Pool) *Adapter {
	return &Adapter{
		conn:       conn,
		eventTable: sql.NewStruct(new(domain.Event)).For(sql.PostgreSQL),
	}
}
