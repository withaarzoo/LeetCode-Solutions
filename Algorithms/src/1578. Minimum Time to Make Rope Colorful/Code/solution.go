package main

// minCost returns the minimal time to make the rope colorful.
// It uses O(n) time and O(1) extra space.
func minCost(colors string, neededTime []int) int {
    var ans int64 = 0
    var blockSum int64 = 0
    var blockMax int64 = 0
    n := len(colors)
    
    for i := 0; i < n; i++ {
        if i > 0 && colors[i] != colors[i-1] {
            ans += blockSum - blockMax
            blockSum = 0
            blockMax = 0
        }
        blockSum += int64(neededTime[i])
        if int64(neededTime[i]) > blockMax {
            blockMax = int64(neededTime[i])
        }
    }
    ans += blockSum - blockMax
    return int(ans)
}
