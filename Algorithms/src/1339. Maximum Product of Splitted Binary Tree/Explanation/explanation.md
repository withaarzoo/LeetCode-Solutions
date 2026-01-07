# Maximum Product of Splitted Binary Tree (LeetCode 1339)

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
I can remove **only one edge** from the tree.
After removing that edge, the tree splits into **two subtrees**.

Each subtree has a **sum of its node values**.

My goal is to **maximize the product** of the sums of these two subtrees.

Because the result can be very large, I must return the answer **modulo (10⁹ + 7)**.

Important rule:
I must **maximize the product first**, then apply the modulo.

---

## Constraints

* Number of nodes: `2 ≤ n ≤ 5 × 10⁴`
* Node value range: `1 ≤ Node.val ≤ 10⁴`
* Tree is a valid binary tree

---

## Intuition

When I remove one edge, the tree breaks into two parts.

If I know:

* The **total sum** of the entire tree
* The **sum of one subtree**

Then the second subtree sum is:

```bash
totalSum − subtreeSum
```

So for every possible cut, the product becomes:

```bash
subtreeSum × (totalSum − subtreeSum)
```

That means my job is simple:

* Calculate the sum of every subtree
* Try the product formula at every node
* Track the maximum result

---

## Approach

1. First, I calculate the **total sum** of the whole tree using DFS.
2. Then, I run another DFS to calculate **subtree sums**.
3. At each node, I treat its subtree as one part of the split.
4. I calculate the product:

   ```bash
   subtreeSum × (totalSum − subtreeSum)
   ```

5. I update the maximum product.
6. Finally, I return the maximum product modulo `10⁹ + 7`.

DFS is perfect here because it naturally gives subtree sums.

---

## Data Structures Used

* Binary Tree
* Recursion (DFS)
* Integer / Long variables for large values

No extra arrays or maps are needed.

---

## Operations & Behavior Summary

* Traverse the tree twice
* First pass gets total sum
* Second pass evaluates every possible split
* Keeps track of the maximum product
* Applies modulo only at the end

---

## Complexity

**Time Complexity:**
O(n)
Each node is visited a constant number of times.

**Space Complexity:**
O(h)
`h` is the height of the tree due to recursion stack.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long totalSum = 0;
    long long maxProd = 0;
    const int MOD = 1e9 + 7;

    long long getTotalSum(TreeNode* root) {
        if (!root) return 0;
        return root->val + getTotalSum(root->left) + getTotalSum(root->right);
    }

    long long dfs(TreeNode* root) {
        if (!root) return 0;

        long long left = dfs(root->left);
        long long right = dfs(root->right);

        long long subtreeSum = root->val + left + right;
        maxProd = max(maxProd, subtreeSum * (totalSum - subtreeSum));

        return subtreeSum;
    }

    int maxProduct(TreeNode* root) {
        totalSum = getTotalSum(root);
        dfs(root);
        return maxProd % MOD;
    }
};
```

---

### Java

```java
class Solution {
    long totalSum = 0;
    long maxProduct = 0;
    int MOD = 1000000007;

    private long getTotalSum(TreeNode root) {
        if (root == null) return 0;
        return root.val + getTotalSum(root.left) + getTotalSum(root.right);
    }

    private long dfs(TreeNode root) {
        if (root == null) return 0;

        long left = dfs(root.left);
        long right = dfs(root.right);

        long subtreeSum = root.val + left + right;
        maxProduct = Math.max(maxProduct, subtreeSum * (totalSum - subtreeSum));

        return subtreeSum;
    }

    public int maxProduct(TreeNode root) {
        totalSum = getTotalSum(root);
        dfs(root);
        return (int)(maxProduct % MOD);
    }
}
```

---

### JavaScript

```javascript
var maxProduct = function(root) {
    const MOD = 1e9 + 7;
    let totalSum = 0;
    let maxProd = 0;

    function getTotalSum(node) {
        if (!node) return 0;
        return node.val + getTotalSum(node.left) + getTotalSum(node.right);
    }

    function dfs(node) {
        if (!node) return 0;

        let left = dfs(node.left);
        let right = dfs(node.right);

        let subtreeSum = node.val + left + right;
        maxProd = Math.max(maxProd, subtreeSum * (totalSum - subtreeSum));

        return subtreeSum;
    }

    totalSum = getTotalSum(root);
    dfs(root);

    return maxProd % MOD;
};
```

---

### Python3

```python
class Solution:
    def maxProduct(self, root):
        MOD = 10**9 + 7
        self.max_prod = 0

        def total_sum(node):
            if not node:
                return 0
            return node.val + total_sum(node.left) + total_sum(node.right)

        def dfs(node):
            if not node:
                return 0

            left = dfs(node.left)
            right = dfs(node.right)

            subtree = node.val + left + right
            self.max_prod = max(self.max_prod, subtree * (total - subtree))
            return subtree

        total = total_sum(root)
        dfs(root)
        return self.max_prod % MOD
```

---

### Go

```go
func maxProduct(root *TreeNode) int {
    const MOD int64 = 1e9 + 7
    var totalSum int64 = 0
    var maxProd int64 = 0

    var getTotalSum func(*TreeNode) int64
    getTotalSum = func(node *TreeNode) int64 {
        if node == nil {
            return 0
        }
        return int64(node.Val) + getTotalSum(node.Left) + getTotalSum(node.Right)
    }

    var dfs func(*TreeNode) int64
    dfs = func(node *TreeNode) int64 {
        if node == nil {
            return 0
        }

        left := dfs(node.Left)
        right := dfs(node.Right)

        subtreeSum := int64(node.Val) + left + right
        product := subtreeSum * (totalSum - subtreeSum)
        if product > maxProd {
            maxProd = product
        }

        return subtreeSum
    }

    totalSum = getTotalSum(root)
    dfs(root)

    return int(maxProd % MOD)
}
```

---

## Step-by-step Detailed Explanation

1. I calculate the total sum of all nodes.
2. I run DFS again to calculate subtree sums.
3. Each subtree represents a possible cut.
4. I calculate the product using the remaining tree.
5. I store the maximum product.
6. I apply modulo at the end.

---

## Examples

**Input**

```bash
root = [1,2,3,4,5,6]
```

**Output**

```bash
110
```

**Explanation**
Split sums are 11 and 10
Product = 11 × 10 = 110

---

## How to use / Run locally

1. Clone the repository
2. Copy the solution into your LeetCode editor
3. Run with sample test cases
4. Submit and verify

---

## Notes & Optimizations

* DFS is optimal for tree problems
* Avoid applying modulo during comparison
* Use `long long` or `int64` for large values
* Works efficiently even for large trees

---

## Author

* **Md Aarzoo Islam**
  [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
