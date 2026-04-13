class Solution {
public:
    int getMinDistance(vector<int>& nums, int target, int start) {
        // Store the minimum distance found so far
        int answer = INT_MAX;

        // Traverse through the array
        for (int i = 0; i < nums.size(); i++) {
            // Check if current element is the target
            if (nums[i] == target) {
                // Update the minimum distance
                answer = min(answer, abs(i - start));
            }
        }

        return answer;
    }
};