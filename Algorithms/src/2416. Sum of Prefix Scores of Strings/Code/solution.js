// Definition of the TrieNode class, which represents a node in the Trie.
class TrieNode {
  constructor() {
    // 'children' stores the next character nodes in the Trie.
    // It is an object (hashmap) where each key is a character and the value is another TrieNode.
    this.children = {};
    // 'count' keeps track of how many times this node has been visited during insertions.
    this.count = 0;
  }
}

// Definition of the Trie class, which encapsulates the root node and Trie operations.
class Trie {
  constructor() {
    // Initialize the root node of the Trie as an empty TrieNode.
    this.root = new TrieNode();
  }

  // Method to insert a word into the Trie.
  insert(word) {
    // Start traversing from the root of the Trie.
    let node = this.root;

    // Traverse each character in the word.
    for (let ch of word) {
      // If the current character is not already a child of the current node, create a new TrieNode.
      if (!node.children[ch]) {
        node.children[ch] = new TrieNode();
      }
      // Move to the next node (child node corresponding to the current character).
      node = node.children[ch];
      // Increment the count for this node, as this prefix (or part of it) has been inserted.
      node.count++;
    }
  }

  // Method to get the sum of scores for all prefixes of the given word.
  getPrefixScoreSum(word) {
    // Start from the root node of the Trie.
    let node = this.root;
    // Variable to accumulate the score (sum of counts for each prefix).
    let scoreSum = 0;

    // Traverse each character in the word to compute the prefix scores.
    for (let ch of word) {
      // Move to the next node corresponding to the current character.
      node = node.children[ch];
      // Add the count of this node to the score sum (count represents how many times this prefix was seen).
      scoreSum += node.count;
    }

    // Return the accumulated score sum.
    return scoreSum;
  }
}

/**
 * Function to compute the sum of prefix scores for an array of words.
 *
 * @param {string[]} words - The list of words for which we want to calculate prefix scores.
 * @return {number[]} - An array of numbers where each element corresponds to the sum of prefix scores for a word.
 */
var sumPrefixScores = function (words) {
  // Create a new Trie.
  const trie = new Trie();

  // Insert each word from the input array into the Trie.
  for (let word of words) {
    trie.insert(word);
  }

  // Create an array to store the results (sum of prefix scores for each word).
  const result = [];

  // For each word in the input array, calculate the sum of prefix scores.
  for (let word of words) {
    result.push(trie.getPrefixScoreSum(word)); // Get the prefix score sum and store it in the result array.
  }

  // Return the result array containing the sum of prefix scores for each word.
  return result;
};
