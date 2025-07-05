package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang/glog"
	_ "github.com/lib/pq"
)

var db *sql.DB
var Pg_host string
var Pg_db string

func Open() {
	var err error
	pg_user, u := os.LookupEnv("pg_user")
	pg_pass, p := os.LookupEnv("pg_pass")
	if !u || !p {
		glog.Infof("Environment variables pg_user and/or pg_pass weren't defined")
		os.Exit(1)
	}

	connStr := "host=" + Pg_host + " user=" + pg_user + " password=" + pg_pass + " dbname=" + Pg_db + " sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func Close() {
	db.Close()
}

func GetRidByHostname(hostname string) (id int, rid string) {
	id = 0
	rid = "x"
	rows, err := db.Query("select node.id, node.ip from node where node.hostname like '" + hostname + "'")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		switch err := rows.Scan(&id, &rid); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			return id, rid
		default:
		}
	}
	rows.Close()
	return id, rid
}
