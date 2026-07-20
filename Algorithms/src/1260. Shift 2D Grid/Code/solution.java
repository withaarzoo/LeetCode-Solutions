class Solution {
    public List<List<Integer>> shiftGrid(int[][] grid, int k) {

        // Get the grid dimensions
        int m = grid.length;
        int n = grid[0].length;

        // Total number of elements
        int total = m * n;

        // Ignore unnecessary complete rotations
        k %= total;

        // Create the answer grid
        List<List<Integer>> ans = new ArrayList<>();

        for (int i = 0; i < m; i++) {
            List<Integer> row = new ArrayList<>();

            // Fill with dummy values so we can use set()
            for (int j = 0; j < n; j++) {
                row.add(0);
            }

            ans.add(row);
        }

        // Move every element to its final position
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {

                // Flatten the current position
                int oldIndex = i * n + j;

                // Compute shifted position
                int newIndex = (oldIndex + k) % total;

                // Convert back to row and column
                int newRow = newIndex / n;
                int newCol = newIndex % n;

                // Store the value
                ans.get(newRow).set(newCol, grid[i][j]);
            }
        }

        return ans;
    }
}