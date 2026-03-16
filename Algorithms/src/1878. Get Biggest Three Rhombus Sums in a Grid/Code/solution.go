func getBiggestThree(grid [][]int) []int {

    m := len(grid)
    n := len(grid[0])

    sums := map[int]bool{}

    for r:=0;r<m;r++{
        for c:=0;c<n;c++{

            sums[grid[r][c]] = true

            maxSize := min(min(r,c), min(m-1-r,n-1-c))

            for k:=1;k<=maxSize;k++{

                sum := 0

                for i:=0;i<k;i++{
                    sum += grid[r-k+i][c+i]
                }

                for i:=0;i<k;i++{
                    sum += grid[r+i][c+k-i]
                }

                for i:=0;i<k;i++{
                    sum += grid[r+k-i][c-i]
                }

                for i:=0;i<k;i++{
                    sum += grid[r-i][c-k+i]
                }

                sums[sum] = true
            }
        }
    }

    res := []int{}

    for k := range sums{
        res = append(res,k)
    }

    sort.Sort(sort.Reverse(sort.IntSlice(res)))

    if len(res) > 3 {
        res = res[:3]
    }

    return res
}

func min(a,b int) int{
    if a<b {return a}
    return b
}