class Solution:
    def simulate(self, nums: list[int], start: int, dir: int) -> bool:
        n = len(nums)
        a = nums.copy()   # copy so original remains unchanged
        curr = start
        while 0 <= curr < n:
            if a[curr] == 0:
                curr += dir    # move same direction
            else:
                a[curr] -= 1   # decrement
                dir = -dir     # reverse direction
                curr += dir    # step in new direction
        # check all zero
        return all(v == 0 for v in a)

    def countValidSelections(self, nums: list[int]) -> int:
        n = len(nums)
        ans = 0
        for i in range(n):
            if nums[i] != 0:
                continue
            if self.simulate(nums, i, -1):
                ans += 1   # start i, go left
            if self.simulate(nums, i, +1):
                ans += 1   # start i, go right
        return ans
