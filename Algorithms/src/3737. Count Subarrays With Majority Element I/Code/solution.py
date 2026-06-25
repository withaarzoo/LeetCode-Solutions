class Solution:
    def countMajoritySubarrays(self, nums: List[int], target: int) -> int:
        n = len(nums)
        ans = 0

        # Try every possible starting index
        for left in range(n):
            count_target = 0

            # Extend the subarray
            for right in range(left, n):

                # Update target frequency
                if nums[right] == target:
                    count_target += 1

                # Current subarray length
                length = right - left + 1

                # Check majority condition
                if 2 * count_target > length:
                    ans += 1

        return ans