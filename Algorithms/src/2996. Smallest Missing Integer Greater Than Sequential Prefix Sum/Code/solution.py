from typing import List

class Solution:
    def missingInteger(self, nums: List[int]) -> int:
        # I start with the first value because nums[0]
        # is always part of the sequential prefix.
        total = nums[0]

        # I scan the remaining values to find the longest sequential prefix.
        for i in range(1, len(nums)):
            # The sequence continues only if the current value
            # is exactly one greater than the previous value.
            if nums[i] == nums[i - 1] + 1:
                total += nums[i]
            else:
                # The sequence breaks here, so I stop the prefix scan.
                break

        # I create a set containing every array value for fast lookup.
        seen = set(nums)

        # I start searching from the sequential prefix sum.
        answer = total

        # If the current number exists in the array, it is not missing,
        # so I try the next number.
        while answer in seen:
            answer += 1

        # The first missing value is the required answer.
        return answer