# Longest Strictly Increasing or Strictly Decreasing Subarray - LeetCode 3105 Solution

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

You are given an array of integers called nums. Your task is to find the length of the longest subarray that is either strictly increasing or strictly decreasing.

A strictly increasing subarray means every next number is bigger than the one before it. A strictly decreasing subarray means every next number is smaller than the one before it. Equal numbers break both patterns.

You need to return only the length of the longest such subarray. Single elements are always valid and have length 1.

This is a classic LeetCode easy problem that tests your ability to track monotonic sequences efficiently.

## Constraints

- 1 <= nums.length <= 50
- 1 <= nums[i] <= 50

The small input size means even a simple O(n) scan is more than enough, but the clean one-pass solution is still preferred.

## Intuition

The first thing I noticed is that a monotonic subarray can only grow while consecutive numbers keep moving in the same direction. The moment the direction flips or two numbers become equal, the current streak ends and a new one must start.

Because the array is short, I could have checked every possible subarray. But a smarter observation is that I only need to remember two running lengths: the current increasing streak and the current decreasing streak. Whichever is larger at any point becomes a candidate for the answer. This leads to a clean linear scan.

## Approach

I keep three variables. maxLen starts at 1. inc tracks the length of the strictly increasing streak ending at the current index. dec tracks the length of the strictly decreasing streak ending at the current index.

Both streaks begin at 1 for the first element.

I then walk from the second element to the end of the array.

- If the current number is greater than the previous one, I grow the increasing streak by one and reset the decreasing streak to 1.
- If the current number is smaller, I grow the decreasing streak and reset the increasing streak to 1.
- If the numbers are equal, both streaks reset to 1.

After every comparison I update maxLen with the larger of the two current streaks.

When the loop finishes, maxLen holds the length of the longest strictly increasing or strictly decreasing subarray.

## Data Structures Used

Only a few integer variables are needed. No extra arrays, maps, or stacks are required. This keeps the solution extremely light on memory and easy to understand for beginners.

## Operations & Behavior Summary

1. Initialize maxLen, inc, and dec to 1.
2. Loop from index 1 to the end of the array.
3. Compare the current element with the previous one.
4. Update the appropriate streak and reset the other.
5. Keep the global maximum after every update.
6. Return the final maximum length.

