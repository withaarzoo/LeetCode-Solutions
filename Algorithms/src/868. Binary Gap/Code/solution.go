func binaryGap(n int) int {
    lastPosition := -1      // last seen 1 index
    maxDistance := 0        // max gap
    currentPosition := 0    // bit index
    
    for n > 0 {
        // Check if last bit is 1
        if n&1 == 1 {
            if lastPosition != -1 {
                if currentPosition-lastPosition > maxDistance {
                    maxDistance = currentPosition - lastPosition
                }
            }
            lastPosition = currentPosition
        }
        
        n >>= 1 // shift right
        currentPosition++
    }
    
    return maxDistance
}