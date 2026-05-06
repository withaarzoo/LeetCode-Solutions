func rotateTheBox(boxGrid [][]byte) [][]byte {

    m := len(boxGrid)
    n := len(boxGrid[0])

    // Process every row
    for row := 0; row < m; row++ {

        // Rightmost empty position
        emptyCol := n - 1

        // Traverse from right to left
        for col := n - 1; col >= 0; col-- {

            // Obstacle found
            if boxGrid[row][col] == '*' {

                // Reset valid position
                emptyCol = col - 1

            // Stone found
            } else if boxGrid[row][col] == '#' {

                // Remove current stone
                boxGrid[row][col] = '.'

                // Move stone
                boxGrid[row][emptyCol] = '#'

                // Update next empty position
                emptyCol--
            }
        }
    }

    // Create rotated matrix
    rotated := make([][]byte, n)

    for i := 0; i < n; i++ {
        rotated[i] = make([]byte, m)
    }

    // Rotate clockwise
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {

            rotated[j][m-1-i] = boxGrid[i][j]
        }
    }

    return rotated
}