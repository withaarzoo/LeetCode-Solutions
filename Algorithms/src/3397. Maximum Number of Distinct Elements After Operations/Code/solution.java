import java.util.*;

class Solution {
    public int maxDistinctElements(int[] nums, int k) {
        int n = nums.length;
        long[][] intervals = new long[n][2];
        for (int i = 0; i < n; ++i) {
            intervals[i][0] = (long)nums[i] - k;
            intervals[i][1] = (long)nums[i] + k;
        }
        Arrays.sort(intervals, (a,b) -> {
            if (a[1] != b[1]) return Long.compare(a[1], b[1]);
            return Long.compare(a[0], b[0]);
        });

        long lastAssigned = Long.MIN_VALUE / 4;
        int ans = 0;
        for (int i = 0; i < n; ++i) {
            long l = intervals[i][0], r = intervals[i][1];
            long assigned = Math.max(l, lastAssigned + 1);
            if (assigned <= r) {
                ans++;
                lastAssigned = assigned;
            }
        }
        return ans;
    }
}
