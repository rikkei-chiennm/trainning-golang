package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username = "chiennm"
	password = "chiennm0110@123"
	hostname = "127.0.0.1:3306"
	dbname   = "product"
)

type User struct {
	gorm.Model
	Name  string
	Email string
	Age   int
}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?&parseTime=True", username, password, hostname, dbname)
}

func main() {
	inforDB := dsn()
	db, err := gorm.Open(mysql.Open(inforDB), &gorm.Config{})

	if err != nil {
		log.Println("error connect!")
	}

	db.AutoMigrate(&User{})

	var users = []User{{Name: "Chien", Email: "chiennm0110@gmail.com", Age: 18}, {Name: "Chi", Email: "chi@gmail.com", Age: 32}, {Name: "Dung", Email: "dung@gmail.com", Age: 19}, {Name: "Hoang", Email: "hoang@gmail.com", Age: 18}}
	// create list user
	db.Create(&users)

	// get infor user with id = 1
	var user User
	resG := db.First(&user, 1)
	if resG.Error == nil {
		fmt.Println("user first DB :", user)
	} else {
		fmt.Println(resG.Error)
	}

	// // get all users
	var listUser []User
	resL := db.Find(&listUser)
	if resL.Error == nil {
		fmt.Println("List user in DB :", listUser)
	} else {
		fmt.Println(resL.Error)
	}

	// // update user with id = 4 , email = tu123@gmai.com

	var userUpdate User
	db.First(&userUpdate, 4)
	errU := db.Model(&userUpdate).Update("email", "chi123@gmai.com")
	if errU.Error == nil {
		fmt.Println("data after update :", userUpdate)
	}

	// delete user with id = 1
	db.Delete(&User{}, 1)
}
