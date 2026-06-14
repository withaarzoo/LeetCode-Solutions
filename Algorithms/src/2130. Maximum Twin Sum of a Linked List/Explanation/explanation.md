# 2130. Maximum Twin Sum of a Linked List

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

LeetCode 2130: Maximum Twin Sum of a Linked List is a classic Linked List and Two Pointers problem.

We are given the head of a linked list with an even number of nodes. Every node has a twin node located the same distance from the opposite end of the list.

For a list of length `n`:

* Node `0` is paired with node `n - 1`
* Node `1` is paired with node `n - 2`
* Node `2` is paired with node `n - 3`
* and so on

The twin sum is simply the sum of the values of a node and its twin.

The goal is to find and return the maximum twin sum among all possible twin pairs.

This problem is commonly solved using linked list reversal, fast and slow pointers, and in-place traversal techniques.

## Constraints

| Constraint       | Value                   |
| ---------------- | ----------------------- |
| Number of nodes  | `2 <= n <= 10^5`        |
| `n` is           | Even                    |
| Node value range | `1 <= Node.val <= 10^5` |

## Intuition

My first thought was to store all node values in an array.

Once the values are inside an array, finding twin sums becomes easy because I can access both ends using indices.

However, that requires extra memory.

Since this is a linked list problem, I looked for a way to solve it using constant extra space.

I noticed that if I reverse the second half of the linked list, then each node in the first half lines up directly with its twin node in the reversed half.

After that, I only need one traversal to calculate every twin sum and keep track of the maximum value.

## Approach

1. Find the middle of the linked list using the fast and slow pointer technique.
2. Reverse the second half of the linked list.
3. Keep one pointer at the start of the first half.
4. Keep another pointer at the start of the reversed second half.
5. Move both pointers forward together.
6. Calculate the twin sum for each pair.
7. Continuously update the maximum twin sum.
8. Return the final answer.

This approach avoids using extra arrays and keeps memory usage minimal.

## Data Structures Used

### Linked List

The input is already provided as a singly linked list.

### Pointers

I use multiple pointers for:

* Finding the middle node
* Reversing the second half
* Traversing both halves simultaneously

No additional data structures such as arrays, stacks, queues, or hash maps are required.

## Operations & Behavior Summary

1. Start from the head of the linked list.
2. Use fast and slow pointers to locate the midpoint.
3. Reverse all nodes from the midpoint to the end.
4. Create two traversals:

   * One from the beginning
   * One from the reversed second half
5. Calculate every twin pair sum.
6. Track the largest sum encountered.
7. Return the maximum value after processing all pairs.

## Complexity

