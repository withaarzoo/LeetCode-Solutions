func numMagicSquaresInside(grid [][]int) int {
    // Get the number of rows and columns in the grid
    rows := len(grid)
    cols := len(grid[0])
    
    // Initialize a counter to keep track of the number of magic squares found
    count := 0

    // Function to check if the 3x3 subgrid starting at (r, c) is a magic square
    isMagicSquare := func(r, c int) bool {
        // Create a boolean array to track the numbers 1-9
        vals := make([]bool, 10)

        // Iterate through the 3x3 subgrid
        for i := 0; i < 3; i++ {
            for j := 0; j < 3; j++ {
                // Get the number at the current position
                num := grid[r+i][c+j]
                
                // Check if the number is within the valid range (1-9) and not repeated
                if num < 1 || num > 9 || vals[num] {
                    return false
                }
                
                // Mark the number as seen
                vals[num] = true
            }
        }

        // Check all rows, columns, and diagonals to ensure they sum to 15
        return grid[r][c] + grid[r][c+1] + grid[r][c+2] == 15 && // First row
               grid[r+1][c] + grid[r+1][c+1] + grid[r+1][c+2] == 15 && // Second row
               grid[r+2][c] + grid[r+2][c+1] + grid[r+2][c+2] == 15 && // Third row
               grid[r][c] + grid[r+1][c] + grid[r+2][c] == 15 && // First column
               grid[r][c+1] + grid[r+1][c+1] + grid[r+2][c+1] == 15 && // Second column
               grid[r][c+2] + grid[r+1][c+2] + grid[r+2][c+2] == 15 && // Third column
               grid[r][c] + grid[r+1][c+1] + grid[r+2][c+2] == 15 && // Main diagonal
               grid[r][c+2] + grid[r+1][c+1] + grid[r+2][c] == 15 // Anti-diagonal
    }

    // Loop through each possible 3x3 subgrid in the grid
    for i := 0; i < rows-2; i++ {
        for j := 0; j < cols-2; j++ {
            // If the subgrid is a magic square, increment the count
            if isMagicSquare(i, j) {
                count++
            }
        }
    }

    // Return the total number of magic squares found
    return count
}
