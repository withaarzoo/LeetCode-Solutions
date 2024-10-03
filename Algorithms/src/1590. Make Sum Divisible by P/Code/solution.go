func minSubarray(nums []int, p int) int {
    totalSum := 0
    for _, num := range nums {
        totalSum += num
    }

    rem := totalSum % p
    if rem == 0 {
        return 0
    }

    prefixMod := make(map[int]int)
    prefixMod[0] = -1
    prefixSum, minLength := 0, len(nums)

    for i, num := range nums {
        prefixSum += num
        currentMod := prefixSum % p
        targetMod := (currentMod - rem + p) % p

        if idx, ok := prefixMod[targetMod]; ok {
            if i-idx < minLength {
                minLength = i - idx
            }
        }

        prefixMod[currentMod] = i
    }

    if minLength == len(nums) {
        return -1
    }
    return minLength
}