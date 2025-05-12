package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbm *sql.DB

func connectDB(){
	db, err := sql.Open("mysql","root:root@tcp(localhost:3308)/practicedb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	
	dbm = db
}

// func GetTables(){
// 	query := `SHOW TABLES`
// 	result,err := dbm.Exec(query)
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	fmt.Println(result)
// }

func GetTables() {
	rows, err := dbm.Query(`SHOW TABLES`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tableName string
	fmt.Println("Tables in the database:")
	for rows.Next() {
		err := rows.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("- " + tableName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func InsertInto(id int, name, email, city string, createdAt time.Time) {
	stmt := `INSERT INTO CUSTOMERS (customer_id, name, email, city, created_at) VALUES (?, ?, ?, ?, ?)`
	_, err := dbm.Exec(stmt, id, name, email, city, createdAt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User inserted successfully.")
}

func UpdateUserEmail(id int, newEmail string) {
	_, err := dbm.Exec(`UPDATE CUSTOMERS SET email = ? WHERE customer_id = ?`, newEmail, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated user email.")
}

func DeleteUser(id int) {
	_, err := dbm.Exec(`DELETE FROM CUSTOMERS WHERE customer_id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted user.")
}


