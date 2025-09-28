#include <vector>
#include <algorithm>

class Solution {
public:
    int largestPerimeter(std::vector<int>& nums) {
        // Sort the side lengths in non-decreasing order
        std::sort(nums.begin(), nums.end());
        int n = nums.size();
        // Check triples from largest side downwards
        for (int i = n - 1; i >= 2; --i) {
            int a = nums[i];       // largest in this triple
            int b = nums[i - 1];
            int c = nums[i - 2];
            // triangle condition: sum of two smaller sides > largest side
            if (b + c > a) return a + b + c;
        }
        return 0;
    }
};
