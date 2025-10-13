package main

import (
	"sort"
	"strings"
)

// helper: returns sorted-character signature of s
func sortedSig(s string) string {
    // convert to slice of bytes, sort, and return string
    b := []byte(s)
    sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
    return string(b)
}

func removeAnagrams(words []string) []string {
    var result []string
    prevSig := ""
    for _, w := range words {
        sig := sortedSig(w)
        if sig != prevSig {
            result = append(result, w)
            prevSig = sig
        }
    }
    return result
}
