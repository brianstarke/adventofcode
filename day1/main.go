package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := getInput()

	var freq int64

	for _, i := range input {
		num, err := strconv.ParseInt(i, 10, 32)
		if err != nil {
			panic(err)
		}
		freq += num
	}

	log.Println(freq)
	return
}

func getInput() []string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(input), "\n")
}
