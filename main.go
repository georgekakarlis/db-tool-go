package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"tastetable.com/recipe-db/database"
)




func main() {

	db := database.ConnectDB()
	defer db.Close()
	fmt.Println("Tables created successfully")
	
	// router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// routes - empty for now
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello API!"))
	})
	r.Get("/api/getrecipes", func (w http.ResponseWriter, r *http.Request)  {
		
	})
	r.Post("/api/addrecipe", func (w http.ResponseWriter, r *http.Request)  {
		
	})


	port := os.Getenv("PORT")
	if port == "" {
    port = "8080" 
	}
	
	fmt.Println("Server is running in localhost:8080")
	err := http.ListenAndServe(":" + port , r)
	if err != nil {
    log.Fatalf("Error starting server: %v", err)
	}
}

