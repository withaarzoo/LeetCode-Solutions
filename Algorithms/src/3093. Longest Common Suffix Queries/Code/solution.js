/**
 * @param {string[]} wordsContainer
 * @param {string[]} wordsQuery
 * @return {number[]}
 */
var stringIndices = function (wordsContainer, wordsQuery) {
  // Structure of a Trie Node using a Map for better JS object performance
  class TrieNode {
    constructor() {
      this.children = new Array(26).fill(null);
      this.bestLen = Infinity;
      this.bestIdx = Infinity;
    }
  }

  const root = new TrieNode();

  // Populate the Trie
  for (let i = 0; i < wordsContainer.length; i++) {
    const word = wordsContainer[i];
    const len = word.length;
    let curr = root;

    // Check if this word is the overall best fallback
    if (len < curr.bestLen || (len === curr.bestLen && i < curr.bestIdx)) {
      curr.bestLen = len;
      curr.bestIdx = i;
    }

    // Add characters from back to front
    for (let j = len - 1; j >= 0; j--) {
      const charIdx = word.charCodeAt(j) - 97; // 97 is ASCII for 'a'

      if (curr.children[charIdx] === null) {
        curr.children[charIdx] = new TrieNode();
      }

      curr = curr.children[charIdx];

      // Update the node's record of the best matching container string
      if (len < curr.bestLen || (len === curr.bestLen && i < curr.bestIdx)) {
        curr.bestLen = len;
        curr.bestIdx = i;
      }
    }
  }

  const ans = new Array(wordsQuery.length);

  // Answer each query
  for (let i = 0; i < wordsQuery.length; i++) {
    const query = wordsQuery[i];
    const len = query.length;
    let curr = root;

    // Search the Trie in reverse
    for (let j = len - 1; j >= 0; j--) {
      const charIdx = query.charCodeAt(j) - 97;
      if (curr.children[charIdx] === null) {
        break;
      }
      curr = curr.children[charIdx];
    }

    // Grab the pre-calculated best index from the node
    ans[i] = curr.bestIdx;
  }

  return ans;
};
