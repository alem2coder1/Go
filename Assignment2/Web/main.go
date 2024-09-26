package main

import (
	"Assignment2/Web/apihelper"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users", apihelper.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/api/users", apihelper.AddUserHandler).Methods("POST")
	router.HandleFunc("/api/users/{id}", apihelper.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/api/users/{id}", apihelper.DeleteUserHandler).Methods("DELETE")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
