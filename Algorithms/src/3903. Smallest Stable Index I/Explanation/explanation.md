# 3903. Smallest Stable Index I

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

In this problem, I am given an integer array `nums` and an integer `k`.

For every index `i`, I need to calculate its instability score:

```text
max(nums[0..i]) - min(nums[i..n-1])
```

Here:

* `max(nums[0..i])` is the largest value from index `0` to `i`.
* `min(nums[i..n-1])` is the smallest value from index `i` to the last index.

An index is called **stable** when its instability score is less than or equal to `k`.

My task is to return the **smallest stable index**.

If there is no stable index, I return `-1`.

For example:

```text
nums = [5, 0, 1, 4]
k = 3
```

The instability scores are:

```text
index 0 -> 5
index 1 -> 5
index 2 -> 4
index 3 -> 1
```

The first score that is at most `3` occurs at index `3`, so the answer is `3`.

## Constraints

* `1 <= nums.length <= 100`
* `0 <= nums[i] <= 10^9`
* `0 <= k <= 10^9`

## Intuition

My first thought was to directly calculate the maximum on the left and the minimum on the right for every index.

That would work, but doing those calculations again and again is unnecessary.

I noticed that the maximum from the left can be maintained while scanning the array from left to right.

For the minimum from the current index to the end, I can precompute a **suffix minimum array**. Once that array is ready, I can get the required right-side minimum in constant time for every index.

So I use two passes:

1. Build suffix minimums from right to left.
2. Scan from left to right while maintaining the prefix maximum.

As soon as the instability score becomes `<= k`, I return that index.

Because I check the indices in increasing order, the first valid index is automatically the smallest stable index.

## Approach

I solve the problem in three simple steps.

1. I create a `suffixMin` array.

   `suffixMin[i]` stores the minimum value from `nums[i]` through `nums[n - 1]`.

2. I scan the array from left to right and maintain `prefixMax`.

   `prefixMax` stores the maximum value from `nums[0]` through the current index.

3. At every index `i`, I calculate:

```text
instability = prefixMax - suffixMin[i]
```

If:

```text
instability <= k
```

I return `i`.

If the complete array is checked without finding a stable index, I return `-1`.

## Data Structures Used

### Suffix Minimum Array

I use one array called `suffixMin`.

For every index `i`:

```text
suffixMin[i] = min(nums[i], nums[i + 1], ..., nums[n - 1])
```

This allows me to find the minimum value on the right side in `O(1)` time during the final scan.

### Prefix Maximum Variable

I do not need another array for the prefix maximum.

I store it in a single variable called `prefixMax`.

Whenever I move to the next index, I update it using the current element.

## Operations & Behavior Summary

The algorithm works like this:

1. Find the length of `nums`.
2. Create a `suffixMin` array of the same size.
3. Set the last suffix minimum to the last element.
4. Move from right to left and calculate every suffix minimum.
5. Set `prefixMax` to the first element.
6. Move from left to right.
7. Update `prefixMax` with the current element.
8. Calculate the instability score using:
   `prefixMax - suffixMin[i]`.
9. If the score is at most `k`, return the current index.
10. If no index works, return `-1`.

This gives a linear-time solution and avoids repeatedly scanning the same parts of the array.

## Complexity

