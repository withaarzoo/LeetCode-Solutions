class Solution {
    public int minimumTotal(List<List<Integer>> triangle) {
        int n = triangle.size();
        // dp array initialized to last row values
        int[] dp = new int[n];
        List<Integer> last = triangle.get(n - 1);
        for (int i = 0; i < n; i++) dp[i] = last.get(i);

        // iterate from second-last row up
        for (int i = n - 2; i >= 0; --i) {
            for (int j = 0; j <= i; ++j) {
                dp[j] = triangle.get(i).get(j) + Math.min(dp[j], dp[j + 1]);
            }
        }
        return dp[0];
    }
}
