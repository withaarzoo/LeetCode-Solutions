# 3217. Delete Nodes From Linked List Present in Array

## Table of Contents

* ## Problem Summary

* ## Constraints

* ## Intuition

* ## Approach

* ## Data Structures Used

* ## Operations & Behavior Summary

* ## Complexity

* ## Multi-language Solutions

  * ### C++

  * ### Java

  * ### JavaScript

  * ### Python3

  * ### Go

* ## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

* ## Examples

* ## How to use / Run locally

* ## Notes & Optimizations

* ## Author

---

## Problem Summary

You are given an array of integers `nums` and the head of a singly linked list `head`. Remove all nodes from the linked list whose value appears in `nums`, and return the head of the modified linked list.

In short: delete every node from the linked list that has a value present in `nums`.

---

## Constraints

* `1 <= nums.length <= 10^5`
* `1 <= nums[i] <= 10^5`
* All elements in `nums` are unique.
* The number of nodes in the linked list is in the range `[1, 10^5]`.
* `1 <= Node.val <= 10^5`
* The input guarantees there is at least one node in the linked list whose value is **not** present in `nums`.

---

## Intuition

I thought about how to check quickly whether a node's value needs to be deleted. Checking membership in `nums` repeatedly would be slow if done with linear search. So I convert `nums` into a hash set for constant-time membership checks.
Then I saw that deleting nodes in a singly linked list is easiest when I keep a `prev` pointer that can re-link around `curr` if `curr` must be removed. To handle the case where the head needs deletion, I create a `dummy` node pointing to head — this avoids special-case code for head deletions.

---

## Approach

1. Convert `nums` into a hash set `toDelete` for O(1) lookups.
2. Create a `dummy` node with `dummy.next = head`. Initialize two pointers: `prev = dummy`, `curr = head`.
3. Iterate through the list with `curr`:

   * If `curr.val` is in `toDelete`, remove `curr` by `prev.next = curr.next` (do not advance `prev`).
   * Otherwise, keep `curr` and set `prev = curr`.
   * Move `curr = curr.next` after each step.
4. Return `dummy.next` which is the (possibly new) head.

---

## Data Structures Used

* **Hash set** (`unordered_set` / `HashSet` / `Set` / `map`-based set in Go): for O(1) membership checking of values that should be removed.
* **Singly linked list pointers** (`dummy`, `prev`, `curr`): to traverse and perform deletions without special-case head handling.

---

## Operations & Behavior Summary

* Convert `nums` -> set: `O(m)` time & `O(m)` space (m = `nums.length`).
* Single pass over linked list: `O(n)` time (n = number of nodes).
* Deletion is performed by relinking `prev.next` to `curr.next` whenever `curr` is to be removed.
* `prev` advances only when current node is kept; this allows skipping multiple consecutive removable nodes.

---

## Complexity

* **Time Complexity:** `O(n + m)`

  * `n` = number of nodes in the linked list.
  * `m` = `nums.length` (to build the set).
* **Space Complexity:** `O(m)` for the hash set. Apart from that we use `O(1)` extra pointers (`dummy`, `prev`, `curr`).

---

## Multi-language Solutions

### C++

```c++
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
#include <unordered_set>
#include <vector>
using namespace std;

class Solution {
public:
    ListNode* modifiedList(vector<int>& nums, ListNode* head) {
        unordered_set<int> toDelete(nums.begin(), nums.end());
        ListNode dummy(0, head);
        ListNode* prev = &dummy;
        ListNode* curr = head;
        while (curr) {
            if (toDelete.find(curr->val) != toDelete.end()) {
                prev->next = curr->next; // remove curr
            } else {
                prev = curr;            // keep curr
            }
            curr = curr->next;
        }
        return dummy.next;
    }
};
```

---

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
import java.util.HashSet;
import java.util.Set;

class Solution {
    public ListNode modifiedList(int[] nums, ListNode head) {
        Set<Integer> toDelete = new HashSet<>();
        for (int x : nums) toDelete.add(x);

        ListNode dummy = new ListNode(0, head);
        ListNode prev = dummy;
        ListNode curr = head;
        while (curr != null) {
            if (toDelete.contains(curr.val)) {
                prev.next = curr.next; // remove curr
            } else {
                prev = curr;           // keep curr
            }
            curr = curr.next;
        }
        return dummy.next;
    }
}
```

---

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
 * @param {number[]} nums
 * @param {ListNode} head
 * @return {ListNode}
 */
var modifiedList = function(nums, head) {
    const toDelete = new Set(nums);
    const dummy = new ListNode(0, head);
    let prev = dummy;
    let curr = head;
    while (curr !== null) {
        if (toDelete.has(curr.val)) {
            prev.next = curr.next; // remove curr
        } else {
            prev = curr;           // keep curr
        }
        curr = curr.next;
    }
    return dummy.next;
};
```

