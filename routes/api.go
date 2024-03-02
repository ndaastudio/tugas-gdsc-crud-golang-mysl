package routes

import (
	UserController "crud-golang-mysql/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", UserController.GetUser).Methods("GET")
	router.HandleFunc("/api/user", UserController.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", UserController.GetUserByID).Methods("GET")
	router.HandleFunc("/api/user/{id}", UserController.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", UserController.DeleteUser).Methods("DELETE")

	return router
}
