func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
    // Initialize the result slice to store the coordinates
    result := [][]int{}

    // Define the four possible movement directions: right, down, left, and up
    directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

    // Initialize the number of steps to take in the current direction
    steps := 1

    // Initialize the direction index to start moving right
    d := 0

    // Initialize the starting position
    r, c := rStart, cStart

    // Add the starting position to the result
    result = append(result, []int{r, c})

    // Loop until all cells in the matrix are visited
    for len(result) < rows*cols {
        // For each direction, repeat the movement twice (except for the first time)
        for i := 0; i < 2; i++ {
            // Move 'steps' times in the current direction
            for j := 0; j < steps; j++ {
                // Update the current position based on the current direction
                r += directions[d][0]
                c += directions[d][1]
                
                // Check if the new position is within the matrix boundaries
                if r >= 0 && r < rows && c >= 0 && c < cols {
                    // Add the valid position to the result
                    result = append(result, []int{r, c})
                }
            }
            // Change direction to the next one (right -> down -> left -> up)
            d = (d + 1) % 4
        }
        // Increase the number of steps after completing a full cycle of directions
        steps++
    }

    // Return the list of coordinates in the spiral order
    return result
}
