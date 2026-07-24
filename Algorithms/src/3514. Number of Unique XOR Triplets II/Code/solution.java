class Solution {
    public int uniqueXorTriplets(int[] nums) {
        final int MAX_XOR = 2048;

        // Store whether a value exists in the array.
        boolean[] present = new boolean[MAX_XOR];
        for (int x : nums) {
            present[x] = true;
        }

        // Initially only XOR = 0 is reachable.
        boolean[] dp = new boolean[MAX_XOR];
        dp[0] = true;

        // Pick exactly 3 values.
        for (int step = 0; step < 3; step++) {
            boolean[] next = new boolean[MAX_XOR];

            // Extend every reachable XOR.
            for (int cur = 0; cur < MAX_XOR; cur++) {
                if (!dp[cur])
                    continue;

                // Try every value that exists.
                for (int v = 0; v < MAX_XOR; v++) {
                    if (present[v]) {
                        next[cur ^ v] = true;
                    }
                }
            }

            dp = next;
        }

        // Count unique XOR values.
        int ans = 0;
        for (boolean ok : dp) {
            if (ok)
                ans++;
        }

        return ans;
    }
}