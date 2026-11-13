class Solution {
public:
    long long countNonDecreasingSubarrays(vector<int>& nums, int k) {
        long long ans = 0;
        long long cost = 0;
        deque<pair<int, int>> dq;
        int n = nums.size();
        for (int i = n - 1, j = n - 1; i >= 0; --i) {
            int num = nums[i];
            int count = 1;
            while (!dq.empty() && dq.back().first < num) {
                auto [nextNum, nextCount] = dq.back();
                dq.pop_back();
                count += nextCount;
                cost += 1LL * (num - nextNum) * nextCount;
            }
            dq.emplace_back(num, count);
            while (cost > k) {
                auto [rightmostNum, rightmostCount] = dq.front();
                dq.pop_front();
                cost -= 1LL * (rightmostNum - nums[j]);
                --j;
                if (rightmostCount > 1) {
                    dq.emplace_front(rightmostNum, rightmostCount - 1);
                }
            }
            ans += j - i + 1;
        }
        return ans;
    }
};