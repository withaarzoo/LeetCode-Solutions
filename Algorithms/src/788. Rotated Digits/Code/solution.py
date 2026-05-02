class Solution:
    def rotatedDigits(self, n: int) -> int:
        count = 0  # total good numbers
        
        for i in range(1, n + 1):
            num = i
            isValid = True   # assume valid
            hasChange = False  # check if it changes
            
            while num > 0:
                digit = num % 10  # extract last digit
                
                # invalid digits
                if digit in [3, 4, 7]:
                    isValid = False
                    break
                
                # digits that change
                if digit in [2, 5, 6, 9]:
                    hasChange = True
                
                num //= 10  # remove last digit
            
            if isValid and hasChange:
                count += 1
        
        return count