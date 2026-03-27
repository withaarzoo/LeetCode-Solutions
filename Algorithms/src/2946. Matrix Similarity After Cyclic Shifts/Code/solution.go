func areSimilar(mat [][]int, k int) bool {
    m := len(mat)
    n := len(mat[0])
    
    k %= n
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            var newCol int
            
            if i%2 == 0 {
                newCol = (j + k) % n
            } else {
                newCol = (j - k + n) % n
            }
            
            if mat[i][j] != mat[i][newCol] {
                return false
            }
        }
    }
    
    return true
}