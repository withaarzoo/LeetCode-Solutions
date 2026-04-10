func minimumDistance(nums []int) int {
    pos := make(map[int][]int)
    
    // Store all indices for each value
    for i, num := range nums {
        pos[num] = append(pos[num], i)
    }
    
    ans := int(^uint(0) >> 1) // Maximum int value
    
    // Process each value's indices
    for _, indices := range pos {
        if len(indices) < 3 {
            continue
        }
        
        // Check every consecutive group of 3 indices
        for i := 0; i+2 < len(indices); i++ {
            distance := 2 * (indices[i+2] - indices[i])
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