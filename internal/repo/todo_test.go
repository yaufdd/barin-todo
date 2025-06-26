package repo

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestInsert_Positive(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening stub database connection: %v", err)
	}
	defer db.Close()
	boil.SetDB(db)

	repo := &TodoRepo{DB: db}
	ctx := context.Background()

	future := time.Now().Add(24 * time.Hour).Format(time.RFC3339)
	mock.ExpectExec("INSERT INTO `todos`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("SELECT `id`,`status`,`deleted` FROM `todos` WHERE `id`=?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "status", "deleted"}).AddRow(1, 0, false))

	id, err := repo.Insert(ctx, "title", "desc", future)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if id == 0 {
		t.Errorf("expected non-zero id")
	}
}

func TestInsert_EmptyTitleOrDesc(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()
	boil.SetDB(db)

	repo := &TodoRepo{DB: db}
	ctx := context.Background()
	future := time.Now().Add(24 * time.Hour).Format(time.RFC3339)

	_, err := repo.Insert(ctx, "", "desc", future)
	if err == nil || err.Error() != "title or desc is empty" {
		t.Errorf("expected error for empty title")
	}
	_, err = repo.Insert(ctx, "title", "", future)
	if err == nil || err.Error() != "title or desc is empty" {
		t.Errorf("expected error for empty desc")
	}
}

func TestInsert_PastDeadline(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()
	boil.SetDB(db)

	repo := &TodoRepo{DB: db}
	ctx := context.Background()
	past := time.Now().Add(-24 * time.Hour).Format(time.RFC3339)

	_, err := repo.Insert(ctx, "title", "desc", past)
	if err == nil || err.Error() != "deadline is before today" {
		t.Errorf("expected error for past deadline")
	}
}

func TestInsert_EmptyDeadline(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	boil.SetDB(db)

	repo := &TodoRepo{DB: db}
	ctx := context.Background()

	mock.ExpectExec("INSERT INTO `todos`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("SELECT `id`,`status`,`deleted` FROM `todos` WHERE `id`=?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "status", "deleted"}).AddRow(1, 0, false))

	id, err := repo.Insert(ctx, "title", "desc", "")
	if err != nil {
		t.Errorf("unexpected error for empty deadline: %v", err)
	}
	if id == 0 {
		t.Errorf("expected non-zero id for empty deadline")
	}
}

func TestUpdate_Positive(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening stub database connection: %v", err)
	}
	defer db.Close()
	boil.SetDB(db)

	repo := &TodoRepo{DB: db}
	ctx := context.Background()

	mock.ExpectQuery("(?i)select .* from `todos` where `id`=?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "status", "deadline", "created_at", "updated_at", "deleted"}).
			AddRow(1, "old", "old", 0, nil, time.Now(), time.Now(), false))

	mock.ExpectExec("(?i)update `todos` set .* where `id`=?").
		WillReturnResult(sqlmock.NewResult(1, 1))

	title := "new"
	desc := "new desc"
	status := int8(1)
	err = repo.Update(ctx, 1, &title, &desc, &status)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestUpdate_EmptyTitleOrDesc(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening stub database connection: %v", err)
	}
	defer db.Close()
	boil.SetDB(db)

	repo := &TodoRepo{DB: db}
	ctx := context.Background()

	mock.ExpectQuery("(?i)select .* from `todos` where `id`=?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "status", "deadline", "created_at", "updated_at", "deleted"}).
			AddRow(1, "old", "old", 0, nil, time.Now(), time.Now(), false))

	mock.ExpectExec("(?i)update `todos` set .* where `id`=?").
		WillReturnResult(sqlmock.NewResult(1, 1))

	empty := ""
	desc := "desc"
	status := int8(1)
	err = repo.Update(ctx, 1, &empty, &desc, &status)
	if err == nil {
		t.Errorf("expected error for empty title")
	}

	title := "title"
	err = repo.Update(ctx, 1, &title, &empty, &status)
	if err == nil {
		t.Errorf("expected error for empty desc")
	}
}

func TestUpdate_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening stub database connection: %v", err)
	}
	defer db.Close()
	boil.SetDB(db)

	repo := &TodoRepo{DB: db}
	ctx := context.Background()

	// Мокаем SELECT для поиска задачи (не найдено)
	mock.ExpectQuery("SELECT (.+) FROM `todos` WHERE `todos`.`id` = ?").
		WithArgs(999).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "status", "deadline", "created_at", "updated_at", "deleted"}))

	title := "title"
	desc := "desc"
	status := int8(1)
	err = repo.Update(ctx, 999, &title, &desc, &status)
	if err == nil {
		t.Errorf("expected error for not found id")
	}
}

func TestSoftDelete_Positive(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening stub database connection: %v", err)
	}
	defer db.Close()
	boil.SetDB(db)

	repo := &TodoRepo{DB: db}
	ctx := context.Background()

	// Мокаем SELECT для поиска задачи
	mock.ExpectQuery("(?i)select .* from `todos` where `id`=?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "status", "deadline", "created_at", "updated_at", "deleted"}).
			AddRow(1, "old", "old", 0, nil, time.Now(), time.Now(), false))

	// Мокаем UPDATE
	mock.ExpectExec("(?i)update `todos` set .* where `id`=?").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.SoftDelete(ctx, 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSoftDelete_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening stub database connection: %v", err)
	}
	defer db.Close()
	boil.SetDB(db)

	repo := &TodoRepo{DB: db}
	ctx := context.Background()

	// Мокаем SELECT для поиска задачи (не найдено)
	mock.ExpectQuery("(?i)select .* from `todos` where `id`=?").
		WithArgs(999).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "status", "deadline", "created_at", "updated_at", "deleted"}))

	err = repo.SoftDelete(ctx, 999)
	if err == nil {
		t.Errorf("expected error for not found id")
	}
}

func TestList_Positive(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening stub database connection: %v", err)
	}
	defer db.Close()
	boil.SetDB(db)

	repo := &TodoRepo{DB: db}
	ctx := context.Background()

	// Мокаем SELECT для списка задач
	mock.ExpectQuery("(?i)select .* from `todos`.*").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "status", "deadline", "created_at", "updated_at", "deleted"}).
			AddRow(1, "test", "desc", 0, nil, time.Now(), time.Now(), false))

	todos, err := repo.List(ctx, -1, 0, 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(todos) != 1 {
		t.Errorf("expected 1 todo, got %d", len(todos))
	}
}

func TestList_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening stub database connection: %v", err)
	}
	defer db.Close()
	boil.SetDB(db)

	repo := &TodoRepo{DB: db}
	ctx := context.Background()

	// Мокаем ошибку при SELECT
	mock.ExpectQuery("(?i)select .* from `todos`.*").
		WillReturnError(sqlmock.ErrCancelled)

	_, err = repo.List(ctx, -1, 0, 0)
	if err == nil {
		t.Errorf("expected db error")
	}
}
