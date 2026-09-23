class Solution {
public:
    int minOperations(vector<int>& nums, int x) {
        long long total = 0;
        for (int num : nums) total += num;
        if (total < x) return -1;
        long long target = total - x;
        if (target == 0) return nums.size();
        int n = nums.size();
        int left = 0;
        long long sum = 0;
        int maxLen = -1;
        for (int right = 0; right < n; ++right) {
            sum += nums[right];
            while (sum > target && left <= right) {
                sum -= nums[left];
                ++left;
            }
            if (sum == target) {
                maxLen = max(maxLen, right - left + 1);
            }
        }
        return maxLen == -1 ? -1 : n - maxLen;
    }
};