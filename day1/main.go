package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func importData(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())

		if len(numbers) == 2 {
			num1, err1 := strconv.Atoi(numbers[0])
			num2, err2 := strconv.Atoi(numbers[1])

			if err1 == nil && err2 == nil {
				list1 = append(list1, num1)
				list2 = append(list2, num2)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return list1, list2
}

// sorts lists
func sortLists(list1 []int, list2 []int) {
	slices.Sort(list1)
	slices.Sort(list2)
}

// Abs returns the absolute value of x.
func absVal(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func listDiff(list1 []int, list2 []int) []int {
	deltaList := make([]int, len(list1))
	for idx, val := range list1 {
		// diff := val - list2[idx]
		deltaList[idx] = absVal(val - list2[idx])
	}
	return deltaList
}

func addElements(diffList []int) int {
	var sigma int
	for _, val := range diffList {
		sigma += val
	}
	return sigma
}

func uniqueInts(list1 []int) []int {
	result := []int{}
	seen := make(map[int]bool)

	for _, val := range list1 {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}

func findCommon(list1 []int, list2 []int) map[int]int {
	counts := make(map[int]int)
	for _, val := range list1 {
		i := 0
		for _, value := range list2 {
			if val == value {
				i += 1
			}
		}
		counts[val] = i
	}
	return counts
}

func simScore(list1 []int, list2 []int) int {
	// distinct := uniqueInts(list1)
	i := 0
	counter := findCommon(list1, list2)
	for key, val := range counter {
		product := key * val
		i += product
	}
	return i
}

func main() {
	list1, list2 := importData("input.txt")
	// list1 := []int{3, 4, 2, 1, 3, 3}
	// list2 := []int{4, 3, 5, 3, 9, 3}
	sortLists(list1, list2)
	deltas := listDiff(list1, list2)
	sigma := addElements(deltas)
	fmt.Println("Sigma:", sigma)
	score := simScore(list1, list2)
	fmt.Println("Sim Score:", score)

}
