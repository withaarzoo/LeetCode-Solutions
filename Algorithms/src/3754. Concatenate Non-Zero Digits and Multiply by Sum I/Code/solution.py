class Solution:
    def sumAndMultiply(self, n: int) -> int:
        # This stores the number formed by all non-zero digits.
        x = 0
        
        # This stores the sum of all non-zero digits.
        digit_sum = 0
        
        # Find the highest place value to read digits left to right.
        divisor = 1
        while n // divisor >= 10:
            divisor *= 10
        
        # Process every digit from left to right.
        while divisor > 0:
            # Extract the digit at the current place value.
            digit = n // divisor
            
            # Remove the current digit from n.
            n %= divisor
            
            # Ignore zero digits because they should not be part of x.
            if digit != 0:
                # Append the current digit to x.
                x = x * 10 + digit
                
                # Add the current digit to the sum.
                digit_sum += digit
            
            # Move to the next smaller place value.
            divisor //= 10
        
        # Return the required product.
        return x * digit_sum