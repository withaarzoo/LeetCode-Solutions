# Minimum Moves to Equal Array Elements II - LeetCode 462 Solution

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

You are given an integer array called nums. Your task is to make every element in the array equal using the least number of moves possible.

In one move you can pick any element and either increase it by 1 or decrease it by 1.

The goal is to find that minimum number of moves. The answer will always fit inside a 32-bit integer.

This is the classic “minimum moves to equal array elements II” problem that appears on LeetCode as problem 462. It forces you to think about the best target value instead of trying every possible number.

## Constraints

- 1 ≤ nums.length ≤ 10⁵
- -10⁹ ≤ nums[i] ≤ 10⁹
- The final answer always fits in a 32-bit signed integer

## Intuition

When I first looked at the problem I asked myself: “What number should I try to turn every element into?”

If I pick a number that is too small, the larger values cost a lot of moves. If I pick a number that is too large, the smaller values cost a lot of moves. The sweet spot is right in the middle of the sorted list.

That middle value is the median. The median is the number that minimizes the total distance (absolute difference) from all other numbers. Once I realized that, the rest of the solution became straightforward.

## Approach

1. Sort the array so the numbers sit in increasing order.
2. Pick the median. For an array of length n the median is the element at index n/2 (integer division). When n is even this chooses the lower of the two middle values; either works.
3. Walk through every element and add up how far it is from the median.
4. That sum is the minimum number of moves required.

No extra data structures are needed beyond the sorted array itself. The whole idea rests on the mathematical fact that the median minimizes the sum of absolute deviations.

## Data Structures Used

- Array / Vector – the input itself. We sort it in place so we can reach the median in constant time after sorting.
- No auxiliary arrays, maps, or heaps are required. The algorithm only needs the ordered list of numbers.

## Operations & Behavior Summary

- Sort the input array in non-decreasing order.
- Locate the median at the middle index.
- Initialize a counter to zero.
- For every number compute the absolute difference from the median and add it to the counter.
- Return the final counter value.

That short sequence is the entire algorithm.

## Complexity

| Type              | Complexity   | Explanation                                                                 |
|-------------------|--------------|-----------------------------------------------------------------------------|
| Time Complexity   | O(n log n)   | Sorting dominates. The later linear scan is O(n). n is the length of nums. |
| Space Complexity  | O(1) extra   | Sorting can be done in-place. Only a few variables are used besides the input. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int minMoves2(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int median = nums[nums.size() / 2];
        long long moves = 0;
        for (int num : nums) {
            moves += abs(num - median);
        }
        return moves;
    }
};
```

### Java

```java
class Solution {
    public int minMoves2(int[] nums) {
        Arrays.sort(nums);
        int median = nums[nums.length / 2];
        long moves = 0;
        for (int num : nums) {
            moves += Math.abs(num - median);
        }
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
var minMoves2 = function(nums) {
    nums.sort((a, b) => a - b);
    const median = nums[Math.floor(nums.length / 2)];
    let moves = 0;
    for (const num of nums) {
        moves += Math.abs(num - median);
    }
    return moves;
};
```

### TypeScript

```typescript
function minMoves2(nums: number[]): number {
    nums.sort((a, b) => a - b);
    const median = nums[Math.floor(nums.length / 2)];
    let moves = 0;
    for (const num of nums) {
        moves += Math.abs(num - median);
    }
    return moves;
};
```

### Python3

```python
class Solution:
    def minMoves2(self, nums: list[int]) -> int:
        nums.sort()
        median = nums[len(nums) // 2]
        moves = 0
        for num in nums:
            moves += abs(num - median)
        return moves
```

### Go

```go
func minMoves2(nums []int) int {
    sort.Ints(nums)
    median := nums[len(nums)/2]
    moves := 0
    for _, num := range nums {
        diff := num - median
        if diff < 0 {
            diff = -diff
        }
        moves += diff
    }
    return moves
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical in every language; only the syntax changes.

First the array is sorted. Sorting guarantees that the element sitting at index length/2 is the median. Integer division automatically handles both odd and even lengths correctly.

After the median is known, a single pass computes the total distance. Absolute difference is used because a move can go either up or down; the direction does not matter, only the distance.

In languages that have a 32-bit integer limit (Java, Go, C++ when using int), the accumulator is kept as a 64-bit type so intermediate sums never overflow. The final return value is then cast back to 32-bit because the problem guarantees it will fit.

Edge cases are handled naturally:

- A single-element array needs zero moves.
- An array with two elements needs the absolute difference between them.
- Negative numbers and large positive numbers are treated the same way because absolute difference works for any integer.

No special handling is required for duplicates; the median still works.

## Examples

**Example 1**  
Input: nums = [1, 2, 3]  
Sorted: [1, 2, 3]  
Median: 2  
Moves: |1-2| + |2-2| + |3-2| = 1 + 0 + 1 = 2  
Output: 2

**Example 2**  
Input: nums = [1, 10, 2, 9]  
Sorted: [1, 2, 9, 10]  
Median: 9 (index 4/2 = 2)  
Moves: |1-9| + |2-9| + |9-9| + |10-9| = 8 + 7 + 0 + 1 = 16  
Output: 16

**Example 3**  
Input: nums = [1, 0, 0, 8, 6]  
Sorted: [0, 0, 1, 6, 8]  
Median: 1  
Moves: |0-1| + |0-1| + |1-1| + |6-1| + |8-1| = 1 + 1 + 0 + 5 + 7 = 14  
Output: 14

## How to Use / Run Locally

**C++**  
Save the code in a file named main.cpp. Compile with  
`g++ -std=c++17 main.cpp -o main`  
Run with `./main` after adding a simple driver that reads the array and prints the result.

**Java**  
Save as Solution.java. Compile with `javac Solution.java`.  
Run with `java Solution` (add a main method that creates a Solution object and calls minMoves2).

**JavaScript**  
Save as solution.js. Run with Node.js:  
`node solution.js`  
(Add a few console.log calls with sample arrays to test.)

**TypeScript**  
Save as solution.ts. Compile with `tsc solution.ts` then run the generated JavaScript file with Node.

**Python3**  
Save as solution.py. Run with  
`python3 solution.py`  
(Add a short if **name** == "**main**" block to test the method.)

**Go**  
Save as main.go. Run with  
`go run main.go`  
(Add a main function that creates a slice and prints the result of minMoves2.)

## Notes & Optimizations

- Using the mean instead of the median produces a wrong answer on many test cases. The median is the only correct target for the sum of absolute differences.
- For even-length arrays any integer between the two middle values works; the code simply picks the lower one for simplicity.
- The O(n log n) sorting step is already optimal for this problem under the given constraints. A linear-time selection algorithm could find the median in O(n) expected time, but the constant factors and coding complexity are higher and not worth it for n ≤ 10⁵.
- Negative numbers and the full 32-bit range are handled correctly by absolute difference; no extra clamping or special cases are required.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
