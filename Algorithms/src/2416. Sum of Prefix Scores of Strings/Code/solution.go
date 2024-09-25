package main

// TrieNode represents a single node in the Trie structure.
// Each node contains a map of children (to track the next characters) and a count to store how many words share the prefix up to this point.
type TrieNode struct {
    children map[byte]*TrieNode // Map to store the children nodes, where each key is a character (byte)
    count    int                // Count of words sharing this prefix
}

// Trie represents the overall structure of the prefix tree.
// It has a root node that serves as the entry point for all word insertions and prefix queries.
type Trie struct {
    root *TrieNode // Root node of the Trie
}

// newTrieNode creates a new TrieNode with an empty map of children.
// It is used to initialize new nodes during word insertion.
func newTrieNode() *TrieNode {
    return &TrieNode{children: make(map[byte]*TrieNode)} // Initialize with an empty map for children
}

// NewTrie initializes a new Trie with a root node.
// This function is the starting point for creating a Trie structure.
func NewTrie() *Trie {
    return &Trie{root: newTrieNode()} // Create a Trie with a new root node
}

// Insert adds a word into the Trie character by character.
// As each character is inserted, if it doesn't already exist in the children map, a new TrieNode is created.
func (trie *Trie) Insert(word string) {
    node := trie.root // Start from the root node of the Trie
    // Iterate over each character in the word
    for i := 0; i < len(word); i++ {
        ch := word[i] // Get the current character
        // If the character doesn't exist as a child node, create a new TrieNode
        if _, exists := node.children[ch]; !exists {
            node.children[ch] = newTrieNode() // Add a new node for the character
        }
        node = node.children[ch] // Move to the child node representing the current character
        node.count++             // Increment the count to indicate how many words share this prefix
    }
}

// GetPrefixScoreSum calculates the sum of the prefix scores for a given word.
// The score for each prefix is the count of how many words share that prefix in the Trie.
func (trie *Trie) GetPrefixScoreSum(word string) int {
    node := trie.root // Start from the root node of the Trie
    scoreSum := 0     // Variable to accumulate the sum of prefix scores
    
    // Iterate over each character in the word
    for i := 0; i < len(word); i++ {
        ch := word[i]        // Get the current character
        node = node.children[ch] // Move to the child node representing the current character
        scoreSum += node.count   // Add the count of the current node (prefix) to the score sum
    }
    return scoreSum // Return the total sum of prefix scores for the word
}

// sumPrefixScores calculates the sum of prefix scores for each word in the input list.
// It uses the Trie to store the words and then computes the score for each word.
func sumPrefixScores(words []string) []int {
    trie := NewTrie() // Create a new Trie
    
    // Insert all words into the Trie
    for _, word := range words {
        trie.Insert(word) // Add each word to the Trie
    }
    
    // Initialize a result array to store the prefix score sums for each word
    result := make([]int, len(words))
    
    // Calculate the sum of prefix scores for each word
    for i, word := range words {
        result[i] = trie.GetPrefixScoreSum(word) // Compute the prefix score sum for each word
    }
    
    return result // Return the array containing the prefix score sums for all words
}