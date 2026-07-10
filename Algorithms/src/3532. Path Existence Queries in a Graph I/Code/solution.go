func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
    // component[i] stores which connected component node i belongs to.
    component := make([]int, n)

    // Start with component 0 for the first node.
    componentId := 0

    // Check every gap between two consecutive sorted values.
    for i := 1; i < n; i++ {
        // A gap larger than maxDiff separates the graph into two parts.
        if nums[i]-nums[i-1] > maxDiff {
            componentId++
        }

        // Store the component of the current node.
        component[i] = componentId
    }

    // Create one answer for every query.
    answer := make([]bool, len(queries))

    // Two nodes have a path exactly when their component IDs are equal.
    for i, query := range queries {
        u := query[0]
        v := query[1]

        answer[i] = component[u] == component[v]
    }

    return answer
} 