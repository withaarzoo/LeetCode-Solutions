# 1477. Find Two Non-overlapping Sub-arrays Each With Target Sum

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
  * [TypeScript](#typescript)
  * [Python3](#python3)
  * [Go](#go)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

LeetCode 1477, **Find Two Non-overlapping Sub-arrays Each With Target Sum**, gives me an integer array `arr` and an integer `target`.

I need to find two sub-arrays whose sums are both exactly equal to `target`.

The two sub-arrays must not overlap, and among all possible pairs, I need to return the minimum possible sum of their lengths.

If I cannot find two valid non-overlapping sub-arrays, I return `-1`.

### Input

* `arr`: An array of positive integers.
* `target`: The required sum for each sub-array.

### Output

Return the minimum total length of two non-overlapping sub-arrays whose sums are both equal to `target`.

If no such pair exists, return `-1`.

This problem is a good combination of **sliding window**, **prefix-style dynamic programming**, and **array optimization**.

## Constraints

| Constraint   | Value                     |
| ------------ | ------------------------- |
| Array length | `1 <= arr.length <= 10^5` |
| Array values | `1 <= arr[i] <= 1000`     |
| Target       | `1 <= target <= 10^8`     |

Because the array can contain up to `100,000` elements, an `O(n^2)` solution can become too slow. I need an approach that runs in linear time.

## Intuition

My first thought was to find every sub-array whose sum is `target`, then try to combine every possible pair.

The problem is that there can be many valid sub-arrays. Comparing all pairs could take too much time.

The useful observation is that every `arr[i]` is positive.

That means I can use a **sliding window** to find target-sum sub-arrays efficiently:

* Expand the window from the right.
* If the sum becomes too large, move the left pointer forward.
* When the sum becomes exactly `target`, I found a valid sub-array.

The next question is how to combine two such sub-arrays efficiently.

For every index, I keep the shortest target-sum sub-array found anywhere before or at that index.

Then, whenever I find a new valid sub-array, I look at the shortest valid sub-array that ends before it starts.

So the main idea becomes:

```text
current valid sub-array
          +
shortest valid sub-array before it
          =
candidate answer
```

This avoids checking every possible pair.

## Approach

I solve the problem using two ideas together: a sliding window and a prefix minimum array.

### 1. Use a sliding window

I maintain two pointers:

* `left` for the beginning of the current window.
* `right` for the end of the current window.

I keep the current window sum in `sum`.

For every new `right` index, I add `arr[right]`.

If:

```text
sum > target
```

I keep removing elements from the left until:

```text
sum <= target
```

Because every number is positive, removing elements always decreases the sum.

When:

```text
sum == target
```

the current window is a valid sub-array.

Its length is:

```text
right - left + 1
```

### 2. Store the best previous sub-array

I use an array called `best`.

`best[i]` represents the shortest target-sum sub-array that appears completely inside:

```text
arr[0 ... i]
```

For example:

```text
arr = [3, 2, 4, 3]
target = 3
```

The valid sub-arrays are:

```text
[3]             [3]
  ^               ^
 index 0          index 3

length = 1       length = 1
```

When I reach the second `[3]`, I need a previous sub-array that ends before index `3`.

So I check:

```text
best[2]
```

That gives me the shortest valid sub-array in the part before the current window.

Therefore:

```text
previous length + current length
1 + 1
= 2
```

### 3. Prevent overlap

Suppose the current valid sub-array is:

```text
[left ........ right]
```

A previous sub-array can only end at:

```text
left - 1
```

or earlier.

That is why I use:

```text
best[left - 1]
```

This guarantees that the two sub-arrays do not overlap.

### 4. Keep the minimum answer

Every time I find a valid current window and a valid previous window, I calculate:

```text
currentLength + best[left - 1]
```

Then I keep the smallest value.

If no valid pair is found, the result is `-1`.

## Data Structures Used

### `best` array

I use a one-dimensional array of size `n`.

```text
best[i] = shortest target-sum sub-array found in arr[0 ... i]
```

I need this because I want constant-time access to the best previous sub-array when a new valid window is found.

### Sliding window variables

I also use a few integer variables:

* `left` for the window's left boundary.
* `right` for the window's right boundary.
* `sum` for the current window sum.
* `answer` for the minimum combined length.

No hash map, stack, queue, or tree is required for this optimized solution.

## Operations & Behavior Summary

The algorithm works in the following order:

1. Create a `best` array filled with infinity or a very large value.
2. Set `left = 0` and `sum = 0`.
3. Move `right` from left to right through the array.
4. Add `arr[right]` to `sum`.
5. While `sum > target`, remove elements from the left.
6. When `sum == target`, calculate the current sub-array length.
7. Check `best[left - 1]` to find the shortest valid non-overlapping sub-array before the current one.
8. Update the global minimum answer.
9. Store the shortest valid sub-array for the current prefix in `best[right]`.
10. Carry the previous prefix minimum forward when necessary.
11. Return the answer.
12. If two non-overlapping sub-arrays were never found, return `-1`.

In simple pseudocode:

```text
for every right index:
    add arr[right] to window sum

    while window sum > target:
        remove arr[left]
        move left forward

    if window sum == target:
        current length = right - left + 1

        if a previous valid sub-array exists:
            update answer

        save current length in best[right]

    keep the best previous value in best[right]
```

## Complexity

| Metric           | Complexity | Explanation                                                                                                  |
| ---------------- | ---------: | ------------------------------------------------------------------------------------------------------------ |
| Time Complexity  |     `O(n)` | `right` moves through the array once, and `left` also moves only forward. Together they perform linear work. |
| Space Complexity |     `O(n)` | The `best` array stores one value for each index.                                                            |

Here, `n` is the number of elements in `arr`.

Even though the sliding window contains a `while` loop, the solution is still `O(n)` because the `left` pointer never moves backward.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minSumOfLengths(vector<int>& arr, int target) {
        int n = arr.size();

        // best[i] stores the shortest target-sum sub-array
        // found completely inside arr[0..i].
        vector<int> best(n, INT_MAX);

        int left = 0;       // Left boundary of the sliding window.
        int sum = 0;        // Sum of the current sliding window.
        int answer = INT_MAX;

        for (int right = 0; right < n; ++right) {
            // Extend the window by including arr[right].
            sum += arr[right];

            // Because all values are positive, shrinking from the left
            // is enough whenever the window sum becomes too large.
            while (sum > target) {
                sum -= arr[left];
                ++left;
            }

            // If the current window has exactly the target sum,
            // it is a valid sub-array.
            if (sum == target) {
                int currentLength = right - left + 1;

                // A previous sub-array must end before 'left'.
                // best[left - 1] gives the shortest such sub-array.
                if (left > 0 && best[left - 1] != INT_MAX) {
                    answer = min(answer,
                                 currentLength + best[left - 1]);
                }

                // If this is the first valid sub-array in the prefix,
                // or it is shorter than the previous best, store it.
                best[right] = currentLength;
            }

            // Even when the current position does not end a valid window,
            // keep the best valid sub-array seen anywhere in the prefix.
            if (right > 0) {
                best[right] = min(best[right], best[right - 1]);
            }
        }

        // If two valid non-overlapping sub-arrays were never found,
        // return -1 as required.
        return answer == INT_MAX ? -1 : answer;
    }
};
```

### Java

```java
class Solution {
    public int minSumOfLengths(int[] arr, int target) {
        int n = arr.length;

        // best[i] stores the shortest target-sum sub-array
        // found completely inside arr[0..i].
        int[] best = new int[n];

        // Use a large value to represent "no valid sub-array".
        java.util.Arrays.fill(best, Integer.MAX_VALUE);

        int left = 0;              // Left boundary of the sliding window.
        int sum = 0;                // Sum of the current sliding window.
        int answer = Integer.MAX_VALUE;

        for (int right = 0; right < n; right++) {
            // Add the new element to the current window.
            sum += arr[right];

            // Since all values are positive, move left forward
            // until the window sum becomes at most target.
            while (sum > target) {
                sum -= arr[left];
                left++;
            }

            // A valid target-sum sub-array ends at right.
            if (sum == target) {
                int currentLength = right - left + 1;

                // The previous sub-array must finish before left.
                if (left > 0 && best[left - 1] != Integer.MAX_VALUE) {
                    answer = Math.min(
                        answer,
                        currentLength + best[left - 1]
                    );
                }

                // Store this valid sub-array as the current best
                // for the prefix ending at right.
                best[right] = currentLength;
            }

            // Carry forward the best sub-array found in the prefix
            // even if no valid sub-array ends exactly at right.
            if (right > 0) {
                best[right] = Math.min(best[right], best[right - 1]);
            }
        }

        // Return -1 when two non-overlapping valid sub-arrays do not exist.
        return answer == Integer.MAX_VALUE ? -1 : answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} arr
 * @param {number} target
 * @return {number}
 */
var minSumOfLengths = function(arr, target) {
    const n = arr.length;

    // best[i] stores the shortest target-sum sub-array
    // found completely inside arr[0..i].
    const best = new Array(n).fill(Infinity);

    let left = 0;       // Left boundary of the sliding window.
    let sum = 0;        // Sum of the current window.
    let answer = Infinity;

    for (let right = 0; right < n; right++) {
        // Expand the window by adding arr[right].
        sum += arr[right];

        // All values are positive, so shrinking from the left
        // is guaranteed to reduce the sum.
        while (sum > target) {
            sum -= arr[left];
            left++;
        }

        // The current window has exactly the target sum.
        if (sum === target) {
            const currentLength = right - left + 1;

            // best[left - 1] represents a valid sub-array
            // that ends before the current one starts.
            if (left > 0 && best[left - 1] !== Infinity) {
                answer = Math.min(
                    answer,
                    currentLength + best[left - 1]
                );
            }

            // Save the current valid sub-array length for this prefix.
            best[right] = currentLength;
        }

        // Keep the shortest valid sub-array seen anywhere up to right.
        if (right > 0) {
            best[right] = Math.min(best[right], best[right - 1]);
        }
    }

    // Return -1 if we could not build two non-overlapping sub-arrays.
    return answer === Infinity ? -1 : answer;
};
```

### TypeScript

```typescript
function minSumOfLengths(arr: number[], target: number): number {
    const n = arr.length;

    // best[i] stores the shortest target-sum sub-array
    // found completely inside arr[0..i].
    const best: number[] = new Array(n).fill(Infinity);

    let left = 0;       // Left boundary of the sliding window.
    let sum = 0;        // Sum of the current sliding window.
    let answer = Infinity;

    for (let right = 0; right < n; right++) {
        // Expand the window to include arr[right].
        sum += arr[right];

        // Since all elements are positive, removing from the left
        // always decreases the sum.
        while (sum > target) {
            sum -= arr[left];
            left++;
        }

        // We found a valid sub-array with sum equal to target.
        if (sum === target) {
            const currentLength = right - left + 1;

            // best[left - 1] is guaranteed to be non-overlapping
            // with the current window because it ends before left.
            if (left > 0 && best[left - 1] !== Infinity) {
                answer = Math.min(
                    answer,
                    currentLength + best[left - 1]
                );
            }

            // Store the current valid length for this prefix.
            best[right] = currentLength;
        }

        // Keep the shortest valid sub-array seen anywhere in the prefix.
        if (right > 0) {
            best[right] = Math.min(best[right], best[right - 1]);
        }
    }

    // Return -1 when fewer than two non-overlapping valid sub-arrays exist.
    return answer === Infinity ? -1 : answer;
}
```

### Python3

```python
from typing import List


class Solution:
    def minSumOfLengths(self, arr: List[int], target: int) -> int:
        n = len(arr)

        # best[i] stores the shortest target-sum sub-array
        # found completely inside arr[0..i].
        best = [float("inf")] * n

        left = 0              # Left boundary of the sliding window.
        total = 0             # Sum of the current sliding window.
        answer = float("inf") # Stores the minimum combined length.

        for right in range(n):
            # Expand the window by adding arr[right].
            total += arr[right]

            # All elements are positive, so moving left forward
            # always decreases the window sum.
            while total > target:
                total -= arr[left]
                left += 1

            # The current window is a valid target-sum sub-array.
            if total == target:
                current_length = right - left + 1

                # A previous sub-array must end before left,
                # so best[left - 1] guarantees no overlap.
                if left > 0 and best[left - 1] != float("inf"):
                    answer = min(
                        answer,
                        current_length + best[left - 1]
                    )

                # Save the current valid sub-array length
                # for the prefix ending at right.
                best[right] = current_length

            # Keep the shortest valid sub-array found anywhere
            # in the prefix arr[0..right].
            if right > 0:
                best[right] = min(best[right], best[right - 1])

        # If two non-overlapping sub-arrays were not found, return -1.
        return -1 if answer == float("inf") else answer
```

### Go

```go
func minSumOfLengths(arr []int, target int) int {
	n := len(arr)

	// best[i] stores the shortest target-sum sub-array
	// found completely inside arr[0..i].
	const INF = int(^uint(0) >> 1)
	best := make([]int, n)

	// Fill best with INF to represent "no valid sub-array yet".
	for i := 0; i < n; i++ {
		best[i] = INF
	}

	left := 0      // Left boundary of the sliding window.
	sum := 0       // Sum of the current sliding window.
	answer := INF  // Minimum combined length found so far.

	for right := 0; right < n; right++ {
		// Expand the window by adding arr[right].
		sum += arr[right]

		// All elements are positive, so shrinking from the left
		// always makes the window sum smaller.
		for sum > target {
			sum -= arr[left]
			left++
		}

		// The current window has exactly the target sum.
		if sum == target {
			currentLength := right - left + 1

			// best[left-1] represents a valid sub-array
			// that ends before the current window starts.
			if left > 0 && best[left-1] != INF {
				candidate := currentLength + best[left-1]

				// Keep the smallest combined length.
				if candidate < answer {
					answer = candidate
				}
			}

			// Store the current valid sub-array length.
			best[right] = currentLength
		}

		// Carry the best answer from the previous prefix.
		if right > 0 && best[right-1] < best[right] {
			best[right] = best[right-1]
		}
	}

	// Return -1 when two non-overlapping target-sum sub-arrays do not exist.
	if answer == INF {
		return -1
	}

	return answer
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The implementation is available in six languages, but the core logic is the same in each one.

### C++

I first create the `best` array and fill it with a large value. This means that initially no valid target-sum sub-array has been found.

Then I use `left`, `sum`, and `answer` for the sliding window and final result.

The `right` pointer moves through the array one element at a time.

Whenever I add a new value, I update the current window sum.

If the sum becomes larger than `target`, I move `left` forward. Since all values are positive, this always reduces the sum.

When the sum becomes equal to `target`, I calculate the length of the current sub-array.

At that point, I check `best[left - 1]`. This gives me the shortest valid sub-array completely before the current one.

I combine the two lengths and update the minimum answer.

Finally, I update `best[right]` so that future windows can use this result.

### Java

The Java version follows exactly the same algorithm.

The `int[] best` array stores the shortest valid target-sum sub-array for every prefix.

I use `Integer.MAX_VALUE` as the representation of "not found yet".

The sliding window variables have the same meaning as in C++.

The important part is still the relationship:

```text
best[left - 1] + currentLength
```

The first part comes entirely before the current window, so the two ranges cannot overlap.

### JavaScript

In JavaScript, the `best` array is initialized using `Infinity`.

I use the same two-pointer sliding window.

Whenever `sum` becomes greater than `target`, I move `left` forward until the window becomes valid again.

When `sum === target`, I calculate the current length and compare it with the shortest valid previous sub-array.

The algorithm remains `O(n)`.

### TypeScript

The TypeScript solution has the same logic as the JavaScript version, but the array and function parameters are explicitly typed.

The `best` array is a `number[]`.

`Infinity` is used to represent an unavailable previous sub-array.

The sliding window works exactly the same way because the input values are positive.

### Python3

In Python, I use a list called `best`.

I initialize it with `float("inf")` because I want a value larger than any valid answer.

The `left` and `right` pointers create the sliding window.

For every `right` position, I add the current array value to `total`.

While the total is greater than the target, I remove values from the left.

When I reach the target sum, I use the current length together with `best[left - 1]`.

The result is returned as `-1` if no valid pair exists.

### Go

In Go, I create a slice named `best`.

Since Go does not have a built-in infinity value for integers, I use a very large integer value as `INF`.

The rest of the algorithm is unchanged:

* Expand the window.
* Shrink it when the sum is too large.
* Detect a target-sum sub-array.
* Combine it with the best valid previous sub-array.
* Store the prefix minimum.

The six implementations differ mainly in syntax and language-specific types. The actual algorithm remains the same.

## Examples

### Example 1

Input:

```text
arr = [3,2,4,3]
target = 3
```

Valid target-sum sub-arrays:

```text
[3]       [3]
 ^         ^
 index 0   index 3
```

Both sub-arrays have length `1`.

They do not overlap.

```text
1 + 1 = 2
```

Output:

```text
2
```

### Example 2

Input:

```text
arr = [7,3,4,7]
target = 7
```

Possible target-sum sub-arrays include:

```text
[7]
[3,4]
[7]
```

Their lengths are:

```text
1, 2, 1
```

The first and third sub-arrays are non-overlapping:

```text
[7] [3,4] [7]
 ^           ^
```

So:

```text
1 + 1 = 2
```

Output:

```text
2
```

### Example 3

Input:

```text
arr = [4,3,2,6,2,3,4]
target = 6
```

Only one target-sum sub-array can be used.

There are not two non-overlapping valid sub-arrays.

Output:

```text
-1
```

## How to Use / Run Locally

The repository contains the same algorithm in C++, Java, JavaScript, TypeScript, Python3, and Go.

### C++

Save the solution in a file such as:

```text
solution.cpp
```

Compile it with:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it with:

```bash
./solution
```

### Java

Save the solution as:

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

### JavaScript

Save the solution as:

```text
solution.js
```

Run it with:

```bash
node solution.js
```

### TypeScript

Install TypeScript if it is not already installed:

```bash
npm install -g typescript
```

Save the solution as:

```text
solution.ts
```

Compile it:

```bash
tsc solution.ts
```

Then run the generated JavaScript:

```bash
node solution.js
```

### Python3

Save the solution as:

```text
solution.py
```

Run it with:

```bash
python3 solution.py
```

### Go

Save the solution as:

```text
solution.go
```

Run it directly with:

```bash
go run solution.go
```

For local testing, I recommend adding a small driver program for the language you are using and passing the sample inputs from the problem statement.

## Notes & Optimizations

The most important observation is that all values in `arr` are positive.

That is what makes the sliding window approach possible.

If negative numbers were allowed, simply moving `left` whenever the sum became too large would no longer give the same guarantee.

A brute-force approach would check many possible sub-arrays and then many pairs of them. That can become too slow for `n = 10^5`.

A prefix sum with binary search can also be considered because the values are positive, but the sliding window gives a simpler linear-time solution here.

The `best` array is what makes the non-overlapping condition easy to handle.

For a current window `[left ... right]`, I only use information from `best[left - 1]`. This guarantees that the earlier sub-array finishes before the current one starts.

Important edge cases include:

* Only one target-sum sub-array exists.
* Two target-sum sub-arrays exist but overlap.
* Two single-element target-sum sub-arrays exist.
* The valid sub-arrays occur far apart in the array.
* No target-sum sub-array exists at all.
* Multiple valid sub-arrays exist, but only one pair gives the minimum combined length.

The final solution achieves `O(n)` time and `O(n)` extra space, which is suitable for the given constraints.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
