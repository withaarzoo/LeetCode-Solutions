// Go (sorted slice + sort.Search)
// Note: deletion from slice costs O(n) so worst-case can be O(n^2). Using a balanced tree would be O(n log n).
package main
import "sort"

func avoidFlood(rains []int) []int {
    n := len(rains)
    ans := make([]int, n)
    for i := range ans { ans[i] = 1 }
    last := map[int]int{} // lake -> last day it rained
    var dry []int         // sorted list of dry day indices

    for i, lake := range rains {
        if lake > 0 {
            ans[i] = -1
            if prev, ok := last[lake]; ok {
                // find first dry day > prev
                j := sort.Search(len(dry), func(k int) bool { return dry[k] > prev })
                if j == len(dry) {
                    return []int{} // impossible
                }
                ans[dry[j]] = lake
                // remove dry[j]
                dry = append(dry[:j], dry[j+1:]...)
            }
            last[lake] = i
        } else {
            dry = append(dry, i)
        }
    }
    return ans
}
