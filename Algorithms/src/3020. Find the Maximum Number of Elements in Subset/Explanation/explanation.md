# Find the Maximum Number of Elements in Subset

A beginner-friendly solution and explanation for **LeetCode 3020 - Find the Maximum Number of Elements in Subset**. This repository explains the idea behind the algorithm, walks through the solution step by step, and provides implementations in **C++, Java, JavaScript, Python3, and Go**.

The goal is to understand both the reasoning and the implementation, making it useful for coding interviews, competitive programming, and DSA practice.

---

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
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

In this problem, we are given an array of positive integers.

Our task is to choose the largest possible subset that can be rearranged into a symmetric sequence of powers, following this pattern:

```
x, x², x⁴, ..., center, ..., x⁴, x², x
```

Every value except the center must appear exactly twice, while the center appears only once.

The challenge is to determine the maximum number of elements that can be included in such a subset.

This problem combines **hashing**, **frequency counting**, and **greedy simulation**, making it a great practice problem for competitive programming and coding interviews.

---

## Constraints

| Constraint            | Value                 |
| --------------------- | --------------------- |
| Number of elements    | 2 ≤ nums.length ≤ 10⁵ |
| Value of each element | 1 ≤ nums[i] ≤ 10⁹     |

---

## Intuition

The first thing I noticed was that every value before the middle of the sequence appears twice.

That means I cannot continue extending the sequence unless I have at least two copies of the current number.

The only exception is the final value, which becomes the center and only needs one occurrence.

I also noticed that the number **1** behaves differently because:

```
1² = 1
```

Unlike every other value, the chain never changes. Because of that, I simply need the largest odd number of ones.

For every value greater than one, I repeatedly square it and check whether the next value exists in the frequency map.

Since squaring grows extremely fast, every chain is very short.

---

## Approach

1. Count the frequency of every number.
2. Handle the number `1` separately because it keeps squaring to itself.
3. Try every distinct value greater than `1` as the starting point.
4. Continue squaring the current value.
5. If the current value appears at least twice, use both copies and continue.
6. If only one copy exists, make it the center and stop.
7. If the chain ends without finding a center, remove one element because the sequence length must be odd.
8. Keep track of the maximum valid subset length.

---

## Data Structures Used

| Data Structure           | Purpose                                                          |
| ------------------------ | ---------------------------------------------------------------- |
| Hash Map / Frequency Map | Stores how many times every number appears.                      |
| Integer Variables        | Track the current value, current chain length, and final answer. |

A hash map allows constant-time frequency lookup, which is exactly what this problem needs.

---

## Operations & Behavior Summary

The algorithm works in four simple stages.

1. Build a frequency map from the input array.
2. Compute the best answer using only the value `1`.
3. For every other unique number:

   * Start a new chain.
   * Keep squaring the current value.
   * Continue only if at least two copies exist.
   * Stop once a single occurrence becomes the center.
4. Compare every valid chain and return the longest one.

Since squaring increases numbers very quickly, each chain contains only a handful of values.

---

## Complexity

