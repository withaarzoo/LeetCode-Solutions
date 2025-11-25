func smallestRepunitDivByK(k int) int {
    // If k is divisible by 2 or 5, impossible to be divisible by a number like 111...
    if k%2 == 0 || k%5 == 0 {
        return -1
    }

    rem := 0 // current remainder

    for length := 1; length <= k; length++ {
        // Update remainder when we append '1'
        rem = (rem*10 + 1) % k
        if rem == 0 {
            return length
        }
    }

    return -1
}
