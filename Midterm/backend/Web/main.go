package main

import (
	"backend/Web/apihelper"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users", apihelper.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/api/users", apihelper.AddUserHandler).Methods("POST")
	router.HandleFunc("/api/users/{id}", apihelper.GetUserHandler).Methods("GET")
	router.HandleFunc("/api/users/{id}", apihelper.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/api/users/{id}", apihelper.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/api/users/login", apihelper.LoginHandler).Methods("POST")

	corsObj := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)
	http.ListenAndServe(":8080", corsObj(router))
}
