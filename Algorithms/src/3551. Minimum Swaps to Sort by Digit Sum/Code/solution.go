func minSwaps(nums []int) int {
    digitSum := func(x int) int {
        s := 0
        for x > 0 {
            s += x % 10
            x /= 10
        }
        return s
    }
    n := len(nums)
    sortedNums := make([]int, n)
    copy(sortedNums, nums)
    sort.Slice(sortedNums, func(i, j int) bool {
        sa, sb := digitSum(sortedNums[i]), digitSum(sortedNums[j])
        if sa != sb {
            return sa < sb
        }
        return sortedNums[i] < sortedNums[j]
    })
    pos := make(map[int]int)
    for i, v := range sortedNums {
        pos[v] = i
    }
    visited := make([]bool, n)
    cycles := 0
    for i := 0; i < n; i++ {
        if !visited[i] {
            cycles++
            j := i
            for !visited[j] {
                visited[j] = true
                j = pos[nums[j]]
            }
        }
    }
    return n - cycles
}