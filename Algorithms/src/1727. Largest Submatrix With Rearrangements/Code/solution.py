class Solution:
    def largestSubmatrix(self, matrix):
        m, n = len(matrix), len(matrix[0])

        # Step 1: Build heights
        for i in range(1, m):
            for j in range(n):
                if matrix[i][j] == 1:
                    matrix[i][j] += matrix[i-1][j]

        max_area = 0

        # Step 2 & 3
        for i in range(m):
            row = sorted(matrix[i], reverse=True)

            for j in range(n):
                area = row[j] * (j + 1)
                max_area = max(max_area, area)

        return max_area