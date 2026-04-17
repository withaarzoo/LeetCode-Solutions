class Solution {
public:
    // Function to reverse digits of a number
    int reverseNum(int x) {
        int rev = 0;
        
        while (x > 0) {
            rev = rev * 10 + (x % 10);
            x /= 10;
        }
        
        return rev;
    }

    int minMirrorPairDistance(vector<int>& nums) {
        unordered_map<int, int> lastIndex;
        int ans = INT_MAX;

        for (int i = 0; i < nums.size(); i++) {
            // If current number already exists in map,
            // it means we found a mirror pair
            if (lastIndex.count(nums[i])) {
                ans = min(ans, i - lastIndex[nums[i]]);
            }

            // Store reverse(nums[i]) with current index
            int rev = reverseNum(nums[i]);
            lastIndex[rev] = i;
        }

        return (ans == INT_MAX) ? -1 : ans;
    }
};