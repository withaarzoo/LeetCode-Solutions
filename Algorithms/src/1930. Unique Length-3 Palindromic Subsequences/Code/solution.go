package main

func countPalindromicSubsequence(s string) int {
    n := len(s)
    const A = 26
    first := make([]int, A)
    last := make([]int, A)
    for i := 0; i < A; i++ {
        first[i] = 1<<30
        last[i] = -1
    }
    // record first and last occurrence
    for i := 0; i < n; i++ {
        c := int(s[i] - 'a')
        if i < first[c] {
            first[c] = i
        }
        if i > last[c] {
            last[c] = i
        }
    }

    ans := 0
    // for each outer letter, count distinct middle letters between first and last
    for c := 0; c < A; c++ {
        if first[c] < last[c] {
            seen := make([]bool, A)
            for i := first[c] + 1; i < last[c]; i++ {
                seen[int(s[i]-'a')] = true
            }
            for j := 0; j < A; j++ {
                if seen[j] {
                    ans++
                }
            }
        }
    }
    return ans
}
