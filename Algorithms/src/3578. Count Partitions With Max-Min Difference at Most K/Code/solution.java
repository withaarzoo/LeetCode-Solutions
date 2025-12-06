class Solution {
    public int countPartitions(int[] nums, int k) {
        final int MOD = 1_000_000_007;
        int n = nums.length;

        long[] dp = new long[n + 1];
        long[] pref = new long[n + 1];

        dp[0] = 1;
        pref[0] = 1;

        Deque<Integer> maxdq = new ArrayDeque<>();
        Deque<Integer> mindq = new ArrayDeque<>();

        int l = 0;

        for (int r = 0; r < n; r++) {
            int x = nums[r];

            // max deque: decreasing
            while (!maxdq.isEmpty() && nums[maxdq.peekLast()] <= x) {
                maxdq.pollLast();
            }
            maxdq.offerLast(r);

            // min deque: increasing
            while (!mindq.isEmpty() && nums[mindq.peekLast()] >= x) {
                mindq.pollLast();
            }
            mindq.offerLast(r);

            // shrink window until valid
            while (!maxdq.isEmpty() && !mindq.isEmpty()
                    && (long) nums[maxdq.peekFirst()] - nums[mindq.peekFirst()] > k) {
                if (maxdq.peekFirst() == l)
                    maxdq.pollFirst();
                if (mindq.peekFirst() == l)
                    mindq.pollFirst();
                l++;
            }

            int L = l;
            int i = r + 1;

            long ways = pref[i - 1];
            if (L > 0)
                ways -= pref[L - 1];
            ways %= MOD;
            if (ways < 0)
                ways += MOD;

            dp[i] = ways;
            pref[i] = (pref[i - 1] + dp[i]) % MOD;
        }

        return (int) dp[n];
    }
}
