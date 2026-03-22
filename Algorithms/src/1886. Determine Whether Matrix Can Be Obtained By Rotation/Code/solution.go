func findRotation(mat [][]int, target [][]int) bool {

    rotate := func(mat [][]int) {
        n := len(mat)

        // Transpose
        for i := 0; i < n; i++ {
            for j := i; j < n; j++ {
                mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
            }
        }

        // Reverse each row
        for i := 0; i < n; i++ {
            for j := 0; j < n/2; j++ {
                mat[i][j], mat[i][n-j-1] = mat[i][n-j-1], mat[i][j]
            }
        }
    }

    isEqual := func(a, b [][]int) bool {
        n := len(a)
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                if a[i][j] != b[i][j] {
                    return false
                }
            }
        }
        return true
    }

    for k := 0; k < 4; k++ {
        if isEqual(mat, target) {
            return true
        }
        rotate(mat)
    }

    return false
}