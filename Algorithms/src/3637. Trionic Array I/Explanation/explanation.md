# Trionic Array I (LeetCode 3637)

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

I am given an integer array `nums` of length `n`.

The array is called **trionic** if I can choose two indices `p` and `q` such that:

* `0 < p < q < n - 1`
* `nums[0 ... p]` is **strictly increasing**
* `nums[p ... q]` is **strictly decreasing**
* `nums[q ... n - 1]` is **strictly increasing**

I need to return `true` if the array is trionic.
Otherwise, return `false`.

---

## Constraints

* `3 <= n <= 100`
* `-1000 <= nums[i] <= 1000`
* All comparisons must be **strict** (no equal values allowed)

---

## Intuition

When I first read the problem, I noticed one important thing.

The array must follow **one fixed pattern**:

Increasing → Decreasing → Increasing

There are no shortcuts.
No reordering.
No skipping elements.

So instead of trying all possible `p` and `q`, I decided to **walk through the array once** and check if it naturally follows this pattern.

If the order breaks at any point, I immediately know the answer is false.

---

## Approach

This is how I solved it step by step:

1. Start from index `0`
2. Move forward while the array is **strictly increasing**
3. Then move forward while the array is **strictly decreasing**
4. Then move forward while the array is **strictly increasing again**
5. At the end, I must reach the last index
6. Each phase must contain **at least one valid step**

If all conditions are satisfied, I return `true`.
Otherwise, I return `false`.

---

## Data Structures Used

* No extra data structures
* Only index pointers and comparisons

---

## Operations & Behavior Summary

* Linear scan of the array
* Direction changes exactly **two times**
* Immediate failure if:

  * Any segment is missing
  * Any comparison is not strict
  * Final index is not reached

---

## Complexity

**Time Complexity:**
`O(n)`
I traverse the array once, where `n` is the array length.

**Space Complexity:**
`O(1)`
No extra memory is used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool isTrionic(vector<int>& nums) {
        int n = nums.size();
        int i = 0;

        while (i + 1 < n && nums[i] < nums[i + 1]) i++;
        if (i == 0 || i == n - 1) return false;

        int mid = i;
        while (i + 1 < n && nums[i] > nums[i + 1]) i++;
        if (i == mid || i == n - 1) return false;

        while (i + 1 < n && nums[i] < nums[i + 1]) i++;

        return i == n - 1;
    }
};
```

---

### Java

```java
class Solution {
    public boolean isTrionic(int[] nums) {
        int n = nums.length;
        int i = 0;

        while (i + 1 < n && nums[i] < nums[i + 1]) i++;
        if (i == 0 || i == n - 1) return false;

        int mid = i;
        while (i + 1 < n && nums[i] > nums[i + 1]) i++;
        if (i == mid || i == n - 1) return false;

        while (i + 1 < n && nums[i] < nums[i + 1]) i++;

        return i == n - 1;
    }
}
```

---

### JavaScript

```javascript
var isTrionic = function(nums) {
    const n = nums.length;
    let i = 0;

    while (i + 1 < n && nums[i] < nums[i + 1]) i++;
    if (i === 0 || i === n - 1) return false;

    let mid = i;
    while (i + 1 < n && nums[i] > nums[i + 1]) i++;
    if (i === mid || i === n - 1) return false;

    while (i + 1 < n && nums[i] < nums[i + 1]) i++;

    return i === n - 1;
};
```

---

### Python3

```python
class Solution:
    def isTrionic(self, nums):
        n = len(nums)
        i = 0

        while i + 1 < n and nums[i] < nums[i + 1]:
            i += 1
        if i == 0 or i == n - 1:
            return False

        mid = i
        while i + 1 < n and nums[i] > nums[i + 1]:
            i += 1
        if i == mid or i == n - 1:
            return False

        while i + 1 < n and nums[i] < nums[i + 1]:
            i += 1

        return i == n - 1
```

---

### Go

```go
func isTrionic(nums []int) bool {
 n := len(nums)
 i := 0

 for i+1 < n && nums[i] < nums[i+1] {
  i++
 }
 if i == 0 || i == n-1 {
  return false
 }

 mid := i
 for i+1 < n && nums[i] > nums[i+1] {
  i++
 }
 if i == mid || i == n-1 {
  return false
 }

 for i+1 < n && nums[i] < nums[i+1] {
  i++
 }

 return i == n-1
}
```

---

## Step-by-step Detailed Explanation

* I use a pointer `i` to move through the array
* First loop checks the increasing phase
* Second loop checks the decreasing phase
* Third loop checks the final increasing phase
* If any phase is missing, I return false
* If I reach the last index successfully, I return true

This ensures correctness with minimum code and maximum clarity.

---

## Examples

**Input:**
`[1, 3, 5, 4, 2, 6]`

**Output:**
`true`

**Explanation:**
Increasing → Decreasing → Increasing pattern exists.

---

**Input:**
`[2, 1, 3]`

**Output:**
`false`

**Explanation:**
The required three segments cannot be formed.

---

## How to use / Run locally

1. Clone the repository
2. Open the file in your preferred language
3. Run using your compiler or interpreter
4. Modify input inside the main function if needed

---

## Notes & Optimizations

* This solution avoids brute force
* Works efficiently within constraints
* Very easy to debug and explain in interviews
* Pattern-based problems like this always benefit from pointer traversal

---

## Author

* **Md Aarzoo Islam**
  [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
