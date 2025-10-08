import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
    sort.Ints(potions)
    m := len(potions)
    res := make([]int, len(spells))

    for i, s := range spells {
        s64 := int64(s)
        // ceil(success / s)
        need := (success + s64 - 1) / s64

        // binary search first index with value >= need
        l, r := 0, m
        for l < r {
            mid := (l + r) / 2
            if int64(potions[mid]) < need {
                l = mid + 1
            } else {
                r = mid
            }
        }
        res[i] = m - l
    }
    return res
}
