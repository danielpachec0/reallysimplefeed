package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func pingHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("request received in /ping")
	_, err := writer.Write([]byte("pong"))
	if err != nil {
		log.Println("Error when writing response")
	}
}

func randomHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("request received in /random")
	v := rand.Int()
	s := strconv.Itoa(v)
	_, err := writer.Write([]byte(s))
	if err != nil {
		log.Println("Error when writing response")
	}
}

func testHtmlTemplate(writer http.ResponseWriter, request *http.Request) {
	tpl, err := template.ParseFiles("assets/test.gohtml")
	if err != nil {
		log.Println("Error when ..")
	}
	err = tpl.Execute(writer, struct {
		Head string
		Body string
	}{
		Head: "TEST",
		Body: "test",
	},
	)
	if err != nil {
		return
	}
}

func main() {
	fmt.Println("start")
	port := "8080"

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/test", testHtmlTemplate)
	http.HandleFunc("/random", randomHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server could not be started")
	}
}
