# 2265. Count Nodes Equal to Average of Subtree

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

## Problem Summary

Given the root of a binary tree, the goal is to count how many nodes have a value equal to the average of all values in their own subtree.

For every node, its subtree includes:

* The node itself
* Every node below it
* All of its descendants

The average is calculated as:

```text
average = sum of subtree values / number of nodes in subtree
```

The result is rounded down to the nearest integer.

For example:

```text
        4
       / \
      8   5
     / \   \
    0   1   6
```

For node `5`:

```text
Subtree values = [5, 6]
Sum = 11
Count = 2

Average = 11 / 2 = 5
```

Since the value of the node is `5`, this node is counted.

The input is the root of a binary tree, and the output is an integer representing the number of nodes satisfying this condition.

This solution uses a binary tree, depth-first search, recursion, and postorder traversal to calculate the subtree sum and subtree node count efficiently.

## Constraints

| Constraint      | Value                   |
| --------------- | ----------------------- |
| Number of nodes | `1 <= n <= 1000`        |
| Node value      | `0 <= Node.val <= 1000` |
| Tree type       | Binary tree             |

## Intuition

I noticed that calculating the average of a subtree only needs two pieces of information:

```text
1. Sum of all values in the subtree
2. Number of nodes in the subtree
```

So instead of calculating the entire subtree again for every node, I can calculate this information once and pass it to the parent.

I use postorder traversal because I need to process the children before the current node.

For a node, I first calculate:

```text
Left subtree  -> sum, count
Right subtree -> sum, count
```

Then I add the current node:

```text
sum   = leftSum + rightSum + node.val
count = leftCount + rightCount + 1
```

Now I can calculate the average and check whether it matches the current node's value.

## Approach

I solve the problem using recursive depth-first search.

For every node:

1. Recursively process the left subtree.
2. Recursively process the right subtree.
3. Add the current node's value to the two subtree sums.
4. Add the current node to the two subtree counts.
5. Calculate the integer average.
6. Compare the average with the current node's value.
7. Increase the answer if they are equal.
8. Return the sum and count of the current subtree.

The important idea is that every recursive call returns:

```text
(sum, count)
```

For example:

```text
        8
       / \
      0   1
```

The two leaf nodes return:

```text
0 -> (0, 1)
1 -> (1, 1)
```

Then node `8` combines them:

```text
sum   = 0 + 1 + 8 = 9
count = 1 + 1 + 1 = 3

average = 9 / 3 = 3
```

Since `8 != 3`, node `8` is not counted.

The same process continues upward until the root is processed.

## Data Structures Used

### Binary Tree

The input itself is a binary tree represented using `TreeNode`.

Each node contains:

```text
value
left child
right child
```

### Recursion Stack

I use recursion to perform depth-first search.

Each recursive call processes one subtree and returns its sum and node count.

### Pair / Array / Tuple

Depending on the language, I return two values:

```text
sum
count
```

For example:

* C++ uses `pair<int, int>`
* Java uses an `int[]`
* JavaScript uses an array
* Python3 uses a tuple
* Go returns two integers

I do not need a separate array, map, or set for the tree.

## Operations & Behavior Summary

The algorithm works like this:

```text
Start at root
    |
    v
Process left subtree
    |
    v
Process right subtree
    |
    v
Calculate current subtree sum
    |
    v
Calculate current subtree node count
    |
    v
Calculate sum / count
    |
    v
Is average == current node value?
    |
   Yes -----> Increase answer
    |
    No
    |
    v
Return sum and count to parent
```

For a `null` node, I return:

```text
sum = 0
count = 0
```

For a normal node:

```text
sum = leftSum + rightSum + node.val
count = leftCount + rightCount + 1
```

Finally, the answer contains the number of nodes whose value equals the average of their subtree.

## Complexity

| Type             | Complexity | Explanation                                                          |
| ---------------- | ---------- | -------------------------------------------------------------------- |
| Time Complexity  | `O(n)`     | I visit every node exactly once.                                     |
| Space Complexity | `O(h)`     | `h` is the height of the tree, which is the maximum recursion depth. |

Here, `n` is the total number of nodes in the binary tree.

The space complexity is `O(h)` because I only use the recursion stack. For a balanced tree, `h` is `O(log n)`. For a completely skewed tree, `h` can be `O(n)`.

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
    int answer = 0; // Stores how many nodes have a value equal to their subtree average.

    pair<int, int> dfs(TreeNode* node) {
        if (node == nullptr) {
            return {0, 0}; // An empty subtree has sum 0 and contains 0 nodes.
        }

        auto [leftSum, leftCount] = dfs(node->left); // Get sum and count from the left subtree.
        auto [rightSum, rightCount] = dfs(node->right); // Get sum and count from the right subtree.

        int sum = leftSum + rightSum + node->val; // Add both child sums and the current node value.
        int count = leftCount + rightCount + 1; // Add both child counts and the current node.

        if (node->val == sum / count) {
            answer++; // Count this node when its value equals the floored subtree average.
        }

        return {sum, count}; // Return this subtree's sum and node count to its parent.
    }

