package repo

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"sort"
	"time"

	"github.com/ericbamba/barin-todo/internal/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TodoRepo struct{ DB *sql.DB }

func (r *TodoRepo) Insert(ctx context.Context, title, desc, deadline string) (int64, error) {
	t := &models.Todo{
		Title:       title,
		Description: null.StringFrom(desc),
		Deadline:    null.TimeFromPtr(parseDeadline(deadline)),
		Status:      0,
	}
	if title == "" || desc == "" {
		slog.Error("Insert error", "err", "title or desc is empty")
		return 0, errors.New("title or desc is empty")
	}
	if t.Deadline.Valid && t.Deadline.Time.Before(time.Now()) {
		slog.Error("Insert error", "err", "deadline is before today")
		return 0, errors.New("deadline is before today")
	}
	if err := t.Insert(ctx, r.DB, boil.Infer()); err != nil {
		slog.Error("Insert error", "err", err)
		return 0, err
	}
	return t.ID, nil
}

func parseDeadline(deadline string) *time.Time {
	if deadline == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339, deadline)
	if err != nil {
		return nil
	}
	return &t
}

func (r *TodoRepo) Update(ctx context.Context, id int64, title, desc *string, status *int8) error {
	if (title != nil && *title == "") || (desc != nil && *desc == "") {
		return errors.New("title or desc is empty")
	}
	t, err := models.FindTodo(ctx, r.DB, id)
	if err != nil {
		slog.Error("Update: FindTodo error", "err", err, "id", id)
		return err
	}

	if title != nil {
		t.Title = *title
	}
	if desc != nil {
		t.Description = null.StringFrom(*desc)
	}
	if status != nil {
		t.Status = *status
	}
	t.UpdatedAt = time.Now()

	_, err = t.Update(ctx, r.DB, boil.Infer())
	if err != nil {
		slog.Error("Update: Update error", "err", err, "id", id)
	}
	return err
}

func (r *TodoRepo) SoftDelete(ctx context.Context, id int64) error {
	t, err := models.FindTodo(ctx, r.DB, id)
	if err != nil {
		slog.Error("SoftDelete: FindTodo error", "err", err, "id", id)
		return err
	}
	t.Deleted = true
	_, err = t.Update(ctx, r.DB, boil.Whitelist("deleted"))
	if err != nil {
		slog.Error("SoftDelete: Update error", "err", err, "id", id)
	}
	return err
}

func (r *TodoRepo) List(ctx context.Context, statusFilter int8, sortBy, sortOrder int8) ([]*models.Todo, error) {
	var todos []*models.Todo
	var err error

	switch statusFilter {
	case 1, 2: // OPEN, DONE
		todos, err = models.Todos(
			models.TodoWhere.Deleted.EQ(false),
			models.TodoWhere.Status.EQ(statusFilter),
		).All(ctx, r.DB)
	case 0: // ALL (WITHOUT DELETED)
		todos, err = models.Todos(models.TodoWhere.Deleted.EQ(false)).All(ctx, r.DB)
	case -1: // ALL (WITH DELETED)
		todos, err = models.Todos().All(ctx, r.DB)
	default:
		todos, err = models.Todos(models.TodoWhere.Deleted.EQ(false)).All(ctx, r.DB)
	}

	if err != nil {
		slog.Error("List error", "err", err, "statusFilter", statusFilter)
		return nil, err
	}

	sortTodos(todos, sortBy, sortOrder)
	return todos, nil
}

func sortTodos(todos []*models.Todo, sortBy, sortOrder int8) {
	if sortBy == 1 { // DEADLINE
		sort.Slice(todos, func(i, j int) bool {
			if sortOrder == 0 { // ASC
				return todos[i].Deadline.Time.Before(todos[j].Deadline.Time)
			} else { // DESC
				return todos[j].Deadline.Time.Before(todos[i].Deadline.Time)
			}
		})
	}
}
