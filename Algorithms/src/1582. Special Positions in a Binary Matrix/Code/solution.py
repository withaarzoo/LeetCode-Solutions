class Solution:
    def numSpecial(self, mat: List[List[int]]) -> int:
        m = len(mat)
        n = len(mat[0])

        rowCount = [0] * m
        colCount = [0] * n

        # Count number of 1s in each row and column
        for i in range(m):
            for j in range(n):
                if mat[i][j] == 1:
                    rowCount[i] += 1
                    colCount[j] += 1

        special = 0

        # Check for special positions
        for i in range(m):
            for j in range(n):
                if mat[i][j] == 1 and rowCount[i] == 1 and colCount[j] == 1:
                    special += 1

        return special