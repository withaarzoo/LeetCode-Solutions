import java.util.*;

class Solution {
    public int maxFrequency(int[] nums, int k, int numOperations) {
        if (nums.length == 0)
            return 0;
        int mx = Arrays.stream(nums).max().getAsInt();
        int size = mx + k + 2;
        int[] count = new int[size];

        for (int v : nums)
            count[v]++;

        // prefix sums in-place
        for (int i = 1; i < size; ++i)
            count[i] += count[i - 1];

        int ans = 0;
        for (int t = 0; t < size; ++t) {
            int L = Math.max(0, t - k);
            int R = Math.min(size - 1, t + k);
            int total = count[R] - (L > 0 ? count[L - 1] : 0);
            int freq_t = (t > 0) ? (count[t] - count[t - 1]) : count[t];
            int canConvert = total - freq_t;
            int take = Math.min(numOperations, canConvert);
            ans = Math.max(ans, freq_t + take);
        }
        return ans;
    }
}
