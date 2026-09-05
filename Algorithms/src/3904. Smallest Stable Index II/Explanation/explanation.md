# Smallest Stable Index II - LeetCode 3904

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

In LeetCode 3904, **Smallest Stable Index II**, I am given an integer array `nums` and an integer `k`.

For every index `i`, I need to calculate its instability score:

```text
instability(i) = max(nums[0...i]) - min(nums[i...n-1])
```

The first part looks at everything from the beginning of the array up to index `i`.

The second part looks at everything from index `i` to the end of the array.

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
Index:       0   1   2   3
Score:       5   5   4   1
```

The first score that is at most `3` is at index `3`, so the answer is:

```text
3
```

## Constraints

| Constraint    | Value                  |
| ------------- | ---------------------- |
| `nums.length` | `1 <= n <= 10^5`       |
| `nums[i]`     | `0 <= nums[i] <= 10^9` |
| `k`           | `0 <= k <= 10^9`       |

The array can contain up to `100,000` elements, so an `O(n²)` solution would be too slow.

## Intuition

My first thought was to check every index and calculate the maximum on the left and minimum on the right.

The problem with that idea is that I would scan the same elements again and again.

For example, when checking index `2`, I need the maximum of:

```text
[nums[0], nums[1], nums[2]]
```

Then for index `3`, I need:

```text
[nums[0], nums[1], nums[2], nums[3]]
```

Most of the work is repeated.

I realized that these maximum and minimum values can be calculated once and stored.

I use:

* A **prefix maximum array** to store the maximum value from the beginning up to every index.
* A **suffix minimum array** to store the minimum value from every index to the end.

Then the instability score at any index can be calculated in `O(1)` time.

## Approach

I solve the problem in three main steps.

### 1. Build the prefix maximum array

I create `prefixMax`, where:

```text
prefixMax[i] = maximum value from nums[0] to nums[i]
```

For:

```text
nums = [5, 0, 1, 4]
```

the prefix maximum array becomes:

```text
nums:       5   0   1   4
            |   |   |   |
prefixMax:  5   5   5   5
```

I build this from left to right.

### 2. Build the suffix minimum array

I create `suffixMin`, where:

```text
suffixMin[i] = minimum value from nums[i] to nums[n - 1]
```

For the same array:

```text
nums:       5   0   1   4
            |   |   |   |
suffixMin:  0   0   1   4
```

I build this from right to left.

### 3. Find the first stable index

Now I already know both values needed for every index.

For each index:

```text
instability = prefixMax[i] - suffixMin[i]
```

I check whether:

```text
instability <= k
```

Since I scan from left to right, the first valid index is automatically the smallest stable index.

If no index is stable, I return `-1`.

## Data Structures Used

### Prefix Maximum Array

`prefixMax` stores the largest value from index `0` through the current index.

It avoids repeatedly scanning the left side of the array.

### Suffix Minimum Array

`suffixMin` stores the smallest value from the current index through the last index.

It avoids repeatedly scanning the right side of the array.

### No Other Data Structures

I only need these two arrays. The solution does not require sorting, maps, sets, stacks, or queues.

## Operations & Behavior Summary

The algorithm works like this:

1. Store the length of `nums`.
2. Create the `prefixMax` array.
3. Set `prefixMax[0]` to `nums[0]`.
4. Move from left to right and calculate every prefix maximum.
5. Create the `suffixMin` array.
6. Set `suffixMin[n - 1]` to `nums[n - 1]`.
7. Move from right to left and calculate every suffix minimum.
8. Scan the array from index `0`.
9. For each index, calculate:

   ```text
   prefixMax[i] - suffixMin[i]
   ```

10. If the result is at most `k`, return that index.
11. If no index works, return `-1`.

## Complexity

| Type             | Complexity | Explanation                                                                    |
| ---------------- | ---------- | ------------------------------------------------------------------------------ |
| Time Complexity  | `O(n)`     | I build the prefix maximum, build the suffix minimum, and scan the array once. |
| Space Complexity | `O(n)`     | I use two extra arrays of size `n`.                                            |

Here, `n` is the number of elements in `nums`.

This is efficient enough for the constraint `n <= 10^5`.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int firstStableIndex(vector<int>& nums, int k) {
        int n = nums.size(); // Store the number of elements in the array.

        vector<int> prefixMax(n); // prefixMax[i] stores max(nums[0..i]).
        vector<int> suffixMin(n); // suffixMin[i] stores min(nums[i..n-1]).

        prefixMax[0] = nums[0]; // For index 0, the prefix contains only nums[0].

        for (int i = 1; i < n; i++) { // Build prefix maximums from left to right.
            prefixMax[i] = max(prefixMax[i - 1], nums[i]); // Keep the largest value seen so far.
        }

        suffixMin[n - 1] = nums[n - 1]; // For the last index, the suffix contains only nums[n-1].

        for (int i = n - 2; i >= 0; i--) { // Build suffix minimums from right to left.
            suffixMin[i] = min(suffixMin[i + 1], nums[i]); // Keep the smallest value in the suffix.
        }

        for (int i = 0; i < n; i++) { // Check every index from smallest to largest.
            int instability = prefixMax[i] - suffixMin[i]; // Calculate the instability score at i.

            if (instability <= k) { // A score at most k means this index is stable.
                return i; // Since we scan left to right, this is the smallest stable index.
            }
        }

        return -1; // No index satisfies the stability condition.
    }
};
```

