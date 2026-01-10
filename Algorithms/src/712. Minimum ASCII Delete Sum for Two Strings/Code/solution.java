class Solution {
    public int minimumDeleteSum(String s1, String s2) {
        int n = s1.length(), m = s2.length();
        int[] dp = new int[m + 1];

        for (int j = m - 1; j >= 0; j--) {
            dp[j] = dp[j + 1] + s2.charAt(j);
        }

        for (int i = n - 1; i >= 0; i--) {
            int prev = dp[m];
            dp[m] += s1.charAt(i);

            for (int j = m - 1; j >= 0; j--) {
                int temp = dp[j];
                if (s1.charAt(i) == s2.charAt(j)) {
                    dp[j] = prev;
                } else {
                    dp[j] = Math.min(
                        s1.charAt(i) + dp[j],
                        s2.charAt(j) + dp[j + 1]
                    );
                }
                prev = temp;
            }
        }
        return dp[0];
    }
}
