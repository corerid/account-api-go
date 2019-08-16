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
    dbhost = "localhost"
    dbport = 5432
    dbuser = "postgres"
    dbpass = "password1234"
    dbname = "go-test"
)

func main() {
    initDb()
    defer db.Close()
    http.HandleFunc("/api/index", indexHandler)
    http.HandleFunc("/api/repo/", repoHandler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	repos := repositories{}

// 	err := queryRepos(&repos)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	out, err := json.Marshal(repos)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	fmt.Fprintf(w, string(out))
// }

// func queryRepos(repos *repositories) error {
// 	rows, err := db.Query(`
// 		SELECT
// 			id,
// 			name,
// 			owner,
// 			totalstar
// 		FROM account`)
// 	if err != nil {
// 		fmt.Println("HELLO!!!!!!!!!!!!!!!!!!!!!!!")
// 		return err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		repo := repositorySummary{}
// 		err = rows.Scan(
// 			&repo.ID,
// 			&repo.Owner,
// 			&repo.Name,
// 			&repo.TotalStars,
// 		)
// 		if err != nil {
// 			return err
// 		}
// 		repos.Repositories = append(repos.Repositories, repo)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func repoHandler(w http.ResponseWriter, r *http.Request) {
    //...
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