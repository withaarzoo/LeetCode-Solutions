func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    // Initialize the graph
    graph := make([][]bool, numCourses)
    for i := 0; i < numCourses; i++ {
        graph[i] = make([]bool, numCourses)
    }

    // Build the direct edges from prerequisites
    for _, edge := range prerequisites {
        graph[edge[0]][edge[1]] = true
    }

    // Floyd-Warshall to compute transitive closure
    for k := 0; k < numCourses; k++ {
        for i := 0; i < numCourses; i++ {
            for j := 0; j < numCourses; j++ {
                if graph[i][k] && graph[k][j] {
                    graph[i][j] = true
                }
            }
        }
    }

    // Answer the queries
    result := make([]bool, len(queries))
    for i, query := range queries {
        result[i] = graph[query[0]][query[1]]
    }

    return result
}
