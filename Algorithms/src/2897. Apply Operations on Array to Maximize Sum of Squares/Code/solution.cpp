class Solution {
public:
    int maxSum(vector<int>& nums, int k) {
        const int MOD = 1e9 + 7;
        int n = nums.size();
        vector<int> cnt(30, 0);
        for (int x : nums) {
            for (int b = 0; b < 30; ++b) {
                if (x & (1 << b)) cnt[b]++;
            }
        }
        vector<long long> a(n, 0);
        for (int b = 29; b >= 0; --b) {
            int c = cnt[b];
            for (int i = 0; i < c; ++i) {
                a[i] |= (1LL << b);
            }
        }
        long long ans = 0;
        for (int i = 0; i < k; ++i) {
            ans = (ans + a[i] * a[i]) % MOD;
        }
        return (int)ans;
    }
};