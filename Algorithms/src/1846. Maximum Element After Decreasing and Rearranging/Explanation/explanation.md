# 1846. Maximum Element After Decreasing and Rearranging

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

This problem asks us to take an array of positive integers and change it so it follows two rules:

1. The first element must be `1`.
2. The difference between any two adjacent elements must be at most `1`.

We are allowed to do two things:

* decrease any number to a smaller positive integer
* rearrange the array in any order

The goal is to return the maximum possible value of any element after making the array valid.

In simple words, I need to reshape the array into the best possible increasing sequence, while staying inside the rules.

## Constraints

| Constraint    |                     Value |
| ------------- | ------------------------: |
| Array length  | `1 <= arr.length <= 10^5` |
| Element value |     `1 <= arr[i] <= 10^9` |

## Intuition

The main clue is that rearranging is allowed. That means the original order does not matter at all.

I realized that if I sort the array first, I can build the final array from left to right in the safest way. The smallest values should come first, because large values can always be reduced, but small values cannot be increased.

The first element must be `1`, so after sorting, I force the first value to become `1`. Then for every next position, I only keep the value if it is not too large compared to the previous value. If it is too large, I reduce it to `previous + 1`.

This greedy idea works because I always keep each number as large as possible without breaking the rule. That gives me the highest possible final answer.

## Approach

I start by sorting the array in increasing order.

Then I set the first element to `1`, because that is required.

After that, I move from left to right.

For each position `i`, I compare the current value with `arr[i - 1] + 1`.

* If the current value is already small enough, I keep it.
* If it is too large, I reduce it to `arr[i - 1] + 1`.

This keeps the array valid at every step.

At the end, the last element will be the maximum possible value in the rearranged array. I return that value.

## Data Structures Used

* **Array**: I use the given array directly and update it in place.
* **Sorting**: I sort the array first so the greedy choice becomes easy and correct.

I do not need any extra complex data structure here. The solution stays simple and efficient.

## Operations & Behavior Summary

1. Sort the array.
2. Make the first element `1`.
3. Scan the rest of the array from left to right.
4. At each step, keep the current value no bigger than `previous + 1`.
5. Return the last element.

That is the whole flow. It is a greedy build-up from the smallest values to the largest valid answer.

## Complexity

| Complexity       |        Value | Explanation                                                             |
| ---------------- | -----------: | ----------------------------------------------------------------------- |
| Time Complexity  | `O(n log n)` | Sorting takes `O(n log n)`. The single pass after sorting takes `O(n)`. |
| Space Complexity |       `O(1)` | I update the array in place and do not use extra data structures.       |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maximumElementAfterDecrementingAndRearranging(vector<int>& arr) {

        // Sort the array so we process numbers from smallest to largest.
        sort(arr.begin(), arr.end());

        // The first element must always become 1.
        arr[0] = 1;

        // Build the largest valid sequence.
        for (int i = 1; i < arr.size(); i++) {

            // The current value cannot be larger than previous + 1.
            // If it already satisfies this, it stays unchanged.
            arr[i] = min(arr[i], arr[i - 1] + 1);
        }

        // The last element is the maximum possible value.
        return arr.back();
    }
};
```

### Java

```java
class Solution {
    public int maximumElementAfterDecrementingAndRearranging(int[] arr) {

        // Sort the array in ascending order.
        Arrays.sort(arr);

        // The first element must be 1.
        arr[0] = 1;

        // Make every element as large as possible.
        for (int i = 1; i < arr.length; i++) {

            // Limit the current value to previous + 1.
            arr[i] = Math.min(arr[i], arr[i - 1] + 1);
        }

        // The last element is the answer.
        return arr[arr.length - 1];
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} arr
 * @return {number}
 */
var maximumElementAfterDecrementingAndRearranging = function(arr) {

    // Sort the array in increasing order.
    arr.sort((a, b) => a - b);

    // The first element must become 1.
    arr[0] = 1;

    // Build the largest valid sequence.
    for (let i = 1; i < arr.length; i++) {

        // The current value cannot exceed previous + 1.
        arr[i] = Math.min(arr[i], arr[i - 1] + 1);
    }

    // Return the largest value.
    return arr[arr.length - 1];
};
```

### Python3

```python
class Solution:
    def maximumElementAfterDecrementingAndRearranging(self, arr: List[int]) -> int:

        # Sort the array so smaller values come first.
        arr.sort()

        # The first element must always be 1.
        arr[0] = 1

        # Build the largest valid sequence.
        for i in range(1, len(arr)):

            # The current value cannot be greater than previous + 1.
            arr[i] = min(arr[i], arr[i - 1] + 1)

        # The last element is the maximum possible value.
        return arr[-1]
```

### Go

```go
func maximumElementAfterDecrementingAndRearranging(arr []int) int {

    // Sort the array in ascending order.
    sort.Ints(arr)

    // The first element must become 1.
    arr[0] = 1

    // Build the largest valid sequence.
    for i := 1; i < len(arr); i++ {

        // The current value cannot exceed previous + 1.
        if arr[i] > arr[i-1]+1 {
            arr[i] = arr[i-1] + 1
        }
    }

    // The last element is the answer.
    return arr[len(arr)-1]
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five languages.

First, I sort the array. This is important because I want to handle smaller values first. If I process a large value too early, I may waste the chance to build a better sequence.

Next, I set the first element to `1`. That is not optional. The problem demands it, and I can always decrease a number to `1` if needed.

Then I look at each next element. I already know the previous element is valid, so the current one only needs to be checked against `previous + 1`.

If the current number is too large, I reduce it. I never increase anything, because the problem does not allow that. I also never reduce more than necessary, because I want the final answer to be as large as possible.

That is why the formula is always:

`arr[i] = min(arr[i], arr[i - 1] + 1)`

By doing this for every element, I build the best possible valid array. The last value becomes the answer because the array is sorted and each step only increases by at most `1`.

## Examples

### Example 1

**Input:** `arr = [2, 2, 1, 2, 1]`
**Output:** `2`

**Trace:**

* After sorting: `[1, 1, 2, 2, 2]`
* Set first element to `1`
* Build the sequence from left to right
* Final valid array stays close to `[1, 1, 2, 2, 2]`
* Largest element is `2`

### Example 2

**Input:** `arr = [100, 1, 1000]`
**Output:** `3`

**Trace:**

* After sorting: `[1, 100, 1000]`
* First element becomes `1`
* Second element becomes `2`
* Third element becomes `3`
* Final array is `[1, 2, 3]`
* Largest element is `3`

### Example 3

**Input:** `arr = [1, 2, 3, 4, 5]`
**Output:** `5`

**Trace:**

* The array is already sorted
* The first element is already `1`
* Every next element already fits the rule
* No changes are needed
* Largest element is `5`

## How to Use / Run Locally

### C++

```bash
g++ -std=c++17 main.cpp -o main
./main
```

### Java

```bash
javac Main.java
java Main
```

### JavaScript

```bash
node main.js
```

### Python3

```bash
python3 main.py
```

### Go

```bash
go run main.go
```

## Notes & Optimizations

A sorting-based greedy solution is the cleanest approach here. It is easy to reason about and fast enough for the given limits.

One small optimization is to do the work in place, which keeps memory usage low.

A different approach without sorting is possible only in special cases, but sorting is the most reliable and readable method for this problem.

The key idea to remember is simple: keep each number as large as possible, but never let it become more than `1` bigger than the previous number.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
