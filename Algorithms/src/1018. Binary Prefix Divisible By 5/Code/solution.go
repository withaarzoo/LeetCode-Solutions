package main

func prefixesDivBy5(nums []int) []bool {
    ans := make([]bool, len(nums))
    rem := 0 // remainder of current prefix modulo 5

    for i, bit := range nums {
        // binary operation: num = num*2 + bit
        rem = (rem*2 + bit) % 5
        // store if divisible by 5
        ans[i] = (rem == 0)
    }

    return ans
}
