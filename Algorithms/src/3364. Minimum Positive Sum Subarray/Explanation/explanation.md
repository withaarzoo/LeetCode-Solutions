# Minimum Positive Sum Subarray - LeetCode 3364 Solution

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

You are given an integer array called nums and two integers l and r. Your job is to find a contiguous subarray whose length is between l and r (both ends included) and whose sum is strictly greater than zero. Among all such valid subarrays, return the one with the smallest positive sum. If no subarray meets the conditions, return -1.

A subarray means any continuous non-empty stretch of elements inside the original array. This is a classic minimum positive sum subarray problem that appears in many coding interviews and LeetCode contests.

## Constraints

- 1 <= nums.length <= 100
- 1 <= l <= r <= nums.length
- -1000 <= nums[i] <= 1000

Because the array length is tiny, even a simple O(n²) solution runs instantly.

## Intuition

The first thing I noticed is that the constraints are very small. With only 100 elements, I can safely check every possible subarray. I do not need a fancy sliding window or segment tree. I just need to look at every window whose length sits between l and r, calculate its sum, and keep the smallest sum that is still positive. That observation alone is enough to solve the minimum positive sum subarray problem cleanly.

## Approach

I loop over every possible starting index. From each start I keep adding the next numbers one by one while tracking the current sum and the current length. Whenever the length falls inside the range [l, r] and the sum is greater than zero, I compare it with the best answer I have seen so far and keep the smaller value. After checking every possible window I either return the best positive sum I found or -1 if nothing qualified.

## Data Structures Used

I only use the input array itself. No extra arrays, no maps, no heaps. A couple of integer variables hold the running sum and the current minimum. This keeps the solution extremely light on memory.

## Operations & Behavior Summary

1. Set the answer variable to a very large number so any real positive sum will replace it.  
2. For every left index i from 0 to n-1:  
   - Reset the running sum to zero.  
   - For every right index j from i to n-1:  
     - Add nums[j] to the running sum.  
     - Compute the length of the window.  
     - If length is between l and r and the sum is positive, update the answer.  
3. After both loops finish, return the answer if it was updated, otherwise return -1.

## Complexity

