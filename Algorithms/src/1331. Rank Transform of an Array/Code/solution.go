import "sort"

func arrayRankTransform(arr []int) []int {
    if len(arr) == 0 {
        return []int{}
    }
    
    // Step 1: Create a sorted copy of the array
    sortedArr := make([]int, len(arr))
    copy(sortedArr, arr)
    sort.Ints(sortedArr)
    
    // Step 2: Create a map to assign ranks
    rankMap := make(map[int]int)
    rank := 1
    
    // Step 3: Assign ranks to sorted elements
    for _, num := range sortedArr {
        if _, exists := rankMap[num]; !exists {
            rankMap[num] = rank
            rank++
        }
    }
    
    // Step 4: Replace each element in the original array with its rank
    for i, num := range arr {
        arr[i] = rankMap[num]
    }
    
    return arr
}