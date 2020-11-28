package routes

import (
	controller "../controller"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	//books
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controller.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controller.DeleteBook).Methods("DELETE")

	//users
	router.HandleFunc("/user/", controller.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{userId}", controller.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controller.DeleteUser).Methods("DELETE")

	//joinRoutes
	router.HandleFunc("/userBook/{userBookId}", controller.GetUserByBookId).Methods("GET")
	router.HandleFunc("/bookUser/{bookUserId}", controller.GetBooksByUserId).Methods("GET")

}
