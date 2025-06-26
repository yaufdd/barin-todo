package handler

import (
	"context"
	"time"

	todov1 "github.com/ericbamba/barin-todo/gen/proto/todo/v1"
	todov1connect "github.com/ericbamba/barin-todo/gen/proto/todo/v1/todov1connect"
	"github.com/ericbamba/barin-todo/internal/repo"

	"connectrpc.com/connect"
)

type TodoHandler struct{ Repo *repo.TodoRepo }

// CreateTodo RPC
func (h *TodoHandler) CreateTodo(
	ctx context.Context,
	req *connect.Request[todov1.CreateTodoRequest],
) (*connect.Response[todov1.CreateTodoResponse], error) {

	id, err := h.Repo.Insert(ctx, req.Msg.Title, req.Msg.Description, req.Msg.Deadline)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var deadline string
	if req.Msg.Deadline != "" {
		deadline = req.Msg.Deadline
	}

	return connect.NewResponse(&todov1.CreateTodoResponse{Todo: &todov1.Todo{
		Id:          id,
		Title:       req.Msg.Title,
		Description: req.Msg.Description,
		Status:      todov1.Status_UNKNOWN,
		Deadline:    deadline,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}}), nil
}

// UpdateTodo RPC
func (h *TodoHandler) UpdateTodo(
	ctx context.Context,
	req *connect.Request[todov1.UpdateTodoRequest],
) (*connect.Response[todov1.UpdateTodoResponse], error) {

	status := int8(req.Msg.Status)
	err := h.Repo.Update(ctx, req.Msg.Id, &req.Msg.Title, &req.Msg.Description, &status)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&todov1.UpdateTodoResponse{}), nil
}

// DeleteTodo RPC (soft-delete)
func (h *TodoHandler) DeleteTodo(
	ctx context.Context,
	req *connect.Request[todov1.DeleteTodoRequest],
) (*connect.Response[todov1.DeleteTodoResponse], error) {

	if err := h.Repo.SoftDelete(ctx, req.Msg.Id); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&todov1.DeleteTodoResponse{}), nil
}

// ListTodos RPC
func (h *TodoHandler) ListTodos(
	ctx context.Context,
	req *connect.Request[todov1.ListTodosRequest],
) (*connect.Response[todov1.ListTodosResponse], error) {

	rows, err := h.Repo.List(ctx, int8(req.Msg.StatusFilter), int8(req.Msg.SortBy), int8(req.Msg.SortOrder))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	out := make([]*todov1.Todo, len(rows))
	for i, r := range rows {
		var deadline string
		if r.Deadline.Valid {
			deadline = r.Deadline.Time.Format(time.RFC3339)
		}
		out[i] = &todov1.Todo{
			Id:          r.ID,
			Title:       r.Title,
			Description: r.Description.String,
			Status:      todov1.Status(r.Status),
			Deadline:    deadline,
		}
	}
	return connect.NewResponse(&todov1.ListTodosResponse{Todos: out}), nil
}

func NewHandler(repo *repo.TodoRepo) todov1connect.TodoServiceHandler {
	return &TodoHandler{Repo: repo}
}
