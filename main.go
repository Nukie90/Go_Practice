package main

import (
	_ "fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	// V1 routes will only use GET method
	v1Router.Get("/healthz", readinessHandler)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)


		
	srv := &http.Server{
		Handler : router,
		Addr : ":" + portString,
	}

	log.Println("Server is running on port", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}