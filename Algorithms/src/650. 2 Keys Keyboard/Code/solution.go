package main

import "fmt"

// minSteps calculates the minimum number of steps required to get 'n' 'A's on the screen
// starting with a single 'A' and using only two operations: copy all and paste.
func minSteps(n int) int {
    // Initialize the number of operations required
    operations := 0

    // Start checking from the smallest possible factor, i.e., 2
    for i := 2; i <= n; i++ {
        // While 'n' is divisible by 'i', reduce 'n' by dividing it by 'i'
        // This simulates the process of applying the operation 'paste' multiple times
        // Each time we divide 'n' by 'i', it indicates that we've applied a sequence of operations
        // equivalent to copying the existing sequence of 'A's and pasting it 'i' times.
        for n % i == 0 {
            // Add the factor 'i' to the operation count
            // This is because to fully build 'n', we needed to paste 'i' times, which corresponds
            // to performing 'i' operations.
            operations += i

            // Reduce 'n' by dividing it by 'i'
            n /= i
        }
    }

    // Return the total number of operations calculated
    return operations
}

func main() {
    fmt.Println(minSteps(18)) // Example usage: should return 8
}
