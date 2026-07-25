class Solution:
    def maxProduct(self, n: int) -> int:

        # Store the largest digit
        first = 0

        # Store the second largest digit
        second = 0

        # Process every digit
        while n > 0:

            # Extract the last digit
            digit = n % 10

            # Update the two largest digits
            if digit >= first:
                second = first
                first = digit
            elif digit > second:
                second = digit

            # Remove the last digit
            n //= 10

        # Return the maximum product
        return first * second