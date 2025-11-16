package main

import "fmt"

func numSub(s string) int {
    const MOD int64 = 1000000007
    var res int64 = 0
    var cnt int64 = 0

    for i := 0; i < len(s); i++ {
        if s[i] == '1' {
            cnt++
        } else {
            res = (res + (cnt*(cnt+1)/2)%MOD) % MOD
            cnt = 0
        }
    }
    res = (res + (cnt*(cnt+1)/2)%MOD) % MOD
    return int(res)
}

func main() {
    // quick manual test
    fmt.Println(numSub("0110111")) // expected 9
}
