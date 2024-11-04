func compressedString(word string) string {
    var comp string
    count := 1
    
    for i := 1; i <= len(word); i++ {
        if i == len(word) || word[i] != word[i-1] || count == 9 {
            comp += fmt.Sprintf("%d%c", count, word[i-1])
            count = 1
        } else {
            count++
        }
    }
    
    return comp
}
