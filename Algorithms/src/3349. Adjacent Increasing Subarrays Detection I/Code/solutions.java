import java.util.List;

class Solution {
    public boolean hasIncreasingSubarrays(List<Integer> nums, int k) {
        int n = nums.size();
        if (2 * k > n) return false;

        // nextInc[i] = number of consecutive increasing adjacent pairs starting at i
        int[] nextInc = new int[n];
        nextInc[n - 1] = 0;
        for (int i = n - 2; i >= 0; --i) {
            if (nums.get(i) < nums.get(i + 1)) nextInc[i] = nextInc[i + 1] + 1;
            else nextInc[i] = 0;
        }

        int need = k - 1;
        for (int i = 0; i + 2 * k <= n; ++i) {
            if (nextInc[i] >= need && nextInc[i + k] >= need) return true;
        }
        return false;
    }
}
