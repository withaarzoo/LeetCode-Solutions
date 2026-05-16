# 154. Find Minimum in Rotated Sorted Array II

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

---

## Problem Summary

LeetCode 154 - Find Minimum in Rotated Sorted Array II is a classic Binary Search problem where we need to find the minimum element inside a rotated sorted array that may contain duplicate values.

The array was originally sorted in ascending order, but it got rotated multiple times. Because of this rotation, the smallest value is no longer guaranteed to be at index `0`.

The challenge becomes harder because duplicates are allowed.

We need to return the minimum element while using an efficient approach with optimized time complexity.

This problem is commonly asked in coding interviews because it tests:

* Binary Search understanding
* Edge case handling
* Rotated array logic
* Duplicate element handling
* Search space optimization

---

## Constraints

| Constraint        | Value                       |
| ----------------- | --------------------------- |
| Array Length      | `1 <= n <= 5000`            |
| Element Range     | `-5000 <= nums[i] <= 5000`  |
| Sorting Condition | Array is sorted and rotated |
| Duplicates        | Allowed                     |

---

## Intuition

My first instinct was to use Binary Search because the array is still partially sorted even after rotation.

In the normal rotated sorted array problem, comparing the middle element with the rightmost element is enough to decide which side contains the minimum value.

But here duplicates change things.

For example:

```text
[2,2,2,0,1]
```

works nicely with Binary Search.

But cases like:

```text
[1,1,1,1,1]
```

or

```text
[2,2,2,0,2]
```

make it difficult to decide the correct side immediately.

I realized that when both `nums[mid]` and `nums[right]` are equal, I cannot safely discard half the array.

So instead of breaking Binary Search completely, I simply shrink the search space by one element.

That small adjustment solves the duplicate problem cleanly.

---

## Approach

I use two pointers:

* `left`
* `right`

Both pointers represent the current search range.

Then I repeatedly calculate the middle index.

### Step 1

Find the middle element.

### Step 2

Compare `nums[mid]` with `nums[right]`.

### Step 3

Handle three possible cases:

#### Case 1: `nums[mid] < nums[right]`

The right side is sorted properly.

This means the minimum lies on the left side including `mid`.

So I move:

```text
right = mid
```

---

#### Case 2: `nums[mid] > nums[right]`

The rotation point exists on the right side.

So the minimum must be after `mid`.

I move:

```text
left = mid + 1
```

---

#### Case 3: `nums[mid] == nums[right]`

Duplicates create ambiguity.

I cannot determine which side contains the minimum.

So I safely shrink the range:

```text
right--
```

---

The loop continues until both pointers meet.

At the end, the pointer directly points to the minimum value.

---

## Data Structures Used

| Data Structure    | Purpose                                          |
| ----------------- | ------------------------------------------------ |
| Array             | Stores the rotated sorted numbers                |
| Integer Variables | Used for Binary Search pointers and middle index |

No extra data structures are needed because the algorithm works directly on the input array.

---

## Operations & Behavior Summary

1. Start Binary Search with `left = 0` and `right = n - 1`
2. Calculate the middle index
3. Compare middle value with right value
4. Decide which side may contain the minimum
5. Remove unnecessary search space
6. Handle duplicates carefully
7. Continue until only one element remains
8. Return the minimum value

---

## Complexity

