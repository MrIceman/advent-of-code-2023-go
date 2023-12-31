package main

import (
	"aoc-1/day1"
	"aoc-1/day2"
	_ "embed"
	"log"
	"strings"
)

//go:embed day1/input.txt
var input1 string

//go:embed day2/input.txt
var input2 string

func main() {
	runD2()
}

func runD2() {
	splitInput := strings.Split(input2, "\n")
	res1 := day2.GetResultPart1(splitInput)
	log.Printf("total sum id is %d", res1)

	res2 := day2.GetResultPart2(splitInput)
	log.Printf("power of sets is is %d", res2)
}

func runD1() {
	splitInput := strings.Split(input1, "\n")
	day1.CalculateSum(splitInput)
}
