package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func importData(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var inputLists [][]string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for _, v := range line {
			row = append(row, string(v))
		}
		inputLists = append(inputLists, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return inputLists
}

func findGuard(input [][]string) []int {
	location := make([]int, 2)
	for idx, row := range input {
		if slices.Contains(row, "^") {
			for j, elem := range row {
				if elem == "^" {
					location[0] = idx
					location[1] = j
				}
			}
		}
	}
	return location
}

func isMoveBlocked(currLoc []int, direction string, input [][]string) bool {
	switch direction {
	case "up":
		if input[currLoc[0]-1][currLoc[1]] == "#" {
			return true
		}
	case "down":
		if input[currLoc[0]+1][currLoc[1]] == "#" {
			return true
		}
	case "right":
		if input[currLoc[0]][currLoc[1]+1] == "#" {
			return true
		}
	case "left":
		if input[currLoc[0]][currLoc[1]-1] == "#" {
			return true
		}
	}
	return false
}

func moveGuard(currLoc []int, direction string, input [][]string) string {
	switch direction {
	case "up":
		if isMoveBlocked(currLoc, direction, input) {
			direction = "right"
		} else {
			input[currLoc[0]][currLoc[1]] = "X"
			currLoc[0] -= 1
		}
	case "down":
		if isMoveBlocked(currLoc, direction, input) {
			direction = "left"
		} else {
			input[currLoc[0]][currLoc[1]] = "X"
			currLoc[0] += 1
		}
	case "right":
		if isMoveBlocked(currLoc, direction, input) {
			direction = "down"
		} else {
			input[currLoc[0]][currLoc[1]] = "X"
			currLoc[1] += 1
		}
	case "left":
		if isMoveBlocked(currLoc, direction, input) {
			direction = "up"
		} else {
			input[currLoc[0]][currLoc[1]] = "X"
			currLoc[1] -= 1
		}
	}
	return direction
}

func motionGuard(input [][]string, startLoc []int) {
	direction := "up"
	status := "on"
	currLoc := make([]int, len(startLoc))
	copy(currLoc, startLoc)
	for status == "on" {
		if currLoc[0] == 0 || currLoc[0] == len(input)-1 || currLoc[1] == 0 || currLoc[1] == len(input[0])-1 {
			status = "off"
		} else {
			direction = moveGuard(currLoc, direction, input)
			fmt.Println("CurrLoc:", currLoc)
			fmt.Println("Direction:", direction)
		}
	}
}

func countXs(input [][]string) int {
	numX := 0
	for _, row := range input {
		for _, elem := range row {
			if elem == "X" {
				numX += 1
			} else {
				continue
			}
		}
	}
	// Add one for last position
	numX += 1
	return numX
}

func main() {
	data := importData("input.txt")
	guardLoc := findGuard(data)
	motionGuard(data, guardLoc)
	countX := countXs(data)
	fmt.Println("Total Count Xs:", countX)

}