| Metric           | Complexity | Explanation                                                                           |
| ---------------- | ---------- | ------------------------------------------------------------------------------------- |
| Time Complexity  | O(n)       | Finding the middle, reversing the list, and computing twin sums each take linear time |
| Space Complexity | O(1)       | Only a few pointer variables are used, with no extra data structures                  |

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
    int pairSum(ListNode* head) {
        // Find the middle of the linked list
        ListNode* slow = head;
        ListNode* fast = head;

        while (fast && fast->next) {
            slow = slow->next;
            fast = fast->next->next;
        }

        // Reverse the second half
        ListNode* prev = nullptr;
        while (slow) {
            ListNode* nextNode = slow->next;
            slow->next = prev;
            prev = slow;
            slow = nextNode;
        }

        // Compare first half and reversed second half
        int ans = 0;
        ListNode* first = head;
        ListNode* second = prev;

        while (second) {
            ans = max(ans, first->val + second->val);

            first = first->next;
            second = second->next;
        }

        return ans;
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
    public int pairSum(ListNode head) {

        // Find the middle of the linked list
        ListNode slow = head;
        ListNode fast = head;

        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }

        // Reverse the second half
        ListNode prev = null;

        while (slow != null) {
            ListNode nextNode = slow.next;
            slow.next = prev;
            prev = slow;
            slow = nextNode;
        }

        // Calculate maximum twin sum
        int ans = 0;
        ListNode first = head;
        ListNode second = prev;

        while (second != null) {
            ans = Math.max(ans, first.val + second.val);

            first = first.next;
            second = second.next;
        }

        return ans;
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
 * @return {number}
 */
var pairSum = function(head) {

    // Find the middle of the linked list
    let slow = head;
    let fast = head;

    while (fast && fast.next) {
        slow = slow.next;
        fast = fast.next.next;
    }

    // Reverse the second half
    let prev = null;

    while (slow) {
        let nextNode = slow.next;
        slow.next = prev;
        prev = slow;
        slow = nextNode;
    }

    // Calculate maximum twin sum
    let ans = 0;
    let first = head;
    let second = prev;

    while (second) {
        ans = Math.max(ans, first.val + second.val);

        first = first.next;
        second = second.next;
    }

    return ans;
};
```

### Python3

```python
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def pairSum(self, head: Optional[ListNode]) -> int:

        # Find the middle of the linked list
        slow = head
        fast = head

        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next

        # Reverse the second half
        prev = None

        while slow:
            next_node = slow.next
            slow.next = prev
            prev = slow
            slow = next_node

        # Calculate maximum twin sum
        ans = 0
        first = head
        second = prev

        while second:
            ans = max(ans, first.val + second.val)

            first = first.next
            second = second.next

        return ans
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
func pairSum(head *ListNode) int {

    // Find the middle of the linked list
    slow := head
    fast := head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // Reverse the second half
    var prev *ListNode = nil

    for slow != nil {
        nextNode := slow.Next
        slow.Next = prev
        prev = slow
        slow = nextNode
    }

    // Calculate maximum twin sum
    ans := 0
    first := head
    second := prev

    for second != nil {
        sum := first.Val + second.Val

        if sum > ans {
            ans = sum
        }

        first = first.Next
        second = second.Next
    }

    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same in all five languages.

### Step 1: Find the Middle

I start by creating two pointers:

* Slow pointer
* Fast pointer

The slow pointer moves one node at a time.

The fast pointer moves two nodes at a time.

When the fast pointer reaches the end of the list, the slow pointer naturally reaches the middle.

This allows me to split the linked list into two equal halves without counting nodes.

### Step 2: Reverse the Second Half

Once I reach the middle, I reverse all nodes from that point onward.

This is the key observation of the solution.

Originally, twin nodes are located at opposite ends of the list.

After reversing the second half, every twin node becomes directly aligned with its corresponding node in the first half.

That means both can now be processed together in a simple forward traversal.

### Step 3: Start Two Traversals

I create:

* One pointer starting from the head
* One pointer starting from the head of the reversed half

Now both pointers point to matching twin nodes.

### Step 4: Compute Twin Sums

For every step:

1. Add the values from both pointers.
2. Compare the result with the current maximum.
3. Update the maximum if necessary.
4. Move both pointers forward.

Since both halves contain exactly the same number of nodes, every twin pair is processed once.

### Step 5: Return the Maximum Value

After all pairs are checked, the stored maximum value is the answer.

### Why Reversing Works

Consider:

```text
1 -> 2 -> 3 -> 4
```

Twin pairs:

```text
1 + 4
2 + 3
```

After reversing the second half:

```text
1 -> 2

4 -> 3
```

Now:

```text
1 pairs with 4
2 pairs with 3
```

The problem becomes much easier because both pairs can be visited together.

### Edge Cases

#### Smallest Valid List

```text
[1,100000]
```

Only one twin pair exists.

Answer:

```text
100001
```

#### All Values Equal

```text
[5,5,5,5]
```

Every twin sum is identical.

The algorithm still processes every pair correctly.

#### Maximum Input Size

The solution remains efficient even for `100000` nodes because every node is visited only a constant number of times.

## Examples

### Example 1

Input

```text
head = [5,4,2,1]
```

Twin pairs

```text
5 + 1 = 6
4 + 2 = 6
```

Output

```text
6
```

### Example 2

Input

```text
head = [4,2,2,3]
```

Twin pairs

```text
4 + 3 = 7
2 + 2 = 4
```

Output

```text
7
```

### Example 3

Input

```text
head = [1,100000]
```

Twin pairs

```text
1 + 100000 = 100001
```

Output

```text
100001
```

## How to Use / Run Locally

### C++

Compile

```bash
g++ solution.cpp -o solution
```

Run

```bash
./solution
```

### Java

Compile

```bash
javac Solution.java
```

Run

```bash
java Solution
```

### JavaScript

Run

```bash
node solution.js
```

### Python3

Run

```bash
python solution.py
```

### Go

Run

```bash
go run solution.go
```

## Notes & Optimizations

* An array-based solution also works but requires O(n) extra memory.
* Reversing the second half allows the problem to be solved with O(1) extra space.
* Every node is processed only a few times, making the solution highly efficient.
* The fast and slow pointer technique is a common pattern for linked list problems.
* In-place reversal is the main optimization that improves memory usage.
* This approach is considered the optimal solution for LeetCode 2130 Maximum Twin Sum of a Linked List.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
