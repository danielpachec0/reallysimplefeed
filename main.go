package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func error_respond(writer http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	} else if code > 399 {
		log.Println("Responding with 4XX error:", msg)
	} else if code > 299 {
		log.Println("Responding with 3XX error:", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	json_respond(writer, code, errorResponse{
		Error: msg,
	})
}
func json_respond(writer http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		writer.WriteHeader(500)
		return
	}
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(data)
}

func handlreReadinnes(writer http.ResponseWriter, request *http.Request) {
	json_respond(writer, 200, struct{}{})
}

func testErrorResponse(writer http.ResponseWriter, request *http.Request) {
	error_respond(writer, 500, "Server error test")
}

func main() {
	godotenv.Load()
	fmt.Println("start")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not provided in configuration")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost"},
		AllowedMethods:   []string{"http://localhost"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ping", handlreReadinnes)
	v1Router.Get("/error", testErrorResponse)

	router.Mount("/v1", v1Router)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Output(1, "Server starting at Port: "+port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("..")
	}
}
