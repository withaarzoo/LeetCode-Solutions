# Delete the Middle Node of a Linked List

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

Given the head of a singly linked list, the task is to delete the middle node and return the modified linked list.

The middle node is defined as the node located at index `⌊n / 2⌋` using zero-based indexing, where `n` is the total number of nodes in the list.

For example:

* If the list size is `1`, the middle node is index `0`
* If the list size is `4`, the middle node is index `2`
* If the list size is `7`, the middle node is index `3`

The goal is to remove that node efficiently and return the updated linked list.

This is a classic Linked List and Two Pointers problem commonly asked in coding interviews and competitive programming contests.

## Constraints

| Constraint      | Value                   |
| --------------- | ----------------------- |
| Number of nodes | `1 <= n <= 10^5`        |
| Node value      | `1 <= Node.val <= 10^5` |

## Intuition

My first thought was to count the total number of nodes, calculate the middle index, and then traverse the list again to remove that node.

That approach works, but it requires two complete traversals of the linked list.

I noticed that this problem is a perfect fit for the slow and fast pointer technique.

If one pointer moves one step at a time and another moves two steps at a time, the slower pointer naturally reaches the middle when the faster pointer reaches the end.

The only extra thing I need is a pointer that remembers the node before the middle node so I can reconnect the list after deletion.

This gives a clean one-pass solution with constant extra space.

## Approach

1. Handle the special case where the list contains only one node.
2. Create three pointers:

   * Slow pointer
   * Fast pointer
   * Previous pointer
3. Move:

   * Slow by one step
   * Fast by two steps
4. Keep updating the previous pointer so it always stays behind the slow pointer.
5. When the fast pointer reaches the end:

   * Slow points to the middle node.
   * Previous points to the node before the middle.
6. Remove the middle node by changing the next reference.
7. Return the original head.

## Data Structures Used

### Linked List

The input itself is a singly linked list.

I directly modify the existing list instead of creating a new one.

### Pointers / References

I use:

* Slow pointer to locate the middle node
* Fast pointer to determine when the middle is reached
* Previous pointer to reconnect the list after deletion

No additional data structures are required.

## Operations & Behavior Summary

The algorithm performs the following operations:

1. Check if the list has only one node.
2. Initialize slow, fast, and previous pointers.
3. Traverse the linked list.
4. Move slow one step at a time.
5. Move fast two steps at a time.
6. Track the node before the middle.
7. Stop when fast reaches the end.
8. Remove the middle node.
9. Return the modified linked list.

This process guarantees that the middle node is deleted in a single traversal.

## Complexity

| Metric           | Complexity | Explanation                                                     |
| ---------------- | ---------- | --------------------------------------------------------------- |
| Time Complexity  | O(n)       | Each node is visited at most once while finding the middle node |
| Space Complexity | O(1)       | Only a few pointers are used regardless of input size           |

Where:

* `n` = total number of nodes in the linked list

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
    ListNode* deleteMiddle(ListNode* head) {
        // If there is only one node, deleting the middle leaves an empty list
        if (head->next == nullptr) {
            return nullptr;
        }

        // Slow finds middle, fast moves twice as fast
        ListNode* slow = head;
        ListNode* fast = head;

        // Keeps track of node before slow
        ListNode* prev = nullptr;

        while (fast != nullptr && fast->next != nullptr) {
            prev = slow;           // Store previous node
            slow = slow->next;     // Move slow by 1 step
            fast = fast->next->next; // Move fast by 2 steps
        }

        // Skip the middle node
        prev->next = slow->next;

        return head;
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
 *     ListNode(int val, ListNode next) { this.next = next; }
 * }
 */
