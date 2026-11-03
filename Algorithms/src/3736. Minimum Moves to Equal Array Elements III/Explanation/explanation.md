# Minimum Moves to Equal Array Elements III - LeetCode 3736 Problem

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

You are given an integer array called nums. In one move you are allowed to increase any single element by 1. Your task is to make every element in the array equal using the smallest possible number of moves.

Because you can only increase values, the final common value cannot be smaller than the largest number already present in the array. The goal is to return that minimum number of increments.

This is a classic minimum moves to equal array elements problem that appears frequently in coding interviews and online judges.

## Constraints

- 1 <= nums.length <= 100
- 1 <= nums[i] <= 100

The constraints are small, which means even a simple linear scan is more than enough.

## Intuition

The first thing I noticed is that we can only go up. That immediately tells me the target value must be at least as large as the current maximum. Choosing anything higher than the maximum would only add extra work for every element. So the cheapest target is exactly the maximum value already in the array. Once I realized that, the rest became straightforward: just count how many increments are needed to bring every other number up to that maximum.

## Approach

1. Scan the array once and find the largest value. Call it maxVal.
2. Walk through the array a second time.
3. For every number, add the difference (maxVal - current number) to a running total.
4. Return the total. That total is the minimum number of moves required.

No sorting is needed. No extra collections are required. Just two linear passes (or even one combined pass if preferred).

## Data Structures Used

- The input array itself is the only structure we need.
- A few integer variables hold the maximum value and the running sum of differences.
- No stacks, queues, maps, or extra arrays are allocated.

## Operations & Behavior Summary

- Locate the highest number present in nums.
- For each element, calculate how many times it must be incremented to reach that highest number.
- Accumulate all those increments.
- The final accumulated value is the answer.

This is essentially a “raise everything to the ceiling” strategy.

## Complexity

| Type              | Complexity | Explanation                                      |
|-------------------|------------|--------------------------------------------------|
| Time Complexity   | O(n)       | One or two linear scans over the array of size n |
| Space Complexity  | O(1)       | Only a constant number of integer variables used |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minMoves(vector<int>& nums) {
        int mx = *max_element(nums.begin(), nums.end());
        int ans = 0;
        for (int x : nums) ans += mx - x;
        return ans;
    }
};
```

### Java

```java
class Solution {
    public int minMoves(int[] nums) {
        int mx = nums[0];
        for (int x : nums) if (x > mx) mx = x;
        int ans = 0;
        for (int x : nums) ans += mx - x;
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
var minMoves = function(nums) {
    let mx = Math.max(...nums);
    let ans = 0;
    for (let x of nums) ans += mx - x;
    return ans;
};
```

### TypeScript

```typescript
function minMoves(nums: number[]): number {
    let mx = Math.max(...nums);
    let ans = 0;
    for (let x of nums) ans += mx - x;
    return ans;
};
```

### Python3

```python
class Solution:
    def minMoves(self, nums: List[int]) -> int:
        mx = max(nums)
        return sum(mx - x for x in nums)
```

### Go

```go
func minMoves(nums []int) int {
    mx := nums[0]
    for _, x := range nums {
        if x > mx {
            mx = x
        }
    }
    ans := 0
    for _, x := range nums {
        ans += mx - x
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

Across all six languages the core idea stays identical.

First we discover the maximum value. In C++ this is commonly done with max_element or a simple loop. Java, Go, and the others use an ordinary for-loop that keeps updating a variable whenever a larger number appears. JavaScript, TypeScript and Python can also use their built-in max helpers for brevity.

Once the maximum is known, a second pass (or a generator expression in Python) walks every element and adds the gap between the maximum and that element to an answer variable. Because every gap is non-negative, the sum is guaranteed to be the exact number of increments required.

Edge cases are handled naturally:

- If the array contains only one element, the maximum equals that element and the answer is zero.
- If every element is already equal, every difference is zero and the answer is zero.
- The small constraint range means overflow is not a practical concern in any of the languages for this particular problem.

No language-specific tricks are required beyond the usual idioms for finding a maximum and summing differences.

## Examples

Example 1  
Input: nums = [2, 1, 3]  
Maximum = 3  
Differences: 3-2 = 1, 3-1 = 2, 3-3 = 0  
Total moves = 3  
Output: 3

Example 2  
Input: nums = [4, 4, 5]  
Maximum = 5  
Differences: 5-4 = 1, 5-4 = 1, 5-5 = 0  
Total moves = 2  
Output: 2

Example 3  
Input: nums = [7, 7, 7]  
Maximum = 7  
All differences are zero  
Total moves = 0  
Output: 0

## How to Use / Run Locally

1. Copy the solution for the language you want into a file (for example Solution.cpp, Solution.java, etc.).
2. Add a simple main / driver function that creates a test array and prints the result of the minMoves call.
3. Compile and run with the usual commands:
   - C++: g++ -std=c++17 Solution.cpp -o solution && ./solution
   - Java: javac Solution.java && java Solution
   - JavaScript: node Solution.js
   - TypeScript: tsc Solution.ts && node Solution.js
   - Python: python3 Solution.py
   - Go: go run Solution.go
4. Change the test array inside main to try different inputs.

## Notes & Optimizations

- Because we can only increase values, choosing any target larger than the current maximum is never beneficial.
- The algorithm is already optimal in both time and space for the given constraints.
- If the problem later allowed decreasing as well, the optimal target would become the median; that is a different problem.
- For extremely large arrays the same linear approach still holds; the tiny constraints here simply make the solution feel instantaneous.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
