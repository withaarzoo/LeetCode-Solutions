#include <vector>
#include <unordered_map>
using namespace std;

class Solution {
public:
    vector<int> intersect(vector<int>& nums1, vector<int>& nums2) {
        vector<int> result;
        
        // Iterate through nums1
        for (int i = 0; i < nums1.size(); i++) {
            // Iterate through nums2 to find matching elements
            for (int j = 0; j < nums2.size(); j++) {
                if (nums1[i] == nums2[j]) { // Found a common element
                    result.push_back(nums1[i]); // Add it to the result vector
                    nums2[j] = -1; // Mark nums2[j] as visited (using -1)
                    break; // Break the inner loop to avoid duplicates in result
                }
            }
        }
        
        return result; // Return the vector containing intersection elements
    }
};
