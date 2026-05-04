class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        n = len(matrix)

        # Step 1: Transpose
        for i in range(n):
            for j in range(i + 1, n):
                # Swap across diagonal
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]

        # Step 2: Reverse each row
        for i in range(n):
            matrix[i].reverse()