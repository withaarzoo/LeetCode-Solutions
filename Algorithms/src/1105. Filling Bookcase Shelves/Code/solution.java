import java.util.Arrays;

class Solution {
    public int minHeightShelves(int[][] books, int shelfWidth) {
        int n = books.length; // Number of books
        int[] dp = new int[n + 1]; // dp[i] represents the minimum height of the bookshelf to place the first i
                                   // books
        Arrays.fill(dp, Integer.MAX_VALUE); // Initialize dp array with a large value, representing unbounded height
        dp[0] = 0; // No books placed means the total height is 0

        // Iterate through each book
        for (int i = 1; i <= n; ++i) {
            int width = 0; // To keep track of the current width of books on the shelf
            int height = 0; // To keep track of the maximum height of books on the current shelf

            // Try placing books in reverse order to find the optimal arrangement
            for (int j = i; j > 0; --j) {
                width += books[j - 1][0]; // Add the width of the current book
                if (width > shelfWidth) // If adding the current book exceeds the shelf width, break the loop
                    break;

                // Update the maximum height for the current shelf
                height = Math.max(height, books[j - 1][1]);

                // Update dp[i] to the minimum possible height by placing the current set of
                // books on the shelf
                dp[i] = Math.min(dp[i], dp[j - 1] + height);
            }
        }

        return dp[n]; // The minimum height of the bookshelf for all n books
    }
}
