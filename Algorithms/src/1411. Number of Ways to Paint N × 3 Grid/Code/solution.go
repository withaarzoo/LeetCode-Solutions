func numOfWays(n int) int {
    const MOD int = 1e9 + 7

    same := 6
    diff := 6

    for i := 2; i <= n; i++ {
        newSame := (same*3 + diff*2) % MOD
        newDiff := (same*2 + diff*2) % MOD

        same = newSame
        diff = newDiff
    }

    return (same + diff) % MOD
}
