class Solution {
    // Inner class representing a node in our Trie
    class TrieNode {
        TrieNode[] children = new TrieNode[26];
        int bestLen = Integer.MAX_VALUE;
        int bestIdx = Integer.MAX_VALUE;
    }

    public int[] stringIndices(String[] wordsContainer, String[] wordsQuery) {
        TrieNode root = new TrieNode();

        // Build the Trie with the container words
        for (int i = 0; i < wordsContainer.length; i++) {
            String word = wordsContainer[i];
            int len = word.length();
            TrieNode curr = root;

            // Track the overall shortest word at the root for 0-length matches
            if (len < curr.bestLen || (len == curr.bestLen && i < curr.bestIdx)) {
                curr.bestLen = len;
                curr.bestIdx = i;
            }

            // Insert word backwards
            for (int j = len - 1; j >= 0; j--) {
                int charIdx = word.charAt(j) - 'a';

                // Create child node if it doesn't exist
                if (curr.children[charIdx] == null) {
                    curr.children[charIdx] = new TrieNode();
                }

                // Step into the child node
                curr = curr.children[charIdx];

                // Update the running best for this specific suffix path
                if (len < curr.bestLen || (len == curr.bestLen && i < curr.bestIdx)) {
                    curr.bestLen = len;
                    curr.bestIdx = i;
                }
            }
        }

        int[] ans = new int[wordsQuery.length];

        // Find the best match for each query
        for (int i = 0; i < wordsQuery.length; i++) {
            String query = wordsQuery[i];
            int len = query.length();
            TrieNode curr = root;

            // Traverse backwards as far as possible
            for (int j = len - 1; j >= 0; j--) {
                int charIdx = query.charAt(j) - 'a';
                // If path breaks, longest common suffix is found
                if (curr.children[charIdx] == null) {
                    break;
                }
                curr = curr.children[charIdx];
            }
            // Save the stored best index
            ans[i] = curr.bestIdx;
        }

        return ans;
    }
}