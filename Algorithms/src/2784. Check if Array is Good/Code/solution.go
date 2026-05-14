func isGood(nums []int) bool {

    // Sort the array
    sort.Ints(nums)

    // Length of array
    n := len(nums)

    // Maximum element
    mx := nums[n-1]

    // Size must be mx + 1
    if n != mx+1 {
        return false
    }

    // Check sequence from 1 to mx
    for i := 0; i < n-1; i++ {

        // Expected value is i + 1
        if nums[i] != i+1 {
            return false
        }
    }

    // Last element should also be mx
    return nums[n-1] == mx
}