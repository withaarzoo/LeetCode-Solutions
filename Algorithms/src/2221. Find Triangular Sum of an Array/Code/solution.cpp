#include <vector>
using namespace std;

class Solution {
public:
    int triangularSum(vector<int>& nums) {
        int n = nums.size();
        // Reduce array length one by one
        for (int len = n; len > 1; --len) {
            // compute new values in-place into nums[0..len-2]
            for (int i = 0; i < len - 1; ++i) {
                nums[i] = (nums[i] + nums[i + 1]) % 10;
            }
        }
        return nums[0];
    }
};
