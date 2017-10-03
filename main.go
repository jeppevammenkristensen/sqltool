package main

import (
	"log"
)

func main() {
	job := createJobFromFlags()

	result, err := job.Analyze()
	process(result, "Hejsa", "main")

	checkError(err)

}

// checkError checks if the pased in err is nil. If not it will
// it will fail
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
