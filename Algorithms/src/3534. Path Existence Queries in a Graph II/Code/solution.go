func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
    // order stores original node indices before sorting by nums value.
    order := make([]int, n)

    for i := 0; i < n; i++ {
        order[i] = i
    }

    // Sort node indices according to their nums values.
    sort.Slice(order, func(i, j int) bool {
        return nums[order[i]] < nums[order[j]]
    })

    // pos[node] gives the sorted position of an original node.
    pos := make([]int, n)

    // values stores nums in sorted order.
    values := make([]int, n)

    for i := 0; i < n; i++ {
        values[i] = nums[order[i]]
        pos[order[i]] = i
    }

    // Find the number of levels needed for binary lifting.
    LOG := 1

    for (1 << LOG) <= n {
        LOG++
    }

    // jump[p][i] is the farthest position after at most 2^p greedy jumps.
    jump := make([][]int, LOG)

    for p := 0; p < LOG; p++ {
        jump[p] = make([]int, n)
    }

    // Use two pointers to find every one-jump destination.
    r := 0

    for i := 0; i < n; i++ {
        // Keep r at or after the current position.
        if r < i {
            r = i
        }

        // Extend r while a direct edge from i is still possible.
        for r+1 < n && values[r+1]-values[i] <= maxDiff {
            r++
        }

        // Save the farthest position reachable in one edge.
        jump[0][i] = r
    }

    // Build all larger binary jumps.
    for p := 1; p < LOG; p++ {
        for i := 0; i < n; i++ {
            // Apply the previous jump twice to double the number of jumps.
            jump[p][i] = jump[p-1][jump[p-1][i]]
        }
    }

    // Store one result for every query.
    answer := make([]int, len(queries))

    for q, query := range queries {
        // Convert original nodes into sorted positions.
        left := pos[query[0]]
        right := pos[query[1]]

        // The graph is undirected, so always move from left to right.
        if left > right {
            left, right = right, left
        }

        // No edge is needed when both nodes are the same.
        if left == right {
            answer[q] = 0
            continue
        }

        current := left
        distance := 0

        // Take the largest groups of jumps that still stop before the target.
        for p := LOG - 1; p >= 0; p-- {
            if jump[p][current] < right {
                // Skip 2^p greedy jumps at once.
                current = jump[p][current]
                distance += 1 << p
            }
        }

        // One final edge must reach the target.
        if jump[0][current] >= right {
            answer[q] = distance + 1
        } else {
            // Otherwise the two nodes are in different connected components.
            answer[q] = -1
        }
    }

    return answer
}