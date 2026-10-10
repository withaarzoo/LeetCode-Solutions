class Solution {
    public long minSumSquareDiff(int[] nums1, int[] nums2, int k1, int k2) {
        int n = nums1.length;
        int[] d = new int[n];
        long total = 0;
        int mx = 0;
        for (int i = 0; i < n; ++i) {
            d[i] = Math.abs(nums1[i] - nums2[i]);
            total += d[i];
            mx = Math.max(mx, d[i]);
        }
        long k = (long)k1 + k2;
        if (total <= k) return 0;
        int left = 0, right = mx;
        while (left < right) {
            int mid = left + (right - left) / 2;
            long need = 0;
            for (int v : d) need += Math.max(0, v - mid);
            if (need <= k) right = mid;
            else left = mid + 1;
        }
        for (int i = 0; i < n; ++i) {
            k -= Math.max(0, d[i] - left);
            d[i] = Math.min(d[i], left);
        }
        for (int i = 0; i < n && k > 0; ++i) {
            if (d[i] == left) {
                --d[i];
                --k;
            }
        }
        long ans = 0;
        for (int v : d) ans += (long)v * v;
        return ans;
    }
}