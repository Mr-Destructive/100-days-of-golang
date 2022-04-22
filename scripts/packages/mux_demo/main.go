package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", Server)

	http.ListenAndServe(":8000", router)
}

func Server(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello Mux!"))
}
