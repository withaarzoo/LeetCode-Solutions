// Go implementation
package main
import (
    "fmt"
)

const MOD int64 = 1000000007

func magicalSum(m int, k int, nums []int) int {
    n := len(nums)
    C := make([][]int64, m+1)
    for i:=0;i<=m;i++ { C[i] = make([]int64, m+1) }
    for i:=0;i<=m;i++ {
        C[i][0] = 1
        C[i][i] = 1
        for j:=1;j<i;j++ {
            C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD
        }
    }
    powA := make([][]int64, n)
    for i:=0;i<n;i++ {
        powA[i] = make([]int64, m+1)
        powA[i][0] = 1
        a := int64(nums[i]) % MOD
        for t:=1;t<=m;t++ {
            powA[i][t] = (powA[i][t-1] * a) % MOD
        }
    }

    M := m
    cur := make([][][]int64, M+1)
    for r:=0;r<=M;r++ {
        cur[r] = make([][]int64, M+1)
        for c:=0;c<=M;c++ {
            cur[r][c] = make([]int64, M+1)
        }
    }
    cur[M][0][0] = 1

    for i:=0;i<n;i++ {
        nxt := make([][][]int64, M+1)
        for r:=0;r<=M;r++ {
            nxt[r] = make([][]int64, M+1)
            for c:=0;c<=M;c++ { nxt[r][c] = make([]int64, M+1) }
        }
        for r:=0;r<=M;r++ {
            for carry:=0; carry<=M; carry++ {
                for ones:=0; ones<=M; ones++ {
                    val := cur[r][carry][ones]
                    if val == 0 { continue }
                    for t:=0; t<=r; t++ {
                        newr := r - t
                        s := carry + t
                        bit := s & 1
                        newones := ones + bit
                        if newones > M { continue }
                        newcarry := s >> 1
                        mult := (C[r][t] * powA[i][t]) % MOD
                        add := (val * mult) % MOD
                        nxt[newr][newcarry][newones] = (nxt[newr][newcarry][newones] + add) % MOD
                    }
                }
            }
        }
        cur = nxt
    }

    var ans int64 = 0
    for carry:=0; carry<=M; carry++ {
        for ones:=0; ones<=M; ones++ {
            val := cur[0][carry][ones]
            if val == 0 { continue }
            extra := popcount(carry)
            if ones + extra == k {
                ans = (ans + val) % MOD
            }
        }
    }
    return int(ans)
}

func popcount(x int) int {
    cnt := 0
    for x>0 {
        if x&1 == 1 { cnt++ }
        x >>= 1
    }
    return cnt
}

func main() {
    fmt.Println(magicalSum(2,2, []int{5,4,3,2,1})) // quick test
}
