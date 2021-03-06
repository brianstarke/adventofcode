package main

import (
	"log"

	"github.com/brianstarke/adventofcode/common"
)

func main() {
	input := common.GetInput("input.txt")

	twos, threes := 0, 0

	for _, i := range input {
		two, three := analyze(i)
		if two {
			twos++
		}
		if three {
			threes++
		}
	}

	log.Println(twos)
	log.Println(threes)
	log.Println(twos * threes)
}

func analyze(i string) (bool, bool) {
	two, three := false, false

	m := toCountsMap(i)

	for _, v := range m {
		if v == 2 {
			two = true
		}
		if v == 3 {
			three = true
		}
	}

	return two, three
}

func toCountsMap(i string) map[string]int {
	res := make(map[string]int)

	for _, c := range []byte(i) {
		if _, ok := res[string(c)]; !ok {
			res[string(c)] = 0
		}
		res[string(c)]++
	}

	return res
}
