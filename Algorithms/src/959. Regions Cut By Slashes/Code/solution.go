func regionsBySlashes(grid []string) int {
    n := len(grid) // Determine the size of the grid (n x n)
    
    // Initialize the union-find data structures
    parent := make([]int, 4 * n * n) // Each cell is divided into 4 regions, so we have 4*n*n regions in total
    rank := make([]int, 4 * n * n)   // Rank array to optimize the union operation

    // Initialize the parent of each region to itself
    for i := 0; i < len(parent); i++ {
        parent[i] = i
    }

    // Find function with path compression
    var find func(x int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x]) // Path compression: make every node in the path point directly to the root
        }
        return parent[x]
    }

    // Union function with union by rank
    unite := func(x, y int) {
        rootX := find(x) // Find the root of x
        rootY := find(y) // Find the root of y
        if rootX != rootY {
            if rank[rootX] > rank[rootY] { // Union by rank
                parent[rootY] = rootX // Attach the tree with smaller rank to the root of the tree with larger rank
            } else if rank[rootX] < rank[rootY] {
                parent[rootX] = rootY
            } else {
                parent[rootY] = rootX // If ranks are equal, choose one as the root and increase its rank
                rank[rootX]++
            }
        }
    }

    // Iterate over each cell in the grid
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            base := 4 * (i * n + j) // Calculate the base index for the current cell's regions
            c := grid[i][j]         // Get the character in the current cell

            // Connect the parts inside the current cell based on the slash type
            if c == '/' {
                unite(base, base+3) // Connect top-left and bottom-right regions
                unite(base+1, base+2) // Connect top-right and bottom-left regions
            } else if c == '\\' {
                unite(base, base+1) // Connect top-left and top-right regions
                unite(base+2, base+3) // Connect bottom-left and bottom-right regions
            } else {
                // If it's a space, connect all regions inside the cell
                unite(base, base+1)
                unite(base+1, base+2)
                unite(base+2, base+3)
            }

            // Connect with the right cell (adjacent horizontally)
            if j+1 < n {
                unite(base+1, base+7)
            }

            // Connect with the cell below (adjacent vertically)
            if i+1 < n {
                unite(base+2, base+4*n)
            }
        }
    }

    // Count distinct regions by checking how many unique roots there are
    regions := 0
    for i := 0; i < 4 * n * n; i++ {
        if find(i) == i { // If a region is its own root, it represents a distinct region
            regions++
        }
    }

    return regions // Return the number of distinct regions
}
