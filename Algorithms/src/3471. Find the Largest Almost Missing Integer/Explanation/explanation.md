# 3471. Find the Largest Almost Missing Integer

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

Given an integer array `nums` and an integer `k`, the goal is to find the largest integer that appears in exactly one subarray of size `k`.

A subarray is a continuous part of an array.

For example, if:

```text
nums = [3, 9, 2, 1, 7]
k = 3
```

the subarrays of size `3` are:

```text
[3, 9, 2]
[9, 2, 1]
[2, 1, 7]
```

Here, `3` appears in one subarray and `7` appears in one subarray. The largest valid integer is `7`.

If no integer appears in exactly one subarray of size `k`, the answer is `-1`.

This solution uses array frequency counting and handles the important cases `k = 1`, `k = n`, and `1 < k < n` separately.

## Constraints

* `1 <= nums.length <= 50`
* `0 <= nums[i] <= 50`
* `1 <= k <= nums.length`

## Intuition

My first thought was to generate every subarray of size `k` and count how many times each number appears. That would work, but it does more work than necessary.

The key observation is that the position of an element matters.

When `k = 1`, every subarray contains only one element. So a number is valid only when it occurs exactly once in the entire array.

When `k = n`, there is only one subarray: the entire array. Therefore, every distinct number appears in exactly one subarray, so I only need the largest number.

The interesting case is `1 < k < n`.

In this case, an element in the middle of the array belongs to more than one subarray of size `k`. So only the first and last elements can possibly belong to exactly one subarray.

I then check whether the values at those two positions occur only once in the entire array. If they do, they are valid candidates, and I return the larger one.

Because every number is between `0` and `50`, a small frequency array is enough.

## Approach

I solve the problem in three cases.

1. I count the frequency of every number in `nums`.

2. If `k == 1`, I search for the largest number whose frequency is exactly `1`.

3. If `k == n`, I return the largest number in the entire array.

4. Otherwise, `1 < k < n`. I only check `nums[0]` and `nums[n - 1]`.

5. For each endpoint, I check whether its value occurs exactly once in the whole array.

6. I return the larger valid endpoint. If neither endpoint is valid, I return `-1`.

This avoids explicitly creating or checking every subarray.

## Data Structures Used

### Frequency Array

I use a fixed-size frequency array of size `51`.

Each index represents a possible value from `0` to `50`.

For example:

```text
freq[7]
```

stores how many times `7` appears in `nums`.

I use this instead of a hash map because the problem gives a very small value range. This makes the implementation simple and keeps the extra space constant.

## Operations & Behavior Summary

The algorithm can be viewed as this simple pseudocode:

```text
Count the frequency of every value in nums.

If k == 1:
    Find the largest value whose frequency is exactly 1.
    Return it.
    If none exists, return -1.

If k == n:
    Return the largest value in nums.

Otherwise:
    Check nums[0].
    If its frequency is 1, consider it as a candidate.

    Check nums[n - 1].
    If its frequency is 1, consider it as a candidate.

    Return the larger candidate.
    If neither is valid, return -1.
```

## Complexity

| Complexity       |   Cost | Explanation                                                                          |
| ---------------- | -----: | ------------------------------------------------------------------------------------ |
| Time Complexity  | `O(n)` | I scan `nums` to count frequencies and perform only constant-size additional checks. |
| Space Complexity | `O(1)` | The frequency array always has size `51`, regardless of the input size `n`.          |

