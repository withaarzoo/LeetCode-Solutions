class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        # Traverse from the last digit
        for i in range(len(digits) - 1, -1, -1):
            digits[i] += 1
            
            if digits[i] < 10:  # No carry
                return digits
            
            digits[i] = 0      # Carry continues
        
        # All digits were 9
        return [1] + digits
