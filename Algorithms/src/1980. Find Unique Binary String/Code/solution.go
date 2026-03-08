func findDifferentBinaryString(nums []string) string {
    n := len(nums)
    result := make([]byte, n)

    // Flip the diagonal bits
    for i := 0; i < n; i++ {
        if nums[i][i] == '0' {
            result[i] = '1'
        } else {
            result[i] = '0'
        }
    }

    return string(result)
}