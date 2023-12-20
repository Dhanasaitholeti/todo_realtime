package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Todo struct{
	Id int `json "Id"`
	Title string `json "Title"`
	Done bool `json "Done"`
	Body string `json "Body"`
}

func main() {

	app := chi.NewRouter()
	app.Use(middleware.Logger)


	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Dhanasai Tholeti"))
	})

	app.Post("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		body,err := ioutil.ReadAll(r.Body)
		if err!=nil{
			log.Fatal("error occured lol:)")
		}
		log.Printf("Received request body: %s",body)
	})

	port := ":8080"

	go func() {
		log.Printf("Server is listening on port %s...\n", port)
	}()


	err := http.ListenAndServe(port, app)
	if err != nil {
		log.Fatal(err)
	}
	
}