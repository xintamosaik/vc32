package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/a-h/templ"
)

var name string
const homeLink = `<a href="/">Home</a>`

func handleRequestHome(w http.ResponseWriter, r *http.Request) {
	templ := `
	<a href="/hello" onclick="navigate(event)">Hello</a>
    <a href="/new" onclick="navigate(event)">New</a>
	<a href="/username" onclick="navigate(event)">Edit</a>`
	parsed := template.Must(template.ParseFiles("index.html"))
	parsed.Execute(w, template.HTML(templ))
}

 
func handleRequestNew(w http.ResponseWriter, r *http.Request) {
	templ := "<h1>New</h1>" + homeLink
	parsed := template.Must(template.New("new").Parse(templ))
	parsed.Execute(w, nil)
}

 

func main() {
	name = "John Doe"

	http.HandleFunc("/", handleRequestHome)

	http.HandleFunc("/new", handleRequestNew)
	
	http.Handle("GET /username", templ.Handler(usernameEdit()))
	http.HandleFunc("POST /username", handleRequestUpdateUserName)

	http.ListenAndServe(":8080", nil)

	fmt.Println("Server is running on port 8080")
}
