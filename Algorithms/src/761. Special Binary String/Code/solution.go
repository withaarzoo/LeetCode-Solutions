import "sort"

func makeLargestSpecial(s string) string {
    var parts []string
    count := 0
    start := 0
    
    for i := 0; i < len(s); i++ {
        if s[i] == '1' {
            count++
        } else {
            count--
        }
        
        if count == 0 {
            // Recursively solve inner substring
            inner := makeLargestSpecial(s[start+1 : i])
            parts = append(parts, "1"+inner+"0")
            start = i + 1
        }
    }
    
    // Sort descending
    sort.Slice(parts, func(i, j int) bool {
        return parts[i] > parts[j]
    })
    
    result := ""
    for _, p := range parts {
        result += p
    }
    
    return result
}