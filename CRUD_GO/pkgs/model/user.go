package model

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id     int    `orm:"column(id);auto" description:"Increment Id"`
	Name   string `orm:"column(name);size(64);null"`
	BookId *Book  `orm:"column(bookId);rel(fk)"` //iss user ke pass kon sa book hai
}

const (
	Users = "Users"
)

func (t *User) TableName() string {
	return Users
}

func init() {
	orm.RegisterModel(new(User))
}

func CreateUser(u *User) *User {
	o := orm.NewOrm()
	_, err := o.Insert(u) //whatever wrong payload we give if id of book and name is correct record willbe created according to name and book id content
	if err != nil {
		panic(err)
	}
	return u
}

func GetAllUsers() (users []User) {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(User)).RelatedSel().All(&users) //do relatedSel to get all Info
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
	return users
}

func GetUserById(id int) *User {
	o := orm.NewOrm()
	getuser := User{Id: id}
	err := o.QueryTable(new(User)).Filter("id", id).RelatedSel().One(&getuser) //RelatedSel can get u getuser.BookId.Id  visible
	if err != nil {
		panic(err)
	}
	return &getuser
}

func UpdateUserByObject(updatedUser *User) *User {
	o := orm.NewOrm()
	oldUpdatedUser := User{Id: updatedUser.Id}

	if err := o.Read(&oldUpdatedUser); err == nil { //get the user if exists
		var num int64
		if num, err = o.Update(updatedUser); err == nil { //updateUser here is itaslf a pointer
			fmt.Println("Number of records updated in database:", num)
		}
	}

	return updatedUser
}
func DeleteUserById(id int) *User {
	o := orm.NewOrm()
	deleteUser := User{Id: id}
	// ascertain id exists in the database
	if err := o.QueryTable(new(User)).Filter("id", id).RelatedSel().One(&deleteUser); err == nil {
		var num int64
		if num, err = o.Delete(&deleteUser); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return &deleteUser
}

//kon kon user ye book(given bookid) lia
func GetUserByBookId(id int) (v []orm.Params) { // to get different params   v []orm.Params//RelatedSel wont work here bcoz we are not using QueryTable
	o := orm.NewOrm()
	_, errInQuery := o.Raw("SELECT u.id as id,u.name as name,u.bookId as bookId FROM Users u INNER JOIN Books b ON u.bookId = b.id  WHERE b.id = ?", id).Values(&v) //QueryRows not used as we are not returning the table like user or book we are returning paramns
	if errInQuery != nil {
		panic(errInQuery)
	}
	return v
}
