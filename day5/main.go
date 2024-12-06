package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func importData(filename string) ([][]int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var rules [][]int
	var pages [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			continue
		}

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		rules = append(rules, []int{num1, num2})

	}

	// Second section
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var numbers []int
		parts := strings.Split(line, ",")

		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				continue
			}
			numbers = append(numbers, num)
		}

		if len(numbers) > 0 {
			pages = append(pages, numbers)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error in Scan:", err)
	}

	return rules, pages, nil
}
func getRules(rules [][]int, pages []int) [][]int {
	var relevantRules [][]int
	for _, page := range pages {
		fmt.Println("Page:", page)
		for _, rule := range rules {
			if rule[0] == page && slices.Contains(pages, rule[1]) {
				relevantRules = append(relevantRules, rule)
			}
		}
	}
	return relevantRules
}

// func parseData(rules [][]int, pages [][]int) {
// 	for idx, pages := range pages {
// 		relRules := getRules(rules, pages)

// 	}
// }

func checkCorrect(relRule [][]int, pages []int) bool {
	indexMap := make(map[int]int)
	for idx, page := range pages {
		indexMap[page] = idx
	}

	for _, rule := range relRule {
		if indexMap[rule[0]] > indexMap[rule[1]] {
			return false
		}
	}
	return true
}

// func getIndex(relRule [][]int, pages []int) {
// 	sortedList := make([]int, len(pages))
// 	for _, page := range pages {
// 		count := 0
// 		for _, rule := range relRule {
// 			if rule[0] == page {
// 				count += 1
// 			}
// 		}
// 		sortedList[count] = page
// 	}
// 	fmt.Println("Sorted:", sortedList)

// }

func splitCorrect(rules [][]int, pages [][]int) ([][]int, [][]int) {
	var correctLists [][]int
	var incorrectLists [][]int
	for _, page := range pages {
		relRules := getRules(rules, page)
		if checkCorrect(relRules, page) {
			correctLists = append(correctLists, page)
		} else {
			incorrectLists = append(incorrectLists, page)
		}
	}
	return correctLists, incorrectLists
}

func sigmaMiddleVal(correctLists [][]int) int {
	var sum int
	for _, cor := range correctLists {
		mid := len(cor) / 2
		sum += cor[mid]
	}
	return sum
}

func checkRule(relRule []int, pages []int) bool {
	indexMap := make(map[int]int)
	for idx, page := range pages {
		indexMap[page] = idx
	}
	if indexMap[relRule[0]] > indexMap[relRule[1]] {
		return false
	}
	return true
}

func changeIndex(copyList []int, rule []int) []int {
	for idx, num := range copyList {
		if num == rule[0] {
			destValue := copyList[idx-1]
			copyList[idx-1] = copyList[idx]
			copyList[idx] = destValue
		}
	}
	return copyList
}

func sortList(incorrectLists [][]int, rules [][]int) [][]int {
	var sortedList [][]int
	for _, page := range incorrectLists {
		relRules := getRules(rules, page)
		dst := make([]int, len(page))
		copy(dst, page)
		for _, rule := range relRules {
			for !checkRule(rule, dst) {
				dst = changeIndex(dst, rule)
			}
		}
		sortedList = append(sortedList, dst)
	}
	return sortedList
}

func main() {
	rules, pages, err := importData("input.txt")
	if err != nil {
		panic(err)
	}

	correctLists, incorrectLists := splitCorrect(rules, pages)
	sum := sigmaMiddleVal(correctLists)
	fmt.Println("Correct Ordered Sum:", sum)
	sortedList := sortList(incorrectLists, rules)
	sumSort := sigmaMiddleVal(sortedList)
	fmt.Println("Sorted Sum:", sumSort)
}
