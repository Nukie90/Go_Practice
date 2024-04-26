package main

import (
	"database/sql"
	_ "fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nukie90/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	dbURL := os.Getenv("db_url")
	if dbURL == "" {
		log.Fatal("db_url environment variable not set")
	}

	connection, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error opening connection to database: ", err)
	}
	
	apiConfig := apiConfig{
		DB: database.New(connection),
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
	v1Router.Get("/healthz", readinessHandler)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiConfig.handlerCreateUser)
	v1Router.Get("/users", apiConfig.middlewareAuth(apiConfig.handlerGetUser))
	v1Router.Post("/feeds", apiConfig.middlewareAuth(apiConfig.handlerCreateFeeds))
	v1Router.Get("/feeds", apiConfig.handlerGetFeeds)
	v1Router.Post("/feedfollows", apiConfig.middlewareAuth(apiConfig.handlerCreateFeedFollow))
	v1Router.Get("/feedfollows", apiConfig.middlewareAuth(apiConfig.handlerGetFeedFollows))
	v1Router.Delete("/feedfollows/{feedFollowID}", apiConfig.middlewareAuth(apiConfig.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)
		
	srv := &http.Server{
		Handler : router,
		Addr : ":" + portString,
	}

	log.Println("Server is running on port", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}