class Solution:
    def minimumTotalDistance(self, robot: List[int], factory: List[List[int]]) -> int:
        robot.sort()
        factory.sort()

        n = len(robot)
        m = len(factory)
        INF = float('inf')

        dp = [[-1] * (m + 1) for _ in range(n + 1)]

        def solve(i: int, j: int) -> int:
            # All robots repaired
            if i == n:
                return 0

            # No factories left
            if j == m:
                return INF

            if dp[i][j] != -1:
                return dp[i][j]

            # Skip current factory
            ans = solve(i, j + 1)

            distance = 0
            pos, limit = factory[j]

            # Use current factory for next k robots
            for k in range(limit):
                if i + k >= n:
                    break

                distance += abs(robot[i + k] - pos)

                next_cost = solve(i + k + 1, j + 1)

                if next_cost != INF:
                    ans = min(ans, distance + next_cost)

            dp[i][j] = ans
            return ans

        return solve(0, 0)