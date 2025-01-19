package postgres

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"

	"calendar_app/internal/domain"
)

func (a *Adapter) List(ctx context.Context) ([]domain.Event, error) {
	var result []domain.Event

	query, args := a.eventTable.
		SelectFrom(eventTableName).
		Build()

	if err := pgxscan.Select(ctx, a.conn, &result, query, args...); err != nil {
		return nil, fmt.Errorf("pgxscan.Select: %w", err)
	}

	return result, nil
}

func (a *Adapter) Update(ctx context.Context, event domain.Event) error {
	builder := a.eventTable.Update(eventTableName, event)
	builder.Where(builder.Equal(columnID, event.ID))
	query, args := builder.Build()

	if _, err := a.conn.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("db.conn.Exec: %w", err)
	}

	return nil
}

func (a *Adapter) Insert(ctx context.Context, event domain.Event) (int64, error) {
	builder := a.eventTable.InsertInto(eventTableName, event)
	builder.SQL("RETURNING " + columnID)
	query, args := builder.Build()

	fmt.Println(query, args)

	var result int64
	if err := pgxscan.Get(ctx, a.conn, &result, query, args...); err != nil {
		return 0, fmt.Errorf("pgxscan.Get: %w", err)
	}

	return result, nil
}

func (a *Adapter) Delete(ctx context.Context, id int64) error {
	builder := a.eventTable.DeleteFrom(eventTableName)
	builder.Where(builder.Equal(columnID, id))
	query, args := builder.Build()

	if _, err := a.conn.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("a.conn.Exec: %w", err)
	}

	return nil
}
