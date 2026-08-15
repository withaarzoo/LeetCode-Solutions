# 3702. Longest Subsequence With Non-Zero Bitwise XOR

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

The problem gives me an integer array `nums` and asks me to find the length of the longest subsequence whose bitwise XOR is not zero.

A subsequence can be formed by removing any number of elements while keeping the order of the remaining elements.

I need to return the maximum possible length of such a subsequence. If no valid subsequence exists, I return `0`.

For example, with `nums = [1, 2, 3]`, the XOR of the complete array is:

`1 XOR 2 XOR 3 = 0`

So I cannot use all three elements. But I can remove `1` and use `[2, 3]`:

`2 XOR 3 = 1`

Since the XOR is non-zero, the answer is `2`.

This solution uses a simple bitwise XOR observation instead of dynamic programming or generating subsequences.

## Constraints

* `1 <= nums.length <= 10^5`
* `0 <= nums[i] <= 10^9`

## Intuition

I first looked at the XOR of the entire array.

If the XOR of all elements is already non-zero, I do not need to remove anything. The entire array is itself the longest possible subsequence, so the answer is simply `n`.

The interesting case is when the XOR of the entire array is `0`.

Suppose I remove one element `x`. The XOR of the remaining elements becomes:

`0 XOR x = x`

So if I can find any non-zero element, removing that element makes the remaining XOR non-zero.

That means I can always keep `n - 1` elements.

The only time this does not work is when every element is `0`. In that case, every possible subsequence also has XOR `0`, so the answer is `0`.

This gives me a very small set of cases:

* Total XOR is non-zero → return `n`.
* Total XOR is zero and at least one element is non-zero → return `n - 1`.
* All elements are zero → return `0`.

## Approach

I solve the problem with one pass through the array.

1. I calculate the XOR of every element.
2. At the same time, I check whether the array contains at least one non-zero element.
3. After the loop, I check the total XOR.
4. If the total XOR is non-zero, I return the full array length.
5. If the total XOR is zero but a non-zero element exists, I return one less than the array length.
6. If the total XOR is zero and every element is zero, I return `0`.

There is no need to actually construct the subsequence. I only need its maximum possible length.

## Data Structures Used

No extra data structure is required.

I only use:

* `xorValue` — stores the XOR of all elements.
* `hasNonZero` — records whether at least one non-zero element exists.

Both are simple variables, so the extra space remains constant.

## Operations & Behavior Summary

The algorithm can be viewed as this plain-English pseudocode:

1. Set the total XOR to `0`.
2. Set `hasNonZero` to `false`.
3. Visit every number in `nums`.
4. XOR the current number into the total XOR.
5. If the current number is non-zero, mark `hasNonZero` as `true`.
6. After processing the array:

   * If total XOR is not zero, return `n`.
   * Otherwise, if a non-zero element exists, return `n - 1`.
   * Otherwise, return `0`.

The main idea is that when the complete XOR is zero, removing one non-zero value changes the remaining XOR to exactly that non-zero value.

## Complexity

