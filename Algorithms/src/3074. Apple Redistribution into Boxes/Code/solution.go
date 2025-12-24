func minimumBoxes(apple []int, capacity []int) int {
    // Step 1: Calculate total apples
    totalApples := 0
    for _, a := range apple {
        totalApples += a
    }

    // Step 2: Sort capacities in descending order
    sort.Slice(capacity, func(i, j int) bool {
        return capacity[i] > capacity[j]
    })

    // Step 3: Pick boxes greedily
    usedCapacity := 0
    boxes := 0

    for _, cap := range capacity {
        usedCapacity += cap
        boxes++
        if usedCapacity >= totalApples {
            return boxes
        }
    }

    return boxes
}
