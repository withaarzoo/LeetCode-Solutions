func maxBottlesDrunk(numBottles int, numExchange int) int {
    full := numBottles
    empty := 0
    ans := 0
    curEx := numExchange

    for full > 0 {
        // drink all full bottles
        ans += full
        empty += full
        full = 0

        // exchange one-by-one while possible
        for empty >= curEx {
            empty -= curEx
            full += 1
            curEx += 1
        }
    }
    return ans
}
