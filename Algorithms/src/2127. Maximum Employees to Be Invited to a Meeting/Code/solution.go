func maximumInvitations(favorite []int) int {
    n := len(favorite)
    inDegree := make([]int, n)
    chainLengths := make([]int, n)
    visited := make([]bool, n)

    for _, fav := range favorite {
        inDegree[fav]++
    }

    queue := []int{}
    for i := 0; i < n; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        visited[node] = true

        next := favorite[node]
        chainLengths[next] = chainLengths[node] + 1
        inDegree[next]--
        if inDegree[next] == 0 {
            queue = append(queue, next)
        }
    }

    maxCycle, totalChains := 0, 0
    for i := 0; i < n; i++ {
        if !visited[i] {
            current, cycleLength := i, 0
            for !visited[current] {
                visited[current] = true
                current = favorite[current]
                cycleLength++
            }

            if cycleLength == 2 {
                totalChains += 2 + chainLengths[i] + chainLengths[favorite[i]]
            } else {
                if cycleLength > maxCycle {
                    maxCycle = cycleLength
                }
            }
        }
    }

    if maxCycle > totalChains {
        return maxCycle
    }
    return totalChains
}
