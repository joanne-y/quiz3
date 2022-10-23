// Filename: internal/data/todo.go

package data

import (
	"time"

	"todo.joanneyong.net/internal/validator"
)

type Todo struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Details   string    `json:"details"`
	Priority  string    `json:"priority"`
	Status    string    `json:"status"`
	Version   int32     `json:"version"`
}

func ValidateTodo(v *validator.Validator, todo *Todo) {
	// Use the Check() method to execute our validation checks
	v.Check(todo.Name != "", "name", "must be provided")
	v.Check(len(todo.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(todo.Details != "", "details", "must be provided")
	v.Check(len(todo.Details) <= 200, "details", "must not be more than 200 bytes long")

	v.Check(todo.Priority != "", "priority", "must be provided")
	v.Check(len(todo.Priority) <= 200, "priority", "must not be more than 200 bytes long")

	v.Check(todo.Status != "", "status", "must be provided")
	v.Check(len(todo.Status) <= 200, "status", "must not be more than 200 bytes long")

}
