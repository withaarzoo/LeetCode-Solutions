import "strconv"

func readBinaryWatch(turnedOn int) []string {
    var result []string
    
    // Try all possible hours
    for hour := 0; hour < 12; hour++ {
        
        // Try all possible minutes
        for minute := 0; minute < 60; minute++ {
            
            // Count total set bits
            if bitCount(hour) + bitCount(minute) == turnedOn {
                
                time := strconv.Itoa(hour) + ":"
                
                if minute < 10 {
                    time += "0"
                }
                
                time += strconv.Itoa(minute)
                
                result = append(result, time)
            }
        }
    }
    
    return result
}

// Helper function to count set bits
func bitCount(n int) int {
    count := 0
    for n > 0 {
        count += n & 1
        n >>= 1
    }
    return count
}
