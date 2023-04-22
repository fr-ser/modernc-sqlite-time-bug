package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		fmt.Printf("Error opening the database: %v\n", err)
		return
	}
	var (
		utcTime   string
		localTime string
	)

	rows, err := db.Query("select datetime('now','utc') as UTC, datetime('now','localtime') as LOCAL")
	if err != nil {
		fmt.Printf("Error running the query: %v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&utcTime, &localTime)
		if err != nil {
			fmt.Printf("Error scanning the rows: %v\n", err)
			return
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Printf("Error loading the query result: %v\n", err)
		return
	}

	fmt.Printf("UTC: %s - Local: %s\n", utcTime, localTime)
}
