class Solution:
    def minElement(self, nums: List[int]) -> int:

        # Helper function to calculate digit sum
        def digit_sum(num):
            total = 0

            # Process every digit
            while num > 0:
                total += num % 10  # Add last digit
                num //= 10         # Remove last digit

            return total

        ans = float('inf')

        # Check digit sum of every number
        for num in nums:
            ans = min(ans, digit_sum(num))

        return ans