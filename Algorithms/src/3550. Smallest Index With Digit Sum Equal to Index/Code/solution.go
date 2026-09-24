func smallestIndex(nums []int) int {
    for i := 0; i < len(nums); i++ {
        x := nums[i]
        sum := 0
        for x > 0 {
            sum += x % 10
            x /= 10
        }
        if sum == i {
            return i
        }
    }
    return -1
}