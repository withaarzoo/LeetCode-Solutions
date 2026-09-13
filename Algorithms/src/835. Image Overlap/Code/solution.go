func largestOverlap(img1 [][]int, img2 [][]int) int {
    n := len(img1)
    // collect every coordinate that holds a 1
    type pos struct{ x, y int }
    var A, B []pos
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if img1[i][j] == 1 {
                A = append(A, pos{i, j})
            }
            if img2[i][j] == 1 {
                B = append(B, pos{i, j})
            }
        }
    }
    // count how many times each possible shift appears
    // shifts range from -(n-1) to +(n-1), so offset by n
    cnt := make([][]int, 2*n)
    for i := range cnt {
        cnt[i] = make([]int, 2*n)
    }
    best := 0
    for _, a := range A {
        for _, b := range B {
            dx := b.x - a.x + n // make index non-negative
            dy := b.y - a.y + n
            cnt[dx][dy]++
            if cnt[dx][dy] > best {
                best = cnt[dx][dy]
            }
        }
    }
    return best
}