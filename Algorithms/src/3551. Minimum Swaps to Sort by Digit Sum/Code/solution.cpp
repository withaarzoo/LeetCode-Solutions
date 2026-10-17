class Solution {
public:
    int minSwaps(vector<int>& nums) {
        auto digitSum = [](int x) {
            int s = 0;
            while (x) {
                s += x % 10;
                x /= 10;
            }
            return s;
        };
        int n = nums.size();
        vector<int> sortedNums = nums;
        sort(sortedNums.begin(), sortedNums.end(), [&](int a, int b) {
            int sa = digitSum(a), sb = digitSum(b);
            if (sa != sb) return sa < sb;
            return a < b;
        });
        unordered_map<int, int> pos;
        for (int i = 0; i < n; i++) pos[sortedNums[i]] = i;
        vector<bool> visited(n, false);
        int cycles = 0;
        for (int i = 0; i < n; i++) {
            if (!visited[i]) {
                cycles++;
                int j = i;
                while (!visited[j]) {
                    visited[j] = true;
                    j = pos[nums[j]];
                }
            }
        }
        return n - cycles;
    }
};