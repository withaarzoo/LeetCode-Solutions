# Search in Rotated Sorted Array – Binary Search Solution

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

# Problem Summary

This LeetCode problem asks us to search for a target value inside a rotated sorted array.

The array was originally sorted in ascending order, but it may have been rotated at some pivot index.

For example:

```text
[0,1,2,4,5,6,7]
```

might become:

```text
[4,5,6,7,0,1,2]
```

We need to return:

* the index of the target if it exists
* `-1` if the target is not present

The important part is the required time complexity:

```text
O(log n)
```

That means a linear scan is not good enough. This problem is designed for a modified binary search solution.

This is one of the most common Binary Search interview problems on LeetCode and is frequently asked in coding interviews.

---

# Constraints

| Constraint    | Value                          |
| ------------- | ------------------------------ |
| Array size    | `1 <= nums.length <= 5000`     |
| Element range | `-10^4 <= nums[i] <= 10^4`     |
| Target range  | `-10^4 <= target <= 10^4`      |
| Array values  | Unique                         |
| Array order   | Ascending but possibly rotated |

---

# Intuition

The first thing I noticed was that the array is not fully sorted anymore because of the rotation.

A normal binary search only works on a completely sorted array, so I needed a way to still use binary search even after rotation.

Then I realized something important:

At least one half of the array is always sorted.

For every middle element:

* either the left side is sorted
* or the right side is sorted

Once I know which half is sorted, I can check whether the target belongs inside that range or not.

That single observation turns this into a clean binary search problem.

---

# Approach

I used a modified binary search approach.

Step-by-step process:

1. Start with two pointers:

   * `left`
   * `right`

2. Find the middle index.

3. If the middle element is the target, return its index.

4. Check which half is sorted:

   * If `nums[left] <= nums[mid]`, then the left half is sorted.
   * Otherwise, the right half is sorted.

5. Once the sorted half is identified:

   * Check whether the target lies inside that range.
   * If yes, move toward that side.
   * Otherwise, search the opposite side.

6. Continue until:

   * target is found
   * or search space becomes empty

This keeps reducing the search space by half, which gives logarithmic time complexity.

---

# Data Structures Used

| Data Structure    | Why It Was Used                                   |
| ----------------- | ------------------------------------------------- |
| Array             | The input itself is an array                      |
| Integer variables | Used for pointers like `left`, `right`, and `mid` |

No extra data structure is needed for this solution.

---

# Operations & Behavior Summary

The algorithm works like this:

1. Pick the middle element.
2. Check if it matches the target.
3. Detect which side is sorted.
4. Decide whether the target belongs in the sorted side.
5. Remove the unnecessary half.
6. Repeat until the answer is found.

This is basically a Binary Search on a Rotated Sorted Array.

---

# Complexity

| Type             | Complexity | Explanation                                              |
| ---------------- | ---------- | -------------------------------------------------------- |
| Time Complexity  | `O(log n)` | The search space is divided into half in every iteration |
| Space Complexity | `O(1)`     | No extra data structures are used                        |

Where:

* `n` = size of the array

---

# Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int search(vector<int>& nums, int target) {
        
        // Start pointer
        int left = 0;
        
        // End pointer
        int right = nums.size() - 1;

        // Continue until search space becomes empty
        while (left <= right) {

            // Find middle index safely
            int mid = left + (right - left) / 2;

            // If target is found, return index
            if (nums[mid] == target) {
                return mid;
            }

            // Check if left half is sorted
            if (nums[left] <= nums[mid]) {

                // Check whether target lies inside left sorted half
                if (nums[left] <= target && target < nums[mid]) {
                    
                    // Move to left half
                    right = mid - 1;
                } else {
                    
                    // Move to right half
                    left = mid + 1;
                }
            }
            // Otherwise right half must be sorted
            else {

                // Check whether target lies inside right sorted half
                if (nums[mid] < target && target <= nums[right]) {
                    
                    // Move to right half
                    left = mid + 1;
                } else {
                    
                    // Move to left half
                    right = mid - 1;
                }
            }
        }

        // Target not found
        return -1;
    }
};
```

### Java

```java
class Solution {
    public int search(int[] nums, int target) {
        
        // Start pointer
        int left = 0;
        
        // End pointer
        int right = nums.length - 1;

        // Continue until search space becomes empty
        while (left <= right) {

            // Find middle index safely
            int mid = left + (right - left) / 2;

            // If target is found, return index
            if (nums[mid] == target) {
                return mid;
            }

            // Check if left half is sorted
            if (nums[left] <= nums[mid]) {

                // Check whether target lies inside left sorted half
                if (nums[left] <= target && target < nums[mid]) {
                    
                    // Move search to left side
                    right = mid - 1;
                } else {
                    
                    // Move search to right side
                    left = mid + 1;
                }
            }
            // Otherwise right half is sorted
            else {

                // Check whether target lies inside right sorted half
                if (nums[mid] < target && target <= nums[right]) {
                    
                    // Move search to right side
                    left = mid + 1;
                } else {
                    
                    // Move search to left side
                    right = mid - 1;
                }
            }
        }

        // Target not found
        return -1;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function(nums, target) {
    
    // Start pointer
    let left = 0;

    // End pointer
    let right = nums.length - 1;

    // Continue until search space becomes empty
    while (left <= right) {

        // Find middle index safely
        let mid = Math.floor(left + (right - left) / 2);

        // If target is found
        if (nums[mid] === target) {
            return mid;
        }

        // Check if left half is sorted
        if (nums[left] <= nums[mid]) {

            // Check whether target lies inside left sorted half
            if (nums[left] <= target && target < nums[mid]) {
                
                // Search left side
                right = mid - 1;
            } else {
                
                // Search right side
                left = mid + 1;
            }
        }
        // Otherwise right half is sorted
        else {

            // Check whether target lies inside right sorted half
            if (nums[mid] < target && target <= nums[right]) {
                
                // Search right side
                left = mid + 1;
            } else {
                
                // Search left side
                right = mid - 1;
            }
        }
    }

    // Target not found
    return -1;
};
```

### Python3

```python
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        
        # Start pointer
        left = 0

        # End pointer
        right = len(nums) - 1

        # Continue until search space becomes empty
        while left <= right:

            # Find middle index safely
            mid = left + (right - left) // 2

            # If target is found
            if nums[mid] == target:
                return mid

            # Check if left half is sorted
            if nums[left] <= nums[mid]:

                # Check whether target lies inside left sorted half
                if nums[left] <= target < nums[mid]:

                    # Search left side
                    right = mid - 1
                else:

                    # Search right side
                    left = mid + 1

            # Otherwise right half is sorted
            else:

                # Check whether target lies inside right sorted half
                if nums[mid] < target <= nums[right]:

                    # Search right side
                    left = mid + 1
                else:

                    # Search left side
                    right = mid - 1

        # Target not found
        return -1
```

### Go

```go
func search(nums []int, target int) int {
    
    // Start pointer
    left := 0

    // End pointer
    right := len(nums) - 1

    // Continue until search space becomes empty
    for left <= right {

        // Find middle index safely
        mid := left + (right-left)/2

        // If target is found
        if nums[mid] == target {
            return mid
        }

        // Check if left half is sorted
        if nums[left] <= nums[mid] {

            // Check whether target lies inside left sorted half
            if nums[left] <= target && target < nums[mid] {

                // Search left side
                right = mid - 1
            } else {

                // Search right side
                left = mid + 1
            }

        } else {

            // Right half is sorted

            // Check whether target lies inside right sorted half
            if nums[mid] < target && target <= nums[right] {

                // Search right side
                left = mid + 1
            } else {

                // Search left side
                right = mid - 1
            }
        }
    }

    // Target not found
    return -1
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic stays exactly the same in all five languages.

Only syntax changes.

### 1. Initialize Binary Search Pointers

Every version starts with:

* `left = 0`
* `right = n - 1`

These pointers represent the current search range.

---

### 2. Run Binary Search Loop

The loop continues while:

```text
left <= right
```

The moment `left` crosses `right`, the target does not exist.

---

### 3. Find the Middle Index

We calculate:

```text
mid = left + (right - left) / 2
```

This is safer than:

```text
(left + right) / 2
```

because it avoids integer overflow in some languages.

---

### 4. Check If Middle Element Is the Target

If:

```text
nums[mid] == target
```

we immediately return the index.

This is the fastest possible success case.

---

### 5. Detect the Sorted Half

This is the most important step.

We check:

```text
nums[left] <= nums[mid]
```

If true:

* the left half is sorted

Otherwise:

* the right half is sorted

Even after rotation, one side always remains properly sorted.

---

### 6. Decide Which Side to Search

If the left half is sorted:

* check whether the target lies between `nums[left]` and `nums[mid]`

If yes:

* move left

Otherwise:

* move right

Same idea applies when the right half is sorted.

This removes half the array every iteration.

---

### 7. Return -1 If Not Found

If the loop finishes without finding the target:

```text
return -1
```

That means the target does not exist in the rotated sorted array.

---

# Examples

## Example 1

### Input

```text
nums = [4,5,6,7,0,1,2]
target = 0
```

### Output

```text
4
```

### Explanation

* Middle starts at `7`
* Left side is sorted
* Target is not inside left range
* Search moves right
* Target `0` is found at index `4`

---

## Example 2

### Input

```text
nums = [4,5,6,7,0,1,2]
target = 3
```

### Output

```text
-1
```

### Explanation

The algorithm keeps reducing the search space, but `3` never appears.

---

## Example 3

### Input

```text
nums = [1]
target = 0
```

### Output

```text
-1
```

### Explanation

The array contains only one element and it does not match the target.

---

# How to Use / Run Locally

## C++

Compile:

```bash
g++ main.cpp -o main
```

Run:

```bash
./main
```

---

## Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

---

## JavaScript

Run using Node.js:

```bash
node main.js
```

---

## Python3

Run:

```bash
python main.py
```

---

## Go

Run:

```bash
go run main.go
```

---

# Notes & Optimizations

* A normal binary search cannot directly handle rotated arrays.
* The key observation is identifying the sorted half.
* This problem is a classic example of Modified Binary Search.
* Since all values are unique, the logic becomes simpler.
* If duplicates were allowed, extra handling would be required.
* The solution uses constant extra space.
* This is one of the most important Binary Search interview problems for coding interviews and DSA preparation.

---

# Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
