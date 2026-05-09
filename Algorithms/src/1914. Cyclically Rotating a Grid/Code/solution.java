class Solution {
    public int[][] rotateGrid(int[][] grid, int k) {

        int m = grid.length;
        int n = grid[0].length;

        // Total layers
        int layers = Math.min(m, n) / 2;

        for (int layer = 0; layer < layers; layer++) {

            ArrayList<Integer> nums = new ArrayList<>();

            int top = layer;
            int bottom = m - layer - 1;
            int left = layer;
            int right = n - layer - 1;

            // Store top row
            for (int j = left; j <= right; j++) {
                nums.add(grid[top][j]);
            }

            // Store right column
            for (int i = top + 1; i <= bottom - 1; i++) {
                nums.add(grid[i][right]);
            }

            // Store bottom row
            for (int j = right; j >= left; j--) {
                nums.add(grid[bottom][j]);
            }

            // Store left column
            for (int i = bottom - 1; i >= top + 1; i--) {
                nums.add(grid[i][left]);
            }

            int len = nums.size();

            // Effective rotations
            int rotate = k % len;

            int[] rotated = new int[len];

            // Left rotation
            for (int i = 0; i < len; i++) {
                rotated[i] = nums.get((i + rotate) % len);
            }

            int idx = 0;

            // Fill top row
            for (int j = left; j <= right; j++) {
                grid[top][j] = rotated[idx++];
            }

            // Fill right column
            for (int i = top + 1; i <= bottom - 1; i++) {
                grid[i][right] = rotated[idx++];
            }

            // Fill bottom row
            for (int j = right; j >= left; j--) {
                grid[bottom][j] = rotated[idx++];
            }

            // Fill left column
            for (int i = bottom - 1; i >= top + 1; i--) {
                grid[i][left] = rotated[idx++];
            }
        }

        return grid;
    }
}