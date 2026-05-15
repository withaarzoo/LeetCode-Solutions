# 153. Find Minimum in Rotated Sorted Array

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

This LeetCode DSA problem asks us to find the minimum element in a rotated sorted array.

The array was originally sorted in ascending order, but it has been rotated some number of times. Because of this rotation, the smallest value is no longer guaranteed to be at index `0`.

The goal is to return the minimum value in `O(log n)` time complexity, which means a normal linear scan is not the best approach here.

Example:

```text
Input: [4,5,6,7,0,1,2]
Output: 0
```

The important thing to notice is that even after rotation, one side of the array always remains sorted. That observation helps us solve the problem using Binary Search.

This problem is commonly asked in coding interviews because it tests understanding of:

* Binary Search
* Rotated Sorted Array logic
* Search space reduction
* Logarithmic time optimization

## Constraints

| Constraint                  | Value         |
| --------------------------- | ------------- |
| `1 <= nums.length <= 5000`  | Array size    |
| `-5000 <= nums[i] <= 5000`  | Element range |
| All integers are unique     | No duplicates |
| Array is sorted and rotated | Guaranteed    |

## Intuition

When I first looked at the problem, I noticed that the array is not completely random. Even after rotation, parts of it are still sorted.

That immediately made me think about Binary Search.

If I compare the middle element with the rightmost element, I can figure out where the minimum value exists.

* If the middle value is greater than the right value, the minimum must be on the right side.
* Otherwise, the minimum is on the left side or could even be the middle itself.

Instead of checking every element one by one, I keep reducing the search space by half.

That is why Binary Search works perfectly here.

## Approach

I start with two pointers:

* `left`
* `right`

The search space lies between these two pointers.

Then I repeatedly:

1. Find the middle index
2. Compare `nums[mid]` with `nums[right]`
3. Decide which half can be ignored

There are two cases:

### Case 1: `nums[mid] > nums[right]`

This means the rotation point is somewhere after `mid`.

So the minimum value must exist in the right half.

I move:

```text
left = mid + 1
```

### Case 2: `nums[mid] <= nums[right]`

This means the right half is already sorted.

The minimum can either be:

* at `mid`
* or somewhere on the left side

So I move:

```text
right = mid
```

Eventually, both pointers meet at the minimum element.

## Data Structures Used

| Data Structure    | Why It Was Used                    |
| ----------------- | ---------------------------------- |
| Array             | Input storage                      |
| Integer variables | Used for pointers and middle index |

No extra data structure is needed because Binary Search works directly on the given array.

## Operations & Behavior Summary

The algorithm works like this:

1. Start with the full array
2. Find the middle element
3. Compare middle with rightmost value
4. Remove the sorted half that cannot contain the minimum
5. Continue shrinking the search space
6. Stop when both pointers point to the same index
7. Return that value as the minimum element

This process keeps cutting the array size in half, which gives logarithmic time complexity.

## Complexity

| Type             | Complexity | Explanation                                                    |
| ---------------- | ---------- | -------------------------------------------------------------- |
| Time Complexity  | `O(log n)` | Binary Search removes half the search space in every iteration |
| Space Complexity | `O(1)`     | No extra data structure is used                                |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int findMin(vector<int>& nums) {
        
        // Left pointer starts from beginning
        int left = 0;
        
        // Right pointer starts from end
        int right = nums.size() - 1;

        // Continue until both pointers meet
        while (left < right) {
            
            // Find middle index safely
            int mid = left + (right - left) / 2;

            // Minimum lies in right half
            if (nums[mid] > nums[right]) {
                
                // Ignore left sorted part
                left = mid + 1;
            }
            else {
                
                // Minimum may be at mid or left side
                right = mid;
            }
        }

        // Both pointers point to minimum element
        return nums[left];
    }
};
```

### Java

```java
class Solution {
    public int findMin(int[] nums) {
        
        // Left pointer at start
        int left = 0;
        
        // Right pointer at end
        int right = nums.length - 1;

        // Binary Search loop
        while (left < right) {
            
            // Find middle index
            int mid = left + (right - left) / 2;

            // Minimum is in right half
            if (nums[mid] > nums[right]) {
                
                // Move left pointer
                left = mid + 1;
            } 
            else {
                
                // Minimum can be mid itself
                right = mid;
            }
        }

        // Return minimum element
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
    
    // Start pointer
    let left = 0;
    
    // End pointer
    let right = nums.length - 1;

    // Run Binary Search
    while (left < right) {
        
        // Middle index
        let mid = Math.floor((left + right) / 2);

        // Minimum exists on right side
        if (nums[mid] > nums[right]) {
            
            // Remove left sorted half
            left = mid + 1;
        } 
        else {
            
            // Minimum can still be at mid
            right = mid;
        }
    }

    // Final answer
    return nums[left];
};
```

### Python3

```python
class Solution:
    def findMin(self, nums: List[int]) -> int:
        
