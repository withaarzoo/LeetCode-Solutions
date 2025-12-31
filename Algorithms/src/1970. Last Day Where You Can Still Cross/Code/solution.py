class Solution:
    def latestDayToCross(self, row: int, col: int, cells: List[List[int]]) -> int:
        n = row * col
        top, bottom = n, n + 1

        parent = list(range(n + 2))
        rank = [0] * (n + 2)
        grid = [[False] * col for _ in range(row)]

        def find(x):
            if parent[x] != x:
                parent[x] = find(parent[x])
            return parent[x]

        def union(a, b):
            a, b = find(a), find(b)
            if a == b:
                return
            if rank[a] < rank[b]:
                parent[a] = b
            else:
                parent[b] = a
                if rank[a] == rank[b]:
                    rank[a] += 1

        dr = [1, -1, 0, 0]
        dc = [0, 0, 1, -1]

        for d in range(n - 1, -1, -1):
            r, c = cells[d][0] - 1, cells[d][1] - 1
            grid[r][c] = True
            idx = r * col + c

            if r == 0:
                union(idx, top)
            if r == row - 1:
                union(idx, bottom)

            for k in range(4):
                nr, nc = r + dr[k], c + dc[k]
                if 0 <= nr < row and 0 <= nc < col and grid[nr][nc]:
                    union(idx, nr * col + nc)

            if find(top) == find(bottom):
                return d
        return 0
