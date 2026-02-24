# 1022. Sum of Root To Leaf Binary Numbers

---

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

We are given a binary tree where each node contains either 0 or 1.

Each root to leaf path represents a binary number. The root is the most significant bit.

Our task is to return the sum of all binary numbers formed from root to leaf paths.

---

## Constraints

* Number of nodes is in the range [1, 1000]
* Node value is either 0 or 1
* Answer fits in a 32 bit integer

---

## Intuition

When I first saw this problem, I realized something important.

Every root to leaf path forms a binary number.

Instead of storing the whole path and converting later, I thought why not build the number while traversing the tree.

In binary, shifting left means multiplying by 2.

So if my current number is `x` and I see a new bit `b`, I can update it as:

```bash
x = x * 2 + b
```

This way I keep building the binary number while moving down the tree.

When I reach a leaf node, I simply add that number to my answer.

---

## Approach

1. Use DFS traversal.
2. Pass current value while going down.
3. At each node update current as current = current * 2 + node value.
4. If node is a leaf, return current.
5. Otherwise return left subtree sum + right subtree sum.

No extra array.
No string conversion.
Everything is calculated on the fly.

---

## Data Structures Used

* Recursion stack
* Binary Tree

No additional data structures are used.

---

## Operations & Behavior Summary

* Visit every node exactly once.
* Build binary value during traversal.
* Add value only at leaf nodes.
* Return total sum.

---

## Complexity

**Time Complexity:** O(n)

Where n is number of nodes. Every node is visited once.

**Space Complexity:** O(h)

Where h is height of tree.
Worst case O(n) for skewed tree.
Best case O(log n) for balanced tree.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int dfs(TreeNode* node, int current) {
        if (!node) return 0;

        current = current * 2 + node->val;

        if (!node->left && !node->right)
            return current;

        return dfs(node->left, current) + dfs(node->right, current);
    }

    int sumRootToLeaf(TreeNode* root) {
        return dfs(root, 0);
    }
};
```

---

### Java

```java
class Solution {
    private int dfs(TreeNode node, int current) {
        if (node == null) return 0;

        current = current * 2 + node.val;

        if (node.left == null && node.right == null)
            return current;

        return dfs(node.left, current) + dfs(node.right, current);
    }

    public int sumRootToLeaf(TreeNode root) {
        return dfs(root, 0);
    }
}
```

---

### JavaScript

```javascript
var sumRootToLeaf = function(root) {
    function dfs(node, current) {
        if (!node) return 0;

        current = current * 2 + node.val;

        if (!node.left && !node.right)
            return current;

        return dfs(node.left, current) + dfs(node.right, current);
    }

    return dfs(root, 0);
};
```

---

### Python3

```python
class Solution:
    def sumRootToLeaf(self, root):
        def dfs(node, current):
            if not node:
                return 0

            current = current * 2 + node.val

            if not node.left and not node.right:
                return current

            return dfs(node.left, current) + dfs(node.right, current)

        return dfs(root, 0)
```

---

### Go

```go
func sumRootToLeaf(root *TreeNode) int {
    var dfs func(node *TreeNode, current int) int

    dfs = func(node *TreeNode, current int) int {
        if node == nil {
            return 0
        }

        current = current*2 + node.Val

        if node.Left == nil && node.Right == nil {
            return current
        }

        return dfs(node.Left, current) + dfs(node.Right, current)
    }

    return dfs(root, 0)
}
```

---

## Step-by-step Detailed Explanation

1. Start DFS from root with current value 0.
2. At every node multiply current by 2.
3. Add node value.
4. If leaf node return current.
5. Otherwise compute left and right subtree sums.
6. Return left sum + right sum.

This ensures every root to leaf binary number is calculated exactly once.

---

## Examples

Input:

```bash
root = [1,0,1,0,1,0,1]
```

Paths:

100 = 4
101 = 5
110 = 6
111 = 7

Output:

```bash
22
```

---

## How to use / Run locally

1. Copy the code into your local IDE.
2. Create TreeNode structure.
3. Build test tree manually.
4. Call sumRootToLeaf function.
5. Print result.

For LeetCode:

* Paste inside Solution class.
* Submit directly.

---

## Notes & Optimizations

* No need to store path.
* No need to convert binary string.
* Multiply by 2 is equivalent to left shift.
* Fully optimized DFS solution.
* Works within integer limits.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
