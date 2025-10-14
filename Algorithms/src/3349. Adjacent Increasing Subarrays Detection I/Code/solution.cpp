#include <vector>
using namespace std;

class Solution {
public:
    bool hasIncreasingSubarrays(vector<int>& nums, int k) {
        int n = nums.size();
        if (2 * k > n) return false; // not enough elements for two adjacent subarrays
        
        // nextInc[i] = number of consecutive increasing adjacent pairs starting at i
        // i.e., max t such that nums[i] < nums[i+1] < ... < nums[i+t]
        vector<int> nextInc(n, 0);
        for (int i = n - 2; i >= 0; --i) {
            if (nums[i] < nums[i + 1]) nextInc[i] = nextInc[i + 1] + 1;
            else nextInc[i] = 0;
        }
        
        // need k-1 consecutive increasing pairs inside each length-k subarray
        int need = k - 1;
        for (int i = 0; i + 2 * k <= n; ++i) {
            if (nextInc[i] >= need && nextInc[i + k] >= need) return true;
        }
        return false;
    }
};
