func maxAlternatingSum(nums []int) int64 {
    n := len(nums)
    sq := make([]int64, n)
    for i, v := range nums {
        sq[i] = int64(v) * int64(v)
    }
    sort.Slice(sq, func(i, j int) bool {
        return sq[i] > sq[j]
    })
    var ans int64
    pos := (n + 1) / 2
    for i := 0; i < n; i++ {
        if i < pos {
            ans += sq[i]
        } else {
            ans -= sq[i]
        }
    }
    return ans
}