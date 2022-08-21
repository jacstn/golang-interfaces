package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Url struct {
	Id        int
	Name      string
	CreatedAt string
}

func main() {
	db := Connect()

	res, err := db.Query("SELECT * FROM domain")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var url Url
		err := res.Scan(&url.Id, &url.Name, &url.CreatedAt)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", url)
	}
	defer db.Close()
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/urldb")
	if err != nil {
		panic("Error connecting to db")
	}

	return db
}
