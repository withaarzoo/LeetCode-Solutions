# Minimum Moves to Equal Array Elements - LeetCode 453 Problem

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

You are given an integer array called nums. Your task is to make every element in the array equal using the fewest moves possible.

In one move you are allowed to pick any n - 1 elements (where n is the size of the array) and add 1 to each of them. You can perform this operation as many times as you want.

The function should return the minimum number of such moves needed to make all values identical.

This is a classic array problem that appears in coding interviews and competitive programming contests. It tests how well you can reframe an operation into something simpler.

## Constraints

- 1 <= nums.length <= 10^5
- -10^9 <= nums[i] <= 10^9
- The answer is guaranteed to fit inside a 32-bit signed integer

## Intuition

When I first read the move, it felt a bit awkward. Adding 1 to almost every element is the same as subtracting 1 from the single element you left out, if you only care about the relative differences.

That observation changes everything. Instead of thinking about growing numbers, I can think about shrinking the gaps between them. The cheapest target I can aim for is the smallest number already present in the array. Every other number simply needs to be brought down to that minimum value through relative decrements. The total number of decrements is exactly the answer.

## Approach

1. Scan the array once and find the minimum value. Call it the target.
2. Scan the array a second time. For every element, add the difference between that element and the target to a running total.
3. Return the total. That sum is the minimum number of moves required.

No sorting is needed. No extra collections are needed. Two linear passes are enough.

## Data Structures Used

- Plain array / vector / slice: The input itself is the only structure required. We read it directly and never create additional dynamic containers.
- A few scalar variables: One integer to store the minimum and one integer (or long) to accumulate the sum of differences.

These choices keep both time and memory as low as possible for an n = 10^5 problem.

## Operations & Behavior Summary

- Locate the smallest element in the array.
- Initialize a counter to zero.
- Walk through every element again.
- For each element, add (element - minimum) to the counter.
- After the loop finishes, the counter holds the answer.

The algorithm never modifies the original array. It only reads values and computes a simple arithmetic sum.

## Complexity

| Type              | Value | Explanation                                      |
|-------------------|-------|--------------------------------------------------|
| Time Complexity   | O(n)  | Two linear passes over the array of length n     |
| Space Complexity  | O(1)  | Only a constant number of extra variables used   |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minMoves(vector<int>& nums) {
        int mn = *min_element(nums.begin(), nums.end());
        long long moves = 0;
        for (int x : nums) moves += x - mn;
        return moves;
    }
};
```

### Java

```java
class Solution {
    public int minMoves(int[] nums) {
        int mn = Integer.MAX_VALUE;
        for (int x : nums) if (x < mn) mn = x;
        long moves = 0;
        for (int x : nums) moves += x - mn;
        return (int) moves;
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
    let mn = Math.min(...nums);
    let moves = 0;
    for (let x of nums) moves += x - mn;
    return moves;
};
```

### TypeScript

```typescript
function minMoves(nums: number[]): number {
    let mn = Math.min(...nums);
    let moves = 0;
    for (let x of nums) moves += x - mn;
    return moves;
};
```

### Python3

```python
class Solution:
    def minMoves(self, nums: list[int]) -> int:
        mn = min(nums)
        return sum(x - mn for x in nums)
```

### Go

```go
func minMoves(nums []int) int {
    mn := nums[0]
    for _, x := range nums {
        if x < mn {
            mn = x
        }
    }
    moves := 0
    for _, x := range nums {
        moves += x - mn
    }
    return moves
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core idea stays identical across every language. Only the syntax for finding the minimum and summing the differences changes.

First I obtain the minimum value. In C++ I can rely on the standard library min_element. In Java, JavaScript, TypeScript, Python and Go I either call a built-in min function or write a short loop that keeps track of the smallest number seen so far.

Once the minimum is known, I start a second pass. I initialize a variable that will hold the total number of moves. In languages where the sum of differences might temporarily exceed a 32-bit integer, I use a wider type (long long in C++, long in Java) even though the problem guarantees the final answer fits in 32 bits.

For every element I compute element minus minimum and add that positive difference to the accumulator. Because the minimum is the smallest value, every difference is non-negative, so the sum never decreases.

When the second pass ends, the accumulator contains the exact number of relative decrements needed. That number is also the minimum number of original moves, because each relative decrement corresponds to one original “add 1 to n-1 elements” operation.

Edge cases are handled naturally. If every element is already equal, the minimum equals every value, so the sum of differences is zero and the function returns 0. Negative numbers cause no trouble because the arithmetic stays the same; the minimum can simply be a negative value and the differences remain correct.

No language-specific special cases appear. The same two-pass logic works uniformly.

## Examples

Example 1  
Input: nums = [1, 2, 3]  
Output: 3  

Trace:  
Minimum is 1.  
Differences: 1-1 = 0, 2-1 = 1, 3-1 = 2.  
Sum = 0 + 1 + 2 = 3.  
Three moves turn the array into [4, 4, 4].

Example 2  
Input: nums = [1, 1, 1]  
Output: 0  

Trace:  
Minimum is 1.  
All differences are zero.  
Sum = 0.  
No moves are required.

Example 3  
Input: nums = [5, 6, 8, 8, 5]  
Output: 7  

Trace:  
Minimum is 5.  
Differences: 0 + 1 + 3 + 3 + 0 = 7.  
Seven relative decrements bring every element down to 5.

## How to Use / Run Locally

C++  
Compile with g++ -std=c++17 solution.cpp -o solution and run ./solution. Feed the array through standard input or hard-code a test case inside main.

Java  
Save the class in a file named Solution.java. Compile with javac Solution.java and run java Solution. Add a main method that creates a sample array and prints the result.

JavaScript  
Save the function in a .js file. Run it with node solution.js. You can call the function directly from the same file with a test array.

TypeScript  
Write the function in a .ts file. Compile with tsc solution.ts and then run the generated JavaScript with node, or use ts-node for direct execution.

Python3  
Save the class in a .py file. Run python3 solution.py. Instantiating the class and calling minMoves with a list works immediately.

Go  
Place the function in a file with package main. Add a main function that creates a slice and prints the result. Build and run with go run solution.go.

In every language, replace the empty function body with the actual implementation before testing.

## Notes & Optimizations

The solution is already optimal for both time and space. A single pass that tracks the minimum while summing is possible but does not improve the asymptotic complexity and makes the code slightly harder to read, so the clear two-pass version is preferred.

Because the answer is guaranteed to fit in a 32-bit integer, you can safely cast the final sum back to int in languages that used a wider accumulator.

An alternative mathematical view is that the problem is equivalent to making all elements equal to the minimum by repeated decrements. That perspective leads to the same code and helps when explaining the solution in interviews.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
