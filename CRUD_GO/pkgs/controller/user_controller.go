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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &model.User{}       // create user object
	utils.ParseBody(r, CreateUser)    //  parse the request to the user object not responseJson like previous case //we must parse the json to object
	u := model.CreateUser(CreateUser) // insert in db and return the user object
	res, _ := json.Marshal(&u)        // Unparse the b object back to JSON response
	w.WriteHeader(http.StatusOK)      //give 200 status if all fine
	w.Write(res)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	Users := model.GetAllUsers()
	res, _ := json.Marshal(Users)
	w.Header().Set("Content-Type", "pkglication/json") //returns the Response in proper JSON beautified
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	userDetails := model.GetUserById(cast.ToInt(ID))
	res, _ := json.Marshal(&userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &model.User{}
	utils.ParseBody(r, updateUser) //we must parse the json to object..if we want to change the user details from here we must parse to proper JSON..

	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	userDetails := model.GetUserById(cast.ToInt(ID))
	if updateUser.Name != "" {
		userDetails.Name = updateUser.Name
	}
	if updateUser.BookId != nil {
		userDetails.BookId = updateUser.BookId //changes Entire Book content .. we cannot change book parameters from here
	}
	user := model.UpdateUserByObject(userDetails)
	res, _ := json.Marshal(&user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	user := model.DeleteUserById(cast.ToInt(ID))
	res, _ := json.Marshal(&user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//particular book kis kis user ne le rakha hai
func GetUserByBookId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userBoookId := vars["userBookId"]
	ID, err := strconv.ParseInt(userBoookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	totalEntries := model.GetUserByBookId(cast.ToInt(ID)) //RelatedSel wont work here bcoz we are not using QueryTable

	var userBook []*model.User
	for _, index := range totalEntries {
		newUser := &model.User{}
		newUser.Id = cast.ToInt(index["id"])
		newUser.Name = cast.ToString(index["name"])
		newUser.BookId = model.GetBookById(cast.ToInt(index["bookId"])) //to get the values back again
		userBook = append(userBook, newUser)
	}
	res, _ := json.Marshal(userBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
