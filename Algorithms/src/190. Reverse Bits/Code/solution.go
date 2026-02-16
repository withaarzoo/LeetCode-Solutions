func reverseBits(n int) int {
    var result int = 0
    
    for i := 0; i < 32; i++ {
        
        // Shift result left
        result <<= 1
        
        // Add last bit of n
        result |= (n & 1)
        
        // Shift n right
        n >>= 1
    }
    
    return result
}
