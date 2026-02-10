class Solution:
    def longestBalanced(self, nums: List[int]) -> int:
        n = len(nums)
        ans = 0

        for i in range(n):
            even_set = set()
            odd_set = set()

            for j in range(i, n):
                if nums[j] % 2 == 0:
                    even_set.add(nums[j])
                else:
                    odd_set.add(nums[j])

                if len(even_set) == len(odd_set):
                    ans = max(ans, j - i + 1)

        return ans
