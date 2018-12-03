package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var freq int64

func main() {
	freqs := make(map[int64]bool)
	var firstRepeat int64

	for ok := true; ok; ok = true {
		firstRepeat = calculateFrequencies(freqs)
		if firstRepeat != 0 {
			log.Println(firstRepeat)
			return
		}
	}

	return
}

func calculateFrequencies(freqs map[int64]bool) int64 {
	input := getInput()

	for _, i := range input {
		num, err := strconv.ParseInt(i, 10, 32)
		if err != nil {
			panic(err)
		}
		freq += num

		if _, ok := freqs[freq]; ok {
			return freq
		}

		freqs[freq] = true
	}
	return 0
}

func getInput() []string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(input), "\n")
}
