# 3689. Maximum Total Subarray Value I

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

LeetCode 3689 - Maximum Total Subarray Value I is an array and greedy problem where we need to select exactly `k` non-empty subarrays and maximize the total value obtained.

The value of a subarray is defined as:

`maximum element - minimum element`

A subarray may be selected multiple times, and overlapping subarrays are allowed.

The goal is to return the maximum possible total value after choosing exactly `k` subarrays.

### Input

* An integer array `nums`
* An integer `k`

### Output

* Maximum possible total subarray value

This problem looks like a subarray optimization problem at first, but there is a very simple observation that reduces it to finding only the global minimum and maximum elements in the array.

## Constraints

| Constraint                     | Value               |
| ------------------------------ | ------------------- |
| `1 <= nums.length <= 5 * 10^4` | Array size          |
| `0 <= nums[i] <= 10^9`         | Element value       |
| `1 <= k <= 10^5`               | Number of subarrays |

## Intuition

When I first read the problem, I focused on the fact that the same subarray can be chosen more than once.

That immediately changes the problem.

Normally, subarray selection problems require finding multiple different ranges. Here, if I find one subarray with the highest possible value, I can simply choose that same subarray `k` times.

So instead of searching for `k` different subarrays, I only need to know:

"What is the maximum possible value of any single subarray?"

For any subarray:

```text
value = maximum element - minimum element
```

The largest possible maximum comes from the largest number in the entire array.

The smallest possible minimum comes from the smallest number in the entire array.

Since the whole array is itself a valid subarray, the maximum subarray value becomes:

```text
global maximum - global minimum
```

After that, the answer is simply:

```text
k × (global maximum - global minimum)
```

## Approach

1. Traverse the array once.
2. Find the smallest element in the entire array.
3. Find the largest element in the entire array.
4. Compute the best possible subarray value.
5. Multiply that value by `k`.
6. Return the result.

The key observation is that selecting the whole array already achieves the largest possible difference between maximum and minimum values.

Because duplicate selections are allowed, we can select that optimal subarray exactly `k` times.

## Data Structures Used

### Variables

* `minElement` → Stores the smallest value found so far.
* `maxElement` → Stores the largest value found so far.

No additional data structures are required.

This keeps the solution extremely memory efficient.

## Operations & Behavior Summary

The algorithm performs the following operations:

1. Start scanning the array.
2. Keep updating the current minimum element.
3. Keep updating the current maximum element.
4. Finish scanning once all elements are processed.
5. Calculate:

```text
bestValue = maxElement - minElement
```

1. Calculate:

```text
answer = bestValue × k
```

1. Return the final answer.

## Complexity

| Metric           | Complexity | Explanation                                                      |
| ---------------- | ---------- | ---------------------------------------------------------------- |
| Time Complexity  | O(n)       | The array is scanned once to find the minimum and maximum values |
| Space Complexity | O(1)       | Only a few variables are used regardless of input size           |

Where:

* `n` = length of the array

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long maxTotalValue(vector<int>& nums, int k) {
        // Find the smallest element in the array
        long long mn = *min_element(nums.begin(), nums.end());

        // Find the largest element in the array
        long long mx = *max_element(nums.begin(), nums.end());

        // Best subarray value = global maximum - global minimum
        long long best = mx - mn;

        // We can choose the same best subarray k times
        return best * k;
    }
};
```

### Java

```java
class Solution {
    public long maxTotalValue(int[] nums, int k) {
        // Initialize minimum and maximum using first element
        long mn = nums[0];
        long mx = nums[0];

        // Find global minimum and maximum
        for (int num : nums) {
            mn = Math.min(mn, num);
            mx = Math.max(mx, num);
        }

        // Best possible subarray value
        long best = mx - mn;

        // Choose that subarray k times
        return best * k;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxTotalValue = function(nums, k) {
    // Track global minimum and maximum
    let mn = Infinity;
    let mx = -Infinity;

    // Find minimum and maximum element
    for (const num of nums) {
        mn = Math.min(mn, num);
        mx = Math.max(mx, num);
    }

    // Best subarray value
    const best = mx - mn;

    // Choose it k times
    return best * k;
};
```

### Python3

```python
class Solution:
    def maxTotalValue(self, nums: List[int], k: int) -> int:
        # Global minimum element
        mn = min(nums)

        # Global maximum element
        mx = max(nums)

        # Best subarray value
        best = mx - mn

        # Choose the same best subarray k times
        return best * k
```

### Go

```go
func maxTotalValue(nums []int, k int) int64 {
    // Initialize minimum and maximum
    mn := nums[0]
    mx := nums[0]

    // Find global minimum and maximum
    for _, num := range nums {
        if num < mn {
            mn = num
        }
        if num > mx {
            mx = num
        }
    }

    // Best subarray value
    best := int64(mx - mn)

    // Choose the same best subarray k times
    return best * int64(k)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical in all five languages.

### Step 1: Initialize Minimum and Maximum

Before scanning the array, we need variables that track:

* Smallest element seen so far
* Largest element seen so far

These values will eventually become the global minimum and maximum.

### Step 2: Traverse the Array

We process every element exactly once.

For each number:

* If it is smaller than the current minimum, update the minimum.
* If it is larger than the current maximum, update the maximum.

After the loop finishes:

* `minElement` contains the smallest number in the array.
* `maxElement` contains the largest number in the array.

### Step 3: Compute the Best Subarray Value

The largest possible subarray value is:

```text
maxElement - minElement
```

Why?

Because no subarray can contain a value larger than the global maximum or smaller than the global minimum.

The whole array already contains both.

Therefore this value is achievable.

### Step 4: Use the Same Subarray Repeatedly

The problem explicitly allows selecting the same subarray multiple times.

If one subarray gives the best value:

```text
bestValue
```

then choosing it `k` times gives:

```text
bestValue × k
```

Since no subarray can contribute more than `bestValue`, this is always optimal.

### Step 5: Return the Result

The final answer is:

```text
(maxElement - minElement) × k
```

This completes the solution.

## Examples

### Example 1

Input:

```text
nums = [1,3,2]
k = 2
```

Process:

```text
minimum = 1
maximum = 3
bestValue = 3 - 1 = 2
answer = 2 × 2 = 4
```

Output:

```text
4
```

---

### Example 2

Input:

```text
nums = [4,2,5,1]
k = 3
```

Process:

```text
minimum = 1
maximum = 5
bestValue = 4
answer = 4 × 3 = 12
```

Output:

```text
12
```

---

### Example 3

Input:

```text
nums = [7,7,7]
k = 5
```

Process:

```text
minimum = 7
maximum = 7
bestValue = 0
answer = 0 × 5 = 0
```

Output:

```text
0
```

## How to Use / Run Locally

### C++

Compile:

```bash
g++ solution.cpp -o solution
```

Run:

```bash
./solution
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

or

```bash
python3 solution.py
```

### Go

Run:

```bash
go run solution.go
```

Build:

```bash
go build solution.go
```

## Notes & Optimizations

* This is one of those problems that appears harder than it actually is.
* The ability to reuse the same subarray completely changes the solution.
* There is no need for prefix sums.
* There is no need for sliding window techniques.
* There is no need for monotonic stacks.
* There is no need for dynamic programming.
* A single linear scan is enough.
* The solution is already optimal because every element must be inspected at least once.
* The algorithm works efficiently even for the largest constraints.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
