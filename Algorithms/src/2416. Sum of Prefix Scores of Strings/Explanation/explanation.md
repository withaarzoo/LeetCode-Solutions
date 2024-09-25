# Sum of Prefix Scores for Words using Trie Data Structure

This project demonstrates the implementation of a **Trie (prefix tree)** data structure to solve the problem of calculating the sum of prefix scores for a list of words. The Trie is used to efficiently store and retrieve prefixes for each word in the input list, and the prefix score is calculated as the number of words that share a common prefix.

## Key Concepts

1. **TrieNode**: A node in the Trie that stores child nodes and a count to track how many words pass through this node (i.e., how many words share this prefix).
2. **Trie**: A class that implements a Trie with the ability to insert words and calculate prefix scores.

---

## Step-by-Step Explanation for Each Language

### C++ Implementation

1. **Define a TrieNode Structure**:
   - Contains a map to store child nodes, where each character points to the next node.
   - `count` tracks how many words pass through this node (or prefix).

2. **Define a Trie Class**:
   - The Trie contains the root node and supports word insertion and prefix score calculation.

3. **Insert Words into the Trie**:
   - For each word, traverse the Trie character by character.
   - Create a new node if a character is missing and update the `count` for each node passed.

4. **Calculate Prefix Scores**:
   - For each word, traverse the Trie again and sum the `count` values for each character in the word, which represents how many words share the prefix up to that character.

5. **Solution Class**:
   - Insert all words into the Trie.
   - For each word, calculate the prefix score and return the results in an array.

### Java Implementation

1. **Define a TrieNode Class**:
   - Contains a HashMap for storing child nodes (characters) and a `count` variable to track how many words pass through the node.

2. **Define a Trie Class**:
   - Contains the root node and methods to insert words and calculate prefix scores.

3. **Insert Words**:
   - Traverse each character of the word, create a new TrieNode if it doesn't exist, and increment the `count` for each node passed.

4. **Calculate Prefix Scores**:
   - For each word, traverse its characters and sum the `count` values at each step to compute the total score for all prefixes of the word.

5. **Solution Class**:
   - Insert words into the Trie and then compute and return the prefix scores for each word in the input list.

### JavaScript Implementation

1. **Define a TrieNode Class**:
   - `children` stores the next characters as child nodes, and `count` keeps track of how many words pass through the node.

2. **Define a Trie Class**:
   - The Trie class holds the root node and methods to insert words and compute prefix scores.

3. **Insert Words**:
   - Traverse the word, and if a character doesn't exist as a child, create a new TrieNode.
   - Increment the `count` for each node while inserting the word.

4. **Calculate Prefix Scores**:
   - For each word, traverse the Trie character by character, summing the `count` at each node to calculate the total score for the word's prefixes.

5. **Function to Compute Scores**:
   - Insert all words into the Trie and calculate prefix scores, storing them in an array and returning it.

### Python Implementation

1. **Define a TrieNode Class**:
   - `children` is a dictionary to store child nodes, and `count` tracks how many words pass through each prefix.

2. **Define a Trie Class**:
   - Contains the root node and methods to insert words and calculate the prefix score.

3. **Insert Words**:
   - For each word, traverse character by character, creating new TrieNodes if necessary and incrementing the `count` for each node.

4. **Calculate Prefix Scores**:
   - For each word, traverse its characters and sum the `count` values at each step to calculate the total score for all its prefixes.

5. **Solution Class**:
   - Insert words into the Trie, then compute and return the prefix scores for each word in the list.

### Go Implementation

1. **Define a TrieNode Struct**:
   - Contains a map to store child nodes (`children`) and a `count` variable to track how many words pass through each prefix.

2. **Define a Trie Struct**:
   - The Trie contains the root node and methods for inserting words and calculating prefix scores.

3. **Insert Words into the Trie**:
   - For each word, traverse the Trie, and if a character does not have a corresponding child node, create a new one. Increment the `count` for each node.

4. **Calculate Prefix Scores**:
   - For each word, traverse the Trie and sum the `count` for each character to calculate the total score for the word's prefixes.

5. **Solution Function**:
   - Insert all words into the Trie and calculate the sum of prefix scores for each word.

---

### Usage

- **Input**: A list of words.
- **Output**: A list of integers representing the sum of prefix scores for each word.
