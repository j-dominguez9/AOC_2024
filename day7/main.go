package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func import_data(filename string) map[int][]int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// var inputLists []map[int][]int
	inputMap := make(map[int][]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splitString := strings.Split(line, ": ")
		key := splitString[0]
		val := splitString[1]
		keyInt, err := strconv.Atoi(key)
		if err != nil {
			fmt.Println("Error converting string to int", err)
		}
		splitVals := strings.Split(val, " ")

		remVals := []int{}
		// inputMap := make(map[int][]int)
		for _, v := range splitVals {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Error converting string to in", err)
			}
			remVals = append(remVals, num)
		}
		fmt.Println("Key:", keyInt)
		fmt.Println("Val:", remVals)
		inputMap[keyInt] = remVals
		// inputLists = append(inputLists, inputMap)
	}
	fmt.Println(inputMap)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return inputMap

}

func genCombinations(values []int) [][]int {
	total := len(values) - 1
	list0 := make([]int, total)
	list1 := make([]int, total)

	for idx, _ := range list0 {
		list0[idx] = 0
		list1[idx] = 1
	}

	totalCombinations := 1 << total // 2^n combinations
	result := make([][]int, totalCombinations)

	for i := 0; i < totalCombinations; i++ {
		combination := make([]int, total)
		for j := 0; j < total; j++ {
			// Use bit manipulation to determine which slice to take from
			if (i & (1 << j)) != 0 {
				combination[j] = list0[j]
			} else {
				combination[j] = list1[j]
			}
		}
		result[i] = combination
	}

	return result

}

func checkOperation(combos [][]int, values []int, test int) bool {
	for _, set := range combos {
		total := values[0]
		for idx, val := range set {
			if val == 0 {
				total += values[idx+1]
			} else if val == 1 {
				total *= values[idx+1]
			}
			if total == test {
				return true
			}
		}
	}
	return false
}

func main() {
	data := import_data("input.txt")
	totalTestVal := 0
	for key, val := range data {
		combos := genCombinations(val)
		result := checkOperation(combos, val, key)
		fmt.Println(result)
		if result == true {
			totalTestVal += key
		}

	}
	fmt.Println(totalTestVal)

}
