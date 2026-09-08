package main

import (
	"fmt"
	"github.com/a-h/templ"
	"html/template"
	"net/http"
)

func handleRequestHome(w http.ResponseWriter, r *http.Request) {
	templ := `
	<a href="/hello" onclick="navigate(event)">Hello</a>
    <a href="/new" onclick="navigate(event)">New</a>
	<a href="/username" onclick="navigate(event)">Edit</a>`
	parsed := template.Must(template.ParseFiles("index.html"))
	parsed.Execute(w, template.HTML(templ))
}

func main() {

	http.HandleFunc("/", handleRequestHome)

	http.Handle("/new", templ.Handler(new()))

	http.Handle("GET /username", templ.Handler(usernameEdit()))
	http.HandleFunc("POST /username", handleRequestUpdateUserName)

	http.ListenAndServe(":8080", nil)

	fmt.Println("Server is running on port 8080")
}
