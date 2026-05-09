func rotateGrid(grid [][]int, k int) [][]int {

    m := len(grid)
    n := len(grid[0])

    // Total layers in matrix
    layers := min(m, n) / 2

    for layer := 0; layer < layers; layer++ {

        nums := []int{}

        top := layer
        bottom := m - layer - 1
        left := layer
        right := n - layer - 1

        // Store top row
        for j := left; j <= right; j++ {
            nums = append(nums, grid[top][j])
        }

        // Store right column
        for i := top + 1; i <= bottom-1; i++ {
            nums = append(nums, grid[i][right])
        }

        // Store bottom row
        for j := right; j >= left; j-- {
            nums = append(nums, grid[bottom][j])
        }

        // Store left column
        for i := bottom - 1; i >= top+1; i-- {
            nums = append(nums, grid[i][left])
        }

        length := len(nums)

        // Effective rotations only
        rotate := k % length

        rotated := make([]int, length)

        // Left rotation
        for i := 0; i < length; i++ {
            rotated[i] = nums[(i+rotate)%length]
        }

        idx := 0

        // Fill top row
        for j := left; j <= right; j++ {
            grid[top][j] = rotated[idx]
            idx++
        }

        // Fill right column
        for i := top + 1; i <= bottom-1; i++ {
            grid[i][right] = rotated[idx]
            idx++
        }

        // Fill bottom row
        for j := right; j >= left; j-- {
            grid[bottom][j] = rotated[idx]
            idx++
        }

        // Fill left column
        for i := bottom - 1; i >= top+1; i-- {
            grid[i][left] = rotated[idx]
            idx++
        }
    }

    return grid
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}