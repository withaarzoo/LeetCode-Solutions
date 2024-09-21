package main

import "fmt"

// Main function to generate lexicographical numbers from 1 to n
func lexicalOrder(n int) []int {
    // Initialize an empty slice to store the result
    result := []int{}
    
    // Iterate over the first digit (1 to 9) since lexicographical order starts from these
    // We don't start with 0 because numbers in lexicographical order don't start with 0
    for i := 1; i <= 9; i++ {
        // Perform Depth-First Search (DFS) starting with the current digit
        dfs(i, n, &result)
    }
    
    // Return the final result containing numbers in lexicographical order
    return result
}

// Helper function to perform DFS and construct numbers in lexicographical order
// 'curr' represents the current number being constructed
// 'n' is the upper limit (we need numbers from 1 to n)
// 'result' is a pointer to the result slice to accumulate numbers
func dfs(curr, n int, result *[]int) {
    // Base case: if the current number exceeds n, stop the recursion
    if curr > n {
        return
    }
    
    // Add the current number to the result slice
    *result = append(*result, curr)
    
    // Now, we try to generate the next numbers by appending digits 0-9 to 'curr'
    // This simulates going deeper in the lexicographical sequence (e.g., from 1 -> 10, 11, 12, etc.)
    for i := 0; i <= 9; i++ {
        // Calculate the next number by appending digit 'i' to 'curr'
        next := curr * 10 + i
        
        // If the next number exceeds 'n', we stop further exploration from this path
        if next > n {
            return
        }
        
        // Recursive call to explore the next number in the sequence
        dfs(next, n, result)
    }
}

func main() {
    n := 13
    result := lexicalOrder(n)
    fmt.Println(result) // Output: [1 10 11 12 13 2 3 4 5 6 7 8 9]
}