public:
    int averageOfSubtree(TreeNode* root) {
        dfs(root); // Process every node using postorder traversal.
        return answer; // Return the total number of matching nodes.
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
    private int answer = 0; // Stores the number of nodes equal to their subtree average.

    private int[] dfs(TreeNode node) {
        if (node == null) {
            return new int[]{0, 0}; // An empty subtree has sum 0 and count 0.
        }

        int[] left = dfs(node.left); // Get the sum and count of the left subtree.
        int[] right = dfs(node.right); // Get the sum and count of the right subtree.

        int sum = left[0] + right[0] + node.val; // Combine both subtree sums with the current value.
        int count = left[1] + right[1] + 1; // Combine both subtree counts and include the current node.

        if (node.val == sum / count) {
            answer++; // Count the node when its value equals the floored average.
        }

        return new int[]{sum, count}; // Return this subtree's sum and count to its parent.
    }

    public int averageOfSubtree(TreeNode root) {
        dfs(root); // Traverse the entire tree and calculate each subtree's information.
        return answer; // Return the total number of valid nodes.
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
 * @param {TreeNode} root
 * @return {number}
 */
var averageOfSubtree = function(root) {
    let answer = 0; // Stores the number of nodes whose value equals their subtree average.

    function dfs(node) {
        if (node === null) {
            return [0, 0]; // An empty subtree has sum 0 and contains 0 nodes.
        }

        const [leftSum, leftCount] = dfs(node.left); // Get sum and count from the left subtree.
        const [rightSum, rightCount] = dfs(node.right); // Get sum and count from the right subtree.

        const sum = leftSum + rightSum + node.val; // Calculate the sum of the current subtree.
        const count = leftCount + rightCount + 1; // Calculate the number of nodes in the current subtree.

        if (node.val === Math.floor(sum / count)) {
            answer++; // Count the node when its value equals the floored average.
        }

        return [sum, count]; // Return the current subtree's sum and count to its parent.
    }

    dfs(root); // Process every node using postorder traversal.
    return answer; // Return the final count.
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
    def averageOfSubtree(self, root: TreeNode) -> int:
        answer = 0  # Stores the number of nodes equal to their subtree average.

        def dfs(node):
            nonlocal answer  # Allows the recursive function to update the outer answer variable.

            if node is None:
                return 0, 0  # An empty subtree has sum 0 and contains 0 nodes.

            left_sum, left_count = dfs(node.left)  # Get the sum and count of the left subtree.
            right_sum, right_count = dfs(node.right)  # Get the sum and count of the right subtree.

            total_sum = left_sum + right_sum + node.val  # Calculate the sum of the current subtree.
            total_count = left_count + right_count + 1  # Calculate the number of nodes in the current subtree.

            if node.val == total_sum // total_count:
                answer += 1  # Count the node when its value equals the floored average.

            return total_sum, total_count  # Return this subtree's sum and count to its parent.

        dfs(root)  # Traverse the tree and process every node.
        return answer  # Return the total number of matching nodes.
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

func averageOfSubtree(root *TreeNode) int {
    answer := 0 // Stores the number of nodes equal to their subtree average.

    var dfs func(*TreeNode) (int, int) // Defines a recursive function returning subtree sum and count.

    dfs = func(node *TreeNode) (int, int) {
        if node == nil {
            return 0, 0 // An empty subtree has sum 0 and contains 0 nodes.
        }

        leftSum, leftCount := dfs(node.Left) // Get the sum and count of the left subtree.
        rightSum, rightCount := dfs(node.Right) // Get the sum and count of the right subtree.

        sum := leftSum + rightSum + node.Val // Calculate the sum of the current subtree.
        count := leftCount + rightCount + 1 // Calculate the number of nodes in the current subtree.

        if node.Val == sum/count {
            answer++ // Count the node when its value equals the floored average.
        }

        return sum, count // Return the current subtree's sum and count to its parent.
    }

    dfs(root) // Process every node using postorder traversal.
    return answer // Return the final count.
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### C++

I use a recursive helper function that returns two values: the sum of the current subtree and the number of nodes in that subtree.

For a `nullptr`, I return `(0, 0)` because an empty subtree has no values and no nodes.

Then I recursively process the left and right children.

Once both calls return, I calculate:

```text
sum = leftSum + rightSum + currentValue
count = leftCount + rightCount + 1
```

The `+1` in `count` represents the current node.

I then compare:

```text
currentValue == sum / count
```

C++ integer division automatically removes the fractional part, which matches the floor operation here because all node values are non-negative.

If the condition is true, I increase the answer.

Finally, I return the calculated sum and count so the parent can use them.

### Java

The Java solution follows the same recursive logic.

Java does not have a built-in pair type that is convenient for this problem, so I return an `int[]` containing:

```text
[result[0] = sum]
[result[1] = count]
```

The recursive helper first handles the `null` case.

Then it gets the information from both children and combines it with the current node.

The average is calculated using integer division:

```text
sum / count
```

Since the node values are non-negative, this gives the required rounded-down average.

The answer is stored as a class variable so the recursive helper can update it directly.

### JavaScript

In JavaScript, I return an array containing the subtree sum and count:

```text
[sum, count]
```

The helper function recursively processes both children.

For a missing child, it returns:

```text
[0, 0]
```

After getting both results, I calculate the current subtree's sum and count.

JavaScript's `/` operator returns a floating-point number, so I use `Math.floor()` when comparing the average with the node value.

This keeps the behavior consistent with the problem statement.

### Python3

The Python3 solution returns a tuple:

```text
(sum, count)
```

Python makes this convenient because I can unpack the returned values directly.

For example:

```text
left_sum, left_count = dfs(node.left)
```

The helper returns `(0, 0)` when the node is `None`.

After processing both children, I calculate the total sum and count.

Python's `//` operator performs floor division, so:

```text
total_sum // total_count
```

directly gives the required average.

I use `nonlocal` for the answer because the recursive helper is defined inside `averageOfSubtree`.

### Go

The Go solution uses a recursive function that returns two integers:

```text
func dfs(node *TreeNode) (int, int)
```

The first returned value is the subtree sum, and the second is the subtree node count.

For a `nil` node, it returns:

```text
0, 0
```

The left and right recursive calls provide the information needed to calculate the current subtree.

Go's integer division automatically discards the fractional part, which is exactly what is needed here because all values are non-negative.

The answer is maintained outside the recursive helper and returned after the complete tree has been processed.

## Examples

### Example 1

Input:

```text
root = [4,8,5,0,1,null,6]
```

Tree:

```text
        4
       / \
      8   5
     / \   \
    0   1   6
```

Important subtree calculations:

```text
Node 0:
sum = 0
count = 1
average = 0
0 == 0 -> counted

Node 1:
sum = 1
count = 1
average = 1
1 == 1 -> counted

Node 6:
sum = 6
count = 1
average = 6
6 == 6 -> counted

Node 8:
sum = 9
count = 3
average = 3
8 != 3 -> not counted

Node 5:
sum = 11
count = 2
average = 5
5 == 5 -> counted

Node 4:
sum = 24
count = 6
average = 4
4 == 4 -> counted
```

Output:

```text
5
```

### Example 2

Input:

```text
root = [1]
```

Tree:

```text
    1
```

There is only one node:

```text
sum = 1
count = 1
average = 1
```

The node value is equal to the average.

Output:

```text
1
```

### Example 3

Consider this tree:

```text
        2
       / \
      1   4
```

For node `1`:

```text
sum = 1
count = 1
average = 1
1 == 1 -> counted
```

For node `4`:

```text
sum = 4
count = 1
average = 4
4 == 4 -> counted
```

For node `2`:

```text
sum = 2 + 1 + 4 = 7
count = 3

average = 7 / 3 = 2
```

So node `2` is also counted.

Output:

```text
3
```

## How to Use / Run Locally

The code shown in this repository is intended for the LeetCode problem format, where the `TreeNode` structure and test cases are provided by the platform.

If you want to run the solution locally, you need to create the `TreeNode` structure and build a test tree first.

### C++

Save the solution in a `.cpp` file.

Compile it with:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it with:

```bash
./solution
```

### Java

Save the solution in a Java file.

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

You will need a `main` method for local testing because LeetCode provides the test runner automatically.

### JavaScript

Save the solution as:

```text
solution.js
```

Run it with Node.js:

```bash
node solution.js
```

For local testing, create the required `TreeNode` objects and call `averageOfSubtree()`.

### Python3

Save the solution as:

```text
solution.py
```

Run it with:

```bash
python3 solution.py
```

For local execution, add a small test section that creates the binary tree and calls the solution method.

### Go

Save the solution as:

```text
solution.go
```

Run it with:

```bash
go run solution.go
```

The local Go program should include a `main` function for testing.

## Notes & Optimizations

The main optimization is avoiding repeated subtree traversal.

A straightforward but slower approach would be to calculate the sum and number of nodes separately for every node. In the worst case, this can repeat the same work many times.

Instead, I calculate both values during one postorder traversal.

Another important point is integer division. The problem asks for the average rounded down. Since every node value is between `0` and `1000`, all values are non-negative, so integer division gives the correct floor value.

A single node is also an important edge case:

```text
    1
```

Its subtree has one node, so:

```text
1 / 1 = 1
```

Therefore, every leaf node always satisfies the condition because its subtree contains only itself.

The maximum possible subtree sum is:

```text
1000 nodes × 1000 value = 1,000,000
```

So standard integer types are more than enough for this problem.

The core idea is to return `(sum, count)` from every recursive call. Once this pattern becomes familiar, it can also be used for many other binary tree problems involving subtree sums, subtree sizes, averages, and other aggregate values.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