| Type              | Complexity | Explanation |
|-------------------|------------|-------------|
| Time Complexity   | O(n²)      | We examine every possible subarray once. With n ≤ 100 this is more than fast enough. |
| Space Complexity  | O(1)       | Only a handful of integer variables are used. No extra data structures grow with n. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int minimumSumSubarray(vector<int>& nums, int l, int r) {
        int n = nums.size();
        int ans = INT_MAX;
        for (int i = 0; i < n; i++) {
            int sum = 0;
            for (int j = i; j < n; j++) {
                sum += nums[j];
                int len = j - i + 1;
                if (len >= l && len <= r && sum > 0) {
                    ans = min(ans, sum);
                }
            }
        }
        return ans == INT_MAX ? -1 : ans;
    }
};
```

### Java
```java
class Solution {
    public int minimumSumSubarray(List<Integer> nums, int l, int r) {
        int n = nums.size();
        int ans = Integer.MAX_VALUE;
        for (int i = 0; i < n; i++) {
            int sum = 0;
            for (int j = i; j < n; j++) {
                sum += nums.get(j);
                int len = j - i + 1;
                if (len >= l && len <= r && sum > 0) {
                    ans = Math.min(ans, sum);
                }
            }
        }
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @param {number} l
 * @param {number} r
 * @return {number}
 */
var minimumSumSubarray = function(nums, l, r) {
    const n = nums.length;
    let ans = Infinity;
    for (let i = 0; i < n; i++) {
        let sum = 0;
        for (let j = i; j < n; j++) {
            sum += nums[j];
            const len = j - i + 1;
            if (len >= l && len <= r && sum > 0) {
                ans = Math.min(ans, sum);
            }
        }
    }
    return ans === Infinity ? -1 : ans;
};
```

### TypeScript
```typescript
function minimumSumSubarray(nums: number[], l: number, r: number): number {
    const n = nums.length;
    let ans = Infinity;
    for (let i = 0; i < n; i++) {
        let sum = 0;
        for (let j = i; j < n; j++) {
            sum += nums[j];
            const len = j - i + 1;
            if (len >= l && len <= r && sum > 0) {
                ans = Math.min(ans, sum);
            }
        }
    }
    return ans === Infinity ? -1 : ans;
};
```

### Python3
```python
class Solution:
    def minimumSumSubarray(self, nums: List[int], l: int, r: int) -> int:
        n = len(nums)
        ans = float('inf')
        for i in range(n):
            s = 0
            for j in range(i, n):
                s += nums[j]
                length = j - i + 1
                if l <= length <= r and s > 0:
                    ans = min(ans, s)
        return -1 if ans == float('inf') else ans
```

### Go
```go
func minimumSumSubarray(nums []int, l int, r int) int {
    n := len(nums)
    ans := math.MaxInt32
    for i := 0; i < n; i++ {
        sum := 0
        for j := i; j < n; j++ {
            sum += nums[j]
            length := j - i + 1
            if length >= l && length <= r && sum > 0 {
                if sum < ans {
                    ans = sum
                }
            }
        }
    }
    if ans == math.MaxInt32 {
        return -1
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages; only the syntax for infinity and list access changes.

I start by storing the length of the array so I do not keep asking for it inside the loops. I also initialize the answer to the largest possible integer (or positive infinity in languages that support it). That way the first valid positive sum I meet will become the new answer.

The outer loop picks every possible left endpoint. Inside it I reset a running sum to zero. The inner loop then walks rightward, adding one element at a time. After each addition I calculate the current window length with a simple subtraction.

I then ask three questions:  
- Is the length at least l?  
- Is the length at most r?  
- Is the sum strictly greater than zero?  

Only when all three answers are yes do I update the answer with the smaller of the two values.

When both loops are finished I check whether the answer is still the original large sentinel. If it is, no valid subarray existed and I return -1. Otherwise I return the smallest positive sum I recorded.

Edge cases are handled automatically:  
- All negative numbers → answer stays large → return -1.  
- l equals r and that exact length never produces a positive sum → return -1.  
- The whole array is positive and length is inside [l, r] → the smallest window sum is returned.

Because every language follows the exact same control flow, the explanation above covers C++, Java, JavaScript, TypeScript, Python3 and Go equally.

## Examples

**Example 1**  
Input: nums = [3, -2, 1, 4], l = 2, r = 3  
Possible windows of length 2 or 3 with positive sum:  
[3, -2] → 1  
[1, 4] → 5  
[3, -2, 1] → 2  
[-2, 1, 4] → 3  
The smallest positive sum is 1, so the answer is 1.

**Example 2**  
Input: nums = [-2, 2, -3, 1], l = 2, r = 3  
Every window of length 2 or 3 has a non-positive sum. Therefore the function returns -1.

**Example 3**  
Input: nums = [1, 2, 3, 4], l = 2, r = 4  
The shortest window [1, 2] already sums to 3, which is the minimum positive value among all valid lengths. Answer is 3.

## How to Use / Run Locally

**C++**  
g++ -std=c++17 solution.cpp -o solution && ./solution

**Java**  
javac Solution.java && java Solution

**JavaScript**  
node solution.js

**TypeScript**  
tsc solution.ts && node solution.js

**Python3**  
python3 solution.py

**Go**  
go run solution.go

Replace the empty function bodies with the code you want to test, then compile or run with the commands above. All of them expect the standard LeetCode function signature.

## Notes & Optimizations

Because n is at most 100, the O(n²) double loop is the simplest and most readable solution. A prefix-sum array would let each window sum be calculated in O(1), bringing the constant factor down a little, but it would not change the overall complexity and would use extra O(n) space. For interview settings the plain nested loops are usually preferred.

Watch out for the case when every possible window sums to zero or negative; the sentinel check correctly returns -1. Also remember that the subarray must be contiguous; the problem never asks for a subsequence.

This approach works for both positive and negative numbers and never overflows the 32-bit integer range given the constraints.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
