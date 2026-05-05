# Rotate List (LeetCode 61)

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

This problem is about rotating a singly linked list to the right by `k` positions.

You are given the head of a linked list and an integer `k`. You need to move the last `k` nodes of the list to the front while keeping the order intact.

In simple words:

* Take the last `k` elements
* Put them at the beginning
* Keep everything else in order

The output should be the new head of the modified linked list.

This is a classic linked list problem that tests pointer manipulation and understanding of list structure.

## Constraints

* Number of nodes in the list: `0 <= n <= 500`
* Node value range: `-100 <= Node.val <= 100`
* Rotation value: `0 <= k <= 2 * 10^9`

## Intuition

At first, rotating the list one step at a time feels natural. But doing that `k` times is inefficient, especially when `k` is very large.

I noticed something important:
Rotating `k` times is the same as rotating `k % n` times.

Then I thought about how to avoid physically shifting nodes again and again. The trick is to convert the list into a circular linked list. Once it's circular, rotation becomes just a matter of breaking the list at the correct point.

That idea simplifies everything.

## Approach

Here’s how I broke it down step by step:

1. Handle edge cases:

   * If the list is empty
   * If it has only one node
   * If `k` is 0

2. Traverse the list once to:

   * Count total nodes (`n`)
   * Reach the last node

3. Connect the last node back to the head

   * This makes the list circular

4. Reduce unnecessary rotations:

   * `k = k % n`

5. Find the new tail:

   * Move `(n - k - 1)` steps from the head

6. The next node becomes the new head

7. Break the circular link:

   * Set `newTail.next = null`

8. Return the new head

## Data Structures Used

* Singly Linked List
  Used to represent the input. The whole problem is about modifying this structure efficiently.

* Pointers (References)
  Used to traverse nodes, track tail, and reassign links.

No extra data structures are used, which keeps space usage minimal.

## Operations & Behavior Summary

* Traverse list to calculate length
* Connect tail to head (make circular)
* Reduce rotation using modulo
* Move pointer to find new tail
* Break the circle
* Return new head

This avoids repeated rotations and keeps the solution efficient.

## Complexity

| Type             | Complexity | Explanation                                                   |
| ---------------- | ---------- | ------------------------------------------------------------- |
| Time Complexity  | O(n)       | I traverse the list at most twice. `n` is the number of nodes |
| Space Complexity | O(1)       | No extra memory is used apart from a few pointers             |

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
    ListNode* rotateRight(ListNode* head, int k) {
        // edge case: empty list or single node
        if (!head || !head->next || k == 0) return head;

        int n = 1; // length of list
        ListNode* tail = head;

        // find length and last node
        while (tail->next) {
            tail = tail->next;
            n++;
        }

        // make it circular
        tail->next = head;

        // reduce k
        k = k % n;

        // find new tail (n - k - 1 steps)
        int steps = n - k - 1;
        ListNode* newTail = head;

        while (steps--) {
            newTail = newTail->next;
        }

        // new head is next of newTail
        ListNode* newHead = newTail->next;

        // break the circle
        newTail->next = nullptr;

        return newHead;
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
    public ListNode rotateRight(ListNode head, int k) {
        // edge case
        if (head == null || head.next == null || k == 0) return head;

        int n = 1;
        ListNode tail = head;

        // find length
        while (tail.next != null) {
            tail = tail.next;
            n++;
        }

        // make circular
        tail.next = head;

        // reduce k
        k = k % n;

        int steps = n - k - 1;
        ListNode newTail = head;

        // move to new tail
        while (steps-- > 0) {
            newTail = newTail.next;
        }

        ListNode newHead = newTail.next;

        // break circle
        newTail.next = null;

        return newHead;
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
var rotateRight = function(head, k) {
    // edge case
    if (!head || !head.next || k === 0) return head;

    let n = 1;
    let tail = head;

    // find length
    while (tail.next) {
        tail = tail.next;
        n++;
    }

    // make circular
    tail.next = head;

    // reduce k
    k = k % n;

    let steps = n - k - 1;
    let newTail = head;

    // move to new tail
    while (steps-- > 0) {
        newTail = newTail.next;
    }

    let newHead = newTail.next;

    // break circle
    newTail.next = null;

    return newHead;
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
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        # edge case
        if not head or not head.next or k == 0:
            return head

        n = 1
        tail = head

        # find length
        while tail.next:
            tail = tail.next
            n += 1

        # make circular
        tail.next = head

        # reduce k
        k = k % n

        steps = n - k - 1
        newTail = head

        # move to new tail
        while steps > 0:
            newTail = newTail.next
            steps -= 1

        newHead = newTail.next

        # break circle
        newTail.next = None

        return newHead
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
func rotateRight(head *ListNode, k int) *ListNode {
    // edge case
    if head == nil || head.Next == nil || k == 0 {
        return head
    }

    n := 1
    tail := head

    // find length
    for tail.Next != nil {
        tail = tail.Next
        n++
    }

    // make circular
    tail.Next = head

    // reduce k
    k = k % n

    steps := n - k - 1
    newTail := head

    // move to new tail
    for steps > 0 {
        newTail = newTail.Next
        steps--
    }

    newHead := newTail.Next

    // break circle
    newTail.Next = nil

    return newHead
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same across all languages. Only syntax changes.

I start by checking edge cases. If the list is empty or has one node, there is nothing to rotate.

Next, I walk through the list to find its length. At the same time, I keep track of the last node.

Once I reach the end, I connect it back to the head. Now the list becomes circular. This is the key trick.

Then I reduce `k` using modulo. This avoids unnecessary full rotations.

Now I need to find the new tail. I move `(n - k - 1)` steps from the head. This position determines where the list will be cut.

The node after this becomes the new head.

Finally, I break the circular link. This restores the normal singly linked list structure.

If I skip breaking the circle, the list would loop forever.

This approach avoids repeated shifting and works in linear time.

## Examples

### Example 1

Input:
head = [1,2,3,4,5], k = 2

Output:
[4,5,1,2,3]

Explanation:

* Length = 5
* Effective rotation = 2
* New tail is at position 2 (value = 3)
* New head = 4

---

### Example 2

Input:
head = [0,1,2], k = 4

Output:
[2,0,1]

Explanation:

* Length = 3
* Effective rotation = 4 % 3 = 1
* Rotate once
* Result becomes [2,0,1]

---

### Example 3

Input:
head = [1], k = 10

Output:
[1]

Explanation:

* Only one node
* Rotation does not change anything

## How to Use / Run Locally

### C++

1. Save code in `main.cpp`
2. Compile using:

   ```
   g++ main.cpp -o output
   ```

3. Run:

   ```
   ./output
   ```

### Java

1. Save in `Solution.java`
2. Compile:

   ```
   javac Solution.java
   ```

3. Run:

   ```
   java Solution
   ```

### JavaScript

1. Save in `solution.js`
2. Run using Node.js:

   ```
   node solution.js
   ```

### Python3

1. Save in `solution.py`
2. Run:

   ```
   python3 solution.py
   ```

### Go

1. Save in `main.go`
2. Run:

   ```
   go run main.go
   ```

## Notes & Optimizations

* Always reduce `k` using modulo to avoid unnecessary work
* Converting to a circular linked list makes the logic simpler
* Be careful while breaking the loop, otherwise it can cause infinite traversal
* Edge cases like empty list or single node should be handled early
* This is an optimal solution with O(n) time and O(1) space

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
