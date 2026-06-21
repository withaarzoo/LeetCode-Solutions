func maxIceCream(costs []int, coins int) int {

    // Maximum possible cost according to constraints
    const MAX_COST = 100000

    // Frequency array to count occurrences of each cost
    freq := make([]int, MAX_COST+1)

    // Count every ice cream cost
    for _, cost := range costs {
        freq[cost]++
    }

    // Stores total purchased bars
    answer := 0

    // Process costs from smallest to largest
    for cost := 1; cost <= MAX_COST; cost++ {

        // Skip unavailable costs
        if freq[cost] == 0 {
            continue
        }

        // Maximum bars affordable at current cost
        canBuy := freq[cost]
        affordable := coins / cost

        if canBuy > affordable {
            canBuy = affordable
        }

        // Add purchased bars
        answer += canBuy

        // Deduct spent coins
        coins -= canBuy * cost
    }

    return answer
}