func findRedundantConnection(edges [][]int) []int {
    parent := make([]int, len(edges)+1)
    rank := make([]int, len(edges)+1)

    for i := range parent {
        parent[i] = i
    }

    var find func(int) int
    find = func(node int) int {
        if parent[node] != node {
            parent[node] = find(parent[node])
        }
        return parent[node]
    }

    union := func(u, v int) bool {
        rootU, rootV := find(u), find(v)
        if rootU == rootV {
            return false
        }
        if rank[rootU] > rank[rootV] {
            parent[rootV] = rootU
        } else if rank[rootU] < rank[rootV] {
            parent[rootU] = rootV
        } else {
            parent[rootV] = rootU
            rank[rootU]++
        }
        return true
    }

    for _, edge := range edges {
        if !union(edge[0], edge[1]) {
            return edge
        }
    }
    return nil
}
