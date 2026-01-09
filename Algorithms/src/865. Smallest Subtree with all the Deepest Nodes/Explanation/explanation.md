# 865. Smallest Subtree with All the Deepest Nodes

---

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

I am given the root of a binary tree.
The depth of a node is defined as the shortest distance from that node to the root.

My task is to return the **smallest subtree** that contains **all the deepest nodes** in the tree.

A subtree means a node and all of its descendants.

If there is only one deepest node, then that node itself is the answer.

---

## Constraints

* Number of nodes is between 1 and 500
* Node values are unique
* Tree is a valid binary tree

---

## Intuition

When I first read the problem, I noticed that it is not asking for the deepest nodes themselves.
It is asking for the **smallest subtree that contains all of them**.

Then I realized something important.

If I take all the deepest nodes and find their **Lowest Common Ancestor**, that node will be the root of the smallest subtree containing all deepest nodes.

So the problem becomes:

* Find the depth of left and right subtrees
* If both sides have deepest nodes at the same depth, current node is the answer
* Otherwise, move toward the deeper side

This insight allows me to solve everything in **one DFS traversal**.

---

## Approach

1. I use **postorder DFS** because decisions depend on child nodes
2. For each node, I calculate:

   * Maximum depth from this node
   * Subtree root that contains all deepest nodes
3. If left and right depths are equal:

   * Current node becomes the subtree root
4. If one side is deeper:

   * I propagate the deeper subtree upward
5. The final returned node is the answer

This approach avoids extra passes and is fully optimized.

---

## Data Structures Used

* Recursion stack (DFS)
* Binary Tree structure (TreeNode)

No extra data structures like arrays, maps, or sets are required.

---

## Operations & Behavior Summary

* Traverse tree using DFS
* Compare left and right subtree depths
* Decide subtree root dynamically
* Return result in one traversal

---

## Complexity

**Time Complexity:** O(n)
Each node is visited exactly once.

**Space Complexity:** O(h)
Where h is the height of the tree due to recursion stack.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    pair<int, TreeNode*> dfs(TreeNode* root) {
        if (!root) return {0, nullptr};

        auto left = dfs(root->left);
        auto right = dfs(root->right);

        if (left.first == right.first)
            return {left.first + 1, root};
        else if (left.first > right.first)
            return {left.first + 1, left.second};
        else
            return {right.first + 1, right.second};
    }

    TreeNode* subtreeWithAllDeepest(TreeNode* root) {
        return dfs(root).second;
    }
};
```

---

### Java

```java
class Solution {

    static class Pair {
        int depth;
        TreeNode node;

        Pair(int d, TreeNode n) {
            depth = d;
            node = n;
        }
    }

    private Pair dfs(TreeNode root) {
        if (root == null) return new Pair(0, null);

        Pair left = dfs(root.left);
        Pair right = dfs(root.right);

        if (left.depth == right.depth)
            return new Pair(left.depth + 1, root);
        else if (left.depth > right.depth)
            return new Pair(left.depth + 1, left.node);
        else
            return new Pair(right.depth + 1, right.node);
    }

    public TreeNode subtreeWithAllDeepest(TreeNode root) {
        return dfs(root).node;
    }
}
```

---

### JavaScript

```javascript
var subtreeWithAllDeepest = function(root) {

    function dfs(node) {
        if (!node) return [0, null];

        const [ld, ln] = dfs(node.left);
        const [rd, rn] = dfs(node.right);

        if (ld === rd) return [ld + 1, node];
        if (ld > rd) return [ld + 1, ln];
        return [rd + 1, rn];
    }

    return dfs(root)[1];
};
```

---

### Python3

```python
class Solution:
    def subtreeWithAllDeepest(self, root):
        def dfs(node):
            if not node:
                return 0, None

            ld, ln = dfs(node.left)
            rd, rn = dfs(node.right)

            if ld == rd:
                return ld + 1, node
            elif ld > rd:
                return ld + 1, ln
            else:
                return rd + 1, rn

        return dfs(root)[1]
```

---

### Go

```go
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {

    var dfs func(*TreeNode) (int, *TreeNode)

    dfs = func(node *TreeNode) (int, *TreeNode) {
        if node == nil {
            return 0, nil
        }

        ld, ln := dfs(node.Left)
        rd, rn := dfs(node.Right)

        if ld == rd {
            return ld + 1, node
        } else if ld > rd {
            return ld + 1, ln
        }
        return rd + 1, rn
    }

    _, ans := dfs(root)
    return ans
}
```

---

## Step-by-step Detailed Explanation

1. Start DFS from root
2. Process left and right subtrees first
3. Get depth and subtree root from both sides
4. Compare depths:

   * Equal depth → current node is answer
   * Unequal → propagate deeper subtree
5. Return final subtree root from root call

This ensures correctness with minimum traversal.

---

## Examples

Input:

```bash
root = [3,5,1,6,2,0,8,null,null,7,4]
```

Output:

```bash
[2,7,4]
```

Explanation:
Nodes 7 and 4 are deepest. Their smallest common subtree is rooted at node 2.

---

## How to use / Run locally

1. Clone the repository
2. Choose your preferred language file
3. Compile and run using standard compiler
4. Pass test cases via main function or online judge

---

## Notes & Optimizations

* Single DFS traversal
* No extra memory usage
* Works efficiently for large trees
* Same logic as Lowest Common Ancestor of deepest leaves

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
