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

func isMoveBlocked(currentPosition []int, direction string, input [][]string) bool {
	switch direction {
	case "up":
		if input[currentPosition[0]-1][currentPosition[1]] == "#" || input[currentPosition[0]-1][currentPosition[1]] == "O" {
			return true
		}
	case "down":
		if input[currentPosition[0]+1][currentPosition[1]] == "#" || input[currentPosition[0]-1][currentPosition[1]] == "O" {
			return true
		}
	case "right":
		if input[currentPosition[0]][currentPosition[1]+1] == "#" || input[currentPosition[0]-1][currentPosition[1]] == "O" {
			return true
		}
	case "left":
		if input[currentPosition[0]][currentPosition[1]-1] == "#" || input[currentPosition[0]-1][currentPosition[1]] == "O" {
			return true
		}
	}
	return false
}

func moveGuard(currentPosition []int, direction string, input [][]string) string {
	switch direction {
	case "up":
		if isMoveBlocked(currentPosition, direction, input) {
			direction = "right"
		} else {
			input[currentPosition[0]][currentPosition[1]] = "X"
			currentPosition[0] -= 1
		}
	case "down":
		if isMoveBlocked(currentPosition, direction, input) {
			direction = "left"
		} else {
			input[currentPosition[0]][currentPosition[1]] = "X"
			currentPosition[0] += 1
		}
	case "right":
		if isMoveBlocked(currentPosition, direction, input) {
			direction = "down"
		} else {
			input[currentPosition[0]][currentPosition[1]] = "X"
			currentPosition[1] += 1
		}
	case "left":
		if isMoveBlocked(currentPosition, direction, input) {
			direction = "up"
		} else {
			input[currentPosition[0]][currentPosition[1]] = "X"
			currentPosition[1] -= 1
		}
	}
	return direction
}

func motionGuard(input [][]string, startLoc []int) [][]int {
	direction := "up"
	status := "on"
	currentPosition := make([]int, len(startLoc))
	copy(currentPosition, startLoc)
	var xPositions [][]int
	for status == "on" {
		if currentPosition[0] == 0 || currentPosition[0] == len(input)-1 || currentPosition[1] == 0 || currentPosition[1] == len(input[0])-1 {
			// Mark X on last position
			input[currentPosition[0]][currentPosition[1]] = "X"
			status = "off"
		} else {
			direction = moveGuard(currentPosition, direction, input)
			fmt.Println("CurrLoc:", currentPosition)
			fmt.Println("Direction:", direction)
			positionCopy := make([]int, len(currentPosition))
			copy(positionCopy, currentPosition)
			xPositions = append(xPositions, positionCopy)
		}
	}
	return xPositions
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
	return numX
}

// Part 2
func isMoveBlockedPlus(currentPosition []int, direction string, input [][]string) bool {
	switch direction {
	case "up":
		if input[currentPosition[0]-1][currentPosition[1]] == "+" {
			return true
		}
	case "down":
		if input[currentPosition[0]+1][currentPosition[1]] == "+" {
			return true
		}
	case "right":
		if input[currentPosition[0]][currentPosition[1]+1] == "+" {
			return true
		}
	case "left":
		if input[currentPosition[0]][currentPosition[1]-1] == "+" {
			return true
		}
	}
	return false
}

var ErrHitPlus = fmt.Errorf("Hit plus sign, possible infinite loop")

func moveGuard2(currentPosition []int, direction string, input [][]string) (string, error) {
	switch direction {
	case "up":
		if isMoveBlocked(currentPosition, direction, input) {
			direction = "right"
			input[currentPosition[0]][currentPosition[1]] = "+"
		} else if isMoveBlockedPlus(currentPosition, direction, input) {
			return "up", ErrHitPlus
		} else {
			input[currentPosition[0]][currentPosition[1]] = "|"
			currentPosition[0] -= 1
		}
	case "down":
		if isMoveBlocked(currentPosition, direction, input) {
			direction = "left"
			input[currentPosition[0]][currentPosition[1]] = "+"
		} else if isMoveBlockedPlus(currentPosition, direction, input) {
			return "down", ErrHitPlus
		} else {
			input[currentPosition[0]][currentPosition[1]] = "|"
			currentPosition[0] += 1
		}
	case "right":
		if isMoveBlocked(currentPosition, direction, input) {
			direction = "down"
			input[currentPosition[0]][currentPosition[1]] = "+"
		} else if isMoveBlockedPlus(currentPosition, direction, input) {
			return "right", ErrHitPlus
		} else {
			input[currentPosition[0]][currentPosition[1]] = "-"
			currentPosition[1] += 1
		}
	case "left":
		if isMoveBlocked(currentPosition, direction, input) {
			direction = "up"
			input[currentPosition[0]][currentPosition[1]] = "+"
		} else if isMoveBlockedPlus(currentPosition, direction, input) {
			return "left", ErrHitPlus
		} else {
			input[currentPosition[0]][currentPosition[1]] = "-"
			currentPosition[1] -= 1
		}
	}
	return direction, nil
}

func isInfiniteLoop(direction string, input [][]string, startLoc []int) bool {
	status := "on"
	currentPosition := make([]int, len(startLoc))
	copy(currentPosition, startLoc)
	for status == "on" {
		if currentPosition[0] == 0 || currentPosition[0] == len(input)-1 || currentPosition[1] == 0 || currentPosition[1] == len(input[0])-1 {
			// Mark X on last position
			// input[currentPosition[0]][currentPosition[1]] = "X"
			status = "off"
			// return false

		} else {
			direction, err := moveGuard2(currentPosition, direction, input)
			if err != nil {
				return true
			}
			fmt.Println("CurrLoc:", currentPosition)
			fmt.Println("Direction:", direction)
		}
	}
	return false
}

func checkRoute(input [][]string, ogData [][]string, guardLoc []int, Xcoords [][]int) int {
	count := 0
	// fmt.Println(ogData)
	// fmt.Println(input)
	for _, coord := range Xcoords[1:] {
		fmt.Println("Coord tested:", coord)
		loopData := make([][]string, len(ogData))
		copy(loopData, ogData)
		// fmt.Println(ogData)
		loopData[coord[0]][coord[1]] = "O"
		if isInfiniteLoop("up", loopData, guardLoc) {
			fmt.Println("IL found.")
			count += 1
		} else {
			fmt.Println("No IL.")
			continue
		}
	}
	return count
}

func main() {
	data := importData("input.txt")
	ogData := make([][]string, len(data))
	copy(ogData, data)
	// start := time.Now()
	guardLoc := findGuard(data)
	Xcoords := motionGuard(data, guardLoc)
	// countX := countXs(data)
	// duration := time.Since(start)
	// fmt.Println("Time", duration)
	// fmt.Println("Total Count Xs:", countX)
	fmt.Println(Xcoords)
	count := checkRoute(data, ogData, guardLoc, Xcoords)
	fmt.Println("Infinite Loop Count:", count)

}
