package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"github.com/spf13/cobra"

	todov1 "github.com/ericbamba/barin-todo/gen/proto/todo/v1"
	todov1connect "github.com/ericbamba/barin-todo/gen/proto/todo/v1/todov1connect"
)

func main() {
	serverURL := os.Getenv("SERVER_URL")
	if serverURL == "" {
		serverURL = "http://localhost:8080"
	}
	client := todov1connect.NewTodoServiceClient(
		&http.Client{},
		serverURL,
	)

	root := &cobra.Command{Use: "todo"}

	create := &cobra.Command{
		Use:   "create",
		Short: "Add todo",
		Run: func(cmd *cobra.Command, args []string) {
			title, _ := cmd.Flags().GetString("title")
			desc, _ := cmd.Flags().GetString("desc")
			deadline, _ := cmd.Flags().GetString("deadline")
			resp, err := client.CreateTodo(context.Background(),
				connect.NewRequest(&todov1.CreateTodoRequest{
					Title: title, Description: desc, Deadline: deadline,
				}))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Created todo with ID:", resp.Msg.Todo.Id, "Title:", resp.Msg.Todo.Title,
				"Description:", resp.Msg.Todo.Description,
				"Status:", resp.Msg.Todo.Status,
				"Deadline:", resp.Msg.Todo.Deadline,
			)
		},
	}
	create.Flags().String("title", "", "title")
	create.MarkFlagRequired("title")
	create.Flags().String("desc", "", "description")
	create.MarkFlagRequired("desc")
	create.Flags().String("deadline", "", "deadline")
	create.MarkFlagRequired("deadline")
	root.AddCommand(create)

	update := &cobra.Command{
		Use:   "update",
		Short: "Update todo",
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := cmd.Flags().GetInt64("id")
			title, _ := cmd.Flags().GetString("title")
			desc, _ := cmd.Flags().GetString("desc")
			status, _ := cmd.Flags().GetInt32("status")

			_, err := client.UpdateTodo(context.Background(),
				connect.NewRequest(&todov1.UpdateTodoRequest{
					Id: id, Title: title, Description: desc, Status: todov1.Status(status),
				}))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Updated successfully")
		},
	}
	update.Flags().Int64("id", 0, "todo ID")
	update.MarkFlagRequired("id")
	update.Flags().String("title", "", "title")
	update.Flags().String("desc", "", "description")
	update.Flags().Int32("status", 0, "status (0=UNKNOWN, 1=OPEN, 2=DONE)")
	root.AddCommand(update)

	delete := &cobra.Command{
		Use:   "delete",
		Short: "Delete todo (soft delete)",
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := cmd.Flags().GetInt64("id")

			_, err := client.DeleteTodo(context.Background(),
				connect.NewRequest(&todov1.DeleteTodoRequest{
					Id: id,
				}))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Deleted successfully")
		},
	}
	delete.Flags().Int64("id", 0, "todo ID")
	delete.MarkFlagRequired("id")
	root.AddCommand(delete)

	list := &cobra.Command{
		Use:   "list",
		Short: "List todos",
		Run: func(cmd *cobra.Command, args []string) {
			statusFilter, _ := cmd.Flags().GetInt8("status")
			sortBy, _ := cmd.Flags().GetInt8("sort-by")
			sortOrder, _ := cmd.Flags().GetInt8("sort-order")

			resp, err := client.ListTodos(context.Background(),
				connect.NewRequest(&todov1.ListTodosRequest{
					StatusFilter: todov1.Status(statusFilter),
					SortBy:       todov1.SortBy(sortBy),
					SortOrder:    todov1.SortOrder(sortOrder),
				}))
			if err != nil {
				log.Fatal(err)
			}
			for _, t := range resp.Msg.Todos {
				fmt.Printf("[%d] %s\n", t.Id, t.Title)
				if t.Description != "" {
					fmt.Printf("    Description: %s\n", t.Description)
				}
				fmt.Printf("    Status: %s\n", t.Status)
				if t.Deadline != "" {
					fmt.Printf("    Deadline: %s\n", t.Deadline)
				}
				fmt.Println()
			}
		},
	}
	list.Flags().Int8("status", 0, "filter by status (0=ALL, 1=OPEN, 2=DONE, -1=ALL(WITH DELETED))")
	list.Flags().Int8("sort-by", 0, "sort by deadline")
	list.Flags().Int8("sort-order", 0, "sort order (0=ASC, 1=DESC)")
	root.AddCommand(list)

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}

}
