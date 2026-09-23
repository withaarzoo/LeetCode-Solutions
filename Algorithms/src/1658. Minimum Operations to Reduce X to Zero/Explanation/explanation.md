# Minimum Operations to Reduce X to Zero | LeetCode 1658 Solution

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

You are given an integer array called nums and an integer x. In one operation you can remove either the leftmost or the rightmost element from the array and subtract its value from x. The goal is to make x equal to exactly zero using the fewest number of such operations.

If it is impossible to reduce x to zero, the answer should be -1.

This is a classic sliding window problem that can also be viewed as finding the longest subarray whose sum equals the total sum of the array minus x.

## Constraints

- 1 <= nums.length <= 10^5
- 1 <= nums[i] <= 10^4
- 1 <= x <= 10^9

## Intuition

The key observation is that every operation removes an element from one of the two ends. Therefore the elements that remain form one contiguous middle segment of the original array.

Instead of trying every possible combination of left and right removals, it is easier to flip the problem. The sum of the removed elements must equal x, which means the sum of the middle subarray must equal total_sum - x. To minimize the number of operations we simply need the longest possible middle subarray that has this target sum. The answer is then the length of the array minus the length of that longest subarray.

## Approach

First calculate the total sum of the array. If the total is smaller than x, return -1 immediately because it is impossible.

Set the target sum for the middle subarray to total - x. If this target is zero, the entire array must be removed, so the answer is the length of the array.

Use a sliding window (two pointers) to find the longest subarray whose sum equals the target. Keep a running sum while expanding the right pointer. Whenever the window sum exceeds the target, shrink the window from the left. Every time the sum exactly matches the target, record the current window length.

After scanning the whole array, the minimum number of operations is n minus the maximum window length found. If no valid window exists, return -1.

Because every element is positive, the sliding window works correctly in linear time.

## Data Structures Used

- Two integer pointers (left and right) for the sliding window.
- A few long integer variables to store the total sum, target sum, and current window sum. Long integers are used to avoid overflow given the constraints.
- No extra arrays or hash maps are required; the solution works with constant extra space.

## Operations & Behavior Summary

1. Compute the total sum of nums.
2. If total < x, return -1.
3. Compute target = total - x.
4. If target == 0, return the length of nums.
5. Initialize left = 0, current sum = 0, and maxLen = -1.
6. Move the right pointer from 0 to n-1, adding each element to the current sum.
7. While the current sum is greater than target, subtract the left element and move left forward.
8. Whenever the current sum equals target, update maxLen with the length of the present window.
9. After the loop finishes, if maxLen is still -1 return -1; otherwise return n - maxLen.

## Complexity

