# 1848. Minimum Distance to the Target Element

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
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

Given an integer array `nums`, an integer `target`, and an integer `start`, I need to find an index `i` such that:

* `nums[i] == target`
* `abs(i - start)` is minimum

Finally, I return the minimum distance.

The problem guarantees that the target always exists in the array.

## Constraints

* `1 <= nums.length <= 1000`
* `1 <= nums[i] <= 10^4`
* `0 <= start < nums.length`
* `target` exists in `nums`

## Intuition

I thought about checking every element one by one.

Whenever I find an element equal to `target`, I can calculate its distance from `start` using:

```text
abs(i - start)
```

Then I compare it with the minimum distance found so far.

Since the array size is small and I only need the smallest distance, a simple linear scan is enough.

## Approach

1. Create a variable `answer` and initialize it with a very large value.
2. Traverse the array from left to right.
3. Whenever `nums[i] == target`:

   * Calculate `abs(i - start)`
   * Update `answer` with the smaller value
4. Return `answer`

## Data Structures Used

* Array / List
* Integer variable for storing the minimum distance

## Operations & Behavior Summary

* Traverse all elements once
* Compare each value with `target`
* Compute the distance from `start`
* Keep updating the smallest distance
* Return the final minimum distance

## Complexity

* Time Complexity: `O(n)`

  * `n` is the size of the array
  * I only scan the array once

* Space Complexity: `O(1)`

  * I only use a few extra variables

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int getMinDistance(vector<int>& nums, int target, int start) {
        // Store the minimum distance found so far
        int answer = INT_MAX;

        // Traverse through the array
        for (int i = 0; i < nums.size(); i++) {
            // Check if current element is equal to target
            if (nums[i] == target) {
                // Update the minimum distance
                answer = min(answer, abs(i - start));
            }
        }

        return answer;
    }
};
```

### Java

```java
class Solution {
    public int getMinDistance(int[] nums, int target, int start) {
        // Store the minimum distance found so far
        int answer = Integer.MAX_VALUE;

        // Traverse through the array
        for (int i = 0; i < nums.length; i++) {
            // Check if current element is equal to target
            if (nums[i] == target) {
                // Update the minimum distance
                answer = Math.min(answer, Math.abs(i - start));
            }
        }

        return answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} target
 * @param {number} start
 * @return {number}
 */
var getMinDistance = function(nums, target, start) {
    // Store the minimum distance found so far
    let answer = Infinity;

    // Traverse through the array
    for (let i = 0; i < nums.length; i++) {
        // Check if current element is equal to target
        if (nums[i] === target) {
            // Update the minimum distance
            answer = Math.min(answer, Math.abs(i - start));
        }
    }

    return answer;
};
```

### Python3

```python
class Solution:
    def getMinDistance(self, nums: List[int], target: int, start: int) -> int:
        # Store the minimum distance found so far
        answer = float('inf')

        # Traverse through the array
        for i in range(len(nums)):
            # Check if current element is equal to target
            if nums[i] == target:
                # Update the minimum distance
                answer = min(answer, abs(i - start))

        return answer
```

### Go

```go
func getMinDistance(nums []int, target int, start int) int {
    // Store the minimum distance found so far
    answer := int(^uint(0) >> 1)

    // Traverse through the array
    for i := 0; i < len(nums); i++ {
        // Check if current element is equal to target
        if nums[i] == target {
            distance := i - start

            // Convert negative value to positive
            if distance < 0 {
                distance = -distance
            }

            // Update minimum distance
            if distance < answer {
                answer = distance
            }
        }
    }

    return answer
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### Step 1: Initialize the Answer Variable

```cpp
int answer = INT_MAX;
```

I initialize the answer with a very large value.

This helps because later I only want to keep the smallest distance.

---

### Step 2: Traverse the Array

```cpp
for (int i = 0; i < nums.size(); i++)
```

I go through every element of the array.

The variable `i` represents the current index.

---

### Step 3: Check Whether Current Value is Target

```cpp
if (nums[i] == target)
```

Whenever I find the target value, I calculate its distance from `start`.

---

### Step 4: Calculate Distance

```cpp
abs(i - start)
```

This gives the absolute difference between the current index and the start index.

For example:

```text
i = 4, start = 3
abs(4 - 3) = 1
```

---

### Step 5: Update Minimum Distance

```cpp
answer = min(answer, abs(i - start));
```

I compare the current distance with the minimum distance stored so far.

If the current one is smaller, I update `answer`.

---

### Step 6: Return Final Answer

```cpp
return answer;
```

At the end of the loop, `answer` stores the smallest possible distance.

## Examples

### Example 1

```text
Input: nums = [1,2,3,4,5], target = 5, start = 3
Output: 1
```

Explanation:

* Target `5` is found at index `4`
* Distance = `abs(4 - 3) = 1`

### Example 2

```text
Input: nums = [1], target = 1, start = 0
Output: 0
```

Explanation:

* Target `1` is found at index `0`
* Distance = `abs(0 - 0) = 0`

### Example 3

```text
Input: nums = [1,1,1,1,1], target = 1, start = 2
Output: 0
```

Explanation:

* Target exists at index `2`
* Distance = `abs(2 - 2) = 0`

## How to use / Run locally

### C++

```bash
g++ solution.cpp -o solution
./solution
```

### Java

```bash
javac Solution.java
java Solution
```

### JavaScript

```bash
node solution.js
```

### Python3

```bash
python solution.py
```

### Go

```bash
go run solution.go
```

## Notes & Optimizations

* A brute force scan is already efficient enough because the array size is small.
* No sorting is needed.
* No extra array or hash map is needed.
* A single loop is enough to solve the problem.
* This is the most optimal approach for this problem.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
