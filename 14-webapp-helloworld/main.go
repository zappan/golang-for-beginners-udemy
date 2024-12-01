package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const PORT = 8080


func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}

func wannabeHtmlPage(w http.ResponseWriter, r *http.Request) {
	html := `<strong>Hello, world</strong>`
	fmt.Fprintf(w, html)
}

func htmlPage(w http.ResponseWriter, r *http.Request) {
	html := `<strong>Hello, world</strong>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, html)
}

func automaticHtmlPage(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
	<body>
	  <h1>Hello, world</h1>
	</body>
</html>`

	// this one does NOT need an explicit header set
	// unclear why? maybe the fprintf checks for the html markup in the given string...
	fmt.Fprintf(w, html)
}

func staticPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}


func renderTemplate(w http.ResponseWriter, pathToPage string) {
	tpl, err := template.ParseFiles(pathToPage)
	if err != nil {
		log.Println(err)
		return
	}

	err = tpl.Execute(w, nil)   // nil: a viariable used to pass data to the template
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/wannabe", wannabeHtmlPage)
	http.HandleFunc("/html", htmlPage)
	http.HandleFunc("/page", automaticHtmlPage)
	http.HandleFunc("/static", staticPage)

	log.Println("Starting web server on port", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
