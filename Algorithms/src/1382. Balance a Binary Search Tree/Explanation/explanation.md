# Balance a Binary Search Tree

**LeetCode Problem 1382**

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

I am given the root of a **Binary Search Tree (BST)**.
The tree may be unbalanced.

My task is to **return a balanced BST** that contains **the same node values**.

A BST is considered **balanced** if, for every node, the height difference between its left and right subtree is **not more than 1**.

If multiple balanced trees are possible, I can return **any one of them**.

---

## Constraints

* Number of nodes: `1 ≤ n ≤ 10^4`
* Node value range: `1 ≤ Node.val ≤ 10^5`
* Input tree is always a valid BST

---

## Intuition

When I saw this problem, I remembered one very important fact:

👉 **Inorder traversal of a BST always gives a sorted sequence.**

Then I thought:

* If I convert the BST into a sorted array
* And then rebuild the tree like binary search (middle element as root)
* The new tree will automatically become balanced

So instead of fixing the existing tree structure,
I decided to **rebuild a new balanced BST** using its values.

This makes the solution simple, clean, and reliable.

---

## Approach

I solved the problem in two main steps:

1. **Inorder Traversal**

   * Traverse the BST using inorder
   * Store all node values in a sorted array

2. **Build Balanced BST**

   * Choose the middle element as root
   * Recursively build left subtree from left half
   * Recursively build right subtree from right half

This guarantees:

* BST property is preserved
* Tree height stays balanced

---

## Data Structures Used

* Array / List → to store inorder values
* Recursion → to rebuild the balanced BST

---

## Operations & Behavior Summary

* Traverse original tree once
* Store values in sorted order
* Rebuild a new balanced tree
* Return the new root

No node values are lost or changed.

---

## Complexity

**Time Complexity:** `O(n)`

* Inorder traversal takes `O(n)`
* Building BST also takes `O(n)`

**Space Complexity:** `O(n)`

* Extra array for storing values
* Recursive call stack

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> arr;

    void inorder(TreeNode* root) {
        if (!root) return;
        inorder(root->left);
        arr.push_back(root->val);
        inorder(root->right);
    }

    TreeNode* build(int l, int r) {
        if (l > r) return nullptr;
        int mid = (l + r) / 2;
        TreeNode* node = new TreeNode(arr[mid]);
        node->left = build(l, mid - 1);
        node->right = build(mid + 1, r);
        return node;
    }

    TreeNode* balanceBST(TreeNode* root) {
        inorder(root);
        return build(0, arr.size() - 1);
    }
};
```

---

### Java

```java
class Solution {
    List<Integer> arr = new ArrayList<>();

    void inorder(TreeNode root) {
        if (root == null) return;
        inorder(root.left);
        arr.add(root.val);
        inorder(root.right);
    }

    TreeNode build(int l, int r) {
        if (l > r) return null;
        int mid = (l + r) / 2;
        TreeNode node = new TreeNode(arr.get(mid));
        node.left = build(l, mid - 1);
        node.right = build(mid + 1, r);
        return node;
    }

    public TreeNode balanceBST(TreeNode root) {
        inorder(root);
        return build(0, arr.size() - 1);
    }
}
```

---

### JavaScript

```javascript
var balanceBST = function(root) {
    let arr = [];

    function inorder(node) {
        if (!node) return;
        inorder(node.left);
        arr.push(node.val);
        inorder(node.right);
    }

    function build(l, r) {
        if (l > r) return null;
        let mid = Math.floor((l + r) / 2);
        let node = new TreeNode(arr[mid]);
        node.left = build(l, mid - 1);
        node.right = build(mid + 1, r);
        return node;
    }

    inorder(root);
    return build(0, arr.length - 1);
};
```

---

### Python3

```python
class Solution:
    def balanceBST(self, root):
        arr = []

        def inorder(node):
            if not node:
                return
            inorder(node.left)
            arr.append(node.val)
            inorder(node.right)

        def build(l, r):
            if l > r:
                return None
            mid = (l + r) // 2
            node = TreeNode(arr[mid])
            node.left = build(l, mid - 1)
            node.right = build(mid + 1, r)
            return node

        inorder(root)
        return build(0, len(arr) - 1)
```

---

### Go

```go
func balanceBST(root *TreeNode) *TreeNode {
    arr := []int{}

    var inorder func(*TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inorder(node.Left)
        arr = append(arr, node.Val)
        inorder(node.Right)
    }

    var build func(int, int) *TreeNode
    build = func(l, r int) *TreeNode {
        if l > r {
            return nil
        }
        mid := (l + r) / 2
        node := &TreeNode{Val: arr[mid]}
        node.Left = build(l, mid-1)
        node.Right = build(mid+1, r)
        return node
    }

    inorder(root)
    return build(0, len(arr)-1)
}
```

---

## Step-by-step Detailed Explanation

1. I traverse the BST using inorder
2. I store values in a sorted array
3. I pick the middle element as root
4. Left side builds the left subtree
5. Right side builds the right subtree
6. Recursion continues until tree is complete

This is similar to binary search logic.

---

## Examples

**Input**

```bash
[1, null, 2, null, 3, null, 4]
```

**Output**

```bash
[2, 1, 3, null, null, null, 4]
```

Multiple correct outputs are possible.

---

## How to use / Run locally

1. Copy the code for your language
2. Paste it into LeetCode editor or local IDE
3. Use provided test cases
4. Run and verify output

---

## Notes & Optimizations

* This solution is simple and safe
* No need to modify original tree in-place
* Works efficiently for up to `10^4` nodes
* Very common interview pattern

---

## Author

**Md Aarzoo Islam**
🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
