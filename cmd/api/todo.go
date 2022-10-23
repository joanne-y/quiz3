// Filename: cmd/api/todo.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"todo.joanneyong.net/internal/data"
)

// createTodoHandler for the "POST /v1/todo" endpoint
func (app *application) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Our target decode destination
	var input struct {
		Name     string `json:"name"`
		Details  string `json:"details"`
		Priority string `json:"priority"`
		Status   string `json:"status"`
	}
	// Initialize a new json.Decoder instance
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	// Display the request
	fmt.Fprintf(w, "%+v\n", input)
}

// showTodoHandler for the "GET /v1/todo/:id" endpoint
func (app *application) showTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Create a new instance of the Todo struct containing the ID we extracted
	// from our URL and some sample data
	todo := data.Todo{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "Laundry",
		Details:   "Wash jeans",
		Priority:  "High",
		Status:    "Completed",
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"todo": todo}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
