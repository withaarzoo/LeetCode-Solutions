# LeetCode 3093: Longest Common Suffix Queries

## Table of Contents

* [Problem Summary](https://www.google.com/search?q=%23problem-summary)
* [Constraints](https://www.google.com/search?q=%23constraints)
* [Intuition](https://www.google.com/search?q=%23intuition)
* [Approach](https://www.google.com/search?q=%23approach)
* [Data Structures Used](https://www.google.com/search?q=%23data-structures-used)
* [Operations & Behavior Summary](https://www.google.com/search?q=%23operations--behavior-summary)
* [Complexity](https://www.google.com/search?q=%23complexity)
* [Multi-language Solutions](https://www.google.com/search?q=%23multi-language-solutions)
* [C++](https://www.google.com/search?q=%23c)
* [Java](https://www.google.com/search?q=%23java)
* [JavaScript](https://www.google.com/search?q=%23javascript)
* [Python3](https://www.google.com/search?q=%23python3)
* [Go](https://www.google.com/search?q=%23go)

* [Step-by-step Detailed Explanation](https://www.google.com/search?q=%23step-by-step-detailed-explanation)
* [Examples](https://www.google.com/search?q=%23examples)
* [How to Use / Run Locally](https://www.google.com/search?q=%23how-to-use--run-locally)
* [Notes & Optimizations](https://www.google.com/search?q=%23notes--optimizations)
* [Author](https://www.google.com/search?q=%23author)

## Problem Summary

In this problem, we are given two arrays of strings: a container array and a query array. For every string in the query array, we need to find a string in the container array that shares the longest common suffix with it.

If there are multiple container strings that match the longest suffix, we have strict tie-breaking rules:

1. Pick the string with the shortest overall length.
2. If there is still a tie, pick the string that appears earliest in the container array (the one with the smallest index).

Our goal is to return an array of integers representing the index of the best matching container string for each query.

## Constraints

| Constraint | Value |
| --- | --- |
| `wordsContainer.length` | 1 to 10,000 |
| `wordsQuery.length` | 1 to 10,000 |
| `wordsContainer[i].length` | 1 to 5,000 |
| `wordsQuery[i].length` | 1 to 5,000 |
| Sum of `wordsContainer[i].length` | Up to 500,000 |
| Sum of `wordsQuery[i].length` | Up to 500,000 |
| Character set | Lowercase English letters only |

## Intuition

When I first read this problem, comparing suffixes directly for thousands of queries against thousands of container words sounded like a recipe for a Time Limit Exceeded error. Brute force would just be too slow.

But if you think about string algorithms, prefix matching is incredibly fast if you use a Trie (a prefix tree). A suffix is just a prefix if you read the string backwards. My immediate instinct was to reverse all the strings in the container and insert them into a Trie. Then, I could reverse the query strings and just walk down the tree to find the longest matching path.

Because we have tie-breaking rules, I realized I could pre-calculate the "best" answer for every single node in the Trie during the building phase. That way, querying is instant.

## Approach

My strategy is to build a highly optimized Trie where every single node keeps track of the "best" container string index that passes through it.

First, I initialize a root node. This root node will keep track of the absolute best string in the entire container, which acts as our fallback answer if a query matches nothing at all.

Then, I loop through the container strings. For each string, I traverse it backwards character by character. If a character path does not exist in the tree, I create a new node. As I step into each node, I check if the current string is "better" (shorter length, or same length but earlier index) than the string that currently owns that node. If it is, I update the node's record.

Finally, I process the queries. For each query, I traverse it backwards. I walk down the Trie following the characters until I hit a dead end. The node where I stop already contains the pre-computed index of the best possible answer. I grab that index and move to the next query.

## Data Structures Used

* **Trie (Prefix Tree)**: Chosen because it provides optimal time complexity for string matching. Instead of just storing children, our custom Trie nodes store metadata (`best_len` and `best_idx`) to track the optimal answer for any given suffix path.
* **Arrays/Lists**: Used to hold the final answers to be returned. In languages like C++, a flat array or vector of nodes is used instead of pointer-based nodes to heavily optimize memory and prevent Memory Limit Exceeded issues.

## Operations & Behavior Summary

1. Initialize a Trie root node with default "infinity" values for best length and index.
2. Iterate over the container array, reading each word backwards.
3. Traverse or create nodes for each character.
4. At every step (including the root), update the node's stored best length and index if the current word is a better candidate.
5. Iterate over the query array, reading each query backwards.
6. Walk down the Trie. Break the loop early if a character branch does not exist.
7. Record the best index stored at the final valid node reached.

## Complexity

| Metric | Complexity | Explanation |
| --- | --- | --- |
| **Time Complexity** | O(N + M) | N is the total number of characters across all container words. M is the total number of characters across all query words. We touch each character exactly once during insertion and once during querying. |
| **Space Complexity** | O(N) | In the worst-case scenario where there are no shared suffixes, the Trie will create a node for every character in the container array. The alphabet size is fixed at 26, so the branching factor is constant. |

## Multi-language Solutions

### C++

```cpp
class Solution {
    // Structure to define a node in the Trie
    struct TrieNode {
        int children[26];
        int bestLen;
        int bestIdx;
        
        // Initialize children to -1 and best trackers to a large number
        TrieNode() {
            fill(begin(children), end(children), -1);
            bestLen = 1e9;
            bestIdx = 1e9;
        }
    };

public:
    vector<int> stringIndices(vector<string>& wordsContainer, vector<string>& wordsQuery) {
        // Using a vector of nodes instead of pointers to avoid memory overhead and MLE
        vector<TrieNode> trie;
        trie.emplace_back(); // Push root node
        
        // Insert each string from the container into the Trie
        for (int i = 0; i < wordsContainer.size(); i++) {
            int len = wordsContainer[i].length();
            int curr = 0; // Start at root
            
            // Update root with the absolute best string (in case of zero matches later)
            if (len < trie[curr].bestLen || (len == trie[curr].bestLen && i < trie[curr].bestIdx)) {
                trie[curr].bestLen = len;
                trie[curr].bestIdx = i;
            }
            
            // Traverse the string backwards to simulate suffix matching as prefix matching
            for (int j = len - 1; j >= 0; j--) {
                int charIdx = wordsContainer[i][j] - 'a';
                
                // If child path doesn't exist, create a new node
                if (trie[curr].children[charIdx] == -1) {
                    trie[curr].children[charIdx] = trie.size();
                    trie.emplace_back();
                }
                
                // Move down the Trie
                curr = trie[curr].children[charIdx];
                
                // Update the best string properties for this specific prefix path
                if (len < trie[curr].bestLen || (len == trie[curr].bestLen && i < trie[curr].bestIdx)) {
                    trie[curr].bestLen = len;
                    trie[curr].bestIdx = i;
                }
            }
        }
        
        vector<int> ans;
        ans.reserve(wordsQuery.size());
        
        // Process each query
        for (const string& query : wordsQuery) {
            int curr = 0; // Start at root
            int len = query.length();
            
            // Traverse backwards down the Trie
            for (int j = len - 1; j >= 0; j--) {
                int charIdx = query[j] - 'a';
                // Stop if the matching path ends
                if (trie[curr].children[charIdx] == -1) {
                    break;
                }
                curr = trie[curr].children[charIdx];
            }
            // The current node holds the index of the best matching string
            ans.push_back(trie[curr].bestIdx);
        }
        
        return ans;
    }
};
```

### Java

```java
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
```

### JavaScript

```javascript
/**
 * @param {string[]} wordsContainer
 * @param {string[]} wordsQuery
 * @return {number[]}
 */
var stringIndices = function(wordsContainer, wordsQuery) {
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
```

### Python3

```python
class TrieNode:
    # Using slots prevents dynamic dict creation per instance, saving memory to avoid MLE
    __slots__ = ['children', 'bestLen', 'bestIdx']
    
    def __init__(self):
        self.children = {}
        self.bestLen = float('inf')
        self.bestIdx = float('inf')

class Solution:
    def stringIndices(self, wordsContainer: List[str], wordsQuery: List[str]) -> List[int]:
        root = TrieNode()
        
        # Process every string to build the suffix tree
        for i, word in enumerate(wordsContainer):
            n = len(word)
            curr = root
            
            # Root stores the best word globally (when suffix match length is 0)
            if n < curr.bestLen or (n == curr.bestLen and i < curr.bestIdx):
                curr.bestLen = n
                curr.bestIdx = i
                
            # Iterate backwards to treat suffixes as prefixes
            for char in reversed(word):
                if char not in curr.children:
                    curr.children[char] = TrieNode()
                
                curr = curr.children[char]
                
                # Continuously update the best match info down the branch
                if n < curr.bestLen or (n == curr.bestLen and i < curr.bestIdx):
                    curr.bestLen = n
                    curr.bestIdx = i
                    
        ans = []
        
        # Evaluate queries against the constructed tree
        for query in wordsQuery:
            curr = root
            
            # Traverse backwards until we hit a dead end
            for char in reversed(query):
                if char not in curr.children:
                    break
                curr = curr.children[char]
            
            # Append the stored best index at the deepest reachable node
            ans.append(curr.bestIdx)
            
        return ans
```

### Go

```go
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
```

## Step-by-step Detailed Explanation

No matter which language you use, the core logic flows through the same critical steps.

First, we handle the tree structure. Instead of doing expensive string reversal operations that create new strings in memory, we simply loop through the strings using a decreasing index. This saves a massive amount of overhead.

When building the tree, we always evaluate the root node first. The root represents an empty string match (a suffix of length zero). If a query fails to match even a single character, the answer defaults to whatever is stored at the root. We update the root with the shortest string seen so far.

As we move down a branch for a specific word, we compare the current word's length against the node's recorded `best_len`. If the current word is shorter, it overwrites the node's record. If the lengths are tied, we check the index. Since we process the container sequentially from index 0 upwards, if there is a length tie, the earlier string already sitting in the node wins naturally, so we only overwrite on strictly shorter lengths. However, checking both length and index explicitly makes the code robust.

During the query phase, we start at the root. We look at the last character of the query. If a branch exists for that character, we move to that child node. We repeat this for the second-to-last character, and so on. The moment we ask for a character branch that is null or empty, we halt. We do not need to search any further, and we do not need to backtrack. The exact node we are standing on has already done the hard work of tracking the best container index. We push that index into our results array.

## Examples

**Example 1**

* **Container**: `["abcd", "bcd", "xbcd"]`
* **Query**: `["cd"]`
* **Trace**: We build the Trie backwards. "d" -> "c" -> "b" -> "a". The node at the path `d -> c` is shared by all three container words. At this node, the algorithm looks at the lengths: "abcd" is 4, "bcd" is 3, "xbcd" is 4. The word "bcd" is the shortest, so its index (1) is permanently saved at the `d -> c` node. When the query "cd" traverses the tree, it lands on this node and instantly returns index 1.

**Example 2**

* **Container**: `["apple", "pineapple"]`
* **Query**: `["x"]`
* **Trace**: The Trie is built for the suffixes. The query "x" starts at the root and looks for a branch "x". It does not exist. The loop breaks immediately, leaving us at the root node. The root node tracked the overall shortest word in the container, which is "apple" (index 0). It returns 0.

## How to Use / Run Locally

If you want to run these solutions on your local machine, here is how you compile and execute them based on your language of choice. You will need to wrap the provided class methods in a standard entry point (like a `main` function) and feed it dummy arrays.

* **C++**: Save as `solution.cpp`. Compile using `g++ -std=c++17 solution.cpp -o run`. Execute with `./run`.
* **Java**: Save as `Solution.java`. Compile using `javac Solution.java`. Run the class using `java Solution`.
* **Python**: Save as `solution.py`. Run directly using `python3 solution.py`.
* **JavaScript**: Save as `solution.js`. Make sure you have Node installed, then run with `node solution.js`.
* **Go**: Save as `main.go`. Run using the command `go run main.go`.

## Notes & Optimizations

* **Memory Limit Exceeded (MLE)**: This is the most common trap for this LeetCode problem. The sum of characters can be huge. In C++, using a flat `vector` of struct nodes rather than pointers with `new` saves significant memory and prevents memory fragmentation.
* **Python Object Overhead**: In Python, creating thousands of dictionary objects for Trie nodes can be bloated. Using `__slots__` inside the node class prevents dynamic dictionary allocation per instance, drastically reducing the memory footprint.
* **Early Break**: Notice that in the query search loop, there is no need to keep searching if the path breaks. We just stop. This makes the query extremely fast, essentially running in time equal to the length of the matching suffix rather than the whole query string.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
