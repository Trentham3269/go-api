package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
)

var db *sql.DB

const (
    dbhost = "DBHOST"
    dbport = "DBPORT"
    dbuser = "DBUSER"
    dbpass = "DBPASS"
    dbname = "DBNAME"
)

func main() {
    initDb()
    defer db.Close()
    http.HandleFunc("/api/", apiHandler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type apiSummary struct {
	Name       string
	Year       string
	Location   string
}

type winners struct {
    Winners []apiSummary
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    wins := winners{}

	err := queryWinners(&wins)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(wins)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

// queryRepos first fetches the repositories data from the db
func queryWinners(wins *winners) error {
	rows, err := db.Query(`select name, year, location from honourboard where board = 'Commonwealth Games'`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		win := apiSummary{}
		err = rows.Scan(
			&win.Name,
			&win.Year,
			&win.Location,
		)
		if err != nil {
			return err
		}
		wins.Winners = append(wins.Winners, win)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func initDb() {
    config := dbConfig()
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        config[dbhost], config[dbport],
        config[dbuser], config[dbpass], config[dbname])

    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string {
    conf := make(map[string]string)
    host, ok := os.LookupEnv(dbhost)
    if !ok {
        panic("DBHOST environment variable required but not set")
    }
    port, ok := os.LookupEnv(dbport)
    if !ok {
        panic("DBPORT environment variable required but not set")
    }
    user, ok := os.LookupEnv(dbuser)
    if !ok {
        panic("DBUSER environment variable required but not set")
    }
    password, ok := os.LookupEnv(dbpass)
    if !ok {
        panic("DBPASS environment variable required but not set")
    }
    name, ok := os.LookupEnv(dbname)
    if !ok {
        panic("DBNAME environment variable required but not set")
    }
    conf[dbhost] = host
    conf[dbport] = port
    conf[dbuser] = user
    conf[dbpass] = password
    conf[dbname] = name
    return conf
}