        # Left pointer
        left = 0
        
        # Right pointer
        right = len(nums) - 1

        # Binary Search loop
        while left < right:
            
            # Middle index
            mid = left + (right - left) // 2

            # Minimum lies on right side
            if nums[mid] > nums[right]:
                
                # Move left pointer
                left = mid + 1
            else:
                
                # Minimum may be at mid
                right = mid

        # Return minimum element
        return nums[left]
```

### Go

```go
func findMin(nums []int) int {
    
    // Left pointer
    left := 0

    // Right pointer
    right := len(nums) - 1

    // Binary Search loop
    for left < right {

        // Middle index
        mid := left + (right-left)/2

        // Minimum is on right side
        if nums[mid] > nums[right] {

            // Move left pointer
            left = mid + 1
        } else {

            // Minimum may be at mid
            right = mid
        }
    }

    // Return minimum element
    return nums[left]
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same in all five languages. Only syntax changes.

First, two pointers are created:

* one at the beginning
* one at the end

These pointers represent the current search space.

Then a loop runs while:

```text
left < right
```

The loop stops only when both pointers meet.

Inside the loop, the middle index is calculated carefully to avoid overflow problems in some languages like Java and C++.

Then the algorithm compares:

```text
nums[mid]
```

with:

```text
nums[right]
```

This comparison is the most important part of the solution.

### If `nums[mid] > nums[right]`

The minimum value must be on the right side because the left side is properly sorted.

So the left pointer moves forward.

### If `nums[mid] <= nums[right]`

The right side is sorted, meaning the minimum is not beyond `right`.

The minimum may still be at `mid`, so we do not remove it.

That is why:

```text
right = mid
```

is used instead of:

```text
right = mid - 1
```

Finally:

```text
left == right
```

At this point, both pointers are pointing to the minimum element.

The algorithm returns that value.

This approach is much faster than checking every element because the search space becomes half after every step.

## Examples

### Example 1

```text
Input: [3,4,5,1,2]
Output: 1
```

Trace:

* mid = 5
* 5 > 2
* move left to right half
* minimum found at index containing `1`

---

### Example 2

```text
Input: [4,5,6,7,0,1,2]
Output: 0
```

Trace:

* mid = 7
* 7 > 2
* move to right half
* continue Binary Search
* final answer becomes `0`

---

### Example 3

```text
Input: [11,13,15,17]
Output: 11
```

Trace:

* array is already sorted
* Binary Search still works
* left pointer finally reaches `11`

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

Run:

```bash
node main.js
```

---

### Python3

Run:

```bash
python main.py
```

---

### Go

Run:

```bash
go run main.go
```

## Notes & Optimizations

* This problem is a classic Binary Search interview problem.
* A linear scan would work, but its time complexity would be `O(n)`.
* Binary Search improves the solution to `O(log n)`.
* The array contains unique values, which makes the comparison logic simpler.
* If duplicates were allowed, the logic would need additional handling.
* The algorithm works even if the array is not rotated at all.
* No extra memory is used, so the solution is memory efficient.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
