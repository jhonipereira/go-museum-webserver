package main

import (
	"fmt"
	"jhonidev/go/goWebServer/api"
	"jhonidev/go/goWebServer/data"
	"net/http"
	"text/template"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go. HOMEPAGE"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error loading the template"))
		return
	}

	html.Execute(w, data.GetAll())
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/home", handleHome)

	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3334", server)
	if err != nil {
		fmt.Println("Couldn't start the server")
	}
}
