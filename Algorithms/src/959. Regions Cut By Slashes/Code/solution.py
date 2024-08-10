class Solution:
    def regionsBySlashes(self, grid: List[str]) -> int:
        n = len(grid)
        
        # Each cell in the grid is divided into 4 triangles, so we need 4 * n * n elements
        # for the union-find (disjoint-set) array to represent these parts.
        parent = list(range(4 * n * n))  # Initialize parent array for union-find
        rank = [0] * (4 * n * n)  # Initialize rank array for union by rank

        # Find function with path compression to find the root of an element
        def find(x):
            if parent[x] != x:  # If x is not its own parent, it's not the root
                parent[x] = find(parent[x])  # Path compression: point x directly to the root
            return parent[x]  # Return the root of x

        # Union function with union by rank to unite two sets
        def unite(x, y):
            rootX = find(x)  # Find root of x
            rootY = find(y)  # Find root of y
            if rootX != rootY:  # Only unite if they have different roots
                if rank[rootX] > rank[rootY]:  # If rootX's tree is taller
                    parent[rootY] = rootX  # Make rootY a child of rootX
                elif rank[rootX] < rank[rootY]:  # If rootY's tree is taller
                    parent[rootX] = rootY  # Make rootX a child of rootY
                else:
                    parent[rootY] = rootX  # If they have the same height, make rootY a child of rootX
                    rank[rootX] += 1  # Increase the rank of rootX

        # Loop through each cell in the grid
        for i in range(n):
            for j in range(n):
                base = 4 * (i * n + j)  # Base index for the current cell's 4 triangles
                c = grid[i][j]  # Get the character in the current cell

                # Connect the parts inside the current cell based on the character
                if c == '/':
                    unite(base, base + 3)  # Connect top-left with bottom-right
                    unite(base + 1, base + 2)  # Connect top-right with bottom-left
                elif c == '\\':
                    unite(base, base + 1)  # Connect top-left with top-right
                    unite(base + 2, base + 3)  # Connect bottom-left with bottom-right
                else:
                    # Connect all four triangles if the cell is empty
                    unite(base, base + 1)  # Top-left with top-right
                    unite(base + 1, base + 2)  # Top-right with bottom-left
                    unite(base + 2, base + 3)  # Bottom-left with bottom-right

                # Connect with the adjacent right cell (if exists)
                if j + 1 < n:
                    unite(base + 1, 4 * (i * n + (j + 1)))  # Connect right part of current cell with left part of right cell

                # Connect with the cell below (if exists)
                if i + 1 < n:
                    unite(base + 2, 4 * ((i + 1) * n + j))  # Connect bottom part of current cell with top part of cell below

        # Count the number of distinct regions by counting the number of unique roots in the union-find structure
        regions = sum(find(i) == i for i in range(4 * n * n))  # Count roots
        return regions  # Return the total number of regions
