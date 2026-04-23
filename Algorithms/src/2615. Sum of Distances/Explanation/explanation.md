# 2615. Sum of Distances

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

Given a 0-indexed integer array `nums`, for every index `i` we need to calculate the sum of absolute differences between `i` and all other indices `j` such that `nums[i] == nums[j]`.

In other words, for each position `i`, we want:

`arr[i] = sum of |i - j| for all j where nums[j] == nums[i] and j != i`

If no other index has the same value, then `arr[i] = 0`.

## Constraints

* `1 <= nums.length <= 10^5`
* `0 <= nums[i] <= 10^9`

These constraints make an `O(n^2)` solution too slow, so we need an efficient approach.

## Intuition

I first thought about comparing every index with every other index having the same value. That would work, but it would be too slow for large input sizes.

Then I noticed that the problem only depends on indices of equal values. So instead of looking at the whole array again and again, I can group all positions for each number.

Once I have the indices for one number in sorted order, I can compute distances using prefix sums. This lets me calculate the answer for each index without checking every pair separately.

## Approach

1. Group all indices by their value.
2. For each group of equal values:

   * Traverse the indices from left to right.
   * Keep a running prefix sum of the indices already processed.
   * Also compute the total sum of all indices in the group.
3. For each index `idx[i]` in the group:

   * The contribution from the left side is:
     `idx[i] * i - prefixSum`
   * The contribution from the right side is:
     `(totalSum - prefixSum - idx[i]) - idx[i] * (k - i - 1)`
   * Add both parts to get the final answer for that index.
4. Store the answer in the result array.

This works because the indices in each group are already in increasing order, and prefix sums let us compute distance sums in constant time per position.

## Data Structures Used

* `HashMap / unordered_map / Map / dict`

  * Stores all indices for each distinct value.
* `List / vector / array`

  * Stores the positions of each value in increasing order.
* Result array

  * Stores the final answer for every index.

## Operations & Behavior Summary

For each value group:

* `prefixSum` stores the sum of indices before the current one.
* `totalSum` stores the sum of all indices in that group.
* Left-side distance is computed from already processed indices.
* Right-side distance is computed from the remaining indices.
* The final answer is the sum of left and right contributions.

## Complexity

* **Time Complexity:** `O(n)`

  * Every index is added once to a group.
  * Every index is processed once while computing answers.

