class Solution {
public:
    int sumCounts(vector<int>& nums) {
        int n = nums.size();
        int res = 0;
        for (int i = 0; i < n; ++i) {
            unordered_set<int> s;
            for (int j = i; j < n; ++j) {
                s.insert(nums[j]);
                int d = s.size();
                res += d * d;
            }
        }
        return res;
    }
};