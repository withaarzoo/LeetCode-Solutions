func minSwaps(grid [][]int) int {
    n := len(grid)
    
    trailing := make([]int, n)
    
    // Count trailing zeros
    for i := 0; i < n; i++ {
        count := 0
        for j := n - 1; j >= 0; j-- {
            if grid[i][j] == 0 {
                count++
            } else {
                break
            }
        }
        trailing[i] = count
    }
    
    swaps := 0
    
    for i := 0; i < n; i++ {
        required := n - 1 - i
        j := i
        
        for j < n && trailing[j] < required {
            j++
        }
        
        if j == n {
            return -1
        }
        
        for j > i {
            trailing[j], trailing[j-1] = trailing[j-1], trailing[j]
            swaps++
            j--
        }
    }
    
    return swaps
}