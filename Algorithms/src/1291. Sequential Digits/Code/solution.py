class Solution:
    def sequentialDigits(self, low: int, high: int) -> List[int]:

        # String containing all consecutive digits
        digits = "123456789"

        # Result list
        ans = []

        # Number of digits in low and high
        min_len = len(str(low))
        max_len = len(str(high))

        # Try every possible length
        for length in range(min_len, max_len + 1):

            # Generate every substring of current length
            for start in range(10 - length):

                # Convert substring into an integer
                num = int(digits[start:start + length])

                # Keep only numbers inside the range
                if low <= num <= high:
                    ans.append(num)

        return ans