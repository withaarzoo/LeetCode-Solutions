# 3660. Jump Game IX

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

## Problem Summary

In this LeetCode hard array problem, we are given an integer array `nums`.

From any index `i`, we can jump to another index `j` using these rules:

* Move to the right only if `nums[j] < nums[i]`
* Move to the left only if `nums[j] > nums[i]`

For every position in the array, we need to find the maximum value reachable after making any number of valid jumps.

The final output is another array where:

* `answer[i]` = maximum value reachable starting from index `i`

This problem looks like a graph traversal problem at first, but there is a much smarter linear-time solution using array observations, connected segments, and suffix minimum values.

This is a great competitive programming and DSA problem for learning:

* Graph connectivity thinking
* Greedy observations
* Prefix/suffix preprocessing
* Linear scan optimization
* Array segmentation techniques

---

## Constraints

| Constraint                 | Value                    |
| -------------------------- | ------------------------ |
| `1 <= nums.length <= 10^5` | Array size can be large  |
| `1 <= nums[i] <= 10^9`     | Values can be very large |

Because the input size is large, an `O(n^2)` solution will cause TLE. An optimized `O(n)` solution is needed.

---

## Intuition

The first thing I noticed was that jumps are based entirely on comparisons between values.

If `nums[i] > nums[j]` and `i < j`, then:

* I can jump from `i -> j`
* I can also jump back from `j -> i`

That means these positions are connected.

So instead of thinking about individual jumps, I started thinking in terms of connected components inside the array.

Then I realized something important:

If there exists any inversion between two sides of the array, then those sides are still connected.

This turns the problem into finding connected segments and assigning the maximum value inside each segment.

That observation removes the need for DFS, BFS, or graph construction entirely.

---

## Approach

I solve the problem in four main steps.

### Step 1: Build a suffix minimum array

For every index:

* `suffixMin[i]` stores the smallest value from `i` to the end

This helps me quickly check whether a smaller value exists on the right side.

---

### Step 2: Start building segments

I scan the array from left to right.

For every new segment:

* Keep track of the current segment maximum
* Expand the segment while:

`currentMax > suffixMin[nextPosition]`

If this condition is true, then an inversion still exists across the boundary, so the segment is still connected.

---

### Step 3: Close the segment

Once the condition becomes false:

* The current segment is fully separated from the remaining array
* Every index inside this segment can reach the same maximum value

---

### Step 4: Fill the answer

Assign the segment maximum to every index inside that segment.

Repeat until the entire array is processed.

---

## Data Structures Used

| Data Structure       | Purpose                                      |
| -------------------- | -------------------------------------------- |
| Array                | Store the original input                     |
| Suffix Minimum Array | Quickly find the smallest value to the right |
| Answer Array         | Store final maximum reachable values         |

No stacks, queues, graphs, or recursion are needed.

---

## Operations & Behavior Summary

Here is the overall flow of the algorithm in simple terms:

1. Create a suffix minimum array
2. Start scanning from the left
3. Keep expanding the current segment while inversions still exist
4. Track the maximum value inside the segment
5. Once the segment becomes isolated, stop expanding
6. Fill all positions in that segment with the segment maximum
7. Continue with the next segment

The algorithm processes the array only once after preprocessing.

---

## Complexity

| Type             | Complexity | Explanation                                               |
| ---------------- | ---------- | --------------------------------------------------------- |
| Time Complexity  | `O(n)`     | Each index is processed a constant number of times        |
| Space Complexity | `O(n)`     | Extra arrays are used for suffix minimum and final answer |

Where:

* `n` = size of the input array `nums`

