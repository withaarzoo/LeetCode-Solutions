import "sort"

func canBeEqual(s1 string, s2 string) bool {
    // Store even indexed characters
    even1 := []byte{s1[0], s1[2]}
    even2 := []byte{s2[0], s2[2]}
    
    // Store odd indexed characters
    odd1 := []byte{s1[1], s1[3]}
    odd2 := []byte{s2[1], s2[3]}
    
    // Sort all groups
    sort.Slice(even1, func(i, j int) bool {
        return even1[i] < even1[j]
    })
    sort.Slice(even2, func(i, j int) bool {
        return even2[i] < even2[j]
    })
    sort.Slice(odd1, func(i, j int) bool {
        return odd1[i] < odd1[j]
    })
    sort.Slice(odd2, func(i, j int) bool {
        return odd2[i] < odd2[j]
    })
    
    // Compare even and odd groups
    return even1[0] == even2[0] &&
           even1[1] == even2[1] &&
           odd1[0] == odd2[0] &&
           odd1[1] == odd2[1]
}