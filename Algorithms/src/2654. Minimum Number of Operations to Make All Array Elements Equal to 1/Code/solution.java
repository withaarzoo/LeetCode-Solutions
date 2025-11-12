class Solution {
    public int minOperations(int[] nums) {
        int n = nums.length;

        // 1) If there are ones, just turn the others into 1s.
        int ones = 0;
        for (int x : nums)
            if (x == 1)
                ones++;
        if (ones > 0)
            return n - ones;

        // 2) If global gcd > 1, impossible.
        int g = 0;
        for (int x : nums)
            g = gcd(g, x);
        if (g > 1)
            return -1;

        // 3) Shortest subarray with gcd == 1.
        int best = Integer.MAX_VALUE;
        for (int i = 0; i < n; i++) {
            int cur = 0;
            for (int j = i; j < n; j++) {
                cur = gcd(cur, nums[j]);
                if (cur == 1) {
                    best = Math.min(best, j - i + 1);
                    break;
                }
            }
        }
        // 4) Make first 1 (best-1 ops), then spread to all (n-1 ops).
        return (best - 1) + (n - 1);
    }

    private int gcd(int a, int b) {
        while (b != 0) {
            int t = a % b;
            a = b;
            b = t;
        }
        return Math.abs(a);
    }
}