| Complexity       | Result | Explanation                                                                                                          |
| ---------------- | ------ | -------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n)` | I build the suffix minimum array in one pass and search for the answer in another pass. `n` is the length of `nums`. |
| Space Complexity | `O(n)` | I use the `suffixMin` array with `n` elements. The prefix maximum only needs one variable.                           |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int firstStableIndex(vector<int>& nums, int k) {
        int n = nums.size();

        // suffixMin[i] stores the minimum value from i to n - 1.
        vector<int> suffixMin(n);

        // For the last index, the suffix contains only nums[n - 1].
        suffixMin[n - 1] = nums[n - 1];

        // Build suffix minimums from right to left.
        for (int i = n - 2; i >= 0; --i) {
            // Keep the smaller value between the current element
            // and the minimum of the suffix to its right.
            suffixMin[i] = min(nums[i], suffixMin[i + 1]);
        }

        // Store the largest value seen from index 0 to the current index.
        int prefixMax = nums[0];

        // Check every index from smallest to largest.
        for (int i = 0; i < n; ++i) {
            // Update the maximum value in nums[0..i].
            prefixMax = max(prefixMax, nums[i]);

            // Calculate the instability score for this index.
            int instability = prefixMax - suffixMin[i];

            // Since indices are checked from left to right,
            // this is automatically the smallest stable index.
            if (instability <= k) {
                return i;
            }
        }

        // No stable index exists.
        return -1;
    }
};
```

### Java

```java
class Solution {
    public int firstStableIndex(int[] nums, int k) {
        int n = nums.length;

        // suffixMin[i] stores the minimum value from i to n - 1.
        int[] suffixMin = new int[n];

        // For the last index, the suffix contains only nums[n - 1].
        suffixMin[n - 1] = nums[n - 1];

        // Build suffix minimums from right to left.
        for (int i = n - 2; i >= 0; i--) {
            // Keep the smaller value between the current element
            // and the minimum value already calculated to its right.
            suffixMin[i] = Math.min(nums[i], suffixMin[i + 1]);
        }

        // Store the largest value seen from index 0 to the current index.
        int prefixMax = nums[0];

        // Check indices from smallest to largest.
        for (int i = 0; i < n; i++) {
            // Update the maximum value in nums[0..i].
            prefixMax = Math.max(prefixMax, nums[i]);

            // Calculate the instability score at index i.
            int instability = prefixMax - suffixMin[i];

            // The first valid index is the smallest stable index.
            if (instability <= k) {
                return i;
            }
        }

        // No stable index was found.
        return -1;
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
var firstStableIndex = function(nums, k) {
    const n = nums.length;

    // suffixMin[i] stores the minimum value from i to the end.
    const suffixMin = new Array(n);

    // The last suffix contains only the last element.
    suffixMin[n - 1] = nums[n - 1];

    // Build suffix minimums from right to left.
    for (let i = n - 2; i >= 0; i--) {
        // Store the smaller value between nums[i]
        // and the minimum of the suffix to its right.
        suffixMin[i] = Math.min(nums[i], suffixMin[i + 1]);
    }

    // Store the largest value seen from index 0 to the current index.
    let prefixMax = nums[0];

    // Check indices from left to right.
    for (let i = 0; i < n; i++) {
        // Update the maximum value in nums[0..i].
        prefixMax = Math.max(prefixMax, nums[i]);

        // Calculate the instability score at index i.
        const instability = prefixMax - suffixMin[i];

        // Return immediately because this is the first valid index.
        if (instability <= k) {
            return i;
        }
    }

    // No stable index exists.
    return -1;
};
```

### Python3

```python
class Solution:

    def firstStableIndex(self, nums: list[int], k: int) -> int:
        n = len(nums)

        # suffix_min[i] stores the minimum value from i to the end.
        suffix_min = [0] * n

        # The last suffix contains only the last element.
        suffix_min[n - 1] = nums[n - 1]

        # Build suffix minimums from right to left.
        for i in range(n - 2, -1, -1):
            # Keep the smaller value between nums[i]
            # and the minimum value already found to its right.
            suffix_min[i] = min(nums[i], suffix_min[i + 1])

        # Store the largest value seen from index 0 to the current index.
        prefix_max = nums[0]

        # Check every index from smallest to largest.
        for i in range(n):
            # Update the maximum value in nums[0..i].
            prefix_max = max(prefix_max, nums[i])

            # Calculate the instability score at index i.
            instability = prefix_max - suffix_min[i]

            # This is the first stable index because we scan from left to right.
            if instability <= k:
                return i

        # No stable index exists.
        return -1
```

