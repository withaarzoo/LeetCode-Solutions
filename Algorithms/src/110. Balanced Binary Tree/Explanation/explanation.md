# 110. Balanced Binary Tree

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

Given the root of a binary tree, I need to check whether the tree is **height-balanced**.

A binary tree is called height-balanced if:

* For **every node**, the height difference between its left and right subtree is **at most 1**.

If the tree satisfies this condition for all nodes, return `true`.
Otherwise, return `false`.

---

## Constraints

* Number of nodes in the tree: `0 ≤ n ≤ 5000`
* Node values range from `-10^4` to `10^4`
* The input tree may be empty

---

## Intuition

When I saw this problem, the key word that caught my attention was **height-balanced**.

My first thought was:

* To check balance, I must know the **height of left and right subtrees**.
* But checking height again and again for every node would be slow.

So I decided:

* I will calculate the height using **DFS**.
* At the same time, I will check whether the tree is balanced.
* If I ever find an unbalanced subtree, I will **stop early**.

To do this efficiently, I used a simple trick:

* Return `-1` if a subtree is unbalanced.
* Otherwise, return the height of the subtree.

---

## Approach

1. I use **Depth First Search (DFS)**.
2. I create a helper function that returns:

   * Height of the subtree if it is balanced
   * `-1` if it is not balanced
3. For each node:

   * I get left subtree height
   * I get right subtree height
4. If height difference is more than `1`, I return `-1`.
5. If balanced, I return `1 + max(leftHeight, rightHeight)`.
6. Finally:

   * If the helper function returns `-1`, the tree is not balanced
   * Otherwise, it is balanced

---

## Data Structures Used

* **Binary Tree**
* **Recursion Stack** (used by DFS)

No extra data structures are used.

---

## Operations & Behavior Summary

* Traverse tree using DFS
* Compute height bottom-up
* Early termination when imbalance is detected
* Single traversal for optimal performance

---

## Complexity

**Time Complexity:** `O(n)`

* `n` is the number of nodes
* Each node is visited only once

**Space Complexity:** `O(h)`

* `h` is the height of the tree
* Space is used by the recursion stack

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int height(TreeNode* root) {
        if (root == nullptr) return 0;

        int left = height(root->left);
        if (left == -1) return -1;

        int right = height(root->right);
        if (right == -1) return -1;

        if (abs(left - right) > 1) return -1;

        return 1 + max(left, right);
    }

    bool isBalanced(TreeNode* root) {
        return height(root) != -1;
    }
};
```

---

### Java

```java
class Solution {

    private int height(TreeNode root) {
        if (root == null) return 0;

        int left = height(root.left);
        if (left == -1) return -1;

        int right = height(root.right);
        if (right == -1) return -1;

        if (Math.abs(left - right) > 1) return -1;

        return 1 + Math.max(left, right);
    }

    public boolean isBalanced(TreeNode root) {
        return height(root) != -1;
    }
}
```

---

### JavaScript

```javascript
var isBalanced = function(root) {

    function height(node) {
        if (node === null) return 0;

        const left = height(node.left);
        if (left === -1) return -1;

        const right = height(node.right);
        if (right === -1) return -1;

        if (Math.abs(left - right) > 1) return -1;

        return 1 + Math.max(left, right);
    }

    return height(root) !== -1;
};
```

---

### Python3

```python
class Solution:
    def isBalanced(self, root):

        def height(node):
            if not node:
                return 0

            left = height(node.left)
            if left == -1:
                return -1

            right = height(node.right)
            if right == -1:
                return -1

            if abs(left - right) > 1:
                return -1

            return 1 + max(left, right)

        return height(root) != -1
```

---

### Go

```go
func isBalanced(root *TreeNode) bool {

    var height func(node *TreeNode) int
    height = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        left := height(node.Left)
        if left == -1 {
            return -1
        }

        right := height(node.Right)
        if right == -1 {
            return -1
        }

        if left-right > 1 || right-left > 1 {
            return -1
        }

        if left > right {
            return 1 + left
        }
        return 1 + right
    }

    return height(root) != -1
}
```

---

## Step-by-step Detailed Explanation

1. I start DFS from the root.
2. I go to the leftmost node first.
3. If a node is `null`, I return height `0`.
4. I calculate left subtree height.
5. I calculate right subtree height.
6. If height difference is more than `1`, I return `-1`.
7. Otherwise, I return current height.
8. This result bubbles up to the root.
9. If at any point I receive `-1`, the tree is unbalanced.

---

## Examples

**Example 1**

```bash
Input:  [3,9,20,null,null,15,7]
Output: true
```

**Example 2**

```bash
Input:  [1,2,2,3,3,null,null,4,4]
Output: false
```

**Example 3**

```bash
Input:  []
Output: true
```

---

## How to use / Run locally

1. Copy the solution code in your preferred language
2. Paste it into your local editor or LeetCode editor
3. Use the provided `TreeNode` definition
4. Run the program with test cases

---

## Notes & Optimizations

* This solution avoids repeated height calculation
* Early termination improves performance
* Works efficiently even for large trees
* Interview-friendly and easy to explain

---

## Author

* **Md Aarzoo Islam**
  🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
