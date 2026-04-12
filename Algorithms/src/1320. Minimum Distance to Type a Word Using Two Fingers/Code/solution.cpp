class Solution {
public:
    // Calculate Manhattan distance between two letters
    int getDist(int a, int b) {
        // 26 means finger is not placed yet
        if (a == 26 || b == 26) return 0;

        int row1 = a / 6, col1 = a % 6;
        int row2 = b / 6, col2 = b % 6;

        return abs(row1 - row2) + abs(col1 - col2);
    }

    int dp[301][27][27];

    int solve(int idx, int f1, int f2, string &word) {
        // If all characters are typed
        if (idx == word.size()) return 0;

        // Return memoized result
        if (dp[idx][f1][f2] != -1) return dp[idx][f1][f2];

        int cur = word[idx] - 'A';

        // Option 1: Use finger 1
        int useFinger1 = getDist(f1, cur) + solve(idx + 1, cur, f2, word);

        // Option 2: Use finger 2
        int useFinger2 = getDist(f2, cur) + solve(idx + 1, f1, cur, word);

        return dp[idx][f1][f2] = min(useFinger1, useFinger2);
    }

    int minimumDistance(string word) {
        memset(dp, -1, sizeof(dp));

        // Both fingers initially not placed
        return solve(0, 26, 26, word);
    }
};