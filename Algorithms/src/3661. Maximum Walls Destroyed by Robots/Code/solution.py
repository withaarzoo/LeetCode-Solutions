class Solution:
    def maxWalls(self, robots: List[int], distance: List[int], walls: List[int]) -> int:
        n = len(robots)

        arr = sorted(zip(robots, distance))
        walls.sort()

        # Dummy robot
        arr.append((10**9, 0))

        def count_walls(left, right):
            if left > right:
                return 0

            return bisect_right(walls, right) - bisect_left(walls, left)

        dp = [[0, 0] for _ in range(n)]

        # First robot shoots left
        dp[0][0] = count_walls(arr[0][0] - arr[0][1], arr[0][0])

        # First robot shoots right
        if n == 1:
            first_right_end = arr[0][0] + arr[0][1]
        else:
            first_right_end = min(arr[0][0] + arr[0][1], arr[1][0] - 1)

        dp[0][1] = count_walls(arr[0][0], first_right_end)

        for i in range(1, n):
            pos, dist = arr[i]

            # Shoot right
            right_end = min(pos + dist, arr[i + 1][0] - 1)
            right_walls = count_walls(pos, right_end)

            dp[i][1] = max(dp[i - 1][0], dp[i - 1][1]) + right_walls

            # Shoot left
            left_start = max(pos - dist, arr[i - 1][0] + 1)
            left_walls = count_walls(left_start, pos)

            dp[i][0] = dp[i - 1][0] + left_walls

            prev_right_end = min(arr[i - 1][0] + arr[i - 1][1], pos - 1)

            overlap_start = left_start
            overlap_end = min(prev_right_end, pos - 1)

            overlap_walls = count_walls(overlap_start, overlap_end)

            dp[i][0] = max(dp[i][0], dp[i - 1][1] + left_walls - overlap_walls)

        return max(dp[n - 1][0], dp[n - 1][1])