### Java

```java
class Solution {
    public int firstStableIndex(int[] nums, int k) {
        int n = nums.length; // Store the number of elements in the array.

        int[] prefixMax = new int[n]; // prefixMax[i] stores max(nums[0..i]).
        int[] suffixMin = new int[n]; // suffixMin[i] stores min(nums[i..n-1]).

        prefixMax[0] = nums[0]; // For index 0, the prefix contains only nums[0].

        for (int i = 1; i < n; i++) { // Build prefix maximums from left to right.
            prefixMax[i] = Math.max(prefixMax[i - 1], nums[i]); // Keep the largest value seen so far.
        }

        suffixMin[n - 1] = nums[n - 1]; // For the last index, the suffix contains only nums[n-1].

        for (int i = n - 2; i >= 0; i--) { // Build suffix minimums from right to left.
            suffixMin[i] = Math.min(suffixMin[i + 1], nums[i]); // Keep the smallest value in the suffix.
        }

        for (int i = 0; i < n; i++) { // Check indices from smallest to largest.
            int instability = prefixMax[i] - suffixMin[i]; // Calculate the score for index i.

            if (instability <= k) { // The index is stable when its score is at most k.
                return i; // This is the first stable index because we scan left to right.
            }
        }

        return -1; // Return -1 when no stable index exists.
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
    const n = nums.length; // Store the number of elements in the array.

    const prefixMax = new Array(n); // prefixMax[i] stores max(nums[0..i]).
    const suffixMin = new Array(n); // suffixMin[i] stores min(nums[i..n-1]).

    prefixMax[0] = nums[0]; // For index 0, the prefix contains only nums[0].

    for (let i = 1; i < n; i++) { // Build prefix maximums from left to right.
        prefixMax[i] = Math.max(prefixMax[i - 1], nums[i]); // Keep the largest value seen so far.
    }

    suffixMin[n - 1] = nums[n - 1]; // For the last index, the suffix contains only nums[n-1].

    for (let i = n - 2; i >= 0; i--) { // Build suffix minimums from right to left.
        suffixMin[i] = Math.min(suffixMin[i + 1], nums[i]); // Keep the smallest value in the suffix.
    }

    for (let i = 0; i < n; i++) { // Check every index from smallest to largest.
        const instability = prefixMax[i] - suffixMin[i]; // Calculate the score for index i.

        if (instability <= k) { // The index is stable when its score is at most k.
            return i; // This is the smallest stable index because we scan left to right.
        }
    }

    return -1; // No stable index was found.
};
```

### Python3

