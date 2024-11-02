import "strings"

func isCircularSentence(sentence string) bool {
    // Step 1: Split the sentence into words
    words := strings.Split(sentence, " ")
    
    // Step 2: Check adjacent pairs and the circular condition
    for i := 0; i < len(words); i++ {
        lastChar := words[i][len(words[i])-1]
        firstChar := words[(i+1)%len(words)][0]
        if lastChar != firstChar {
            return false
        }
    }
    
    return true
}
