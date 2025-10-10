func maximumEnergy(energy []int, k int) int {
    n := len(energy)
    ans := int64(-1 << 62) // very small initial value
    for r := 0; r < k; r++ {
        var cur int64 = 0
        last := r + ((n-1 - r)/k)*k // last index in the class
        for i := last; i >= r; i -= k {
            cur += int64(energy[i])  // suffix sum from i to class end
            if cur > ans {
                ans = cur
            }
        }
    }
    return int(ans)
}
