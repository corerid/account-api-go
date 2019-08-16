package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type repositorySummary struct {
	ID         int
	Name       string
	Owner      string
	TotalStars int
}

type repositories struct {
	Repositories []repositorySummary
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	repos := repositories{}

	err := queryRepos(&repos)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(repos)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

func queryRepos(repos *repositories) error {
	rows, err := db.Query(`
		SELECT
			id,
			name,
			owner,
			totalstar
		FROM account`)
	if err != nil {
		fmt.Println("HELLO!!!!!!!!!!!!!!!!!!!!!!!")
		return err
	}
	defer rows.Close()
	for rows.Next() {
		repo := repositorySummary{}
		err = rows.Scan(
			&repo.ID,
			&repo.Owner,
			&repo.Name,
			&repo.TotalStars,
		)
		if err != nil {
			return err
		}
		repos.Repositories = append(repos.Repositories, repo)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
