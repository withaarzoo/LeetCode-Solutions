#include <vector>
#include <numeric>
#include <functional>

class Solution
{
public:
    int regionsBySlashes(std::vector<std::string> &grid)
    {
        int n = grid.size(); // Get the size of the grid (n x n).

        // Initialize the union-find (disjoint-set) structures.
        std::vector<int> parent(4 * n * n), rank(4 * n * n, 0);

        // Initialize the parent array where each element is its own parent.
        std::iota(parent.begin(), parent.end(), 0);

        // Define the find function with path compression.
        std::function<int(int)> find = [&](int x)
        {
            // If x is its own parent, return x; otherwise, find the root and compress the path.
            return parent[x] == x ? x : parent[x] = find(parent[x]);
        };

        // Define the union function with union by rank.
        auto unite = [&](int x, int y)
        {
            int rootX = find(x); // Find the root of x.
            int rootY = find(y); // Find the root of y.

            // If the roots are different, perform union.
            if (rootX != rootY)
            {
                // Union by rank: attach the smaller tree under the larger tree.
                if (rank[rootX] > rank[rootY])
                {
                    parent[rootY] = rootX;
                }
                else if (rank[rootX] < rank[rootY])
                {
                    parent[rootX] = rootY;
                }
                else
                {
                    parent[rootY] = rootX; // If ranks are equal, make rootX the parent and increment its rank.
                    rank[rootX]++;
                }
            }
        };

        // Iterate over each cell in the grid.
        for (int i = 0; i < n; ++i)
        {
            for (int j = 0; j < n; ++j)
            {
                int base = 4 * (i * n + j); // Calculate the base index for the current cell.
                char c = grid[i][j];        // Get the character in the current cell.

                // Connect parts within the cell depending on the character ('/', '\\', or ' ').
                if (c == '/')
                {
                    // Connect top-left (0) with bottom-right (3) and top-right (1) with bottom-left (2).
                    unite(base, base + 3);
                    unite(base + 1, base + 2);
                }
                else if (c == '\\')
                {
                    // Connect top-left (0) with top-right (1) and bottom-left (2) with bottom-right (3).
                    unite(base, base + 1);
                    unite(base + 2, base + 3);
                }
                else
                {
                    // If the cell is empty (' '), connect all parts together.
                    unite(base, base + 1);
                    unite(base + 1, base + 2);
                    unite(base + 2, base + 3);
                }

                // Connect with the right cell (if it exists).
                if (j + 1 < n)
                {
                    unite(base + 1, base + 7); // Connect the right edge of the current cell with the left edge of the right cell.
                }

                // Connect with the cell below (if it exists).
                if (i + 1 < n)
                {
                    unite(base + 2, base + 4 * n); // Connect the bottom edge of the current cell with the top edge of the below cell.
                }
            }
        }

        // Count the number of distinct regions.
        int regions = 0;
        for (int i = 0; i < 4 * n * n; ++i)
        {
            // If an element is its own parent, it represents a distinct region.
            if (find(i) == i)
            {
                regions++;
            }
        }

        return regions; // Return the total number of regions.
    }
};