| Complexity       | Value                 | Explanation                                                                                                                                               |
| ---------------- | --------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | O(n + k × log(log M)) | `n` is the array size, `k` is the number of distinct values, and `M` is the maximum value. Every chain grows extremely fast because of repeated squaring. |
| Space Complexity | O(k)                  | A hash map stores the frequency of every distinct number.                                                                                                 |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maximumLength(vector<int>& nums) {
        // Store frequency of every number
        unordered_map<long long, int> freq;
        for (int x : nums) freq[x]++;

        int ans = 1;

        // Handle value 1 separately because 1^2 = 1
        if (freq.count(1)) {
            int cnt = freq[1];

            // We can only take an odd number of ones
            ans = max(ans, (cnt % 2 == 1) ? cnt : cnt - 1);
        }

        // Try every distinct starting value (>1)
        for (auto &[start, cnt] : freq) {
            if (start == 1) continue;

            long long cur = start;
            int len = 0;

            while (freq.count(cur)) {
                // If at least two copies exist, use both
                if (freq[cur] >= 2) {
                    len += 2;

                    // Move to the next squared value
                    cur = cur * cur;
                }
                // Only one copy exists, so it becomes the center
                else {
                    len++;
                    break;
                }
            }

            // If the length is even, we never found a center
            if (len % 2 == 0) len--;

            ans = max(ans, len);
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int maximumLength(int[] nums) {
        // Store frequency of every number
        HashMap<Long, Integer> freq = new HashMap<>();

        for (int x : nums) {
            freq.put((long) x, freq.getOrDefault((long) x, 0) + 1);
        }

        int ans = 1;

        // Handle value 1 separately
        if (freq.containsKey(1L)) {
            int cnt = freq.get(1L);

            // Only odd count of ones is valid
            ans = Math.max(ans, (cnt % 2 == 1) ? cnt : cnt - 1);
        }

        // Try every distinct starting value
        for (long start : freq.keySet()) {
            if (start == 1L) continue;

            long cur = start;
            int len = 0;

            while (freq.containsKey(cur)) {
                // Use two copies if possible
                if (freq.get(cur) >= 2) {
                    len += 2;

                    // Move to the squared value
                    cur = cur * cur;
                } else {
                    // Single copy becomes the center
                    len++;
                    break;
                }
            }

            // No center found
            if ((len & 1) == 0) len--;

            ans = Math.max(ans, len);
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var maximumLength = function(nums) {
    // Store frequency of every number
    const freq = new Map();

    for (const x of nums) {
        freq.set(x, (freq.get(x) || 0) + 1);
    }

    let ans = 1;

    // Handle value 1 separately
    if (freq.has(1)) {
        const cnt = freq.get(1);

        // Only odd count of ones is valid
        ans = Math.max(ans, cnt % 2 ? cnt : cnt - 1);
    }

    // Try every distinct starting value
    for (const [start] of freq) {
        if (start === 1) continue;

        let cur = start;
        let len = 0;

        while (freq.has(cur)) {
            // Use two copies if available
            if (freq.get(cur) >= 2) {
                len += 2;

                // Move to the squared value
                cur = cur * cur;
            } else {
                // Single copy becomes the center
                len++;
                break;
            }
        }

        // No center found
        if (len % 2 === 0) len--;

        ans = Math.max(ans, len);
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def maximumLength(self, nums: List[int]) -> int:
        # Store frequency of every number
        from collections import Counter

        freq = Counter(nums)
        ans = 1

        # Handle value 1 separately
        if 1 in freq:
            cnt = freq[1]

            # Only odd count of ones is valid
            ans = max(ans, cnt if cnt % 2 else cnt - 1)

        # Try every distinct starting value
        for start in freq:
            if start == 1:
                continue

            cur = start
            length = 0

            while cur in freq:
                # Use two copies if available
                if freq[cur] >= 2:
                    length += 2

                    # Move to the squared value
                    cur *= cur
                else:
                    # Single copy becomes the center
                    length += 1
                    break

            # No center found
            if length % 2 == 0:
                length -= 1

            ans = max(ans, length)

        return ans
```

### Go

```go
func maximumLength(nums []int) int {
 // Store frequency of every number
 freq := make(map[int64]int)

 for _, x := range nums {
  freq[int64(x)]++
 }

 ans := 1

 // Handle value 1 separately
 if cnt, ok := freq[1]; ok {
  // Only odd count of ones is valid
  if cnt%2 == 1 {
   if cnt > ans {
    ans = cnt
   }
  } else {
   if cnt-1 > ans {
    ans = cnt - 1
   }
  }
 }

 // Try every distinct starting value
 for start := range freq {
  if start == 1 {
   continue
  }

  cur := start
  length := 0

  for {
   cnt, ok := freq[cur]
   if !ok {
    break
   }

   // Use two copies if available
   if cnt >= 2 {
    length += 2

    // Move to the squared value
    cur = cur * cur
   } else {
    // Single copy becomes the center
    length++
    break
   }
  }

  // No center found
  if length%2 == 0 {
   length--
  }

  if length > ans {
   ans = length
  }
 }

 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical across all five implementations, so learning one version makes the others easy to understand.

### Step 1: Count frequencies

The algorithm begins by counting how many times every number appears.

Without these counts, there is no way to know whether a value can appear on both sides of the sequence.

---

### Step 2: Handle the special case of `1`

The number `1` is unique because its square is still `1`.

A valid sequence always has exactly one middle element, so only an odd number of ones can be used.

If there are an even number of ones, one of them must be discarded.

---

### Step 3: Try every starting value

Each distinct number greater than `1` is treated as a possible starting point.

The algorithm builds the longest possible chain beginning from that value.

---

### Step 4: Extend the chain

While the current value exists:

* If at least two copies are available, both are used.
* The current value is squared.
* The algorithm moves to the next level.

This mirrors the required symmetric structure of the final subset.

---

### Step 5: Choose the center

Eventually, one of two things happens.

If exactly one copy of the current value exists, that value becomes the center.

The chain is now complete.

---

### Step 6: Handle incomplete chains

Sometimes the next squared value does not exist.

In that case, every selected value was paired, meaning the total length is even.

Since every valid answer must have a unique center, one element is removed.

---

### Step 7: Update the answer

After finishing one chain, compare its length with the best answer found so far.

Repeat the process for every distinct starting value.

Finally, return the largest valid subset length.

---

## Examples

### Example 1

**Input**

```
nums = [5,4,1,2,2]
```

**Output**

```
3
```

**Trace**

* Start with `2`
* Square of `2` is `4`
* One copy of `4` becomes the center
* Valid subset:

```
[2,4,2]
```

Length = **3**

---

### Example 2

**Input**

```
nums = [1,3,2,4]
```

**Output**

```
1
```

**Trace**

No value appears twice.

Only a single element can be chosen.

---

### Example 3

**Input**

```
nums = [1,1,1,1,1]
```

**Output**

```
5
```

**Trace**

Since every value is `1`, the entire subset already satisfies the required pattern.

The count is odd, so all five elements are used.

---

## How to Use / Run Locally

Clone the repository:

```bash
git clone <repository-url>
```

Move into the project directory:

```bash
cd <repository-folder>
```

### C++

Compile:

```bash
g++ solution.cpp -o solution
```

Run:

```bash
./solution
```

---

### Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

---

### JavaScript

Run:

```bash
node solution.js
```

---

### Python3

Run:

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Run:

```bash
go run solution.go
```

---

## Notes & Optimizations

* The number `1` must always be handled separately.
* A frequency map makes lookups constant time.
* Repeated squaring grows extremely fast, so every chain remains very short.
* This approach is much faster than trying to build every possible subset.
* Using `long long` in C++ and `long` in Java helps prevent overflow while computing squares.
* The algorithm is already close to the best possible solution for the given constraints.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
