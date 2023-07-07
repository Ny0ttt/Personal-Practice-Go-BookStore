package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //not sure about the _
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:18181818@/go-bookstore?charset=utf8&parseTime=True&loc=Local")
	
	//NOTE MYSQL Part 2
	/*it seems like you need to connect to an existing account for mysql that was installed in the system. you also need to 
	connect to an existing database in order to make this work. checking database in your mysql system will require to use 
	the command promt 
	mysql -u username -p
	*type your password*
	show databases
	quit
	i use root as username when doing this project. 
	i made the go-bookstore database using mysql workbench, make connection and create schema
	*/




	//d, err := sql.Open("mysql", "host=localhost:3306 user=root password=18181818 dbname=bookstore sslmode=disable")
	//d, err := gorm.Open("mysql", "user=root password=18181818/bookstore sslmode=disable")
	//i dont have any gorm database in my system, idk what the user name and the password are
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
