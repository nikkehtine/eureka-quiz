package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/unrolled/render"
)

type Quiz struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var (
	rend *render.Render = render.New()
	db   *pgx.Conn
)

const PORT = ":2137"

func index(w http.ResponseWriter, r *http.Request) {
	var greeting string
	err := db.QueryRow(context.Background(), "select 'Hello Eureka!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("QueryRow failed: %v", err)
	}
	fmt.Println(greeting)
	w.Write([]byte(greeting))
}

func getQuizzes(w http.ResponseWriter, r *http.Request) {
	quizzesList := []Quiz{
		{ID: 1, Title: "Quiz 1"},
		{ID: 2, Title: "Quiz 2"},
		{ID: 3, Title: "Quiz 3"},
	}
	if err := rend.JSON(w, http.StatusOK, quizzesList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close(context.Background())

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", index)
	r.Get("/api/quizzes", getQuizzes)

	log.Printf("Server started on %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}