### Go

```go
func firstStableIndex(nums []int, k int) int {
    n := len(nums)

    // suffixMin[i] stores the minimum value from i to the end.
    suffixMin := make([]int, n)

    // The last suffix contains only the last element.
    suffixMin[n-1] = nums[n-1]

    // Build suffix minimums from right to left.
    for i := n - 2; i >= 0; i-- {
        // Store the smaller value between nums[i]
        // and the minimum value already calculated to its right.
        if nums[i] < suffixMin[i+1] {
            suffixMin[i] = nums[i]
        } else {
            suffixMin[i] = suffixMin[i+1]
        }
    }

    // Store the largest value seen from index 0 to the current index.
    prefixMax := nums[0]

    // Check indices from smallest to largest.
    for i := 0; i < n; i++ {
        // Update the maximum value in nums[0..i].
        if nums[i] > prefixMax {
            prefixMax = nums[i]
        }

        // Calculate the instability score at index i.
        instability := prefixMax - suffixMin[i]

        // Return immediately because this is the first stable index.
        if instability <= k {
            return i
        }
    }

    // No stable index exists.
    return -1
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five languages. Only the syntax changes.

### Step 1: Store the array length

I first store the length of `nums` in `n`.

I need this value to know where the array ends and to create the suffix minimum array.

For an array of length `n`, the final index is `n - 1`.

### Step 2: Create the suffix minimum array

I create an array named `suffixMin` with `n` positions.

The purpose of this array is to avoid calculating the right-side minimum repeatedly.

For example:

```text
nums = [5, 0, 1, 4]
```

The suffix minimum array becomes:

```text
suffixMin = [0, 0, 1, 4]
```

Each position represents the smallest value from that position to the end.

### Step 3: Initialize the last position

The suffix beginning at the last index contains only one element.

For:

```text
nums = [5, 0, 1, 4]
```

the suffix at index `3` is:

```text
[4]
```

So:

```text
suffixMin[3] = 4
```

This gives me the starting point for building the rest of the suffix minimum array.

### Step 4: Build suffix minimums from right to left

I then move from `n - 2` down to `0`.

At every position, I compare:

```text
nums[i]
```

with:

```text
suffixMin[i + 1]
```

The smaller one becomes `suffixMin[i]`.

For example:

```text
nums = [5, 0, 1, 4]
```

At index `2`:

```text
min(1, 4) = 1
```

At index `1`:

```text
min(0, 1) = 0
```

At index `0`:

```text
min(5, 0) = 0
```

So the result is:

```text
[0, 0, 1, 4]
```

### Step 5: Start tracking the prefix maximum

Now I scan from left to right.

I use one variable, `prefixMax`, to store:

```text
max(nums[0..i])
```

For example, with:

```text
nums = [5, 0, 1, 4]
```

the prefix maximum is:

```text
index 0 -> 5
index 1 -> 5
index 2 -> 5
index 3 -> 5
```

I do not need a complete prefix array because I only need the current maximum.

### Step 6: Calculate the instability score

At every index `i`, I now have both values required by the problem:

```text
prefixMax = max(nums[0..i])
suffixMin[i] = min(nums[i..n-1])
```

So the instability score is simply:

```text
prefixMax - suffixMin[i]
```

For:

```text
nums = [5, 0, 1, 4]
k = 3
```

at index `2`:

```text
prefixMax = 5
suffixMin[2] = 1

instability = 5 - 1
            = 4
```

Since `4 > 3`, index `2` is not stable.

At index `3`:

```text
prefixMax = 5
suffixMin[3] = 4

instability = 5 - 4
            = 1
