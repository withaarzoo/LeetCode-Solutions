class Solution {
    public int largestSubmatrix(int[][] matrix) {
        int m = matrix.length, n = matrix[0].length;

        // Step 1: Build heights
        for (int i = 1; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (matrix[i][j] == 1) {
                    matrix[i][j] += matrix[i - 1][j];
                }
            }
        }

        int maxArea = 0;

        // Step 2 & 3
        for (int i = 0; i < m; i++) {
            int[] row = matrix[i].clone();
            Arrays.sort(row);

            // Traverse from end (descending)
            for (int j = 0; j < n; j++) {
                int height = row[n - 1 - j];
                int area = height * (j + 1);
                maxArea = Math.max(maxArea, area);
            }
        }

        return maxArea;
    }
}