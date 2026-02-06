func minRemoval(nums []int, k int) int {
    sort.Ints(nums)

    left := 0
    maxKeep := 1

    for right := 0; right < len(nums); right++ {
        for nums[right] > nums[left]*k {
            left++
        }
        if right-left+1 > maxKeep {
            maxKeep = right - left + 1
        }
    }

    return len(nums) - maxKeep
}
