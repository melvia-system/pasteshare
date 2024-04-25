package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type BaseResponse struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

type DocumentItem struct {
	Id        string    `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"createdat"`
}

var db *sqlx.DB

func main() {
	log.Println("🚀 PasteShare api server starting...")

	// env
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("RPC_PORT")
	if port == "" {
		port = "3001"
	}
	log.Println("🔌 Environment loaded ✅")

	// db
	datasource := "user=" + os.Getenv("DB_USER") + " dbname=" + os.Getenv("DB_NAME") + " password=" + os.Getenv("DB_PASSWORD") + " sslmode=disable"
	log.Println("🔌 Database connecting... (" + datasource + ")")
	db, err = sqlx.Connect("postgres", datasource)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Println("📚 Database connected ✅")

	// Create a new Router
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
	r.Use(middleware.Timeout(60 * time.Second))

	// Routes
	// GET /
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(BaseResponse{Message: "PasteShare api server", Ok: true})
	})
	// POST /doc
	r.Post("/doc", DocSave)
	// GET /doc/{id}
	r.Get("/doc/{id}", DocGet)

	// Start the server
	log.Println("🌐 http server start on :" + port + " ✅")
	http.ListenAndServe(":"+port, r)
}

func DocSave(w http.ResponseWriter, r *http.Request) {
	// get body have json { title: string, content: string }
	var doc DocumentItem
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// save to db and get id of document
	err = db.QueryRow("INSERT INTO documents (title, content) VALUES ($1, $2) RETURNING id", doc.Title, doc.Content).Scan(&doc.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if doc.Id == "" {
		http.Error(w, "Document not saved", http.StatusBadRequest)
		return
	}

	// response json, merge with BaseResponse struct and { data: DocumentItem }
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		BaseResponse
		Data DocumentItem
	}{
		BaseResponse: BaseResponse{Message: "Document saved", Ok: true},
		Data:         doc,
	})
}

func DocGet(w http.ResponseWriter, r *http.Request) {
	docId := chi.URLParam(r, "id")
	log.Println("🔍 Get document id: " + docId)

	// check docId is empty
	if docId == "" {
		http.Error(w, "Document id is required", http.StatusBadRequest)
		return
	}

	// get document from db
	var doc DocumentItem
	err := db.Get(&doc, "SELECT * FROM documents WHERE Id = $1", docId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// response json, merge with BaseResponse struct and { data: DocumentItem }
	json.NewEncoder(w).Encode(struct {
		BaseResponse
		Data DocumentItem
	}{
		BaseResponse: BaseResponse{Message: "Document found", Ok: true},
		Data:         doc,
	})
}
