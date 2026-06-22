func maxNumberOfBalloons(text string) int {
    // Store frequency of all lowercase letters
    freq := make([]int, 26)

    // Count each character
    for _, ch := range text {
        freq[ch-'a']++
    }

    // Start with count of 'b'
    ans := freq['b'-'a']

    // Update answer with the minimum possible value
    if freq['a'-'a'] < ans {
        ans = freq['a'-'a']
    }

    if freq['l'-'a']/2 < ans {
        ans = freq['l'-'a'] / 2
    }

    if freq['o'-'a']/2 < ans {
        ans = freq['o'-'a'] / 2
    }

    if freq['n'-'a'] < ans {
        ans = freq['n'-'a']
    }

    return ans
}