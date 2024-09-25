# Class representing a node in the Trie
class TrieNode:
    def __init__(self):
        # Dictionary to hold child nodes, where keys are characters
        self.children = {}
        # Count to store how many times this node (prefix) is visited
        self.count = 0

# Class representing the Trie (prefix tree) structure
class Trie:
    def __init__(self):
        # The root node of the Trie
        self.root = TrieNode()

    # Method to insert a word into the Trie
    def insert(self, word: str):
        node = self.root  # Start from the root node
        # Loop through each character of the word
        for ch in word:
            # If the character is not a child of the current node, add it
            if ch not in node.children:
                node.children[ch] = TrieNode()
            # Move to the next node (child node corresponding to the character)
            node = node.children[ch]
            # Increment the count for the current node (this prefix)
            node.count += 1

    # Method to calculate the sum of prefix scores for a given word
    def get_prefix_score_sum(self, word: str) -> int:
        node = self.root  # Start from the root node
        score_sum = 0     # Initialize score sum
        # Loop through each character in the word
        for ch in word:
            # Move to the next node (child node corresponding to the character)
            node = node.children[ch]
            # Add the count of the current node (this prefix) to the score sum
            score_sum += node.count
        return score_sum

# Class to solve the problem of finding prefix scores for a list of words
class Solution:
    def sumPrefixScores(self, words: List[str]) -> List[int]:
        trie = Trie()  # Initialize a new Trie instance
        
        # Insert all words into the Trie to build the structure
        for word in words:
            trie.insert(word)
        
        # For each word, calculate the sum of prefix scores and return the result as a list
        return [trie.get_prefix_score_sum(word) for word in words]