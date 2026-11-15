func numOfSubarrays(arr []int) int {
    const MOD = 1000000007
    even, odd := 1, 0
    ans := 0
    sum := 0
    for _, x := range arr {
        sum = (sum + x) % 2
        if sum == 1 {
            ans = (ans + even) % MOD
            odd++
        } else {
            ans = (ans + odd) % MOD
            even++
        }
    }
    return ans
}