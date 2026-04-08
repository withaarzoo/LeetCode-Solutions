func xorAfterQueries(nums []int, queries [][]int) int {
    const MOD int64 = 1000000007

    // Process each query
    for _, q := range queries {
        l, r, k, v := q[0], q[1], q[2], q[3]

        // Visit indices: l, l+k, l+2k, ... <= r
        for i := l; i <= r; i += k {
            nums[i] = int((int64(nums[i]) * int64(v)) % MOD)
        }
    }

    // Compute XOR of all final values
    ans := 0
    for _, num := range nums {
        ans ^= num
    }

    return ans
}