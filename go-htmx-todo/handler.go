package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	items, err := FetchTasks()
	if err != nil {
		log.Printf("error fetching tasks: %v\n", err)
		return
	}
	count, err := FetchCount()
	if err != nil {
		log.Printf("error fetching count: %v\n", err)
	}
	completedCount, err := FetchCountCompleted()
	if err != nil {
		log.Printf("error fetching completed count: %v\n", err)
	}
	data := Tasks{
		Items:          items,
		Count:          count,
		CompletedCount: completedCount,
	}
	Tmpl.ExecuteTemplate(w, "Base", data)
}

func HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		Tmpl.ExecuteTemplate(w, "Form", nil)
		return
	}
	item, err := InsertTask(title)
	if err != nil {
		log.Printf("error inserting task: %v\n", err)
		return
	}
	count, err := FetchCount()
	if err != nil {
		log.Printf("error fetching count: %v\n", err)
	}
	w.WriteHeader(http.StatusCreated)
	Tmpl.ExecuteTemplate(w, "Form", nil)
	Tmpl.ExecuteTemplate(w, "Item", map[string]any{"Item": item, "SwapOOB": true})
	Tmpl.ExecuteTemplate(w, "TotalCount", map[string]any{
		"Count":   count,
		"SwapOOB": true,
	})
}

func HandleToggleTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("error parsing id into int: %v\n", err)
		return
	}
	_, err = ToggleTask(id)
	if err != nil {
		log.Printf("error toggle task: %v\n", err)
		return
	}
	completed, err := FetchCountCompleted()
	if err != nil {
		log.Printf("error fetching completed count: %v\n", err)
	}
	Tmpl.ExecuteTemplate(w, "CompletedCount", map[string]any{
		"Count":   completed,
		"SwapOOB": true,
	})
}

func HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("error parsing id into int: %v\n", err)
		return
	}
	if err := DeleteTask(r.Context(), id); err != nil {
		log.Printf("error deleting task: %v\n", err)
	}
	count, err := FetchCount()
	if err != nil {
		log.Printf("error fetching count: %v\n", err)
	}
	completedCount, err := FetchCountCompleted()
	if err != nil {
		log.Printf("error fetching completed count: %v\n", err)
	}
	Tmpl.ExecuteTemplate(w, "TotalCount", map[string]any{"Count": count, "SwapOOB": true})
	Tmpl.ExecuteTemplate(w, "CompletedCount", map[string]any{"Count": completedCount, "SwapOOB": true})
}

func HandleEditTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("error parsing id into int: %v\n", err)
		return
	}
	task, err := FetchTask(id)
	if err != nil {
		log.Printf("error fetching task with id %d: %v\n", id, err)
		return
	}
	Tmpl.ExecuteTemplate(w, "Item", map[string]any{"Item": task, "Editing": true})
}

func HandleUpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("error parsing id into int: %v\n", err)
		return
	}
	title := r.FormValue("title")
	if title == "" {
		log.Printf("error title must not empty")
		return
	}
	task, err := UpdateTask(id, title)
	if err != nil {
		log.Printf("error update task with id %d: %v\n", id, err)
		return
	}
	Tmpl.ExecuteTemplate(w, "Item", map[string]any{"Item": task})
}

func HandleOrderTasks(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("error parsing form: %v\n", err)
	}
	var values []int
	for k, v := range r.Form {
		if k == "item" {
			for _, v := range v {
				value, err := strconv.Atoi(v)
				if err != nil {
					log.Printf("error parsing id into int: %v\n", err)
					return
				}
				values = append(values, value)
			}
		}
	}
	if err := OrderTasks(r.Context(), values); err != nil {
		log.Printf("error ordering tasks: %v\n", err)
	}
}
