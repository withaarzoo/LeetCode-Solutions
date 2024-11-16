class Solution:
    def resultsArray(self, nums: List[int], k: int) -> List[int]:
        n = len(nums)
        result = []

        for i in range(n - k + 1):
            subarray = nums[i:i + k]
            sorted_subarray = sorted(subarray)

            # Check if elements are consecutive
            is_consecutive = all(
                sorted_subarray[j] - sorted_subarray[j - 1] == 1 for j in range(1, k)
            )

            # Add the result based on conditions
            if is_consecutive and subarray == sorted_subarray:
                result.append(sorted_subarray[-1])  # Max element
            else:
                result.append(-1)

        return result