class Solution {
public:
    int maxDistance(vector<int>& nums1, vector<int>& nums2) {
        int i = 0, j = 0;
        int ans = 0;

        // Traverse both arrays
        while (i < nums1.size() && j < nums2.size()) {
            
            // Make sure i <= j
            if (i > j) {
                j++;
                continue;
            }

            // Valid pair found
            if (nums1[i] <= nums2[j]) {
                ans = max(ans, j - i);
                j++; // Try to get a larger distance
            } 
            else {
                // Invalid pair, move i to smaller value
                i++;
            }
        }

        return ans;
    }
};