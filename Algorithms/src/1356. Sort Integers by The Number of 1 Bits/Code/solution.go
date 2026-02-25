import (
    "math/bits"
    "sort"
)

func sortByBits(arr []int) []int {
    
    sort.Slice(arr, func(i, j int) bool {
        
        bitsI := bits.OnesCount(uint(arr[i]))
        bitsJ := bits.OnesCount(uint(arr[j]))
        
        if bitsI != bitsJ {
            return bitsI < bitsJ
        }
        
        return arr[i] < arr[j]
    })
    
    return arr
}