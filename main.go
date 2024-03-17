package main

import (
	"database/sql"
	"fmt"
	"gotest/db"
)

func main() {
	fmt.Println("Hi, this is go")

	rows, err := db.Query("my_server_name", "user_name", "password", "db_name", 1433, "Select * FROM Countries")

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var (
			countryId   int
			countryName string
			currency    string
			isActive    bool
			createdOn   string
			modifiedBy  string
			modifiedOn  string
			appid       sql.NullString
		)

		if err := rows.Scan(&countryId, &countryName, &currency, &isActive, &createdOn, &modifiedBy, &modifiedOn, &appid); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Println("CountryID:", countryId, "CountryName:", countryName)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return
	}

}
