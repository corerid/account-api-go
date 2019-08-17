package main
import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    _ "github.com/lib/pq"
)

var db *sql.DB

const (
    dbhost = "XXX.XXX.XXX.XXX"
    dbport = 5432
    dbuser = "postgres"
    dbpass = "password1234"
    dbname = "postgres"
)

// init main func
func main() {
    initDb()
    defer db.Close()
    router := NewRouter()
    log.Fatal(
        // start on port 3000 by default
        http.ListenAndServe(":3000", router),
    )
}

func initDb() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",dbhost, dbport, dbuser, dbpass, dbname)
	var err error
    db, err = sql.Open("postgres", psqlInfo)
  	if err != nil {
    	panic(err)
  	}
      //defer db.Close()

  	err = db.Ping()
  	if err != nil {
    	panic(err)
  	}
    fmt.Println("Successfully connected!")
}