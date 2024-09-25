#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

// Structure to define a TrieNode
// Each node will store its children (mapped by characters) and the count of words passing through it
struct TrieNode
{
    // A map to store the children nodes corresponding to each character
    unordered_map<char, TrieNode *> children;

    // This variable keeps track of how many words share this node (i.e., pass through this prefix)
    int count = 0;
};

// Class that implements the Trie (prefix tree) data structure
class Trie
{
public:
    // Root node of the Trie
    TrieNode *root;

    // Constructor to initialize the root node of the Trie
    Trie()
    {
        root = new TrieNode();
    }

    // Function to insert a word into the Trie
    // For each character in the word, it will either create a new node or traverse an existing one
    void insert(const string &word)
    {
        TrieNode *node = root; // Start from the root node
        for (char ch : word)
        { // Iterate through each character in the word
            // If the character is not already a child of the current node, create a new TrieNode
            if (!node->children.count(ch))
            {
                node->children[ch] = new TrieNode();
            }
            // Move to the child node corresponding to the character
            node = node->children[ch];
            // Increment the count to indicate that one more word passes through this prefix
            node->count++;
        }
    }

    // Function to calculate the sum of scores for all prefixes of the given word
    // The score for a prefix is the number of words that have this prefix
    int getPrefixScoreSum(const string &word)
    {
        TrieNode *node = root; // Start from the root node
        int scoreSum = 0;      // Initialize the total score sum to 0
        for (char ch : word)
        { // Iterate through each character in the word
            // Move to the child node corresponding to the character
            node = node->children[ch];
            // Add the count of words passing through this prefix (the current node) to the score sum
            scoreSum += node->count;
        }
        // Return the total score sum for all prefixes of the word
        return scoreSum;
    }
};

// Solution class to calculate the sum of prefix scores for a list of words
class Solution
{
public:
    // Function that calculates the sum of prefix scores for each word in the input vector
    vector<int> sumPrefixScores(vector<string> &words)
    {
        Trie trie; // Create an instance of the Trie

        // Insert all words into the Trie
        for (const string &word : words)
        {
            trie.insert(word); // Insert each word into the Trie
        }

        // Calculate the sum of prefix scores for each word
        vector<int> result; // Vector to store the prefix score sums for all words
        for (const string &word : words)
        {
            // For each word, calculate the sum of scores for its prefixes using the Trie
            result.push_back(trie.getPrefixScoreSum(word));
        }

        // Return the result containing prefix score sums for all words
        return result;
    }
};