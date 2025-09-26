#include <vector>
#include <algorithm>
using namespace std;

class Solution {
public:
    int triangleNumber(vector<int>& nums) {
        int n = nums.size();
        if (n < 3) return 0;
        sort(nums.begin(), nums.end());            // sort ascending
        int count = 0;
        // k is index for the largest side (c)
        for (int k = n - 1; k >= 2; --k) {
            int i = 0, j = k - 1;                  // two pointers for a and b
            while (i < j) {
                // if smallest + current largest > c, then all values from i..j-1 with j work
                if (nums[i] + nums[j] > nums[k]) {
                    count += (j - i);             // add all pairs (i..j-1, j)
                    --j;                          // try smaller b
                } else {
                    ++i;                          // need larger a
                }
            }
        }
        return count;
    }
};