| Complexity       | Result | Explanation                                                                                      |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------ |
| Time Complexity  | `O(n)` | `n` is the number of elements in `nums`. I scan the array exactly once.                          |
| Space Complexity | `O(1)` | I only use a few variables and do not create any extra array, map, set, or other data structure. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int longestSubsequence(vector<int>& nums) {
        int xorValue = 0; // Stores the XOR of all elements.
        bool hasNonZero = false; // Tracks whether at least one element is non-zero.

        for (int x : nums) {
            xorValue ^= x; // Add the current element to the total XOR.

            if (x != 0) {
                hasNonZero = true; // A non-zero element can be removed if needed.
            }
        }

        if (xorValue != 0) {
            return nums.size(); // The entire array already has a non-zero XOR.
        }

        if (hasNonZero) {
            return nums.size() - 1; // Remove one non-zero element to make XOR non-zero.
        }

        return 0; // All elements are zero, so every subsequence has XOR zero.
    }
};
```

### Java

```java
class Solution {
    public int longestSubsequence(int[] nums) {
        int xorValue = 0; // Stores the XOR of all elements.
        boolean hasNonZero = false; // Tracks whether at least one element is non-zero.

        for (int x : nums) {
            xorValue ^= x; // Add the current element to the total XOR.

            if (x != 0) {
                hasNonZero = true; // A non-zero element can be removed if needed.
            }
        }

        if (xorValue != 0) {
            return nums.length; // The entire array already has a non-zero XOR.
        }

        if (hasNonZero) {
            return nums.length - 1; // Remove one non-zero element to make XOR non-zero.
        }

        return 0; // All elements are zero, so every subsequence has XOR zero.
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var longestSubsequence = function(nums) {
    let xorValue = 0; // Stores the XOR of all elements.
    let hasNonZero = false; // Tracks whether at least one element is non-zero.

    for (const x of nums) {
        xorValue ^= x; // Add the current element to the total XOR.

        if (x !== 0) {
            hasNonZero = true; // A non-zero element can be removed if needed.
        }
    }

    if (xorValue !== 0) {
        return nums.length; // The entire array already has a non-zero XOR.
    }

    if (hasNonZero) {
        return nums.length - 1; // Remove one non-zero element to make XOR non-zero.
    }

    return 0; // All elements are zero, so every subsequence has XOR zero.
};
```

### Python3

```python
from typing import List

class Solution:
    def longestSubsequence(self, nums: List[int]) -> int:
        xor_value = 0  # Stores the XOR of all elements.
        has_non_zero = False  # Tracks whether at least one element is non-zero.

        for x in nums:
            xor_value ^= x  # Add the current element to the total XOR.

            if x != 0:
                has_non_zero = True  # A non-zero element can be removed if needed.

        if xor_value != 0:
            return len(nums)  # The entire array already has a non-zero XOR.

        if has_non_zero:
            return len(nums) - 1  # Remove one non-zero element to make XOR non-zero.

        return 0  # All elements are zero, so every subsequence has XOR zero.
```

### Go

```go
func longestSubsequence(nums []int) int {
 xorValue := 0 // Stores the XOR of all elements.
 hasNonZero := false // Tracks whether at least one element is non-zero.

 for _, x := range nums {
  xorValue ^= x // Add the current element to the total XOR.

  if x != 0 {
   hasNonZero = true // A non-zero element can be removed if needed.
  }
 }

 if xorValue != 0 {
  return len(nums) // The entire array already has a non-zero XOR.
 }

 if hasNonZero {
  return len(nums) - 1 // Remove one non-zero element to make XOR non-zero.
 }

 return 0 // All elements are zero, so every subsequence has XOR zero.
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core logic is the same in all five languages. Only the syntax changes.

### Step 1: Calculate the total XOR

I start with a variable containing `0` and XOR every element into it.

For an array such as:

`[2, 3, 4]`

the calculation becomes:

`0 XOR 2 XOR 3 XOR 4`

which gives:

`5`

Since `5` is non-zero, the entire array is already valid.

### Step 2: Check whether a non-zero element exists

While calculating the XOR, I also check every element.

If I find a value other than `0`, I remember that fact.

I do not need to store the actual value because I only care whether at least one non-zero value exists.

### Step 3: Handle a non-zero total XOR

If the XOR of the complete array is non-zero, the answer is the length of the array.

For example:

`nums = [2, 3, 4]`

The XOR is:

`2 XOR 3 XOR 4 = 5`

Since `5 != 0`, I can keep all elements.

The answer is `3`.

### Step 4: Handle a zero total XOR

Now suppose:

`nums = [1, 2, 3]`

The total XOR is:

`1 XOR 2 XOR 3 = 0`

So I cannot use all three elements.

But the array contains non-zero elements.

If I remove `1`, the remaining XOR is:

`2 XOR 3 = 1`

So I can keep `2` elements.

More generally, if the XOR of the complete array is zero and I remove any non-zero element `x`, the remaining XOR is:

`x`

because:

`0 XOR x = x`

Therefore, `n - 1` is always possible when at least one non-zero element exists.

### Step 5: Handle an array containing only zeroes

Consider:

`nums = [0, 0, 0]`

The XOR is:

`0 XOR 0 XOR 0 = 0`

There is also no non-zero element to remove.

Every subsequence consists only of zeroes, so its XOR is always zero.

Therefore, no valid subsequence exists and the answer is `0`.

### C++ Behavior

In C++, I use `^` for bitwise XOR. The range-based `for` loop lets me visit every element without needing an index.

The solution returns an `int`, which is sufficient because the maximum array length is `10^5`.

### Java Behavior

In Java, `^` is also the bitwise XOR operator for integers.

The array length is available through `nums.length`, and the return type is `int`.

The logic is identical to the C++ version.

### JavaScript Behavior

JavaScript uses `^` for bitwise XOR as well.

Because the input values are within the normal 32-bit signed integer range after bitwise conversion, this operation works correctly for the given constraints.

I use `nums.length` to get the number of elements.

### Python3 Behavior

Python uses `^` as the XOR operator.

I can update the XOR variable directly with:

`xor_value ^= x`

Python integers handle the required values without needing a special integer type.

### Go Behavior

Go also uses `^` for XOR.

I iterate through the slice using a `for` loop and keep the XOR in an integer variable.

The slice length is obtained with `len(nums)`.

In every language, the important part is not the syntax. The important part is the same XOR observation that reduces the problem to three simple cases.

## Examples

### Example 1

Input:

```text
nums = [1, 2, 3]
```

Total XOR:

```text
1 XOR 2 XOR 3 = 0
```

The array contains non-zero elements, so I remove one element.

For example:

```text
[2, 3]
```

Its XOR is:

```text
2 XOR 3 = 1
```

The XOR is non-zero.

Expected output:

```text
2
```

### Example 2

Input:

```text
nums = [2, 3, 4]
```

Total XOR:

```text
2 XOR 3 XOR 4 = 5
```

The XOR is already non-zero, so I can keep the complete array.

Expected output:

```text
3
```

### Example 3

Input:

```text
nums = [0, 0, 0]
```

Total XOR:

```text
0 XOR 0 XOR 0 = 0
```

There is no non-zero element.

Every possible subsequence has XOR `0`.

Expected output:

```text
0
```

## How to Use / Run Locally

The repository contains the same algorithm implemented in C++, Java, JavaScript, Python3, and Go.

### C++

Save the solution in a `.cpp` file and compile it with:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Then run:

```bash
./solution
```

For Windows, the generated executable can be run with:

```bash
solution.exe
```

### Java

Save the solution in a Java file.

Compile it with:

```bash
javac Solution.java
```

Then run:

```bash
java Solution
```

For a direct LeetCode submission, the `Solution` class can be submitted without adding a separate `main` method.

### JavaScript

Save the solution in a `.js` file.

Run it with Node.js:

```bash
node solution.js
```

Make sure Node.js is installed and available through your terminal.

### Python3

Save the solution in a `.py` file.

Run:

```bash
python3 solution.py
```

On some Windows installations, the command may be:

```bash
python solution.py
```

### Go

Save the solution in a `.go` file.

Run it directly with:

```bash
go run solution.go
```

You can also build an executable with:

```bash
go build solution.go
```

## Notes & Optimizations

The biggest optimization is avoiding subsequence generation.

There can be an enormous number of subsequences, so trying every possible subsequence would be far too slow for `n = 10^5`.

I also do not need dynamic programming because the answer depends only on the XOR of the complete array and whether at least one non-zero value exists.

An important edge case is an array containing only zeroes. The answer must be `0` because XOR of any number of zeroes is always zero.

Another useful observation is that when the complete XOR is zero and at least one value is non-zero, removing exactly one non-zero element is always enough. There is no need to remove multiple elements.

So the final result can be summarized as:

```text
total XOR != 0  →  n
total XOR == 0 and a non-zero value exists  →  n - 1
all values are zero  →  0
```

This gives an optimal `O(n)` time and `O(1)` extra space solution for LeetCode 3702, Longest Subsequence With Non-Zero Bitwise XOR.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
