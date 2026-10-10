class Solution {
public:
    long long minSumSquareDiff(vector<int>& nums1, vector<int>& nums2, int k1, int k2) {
        int n = nums1.size();
        vector<int> d(n);
        long long total = 0;
        int mx = 0;
        for (int i = 0; i < n; ++i) {
            d[i] = abs(nums1[i] - nums2[i]);
            total += d[i];
            mx = max(mx, d[i]);
        }
        long long k = (long long)k1 + k2;
        if (total <= k) return 0;
        int left = 0, right = mx;
        while (left < right) {
            int mid = left + (right - left) / 2;
            long long need = 0;
            for (int v : d) need += max(0, v - mid);
            if (need <= k) right = mid;
            else left = mid + 1;
        }
        for (int i = 0; i < n; ++i) {
            k -= max(0, d[i] - left);
            d[i] = min(d[i], left);
        }
        for (int i = 0; i < n && k > 0; ++i) {
            if (d[i] == left) {
                --d[i];
                --k;
            }
        }
        long long ans = 0;
        for (int v : d) ans += (long long)v * v;
        return ans;
    }
};