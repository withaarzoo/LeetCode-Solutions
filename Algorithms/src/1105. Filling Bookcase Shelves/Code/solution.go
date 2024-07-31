func minHeightShelves(books [][]int, shelfWidth int) int {
    // Number of books
    n := len(books)
    
    // dp[i] will store the minimum height required to place the first i books
    dp := make([]int, n + 1)
    
    // Initialize dp array with a large value (infinity equivalent)
    for i := range dp {
        dp[i] = 1<<31 - 1
    }
    dp[0] = 0 // No books placed means height is 0

    // Iterate over each book to determine the minimum height needed to place them
    for i := 1; i <= n; i++ {
        width := 0  // Current shelf width used
        height := 0 // Current maximum height of books in the shelf

        // Iterate backwards from the current book to place as many as possible in the current row
        for j := i; j > 0; j-- {
            // Add the width of the current book
            width += books[j - 1][0]
            
            // If adding the current book exceeds the shelf width, stop adding books
            if width > shelfWidth {
                break
            }
            
            // Update the maximum height of the books in the current shelf
            if height < books[j - 1][1] {
                height = books[j - 1][1]
            }
            
            // Update dp[i] to the minimum height by considering the height
            // of the current row plus the best configuration of the previous books
            if dp[i] > dp[j - 1] + height {
                dp[i] = dp[j - 1] + height
            }
        }
    }

    // The last element of dp gives the minimum height required to place all books
    return dp[n]
}
