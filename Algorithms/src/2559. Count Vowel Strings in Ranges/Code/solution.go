package main

func vowelStrings(words []string, queries [][]int) []int {
    vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
    n := len(words)
    prefix := make([]int, n)

    // Precompute the prefix sum
    for i := 0; i < n; i++ {
        if vowels[words[i][0]] && vowels[words[i][len(words[i])-1]] {
            prefix[i] = 1
        }
        if i > 0 {
            prefix[i] += prefix[i-1]
        }
    }

    // Answer the queries
    result := make([]int, len(queries))
    for i, query := range queries {
        l, r := query[0], query[1]
        if l > 0 {
            result[i] = prefix[r] - prefix[l-1]
        } else {
            result[i] = prefix[r]
        }
    }
    return result
}