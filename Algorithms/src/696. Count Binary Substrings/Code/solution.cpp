class Solution {
public:
    int countBinarySubstrings(string s) {
        int n = s.length();
        
        int prevGroup = 0;      // length of previous group
        int currGroup = 1;      // length of current group (start with 1)
        int result = 0;         // final answer
        
        for (int i = 1; i < n; i++) {
            
            if (s[i] == s[i - 1]) {
                // Same character, increase current group size
                currGroup++;
            } else {
                // Character changed, so we finish one group
                
                // Add min of previous and current group
                result += min(prevGroup, currGroup);
                
                // Update previous group
                prevGroup = currGroup;
                
                // Reset current group
                currGroup = 1;
            }
        }
        
        // Add the last comparison
        result += min(prevGroup, currGroup);
        
        return result;
    }
};
