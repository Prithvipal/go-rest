package main

import (
	"io"
	"net/http"

	"github.com/Prithvipal/go-rest/handler"
)

func h2(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello from a HandleFunc #2!\n")
}
func main() {
	http.HandleFunc("/api/v1/todo", handler.CreateToDo)
	http.HandleFunc("/endpoint", h2)
	http.ListenAndServe(":8080", nil)
}

// import (
// 	"io"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	h1 := func(w http.ResponseWriter, _ *http.Request) {
// 		io.WriteString(w, "Hello from a HandleFunc #1!\n")
// 	}
// 	h2 := func(w http.ResponseWriter, _ *http.Request) {
// 		io.WriteString(w, "Hello from a HandleFunc #2!\n")
// 	}

// 	http.HandleFunc("/", h1)
// 	http.HandleFunc("/endpoint", h2)

// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
