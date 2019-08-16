package main

// import (
//   "database/sql"
//   "fmt"

//   _ "github.com/lib/pq"
// )

// const (
//     dbhost = "localhost"
//     dbport = 5432
//     dbuser = "postgres"
//     dbpass = "password1234"
//     dbname = "go-test"
// )

// func main() {
//     psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",dbhost, dbport, dbuser, dbpass, dbname)
// //   db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password1234 dbname=go-test sslmode=disable")
//     db, err := sql.Open("postgres", psqlInfo)
//   if err != nil {
//     panic(err)
//   }
//   defer db.Close()

//   err = db.Ping()
//   if err != nil {
//     panic(err)
//   }

//   fmt.Println("Successfully connected!")
// }