package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if len(os.Args) < 3 {
		println("Syntax: go run main.go <dbname> <table>")
		return
	}
	dbname := os.Args[1]
	table := os.Args[2]

	fmt.Println("MySQL V1")
	dt := time.Now().Format("15:04")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Println("Database Connected [Time: " + dt + "]")
	}
	defer db.Close()

	insert, err := db.Query("CREATE DATABASE " + dbname + ";")
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Println("Created Database: " + dbname + " [Time: " + dt + "]")
	}
	db2, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		log.Println(err.Error())
	} else {
	}
	defer db2.Close()

	intable, err := db2.Query("CREATE TABLE `" + table + "` (`Value` varchar(128) DEFAULT NULL)")

	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Println("Created Table: " + table + " [Time: " + dt + "]")
	}

	defer insert.Close()
	defer intable.Close()
}
