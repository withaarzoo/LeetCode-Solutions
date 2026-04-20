class Solution:
    def maxDistance(self, colors: List[int]) -> int:
        n = len(colors)
        ans = 0

        # Check farthest house from the first house
        for i in range(n - 1, -1, -1):
            if colors[i] != colors[0]:
                ans = max(ans, i)
                break

        # Check farthest house from the last house
        for i in range(n):
            if colors[i] != colors[n - 1]:
                ans = max(ans, n - 1 - i)
                break

        return ans