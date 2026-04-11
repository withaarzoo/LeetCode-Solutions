func minimumDistance(nums []int) int {
    positions := make(map[int][]int)
    
    // Store all indices for each value
    for i, num := range nums {
        positions[num] = append(positions[num], i)
    }
    
    ans := int(^uint(0) >> 1) // Maximum int value
    
    // Check every value's index list
    for _, idx := range positions {
        if len(idx) < 3 {
            continue
        }
        
        // Check every consecutive group of 3 indices
        for i := 0; i+2 < len(idx); i++ {
            distance := 2 * (idx[i+2] - idx[i])
            if distance < ans {
                ans = distance
            }
        }
    }
    
    if ans == int(^uint(0)>>1) {
        return -1
    }
    
    return ans
}