package routes

import(
	"github.com/gorilla/mux"
	"github.com/Ny0ttt/Go-Bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}

//need to read more about mux.router and http
