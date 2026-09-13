class Solution:
    def largestOverlap(self, img1: List[List[int]], img2: List[List[int]]) -> int:
        n = len(img1)
        # collect every coordinate that holds a 1
        A = [(i, j) for i in range(n) for j in range(n) if img1[i][j] == 1]
        B = [(i, j) for i in range(n) for j in range(n) if img2[i][j] == 1]
        # count how many times each possible shift appears
        # shifts range from -(n-1) to +(n-1), so offset by n
        cnt = [[0] * (2 * n) for _ in range(2 * n)]
        best = 0
        for ax, ay in A:
            for bx, by in B:
                dx = bx - ax + n   # make index non-negative
                dy = by - ay + n
                cnt[dx][dy] += 1
                best = max(best, cnt[dx][dy])
        return best