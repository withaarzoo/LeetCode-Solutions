func targetIndices(nums []int, target int) []int {
    less, equal := 0, 0
    for _, x := range nums {
        if x < target {
            less++
        } else if x == target {
            equal++
        }
    }
    ans := make([]int, equal)
    for i := 0; i < equal; i++ {
        ans[i] = less + i
    }
    return ans
}