| Type             | Complexity                            | Explanation                                                                             |
| ---------------- | ------------------------------------- | --------------------------------------------------------------------------------------- |
| Time Complexity  | `O(log n)` average, `O(n)` worst case | Binary Search usually removes half the array, but duplicates may force linear shrinking |
| Space Complexity | `O(1)`                                | No extra memory or additional data structures are used                                  |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int findMin(vector<int>& nums) {
        
        // Initialize binary search boundaries
        int left = 0;
        int right = nums.size() - 1;

        // Continue until both pointers meet
        while (left < right) {

            // Find middle index safely
            int mid = left + (right - left) / 2;

            // Case 1:
            // Minimum lies on left side including mid
            if (nums[mid] < nums[right]) {
                right = mid;
            }

            // Case 2:
            // Minimum lies strictly on right side
            else if (nums[mid] > nums[right]) {
                left = mid + 1;
            }

            // Case 3:
            // Duplicates found, cannot decide direction
            // Safely shrink search space
            else {
                right--;
            }
        }

        // Left now points to minimum element
        return nums[left];
    }
};
```

### Java

```java
class Solution {
    public int findMin(int[] nums) {
        
        // Initialize search boundaries
        int left = 0;
        int right = nums.length - 1;

        // Binary search loop
        while (left < right) {

            // Calculate middle index
            int mid = left + (right - left) / 2;

            // Minimum is on left side including mid
            if (nums[mid] < nums[right]) {
                right = mid;
            }

            // Minimum is on right side
            else if (nums[mid] > nums[right]) {
                left = mid + 1;
            }

            // Duplicate case
            // Remove one duplicate safely
            else {
                right--;
            }
        }

        // Return minimum value
        return nums[left];
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var findMin = function(nums) {
    
    // Initialize pointers
    let left = 0;
    let right = nums.length - 1;

    // Binary search
    while (left < right) {

        // Middle index
        let mid = Math.floor((left + right) / 2);

        // Minimum is on left side including mid
        if (nums[mid] < nums[right]) {
            right = mid;
        }

        // Minimum is on right side
        else if (nums[mid] > nums[right]) {
            left = mid + 1;
        }

        // Duplicate case
        else {
            right--;
        }
    }

    // Return minimum element
    return nums[left];
};
```

### Python3

```python
class Solution:
    def findMin(self, nums: List[int]) -> int:
        
        # Initialize pointers
        left = 0
        right = len(nums) - 1

        # Binary search loop
        while left < right:

            # Middle index
            mid = left + (right - left) // 2

            # Minimum lies on left side including mid
            if nums[mid] < nums[right]:
                right = mid

            # Minimum lies on right side
            elif nums[mid] > nums[right]:
                left = mid + 1

            # Duplicate case
            # Remove one element safely
            else:
                right -= 1

        # Left points to minimum element
        return nums[left]
```

### Go

```go
func findMin(nums []int) int {
    
    // Initialize search boundaries
    left := 0
    right := len(nums) - 1

    // Binary search loop
    for left < right {

        // Calculate middle index
        mid := left + (right-left)/2

        // Minimum is on left side including mid
        if nums[mid] < nums[right] {
            right = mid

        // Minimum is on right side
        } else if nums[mid] > nums[right] {
            left = mid + 1

        // Duplicate case
        // Shrink search space safely
        } else {
            right--
        }
    }

    // Return minimum element
    return nums[left]
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same across all five languages.

Only syntax changes.

### Initial Pointer Setup

I begin with two pointers:

```text
left = 0
right = n - 1
```

These pointers represent the active search area.

The minimum element must exist somewhere between them.

---

### Binary Search Loop

I continue searching while:

```text
left < right
```

As long as both pointers are different, more than one possible answer still exists.

---

### Finding the Middle Index

I calculate:

```text
mid = left + (right - left) / 2
```

This avoids overflow issues that can happen in some languages when using:

```text
(left + right) / 2
```

---

### Understanding the Comparison

The entire algorithm depends on comparing:

```text
nums[mid]
```

with

```text
nums[right]
```

That comparison tells me which side is sorted and where the minimum may exist.

---

### When `nums[mid] < nums[right]`

This means the right side is already sorted.

So the minimum cannot exist after `mid`.

The minimum is either:

* exactly at `mid`
* somewhere before `mid`

So I move:

```text
right = mid
```

I do not remove `mid` because it might actually be the minimum element.

---

### When `nums[mid] > nums[right]`

This means rotation happened somewhere after `mid`.

The minimum must exist on the right side.

So I safely discard the left portion including `mid`.

```text
left = mid + 1
```

---

### When `nums[mid] == nums[right]`

This is the tricky duplicate case.

I cannot decide which side contains the minimum because both values are identical.

Instead of risking the wrong half removal, I shrink the search space carefully:

```text
right--
```

This removes only one duplicate value and keeps the answer safe.

---

### Ending Condition

Eventually:

```text
left == right
```

At this point only one valid candidate remains.

That element is the minimum value in the rotated sorted array.

---

## Examples

### Example 1

### Input

```text
nums = [1,3,5]
```

### Output

```text
1
```

### Explanation

The array is already sorted.

Binary Search quickly narrows down to index `0`.

---

### Example 2

### Input

```text
nums = [2,2,2,0,1]
```

### Output

```text
0
```

### Explanation

* Middle element is compared with right element
* Search moves toward the rotated portion
* Eventually the algorithm reaches `0`

---

### Example 3

### Input

```text
nums = [10,1,10,10,10]
```

### Output

```text
1
```

### Explanation

Duplicates make the search ambiguous.

The algorithm shrinks the search space carefully until the minimum value is found.

---

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

---

### Java

Compile:

```bash
javac Main.java
```

Run:

```bash
java Main
```

---

### JavaScript

Run using Node.js:

```bash
node main.js
```

---

### Python3

Run:

```bash
python main.py
```

or

```bash
python3 main.py
```

---

### Go

Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* This problem is an extension of the classic rotated sorted array minimum problem.
* Duplicates are the main reason worst-case complexity becomes `O(n)`.
* Without duplicates, Binary Search always works in `O(log n)`.
* The algorithm does not modify the array.
* No extra memory is required.
* Edge cases like single-element arrays and fully sorted arrays work naturally.
* Arrays containing all duplicate values are also handled correctly.

Alternative approaches like linear scanning are easier to write but much slower for large inputs.

Binary Search remains the best optimized solution for this problem.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
