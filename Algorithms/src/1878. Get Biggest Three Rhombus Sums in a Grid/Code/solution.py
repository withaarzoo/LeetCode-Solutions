class Solution:
    def getBiggestThree(self, grid: List[List[int]]) -> List[int]:

        m, n = len(grid), len(grid[0])
        sums = set()

        for r in range(m):
            for c in range(n):

                sums.add(grid[r][c])

                maxSize = min(r, c, m-1-r, n-1-c)

                for k in range(1, maxSize+1):

                    s = 0

                    for i in range(k):
                        s += grid[r-k+i][c+i]

                    for i in range(k):
                        s += grid[r+i][c+k-i]

                    for i in range(k):
                        s += grid[r+k-i][c-i]

                    for i in range(k):
                        s += grid[r-i][c-k+i]

                    sums.add(s)

        res = sorted(sums, reverse=True)
        return res[:3]