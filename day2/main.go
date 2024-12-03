package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func importData(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var inputLists [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by whitespace
		numStrings := strings.Fields(line)

		numbers := make([]int, len(numStrings))
		for i, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("Error converting string to number: %v\n", err)
				continue
			}
			numbers[i] = num
		}

		// Append the slice of nums to inputLists
		inputLists = append(inputLists, numbers)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return inputLists
}

func isDesc(inputList []int) bool {
	var prevInt int = inputList[0]
	for _, val := range inputList[1:] {
		if val < prevInt {
			diff := getDiff(val, prevInt)
			prevInt = val
			if diff > 3 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
func isAsc(inputList []int) bool {
	var prevInt int = inputList[0]
	for _, val := range inputList[1:] {
		if val > prevInt {
			diff := getDiff(val, prevInt)
			prevInt = val
			if diff > 3 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// Abs returns the absolute value of x.
func absVal(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getDiff(x int, y int) int {
	sub := x - y
	diff := absVal(sub)
	return diff
}

func main() {
	// var inputLists [][]int = [][]int{
	// 	{7, 6, 4, 2, 1},
	// 	{1, 2, 7, 8, 9},
	// 	{9, 7, 6, 2, 1},
	// 	{1, 3, 2, 4, 5},
	// 	{1, 3, 6, 7, 9},
	// }
	inputLists := importData("input.txt")
	i := 0
	for _, val := range inputLists {
		if isDesc(val) || isAsc(val) {
			i += 1
		} else {
			for idx, _ := range val {
				copyDest := make([]int, len(val))
				copy(copyDest, val)
				copyDest = append(copyDest[:idx], copyDest[idx+1:]...)
				if isDesc(copyDest) || isAsc(copyDest) {
					i += 1
					break
				}
			}
		}

	}
	fmt.Println("Total Safe:", i)

}
