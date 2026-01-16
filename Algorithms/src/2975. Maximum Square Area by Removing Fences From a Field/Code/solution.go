func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
    const MOD int64 = 1000000007

    hFences = append(hFences, 1, m)
    vFences = append(vFences, 1, n)

    sort.Ints(hFences)
    sort.Ints(vFences)

    horizontal := make(map[int]bool)
    vertical := make(map[int]bool)

    for i := 0; i < len(hFences); i++ {
        for j := i + 1; j < len(hFences); j++ {
            horizontal[hFences[j]-hFences[i]] = true
        }
    }

    for i := 0; i < len(vFences); i++ {
        for j := i + 1; j < len(vFences); j++ {
            vertical[vFences[j]-vFences[i]] = true
        }
    }

    maxSide := 0
    for d := range horizontal {
        if vertical[d] {
            if d > maxSide {
                maxSide = d
            }
        }
    }

    if maxSide == 0 {
        return -1
    }

    return int((int64(maxSide) * int64(maxSide)) % MOD)
}
