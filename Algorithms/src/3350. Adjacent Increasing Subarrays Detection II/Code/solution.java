class Solution {
    public int maxIncreasingSubarrays(List<Integer> nums) {
        int n = nums.size();
        if (n < 2) return 0;

        int[] inc = new int[n]; // inc[i] = length of increasing run starting at i
        inc[n - 1] = 1;

        // Step 1: Compute increasing run lengths from right to left
        for (int i = n - 2; i >= 0; --i) {
            if (nums.get(i) < nums.get(i + 1))
                inc[i] = inc[i + 1] + 1;
            else
                inc[i] = 1;
        }

        // Step 2: Helper function to check if k is feasible
        java.util.function.IntPredicate feasible = (k) -> {
            if (k == 0) return true;
            for (int a = 0; a + 2 * k <= n; ++a) {
                if (inc[a] >= k && inc[a + k] >= k)
                    return true;
            }
            return false;
        };

        // Step 3: Binary search for maximum k
        int lo = 0, hi = n / 2, ans = 0;
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;
            if (feasible.test(mid)) {
                ans = mid;
                lo = mid + 1;
            } else {
                hi = mid - 1;
            }
        }
        return ans;
    }
}
