class Solution:
    def smallestNumber(self, n: int, t: int) -> int:
        # Keep checking numbers starting from n
        while True:
            product = 1
            x = n

            # Calculate the product of all digits
            while x > 0:
                product *= x % 10
                x //= 10

            # Return the first valid number
            if product % t == 0:
                return n

            # Try the next number
            n += 1