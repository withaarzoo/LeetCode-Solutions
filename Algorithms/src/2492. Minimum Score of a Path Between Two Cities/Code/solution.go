func minScore(n int, roads [][]int) int {
    // Edge stores the neighboring city and the road distance.
    type Edge struct {
        city     int
        distance int
    }

    // graph[city] stores all roads connected to that city.
    graph := make([][]Edge, n+1)

    // Build the undirected graph by storing every road in both directions.
    for _, road := range roads {
        a := road[0]
        b := road[1]
        distance := road[2]

        graph[a] = append(graph[a], Edge{b, distance})
        graph[b] = append(graph[b], Edge{a, distance})
    }

    // visited prevents processing the same city more than once.
    visited := make([]bool, n+1)

    // I use a slice as a BFS queue and move a front pointer.
    queue := []int{1}
    front := 0
    visited[1] = true

    // Start with the largest possible integer value.
    answer := int(^uint(0) >> 1)

    // Visit every city in the connected component containing city 1.
    for front < len(queue) {
        city := queue[front]
        front++

        // Check every road connected to the current city.
        for _, edge := range graph[city] {
            // Keep the smallest road distance found in this component.
            if edge.distance < answer {
                answer = edge.distance
            }

            // Add the neighboring city only if it has not been visited.
            if !visited[edge.city] {
                visited[edge.city] = true
                queue = append(queue, edge.city)
            }
        }
    }

    // Return the minimum road distance in city 1's component.
    return answer
}