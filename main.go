package main

import (
	"log"
	"database/sql"
	_ "github.com/alexbrainman/odbc"
	)


func main() {
	dsnStr := "DSN=test;"

	// Replace the DSN value with the name of your ODBC data source.
	db, err := sql.Open("odbc", dsnStr)
	if err != nil {
		log.Fatal(err)
	}

	var (
		id int
		name string
	)

        // This is a Impala query.
	rows, err := db.Query("SELECT distinct id, name FROM User WHERE userId = ?", 7)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
		

}