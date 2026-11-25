class Solution {
    public int smallestUniqueSubarray(int[] nums) {
        int n = nums.length;
        final long MOD1 = 1000000007L, MOD2 = 1000000009L;
        final long BASE1 = 1000003L, BASE2 = 1000033L;
        long[] h1 = new long[n + 1], h2 = new long[n + 1];
        long[] p1 = new long[n + 1], p2 = new long[n + 1];
        p1[0] = p2[0] = 1;
        for (int i = 0; i < n; i++) {
            h1[i + 1] = (h1[i] * BASE1 + nums[i]) % MOD1;
            h2[i + 1] = (h2[i] * BASE2 + nums[i]) % MOD2;
            p1[i + 1] = p1[i] * BASE1 % MOD1;
            p2[i + 1] = p2[i] * BASE2 % MOD2;
        }
        int lo = 1, hi = n, ans = n;
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;
            if (hasUnique(nums, mid, h1, h2, p1, p2, MOD1, MOD2)) {
                ans = mid;
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        return ans;
    }
    private boolean hasUnique(int[] nums, int len, long[] h1, long[] h2, long[] p1, long[] p2, long MOD1, long MOD2) {
        Map<Long, Integer> cnt = new HashMap<>();
        int n = nums.length;
        for (int i = 0; i + len <= n; i++) {
            long a = (h1[i + len] - h1[i] * p1[len] % MOD1 + MOD1) % MOD1;
            long b = (h2[i + len] - h2[i] * p2[len] % MOD2 + MOD2) % MOD2;
            long key = a * MOD2 + b;
            cnt.put(key, cnt.getOrDefault(key, 0) + 1);
        }
        for (int v : cnt.values()) if (v == 1) return true;
        return false;
    }
}