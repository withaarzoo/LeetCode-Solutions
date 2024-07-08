// Function to find the winner of the game
func findTheWinner(n int, k int) int {
    // The josephus function returns the position in 0-indexed format
    // Adding 1 to convert it to 1-indexed format for the final result
    return josephus(n, k) + 1
}

// Helper function to solve the Josephus problem
func josephus(n int, k int) int {
    // Base case: when there is only one person, they are the winner
    if n == 1 {
        return 0
    }
    // Recursive case:
    // - Solve the problem for (n-1) people
    // - Adjust the position by adding k (the step count)
    // - Use modulo n to wrap around if the position exceeds the number of people
    return (josephus(n-1, k) + k) % n
}
