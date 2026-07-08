package handler

import (
	"agency-site/internal/components/core"
	"agency-site/internal/components/home"
	"agency-site/internal/types"
	"fmt"
	"net/http"
	"strconv"
)

// Home handles the home page.
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.html(r.Context(), w, http.StatusOK, core.HTML("Example Site", h.Header, home.Page(h.todos)))
}

func (h *Handler) AddTodo(w http.ResponseWriter, r *http.Request) {
	task := r.FormValue("task")
	todo := types.TodoItem{
		Id:        h.nextId,
		Content:   task,
		Completed: false,
	}

	h.nextId++
	h.todos = append(h.todos, todo)
	h.html(r.Context(), w, http.StatusCreated, home.RenderTodos(h.todos))
}

func (h *Handler) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	fmt.Println(id)
	for i := range h.todos {
		if h.todos[i].Id == id {
			h.todos[i].Completed = !h.todos[i].Completed
			break
		}
	}
	fmt.Println(h.todos)

	h.html(
		r.Context(),
		w,
		http.StatusOK,
		home.RenderTodos(h.todos),
	)
}
