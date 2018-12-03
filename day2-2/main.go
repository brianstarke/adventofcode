package main

import (
	"io/ioutil"
	"log"
	"strings"
)

var input []string

func init() {
	input = getInput()
}

func main() {
	var a, b string

	for _, i := range input {
		m := toDiffMap(i)

		for k, v := range m {
			if v == 1 {
				a = k
				b = i
				break
			}
		}
	}

	var res []byte

	// Produce the string with the different characters removed
	for n := 0; n < len(a); n++ {
		if []byte(a)[n] == []byte(b)[n] {
			res = append(res, []byte(a)[n])
		}
	}

	log.Println(a)
	log.Println(b)
	log.Println(string(res))
}

// returns a map containing the diff counts of this item vs all
// others
func toDiffMap(i string) map[string]int {
	res := make(map[string]int)

	for _, c := range input {
		res[c] = getDiffCount(i, c)
	}

	return res
}

func getDiffCount(a, b string) int {
	cnt := 0

	for i := 0; i < len(a); i++ {
		if []byte(a)[i] != []byte(b)[i] {
			cnt++
		}
	}

	return cnt
}

func getInput() []string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(input), "\n")
}
