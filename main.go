package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var name string

func handleRequestHome(w http.ResponseWriter, r *http.Request) {
	templ := `
	<a href="/hello" onclick="navigate(event)">Hello</a>
    <a href="/new" onclick="navigate(event)">New</a>`
	parsed := template.Must(template.ParseFiles("index.html"))
	parsed.Execute(w, template.HTML(templ))
}

func handleRequestHello(w http.ResponseWriter, r *http.Request) {
	templ := "<h1>Hello, {{.}}</h1>"

	parsed := template.Must(template.New("hello").Parse(templ))
	parsed.Execute(w, name)
}

func handleRequestNew(w http.ResponseWriter, r *http.Request) {
	templ := "<h1>New</h1>"
	parsed := template.Must(template.New("new").Parse(templ))
	parsed.Execute(w, nil)
}
func main() {
	name = "John Doe"

	http.HandleFunc("/", handleRequestHome)
	http.HandleFunc("/hello", handleRequestHello)
	http.HandleFunc("/new", handleRequestNew)

	http.ListenAndServe(":8080", nil)

	fmt.Println("Server is running on port 8080")
}
