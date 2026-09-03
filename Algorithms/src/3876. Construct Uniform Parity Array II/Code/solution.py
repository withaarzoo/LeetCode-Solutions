class Solution:
    def uniformArray(self, nums1: list[int]) -> bool:
        # Use infinity to represent that no odd or even value has been found yet.
        min_odd = float('inf')
        min_even = float('inf')

        # Scan the array once because only the smallest values of each parity matter.
        for x in nums1:
            if x % 2 == 0:
                # The smallest even value is the hardest one to convert to odd.
                min_even = min(min_even, x)
            else:
                # The smallest odd value is the best value to subtract.
                min_odd = min(min_odd, x)

        # If there is no odd number, all elements are already even.
        if min_odd == float('inf'):
            return True

        # Every even value must be greater than the smallest odd value.
        # Then subtracting min_odd makes that even value odd.
        return min_odd < min_even