Here, `n` is the length of the input array `nums`.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int largestInteger(vector<int>& nums, int k) {
        int n = nums.size();

        // nums[i] is between 0 and 50, so a fixed frequency array is enough.
        // This uses constant extra space instead of a hash table.
        int freq[51] = {};

        // Count how many times every value appears in nums.
        for (int x : nums) {
            freq[x]++;
        }

        // When k = 1, every subarray contains exactly one element.
        // Therefore, a value is almost missing exactly when it appears once.
        if (k == 1) {
            // Search from 50 down to 0 so the first valid value is the largest.
            for (int x = 50; x >= 0; x--) {
                if (freq[x] == 1) {
                    return x;
                }
            }

            // No value occurs exactly once.
            return -1;
        }

        // When k = n, there is only one subarray: the entire array.
        // Every distinct value appears in exactly one subarray.
        if (k == n) {
            int answer = 0;

            // Find the largest value in the whole array.
            for (int x : nums) {
                answer = max(answer, x);
            }

            return answer;
        }

        // For 1 < k < n, only the first and last elements
        // can belong to exactly one subarray of size k.
        int answer = -1;

        // The first element is valid only if its value occurs once in nums.
        if (freq[nums[0]] == 1) {
            answer = max(answer, nums[0]);
        }

        // The last element is valid only if its value occurs once in nums.
        if (freq[nums[n - 1]] == 1) {
            answer = max(answer, nums[n - 1]);
        }

        // Return the largest valid endpoint, or -1 if neither is valid.
        return answer;
    }
};
```

### Java

```java
class Solution {
    public int largestInteger(int[] nums, int k) {
        int n = nums.length;

        // nums[i] is between 0 and 50, so a fixed array can store all frequencies.
        // This keeps the extra space constant.
        int[] freq = new int[51];

        // Count how many times every value occurs in nums.
        for (int x : nums) {
            freq[x]++;
        }

        // When k = 1, each subarray contains one element.
        // So only values occurring exactly once are valid.
        if (k == 1) {
            // Check from the largest possible value to the smallest.
            for (int x = 50; x >= 0; x--) {
                if (freq[x] == 1) {
                    return x;
                }
            }

            // No value occurs exactly once.
            return -1;
        }

        // When k = n, the entire array is the only subarray.
        // Therefore, the largest value in nums is the answer.
        if (k == n) {
            int answer = 0;

            // Find the maximum value in the array.
            for (int x : nums) {
                answer = Math.max(answer, x);
            }

            return answer;
        }

        // For 1 < k < n, only the first and last elements
        // can occur in exactly one subarray of size k.
        int answer = -1;

        // The first value is valid only if it appears once in the whole array.
        if (freq[nums[0]] == 1) {
            answer = Math.max(answer, nums[0]);
        }

        // The last value is valid only if it appears once in the whole array.
        if (freq[nums[n - 1]] == 1) {
            answer = Math.max(answer, nums[n - 1]);
        }

        // Return the largest valid candidate, or -1 if none exists.
        return answer;
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
var largestInteger = function(nums, k) {
    const n = nums.length;

    // nums[i] is between 0 and 50, so a fixed frequency array is enough.
    // Using an array keeps the extra space constant.
    const freq = new Array(51).fill(0);

    // Count the frequency of every value in nums.
    for (const x of nums) {
        freq[x]++;
    }

    // When k = 1, every subarray contains exactly one element.
    // Therefore, a valid value must occur exactly once in nums.
    if (k === 1) {
        // Search from 50 downwards so the first valid value is the largest.
        for (let x = 50; x >= 0; x--) {
            if (freq[x] === 1) {
                return x;
            }
        }

        // No value appears exactly once.
        return -1;
    }

    // When k = n, the entire array is the only subarray.
    // Hence, every distinct value appears in exactly one subarray.
    if (k === n) {
        let answer = 0;

        // Find the largest value in nums.
        for (const x of nums) {
            answer = Math.max(answer, x);
        }

        return answer;
    }

    // For 1 < k < n, only the first and last elements
    // can belong to exactly one subarray of size k.
    let answer = -1;

    // The first value is valid only if it occurs once in nums.
    if (freq[nums[0]] === 1) {
        answer = Math.max(answer, nums[0]);
    }

    // The last value is valid only if it occurs once in nums.
    if (freq[nums[n - 1]] === 1) {
        answer = Math.max(answer, nums[n - 1]);
    }

    // Return the largest valid endpoint, or -1 if neither is valid.
    return answer;
};
```

### Python3

```python
from typing import List

class Solution:
    def largestInteger(self, nums: List[int], k: int) -> int:
        n = len(nums)

        # nums[i] is between 0 and 50, so a fixed frequency array is enough.
        # Its size never depends on n, so the extra space stays constant.
        freq = [0] * 51

        # Count how many times every value occurs in nums.
        for x in nums:
            freq[x] += 1

        # When k = 1, every subarray contains exactly one element.
        # So a value is valid only when it occurs exactly once in nums.
        if k == 1:
            # Check from the largest possible value to the smallest.
            for x in range(50, -1, -1):
                if freq[x] == 1:
                    return x

            # No value occurs exactly once.
            return -1

        # When k = n, the whole array is the only subarray.
        # Therefore, the largest value in nums is the answer.
        if k == n:
            # max(nums) gives the largest value in the only subarray.
            return max(nums)

        # For 1 < k < n, only the first and last elements
        # can appear in exactly one subarray of size k.
        answer = -1

        # The first value is valid only if it occurs once in the whole array.
        if freq[nums[0]] == 1:
            answer = max(answer, nums[0])

        # The last value is valid only if it occurs once in the whole array.
        if freq[nums[-1]] == 1:
            answer = max(answer, nums[-1])

        # Return the largest valid candidate, or -1 if none exists.
        return answer
```

### Go

```go
func largestInteger(nums []int, k int) int {
 n := len(nums)

 // nums[i] is between 0 and 50, so a fixed-size frequency array is enough.
 // Its size is constant, so it does not grow with n.
 freq := [51]int{}

 // Count how many times every value occurs in nums.
 for _, x := range nums {
  freq[x]++
 }

 // When k = 1, every subarray contains exactly one element.
 // Therefore, only values occurring exactly once are valid.
 if k == 1 {
  // Check from 50 down to 0 to find the largest valid value first.
  for x := 50; x >= 0; x-- {
   if freq[x] == 1 {
    return x
   }
  }

  // No value occurs exactly once.
  return -1
 }

 // When k = n, the whole array is the only subarray.
 // Every distinct value therefore appears in exactly one subarray.
 if k == n {
  answer := 0

  // Find the largest value in nums.
  for _, x := range nums {
   if x > answer {
    answer = x
   }
  }

  return answer
 }

 // For 1 < k < n, only the first and last elements
 // can belong to exactly one subarray of size k.
 answer := -1

 // The first value is valid only if it occurs once in nums.
 if freq[nums[0]] == 1 && nums[0] > answer {
  answer = nums[0]
 }

 // The last value is valid only if it occurs once in nums.
 if freq[nums[n-1]] == 1 && nums[n-1] > answer {
  answer = nums[n-1]
 }

 // Return the largest valid endpoint, or -1 if none exists.
 return answer
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five languages. Only the syntax for arrays, loops, and basic operations changes.

### Step 1: Get the array length

I first store the length of `nums` in `n`.

I need this value to identify whether `k` is `1`, equal to the entire array length, or somewhere between them.

### Step 2: Count frequencies

I create a frequency array with `51` positions.

Since the problem guarantees:

```text
0 <= nums[i] <= 50
```

every value can directly be used as an index.

For every number in `nums`, I increase its frequency.

For example:

```text
nums = [3, 9, 2, 1, 7]
```

gives:

```text
3 -> 1
9 -> 1
2 -> 1
1 -> 1
7 -> 1
```

If a number appears multiple times, its frequency becomes greater than `1`.

### Step 3: Handle `k == 1`

When `k` is `1`, the possible subarrays contain one element each.

For:

```text
nums = [3, 7, 2, 7, 1]
```

the subarrays are:

```text
[3]
[7]
[2]
[7]
[1]
```

The value `7` appears in two subarrays, so it is not almost missing.

The values `3`, `2`, and `1` each appear in exactly one subarray.

I search from the largest possible value down to `0`. The first value whose frequency is `1` is the answer.

### Step 4: Handle `k == n`

When `k` equals `n`, there is only one subarray.

For:

```text
nums = [3, 9, 2, 1, 7]
k = 5
```

the only subarray is:

```text
[3, 9, 2, 1, 7]
```

Every distinct value is therefore present in exactly one subarray.

So I simply find the maximum value in `nums`.

### Step 5: Handle `1 < k < n`

This is the main observation.

An element in the middle of the array can belong to multiple subarrays of size `k`.

For example:

```text
nums = [3, 9, 2, 1, 7]
k = 3
```

The subarrays are:

```text
[3, 9, 2]
[9, 2, 1]
[2, 1, 7]
```

The first element `3` appears only in the first subarray.

The last element `7` appears only in the last subarray.

But `2` appears in all three subarrays.

Therefore, only the first and last positions can give us a valid candidate.

### Step 6: Check the first element

I check:

```text
frequency[nums[0]] == 1
```

The first element belongs to only one subarray, but its value must also appear nowhere else in the array.

For example:

```text
nums = [7, 3, 2, 7]
k = 2
```

Although the first `7` is in one subarray, the value `7` also occurs at the end.

So `7` appears in two subarrays and is not valid.

That is why checking its frequency is necessary.

### Step 7: Check the last element

I perform the same check for:

```text
nums[n - 1]
```

If its frequency is exactly `1`, it is a valid candidate.

### Step 8: Choose the largest candidate

If both endpoints are valid, I compare them and keep the larger value.

If only one is valid, that value becomes the answer.

If neither is valid, the answer remains `-1`.

### Language-specific notes

In C++, I use a fixed integer array for frequencies and range-based loops for simple traversal.

In Java, I use an `int[]` of size `51`. Java arrays are initialized with zero values, so the frequency array starts ready to use.

In JavaScript, I create an array with `51` positions and initialize every position to `0`. This is important because an uninitialized JavaScript array would not behave like a normal integer frequency array.

In Python3, I use a list containing `51` zeroes. Python's `max()` function makes the `k == n` case very short.

In Go, I use an array of type `[51]int`. Go initializes integer arrays with zero values, so no separate initialization loop is required.

## Examples

### Example 1

Input:

```text
nums = [3, 9, 2, 1, 7]
k = 3
```

Subarrays:

```text
[3, 9, 2]
[9, 2, 1]
[2, 1, 7]
```

`3` appears in one subarray.

`9` appears in two subarrays.

`2` appears in three subarrays.

`1` appears in two subarrays.

`7` appears in one subarray.

The largest valid value is:

```text
7
```

Expected output:

```text
7
```

### Example 2

Input:

```text
nums = [3, 9, 7, 2, 1, 7]
k = 4
```

The size-4 subarrays are:

```text
[3, 9, 7, 2]
[9, 7, 2, 1]
[7, 2, 1, 7]
```

The first value is `3`, and it occurs only once.

The last value is `7`, but `7` occurs twice in the array.

Therefore, `3` is the only valid candidate.

Expected output:

```text
3
```

### Example 3

Input:

```text
nums = [0, 0]
k = 1
```

The subarrays are:

```text
[0]
[0]
```

The value `0` occurs twice, so it does not appear in exactly one subarray.

There is no valid integer.

Expected output:

```text
-1
```

## How to Use / Run Locally

### C++

Create a file named `main.cpp` and place the C++ solution inside it.

Compile it with:

```bash
g++ -std=c++17 main.cpp -o main
```

Run it with:

```bash
./main
```

On Windows, you can run:

```bash
main.exe
```

The LeetCode `Solution` class itself is normally submitted directly to LeetCode. For local execution, you can add a `main()` function and provide your own test cases.

### Java

Create a file named `Main.java`.

Compile it with:

```bash
javac Main.java
```

Run it with:

```bash
java Main
```

For local testing, add a `main()` method that creates the input array and calls `largestInteger()`.

### JavaScript

Create a file named `main.js`.

Run it using Node.js:

```bash
node main.js
```

You can create an array such as:

```text
[3, 9, 2, 1, 7]
```

and call the solution function with `k = 3`.

### Python3

Create a file named `main.py`.

Run it with:

```bash
python3 main.py
```

You can add a small test section at the bottom of the file to call the `Solution` class.

### Go

Create a file named `main.go`.

Run it with:

```bash
go run main.go
```

For a compiled executable, use:

```bash
go build main.go
```

Then run the generated executable.

## Notes & Optimizations

The biggest optimization is avoiding the creation of every subarray.

A straightforward solution could examine all `k`-length subarrays and count their values. However, the important structural observation makes that unnecessary.

The three cases are especially important:

* `k == 1`: find the largest value occurring exactly once.
* `k == n`: find the largest value in the entire array.
* `1 < k < n`: only the first and last elements can be candidates.

Another useful optimization comes from the constraints. Since every value is between `0` and `50`, a frequency array is better than a hash map here. It is simpler, faster in practice, and uses constant space.

An important edge case is when the first and last elements have the same value. The frequency check handles this correctly. If that value occurs twice, neither endpoint is considered valid.

For example:

```text
nums = [5, 2, 3, 5]
k = 2
```

The value `5` appears in two different subarrays, so it cannot be the answer.

This approach gives an `O(n)` time complexity and `O(1)` extra space, which is optimal for the given constraints.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
