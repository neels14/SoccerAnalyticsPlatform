package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	id   int    `json:"id"`
	name string `json:"name"`
}

const testDBName string = "testdb"
const prodDBName string = "proddb"

func main() {
	// Open up our database connection.
	// The database is called testdb
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/"+testDBName)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// CREATE table test
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test ( id integer, name varchar(32) )")
	if err != nil {
		panic(err.Error())
	}

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO test VALUES ( 1, 'Neel' ), ( 2, 'Anu' ), ( 3, 'Aashrit') , ( 4, 'Hamza' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	// CHECK DATA EXISTS
	// perform a db.Query
	results, err := db.Query("SELECT * FROM test")
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("ID", "NAME")
	for results.Next() {
		var data Data
		// for each row, scan the result into our tag composite object
		err = results.Scan(&data.id, &data.name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(data.id, data.name)
	}

}
