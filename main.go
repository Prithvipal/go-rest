package main

import (
	"io"
	"net/http"

	"github.com/Prithvipal/go-rest/controller"
	"github.com/gorilla/mux"
)

func h2(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello from a HandleFunc #2!\n")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/todo", controller.CreateTodo).Methods("POST")
	r.HandleFunc("/api/v1/todo", controller.GetTodoList).Methods("GET")
	r.HandleFunc("/api/v1/todo/{ID}", controller.GetTodoByID).Methods("GET")
	r.HandleFunc("/api/v1/todo/{ID}/detail", controller.GetTodoDetailByID).Methods("GET")
	r.HandleFunc("/api/v1/todo/{ID}", controller.DeleteTodoByID).Methods("DELETE")
	r.HandleFunc("/api/v1/todo/{ID}", controller.UpdateTodoByID).Methods("PUT")
	http.ListenAndServe(":8080", r)
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
