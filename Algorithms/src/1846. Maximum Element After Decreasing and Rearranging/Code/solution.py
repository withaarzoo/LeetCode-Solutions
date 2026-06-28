class Solution:
    def maximumElementAfterDecrementingAndRearranging(self, arr: List[int]) -> int:

        # Sort the array so smaller values come first.
        arr.sort()

        # The first element must always be 1.
        arr[0] = 1

        # Build the largest valid sequence.
        for i in range(1, len(arr)):

            # The current value cannot be greater than previous + 1.
            arr[i] = min(arr[i], arr[i - 1] + 1)

        # The last element is the maximum possible value.
        return arr[-1]