func resultArray(nums []int, k int) []int64 {
    result := make([]int64, k)
    curr := make([]int64, k)
    for _, num := range nums {
        m := num % k
        next := make([]int64, k)
        next[m] = 1
        for prev := 0; prev < k; prev++ {
            if curr[prev] > 0 {
                nr := (prev * m) % k
                next[nr] += curr[prev]
            }
        }
        for r := 0; r < k; r++ {
            result[r] += next[r]
        }
        curr = next
    }
    return result
}