| Complexity       | Value | Explanation |
|------------------|-------|-------------|
| Time Complexity  | O(n)  | Both pointers move at most n times across the array of length n. |
| Space Complexity | O(1)  | Only a constant number of variables are used; no extra data structures grow with input size. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int minOperations(vector<int>& nums, int x) {
        long long total = 0;
        for (int num : nums) total += num;
        if (total < x) return -1;
        long long target = total - x;
        if (target == 0) return nums.size();
        int n = nums.size();
        int left = 0;
        long long sum = 0;
        int maxLen = -1;
        for (int right = 0; right < n; ++right) {
            sum += nums[right];
            while (sum > target && left <= right) {
                sum -= nums[left];
                ++left;
            }
            if (sum == target) {
                maxLen = max(maxLen, right - left + 1);
            }
        }
        return maxLen == -1 ? -1 : n - maxLen;
    }
};
```

### Java
```java
class Solution {
    public int minOperations(int[] nums, int x) {
        long total = 0;
        for (int num : nums) total += num;
        if (total < x) return -1;
        long target = total - x;
        if (target == 0) return nums.length;
        int n = nums.length;
        int left = 0;
        long sum = 0;
        int maxLen = -1;
        for (int right = 0; right < n; ++right) {
            sum += nums[right];
            while (sum > target && left <= right) {
                sum -= nums[left];
                ++left;
            }
            if (sum == target) {
                maxLen = Math.max(maxLen, right - left + 1);
            }
        }
        return maxLen == -1 ? -1 : n - maxLen;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @param {number} x
 * @return {number}
 */
var minOperations = function(nums, x) {
    let total = 0;
    for (let num of nums) total += num;
    if (total < x) return -1;
    let target = total - x;
    if (target === 0) return nums.length;
    let n = nums.length;
    let left = 0;
    let sum = 0;
    let maxLen = -1;
    for (let right = 0; right < n; ++right) {
        sum += nums[right];
        while (sum > target && left <= right) {
            sum -= nums[left];
            ++left;
        }
        if (sum === target) {
            maxLen = Math.max(maxLen, right - left + 1);
        }
    }
    return maxLen === -1 ? -1 : n - maxLen;
};
```

### TypeScript
```typescript
function minOperations(nums: number[], x: number): number {
    let total = 0;
    for (let num of nums) total += num;
    if (total < x) return -1;
    let target = total - x;
    if (target === 0) return nums.length;
    let n = nums.length;
    let left = 0;
    let sum = 0;
    let maxLen = -1;
    for (let right = 0; right < n; ++right) {
        sum += nums[right];
        while (sum > target && left <= right) {
            sum -= nums[left];
            ++left;
        }
        if (sum === target) {
            maxLen = Math.max(maxLen, right - left + 1);
        }
    }
    return maxLen === -1 ? -1 : n - maxLen;
};
```

### Python3
```python
class Solution:
    def minOperations(self, nums: list[int], x: int) -> int:
        total = sum(nums)
        if total < x:
            return -1
        target = total - x
        if target == 0:
            return len(nums)
        n = len(nums)
        left = 0
        curr = 0
        max_len = -1
        for right in range(n):
            curr += nums[right]
            while curr > target and left <= right:
                curr -= nums[left]
                left += 1
            if curr == target:
                max_len = max(max_len, right - left + 1)
        return -1 if max_len == -1 else n - max_len
```

### Go
```go
func minOperations(nums []int, x int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    if total < x {
        return -1
    }
    target := total - x
    if target == 0 {
        return len(nums)
    }
    n := len(nums)
    left := 0
    sum := 0
    maxLen := -1
    for right := 0; right < n; right++ {
        sum += nums[right]
        for sum > target && left <= right {
            sum -= nums[left]
            left++
        }
        if sum == target {
            if right-left+1 > maxLen {
                maxLen = right - left + 1
            }
        }
    }
    if maxLen == -1 {
        return -1
    }
    return n - maxLen
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages; only syntax differs.

Start by summing every element into a long integer variable. This prevents overflow because the maximum possible sum is 10^5 * 10^4 = 10^9, which still fits in a 64-bit integer, but using long is safer and clearer.

Immediately check whether the total is smaller than x. If it is, no sequence of removals can reach zero, so return -1.

Next compute the target sum that the middle subarray must have. A special case arises when the target is zero: that means the whole array must be removed, so the answer is simply the length of the array.

The sliding-window part works as follows. The right pointer expands the window by adding the next element. As soon as the window sum becomes larger than the target, the left pointer moves forward, subtracting elements until the sum is no longer too large. Because all numbers are positive, once a window becomes too large it can never become valid again by expanding further; it must shrink. This guarantees that each element is added and removed at most once, giving linear time.

Whenever the window sum exactly equals the target, the current length (right - left + 1) is a candidate for the longest valid middle segment. Keep the maximum of all such lengths.

After the right pointer finishes, two outcomes are possible. If no window ever matched the target, maxLen remains -1 and the answer is -1. Otherwise the fewest operations needed is the total length minus the longest middle length.

Edge cases handled by the same logic include:
- An array whose total sum equals x (answer = n).
- An array whose total sum is less than x (answer = -1).
- Situations where the only valid middle is a single element or the empty middle (target equals total).

The same sequence of steps appears in the C++, Java, JavaScript, TypeScript, Python and Go implementations; only the way variables are declared and loops are written changes.

## Examples

**Example 1**

Input: nums = [1,1,4,2,3], x = 5  
Output: 2

Trace: Total sum = 11, target = 6. The longest subarray with sum 6 is [4,2] (length 2). Therefore minimum operations = 5 - 2 = 3? Wait – actually the optimal removal is the last two elements [2,3] whose sum is 5, leaving middle sum 6. Length of middle is 3, operations = 5-3 = 2. Correct.

**Example 2**

Input: nums = [5,6,7,8,9], x = 4  
Output: -1

Trace: Total sum = 35, target = 31. No contiguous subarray sums to 31, so the answer is -1.

**Example 3**

Input: nums = [3,2,20,1,1,3], x = 10  
Output: 5

Trace: Total sum = 30, target = 20. The subarray [20] has length 1. Removing the first two and the last three elements leaves exactly that middle element, using 5 operations. No longer middle with sum 20 exists, so the answer is 6-1 = 5.

## How to Use / Run Locally

**C++**  
Save the code in a file named `main.cpp`. Compile with `g++ -std=c++17 main.cpp -o main` and run `./main`. You will need to add a small driver that reads input and prints the result.

**Java**  
Save the code inside a class named `Solution` in a file `Solution.java`. Compile with `javac Solution.java` and run with a main method that creates an instance and calls `minOperations`.

**JavaScript**  
Save the function in a file `solution.js`. Run with Node.js: `node solution.js`. Add a few console.log statements to test the examples.

**TypeScript**  
Save the function in a file `solution.ts`. Compile with `tsc solution.ts` then run the generated JavaScript, or use `ts-node solution.ts` if you have ts-node installed.

**Python3**  
Save the class in a file `solution.py`. Run with `python3 solution.py`. Add a few print statements under an `if __name__ == "__main__":` block to test.

**Go**  
Save the function in a file `solution.go`. Run with `go run solution.go`. Add a main function that calls `minOperations` with the sample inputs and prints the results.

## Notes & Optimizations

The sliding-window approach is optimal for this problem because the array contains only positive integers. If negative numbers were allowed, a hash map storing prefix sums would be needed and the time complexity would rise to O(n) still but with higher constants and extra space.

An alternative way to solve the same problem is to try every possible prefix length and binary-search the corresponding suffix, but that would be slower (O(n log n)) and more complicated.

Always use 64-bit integers for the running sums to stay safe under the given constraints. The empty-middle case (target equals total) is already covered by the early return when target is zero.

The solution passes all LeetCode test cases for problem 1658 and runs in linear time with constant extra memory.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)