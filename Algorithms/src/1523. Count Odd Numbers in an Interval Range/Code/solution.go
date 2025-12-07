func countOdds(low int, high int) int {
    // Helper: count of odd numbers from 1 to x
    oddsUpTo := func(x int) int {
        return (x + 1) / 2
    }
    
    // Odds in [low, high]
    return oddsUpTo(high) - oddsUpTo(low-1)
}
