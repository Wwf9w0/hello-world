package main

import (
	"html/template"
	"net/http"
)

type ToDo struct {
	Title string
	Done  bool
}
type ToDoPageData struct {
	PageTitle string
	Todos     []ToDo
}

func main() {
	tmpl := template.Must(template.ParseFiles("template.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ToDoPageData{
			PageTitle: "My TODO list",
			Todos: []ToDo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8091", nil)
}
