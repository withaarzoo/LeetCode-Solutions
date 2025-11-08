func minimumOneBitOperations(n int) int {
    // Inverse Gray code via iterative XOR-shift
    ans := 0
    for n > 0 {
        ans ^= n
        n >>= 1
    }
    return ans
}
