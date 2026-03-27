class Solution:
    def areSimilar(self, mat: List[List[int]], k: int) -> bool:
        m, n = len(mat), len(mat[0])
        
        k %= n
        
        for i in range(m):
            for j in range(n):
                if i % 2 == 0:
                    new_col = (j + k) % n
                else:
                    new_col = (j - k + n) % n
                
                if mat[i][j] != mat[i][new_col]:
                    return False
        
        return True