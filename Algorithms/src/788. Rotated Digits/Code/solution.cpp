class Solution {
public:
    int rotatedDigits(int n) {
        int count = 0; // this will store the number of good numbers
        
        for (int i = 1; i <= n; i++) {
            int num = i;
            bool isValid = true;   // assume number is valid initially
            bool hasChange = false; // to check if at least one digit changes
            
            while (num > 0) {
                int digit = num % 10; // extract last digit
                
                // if digit is invalid after rotation
                if (digit == 3 || digit == 4 || digit == 7) {
                    isValid = false;
                    break; // no need to check further
                }
                
                // if digit changes after rotation
                if (digit == 2 || digit == 5 || digit == 6 || digit == 9) {
                    hasChange = true;
                }
                
                num /= 10; // remove last digit
            }
            
            // count only if valid and has at least one changing digit
            if (isValid && hasChange) {
                count++;
            }
        }
        
        return count;
    }
};