package main

import (
    "strings"
)

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    // Split sentences into words
    words1 := strings.Split(sentence1, " ")
    words2 := strings.Split(sentence2, " ")

    // Ensure words1 is the longer sentence
    if len(words1) < len(words2) {
        words1, words2 = words2, words1
    }

    start, end := 0, 0
    n1, n2 := len(words1), len(words2)

    // Compare from the start
    for start < n2 && words1[start] == words2[start] {
        start++
    }

    // Compare from the end
    for end < n2 && words1[n1-end-1] == words2[n2-end-1] {
        end++
    }

    // Check if the remaining unmatched part is in the middle
    return start+end >= n2
}