---

### Python3

```python
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def modifiedList(self, nums: List[int], head: Optional[ListNode]) -> Optional[ListNode]:
        to_delete = set(nums)
        dummy = ListNode(0, head)
        prev, curr = dummy, head
        while curr:
            if curr.val in to_delete:
                prev.next = curr.next  # remove curr
            else:
                prev = curr            # keep curr
            curr = curr.next
        return dummy.next
```

---

### Go

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
    toDelete := make(map[int]struct{}, len(nums))
    for _, v := range nums {
        toDelete[v] = struct{}{}
    }

    dummy := &ListNode{Val: 0, Next: head}
    prev := dummy
    curr := head
    for curr != nil {
        if _, ok := toDelete[curr.Val]; ok {
            prev.Next = curr.Next // remove curr
        } else {
            prev = curr           // keep curr
        }
        curr = curr.Next
    }
    return dummy.Next
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the algorithm (these steps map to every language's implementation):

1. **Build a set of values to delete**

   * C++: `unordered_set<int> toDelete(nums.begin(), nums.end());`
   * Java: `Set<Integer> toDelete = new HashSet<>(); for (int x : nums) toDelete.add(x);`
   * JavaScript: `const toDelete = new Set(nums);`
   * Python: `to_delete = set(nums)`
   * Go: `toDelete := make(map[int]struct{}, len(nums)); for _, v := range nums { toDelete[v] = struct{}{} }`
     *Why:* membership tests become O(1).

2. **Create a dummy node**

   * Purpose: makes deletion of the head node simple.
   * All languages: `dummy` node points to `head`. `prev = dummy`, `curr = head`.

3. **Traverse the linked list with two pointers**

   * Loop: `while curr != null` / `while (curr)` etc.
   * If `curr.val` is in the set:

     * Do `prev.next = curr.next` (skip `curr`, effectively deleting it).
     * Do **not** advance `prev`. Keeping `prev` unmoved lets us delete multiple consecutive nodes because `prev` still refers to the last kept node.
   * Else (value is not in set):

     * Move `prev = curr` (we keep the node).
   * Always advance `curr = curr.next` to continue scanning.

4. **Return `dummy.next`**

   * This is the head of the modified list (handles the original head being deleted or not).

### Why `prev` doesn't move on deletion?

If we delete the `curr`, `prev` must still point to the last kept node so that subsequent deletions can be linked around. Consider a stretch of deletable nodes: we need `prev` to link across the whole deletable block to the next kept node.

---

## Examples

**Example 1**

* Input: `nums = [1,2,3]`, `head = 1 -> 2 -> 3 -> 4 -> 5`
* Output: `4 -> 5`
* Explanation: Remove nodes with values 1, 2 and 3.

**Example 2**

* Input: `nums = [1]`, `head = 1 -> 2 -> 1 -> 2 -> 1 -> 2`
* Output: `2 -> 2 -> 2`
* Explanation: Remove nodes with value 1.

**Example 3**

* Input: `nums = [5]`, `head = 1 -> 2 -> 3 -> 4`
* Output: `1 -> 2 -> 3 -> 4`
* Explanation: No node has value 5.

---

## How to use / Run locally

### C++

1. Create a `main.cpp` where you build a sample list and call `Solution().modifiedList(nums, head)`.
2. Compile and run:

   ```bash
   g++ -std=c++17 main.cpp -o run && ./run
   ```

### Java

1. Place `Solution` and `ListNode` definitions in a `.java` file.
2. Compile and run:

   ```bash
   javac Solution.java
   java Solution
   ```

### JavaScript (Node)

1. Save the JS code in a file `solution.js` and add a simple test harness that builds `ListNode`s and prints result.
2. Run:

   ```bash
   node solution.js
   ```

### Python3

1. Create a Python script `solution.py` with `ListNode` and `Solution`.
2. Run:

   ```bash
   python3 solution.py
   ```

### Go

1. Add the `modifiedList` function into a `main.go` with a test harness.
2. Run:

   ```bash
   go run main.go
   ```

*(The exact test harness is left to you — typical steps: build linked list nodes, call function, traverse and print the result.)*

---

## Notes & Optimizations

* Using a hash set is optimal for membership checks, giving O(1) per check. Because `nums` values are unique, building the set is straightforward.
* The algorithm performs one pass over the linked list, which is optimal — every node must be inspected at least once to decide whether to remove it.
* Memory use is dominated by the hash set `O(m)`. If `nums` is very large and memory is constrained, alternative approaches (like sorting `nums` and using binary search) trade memory for `O(m log m)` preprocessing and are usually slower. Given constraints, the hash-set approach is most practical.
* The dummy-node trick simplifies code and avoids head-special-case logic; this is a standard idiom for linked-list manipulation.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
