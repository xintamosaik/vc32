package main

import (
	"fmt"
	"html/template"
	"net/http"
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

func handleRequestHello(w http.ResponseWriter, r *http.Request) {
	templ := "<h1>Hello, {{.}}</h1>" + homeLink

	parsed := template.Must(template.New("hello").Parse(templ))
	parsed.Execute(w, name)
}

func handleRequestNew(w http.ResponseWriter, r *http.Request) {
	templ := "<h1>New</h1>" + homeLink
	parsed := template.Must(template.New("new").Parse(templ))
	parsed.Execute(w, nil)
}

func handleRequestUpdateUserName(w http.ResponseWriter, r *http.Request) {
	// check if the request method is POST
	if r.Method == http.MethodPost {
		r.ParseForm()
		newName := r.FormValue("name")
		if newName != "" {
			name = newName
		}
		http.Redirect(w, r, "/hello", http.StatusSeeOther)
	}
}

func handleRequestEditUserName(w http.ResponseWriter, r *http.Request) {
	templ := `
	<h1>Change Name</h1>
	<form action="/edit-name" method="POST">
	<label for="username">Name</label>
		<input name="username" id="username" type="text" /> 
	</form>
	` + homeLink

	parsed, err := template.New("hello").Parse(templ)
	if (err !=nil) {
		fmt.Fprint(w, "no")
	}
	parsed.Execute(w, name)
}

func main() {
	name = "John Doe"

	http.HandleFunc("/", handleRequestHome)
	http.HandleFunc("/hello", handleRequestHello)
	http.HandleFunc("/new", handleRequestNew)
	http.HandleFunc("GET /username", handleRequestEditUserName)
	http.HandleFunc("POST /username", handleRequestUpdateUserName)

	http.ListenAndServe(":8080", nil)

	fmt.Println("Server is running on port 8080")
}
