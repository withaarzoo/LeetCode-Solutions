func findLengthOfShortestSubarray(arr []int) int {
    n := len(arr)
    
    // Step 1: Find the longest non-decreasing prefix
    left := 0
    for left+1 < n && arr[left] <= arr[left+1] {
        left++
    }
    
    // If the entire array is already sorted
    if left == n-1 {
        return 0
    }
    
    // Step 2: Find the longest non-decreasing suffix
    right := n - 1
    for right > 0 && arr[right-1] <= arr[right] {
        right--
    }
    
    // Step 3: Find the minimum length to remove by comparing prefix and suffix
    result := min(n-left-1, right)
    
    // Step 4: Use two pointers to find the smallest middle part to remove
    i, j := 0, right
    for i <= left && j < n {
        if arr[i] <= arr[j] {
            result = min(result, j-i-1)
            i++
        } else {
            j++
        }
    }
    
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}