This is essentially a single-pass state machine that tracks the two possible directions.

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | O(n) | We examine each pair of consecutive elements exactly once. n is the length of nums. |
| Space Complexity | O(1) | Only a constant number of integer variables are used. No extra data structures grow with input size. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int longestMonotonicSubarray(vector<int>& nums) {
        int n = nums.size();
        int maxLen = 1, inc = 1, dec = 1;
        for (int i = 1; i < n; i++) {
            if (nums[i] > nums[i - 1]) {
                inc++;
                dec = 1;
            } else if (nums[i] < nums[i - 1]) {
                dec++;
                inc = 1;
            } else {
                inc = 1;
                dec = 1;
            }
            maxLen = max(maxLen, max(inc, dec));
        }
        return maxLen;
    }
};
```

### Java
```java
class Solution {
    public int longestMonotonicSubarray(int[] nums) {
        int n = nums.length;
        int maxLen = 1, inc = 1, dec = 1;
        for (int i = 1; i < n; i++) {
            if (nums[i] > nums[i - 1]) {
                inc++;
                dec = 1;
            } else if (nums[i] < nums[i - 1]) {
                dec++;
                inc = 1;
            } else {
                inc = 1;
                dec = 1;
            }
            maxLen = Math.max(maxLen, Math.max(inc, dec));
        }
        return maxLen;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var longestMonotonicSubarray = function(nums) {
    let n = nums.length;
    let maxLen = 1, inc = 1, dec = 1;
    for (let i = 1; i < n; i++) {
        if (nums[i] > nums[i - 1]) {
            inc++;
            dec = 1;
        } else if (nums[i] < nums[i - 1]) {
            dec++;
            inc = 1;
        } else {
            inc = 1;
            dec = 1;
        }
        maxLen = Math.max(maxLen, Math.max(inc, dec));
    }
    return maxLen;
};
```

### TypeScript
```typescript
function longestMonotonicSubarray(nums: number[]): number {
    let n = nums.length;
    let maxLen = 1, inc = 1, dec = 1;
    for (let i = 1; i < n; i++) {
        if (nums[i] > nums[i - 1]) {
            inc++;
            dec = 1;
        } else if (nums[i] < nums[i - 1]) {
            dec++;
            inc = 1;
        } else {
            inc = 1;
            dec = 1;
        }
        maxLen = Math.max(maxLen, Math.max(inc, dec));
    }
    return maxLen;
};
```

### Python3
```python
class Solution:
    def longestMonotonicSubarray(self, nums: List[int]) -> int:
        n = len(nums)
        maxLen = 1
        inc = 1
        dec = 1
        for i in range(1, n):
            if nums[i] > nums[i - 1]:
                inc += 1
                dec = 1
            elif nums[i] < nums[i - 1]:
                dec += 1
                inc = 1
            else:
                inc = 1
                dec = 1
            maxLen = max(maxLen, inc, dec)
        return maxLen
```

### Go
```go
func longestMonotonicSubarray(nums []int) int {
    n := len(nums)
    maxLen, inc, dec := 1, 1, 1
    for i := 1; i < n; i++ {
        if nums[i] > nums[i-1] {
            inc++
            dec = 1
        } else if nums[i] < nums[i-1] {
            dec++
            inc = 1
        } else {
            inc = 1
            dec = 1
        }
        if inc > maxLen {
            maxLen = inc
        }
        if dec > maxLen {
            maxLen = dec
        }
    }
    return maxLen
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages. Here is the reasoning behind every decision.

I start by recording the array length and setting three variables to 1. maxLen must start at 1 because every single element is a valid subarray. The two streak counters also start at 1 so that the first element is already counted.

The loop begins at index 1. Starting earlier would try to read a non-existent previous element and cause an error.

Inside the loop I compare nums[i] with nums[i-1].

When the numbers are strictly increasing I increment the increasing streak and force the decreasing streak back to 1. The opposite happens when the numbers are strictly decreasing. Equality forces both streaks to restart because neither pattern can continue.

After updating the streaks I always take the maximum of the current maxLen and the two streak values. This guarantees I never miss a longer streak that ended earlier in the array.

If I forgot to reset the opposite streak, the counters would keep growing across direction changes and produce a wrong answer. Keeping everything in plain integers avoids any memory allocation and makes the code easy to translate between languages.

Edge cases such as an array of length 1, an array of all equal numbers, or a fully monotonic array are all handled correctly by the same logic.

## Examples

Example 1  
Input: nums = [1, 4, 3, 3, 2]  
Output: 2  

Trace:  
- 1 → 4 increases → increasing streak becomes 2  
- 4 → 3 decreases → decreasing streak becomes 2  
- 3 → 3 equal → both streaks reset to 1  
- 3 → 2 decreases → decreasing streak becomes 2  
The longest streak seen is 2.

Example 2  
Input: nums = [3, 3, 3, 3]  
Output: 1  

Trace:  
Every pair of consecutive numbers is equal, so both streaks stay at 1 the entire time. The answer is correctly 1.

Example 3  
Input: nums = [3, 2, 1]  
Output: 3  

Trace:  
- 3 → 2 decreases → decreasing streak becomes 2  
- 2 → 1 decreases → decreasing streak becomes 3  
The entire array is strictly decreasing, so the answer is 3.

## How to Use / Run Locally

C++  
Compile with g++ and run the executable. Paste the solution into a file that contains a main function calling the method with sample inputs.

Java  
Save the class in a file named Solution.java, compile with javac, and run with java. Add a main method for local testing.

JavaScript  
Paste the function into a .js file and run it with Node.js. Call the function with different arrays and log the results.

TypeScript  
Compile the .ts file with tsc or run it directly with ts-node. The function signature is ready for testing.

Python3  
Save the class in a .py file and run it with python3. Instantiating the class and calling the method is straightforward.

Go  
Place the function in a package main file, add a main function that prints the results, then run with go run.

In every language the core function is self-contained and only needs the input array.

## Notes & Optimizations

The O(n) one-pass solution is already optimal for this problem. Because the constraints are tiny, a nested loop checking every subarray would also pass, but it is slower and less elegant.

Watch out for the all-equal case. Many people forget that equal numbers break both increasing and decreasing patterns.

Another possible approach is to compute the longest increasing and longest decreasing lengths in two separate passes and take the maximum. The single-pass version above is cleaner and uses the same amount of work.

The solution works correctly for the minimum array length of 1 and never produces an answer smaller than 1.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
