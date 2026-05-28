// Define the TrieNode globally to avoid redeclaring inside the function
type TrieNode struct {
    children [26]*TrieNode
    bestLen  int
    bestIdx  int
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
    root := &TrieNode{
        bestLen: 1e9,
        bestIdx: 1e9,
    }
    
    // Construct the Trie from wordsContainer
    for i, word := range wordsContainer {
        n := len(word)
        curr := root
        
        // Keep track of the globally shortest word at the root
        if n < curr.bestLen || (n == curr.bestLen && i < curr.bestIdx) {
            curr.bestLen = n
            curr.bestIdx = i
        }
        
        // Insert the word right-to-left
        for j := n - 1; j >= 0; j-- {
            charIdx := word[j] - 'a'
            
            if curr.children[charIdx] == nil {
                curr.children[charIdx] = &TrieNode{
                    bestLen: 1e9,
                    bestIdx: 1e9,
                }
            }
            
            curr = curr.children[charIdx]
            
            // Record the shortest length and smallest index that flows through here
            if n < curr.bestLen || (n == curr.bestLen && i < curr.bestIdx) {
                curr.bestLen = n
                curr.bestIdx = i
            }
        }
    }
    
    ans := make([]int, len(wordsQuery))
    
    // Process all queries to find their matching index
    for i, query := range wordsQuery {
        curr := root
        n := len(query)
        
        // Go down the Trie right-to-left
        for j := n - 1; j >= 0; j-- {
            charIdx := query[j] - 'a'
            
            if curr.children[charIdx] == nil {
                break
            }
            curr = curr.children[charIdx]
        }
        
        ans[i] = curr.bestIdx
    }
    
    return ans
}