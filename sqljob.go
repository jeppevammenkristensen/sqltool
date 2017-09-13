package main

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/lib/pq"
)

type analyzer struct {
	conn *connection
}

type connection struct {
	db *sql.DB
}

func initiateconnection(connectionstring string, database string) (*connection, error) {
	db, err := sql.Open(database, connectionstring)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return &connection{db}, nil
}

type SqlJob struct {
	query      string
	connection *connection
}

func (s SqlJob) Hey() {
	return
}

func NewJob(query string, connectionstring string, database string) SqlJob {
	conn, err := initiateconnection(connectionstring, database)
	if err != nil {
		log.Panicln("Failed to establish connection", err)
	}

	log.Print("Initated connection to database")

	return SqlJob{query, conn}
}

func createJobFromFlags() SqlJob {
	var connectionstring, database, query string

	flag.StringVar(&connectionstring, "c", "", "The connectionstring to use")
	flag.StringVar(&database, "d", "postgres", "The database to use.")
	flag.StringVar(&query, "q", "", "The query to use")

	flag.Parse()

	return NewJob(query, connectionstring, database)
}
