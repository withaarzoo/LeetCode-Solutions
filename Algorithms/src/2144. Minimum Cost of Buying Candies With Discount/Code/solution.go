func minimumCost(cost []int) int {
    // Sort candies from highest cost to lowest cost
    sort.Slice(cost, func(i, j int) bool {
        return cost[i] > cost[j]
    })

    total := 0

    // Skip every third candy
    for i := 0; i < len(cost); i++ {
        if i%3 == 2 {
            continue
        }

        total += cost[i]
    }

    return total
}