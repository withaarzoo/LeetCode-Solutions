class Solution {
    public long maxPoints(int[][] points) {
        // Get the number of rows (m) and columns (n) from the points array
        int m = points.length;
        int n = points[0].length;

        // Create a dp array to store the maximum points for each column in the current
        // row
        long[] dp = new long[n];

        // Initialize dp with the values from the first row
        for (int j = 0; j < n; ++j) {
            dp[j] = points[0][j];
        }

        // Iterate over each row starting from the second one
        for (int i = 1; i < m; ++i) {
            // Arrays to store the maximum values while traversing from left to right and
            // right to left
            long[] leftMax = new long[n];
            long[] rightMax = new long[n];
            // Array to store the new dp values for the current row
            long[] newDp = new long[n];

            // Calculate the maximum points that can be obtained by moving from left to
            // right
            leftMax[0] = dp[0]; // Initialize the first element
            for (int j = 1; j < n; ++j) {
                // Maximum of the previous leftMax value or the current dp value plus the column
                // index
                leftMax[j] = Math.max(leftMax[j - 1], dp[j] + j);
            }

            // Calculate the maximum points that can be obtained by moving from right to
            // left
            rightMax[n - 1] = dp[n - 1] - (n - 1); // Initialize the last element
            for (int j = n - 2; j >= 0; --j) {
                // Maximum of the next rightMax value or the current dp value minus the column
                // index
                rightMax[j] = Math.max(rightMax[j + 1], dp[j] - j);
            }

            // Calculate the dp values for the current row
            for (int j = 0; j < n; ++j) {
                // The new dp value is the maximum of leftMax[j] - j or rightMax[j] + j plus the
                // points at the current position
                newDp[j] = Math.max(leftMax[j] - j, rightMax[j] + j) + points[i][j];
            }

            // Update dp to the new values calculated for the current row
            dp = newDp;
        }

        // Find the maximum value in the dp array after processing all rows
        long maxPoints = dp[0];
        for (int j = 1; j < n; ++j) {
            maxPoints = Math.max(maxPoints, dp[j]);
        }

        // Return the maximum points that can be obtained
        return maxPoints;
    }
}
