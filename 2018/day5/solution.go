package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

const asciiPolarityDistance = 32

func main() {
	polymer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	// Solution 1
	resultPolymer := processPolarity(polymer)
	fmt.Println(len(resultPolymer))


	// Solution 2
	polymer, err = ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	shortestLength := len(processPolarity(removeUnit(polymer, 97)))
	shortestUnit := 65
	for i := 66; i <= 90; i++ {
		unitsRemoved := removeUnit(polymer, i)
		reacted := processPolarity(unitsRemoved)
		thisLen := len(reacted)

		if thisLen < shortestLength {
			shortestLength = thisLen
			shortestUnit = i
		}
	}

	fmt.Println(string(shortestUnit))
	fmt.Println(shortestLength)
}

func removeUnit(polymer []byte, lowerByteVal int) []byte {
	units := []byte{
		byte(lowerByteVal),
		byte(lowerByteVal+ asciiPolarityDistance),
	}

	return bytes.Map(func(r rune) rune {
		if bytes.IndexRune(units, r) < 0 {
			return r
		}
		return -1
	}, polymer)
}

func processPolarity(polymer []byte) []byte {
	for i := 0; i < len(polymer) - 1; i++ {
		if Abs(int(polymer[i]) - int(polymer[i+1])) == asciiPolarityDistance {
			return processPolarity(append(polymer[:i], polymer[i+2:]...))
		}
	}

	return polymer
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
