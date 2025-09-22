package main

func maxFrequencyElements(nums []int) int {
    freq := make(map[int]int)
    // Build frequency map
    for _, x := range nums {
        freq[x]++
    }

    // Find maximum frequency
    maxFreq := 0
    for _, v := range freq {
        if v > maxFreq {
            maxFreq = v
        }
    }

    // Sum frequencies of those elements that have frequency == maxFreq
    result := 0
    for _, v := range freq {
        if v == maxFreq {
            result += v
        }
    }
    return result
}
