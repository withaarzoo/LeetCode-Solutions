class Solution {
public:
    int maxIncreasingSubarrays(vector<int>& nums) {
        int n = nums.size();
        if (n < 2) return 0;

        vector<int> inc(n, 1);
        for (int i = n - 2; i >= 0; --i)
            inc[i] = (nums[i] < nums[i+1]) ? inc[i+1] + 1 : 1;

        auto feasible = [&](int k) {
            if (k == 0) return true;
            for (int a = 0; a + 2*k <= n; ++a)
                if (inc[a] >= k && inc[a + k] >= k)
                    return true;
            return false;
        };

        int lo = 0, hi = n / 2, ans = 0;
        while (lo <= hi) {
            int mid = (lo + hi) / 2;
            if (feasible(mid)) ans = mid, lo = mid + 1;
            else hi = mid - 1;
        }
        return ans;
    }
};
