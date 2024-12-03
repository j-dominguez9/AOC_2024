package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func importData(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	st := string(data)
	return st
}

func getActiveString(input string) string {
	fmt.Println("Str head:", input[:500])
	donts := strings.Split(input, "don't()")
	// donts := strings.SplitAfter(input, "don't()")

	fmt.Println("Donts len:", len(donts))
	fmt.Println("Donts:", donts[0])
	var dosList []string
	dosList = append(dosList, donts[0])
	for _, val := range donts[1:] {
		if strings.Contains(val, "do()") {
			dos := strings.Split(val, "do()")
			// fmt.Println("Do", dos)
			dosList = append(dosList, dos[1:]...)

		}
	}

	fmt.Println("Dos Len:", len(dosList))
	fmt.Println("Do", dosList[0])

	finalString := strings.Join(dosList, "_")
	return finalString
}

func cleanString(input string) []string {
	pattern := `mul\(\d+,\s*\d+\)`

	regex := regexp.MustCompile(pattern)

	match := regex.FindAllString(input, -1)

	return match
}

func cleanElem(input string) []int {
	pattern := `\d+`
	regex := regexp.MustCompile(pattern)

	matches := regex.FindAllString(input, -1)
	numbers := make([]int, len(matches))
	for i, match := range matches {
		numbers[i], _ = strconv.Atoi(match)
	}

	return numbers
}

func mulElem(input []int) int {
	product := input[0] * input[1]
	return product
}

func main() {

	data := importData("input.txt")
	procData := getActiveString(data)
	fmt.Println("Proc Data:", procData[:5])
	cleanData := cleanString(procData)
	fmt.Println(len(cleanData))
	var total int = 0
	for _, val := range cleanData {
		nums := cleanElem(val)
		product := mulElem(nums)
		total += product
	}
	fmt.Println("Total:", total)

}
