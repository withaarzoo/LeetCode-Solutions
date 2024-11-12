import (
    "sort"
)

func maximumBeauty(items [][]int, queries []int) []int {
    sort.Slice(items, func(i, j int) bool {
        return items[i][0] < items[j][0]
    })

    priceBeauty := make([][2]int, 0)
    maxBeauty := 0
    for _, item := range items {
        price, beauty := item[0], item[1]
        maxBeauty = max(maxBeauty, beauty)
        priceBeauty = append(priceBeauty, [2]int{price, maxBeauty})
    }

    result := make([]int, len(queries))
    for i, query := range queries {
        idx := sort.Search(len(priceBeauty), func(j int) bool {
            return priceBeauty[j][0] > query
        }) - 1
        if idx >= 0 {
            result[i] = priceBeauty[idx][1]
        } else {
            result[i] = 0
        }
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}