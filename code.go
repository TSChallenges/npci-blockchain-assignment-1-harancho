package main

import (
	"errors"
	"fmt"
	"sort"
)

// FilterAndSort filters numbers greater than or equal to the threshold and sorts them.
func FilterAndSort(nums []int, threshold int) []int {
	// TODO: Implement this function
	sort.Ints(nums)

	var index int
	for key, value := range nums {
		if value >= threshold {
			index = key
			break
		}
	}
	return nums[index:]
}

// FindMostFrequent finds the most frequent word in a slice of strings.
// Returns an error if the slice is empty.
func FindMostFrequent(words []string) (string, error) {
	// TODO: Implement this function
	var newMap = map[string]int{}
	if len(words) == 0 {
		return "", errors.New("not implemented")
	}
	
	for _, value := range words {
		newMap[value]++
	}

	maxCount :=0
	ans := ""

	for key, value := range newMap {
		if value > maxCount {
			maxCount = value
			ans = key
		}
	}

	return ans, nil
}

func main() {
	// Test FilterAndSort
	fmt.Println("Testing FilterAndSort:")
	nums := []int{3, 10, 1, 7, 8, 2}
	threshold := 5
	fmt.Printf("Input: %v, Threshold: %d, Output: %v\n", nums, threshold, FilterAndSort(nums, threshold))

	// Test FindMostFrequent
	fmt.Println("\nTesting FindMostFrequent:")
	words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	result, err := FindMostFrequent(words)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Input: %v, Most Frequent: %s\n", words, result)
	}

	// Edge case: empty slice
	emptyWords := []string{}
	result, err = FindMostFrequent(emptyWords)
	if err != nil {
		fmt.Printf("Error: %v (empty input case handled)\n", err)
	} else {
		fmt.Printf("Input: %v, Most Frequent: %s\n", emptyWords, result)
	}
}
