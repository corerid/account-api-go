package main
import (
    "encoding/json"
    "github.com/gorilla/mux"
    "fmt"
    "net/http"
    _ "github.com/lib/pq"
    "strconv"
)

// Our first handler
func accountFindById(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    //fmt.Println("HELLO ", params["id"])
    repos := repositories{}

	err := queryReposById(&repos, params["id"])
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

// Our first json response handler
func accounts(w http.ResponseWriter, r *http.Request) {
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

func insert(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    
    var data repositorySummary
    err := decoder.Decode(&data)
    if err != nil {
		panic(err)
    }
    fmt.Println(data.ID, data.Name, data.Owner, data.TotalStars)

    _, err = db.Exec(
        "INSERT INTO account VALUES ($1, $2, $3, $4)", 
        data.ID, 
        data.Name,  
        data.Owner, 
        data.TotalStars,
    )
    if err != nil {
        http.Error(w, err.Error(), 500)
        return 
    }

    out, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "Deleted successfully!\n" + string(out))
//	fmt.Fprintf(w, string(out))
}

func update(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data repositorySummary

    err := decoder.Decode(&data)
    if err != nil {
		panic(err)
    }
    fmt.Println(data.ID, data.Name, data.Owner, data.TotalStars)

    _, err = db.Exec(
		`UPDATE account
		SET id = $1, name = $2, owner = $3, totalstar = $4
        WHERE id = $1`, data.ID, data.Name, data.Owner, data.TotalStars,
	)
	if err != nil {
        http.Error(w, err.Error(), 500)
        return 
    }

    out, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	fmt.Fprintf(w, "Updated successfully!\n" + string(out))

    // json.NewEncoder(w).Encode(BasicResponse{
    //     200,
    //     "Updated Successfully!",
    // })

}

func delete(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    _, err := db.Exec(
        `DELETE FROM account
        WHERE id = $1`, params["id"],
    )
    if err != nil {
        http.Error(w, err.Error(), 500)
        json.NewEncoder(w).Encode(BasicResponse{
            0,
            err.Error(),
        })
        return 
    }

    json.NewEncoder(w).Encode(BasicResponse{
        200,
        "Deleted Successfully!",
    })

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
		return err
	}
	defer rows.Close()
	for rows.Next() {
		repo := repositorySummary{}
		err = rows.Scan(
			&repo.ID,
			&repo.Name,
			&repo.Owner,
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

func queryReposById(repos *repositories, id string) error {
    idInt, err := strconv.Atoi(id)

	rows, err := db.Query(`
		SELECT
			id,
			name,
			owner,
			totalstar
        FROM account
        WHERE id = $1`, idInt)

	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		repo := repositorySummary{}
		err = rows.Scan(
			&repo.ID,
			&repo.Name,
			&repo.Owner,
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