This is the optimal solution for the problem.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> maxValue(vector<int>& nums) {
        int n = (int)nums.size();

        // suffixMin[i] = smallest value in nums[i...n-1]
        // I keep one extra cell at the end so the last segment can stop cleanly.
        vector<int> suffixMin(n + 1, INT_MAX);
        for (int i = n - 1; i >= 0; --i) {
            suffixMin[i] = min(nums[i], suffixMin[i + 1]);
        }

        vector<int> ans(n);
        int l = 0;

        // I build one connected component at a time.
        while (l < n) {
            int r = l;
            int componentMax = nums[l];

            // I keep extending this segment while some value on the left
            // is bigger than some value on the right, which means an inversion
            // crosses the cut and the two parts are still connected.
            while (r + 1 < n && componentMax > suffixMin[r + 1]) {
                ++r;
                componentMax = max(componentMax, nums[r]);
            }

            // Every index in this connected component can reach the same maximum.
            for (int i = l; i <= r; ++i) {
                ans[i] = componentMax;
            }

            // Move to the next segment.
            l = r + 1;
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int[] maxValue(int[] nums) {
        int n = nums.length;

        // suffixMin[i] = smallest value in nums[i...n-1]
        // I keep one extra slot so the boundary case is easy to handle.
        int[] suffixMin = new int[n + 1];
        suffixMin[n] = Integer.MAX_VALUE;
        for (int i = n - 1; i >= 0; i--) {
            suffixMin[i] = Math.min(nums[i], suffixMin[i + 1]);
        }

        int[] ans = new int[n];
        int l = 0;

        // I process one connected component at a time.
        while (l < n) {
            int r = l;
            int componentMax = nums[l];

            // I extend the current segment while an inversion crosses the next cut.
            while (r + 1 < n && componentMax > suffixMin[r + 1]) {
                r++;
                componentMax = Math.max(componentMax, nums[r]);
            }

            // All positions in this component share the same maximum reachable value.
            for (int i = l; i <= r; i++) {
                ans[i] = componentMax;
            }

            // Jump to the next untouched segment.
            l = r + 1;
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number[]}
 */
var maxValue = function(nums) {
    const n = nums.length;

    // suffixMin[i] = smallest value in nums[i...n-1]
    // I add one extra slot so the last segment stops naturally.
    const suffixMin = new Array(n + 1);
    suffixMin[n] = Infinity;
    for (let i = n - 1; i >= 0; i--) {
        suffixMin[i] = Math.min(nums[i], suffixMin[i + 1]);
    }

    const ans = new Array(n);
    let l = 0;

    // I split the array into connected components.
    while (l < n) {
        let r = l;
        let componentMax = nums[l];

        // I keep expanding while the current segment still has an inversion
        // crossing the next boundary.
        while (r + 1 < n && componentMax > suffixMin[r + 1]) {
            r++;
            componentMax = Math.max(componentMax, nums[r]);
        }

        // Every index inside this component gets the same best reachable value.
        for (let i = l; i <= r; i++) {
            ans[i] = componentMax;
        }

        // Move to the next component.
        l = r + 1;
    }

    return ans;
};
```

### Python3

```python
from typing import List

class Solution:
    def maxValue(self, nums: List[int]) -> List[int]:
        n = len(nums)

        # suffixMin[i] = smallest value in nums[i...n-1]
        # I keep one extra cell so the last segment can close cleanly.
        suffixMin = [0] * (n + 1)
        suffixMin[n] = float("inf")
        for i in range(n - 1, -1, -1):
            suffixMin[i] = min(nums[i], suffixMin[i + 1])

        ans = [0] * n
        l = 0

        # I walk through the array and build one connected segment at a time.
        while l < n:
            r = l
            component_max = nums[l]

            # I extend the segment while some inversion still crosses the next cut.
            while r + 1 < n and component_max > suffixMin[r + 1]:
                r += 1
                component_max = max(component_max, nums[r])

            # Every index in this segment can reach this segment maximum.
            for i in range(l, r + 1):
                ans[i] = component_max

            # Continue with the next segment.
            l = r + 1

        return ans
```

### Go

```go
func maxValue(nums []int) []int {
 n := len(nums)

 // suffixMin[i] = smallest value in nums[i...n-1]
 // I keep one extra slot so the boundary check is simple.
 suffixMin := make([]int, n+1)
 suffixMin[n] = int(^uint(0) >> 1) // MaxInt for the current Go architecture

 for i := n - 1; i >= 0; i-- {
  if nums[i] < suffixMin[i+1] {
   suffixMin[i] = nums[i]
  } else {
   suffixMin[i] = suffixMin[i+1]
  }
 }

 ans := make([]int, n)
 l := 0

 // I process each connected component as one contiguous block.
 for l < n {
  r := l
  componentMax := nums[l]

  // I keep growing the block while an inversion crosses the next cut.
  for r+1 < n && componentMax > suffixMin[r+1] {
   r++
   if nums[r] > componentMax {
    componentMax = nums[r]
   }
  }

  // Every index in this block gets the same answer.
  for i := l; i <= r; i++ {
   ans[i] = componentMax
  }

  // Move to the next block.
  l = r + 1
 }

 return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical in all five languages. Only syntax changes.

### Building the suffix minimum array

I first preprocess the array from right to left.

For every position:

* Store the minimum value seen so far

This allows me to instantly know whether a smaller value exists somewhere on the right side.

Without this preprocessing step, I would need nested loops, which would be too slow.

---

### Starting a new connected segment

I begin from the leftmost unprocessed index.

For the current segment:

* Keep a running maximum
* Expand the segment while an inversion still exists across the boundary

The condition:

`currentMax > suffixMin[next]`

means there is still a valid connection between the current segment and the next part of the array.

So I must continue expanding.

---

### Updating the segment maximum

Whenever I add a new value into the segment:

* Update the maximum value

This matters because the segment maximum determines the final answer for all indices inside that component.

If I forget to update it, the final result becomes incorrect.

---

### Detecting segment boundaries

The segment ends when:

`currentMax <= suffixMin[next]`

At that point:

* No inversion crosses the boundary
* No jumps can connect the two sides anymore
* The current connected component is complete

This is the key observation that makes the linear solution possible.

---

### Filling the answer array

Once the segment is complete:

* Every index inside that segment receives the same answer
* That answer is simply the segment maximum

Because all positions inside the same component can eventually reach one another.

---

### Why graph traversal is avoided

A brute-force graph approach would be extremely expensive.

If every pair is checked:

* Time complexity becomes `O(n^2)`

That will fail for large constraints.

The optimized approach avoids building edges completely and instead uses mathematical observations about inversions and connectivity.

---

## Examples

### Example 1

Input:

```text
nums = [2,1,3]
```

Output:

```text
[2,2,3]
```

Explanation:

* Index 0 already has value 2 and cannot reach 3
* Index 1 can jump to index 0
* Index 2 already contains the maximum value

Final answer:

```text
[2,2,3]
```

---

### Example 2

Input:

```text
nums = [2,3,1]
```

Output:

```text
[3,3,3]
```

Explanation:

* Index 0 can jump to index 2
* Index 2 can jump back to index 1
* All positions become connected

Maximum reachable value becomes 3 for every index.

---

### Example 3

Input:

```text
nums = [5,4,3,2,1]
```

Output:

```text
[5,5,5,5,5]
```

Explanation:

Every position can eventually reach the largest value because the entire array forms one connected component.

---

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

Run using Node.js:

```bash
node solution.js
```

---

### Python3

Run:

```bash
python solution.py
```

---

### Go

Run:

```bash
go run solution.go
```

---

## Notes & Optimizations

* The biggest optimization is avoiding graph construction entirely.
* The suffix minimum preprocessing step is what enables the `O(n)` solution.
* This problem is a good example of transforming graph connectivity into array segmentation.
* A brute-force simulation of jumps will not work for large inputs.
* DFS and BFS are unnecessary once the connectivity observation is understood.
* The solution works efficiently even for the maximum constraint size of `10^5`.

Possible alternative approaches:

* Union Find / DSU
* Graph traversal
* Monotonic structures

But the segment-based linear scan is simpler and faster.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
