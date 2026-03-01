class Solution:
    def minPartitions(self, n: str) -> int:
        max_digit = 0  # Store maximum digit
        
        for ch in n:
            digit = int(ch)  # Convert character to integer
            
            if digit > max_digit:
                max_digit = digit
            
            # Early stopping if 9 is found
            if max_digit == 9:
                break
        
        return max_digit