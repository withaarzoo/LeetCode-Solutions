class Solution:
    def maxBuilding(self, n: int, restrictions: List[List[int]]) -> int:
        # Building 1 must have height 0
        restrictions.append([1, 0])

        # Building n can be at most n - 1
        restrictions.append([n, n - 1])

        # Sort by building index
        restrictions.sort()

        m = len(restrictions)

        # Left to right pass
        for i in range(1, m):
            dist = restrictions[i][0] - restrictions[i - 1][0]

            restrictions[i][1] = min(
                restrictions[i][1],
                restrictions[i - 1][1] + dist
            )

        # Right to left pass
        for i in range(m - 2, -1, -1):
            dist = restrictions[i + 1][0] - restrictions[i][0]

            restrictions[i][1] = min(
                restrictions[i][1],
                restrictions[i + 1][1] + dist
            )

        ans = 0

        # Compute peak for every interval
        for i in range(1, m):
            x1, h1 = restrictions[i - 1]
            x2, h2 = restrictions[i]

            dist = x2 - x1

            peak = max(h1, h2) + (
                dist - abs(h1 - h2)
            ) // 2

            ans = max(ans, peak)

        return ans