```

Since `1 <= 3`, index `3` is stable.

### Step 7: Return immediately

I return the current index as soon as the instability score is at most `k`.

This is safe because I scan from index `0` toward the end.

That means I never skip a smaller index.

So the first valid index is always the smallest stable index.

### Step 8: Handle the case where no index works

If every index has an instability score greater than `k`, I reach the end of the loop.

In that case, there is no stable index, so I return:

```text
-1
```

### Why this works in all five languages

The C++, Java, JavaScript, Python3, and Go implementations all follow the same algorithm:

```text
Build suffix minimums
        |
        v
Scan from left to right
        |
        v
Update prefix maximum
        |
        v
Calculate instability score
        |
        v
score <= k ?
   /       \
 yes        no
  |          |
return i    continue
```

The only difference is how arrays, loops, minimum functions, and maximum functions are written in each language.

## Examples

### Example 1

Input:

```text
nums = [5, 0, 1, 4]
k = 3
```

Suffix minimums:

```text
[0, 0, 1, 4]
```

Now I scan from left to right:

| Index | Prefix Maximum | Suffix Minimum | Instability | Stable? |
| ----: | -------------: | -------------: | ----------: | ------- |
|     0 |              5 |              0 |           5 | No      |
|     1 |              5 |              0 |           5 | No      |
|     2 |              5 |              1 |           4 | No      |
|     3 |              5 |              4 |           1 | Yes     |

The first stable index is:

```text
3
```

Expected output:

```text
3
```

### Example 2

Input:

```text
nums = [3, 2, 1]
k = 1
```

The suffix minimums are:

```text
[1, 1, 1]
```

The prefix maximum is always `3`.

So:

```text
index 0 -> 3 - 1 = 2
index 1 -> 3 - 1 = 2
index 2 -> 3 - 1 = 2
```

Every score is greater than `k = 1`.

Expected output:

```text
-1
```

### Example 3

Input:

```text
nums = [0]
k = 0
```

There is only one index.

The prefix maximum is `0`, and the suffix minimum is also `0`.

So:

```text
instability = 0 - 0 = 0
```

Since:

```text
0 <= 0
```

index `0` is stable.

Expected output:

```text
0
```

## How to Use / Run Locally

The LeetCode solution is written inside the required `Solution` class or function. To run it locally, I can place the same logic inside a small driver program with sample input.

### C++

Save the solution as `main.cpp`.

Compile it with:

```bash
g++ -std=c++17 -O2 -o main main.cpp
```

Run it with:

```bash
./main
```

On Windows, the generated executable can be run with:

```bash
main.exe
```

### Java

Save the solution as `Main.java`.

Compile it with:

```bash
javac Main.java
```

Run it with:

```bash
java Main
```

For local testing, I can create a `main` method that builds the input array and calls `firstStableIndex`.

### JavaScript

Save the solution as:

```text
main.js
```

Make sure Node.js is installed, then run:

```bash
node main.js
```

I can create an array such as `[5, 0, 1, 4]`, set `k` to `3`, and call the solution function to test the result.

### Python3

Save the solution as:

```text
main.py
```

Run it with:

```bash
python3 main.py
```

On some Windows installations, the command may be:

```bash
python main.py
```

I can test the function with the sample inputs directly.

### Go

Save the program as:

```text
main.go
```

Run it directly with:

```bash
go run main.go
```

To build an executable first:

```bash
go build main.go
```

Then run the generated executable.

## Notes & Optimizations

The most important part of this solution is avoiding repeated range calculations.

A simple brute-force approach would calculate:

```text
max(nums[0..i])
```

and:

```text
min(nums[i..n-1])
```

again for every index. That can take `O(n²)` time.

Instead, I precompute all suffix minimums once and maintain the prefix maximum while scanning. This reduces the total running time to `O(n)`.

The single-element case is also handled naturally. If `nums` contains only one value, both the prefix maximum and suffix minimum are that same value, so the instability score is always `0`.

I also need to use `<= k`, not `< k`, because an index is stable when its instability score is equal to `k` as well.

The constraints are small, but the linear-time approach is still cleaner and avoids unnecessary repeated work.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
