class Solution {
public:
    long long maxAlternatingSum(vector<int>& nums) {
        sort(nums.begin(), nums.end(), [](int a, int b) {
            return abs(a) > abs(b);
        });
        long long ans = 0;
        for (int i = 0; i < (int)nums.size(); ++i) {
            long long sq = 1LL * nums[i] * nums[i];
            if (i % 2 == 0) ans += sq;
            else ans -= sq;
        }
        return ans;
    }
};