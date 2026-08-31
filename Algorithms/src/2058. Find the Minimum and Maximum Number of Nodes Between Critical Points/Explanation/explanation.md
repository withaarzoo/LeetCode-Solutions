# 2058. Find the Minimum and Maximum Number of Nodes Between Critical Points

A beginner-friendly solution for **LeetCode 2058 – Find the Minimum and Maximum Number of Nodes Between Critical Points**. This repository explains the intuition, approach, complexity, examples, and provides solutions in **C++, Java, JavaScript, Python, and Go**.

This solution uses a **single traversal of the linked list** to find all critical points while keeping the extra space constant. If you're preparing for coding interviews or learning linked list problems in Data Structures and Algorithms (DSA), this is a great problem to practice.

---

## Table of Contents

- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
  - [C++](#c)
  - [Java](#java)
  - [JavaScript](#javascript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

You are given the **head of a singly linked list**.

A node is called a **critical point** if it is either:

- A **local maximum**, meaning its value is strictly greater than both the previous and next nodes.
- A **local minimum**, meaning its value is strictly smaller than both the previous and next nodes.

The first and last nodes can never be critical points because they do not have both neighbors.

Your task is to return an array containing two values:

- **Minimum distance** between any two different critical points.
- **Maximum distance** between any two different critical points.

If the linked list contains fewer than two critical points, return `[-1, -1]`.

### Input

- Head of a singly linked list.

### Output

- An integer array `[minDistance, maxDistance]`.

This is a classic **Linked List + One Pass Traversal** problem that appears frequently in coding interviews.

---

## Constraints

| Constraint | Value |
|------------|-------|
| Number of nodes | `2 <= n <= 10^5` |
| Node value | `1 <= Node.val <= 10^5` |

---

## Intuition

The first thing I noticed was that a node can only be a critical point if it has both a previous node and a next node. So I only need to examine the middle nodes of the linked list.

While traversing the list, I compare each node with its neighbors.

If I find a critical point, I do not need to store every position in a separate array. I only need to remember:

- The first critical point.
- The previous critical point.
- The latest critical point.
- The smallest distance found so far.

The maximum distance will always be between the **first** and **last** critical points, while the minimum distance is always found between **consecutive** critical points.

That observation lets me solve the entire problem in one traversal.

---

## Approach

I solve the problem using a single pass through the linked list.

### Step 1

Start traversing from the **second node**, because the first node cannot be a critical point.

### Step 2

For every node, compare its value with the previous and next nodes.

A node is critical if:

- It is greater than both neighbors.
- It is smaller than both neighbors.

### Step 3

Whenever I find a critical point:

- Save the position if it is the first one.
- Calculate the distance from the previous critical point.
- Update the minimum distance if needed.
- Update the latest critical point position.

### Step 4

After the traversal:

- If fewer than two critical points exist, return `[-1, -1]`.
- Otherwise:
  - Maximum distance = last critical position − first critical position.
  - Return both distances.

This approach visits every node exactly once and uses constant extra memory.

---

## Data Structures Used

| Data Structure | Why I Used It |
| --------------- | --------------- |
| Singly Linked List | The input structure provided in the problem. |
| Integer Variables | To store node positions and track minimum and maximum distances. |
| Result Array | To return the final answer in the required format. |

No additional arrays or hash maps are needed.

---

## Operations & Behavior Summary

Here is what the algorithm does during traversal.

1. Start from the second node.
2. Keep track of the current node index.
3. Compare the current node with its neighbors.
4. Detect whether it is a local maximum or local minimum.
5. Record the first critical point.
6. For every new critical point:
   - Calculate distance from the previous critical point.
   - Update the minimum distance.
   - Update the latest critical point.
7. Continue until the end of the linked list.
8. Return `[-1, -1]` if fewer than two critical points exist.
9. Otherwise return `[minimum distance, maximum distance]`.

This behavior guarantees the answer in one pass.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | **O(n)** | I traverse the linked list only once, where `n` is the number of nodes. |
| Space Complexity | **O(1)** | I only use a few integer variables and no extra data structures proportional to the input size. |

---

## Multi-language Solutions

### C++

```cpp
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    vector<int> nodesBetweenCriticalPoints(ListNode* head) {
        // The first critical point position; -1 means we have not found one yet.
        int first = -1;

        // The most recently found critical point position.
        int last = -1;

        // The minimum distance between two consecutive critical points.
        int minDistance = INT_MAX;

        // Position of the current node in the linked list.
        int position = 1;

        // Start from the second node because a critical point needs a previous node.
        ListNode* prev = head;

        // The current node starts from the second node.
        ListNode* curr = head->next;

        // We need curr->next to exist, so the last node is not checked.
        while (curr != nullptr && curr->next != nullptr) {
            // Check whether curr is a local maximum or a local minimum.
            bool isCritical =
                (curr->val > prev->val && curr->val > curr->next->val) ||
                (curr->val < prev->val && curr->val < curr->next->val);

            // Process the node only if it is a critical point.
            if (isCritical) {
                // This is the first critical point we have found.
                if (first == -1) {
                    first = position;
                } else {
                    // Compare this critical point with the previous critical point.
                    minDistance = min(minDistance, position - last);
                }

                // Store the current critical point as the latest one.
                last = position;
            }

            // Move the previous pointer forward for the next iteration.
            prev = curr;

            // Move the current pointer forward for the next iteration.
            curr = curr->next;

            // Move to the next position in the linked list.
            position++;
        }

        // Fewer than two critical points means no valid distance exists.
        if (first == -1 || first == last) {
            return {-1, -1};
        }

        // The distance between the first and last critical points is the maximum.
        int maxDistance = last - first;

        // Return the minimum and maximum distances.
        return {minDistance, maxDistance};
    }
};
```

### Java

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public int[] nodesBetweenCriticalPoints(ListNode head) {
        // Store the position of the first critical point.
        int first = -1;

        // Store the position of the most recent critical point.
        int last = -1;

        // Start with the largest possible integer for the minimum distance.
        int minDistance = Integer.MAX_VALUE;

        // The current node starts at position 1 because head is position 0.
        int position = 1;

        // The previous node is initially the head.
        ListNode prev = head;

        // The current node starts from the second node.
        ListNode curr = head.next;

        // The last node cannot be critical because it has no next node.
        while (curr != null && curr.next != null) {
            // Check whether the current node is a local maximum or local minimum.
            boolean isCritical =
                (curr.val > prev.val && curr.val > curr.next.val) ||
                (curr.val < prev.val && curr.val < curr.next.val);

            // Process the current node if it is a critical point.
            if (isCritical) {
                // Save the position when this is the first critical point.
                if (first == -1) {
                    first = position;
                } else {
                    // Update the minimum distance using consecutive critical points.
                    minDistance = Math.min(minDistance, position - last);
                }

                // Make the current critical point the latest critical point.
                last = position;
            }

            // Move prev to the current node for the next iteration.
            prev = curr;

            // Move curr to the next node.
            curr = curr.next;

            // Move to the next position.
            position++;
        }

        // If there are fewer than two critical points, no distance can be calculated.
        if (first == -1 || first == last) {
            return new int[]{-1, -1};
        }

        // The distance between the first and last critical points is the maximum.
        int maxDistance = last - first;

        // Return the minimum and maximum distances.
        return new int[]{minDistance, maxDistance};
    }
}
```

### JavaScript

```javascript
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {number[]}
 */
var nodesBetweenCriticalPoints = function(head) {
    // Store the position of the first critical point.
    let first = -1;

    // Store the position of the most recent critical point.
    let last = -1;

    // Start with Infinity so the first valid distance can replace it.
    let minDistance = Infinity;

    // The current node starts at position 1 because head is position 0.
    let position = 1;

    // prev starts at the head because the current node needs a previous node.
    let prev = head;

    // curr starts from the second node.
    let curr = head.next;

    // The last node cannot be critical because it has no next node.
    while (curr !== null && curr.next !== null) {
        // Check whether curr is a local maximum or local minimum.
        const isCritical =
            (curr.val > prev.val && curr.val > curr.next.val) ||
            (curr.val < prev.val && curr.val < curr.next.val);

        // Process the node only when it is a critical point.
        if (isCritical) {
            // Save this position if it is the first critical point.
            if (first === -1) {
                first = position;
            } else {
                // Update the minimum distance from the previous critical point.
                minDistance = Math.min(minDistance, position - last);
            }

            // Store this critical point as the latest one.
            last = position;
        }

        // Move prev forward for the next comparison.
        prev = curr;

        // Move curr forward to the next node.
        curr = curr.next;

        // Move to the next position.
        position++;
    }

    // Fewer than two critical points means no valid answer exists.
    if (first === -1 || first === last) {
        return [-1, -1];
    }

    // The first-to-last distance is always the maximum distance.
    const maxDistance = last - first;

    // Return the minimum and maximum distances.
    return [minDistance, maxDistance];
};
```

### Python3

```python
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

from typing import Optional, List

class Solution:
    def nodesBetweenCriticalPoints(self, head: Optional[ListNode]) -> List[int]:
        # Store the position of the first critical point.
        first = -1

        # Store the position of the most recent critical point.
        last = -1

        # Start with infinity so the first valid distance can replace it.
        min_distance = float("inf")

        # The current node starts at position 1 because head is position 0.
        position = 1

        # prev starts at head because curr needs a previous node.
        prev = head

        # curr starts from the second node.
        curr = head.next

        # The last node cannot be critical because it has no next node.
        while curr is not None and curr.next is not None:
            # Check whether curr is a local maximum or local minimum.
            is_critical = (
                (curr.val > prev.val and curr.val > curr.next.val)
                or
                (curr.val < prev.val and curr.val < curr.next.val)
            )

            # Process the node only if it is a critical point.
            if is_critical:
                # Save this position if it is the first critical point.
                if first == -1:
                    first = position
                else:
                    # Update the minimum distance from the previous critical point.
                    min_distance = min(min_distance, position - last)

                # Store this critical point as the latest one.
                last = position

            # Move prev forward for the next comparison.
            prev = curr

            # Move curr forward to the next node.
            curr = curr.next

            # Move to the next position.
            position += 1

        # Fewer than two critical points means no valid answer exists.
        if first == -1 or first == last:
            return [-1, -1]

        # The first-to-last distance is always the maximum distance.
        max_distance = last - first

        # Return the minimum and maximum distances.
        return [min_distance, max_distance]
```

### Go

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
 // Store the position of the first critical point.
 first := -1

 // Store the position of the most recent critical point.
 last := -1

 // Use the largest integer value so the first valid distance can replace it.
 minDistance := int(^uint(0) >> 1)

 // The current node starts at position 1 because head is position 0.
 position := 1

 // prev starts at head because the current node needs a previous node.
 prev := head

 // curr starts from the second node.
 curr := head.Next

 // The last node cannot be critical because it has no next node.
 for curr != nil && curr.Next != nil {
  // Check whether curr is a local maximum or local minimum.
  isCritical := (curr.Val > prev.Val && curr.Val > curr.Next.Val) ||
   (curr.Val < prev.Val && curr.Val < curr.Next.Val)

  // Process the node only if it is a critical point.
  if isCritical {
   // Save this position if it is the first critical point.
   if first == -1 {
    first = position
   } else {
    // Calculate the distance from the previous critical point.
    distance := position - last

    // Keep the smallest distance found so far.
    if distance < minDistance {
     minDistance = distance
    }
   }

   // Store this critical point as the latest one.
   last = position
  }

  // Move prev forward for the next comparison.
  prev = curr

  // Move curr forward to the next node.
  curr = curr.Next

  // Move to the next position.
  position++
 }

 // Fewer than two critical points means no valid distance exists.
 if first == -1 || first == last {
  return []int{-1, -1}
 }

 // The distance between the first and last critical points is the maximum.
 maxDistance := last - first

 // Return the minimum and maximum distances.
 return []int{minDistance, maxDistance}
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical in all five languages. Only the syntax changes.

### C++

I keep three pointers while traversing the linked list:

- Previous node.
- Current node.
- Next node.

The traversal starts from the second node because every critical point needs a previous node.

Every time I find a critical point, I update the first position, previous position, minimum distance, and latest position.

### Java

The Java solution follows the same traversal pattern.

Instead of pointers, Java uses object references to move through the linked list.

The minimum distance is initialized with `Integer.MAX_VALUE` so that the first valid distance replaces it immediately.

### JavaScript

JavaScript uses object references for linked list nodes.

I move through the list using `prev`, `curr`, and `curr.next`.

`Infinity` is used for the initial minimum distance because any real distance will be smaller.

### Python3

Python makes the traversal very readable.

I use `float("inf")` as the initial minimum distance and update it whenever a smaller consecutive distance appears.

The algorithm still uses constant extra space.

### Go

The Go implementation follows the same one-pass traversal.

Pointers are moved node by node while integer variables store the positions of critical points.

Go uses the maximum integer value as the initial minimum distance.

### Why Only Consecutive Critical Points Matter

Suppose the critical point positions are:

```text
2, 5, 9, 15
```

Possible distances are:

| Pair | Distance |
| ------ | ---------- |
| 2 → 5 | 3 |
| 5 → 9 | 4 |
| 9 → 15 | 6 |
| 2 → 9 | 7 |
| 2 → 15 | 13 |
| 5 → 15 | 10 |

The smallest distance must appear between consecutive critical points.

The largest distance must appear between the first and last critical points.

This observation removes the need to compare every pair.

### Edge Cases Covered

#### Case 1 — No Critical Points

Input

```text
[3,1]
```

Output

```text
[-1,-1]
```

The list has fewer than two critical points.

#### Case 2 — Exactly One Critical Point

Input

```text
[1,3,2]
```

Output

```text
[-1,-1]
```

Only one critical point exists.

#### Case 3 — Multiple Critical Points

Input

```text
[5,3,1,2,5,1,2]
```

Output

```text
[1,3]
```

The algorithm updates both minimum and maximum distances during traversal.

---

## Examples

### Example 1

**Input**

```text
head = [3,1]
```

**Output**

```text
[-1,-1]
```

**Explanation**

There are no critical points because neither node has both neighbors.

---

### Example 2

**Input**

```text
head = [5,3,1,2,5,1,2]
```

**Output**

```text
[1,3]
```

**Trace**

| Position | Value | Critical Point |
| ---------- | ------- | ---------------- |
| 2 | 1 | Local Minimum |
| 4 | 5 | Local Maximum |
| 5 | 1 | Local Minimum |

Distances between consecutive critical points:

- `4 - 2 = 2`
- `5 - 4 = 1`

Minimum distance becomes `1`.

Maximum distance is `5 - 2 = 3`.

---

### Example 3

**Input**

```text
head = [1,3,2,2,3,2,2,2,7]
```

**Output**

```text
[3,3]
```

**Trace**

Critical points appear at positions `1` and `4`.

Distance between them:

```text
4 - 1 = 3
```

Both minimum and maximum distances are `3`.

---

## How to Use / Run Locally

Clone this repository first.

```bash
git clone https://github.com/your-username/find-minimum-and-maximum-number-of-nodes-between-critical-points.git

cd find-minimum-and-maximum-number-of-nodes-between-critical-points
```

### Run C++

Compile the C++ solution.

```bash
g++ solution.cpp -o solution

./solution
```

### Run Java

Compile and run the Java solution.

```bash
javac Solution.java

java Solution
```

### Run JavaScript

Run using Node.js.

```bash
node solution.js
```

### Run Python3

Execute the Python solution.

```bash
python solution.py
```

### Run Go

Compile and run the Go solution.

```bash
go run solution.go
```

---

## Notes & Optimizations

- The first and last nodes are never checked because they cannot be critical points.
- Equal neighboring values never create a critical point because comparisons must be **strict**.
- The solution uses **one traversal**, making it efficient for linked lists with up to `100000` nodes.
- No extra array is needed to store critical point positions, so the space complexity stays constant.
- An alternative solution could store all critical point positions in an array and compute distances afterward, but that requires `O(k)` extra space, where `k` is the number of critical points.

### Common Mistakes

- Checking the first or last node as a critical point.
- Using `>=` or `<=` instead of strict comparisons.
- Calculating the minimum distance using every pair instead of consecutive critical points.
- Forgetting to return `[-1,-1]` when fewer than two critical points exist.

---

## Author

**Md Aarzoo Islam**

Instagram: <https://www.instagram.com/codewithaarzoo.in/>
