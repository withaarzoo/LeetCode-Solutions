# 1161. Maximum Level Sum of a Binary Tree

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
* Step-by-step Detailed Explanation
* Examples
* How to Use / Run Locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given the root of a binary tree.
The root is at level 1, its children are at level 2, and so on.

My task is to find **the smallest level number** where the **sum of all node values is maximum**.

If multiple levels have the same maximum sum, I must return the **smallest level index**.

---

## Constraints

* Number of nodes is between 1 and 10⁴
* Node values range from −10⁵ to 10⁵
* Tree can contain negative values
* Tree may not be balanced

---

## Intuition

When I read the problem, I realized one important thing.

This problem is not about paths or recursion depth.
It is about **levels**.

So instead of going deep into the tree, I decided to go **level by level**.

For each level:

* I add all node values
* I compare the sum with the maximum sum seen so far
* I store the level number if the sum is larger

To move level by level, **Breadth First Search (BFS)** is the best and most natural approach.

---

## Approach

1. I use a queue to perform level-order traversal.
2. I start with level 1 and push the root into the queue.
3. While the queue is not empty:

   * I calculate the number of nodes in the current level.
   * I sum all node values of this level.
   * I push left and right children into the queue.
4. If the current level sum is greater than the previous maximum:

   * I update the maximum sum.
   * I store the current level number.
5. After processing all levels, I return the stored level number.

---

## Data Structures Used

* Queue
  Used to process nodes level by level in BFS order.

---

## Operations & Behavior Summary

* Each node is visited exactly once.
* Level sums are calculated independently.
* The first maximum level is always preferred.
* Works correctly even with negative values.

---

## Complexity

**Time Complexity:**
O(n)
n is the number of nodes. Each node is processed once.

**Space Complexity:**
O(w)
w is the maximum width of the tree.
The queue stores nodes from one level at a time.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxLevelSum(TreeNode* root) {
        queue<TreeNode*> q;
        q.push(root);

        int level = 1, answerLevel = 1;
        long long maxSum = LLONG_MIN;

        while (!q.empty()) {
            int size = q.size();
            long long levelSum = 0;

            for (int i = 0; i < size; i++) {
                TreeNode* node = q.front();
                q.pop();

                levelSum += node->val;

                if (node->left) q.push(node->left);
                if (node->right) q.push(node->right);
            }

            if (levelSum > maxSum) {
                maxSum = levelSum;
                answerLevel = level;
            }

            level++;
        }

        return answerLevel;
    }
};
```

---

### Java

```java
class Solution {
    public int maxLevelSum(TreeNode root) {
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);

        int level = 1, answerLevel = 1;
        long maxSum = Long.MIN_VALUE;

        while (!queue.isEmpty()) {
            int size = queue.size();
            long levelSum = 0;

            for (int i = 0; i < size; i++) {
                TreeNode node = queue.poll();
                levelSum += node.val;

                if (node.left != null) queue.add(node.left);
                if (node.right != null) queue.add(node.right);
            }

            if (levelSum > maxSum) {
                maxSum = levelSum;
                answerLevel = level;
            }

            level++;
        }

        return answerLevel;
    }
}
```

---

### JavaScript

```javascript
var maxLevelSum = function(root) {
    let queue = [root];
    let level = 1;
    let answerLevel = 1;
    let maxSum = -Infinity;

    while (queue.length > 0) {
        let size = queue.length;
        let levelSum = 0;

        for (let i = 0; i < size; i++) {
            let node = queue.shift();
            levelSum += node.val;

            if (node.left) queue.push(node.left);
            if (node.right) queue.push(node.right);
        }

        if (levelSum > maxSum) {
            maxSum = levelSum;
            answerLevel = level;
        }

        level++;
    }

    return answerLevel;
};
```

---

### Python3

```python
class Solution:
    def maxLevelSum(self, root):
        from collections import deque

        queue = deque([root])
        level = 1
        answer_level = 1
        max_sum = float('-inf')

        while queue:
            size = len(queue)
            level_sum = 0

            for _ in range(size):
                node = queue.popleft()
                level_sum += node.val

                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

            if level_sum > max_sum:
                max_sum = level_sum
                answer_level = level

            level += 1

        return answer_level
```

---

## Step-by-step Detailed Explanation

1. I push the root node into a queue.
2. I treat each iteration of the while loop as one tree level.
3. I use the queue size to know how many nodes belong to that level.
4. I sum all node values in that level.
5. I push children nodes for the next level.
6. I update the maximum sum and level if needed.
7. After all levels are processed, I return the answer.

---

## Examples

**Example 1**

Input
`[1,7,0,7,-8,null,null]`

Level sums

* Level 1 → 1
* Level 2 → 7
* Level 3 → -1

Output
`2`

---

## How to Use / Run Locally

1. Copy the solution code for your preferred language.
2. Paste it into your local compiler or LeetCode editor.
3. Provide the binary tree input.
4. Run the program to get the result.

---

## Notes & Optimizations

* BFS is optimal for level-based problems.
* DFS would require extra tracking of levels.
* No extra arrays are needed.
* Works well within constraints.

---

## Author

* **Md Aarzoo Islam**
  [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
