/**
 * @param {number} n
 * @return {number}
 */
var rotatedDigits = function(n) {
    let count = 0; // total good numbers
    
    for (let i = 1; i <= n; i++) {
        let num = i;
        let isValid = true;   // assume valid
        let hasChange = false; // track if it changes
        
        while (num > 0) {
            let digit = num % 10; // get last digit
            
            // invalid digits
            if (digit === 3 || digit === 4 || digit === 7) {
                isValid = false;
                break;
            }
            
            // changing digits
            if (digit === 2 || digit === 5 || digit === 6 || digit === 9) {
                hasChange = true;
            }
            
            num = Math.floor(num / 10); // remove last digit
        }
        
        if (isValid && hasChange) {
            count++;
        }
    }
    
    return count;
};