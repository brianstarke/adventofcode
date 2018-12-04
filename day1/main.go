package main

import (
	"log"
	"strconv"

	"github.com/brianstarke/adventofcode/common"
)

func main() {
	input := common.GetInput("input.txt")

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
