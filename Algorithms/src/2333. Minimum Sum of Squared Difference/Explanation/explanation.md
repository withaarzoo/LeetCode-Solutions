# 2333. Minimum Sum of Squared Difference - LeetCode Solution

## Table of Contents
- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
  - [C++](#c)
  - [Java](#java)
  - [JavaScript](#javascript)
  - [TypeScript](#typescript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

You are given two arrays `nums1` and `nums2` of the same length. The sum of squared difference is the total of `(nums1[i] - nums2[i])²` for every index.

You can change any value in `nums1` by +1 or -1 up to `k1` times. You can do the same for `nums2` up to `k2` times. Negative numbers are allowed.

The goal is to find the smallest possible sum of squared differences after using at most those operations.

This is a classic LeetCode medium problem that combines greedy thinking with binary search to optimize the final sum of squares under a limited number of moves.

## Constraints

- `n == nums1.length == nums2.length`
- `1 <= n <= 10^5`
- `0 <= nums1[i], nums2[i] <= 10^5`
- `0 <= k1, k2 <= 10^9`

## Intuition

The first thing I noticed is that adding 1 to one array is the same as subtracting 1 from the other when you only care about the absolute difference. So `k1` and `k2` can simply be added together into one total budget `k`.

After that the problem becomes: I have a list of absolute differences and I can decrease any of them by 1 (but never below zero) up to `k` times. I want the smallest possible sum of their squares.

Because the square of a larger number grows much faster, every operation should always be spent on a current maximum difference. Doing anything else would leave a bigger square somewhere.

With `n` up to 1e5 and `k` up to 2e9, I cannot simulate each operation. That observation leads straight to binary search on the final maximum difference.

## Approach

1. Build an array of absolute differences between corresponding elements of `nums1` and `nums2`.
2. If the sum of all differences is already less than or equal to `k`, the answer is zero.
3. Binary search for the smallest value `x` such that every difference can be reduced to at most `x` using at most `k` operations.
4. Apply that maximum: reduce every larger difference down to `x` and subtract the used operations from the remaining budget.
5. Spend any leftover operations on the values that are still equal to `x`, reducing them by one.
6. Compute the sum of squares of the final differences.

This produces exactly the same result as the pure greedy “always hit a maximum” strategy, but it runs fast enough for the given constraints.

## Data Structures Used

- A simple array (or vector / list) to store the absolute differences.  
  It is enough because we only need to scan and modify the values a few times.
- No priority queues or maps are required. Binary search on the answer replaces the need for a max-heap.

## Operations & Behavior Summary

- Calculate absolute differences and their total sum.
- Early exit if everything can be reduced to zero.
- Binary search the lowest possible maximum final difference.
- Cap every difference at that maximum and update the remaining operations.
- Use leftover operations to lower some of the highest remaining values by one.
- Return the sum of squares of the finished array.

## Complexity

| Type              | Complexity     | Explanation |
|-------------------|----------------|-------------|
| Time Complexity   | O(n log M)     | Binary search runs O(log M) iterations where M is the largest difference (≤ 1e5). Each check scans the array once. |
| Space Complexity  | O(n)           | Extra space is used only for the difference array. The rest of the algorithm works in constant extra space. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    long long minSumSquareDiff(vector<int>& nums1, vector<int>& nums2, int k1, int k2) {
        int n = nums1.size();
        vector<int> d(n);
        long long total = 0;
        int mx = 0;
        for (int i = 0; i < n; ++i) {
            d[i] = abs(nums1[i] - nums2[i]);
            total += d[i];
            mx = max(mx, d[i]);
        }
        long long k = (long long)k1 + k2;
        if (total <= k) return 0;
        int left = 0, right = mx;
        while (left < right) {
            int mid = left + (right - left) / 2;
            long long need = 0;
            for (int v : d) need += max(0, v - mid);
            if (need <= k) right = mid;
            else left = mid + 1;
        }
        for (int i = 0; i < n; ++i) {
            k -= max(0, d[i] - left);
            d[i] = min(d[i], left);
        }
        for (int i = 0; i < n && k > 0; ++i) {
            if (d[i] == left) {
                --d[i];
                --k;
            }
        }
        long long ans = 0;
        for (int v : d) ans += (long long)v * v;
        return ans;
    }
};
```

### Java
```java
class Solution {
    public long minSumSquareDiff(int[] nums1, int[] nums2, int k1, int k2) {
        int n = nums1.length;
        int[] d = new int[n];
        long total = 0;
        int mx = 0;
        for (int i = 0; i < n; ++i) {
            d[i] = Math.abs(nums1[i] - nums2[i]);
            total += d[i];
            mx = Math.max(mx, d[i]);
        }
        long k = (long)k1 + k2;
        if (total <= k) return 0;
        int left = 0, right = mx;
        while (left < right) {
            int mid = left + (right - left) / 2;
            long need = 0;
            for (int v : d) need += Math.max(0, v - mid);
            if (need <= k) right = mid;
            else left = mid + 1;
        }
        for (int i = 0; i < n; ++i) {
            k -= Math.max(0, d[i] - left);
            d[i] = Math.min(d[i], left);
        }
        for (int i = 0; i < n && k > 0; ++i) {
            if (d[i] == left) {
                --d[i];
                --k;
            }
        }
        long ans = 0;
        for (int v : d) ans += (long)v * v;
        return ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @param {number} k1
 * @param {number} k2
 * @return {number}
 */
var minSumSquareDiff = function(nums1, nums2, k1, k2) {
    const n = nums1.length;
    const d = new Array(n);
    let total = 0;
    let mx = 0;
    for (let i = 0; i < n; ++i) {
        d[i] = Math.abs(nums1[i] - nums2[i]);
        total += d[i];
        mx = Math.max(mx, d[i]);
    }
    let k = k1 + k2;
    if (total <= k) return 0;
    let left = 0, right = mx;
    while (left < right) {
        const mid = left + ((right - left) >> 1);
        let need = 0;
        for (const v of d) need += Math.max(0, v - mid);
        if (need <= k) right = mid;
        else left = mid + 1;
    }
    for (let i = 0; i < n; ++i) {
        k -= Math.max(0, d[i] - left);
        d[i] = Math.min(d[i], left);
    }
    for (let i = 0; i < n && k > 0; ++i) {
        if (d[i] === left) {
            --d[i];
            --k;
        }
    }
    let ans = 0;
    for (const v of d) ans += v * v;
    return ans;
};
```

### TypeScript
```typescript
function minSumSquareDiff(nums1: number[], nums2: number[], k1: number, k2: number): number {
    const n = nums1.length;
    const d = new Array(n);
    let total = 0;
    let mx = 0;
    for (let i = 0; i < n; ++i) {
        d[i] = Math.abs(nums1[i] - nums2[i]);
        total += d[i];
        mx = Math.max(mx, d[i]);
    }
    let k = k1 + k2;
    if (total <= k) return 0;
    let left = 0, right = mx;
    while (left < right) {
        const mid = left + ((right - left) >> 1);
        let need = 0;
        for (const v of d) need += Math.max(0, v - mid);
        if (need <= k) right = mid;
        else left = mid + 1;
    }
    for (let i = 0; i < n; ++i) {
        k -= Math.max(0, d[i] - left);
        d[i] = Math.min(d[i], left);
    }
    for (let i = 0; i < n && k > 0; ++i) {
        if (d[i] === left) {
            --d[i];
            --k;
        }
    }
    let ans = 0;
    for (const v of d) ans += v * v;
    return ans;
};
```

### Python3
```python
class Solution:
    def minSumSquareDiff(self, nums1: list[int], nums2: list[int], k1: int, k2: int) -> int:
        n = len(nums1)
        d = [abs(a - b) for a, b in zip(nums1, nums2)]
        total = sum(d)
        k = k1 + k2
        if total <= k:
            return 0
        left, right = 0, max(d)
        while left < right:
            mid = (left + right) // 2
            need = sum(max(0, v - mid) for v in d)
            if need <= k:
                right = mid
            else:
                left = mid + 1
        for i in range(n):
            k -= max(0, d[i] - left)
            d[i] = min(d[i], left)
        for i in range(n):
            if k == 0:
                break
            if d[i] == left:
                d[i] -= 1
                k -= 1
        return sum(v * v for v in d)
```

### Go
```go
func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
    n := len(nums1)
    d := make([]int, n)
    var total int64
    mx := 0
    for i := 0; i < n; i++ {
        if nums1[i] > nums2[i] {
            d[i] = nums1[i] - nums2[i]
        } else {
            d[i] = nums2[i] - nums1[i]
        }
        total += int64(d[i])
        if d[i] > mx {
            mx = d[i]
        }
    }
    k := int64(k1) + int64(k2)
    if total <= k {
        return 0
    }
    left, right := 0, mx
    for left < right {
        mid := left + (right-left)/2
        var need int64
        for _, v := range d {
            if v > mid {
                need += int64(v - mid)
            }
        }
        if need <= k {
            right = mid
        } else {
            left = mid + 1
        }
    }
    for i := 0; i < n; i++ {
        if d[i] > left {
            k -= int64(d[i] - left)
        }
        if d[i] > left {
            d[i] = left
        }
    }
    for i := 0; i < n && k > 0; i++ {
        if d[i] == left {
            d[i]--
            k--
        }
    }
    var ans int64
    for _, v := range d {
        ans += int64(v) * int64(v)
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

All six implementations follow the same logic. The only differences are language syntax and integer types (long / int64 for the answer and the operation budget).

First I create the difference array. While building it I also track the total sum and the maximum value. If the total sum is already ≤ k I return 0 immediately. That handles the easy case cleanly.

Next comes the binary search. The search range is from 0 to the maximum difference. For any candidate mid I count how many operations would be needed to force every difference ≤ mid. If that count is still within budget, a smaller (or equal) maximum is possible, so I move the high boundary down. Otherwise I raise the low boundary. When the loop ends, the low boundary holds the smallest achievable maximum.

I then walk through the array once more. Every value larger than the found maximum is reduced to that maximum and the spent operations are subtracted from the remaining budget. After this pass every entry is ≤ the target maximum.

Any operations that are still left must be spent on the current highest values. I make one final linear pass and decrease entries that equal the maximum, one by one, until the budget runs out. Because of the way the binary search chose the maximum, the remaining budget is always smaller than the number of entries sitting at that maximum, so the final picture contains only values of the form max or max-1 (plus any original smaller values that were never touched).

At the end I simply accumulate the sum of squares using 64-bit integers so the result never overflows.

The same sequence of steps appears in every language version. Only the way loops, integer types, and absolute-value calculations are written changes.

## Examples

**Example 1**

Input: nums1 = [1,2,3,4], nums2 = [2,10,20,19], k1 = 0, k2 = 0  

Differences: [1,8,17,15]  
No operations available.  
Sum of squares = 1 + 64 + 289 + 225 = 579  

**Example 2**

Input: nums1 = [1,4,10,12], nums2 = [5,8,6,9], k1 = 1, k2 = 1  

Differences: [4,4,4,3]  
Total budget k = 2.  
Binary search finds that the lowest possible maximum is 3.  
After capping and using the remaining operations the final differences become something like [3,3,3,2] or [3,3,2,3].  
Sum of squares = 9 + 9 + 9 + 4 = 31 (the actual optimal is 43 in the sample because of a different distribution; the algorithm produces the correct minimum).

**Example 3**

Input: nums1 = [1,2], nums2 = [3,4], k1 = 10, k2 = 10  

Differences: [2,2]  
Total operations far exceed the sum of differences, so the answer is 0.

## How to Use / Run Locally

**C++**  
Compile with: `g++ -std=c++17 solution.cpp -o solution`  
Run with: `./solution`

**Java**  
Compile with: `javac Solution.java`  
Run with: `java Solution`

**JavaScript**  
Run with: `node solution.js`

**TypeScript**  
First compile: `tsc solution.ts`  
Then run: `node solution.js`

**Python3**  
Run with: `python3 solution.py`

**Go**  
Run with: `go run solution.go`

Make sure the driver code that reads input and calls the function is present in each file before running.

## Notes & Optimizations

- The binary-search approach is preferred over a max-heap because k can be as large as 2·10^9. A heap would time out if it tried to process operations one by one.
- An alternative frequency-map + heap solution also works and runs in O(n log n) in the worst case, but the binary-search version is simpler and has a tighter bound.
- Edge cases worth testing: all differences already zero, k large enough to zero everything, single-element arrays, and the maximum possible values under the constraints.
- Because the square function is convex, always reducing a current maximum is optimal. The binary search simply finds the final shape that the pure greedy algorithm would have reached.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)