package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	// Sort the array to facilitate duplicate management and early stopping
	sort.Ints(candidates)
	
	// This will store all the unique combinations that sum to the target
	var result [][]int
	
	// This slice will hold the current combination being explored
	var current []int
	
	// Define the backtracking function that will recursively explore combinations
	var backtrack func(target, start int)
	backtrack = func(target, start int) {
		// Base case: If the target is exactly zero, we found a valid combination
		if target == 0 {
			// Add a copy of the current combination to the result
			result = append(result, append([]int{}, current...))
			return
		}
		
		// Iterate over the candidates starting from 'start'
		for i := start; i < len(candidates); i++ {
			// Skip duplicate elements to avoid duplicate combinations
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			
			// If the current candidate is greater than the remaining target, stop further exploration
			if candidates[i] > target {
				break
			}
			
			// Include the current candidate in the combination
			current = append(current, candidates[i])
			
			// Recursively explore further with the reduced target and the next candidate
			backtrack(target-candidates[i], i+1)
			
			// Backtrack: remove the last candidate added to explore other possibilities
			current = current[:len(current)-1]
		}
	}
	
	// Start the backtracking with the initial target and starting index 0
	backtrack(target, 0)
	
	// Return the final list of unique combinations
	return result
}

func main() {
	// Example usage
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	result := combinationSum2(candidates, target)
	fmt.Println(result) // Expected output: [[1 1 6] [1 2 5] [1 7] [2 6]]
}
