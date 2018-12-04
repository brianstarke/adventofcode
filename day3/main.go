package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input []string
var fabric [1000][1000]int

func main() {
	// First create the fabric...
	fabric = [1000][1000]int{}

	for _, claim := range input {
		c := parseClaim(claim)
		placeClaim(c.ID, c.X, c.Y, c.Width, c.Height)
	}

	// printFabric()
	fmt.Println(determineOverlap())

	for _, claim := range input {
		c := parseClaim(claim)
		o := determineIntact(c)

		if o == 0 {
			fmt.Printf("%d - %d\n", c.ID, o)
			break
		}
	}
}

func determineIntact(c *Claim) int {
	overlap := 0

	for i := c.X; i < c.Width+c.X; i++ {
		for j := c.Y; j < c.Height+c.Y; j++ {
			if fabric[j][i] == -1 {
				overlap++
			}
		}
	}
	return overlap
}

func determineOverlap() int {
	overlap := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabric[i][j] == -1 {
				overlap++
			}
		}
	}
	return overlap
}

// Claim gets parsed out of the input data
type Claim struct {
	ID     int
	X      int
	Y      int
	Width  int
	Height int
}

// #1 @ 1,3: 4x4
func parseClaim(s string) *Claim {
	var c Claim

	var buf []byte

	for i := 0; i < len([]byte(s)); i++ {
		x := string([]byte(s)[i])

		switch x {
		case "#":
			// start collecting ID field bytes
			buf = []byte{}
			continue
		case "@":
			// flush buffer to ID field and start collecting
			// next field
			i, err := strconv.ParseInt(string(buf), 10, 32)
			if err != nil {
				panic(err)
			}
			c.ID = int(i)
			buf = []byte{}
			continue
		case ",":
			i, err := strconv.ParseInt(string(buf), 10, 32)
			if err != nil {
				panic(err)
			}
			c.X = int(i)
			buf = []byte{}
			continue
		case ":":
			i, err := strconv.ParseInt(string(buf), 10, 32)
			if err != nil {
				panic(err)
			}
			c.Y = int(i)
			buf = []byte{}
			continue
		case "x":
			i, err := strconv.ParseInt(string(buf), 10, 32)
			if err != nil {
				panic(err)
			}
			c.Width = int(i)
			buf = []byte{}
			continue
		}

		// Don't buffer whitespace
		if string([]byte(s)[i]) == " " {
			continue
		}

		buf = append(buf, []byte(s)[i])
	}

	// Grab the last value (height)
	i, err := strconv.ParseInt(string(buf), 10, 32)
	if err != nil {
		panic(err)
	}
	c.Height = int(i)

	return &c
}

// This determines what space on the fabric is taken up by this claim and
// tags the area with the claim ID
func placeClaim(id, x, y, width, height int) {
	// Start from the top left and fill in
	for i := x; i < width+x; i++ {
		for j := y; j < height+y; j++ {
			// Check collision
			if fabric[j][i] != 0 {
				fabric[j][i] = -1
			} else {
				fabric[j][i] = id
			}
		}
	}
}

// Test condition laid out in the problem description...
func testOverlap() {
	placeClaim(1, 1, 3, 4, 4)
	placeClaim(2, 3, 1, 4, 4)
	placeClaim(3, 5, 5, 2, 2)
	printFabric()

	// [0 0 0 0 0 0 0 0 0 0]
	// [0 0 0 2 2 2 2 0 0 0]
	// [0 0 0 2 2 2 2 0 0 0]
	// [0 1 1 -1 -1 2 2 0 0 0]
	// [0 1 1 -1 -1 2 2 0 0 0]
	// [0 1 1 1 1 3 3 0 0 0]
	// [0 1 1 1 1 3 3 0 0 0]
	// [0 0 0 0 0 0 0 0 0 0]
	// [0 0 0 0 0 0 0 0 0 0]
	// [0 0 0 0 0 0 0 0 0 0]
}

// Just for visualizing... eg..
// [0 0 0 0 0 0 0 0 0 0]
// [0 0 0 0 0 0 0 0 0 0]
// [0 0 0 0 0 0 0 0 0 0]
// [0 0 0 0 0 0 0 0 0 0]
// [0 0 0 0 0 0 0 0 0 0]
// [0 0 0 0 0 0 0 0 0 0]
// [0 0 0 0 0 0 0 0 0 0]
// [0 0 0 0 0 0 0 0 0 0]
// [0 0 0 0 0 0 0 0 0 0]
// [0 0 0 0 0 0 0 0 0 0]
func printFabric() {
	fmt.Println("")
	for _, row := range fabric {
		fmt.Printf("\t%v\n", row)
	}
}

func init() {
	input = getInput()
}

func getInput() []string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(input), "\n")
}
