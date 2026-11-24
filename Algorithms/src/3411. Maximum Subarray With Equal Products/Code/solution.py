class Solution:
    def maxLength(self, nums: List[int]) -> int:
        n = len(nums)
        ans = 0
        for i in range(n):
            prod = 1
            g = 0
            lc = 1
            for j in range(i, n):
                x = nums[j]
                if g == 0:
                    g = x
                else:
                    a, b = g, x
                    while b:
                        a, b = b, a % b
                    g = a
                gg, b = lc, x
                while b:
                    gg, b = b, gg % b
                if lc // gg > (10**18) // x:
                    break
                lc = lc // gg * x
                if prod > (10**18) // x:
                    break
                prod *= x
                if prod == lc * g:
                    ans = max(ans, j - i + 1)
        return ans