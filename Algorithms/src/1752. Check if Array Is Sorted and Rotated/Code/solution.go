func check(nums []int) bool {
    
    n := len(nums)

    // Counts how many times order decreases
    count := 0

    // Traverse the array
    for i := 0; i < n; i++ {

        // Compare current element with next element
        // % n connects last element with first
        if nums[i] > nums[(i+1)%n] {
            count++
        }

        // More than one decrease means invalid
        if count > 1 {
            return false
        }
    }

    // Valid sorted rotated array
    return true
}