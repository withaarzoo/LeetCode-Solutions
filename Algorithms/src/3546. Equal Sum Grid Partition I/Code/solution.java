class Solution {
    public boolean canPartitionGrid(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        
        long total = 0;
        
        // Step 1: Total sum
        for (int[] row : grid) {
            for (int val : row) {
                total += val;
            }
        }
        
        // Step 2: Odd check
        if (total % 2 != 0) return false;
        
        long target = total / 2;
        
        // Step 3: Horizontal cut
        long rowSum = 0;
        for (int i = 0; i < m - 1; i++) {
            for (int j = 0; j < n; j++) {
                rowSum += grid[i][j];
            }
            if (rowSum == target) return true;
        }
        
        // Step 4: Column sums
        long[] colSum = new long[n];
        for (int j = 0; j < n; j++) {
            for (int i = 0; i < m; i++) {
                colSum[j] += grid[i][j];
            }
        }
        
        // Step 5: Vertical cut
        long curr = 0;
        for (int j = 0; j < n - 1; j++) {
            curr += colSum[j];
            if (curr == target) return true;
        }
        
        return false;
    }
}