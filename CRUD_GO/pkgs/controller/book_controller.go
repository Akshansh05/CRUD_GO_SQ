package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cast"

	model "../model"
	utils "../utils"
	"github.com/gorilla/mux"
)

type BookJson struct {
	Name  string      `json:"name"`
	About Description `json:"about"`
}
type Description struct {
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

//converts JSON to sql table type
func ConvertBookJsonToBookObject(bookJson *BookJson) (book model.Book) { //return must be object or address of pointer(i.e &object)
	book.Name = bookJson.Name
	book.Author = bookJson.About.Author
	book.Publication = bookJson.About.Publication
	return book
}

//converts sql table type to JSON
func ConvertBookObjectToBookJson(book *model.Book) (bookJson BookJson) {
	bookJson.Name = book.Name
	bookJson.About.Author = book.Author
	bookJson.About.Publication = book.Publication
	return bookJson
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	bookJson := &BookJson{} //create a book object

	utils.ParseBody(r, bookJson) //parse according to the interface passed

	//to get request string
	requestString, _ := json.Marshal(bookJson)
	fmt.Println(string(requestString)) //  parse the request to the BookJson object //we must parse the json to object

	CreateBook := ConvertBookJsonToBookObject(bookJson) // convert Book Json to required Db

	b := model.CreateBook(&CreateBook) // insert in db and return the book object

	ConvertedbookJson := ConvertBookObjectToBookJson(b) // convert the book object to book json

	res, _ := json.Marshal(ConvertedbookJson) // Unparse the b object back to JSON response ..get the response as BookJSON object

	w.WriteHeader(http.StatusOK) //give 200 status if all fine
	w.Write(res)                 // write the response back as response
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := model.GetAllBooks()

	res, _ := json.Marshal(newBooks) //get response an book object not BookJson object

	w.Header().Set("Content-Type", "pkglication/json") //returns the Response in proper JSON beautified
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails := model.GetBookById(cast.ToInt(ID))

	res, _ := json.Marshal(&bookDetails) //get response an book object not BookJson object

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookJson := &BookJson{}
	utils.ParseBody(r, bookJson) //we must parse the json to object

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails := model.GetBookById(cast.ToInt(ID))

	BookFormBookDetails := ConvertBookJsonToBookObject(bookJson)

	if BookFormBookDetails.Name != "" {
		bookDetails.Name = BookFormBookDetails.Name
	}
	if BookFormBookDetails.Author != "" {
		bookDetails.Author = BookFormBookDetails.Author
	}
	if BookFormBookDetails.Publication != "" {
		bookDetails.Publication = BookFormBookDetails.Publication
	}

	book := model.UpdateBookByObject(bookDetails)

	JsonFormBookDetails := ConvertBookObjectToBookJson(book)

	res, _ := json.Marshal(JsonFormBookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	book := model.DeleteBookById(cast.ToInt(ID))

	res, _ := json.Marshal(&book) //get response an book object not BookJson object

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//kon sa book user ne rakha h
func GetBooksByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookUserId := vars["bookUserId"]
	ID, err := strconv.ParseInt(bookUserId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := model.GetBooksByUserID(cast.ToInt(ID))
	res, _ := json.Marshal(&book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
