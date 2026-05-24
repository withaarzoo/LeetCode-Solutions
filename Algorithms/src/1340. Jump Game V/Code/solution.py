class Solution:
    def maxJumps(self, arr: List[int], d: int) -> int:

        n = len(arr)

        # dp[i] stores maximum jumps starting from i
        dp = [-1] * n

        # DFS function
        def dfs(i):

            # Return stored answer if already calculated
            if dp[i] != -1:
                return dp[i]

            # Minimum answer is 1 (current index)
            ans = 1

            # Move right
            for j in range(i + 1, min(n, i + d + 1)):

                # Stop if blocked
                if arr[j] >= arr[i]:
                    break

                # Update best answer
                ans = max(ans, 1 + dfs(j))

            # Move left
            for j in range(i - 1, max(-1, i - d - 1), -1):

                # Stop if blocked
                if arr[j] >= arr[i]:
                    break

                # Update best answer
                ans = max(ans, 1 + dfs(j))

            # Store answer
            dp[i] = ans

            return ans

        answer = 1

        # Try every starting index
        for i in range(n):
            answer = max(answer, dfs(i))

        return answer