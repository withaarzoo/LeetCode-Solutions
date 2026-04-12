class Solution {
    private int[][][] dp;

    // Calculate Manhattan distance between two letters
    private int getDist(int a, int b) {
        // 26 means finger is not placed yet
        if (a == 26 || b == 26)
            return 0;

        int row1 = a / 6, col1 = a % 6;
        int row2 = b / 6, col2 = b % 6;

        return Math.abs(row1 - row2) + Math.abs(col1 - col2);
    }

    private int solve(int idx, int f1, int f2, String word) {
        // If all characters are typed
        if (idx == word.length())
            return 0;

        // Return memoized result
        if (dp[idx][f1][f2] != -1)
            return dp[idx][f1][f2];

        int cur = word.charAt(idx) - 'A';

        // Option 1: Use finger 1
        int useFinger1 = getDist(f1, cur) + solve(idx + 1, cur, f2, word);

        // Option 2: Use finger 2
        int useFinger2 = getDist(f2, cur) + solve(idx + 1, f1, cur, word);

        return dp[idx][f1][f2] = Math.min(useFinger1, useFinger2);
    }

    public int minimumDistance(String word) {
        dp = new int[word.length()][27][27];

        for (int i = 0; i < word.length(); i++) {
            for (int j = 0; j < 27; j++) {
                for (int k = 0; k < 27; k++) {
                    dp[i][j][k] = -1;
                }
            }
        }

        // Both fingers initially not placed
        return solve(0, 26, 26, word);
    }
}