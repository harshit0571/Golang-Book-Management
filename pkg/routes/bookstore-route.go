package routes

import (
	"github.com/gorilla/mux"
	"github.com/harshit0571/Golang-Book-Management/pkg/controllers"
)

var RegisterBookStoreRoutes = func (router * mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("PUT")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")
}