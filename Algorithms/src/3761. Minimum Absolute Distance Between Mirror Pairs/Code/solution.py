class Solution:
    def reverseNum(self, x: int) -> int:
        rev = 0

        while x > 0:
            rev = rev * 10 + (x % 10)
            x //= 10

        return rev

    def minMirrorPairDistance(self, nums: List[int]) -> int:
        last_index = {}
        ans = float('inf')

        for i, num in enumerate(nums):
            # If current number exists in map,
            # then we found a mirror pair
            if num in last_index:
                ans = min(ans, i - last_index[num])

            # Store reverse(num) with current index
            rev = self.reverseNum(num)
            last_index[rev] = i

        return -1 if ans == float('inf') else ans