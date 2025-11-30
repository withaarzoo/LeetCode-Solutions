package main

func minSubarray(nums []int, p int) int {
    n := len(nums)
    total := 0
    for _, x := range nums {
        total = (total + x) % p
    }
    
    need := total
    if need == 0 {
        return 0
    }
    
    // map remainder -> latest index
    lastIndex := make(map[int]int, n*2)
    lastIndex[0] = -1
    
    ans := n
    prefix := 0
    
    for i, x := range nums {
        prefix = (prefix + x) % p
        prefMod := prefix
        
        target := prefMod - need
        if target < 0 {
            target += p
        }
        
        if j, ok := lastIndex[target]; ok {
            if i-j < ans {
                ans = i - j
            }
        }
        
        lastIndex[prefMod] = i
    }
    
    if ans == n {
        return -1
    }
    return ans
}
