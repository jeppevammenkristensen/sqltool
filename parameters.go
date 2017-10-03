package main

import "flag"

type jobParameters struct {
	connectionstring, database, query string
}

func createJobParametersFromFlags() jobParameters {
	var connectionstring, database, query string

	flag.StringVar(&connectionstring, "c", "", "The connectionstring to use")
	flag.StringVar(&database, "d", "postgres", "The database to use.")
	flag.StringVar(&query, "q", "", "The query to use")

	flag.Parse()

	return jobParameters{connectionstring, database, query}
}
