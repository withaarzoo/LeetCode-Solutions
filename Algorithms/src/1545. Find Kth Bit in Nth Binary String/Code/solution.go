func findKthBit(n int, k int) byte {
    // Base case: When n = 1, the binary string is "0"
    if n == 1 {
        return '0'
    }
    
    // Find the length of the current string Sn, which is 2^n - 1
    length := (1 << n) - 1
    
    // Find the middle position
    mid := length / 2 + 1
    
    // If k is the middle position, return '1'
    if k == mid {
        return '1'
    }
    
    // If k is in the first half, find the bit in Sn-1
    if k < mid {
        return findKthBit(n - 1, k)
    }
    
    // If k is in the second half, find the bit in Sn-1 and invert it
    if findKthBit(n - 1, length - k + 1) == '0' {
        return '1'
    }
    return '0'
}
