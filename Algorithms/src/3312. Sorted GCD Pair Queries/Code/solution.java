class Solution {
    public int[] gcdValues(int[] nums, long[] queries) {

        // Find maximum value.
        int mx = 0;
        for (int x : nums) mx = Math.max(mx, x);

        // Frequency of every value.
        int[] freq = new int[mx + 1];
        for (int x : nums) freq[x]++;

        // exact[g] = pairs having GCD exactly g.
        long[] exact = new long[mx + 1];

        // Process divisors from large to small.
        for (int g = mx; g >= 1; g--) {

            // Count numbers divisible by g.
            long cnt = 0;
            for (int m = g; m <= mx; m += g)
                cnt += freq[m];

            // Total pairs with GCD multiple of g.
            long pairs = cnt * (cnt - 1) / 2;

            // Remove larger exact GCD counts.
            for (int m = g * 2; m <= mx; m += g)
                pairs -= exact[m];

            exact[g] = pairs;
        }

        // Prefix sums.
        long[] prefix = new long[mx + 1];
        for (int g = 1; g <= mx; g++)
            prefix[g] = prefix[g - 1] + exact[g];

        int[] ans = new int[queries.length];

        for (int i = 0; i < queries.length; i++) {

            int l = 1, r = mx;

            // Binary search answer.
            while (l < r) {
                int mid = (l + r) / 2;

                if (prefix[mid] >= queries[i] + 1)
                    r = mid;
                else
                    l = mid + 1;
            }

            ans[i] = l;
        }

        return ans;
    }
}