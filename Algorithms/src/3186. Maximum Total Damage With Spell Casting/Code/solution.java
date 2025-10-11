class Solution {
    public long maximumTotalDamage(int[] power) {
        if (power == null || power.length == 0) return 0L;
        // frequency map
        HashMap<Long, Long> freq = new HashMap<>();
        for (int v : power) freq.put((long)v, freq.getOrDefault((long)v, 0L) + 1L);
        // unique sorted values
        long[] vals = new long[freq.size()];
        int idx = 0;
        for (long k : freq.keySet()) vals[idx++] = k;
        Arrays.sort(vals);
        int m = vals.length;
        long[] valueSum = new long[m];
        for (int i = 0; i < m; ++i) valueSum[i] = vals[i] * freq.get(vals[i]);
        long[] dp = new long[m];
        dp[0] = valueSum[0];
        for (int i = 1; i < m; ++i) {
            long need = vals[i] - 3; // <= need allowed
            // binary search to find last index j with vals[j] <= need
            int lo = 0, hi = i - 1, j = -1;
            while (lo <= hi) {
                int mid = (lo + hi) >>> 1;
                if (vals[mid] <= need) { j = mid; lo = mid + 1; }
                else hi = mid - 1;
            }
            long take = valueSum[i] + (j >= 0 ? dp[j] : 0L);
            long skip = dp[i - 1];
            dp[i] = Math.max(skip, take);
        }
        return dp[m - 1];
    }
}
