/**
 * @param {string[]} grid - The grid of slashes and backslashes representing the regions
 * @return {number} - The number of distinct regions
 */
var regionsBySlashes = function (grid) {
  const n = grid.length; // The size of the grid (number of rows/columns)

  // Initialize the parent array for Union-Find, treating each grid cell as 4 sub-cells
  const parent = Array.from({ length: 4 * n * n }, (_, i) => i);

  // Initialize the rank array to keep track of the depth of each tree in the Union-Find structure
  const rank = Array(4 * n * n).fill(0);

  // Helper function to find the root of the set that element x belongs to
  // Uses path compression to flatten the structure for faster future queries
  function find(x) {
    if (parent[x] !== x) {
      parent[x] = find(parent[x]); // Path compression
    }
    return parent[x];
  }

  // Helper function to union two sets, x and y
  // Uses union by rank to attach the smaller tree to the root of the larger tree
  function unite(x, y) {
    const rootX = find(x); // Find the root of x
    const rootY = find(y); // Find the root of y

    // If they have different roots, union them
    if (rootX !== rootY) {
      if (rank[rootX] > rank[rootY]) {
        parent[rootY] = rootX; // Attach rootY to rootX
      } else if (rank[rootX] < rank[rootY]) {
        parent[rootX] = rootY; // Attach rootX to rootY
      } else {
        parent[rootY] = rootX; // Attach rootY to rootX and increase rank of rootX
        rank[rootX]++;
      }
    }
  }

  // Iterate through each cell in the grid
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
      const base = 4 * (i * n + j); // Calculate the base index for the current cell's sub-cells
      const c = grid[i][j]; // The character in the current cell ('/', '\\', or ' ')

      // Connect the internal sub-cells based on the character
      if (c === "/") {
        unite(base, base + 3); // Connect top-left and bottom-right sub-cells
        unite(base + 1, base + 2); // Connect top-right and bottom-left sub-cells
      } else if (c === "\\") {
        unite(base, base + 1); // Connect top-left and top-right sub-cells
        unite(base + 2, base + 3); // Connect bottom-left and bottom-right sub-cells
      } else {
        // Connect all sub-cells for a blank space ' '
        unite(base, base + 1); // Top-left to top-right
        unite(base + 1, base + 2); // Top-right to bottom-right
        unite(base + 2, base + 3); // Bottom-right to bottom-left
      }

      // Connect the right neighbor's left sub-cell to the current cell's right sub-cell
      if (j + 1 < n) {
        unite(base + 1, base + 7);
      }

      // Connect the bottom neighbor's top sub-cell to the current cell's bottom sub-cell
      if (i + 1 < n) {
        unite(base + 2, base + 4 * n);
      }
    }
  }

  // Count the number of distinct regions by checking how many unique roots exist
  let regions = 0;
  for (let i = 0; i < 4 * n * n; i++) {
    if (find(i) === i) regions++; // If the element is its own root, it represents a distinct region
  }

  return regions; // Return the total number of distinct regions
};
