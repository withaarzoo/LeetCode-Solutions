class Solution {
public:
    int longestCommonPrefix(vector<int>& arr1, vector<int>& arr2) {
        
        // Hash set to store all prefixes from arr1
        unordered_set<int> prefixes;

        // Generate all prefixes from arr1
        for (int num : arr1) {

            int x = num;

            // Keep removing last digit
            while (x > 0) {

                // Store current prefix
                prefixes.insert(x);

                // Remove last digit
                x /= 10;
            }
        }

        int ans = 0;

        // Check every number from arr2
        for (int num : arr2) {

            int x = num;

            // Keep shortening the number
            while (x > 0) {

                // If prefix exists in arr1 prefixes
                if (prefixes.count(x)) {

                    // Convert to string to get digit length
                    ans = max(ans, (int)to_string(x).size());

                    // No need to check smaller prefixes
                    break;
                }

                // Remove last digit
                x /= 10;
            }
        }

        return ans;
    }
};