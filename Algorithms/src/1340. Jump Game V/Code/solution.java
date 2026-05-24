class Solution {

    int[] dp;

    // DFS function returns maximum jumps starting from index i
    private int dfs(int i, int[] arr, int d) {

        // Return stored answer if already computed
        if (dp[i] != -1)
            return dp[i];

        // At least current index can be visited
        int ans = 1;

        // Move right
        for (int j = i + 1; j <= Math.min(arr.length - 1, i + d); j++) {

            // Stop if blocked
            if (arr[j] >= arr[i])
                break;

            // Update best answer
            ans = Math.max(ans, 1 + dfs(j, arr, d));
        }

        // Move left
        for (int j = i - 1; j >= Math.max(0, i - d); j--) {

            // Stop if blocked
            if (arr[j] >= arr[i])
                break;

            // Update best answer
            ans = Math.max(ans, 1 + dfs(j, arr, d));
        }

        // Store result
        return dp[i] = ans;
    }

    public int maxJumps(int[] arr, int d) {

        int n = arr.length;

        // DP array initialized with -1
        dp = new int[n];

        for (int i = 0; i < n; i++) {
            dp[i] = -1;
        }

        int answer = 1;

        // Try starting from every index
        for (int i = 0; i < n; i++) {
            answer = Math.max(answer, dfs(i, arr, d));
        }

        return answer;
    }
}