* **Space Complexity:** `O(n)`

  * We store indices grouped by value.
  * The result array also uses `O(n)` space.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<long long> distance(vector<int>& nums) {
        int n = nums.size();
        unordered_map<int, vector<int>> groups;

        // Group all indices by value.
        for (int i = 0; i < n; i++) {
            groups[nums[i]].push_back(i);
        }

        vector<long long> ans(n, 0);

        // Process every group separately.
        for (auto &entry : groups) {
            vector<int> &idx = entry.second;
            int k = (int)idx.size();

            long long totalSum = 0;
            for (int x : idx) {
                totalSum += x;
            }

            long long prefixSum = 0;
            for (int i = 0; i < k; i++) {
                long long cur = idx[i];

                // Sum of distances to indices on the left.
                long long left = cur * i - prefixSum;

                // Sum of distances to indices on the right.
                long long right = (totalSum - prefixSum - cur) - cur * (k - i - 1);

                ans[cur] = left + right;
                prefixSum += cur;
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public long[] distance(int[] nums) {
        int n = nums.length;
        Map<Integer, List<Integer>> groups = new HashMap<>();

        // Group all indices by value.
        for (int i = 0; i < n; i++) {
            groups.computeIfAbsent(nums[i], k -> new ArrayList<>()).add(i);
        }

        long[] ans = new long[n];

        // Process every group separately.
        for (List<Integer> idx : groups.values()) {
            int k = idx.size();

            long totalSum = 0;
            for (int x : idx) {
                totalSum += x;
            }

            long prefixSum = 0;
            for (int i = 0; i < k; i++) {
                long cur = idx.get(i);

                // Sum of distances to indices on the left.
                long left = cur * i - prefixSum;

                // Sum of distances to indices on the right.
                long right = (totalSum - prefixSum - cur) - cur * (k - i - 1);

                ans[(int) cur] = left + right;
                prefixSum += cur;
            }
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
var distance = function(nums) {
    const n = nums.length;
    const groups = new Map();

    // Group all indices by value.
    for (let i = 0; i < n; i++) {
        if (!groups.has(nums[i])) {
            groups.set(nums[i], []);
        }
        groups.get(nums[i]).push(i);
    }

    const ans = new Array(n).fill(0);

    // Process every group separately.
    for (const idx of groups.values()) {
        const k = idx.length;

        let totalSum = 0;
        for (const x of idx) {
            totalSum += x;
        }

        let prefixSum = 0;
        for (let i = 0; i < k; i++) {
            const cur = idx[i];

            // Sum of distances to indices on the left.
            const left = cur * i - prefixSum;

            // Sum of distances to indices on the right.
            const right = (totalSum - prefixSum - cur) - cur * (k - i - 1);

            ans[cur] = left + right;
            prefixSum += cur;
        }
    }

    return ans;
};
```

### Python3

```python
from collections import defaultdict
from typing import List

class Solution:
    def distance(self, nums: List[int]) -> List[int]:
        n = len(nums)
        groups = defaultdict(list)

        # Group all indices by value.
        for i, num in enumerate(nums):
            groups[num].append(i)

        ans = [0] * n

        # Process every group separately.
        for idx in groups.values():
            k = len(idx)
            total_sum = sum(idx)
            prefix_sum = 0

            for i in range(k):
                cur = idx[i]

                # Sum of distances to indices on the left.
                left = cur * i - prefix_sum

                # Sum of distances to indices on the right.
                right = (total_sum - prefix_sum - cur) - cur * (k - i - 1)

                ans[cur] = left + right
                prefix_sum += cur

        return ans
```

### Go

```go
func distance(nums []int) []int64 {
    n := len(nums)
    groups := make(map[int][]int)

    // Group all indices by value.
    for i := 0; i < n; i++ {
        groups[nums[i]] = append(groups[nums[i]], i)
    }

    ans := make([]int64, n)

    // Process every group separately.
    for _, idx := range groups {
        k := len(idx)

        var totalSum int64 = 0
        for _, x := range idx {
            totalSum += int64(x)
        }

        var prefixSum int64 = 0
        for i := 0; i < k; i++ {
            cur := int64(idx[i])

            // Sum of distances to indices on the left.
            left := cur*int64(i) - prefixSum

            // Sum of distances to indices on the right.
            right := (totalSum - prefixSum - cur) - cur*int64(k-i-1)

            ans[idx[i]] = left + right
            prefixSum += cur
        }
    }

    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### 1. Group indices by value

We first store all positions of the same number together.

For example, if:

```text
nums = [1, 3, 1, 1, 2]
```

Then the groups become:

```text
1 -> [0, 2, 3]
3 -> [1]
2 -> [4]
```

This is important because the answer for one index only depends on indices with the same value.

### 2. Compute the total sum of indices in a group

For the group `[0, 2, 3]`, the total sum is:

```text
0 + 2 + 3 = 5
```

This helps us compute right-side distances efficiently.

### 3. Walk through the group from left to right

Suppose we are at index `2` in the group `[0, 2, 3]`.

* `prefixSum` stores the sum of earlier indices.
* `i` tells how many indices are on the left.

For `2`:

* Left side: `|2 - 0| = 2`
* Right side: `|2 - 3| = 1`
* Total: `3`

Instead of calculating each absolute value one by one, we use formulas.

### 4. Left-side formula

For a current index `cur`, if there are `i` earlier indices, then:

```text
left = cur * i - prefixSum
```

Why this works:

* `cur * i` means we imagine all left positions are equal to `cur`.
* Then we subtract their real sum, which gives total distance.

### 5. Right-side formula

For the remaining indices on the right:

```text
right = (totalSum - prefixSum - cur) - cur * (k - i - 1)
```

Here:

* `totalSum - prefixSum - cur` gives the sum of the indices on the right.
* `cur * (k - i - 1)` is the same current index repeated for all those right positions.
* Their difference gives total distance on the right.

### 6. Combine both sides

```text
answer = left + right
```

This gives the final distance sum for the current index.

### 7. Store the result

We place the computed value into the result array at the current position.

That way, after processing all groups, we get the full answer for every index.

## Examples

### Example 1

Input:

```text
nums = [1, 3, 1, 1, 2]
```

Output:

```text
[5, 0, 3, 4, 0]
```

Explanation:

* For index `0`, same values are at `2` and `3`, so distance sum is `|0-2| + |0-3| = 5`
* For index `1`, no other `3` exists, so answer is `0`
* For index `2`, same values are at `0` and `3`, so distance sum is `|2-0| + |2-3| = 3`
* For index `3`, same values are at `0` and `2`, so distance sum is `|3-0| + |3-2| = 4`
* For index `4`, no other `2` exists, so answer is `0`

### Example 2

Input:

```text
nums = [0, 5, 3]
```

Output:

```text
[0, 0, 0]
```

Explanation:

* Every number is unique.
* So there are no equal elements to compare with.
* Therefore every answer is `0`.

## How to use / Run locally

### C++

```bash
g++ -std=c++17 -O2 -o main main.cpp
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

* Use `long long` in C++, `long` in Java, and `int64` in Go because the sum of distances can be large.
* The grouping step is the most important optimization.
* Prefix sums remove the need for nested loops.
* This solution is much faster than brute force and fits the problem constraints comfortably.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
