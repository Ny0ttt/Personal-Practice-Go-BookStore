// package main

// import(
// 	"log"
// 	"net/http"
// 	"github.com/gorilla/mux"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	"github.com/Ny0ttt/Go-Bookstore/pkg/routes"
// )

// func main(){
// 	r := mux.NewRouter()
// 	routes.RegisterBookStoreRoutes(r)
// 	http.Handle("/", r)
// 	log.Fatal(http.ListenAndServe(":8080", r)) // creating a server
// }

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" // _ at the front is for use in side effects. it is imported but not used
	
	//"github.com/Ny0ttt/Go-Bookstore/pkg/routes"

	"github.com/Ny0ttt/Go-Bookstore/pkg/controllers"
)

func main() {
	r := mux.NewRouter()
	//routes.RegisterBookStoreRoutes(r)
	r.HandleFunc("/book", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":2222", r)) // creating a server
}

/*	MYSQL NOTE Part 1 
	for ListenAndServe, then the port number, it does not require you to use the same port like the usual ones like 3306 (this will give you the error 'An attempt was made to access a socket in a way forbidden by its access permissions. even if you grant access'),  
	you may use any other port number. you dont need to make a connection in mysql workbench to make a port. 

*/

/*	Ny0t's code NOTE 
	a slightly change on the file architecture is because it shows an error if you trying to pass r, which is the mux.NewRouter() to another file
	the error looks like 'cannot pass the variable 'vendor/github.com/gorilla/mux'.Router as 'github.com/gorilla/mux'.Router in the _____ argument
	i am still not sure why i cannot do it. multiple checking has been made like:-
	check the type using refelct(typeof(r)) 
	make sure import mux using go get github.com/gorilla/mux

	at the end, i just compile it in a single file
*/