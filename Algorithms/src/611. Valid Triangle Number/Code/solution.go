import "sort"

// triangleNumber returns the number of triplets that can form a triangle.
func triangleNumber(nums []int) int {
    n := len(nums)
    if n < 3 {
        return 0
    }
    sort.Ints(nums)                            // sort ascending
    count := 0
    for k := n - 1; k >= 2; k-- {
        i, j := 0, k - 1                       // two pointers
        for i < j {
            if nums[i] + nums[j] > nums[k] {
                count += j - i                 // all pairs (i..j-1, j) are valid
                j--                            // try smaller b
            } else {
                i++                            // need larger a
            }
        }
    }
    return count
}
