class Solution:
    def findRotation(self, mat, target):
        
        def rotate(mat):
            n = len(mat)
            # Transpose
            for i in range(n):
                for j in range(i, n):
                    mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
            
            # Reverse each row
            for row in mat:
                row.reverse()

        for _ in range(4):
            if mat == target:
                return True
            rotate(mat)

        return False