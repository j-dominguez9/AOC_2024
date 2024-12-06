package main

import (
	"bufio"
	"fmt"
	"os"
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

func parseData(data [][]string) int {
	ans := 0
	word := "XMAS"
	for x := range len(data) {
		for y := range len(data[0]) {
			if data[x][y] == "X" {
				result := dfs(data, 0, word, x, y)
				ans += result
			}
		}
	}
	return ans
}

func parseData2(data [][]string) int {
	ans := 0

	for x := range len(data) {
		for y := range len(data[0]) {
			if data[x][y] == "A" {
				result := helper(data, x, y)

				ans += result
			}
		}
	}
	return ans
}

func helper(data [][]string, x int, y int) int {
	if inBounds(data, x-1, y-1) && data[x-1][y-1] == "S" && inBounds(data, x-1, y+1) && data[x-1][y+1] == "S" && inBounds(data, x+1, y-1) && data[x+1][y-1] == "M" && inBounds(data, x+1, y+1) && data[x+1][y+1] == "M" {
		return 1
	}
	if inBounds(data, x-1, y-1) && data[x-1][y-1] == "S" && inBounds(data, x-1, y+1) && data[x-1][y+1] == "M" && inBounds(data, x+1, y-1) && data[x+1][y-1] == "S" && inBounds(data, x+1, y+1) && data[x+1][y+1] == "M" {
		return 1
	}
	if inBounds(data, x-1, y-1) && data[x-1][y-1] == "M" && inBounds(data, x-1, y+1) && data[x-1][y+1] == "M" && inBounds(data, x+1, y-1) && data[x+1][y-1] == "S" && inBounds(data, x+1, y+1) && data[x+1][y+1] == "S" {
		return 1
	}
	if inBounds(data, x-1, y-1) && data[x-1][y-1] == "M" && inBounds(data, x-1, y+1) && data[x-1][y+1] == "S" && inBounds(data, x+1, y-1) && data[x+1][y-1] == "M" && inBounds(data, x+1, y+1) && data[x+1][y+1] == "S" {
		return 1
	}
	return 0

}

func inBounds(data [][]string, x, y int) bool {
	if x < 0 || x >= len(data) || y < 0 || y >= len(data[0]) {
		return false
	}
	return true
}

func dfs(data [][]string, currIndex int, word string, x int, y int) int {
	direction := []string{
		"T", "TR", "R", "BR", "B", "BL", "L", "TL",
	}

	ans := 0
	for _, dir := range direction {
		result := dfsHelper(data, currIndex, word, dir, x, y)
		ans += result
	}
	return ans
}

func dfsHelper(data [][]string, currIndex int, word string, direction string, x int, y int) int {

	//went past the length of last word so we're good
	if currIndex >= len(word) {
		return 1
	}

	// if we out of bounds, in grid, return
	if x < 0 || x >= len(data) || y < 0 || y >= len(data[0]) {
		return 0
	}

	if data[x][y] != string(word[currIndex]) {
		return 0
	}

	switch direction {
	case "T":
		return dfsHelper(data, currIndex+1, word, direction, x-1, y)

	case "TR":
		return dfsHelper(data, currIndex+1, word, direction, x-1, y+1)

	case "R":
		return dfsHelper(data, currIndex+1, word, direction, x, y+1)

	case "BR":
		return dfsHelper(data, currIndex+1, word, direction, x+1, y+1)

	case "B":
		return dfsHelper(data, currIndex+1, word, direction, x+1, y)

	case "BL":
		return dfsHelper(data, currIndex+1, word, direction, x+1, y-1)

	case "L":
		return dfsHelper(data, currIndex+1, word, direction, x, y-1)

	default:
		return dfsHelper(data, currIndex+1, word, direction, x-1, y-1)
	}
}

// func checkCondS(input [][]string) {

// 	ans :=
// 	word := "XMAS"
// 	for x := range len(data) {

// 	}
// 	//horizontal
// 	for _, row := range input {
// 		fmt.Println(row)
// 		arr := strings.Split(row[0], "")
// 		for j, letter := range arr {
// 			fmt.Println(letter)
// 			if letter == "S" {
// 				if j != len(arr) && arr[j+1] == "A" {
// 					fmt.Println("yay")
// 				}
// 			}
// 		}
// 	}
// }

func main() {
	data := importData("input.txt")
	ans := parseData2(data)
	fmt.Println("Answer:", ans)
	// checkCondS(data)
}
