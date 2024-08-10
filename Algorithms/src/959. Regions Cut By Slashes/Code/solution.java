class Solution {
    public int regionsBySlashes(String[] grid) {
        int n = grid.length;

        // The grid is divided into n*n cells, each cell divided into 4 triangles.
        // We need 4 * n * n to represent the 4 triangles in each cell.
        int[] parent = new int[4 * n * n];
        int[] rank = new int[4 * n * n];

        // Initialize the union-find structure
        // Each triangle is its own parent initially
        for (int i = 0; i < parent.length; i++) {
            parent[i] = i;
        }

        // Function to find the root of a set with path compression
        int find(int x) {
            // If the current node is not the root, find its parent recursively
            // Path compression: flatten the structure by setting the parent to the root
            if (parent[x] != x) {
                parent[x] = find(parent[x]);
            }
            return parent[x];
        }

        // Function to unite two sets using union by rank
        void unite(int x, int y) {
            int rootX = find(x);
            int rootY = find(y);

            // Only unite if they are in different sets
            if (rootX != rootY) {
                // Attach the shorter tree under the taller tree
                if (rank[rootX] > rank[rootY]) {
                    parent[rootY] = rootX;
                } else if (rank[rootX] < rank[rootY]) {
                    parent[rootX] = rootY;
                } else {
                    // If they are the same height, make one root and increase its rank
                    parent[rootY] = rootX;
                    rank[rootX]++;
                }
            }
        }

        // Traverse each cell in the grid
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < n; ++j) {
                // Base index for the four triangles in the current cell
                int base = 4 * (i * n + j);
                char c = grid[i].charAt(j);

                // Depending on the slash type, connect the corresponding triangles
                if (c == '/') {
                    // '/' splits the cell diagonally from top-right to bottom-left
                    unite(base, base + 3); // Connect top triangle with bottom triangle
                    unite(base + 1, base + 2); // Connect left triangle with right triangle
                } else if (c == '\\') {
                    // '\' splits the cell diagonally from top-left to bottom-right
                    unite(base, base + 1); // Connect top triangle with right triangle
                    unite(base + 2, base + 3); // Connect left triangle with bottom triangle
                } else {
                    // ' ' (space) means the cell is empty, connect all four triangles
                    unite(base, base + 1); // Connect top triangle with right triangle
                    unite(base + 1, base + 2); // Connect right triangle with bottom triangle
                    unite(base + 2, base + 3); // Connect bottom triangle with left triangle
                }

                // Connect current cell's right triangle with the left triangle of the cell to the right
                if (j + 1 < n) {
                    unite(base + 1, base + 7);
                }

                // Connect current cell's bottom triangle with the top triangle of the cell below
                if (i + 1 < n) {
                    unite(base + 2, base + 4 * n);
                }
            }
        }

        // Count the number of distinct regions by checking the number of unique roots
        int regions = 0;
        for (int i = 0; i < 4 * n * n; ++i) {
            // If a node is its own parent, it is the root of a distinct region
            if (find(i) == i) {
                regions++;
            }
        }

        // Return the number of distinct regions
        return regions;
    }
}
