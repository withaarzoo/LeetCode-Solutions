import java.util.HashMap;

// Class representing a single node in the Trie
class TrieNode {
    // HashMap to store child nodes for each character
    HashMap<Character, TrieNode> children = new HashMap<>();

    // Count stores how many times this node has been visited (useful for scoring
    // prefixes)
    int count = 0;
}

// Class representing the Trie (prefix tree) structure
class Trie {
    TrieNode root; // Root node of the Trie

    // Constructor to initialize the Trie with an empty root node
    public Trie() {
        root = new TrieNode();
    }

    // Method to insert a word into the Trie
    public void insert(String word) {
        TrieNode node = root; // Start at the root node

        // Loop through each character in the word
        for (char ch : word.toCharArray()) {
            // If the current character isn't already a child of the current node, add it
            node.children.putIfAbsent(ch, new TrieNode());

            // Move to the child node corresponding to the current character
            node = node.children.get(ch);

            // Increment the count to signify that this node has been visited
            node.count++;
        }
    }

    // Method to get the sum of scores for all prefixes of a given word
    public int getPrefixScoreSum(String word) {
        TrieNode node = root; // Start at the root node
        int scoreSum = 0; // Initialize the sum of prefix scores

        // Loop through each character in the word
        for (char ch : word.toCharArray()) {
            // Move to the child node corresponding to the current character
            node = node.children.get(ch);

            // Add the count of the current node to the score sum
            scoreSum += node.count;
        }

        // Return the total sum of scores for all prefixes of the word
        return scoreSum;
    }
}

// Solution class that solves the problem using the Trie
class Solution {
    // Method to calculate the sum of prefix scores for each word in the array
    public int[] sumPrefixScores(String[] words) {
        Trie trie = new Trie(); // Create a new Trie

        // Insert all words from the array into the Trie
        for (String word : words) {
            trie.insert(word);
        }

        // Create an array to store the result (sum of prefix scores for each word)
        int[] result = new int[words.length];

        // Calculate the sum of prefix scores for each word
        for (int i = 0; i < words.length; i++) {
            result[i] = trie.getPrefixScoreSum(words[i]);
        }

        // Return the result array
        return result;
    }
}