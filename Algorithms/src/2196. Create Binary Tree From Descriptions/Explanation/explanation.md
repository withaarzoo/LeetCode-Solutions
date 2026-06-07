# 2196. Create Binary Tree From Descriptions

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)

  * [C++](#c)
  * [Java](#java)
  * [JavaScript](#javascript)
  * [Python3](#python3)
  * [Go](#go)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

Given a list of parent-child relationships, we need to construct the complete binary tree and return its root node.

Each description contains three values:

* `parent`
* `child`
* `isLeft`

If `isLeft = 1`, the child becomes the left child of the parent.

If `isLeft = 0`, the child becomes the right child of the parent.

The challenge is not only building the tree correctly but also identifying the root node. Since every node except the root appears as someone's child, we can use that observation to find the answer efficiently.

This problem is a classic Binary Tree Construction problem that combines tree building, hash maps, and parent-child relationship tracking.

---

## Constraints

| Constraint             | Value             |
| ---------------------- | ----------------- |
| descriptions.length    | 1 ≤ n ≤ 10⁴       |
| descriptions[i].length | = 3               |
| parentᵢ                | 1 ≤ parentᵢ ≤ 10⁵ |
| childᵢ                 | 1 ≤ childᵢ ≤ 10⁵  |
| isLeftᵢ                | 0 ≤ isLeftᵢ ≤ 1   |
| Tree Validity          | Guaranteed Valid  |

---

## Intuition

My first observation was that every description directly gives one edge of the binary tree.

The tree itself is not difficult to build. The real question is how to find the root.

I noticed that every node except the root appears as a child at least once. The root is the only node that never appears as a child.

So I decided to:

1. Create nodes as I encounter them.
2. Connect parent and child nodes.
3. Keep track of every child node.
4. Find the node that never appeared as a child.

That node must be the root.

---

## Approach

1. Create a hash map that stores node value → TreeNode.
2. Create a hash set to store all child values.
3. Iterate through every description.
4. Create parent and child nodes if they do not already exist.
5. Connect the child to the left or right side based on `isLeft`.
6. Add the child value to the child set.
7. After processing all descriptions, scan all created nodes.
8. The node whose value is not present in the child set is the root.
9. Return that node.

This builds the tree and finds the root in a single pass.

---

## Data Structures Used

### Hash Map

Used to store:

```text
node value -> TreeNode
```

This allows instant access to any node while building the tree.

### Hash Set

Used to store all child values.

This helps identify the root because the root never appears as a child.

### Binary Tree

The final structure being constructed from the descriptions.

---

## Operations & Behavior Summary

The algorithm performs the following steps:

1. Read a description.
2. Create missing nodes.
3. Connect parent and child.
4. Record the child node.
5. Repeat for all descriptions.
6. Search for a node that never appeared as a child.
7. Return that node as the root.

High-level pseudocode:

```text
Create node map
Create child set

For every description:
    Create parent if needed
    Create child if needed
    Connect nodes
    Mark child

Find node not present in child set

Return root
```

---

## Complexity

| Metric           | Complexity | Explanation                                      |
| ---------------- | ---------- | ------------------------------------------------ |
| Time Complexity  | O(n)       | Each description is processed once               |
| Space Complexity | O(n)       | Hash map and hash set store up to n unique nodes |

Where:

* `n` = number of descriptions

---

## Multi-language Solutions

### C++

```cpp
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* createBinaryTree(vector<vector<int>>& descriptions) {
        
        // Stores value -> TreeNode mapping
        unordered_map<int, TreeNode*> nodes;
        
        // Stores all nodes that appear as children
        unordered_set<int> children;
        
        for (auto &d : descriptions) {
            int parent = d[0];
            int child = d[1];
            int isLeft = d[2];
            
            // Create parent node if not present
            if (!nodes.count(parent))
                nodes[parent] = new TreeNode(parent);
            
            // Create child node if not present
            if (!nodes.count(child))
                nodes[child] = new TreeNode(child);
            
            // Connect child to correct side
            if (isLeft)
                nodes[parent]->left = nodes[child];
            else
                nodes[parent]->right = nodes[child];
            
            // Mark child node
            children.insert(child);
        }
        
        // Root is the node that never appeared as a child
        for (auto &[value, node] : nodes) {
            if (!children.count(value))
                return node;
        }
        
        return nullptr;
    }
};
```

### Java

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public TreeNode createBinaryTree(int[][] descriptions) {
        
        // Stores value -> TreeNode mapping
        Map<Integer, TreeNode> nodes = new HashMap<>();
        
        // Stores all child values
        Set<Integer> children = new HashSet<>();
        
        for (int[] d : descriptions) {
            int parent = d[0];
            int child = d[1];
            int isLeft = d[2];
            
            // Create parent node if needed
            nodes.putIfAbsent(parent, new TreeNode(parent));
            
            // Create child node if needed
            nodes.putIfAbsent(child, new TreeNode(child));
            
            // Attach child to correct side
            if (isLeft == 1)
                nodes.get(parent).left = nodes.get(child);
            else
                nodes.get(parent).right = nodes.get(child);
            
            // Mark child
            children.add(child);
        }
        
        // Find root
        for (int value : nodes.keySet()) {
            if (!children.contains(value))
                return nodes.get(value);
        }
        
        return null;
    }
}
```

### JavaScript

```javascript
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number[][]} descriptions
 * @return {TreeNode}
 */
var createBinaryTree = function(descriptions) {
    
    // Stores value -> TreeNode mapping
    const nodes = new Map();
    
    // Stores all child values
    const children = new Set();
    
    for (const [parent, child, isLeft] of descriptions) {
        
        // Create parent node if needed
        if (!nodes.has(parent))
            nodes.set(parent, new TreeNode(parent));
        
        // Create child node if needed
        if (!nodes.has(child))
            nodes.set(child, new TreeNode(child));
        
        // Connect child to parent
        if (isLeft === 1)
            nodes.get(parent).left = nodes.get(child);
        else
            nodes.get(parent).right = nodes.get(child);
        
        // Mark child
        children.add(child);
    }
    
    // Root never appears as a child
    for (const [value, node] of nodes) {
        if (!children.has(value))
            return node;
    }
    
    return null;
};
```

### Python3

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def createBinaryTree(self, descriptions: List[List[int]]) -> Optional[TreeNode]:
        
        # Stores value -> TreeNode mapping
        nodes = {}
        
        # Stores all child values
        children = set()
        
        for parent, child, isLeft in descriptions:
            
            # Create parent node if needed
            if parent not in nodes:
                nodes[parent] = TreeNode(parent)
            
            # Create child node if needed
            if child not in nodes:
                nodes[child] = TreeNode(child)
            
            # Connect child to correct side
            if isLeft:
                nodes[parent].left = nodes[child]
            else:
                nodes[parent].right = nodes[child]
            
            # Mark child
            children.add(child)
        
        # Root never appears as a child
        for value, node in nodes.items():
            if value not in children:
                return node
```

### Go

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    
    // Stores value -> TreeNode mapping
    nodes := make(map[int]*TreeNode)
    
    // Stores all child values
    children := make(map[int]bool)
    
    for _, d := range descriptions {
        parent := d[0]
        child := d[1]
        isLeft := d[2]
        
        // Create parent node if needed
        if _, exists := nodes[parent]; !exists {
            nodes[parent] = &TreeNode{Val: parent}
        }
        
        // Create child node if needed
        if _, exists := nodes[child]; !exists {
            nodes[child] = &TreeNode{Val: child}
        }
        
        // Attach child to correct side
        if isLeft == 1 {
            nodes[parent].Left = nodes[child]
        } else {
            nodes[parent].Right = nodes[child]
        }
        
        // Mark child
        children[child] = true
    }
    
    // Find root node
    for value, node := range nodes {
        if !children[value] {
            return node
        }
    }
    
    return nil
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains identical across all five languages.

### Step 1: Store Every Node

Whenever a value appears, we check whether a node already exists.

If not, we create it.

This guarantees that every value corresponds to exactly one tree node.

---

### Step 2: Track Child Nodes

Every child value is inserted into a set.

For example:

```text
[50,20,1]
[50,80,0]
```

The child set becomes:

```text
{20,80}
```

Later, this information helps us find the root.

---

### Step 3: Connect Parent and Child

For each description:

```text
[parent, child, isLeft]
```

If:

```text
isLeft = 1
```

we connect:

```text
parent.left = child
```

Otherwise:

```text
parent.right = child
```

This gradually builds the complete binary tree.

---

### Step 4: Find the Root

Once the tree is built, we look through every created node.

Every node except the root appears in the child set.

The only node missing from that set is the root.

---

### Why This Works

A valid binary tree contains exactly one root.

Every non-root node has exactly one parent.

Therefore:

```text
Root = Node that never appears as a child
```

The algorithm uses this property directly.

---

## Examples

### Example 1

Input

```text
descriptions = [[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]]
```

Output

```text
[50,20,80,15,17,19]
```

Trace

```text
50 -> root candidate
20 -> child
80 -> child
15 -> child
17 -> child
19 -> child
```

Only `50` never appears as a child.

Therefore:

```text
Root = 50
```

---

### Example 2

Input

```text
descriptions = [[1,2,1],[2,3,0],[3,4,1]]
```

Output

```text
[1,2,null,3,4]
```

Trace

```text
Children = {2,3,4}
```

Node `1` never appears as a child.

Therefore:

```text
Root = 1
```

---

### Example 3

Input

```text
descriptions = [[10,5,1],[10,20,0]]
```

Output

```text
[10,5,20]
```

Trace

```text
Children = {5,20}
```

Node `10` is never a child.

Therefore:

```text
Root = 10
```

---

## How to Use / Run Locally

### C++

Compile

```bash
g++ main.cpp -o main
```

Run

```bash
./main
```

---

### Java

Compile

```bash
javac Main.java
```

Run

```bash
java Main
```

---

### JavaScript

Run

```bash
node main.js
```

---

### Python3

Run

```bash
python main.py
```

or

```bash
python3 main.py
```

---

### Go

Run

```bash
go run main.go
```

Build

```bash
go build
```

---

## Notes & Optimizations

* The solution already achieves optimal time complexity.
* No tree traversal is required to find the root.
* Hash maps provide constant-time node lookup.
* Hash sets provide constant-time child existence checks.
* Since the tree is guaranteed to be valid, additional validation logic is unnecessary.
* An alternative approach could store parent references directly, but it would not improve complexity.
* This approach is simple, clean, and commonly used in binary tree construction problems.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
