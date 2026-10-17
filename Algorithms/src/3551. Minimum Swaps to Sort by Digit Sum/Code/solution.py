class Solution:
    def minSwaps(self, nums: List[int]) -> int:
        def digit_sum(x):
            s = 0
            while x:
                s += x % 10
                x //= 10
            return s
        n = len(nums)
        sorted_nums = sorted(nums, key=lambda x: (digit_sum(x), x))
        pos = {val: i for i, val in enumerate(sorted_nums)}
        visited = [False] * n
        cycles = 0
        for i in range(n):
            if not visited[i]:
                cycles += 1
                j = i
                while not visited[j]:
                    visited[j] = True
                    j = pos[nums[j]]
        return n - cycles