package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	job := createJobFromFlags()

	result, err := job.Analyze()
	spew.Dump(result)
	checkError(err)

}

// checkError checks if the pased in err is nil. If not it will
// it will fail
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
