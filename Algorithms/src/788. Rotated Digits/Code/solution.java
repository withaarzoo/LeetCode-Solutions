class Solution {
    public int rotatedDigits(int n) {
        int count = 0; // total good numbers
        
        for (int i = 1; i <= n; i++) {
            int num = i;
            boolean isValid = true;   // assume valid
            boolean hasChange = false; // check if it changes
            
            while (num > 0) {
                int digit = num % 10; // extract last digit
                
                // invalid digits
                if (digit == 3 || digit == 4 || digit == 7) {
                    isValid = false;
                    break;
                }
                
                // digits that change
                if (digit == 2 || digit == 5 || digit == 6 || digit == 9) {
                    hasChange = true;
                }
                
                num /= 10; // remove last digit
            }
            
            if (isValid && hasChange) {
                count++;
            }
        }
        
        return count;
    }
}