class Solution {
    public ListNode deleteMiddle(ListNode head) {

        // If there is only one node, return an empty list
        if (head.next == null) {
            return null;
        }

        // Slow finds middle, fast moves twice as fast
        ListNode slow = head;
        ListNode fast = head;

        // Node before slow
        ListNode prev = null;

        while (fast != null && fast.next != null) {
            prev = slow;          // Store previous node
            slow = slow.next;     // Move slow by 1 step
            fast = fast.next.next; // Move fast by 2 steps
        }

        // Remove middle node
        prev.next = slow.next;

        return head;
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
 * @return {ListNode}
 */
var deleteMiddle = function(head) {

    // If there is only one node, return empty list
    if (head.next === null) {
        return null;
    }

    // Slow finds middle, fast moves twice as fast
    let slow = head;
    let fast = head;

    // Node before slow
    let prev = null;

    while (fast !== null && fast.next !== null) {
        prev = slow;              // Store previous node
        slow = slow.next;         // Move slow by 1 step
        fast = fast.next.next;    // Move fast by 2 steps
    }

    // Remove middle node
    prev.next = slow.next;

    return head;
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
    def deleteMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:

        # If there is only one node, deleting it leaves an empty list
        if head.next is None:
            return None

        # Slow finds middle, fast moves twice as fast
        slow = head
        fast = head

        # Node before slow
        prev = None

        while fast and fast.next:
            prev = slow          # Store previous node
            slow = slow.next     # Move slow by 1 step
            fast = fast.next.next  # Move fast by 2 steps

        # Remove middle node
        prev.next = slow.next

        return head
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
func deleteMiddle(head *ListNode) *ListNode {

    // If there is only one node, return an empty list
    if head.Next == nil {
        return nil
    }

    // Slow finds middle, fast moves twice as fast
    slow := head
    fast := head

    // Node before slow
    var prev *ListNode = nil

    for fast != nil && fast.Next != nil {
        prev = slow           // Store previous node
        slow = slow.Next      // Move slow by 1 step
        fast = fast.Next.Next // Move fast by 2 steps
    }

    // Remove middle node
    prev.Next = slow.Next

    return head
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical in all five languages.

### Step 1: Handle the Single Node Case

If the linked list contains only one node, that node is also the middle node.

Deleting it leaves an empty list.

So the function immediately returns `null` (or equivalent).

### Step 2: Initialize Three References

I maintain:

* Slow pointer
* Fast pointer
* Previous pointer

Initially:

* Slow starts at the head.
* Fast starts at the head.
* Previous starts as null.

### Step 3: Move Through the List

During traversal:

* Slow moves one node forward.
* Fast moves two nodes forward.
* Previous follows slow.

This pattern ensures that slow reaches the middle exactly when fast reaches the end.

### Step 4: Identify the Middle Node

Because fast moves twice as quickly:

* For odd-length lists, slow lands on the exact center.
* For even-length lists, slow lands on index `⌊n / 2⌋`.

This matches the problem definition perfectly.

### Step 5: Remove the Middle Node

Once the middle node is found:

Previous is pointing to the node immediately before it.

Instead of physically shifting nodes, I simply reconnect:

* Previous skips the middle node.
* Previous points directly to the next node after the middle.

This effectively removes the middle node from the linked list.

### Step 6: Return the Updated List

The head remains unchanged except for the single-node edge case.

So I return the original head pointer.

## Examples

### Example 1

**Input**

```text
head = [1,3,4,7,1,2,6]
```

**Output**

```text
[1,3,4,1,2,6]
```

**Explanation**

* Total nodes = 7
* Middle index = 3
* Middle value = 7
* Remove node 7

Result:

```text
1 → 3 → 4 → 1 → 2 → 6
```

### Example 2

**Input**

```text
head = [1,2,3,4]
```

**Output**

```text
[1,2,4]
```

**Explanation**

* Total nodes = 4
* Middle index = 2
* Node value = 3

After deletion:

```text
1 → 2 → 4
```

### Example 3

**Input**

```text
head = [2,1]
```

**Output**

```text
[2]
```

**Explanation**

* Total nodes = 2
* Middle index = 1
* Remove value 1

Remaining list:

```text
2
```

## How to Use / Run Locally

### C++

Compile:

```bash
g++ main.cpp -o main
```

Run:

```bash
./main
```

### Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

### JavaScript

Run:

```bash
node solution.js
```

### Python3

Run:

```bash
python solution.py
```

### Go

Run:

```bash
go run main.go
```

## Notes & Optimizations

* The slow and fast pointer technique is the most efficient solution for this problem.
* No extra array or auxiliary linked list is required.
* The solution uses constant memory.
* It works for both odd-length and even-length linked lists.
* A two-pass solution is possible but less elegant.
* The single-node case must be handled separately.
* This pattern frequently appears in linked list interview questions involving middle node detection.

Related keywords:

* Delete Middle Node of a Linked List
* Linked List Solution
* Two Pointers Algorithm
* Fast and Slow Pointer Technique
* LeetCode 2095 Solution
* DSA Linked List Problems
* Competitive Programming
* Interview Preparation
* Optimal Linked List Deletion

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
