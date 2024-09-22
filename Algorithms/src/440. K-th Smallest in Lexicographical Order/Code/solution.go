// Function to count how many numbers exist in the lexicographical range
// starting from 'curr' up to 'n'. This function helps us determine how many
// numbers exist within the current prefix level.
func countSteps(curr, n int) int {
    // Initialize the number of steps (i.e., the number of valid numbers within this range)
    steps := 0

    // The `first` and `last` represent the current range of numbers
    // at the same prefix level. Initially, both are set to 'curr'.
    first, last := curr, curr

    // Loop through all levels of numbers starting from `curr`.
    // The range keeps increasing by powers of 10 (moving deeper in the lexicographical tree).
    for first <= n {
        // Add the number of valid numbers at this prefix level.
        // It calculates the count of numbers between `first` and `last` but ensures not to exceed `n`.
        steps += min(n+1, last+1) - first

        // Move to the next level by multiplying by 10.
        // This means moving from numbers like `1, 10, 100` to deeper levels.
        first *= 10

        // Set `last` to the maximum possible number at this level by adding 9.
        // For example, if `last` was 1, now it becomes 19, then 199, and so on.
        last = last*10 + 9
    }

    // Return the total number of steps (valid numbers) found within the current prefix level.
    return steps
}

// Function to find the k-th smallest number in lexicographical order from 1 to n.
func findKthNumber(n int, k int) int {
    // Start with the current number set to 1 (since lexicographically, 1 is the smallest).
    curr := 1
    
    // Decrement k because we are using a zero-indexed approach.
    // In other words, if we are looking for the k-th number, it corresponds to the (k-1)-th index.
    k--

    // Loop until we find the k-th number
    for k > 0 {
        // Count how many numbers start with the current prefix (`curr`).
        steps := countSteps(curr, n)

        // If the number of numbers with the current prefix (`steps`) is less than or equal to `k`,
        // it means the k-th number is not within the current prefix.
        if steps <= k {
            // Move to the next prefix by incrementing `curr`.
            curr++
            
            // Since we are skipping over `steps` numbers, subtract `steps` from `k`.
            k -= steps
        } else {
            // If the k-th number is within the current prefix, we "dive" deeper into the lexicographical tree.
            // This is done by multiplying `curr` by 10 (e.g., going from 1 to 10).
            curr *= 10
            
            // Since we are diving into the next level, decrement `k` to reflect the movement.
            k--
        }
    }

    // Return the k-th number found.
    return curr
}

// Helper function to return the minimum of two integers.
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
