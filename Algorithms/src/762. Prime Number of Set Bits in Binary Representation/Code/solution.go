import "math/bits"

func countPrimeSetBits(left int, right int) int {
    
    primeSet := map[int]bool{
        2:true, 3:true, 5:true, 7:true,
        11:true, 13:true, 17:true, 19:true,
    }
    
    ans := 0
    
    for num := left; num <= right; num++ {
        // Count set bits
        setBits := bits.OnesCount(uint(num))
        
        if primeSet[setBits] {
            ans++
        }
    }
    
    return ans
}