package common

import (
	"io/ioutil"
	"strings"
)

// GetInput reads in the puzzle input file and returns an
// array of each line as a string.
func GetInput(filename string) []string {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(input), "\n")
}
