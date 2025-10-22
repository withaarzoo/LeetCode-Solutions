import java.util.*;

class Solution {
    public int maxFrequency(int[] nums, int k, int numOperations) {
        int n = nums.length;
        if (n == 0)
            return 0;
        Arrays.sort(nums);

        // frequency map
        Map<Long, Integer> freq = new HashMap<>();
        for (int x : nums)
            freq.put((long) x, freq.getOrDefault((long) x, 0) + 1);

        int ans = 1;

        // Case A: existing values as target
        for (Map.Entry<Long, Integer> entry : freq.entrySet()) {
            long v = entry.getKey();
            int already = entry.getValue();

            long lowVal = v - k;
            long highVal = v + k;
            int L = lowerBound(nums, lowVal);
            int R = upperBound(nums, highVal);
            int totalInRange = R - L;
            int need = totalInRange - already;
            int canFix = Math.min(need, numOperations);
            ans = Math.max(ans, already + canFix);
        }

        // Case B: sliding window for 2*k range
        int l = 0;
        for (int r = 0; r < n; ++r) {
            while (l <= r && (long) nums[r] - nums[l] > 2L * k)
                l++;
            int w = r - l + 1;
            ans = Math.max(ans, Math.min(w, numOperations));
        }

        return ans;
    }

    private int lowerBound(int[] arr, long target) {
        int l = 0, r = arr.length;
        while (l < r) {
            int mid = (l + r) >>> 1;
            if ((long) arr[mid] < target)
                l = mid + 1;
            else
                r = mid;
        }
        return l;
    }

    private int upperBound(int[] arr, long target) {
        int l = 0, r = arr.length;
        while (l < r) {
            int mid = (l + r) >>> 1;
            if ((long) arr[mid] <= target)
                l = mid + 1;
            else
                r = mid;
        }
        return l;
    }
}
