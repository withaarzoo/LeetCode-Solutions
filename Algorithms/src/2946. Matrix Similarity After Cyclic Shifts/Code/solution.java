class Solution {
    public boolean areSimilar(int[][] mat, int k) {
        int m = mat.length;
        int n = mat[0].length;
        
        k %= n;
        
        for(int i = 0; i < m; i++) {
            for(int j = 0; j < n; j++) {
                int newCol;
                
                if(i % 2 == 0) {
                    newCol = (j + k) % n;
                } else {
                    newCol = (j - k + n) % n;
                }
                
                if(mat[i][j] != mat[i][newCol]) {
                    return false;
                }
            }
        }
        return true;
    }
}