```python
class Solution:

    def firstStableIndex(self, nums: list[int], k: int) -> int:
        n = len(nums)  # Store the number of elements in the array.

        prefixMax = [0] * n  # prefixMax[i] stores max(nums[0..i]).
        suffixMin = [0] * n  # suffixMin[i] stores min(nums[i..n-1]).

        prefixMax[0] = nums[0]  # For index 0, the prefix contains only nums[0].

        for i in range(1, n):  # Build prefix maximums from left to right.
            prefixMax[i] = max(prefixMax[i - 1], nums[i])  # Keep the largest value seen so far.

        suffixMin[n - 1] = nums[n - 1]  # For the last index, the suffix contains only nums[n-1].

        for i in range(n - 2, -1, -1):  # Build suffix minimums from right to left.
            suffixMin[i] = min(suffixMin[i + 1], nums[i])  # Keep the smallest value in the suffix.

        for i in range(n):  # Check indices from smallest to largest.
            instability = prefixMax[i] - suffixMin[i]  # Calculate the score for index i.

            if instability <= k:  # The index is stable when its score is at most k.
                return i  # This is the smallest stable index because we scan left to right.

        return -1  # No stable index was found.
```

### Go

```go
func firstStableIndex(nums []int, k int) int {
 n := len(nums) // Store the number of elements in the array.

 prefixMax := make([]int, n) // prefixMax[i] stores max(nums[0..i]).
 suffixMin := make([]int, n) // suffixMin[i] stores min(nums[i..n-1]).

 prefixMax[0] = nums[0] // For index 0, the prefix contains only nums[0].

 for i := 1; i < n; i++ { // Build prefix maximums from left to right.
  prefixMax[i] = prefixMax[i-1] // Start with the previous prefix maximum.
  if nums[i] > prefixMax[i] { // Check whether the current value is larger.
   prefixMax[i] = nums[i] // Update the maximum when nums[i] is larger.
  }
 }

 suffixMin[n-1] = nums[n-1] // For the last index, the suffix contains only nums[n-1].

 for i := n - 2; i >= 0; i-- { // Build suffix minimums from right to left.
  suffixMin[i] = suffixMin[i+1] // Start with the next suffix minimum.
  if nums[i] < suffixMin[i] { // Check whether the current value is smaller.
   suffixMin[i] = nums[i] // Update the minimum when nums[i] is smaller.
  }
 }

 for i := 0; i < n; i++ { // Check every index from smallest to largest.
  instability := prefixMax[i] - suffixMin[i] // Calculate the score for index i.

  if instability <= k { // The index is stable when its score is at most k.
   return i // This is the smallest stable index because we scan left to right.
  }
 }

 return -1 // No stable index was found.
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The main logic is the same in all five languages. Only the syntax used to create arrays, find maximum or minimum values, and loop through the array changes.

### C++

I first get the size of the array.

```text
n = nums.size()
```

Then I create two vectors of size `n`.

The first vector stores prefix maximums, while the second stores suffix minimums.

I initialize the first prefix value with `nums[0]` because the prefix at index `0` contains only one element.

Then I move from index `1` toward the end.

At every index, I compare the previous prefix maximum with the current number. The larger value becomes the new prefix maximum.

For the suffix array, I start from the last index because the suffix at the last position contains only that element.

I then move from right to left. At every position, I compare the current number with the suffix minimum already calculated for the next position.

Once both arrays are ready, I scan from left to right.

At each index:

```text
instability = prefixMax[i] - suffixMin[i]
```

If the instability is at most `k`, I immediately return the index.

Returning immediately is important because the problem asks for the smallest stable index.

### Java

The Java solution follows exactly the same idea.

I use two integer arrays:

```text
int[] prefixMax
int[] suffixMin
```

The prefix maximum is calculated from left to right using `Math.max`.

The suffix minimum is calculated from right to left using `Math.min`.

After preprocessing both arrays, I check every index in increasing order.

The expression:

```text
prefixMax[i] - suffixMin[i]
```

gives the exact instability score required by the problem.

When the score becomes less than or equal to `k`, I return that index.

Java's `int` type is enough here because both `nums[i]` and `k` are at most `10^9`, and the difference between two values in the allowed range also fits within a 32-bit signed integer.

### JavaScript

In JavaScript, I use `Array` objects for `prefixMax` and `suffixMin`.

I initialize their values as I process the input rather than filling them with meaningful values beforehand.

`Math.max` is used for the prefix calculation:

```text
prefixMax[i] = Math.max(prefixMax[i - 1], nums[i])
```

For the suffix calculation:

```text
suffixMin[i] = Math.min(suffixMin[i + 1], nums[i])
```

After that, the final scan checks the instability score for every index.

JavaScript uses the `Number` type for these values. The largest input value here is `10^9`, which is safely represented as an integer by JavaScript.

### Python3

In Python, I create the two arrays with:

```text
[0] * n
```

I calculate `prefixMax` from left to right using Python's built-in `max` function.

Then I calculate `suffixMin` from right to left using `min`.

Python's range allows me to write the reverse loop clearly:

```text
range(n - 2, -1, -1)
```

This visits:

```text
n - 2, n - 3, ..., 1, 0
```

Finally, I scan from `0` to `n - 1` and return the first index whose instability score is at most `k`.

### Go

In Go, I create two integer slices using `make`.

Go does not have a built-in `max` or `min` call that I need to rely on for this simple logic, so I compare the values directly.

For the prefix maximum, I start with the previous maximum and replace it when `nums[i]` is larger.

For the suffix minimum, I start with the next suffix minimum and replace it when `nums[i]` is smaller.

The final scan is the same as in the other languages.

I return the index as soon as I find a stable position. If the loop ends without finding one, I return `-1`.

## Examples

### Example 1

Input:

```text
nums = [5, 0, 1, 4]
k = 3
```

Prefix maximums:

```text
[5, 5, 5, 5]
```

Suffix minimums:

```text
[0, 0, 1, 4]
```

Now I calculate:

```text
Index 0: 5 - 0 = 5
Index 1: 5 - 0 = 5
Index 2: 5 - 1 = 4
Index 3: 5 - 4 = 1
```

Since `1 <= 3`, index `3` is stable.

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

Prefix maximums:

```text
[3, 3, 3]
```

Suffix minimums:

```text
[1, 1, 1]
```

Instability scores:

```text
Index 0: 3 - 1 = 2
Index 1: 3 - 1 = 2
Index 2: 3 - 1 = 2
```

Every score is greater than `1`.

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

```text
prefixMax = [0]
suffixMin = [0]
```

The instability score is:

```text
0 - 0 = 0
```

Since `0 <= 0`, index `0` is stable.

Expected output:

```text
0
```

## How to Use / Run Locally

The code blocks above are intentionally empty so the actual implementations can be added separately.

### C++

Save the completed solution in a file such as:

```text
solution.cpp
```

Compile it with:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Then run:

```bash
./solution
```

On Windows, the generated executable can be run with:

```bash
solution.exe
```

### Java

Save the completed solution in:

```text
Solution.java
```

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

For a LeetCode submission, the `Solution` class and method can be submitted directly in the LeetCode editor.

### JavaScript

Save the completed solution as:

```text
solution.js
```

Run it with Node.js:

```bash
node solution.js
```

You can also paste the function directly into the LeetCode JavaScript editor.

### Python3

Save the completed solution as:

```text
solution.py
```

Run it with:

```bash
python3 solution.py
```

On some Windows installations, this may be:

```bash
python solution.py
```

### Go

Save the completed solution as:

```text
solution.go
```

Run it with:

```bash
go run solution.go
```

You can also compile it with:

```bash
go build solution.go
```

## Notes & Optimizations

The main optimization is avoiding repeated range scans.

A direct solution could calculate the maximum and minimum for every index separately, but that can take `O(n²)` time in the worst case.

The prefix maximum array removes repeated work on the left side.

The suffix minimum array removes repeated work on the right side.

Another important detail is that index `i` belongs to both ranges:

```text
nums[0...i]
nums[i...n-1]
```

So `nums[i]` must be considered when calculating both `prefixMax[i]` and `suffixMin[i]`.

The single-element case also needs to work correctly:

```text
nums = [x]
```

For index `0`, both the maximum and minimum are `x`, so the instability score is always:

```text
x - x = 0
```

Therefore, index `0` is stable whenever `k >= 0`, which is always true under the given constraints.

The final scan should always go from left to right. If I scanned in another order, I could find a stable index but not necessarily the smallest one.

The overall solution uses `O(n)` time and `O(n)` extra space, which is suitable for arrays containing up to `10^5` elements.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)

---
