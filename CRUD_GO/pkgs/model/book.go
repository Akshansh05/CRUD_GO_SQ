package model

import (
	"fmt"

	"../config"
	"github.com/astaxie/beego/orm"
)

type Book struct {
	Id          int    `orm:"column(id);auto" description:"Increment Id"`
	Name        string `orm:"column(Name);size(64);null"`
	Author      string `orm:"column(Author);size(64);null"`
	Publication string `orm:"column(Publication);size(64);null"`
}

const (
	Books = "Books"
)

func (t *Book) TableName() string {
	return Books
}

func init() {
	config.Connect()
	orm.RegisterModel(new(Book))
}

func CreateBook(b *Book) *Book {
	o := orm.NewOrm()
	_, err := o.Insert(b)
	if err != nil {
		panic(err)
	}
	return b
}

func GetAllBooks() (books []Book) {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Book)).All(&books)
	if err != nil {
		panic(err)
	}
	fmt.Println(books)
	return books
}
func GetBookById(id int) *Book {
	o := orm.NewOrm()
	getBook := Book{Id: id}
	err := o.QueryTable(new(Book)).Filter("id", id).One(&getBook)
	if err != nil {
		panic(err)
	}
	return &getBook

}

func UpdateBookByObject(updatedBook *Book) *Book {
	o := orm.NewOrm()
	oldUpdatedBook := Book{Id: updatedBook.Id}

	if err := o.Read(&oldUpdatedBook); err == nil { //get the book if exists
		var num int64
		if num, err = o.Update(updatedBook); err == nil { //updateBook here is itaslf a pointer
			fmt.Println("Number of records updated in database:", num)
		}
	}

	return updatedBook
}

func DeleteBookById(id int) *Book {
	o := orm.NewOrm()
	deleteBook := Book{Id: id}
	// ascertain id exists in the database
	if err := o.Read(&deleteBook); err == nil {
		var num int64
		if num, err = o.Delete(&deleteBook); err == nil { //deleteBook is object thats why we pass &&deleteBook
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return &deleteBook
}

//kon sa book given userId lia hua h
func GetBooksByUserID(id int) (books *Book) {
	o := orm.NewOrm()
	getuser := User{Id: id}
	// ascertain id exists in the database
	var err error
	if err = o.Read(&getuser); err == nil {
		_, errInQuery := o.Raw("SELECT * FROM Books b INNER JOIN Users u ON b.id = u.bookId WHERE u.id = ? ORDER BY b.id", getuser.Id).QueryRows(&books)
		if errInQuery != nil {
			panic(errInQuery)
		}
		return books
	}
	panic(err)
}
