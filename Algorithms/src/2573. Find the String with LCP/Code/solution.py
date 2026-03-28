class Solution:
    def findTheString(self, lcp: List[List[int]]) -> str:
        n = len(lcp)

        # group[i] = which character group position i belongs to
        group = [-1] * n
        cur_group = 0

        # Build groups
        for i in range(n):
            if group[i] == -1:
                if cur_group == 26:
                    return ""

                group[i] = cur_group
                cur_group += 1

                for j in range(i + 1, n):
                    if lcp[i][j] > 0:
                        group[j] = group[i]

        # Build answer string
        ans = [''] * n
        for i in range(n):
            ans[i] = chr(ord('a') + group[i])

        ans = ''.join(ans)

        # Verify using DP
        dp = [[0] * (n + 1) for _ in range(n + 1)]

        for i in range(n - 1, -1, -1):
            for j in range(n - 1, -1, -1):
                if ans[i] == ans[j]:
                    dp[i][j] = 1 + dp[i + 1][j + 1]

        # Compare matrices
        for i in range(n):
            for j in range(n):
                if dp[i][j] != lcp[i][j]:
                    return ""

        return ans