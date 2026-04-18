class Solution:
    def mirrorDistance(self, n: int) -> int:
        rev = 0
        temp = n

        # Reverse the digits of n
        while temp > 0:
            digit = temp % 10       # Get last digit
            rev = rev * 10 + digit  # Add digit to reversed number
            temp //= 10             # Remove last digit

        # Return absolute difference
        return abs(n - rev)