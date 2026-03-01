class Solution {
public:
    int minPartitions(string n) {
        int maxDigit = 0;  // This will store the maximum digit
        
        // Traverse each character in the string
        for(char c : n) {
            int digit = c - '0';  // Convert char to integer
            
            // Update maximum digit
            if(digit > maxDigit) {
                maxDigit = digit;
            }
            
            // Optimization: if we find 9, we can stop early
            if(maxDigit == 9) {
                break;
            }
        }
        
        return maxDigit;
    }
};