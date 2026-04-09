class Solution {
    private static final int MOD = 1_000_000_007;

    // Fast exponentiation
    private long modPow(long base, long exp) {
        long result = 1;

        while (exp > 0) {
            if ((exp & 1) == 1) {
                result = (result * base) % MOD;
            }

            base = (base * base) % MOD;
            exp >>= 1;
        }

        return result;
    }

    // Modular inverse
    private long modInverse(long x) {
        return modPow(x, MOD - 2);
    }

    public int xorAfterQueries(int[] nums, int[][] queries) {
        int n = nums.length;
        int limit = (int) Math.sqrt(n) + 1;

        Map<Integer, List<int[]>> smallQueries = new HashMap<>();

        for (int[] q : queries) {
            int l = q[0];
            int r = q[1];
            int k = q[2];
            int v = q[3];

            // Large k -> process directly
            if (k >= limit) {
                for (int i = l; i <= r; i += k) {
                    nums[i] = (int) ((1L * nums[i] * v) % MOD);
                }
            } else {
                smallQueries.computeIfAbsent(k, x -> new ArrayList<>()).add(q);
            }
        }

        for (Map.Entry<Integer, List<int[]>> entry : smallQueries.entrySet()) {
            int k = entry.getKey();
            List<int[]> group = entry.getValue();

            long[] diff = new long[n];
            Arrays.fill(diff, 1L);

            for (int[] q : group) {
                int l = q[0];
                int r = q[1];
                int v = q[3];

                diff[l] = (diff[l] * v) % MOD;

                int steps = (r - l) / k;
                int nextPos = l + (steps + 1) * k;

                if (nextPos < n) {
                    diff[nextPos] = (diff[nextPos] * modInverse(v)) % MOD;
                }
            }

            for (int i = 0; i < n; i++) {
                if (i >= k) {
                    diff[i] = (diff[i] * diff[i - k]) % MOD;
                }

                nums[i] = (int) ((1L * nums[i] * diff[i]) % MOD);
            }
        }

        int answer = 0;

        for (int num : nums) {
            answer ^= num;
        }

        return answer;
    }
}