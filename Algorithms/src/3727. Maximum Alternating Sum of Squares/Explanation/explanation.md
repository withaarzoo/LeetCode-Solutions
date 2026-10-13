# Maximum Alternating Sum of Squares | LeetCode 3727 Solution

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

You are given an integer array called nums. You can rearrange the elements in any order you like.

The alternating score of an array is calculated as:

score = arr[0]² - arr[1]² + arr[2]² - arr[3]² + ...

Your task is to rearrange the array so that this alternating score becomes as large as possible, then return that maximum score.

In simple words, you need to place the numbers in an order that makes the positive squared terms as big as possible and the negative squared terms as small as possible.

## Constraints

- 1 <= nums.length <= 10⁵
- -4 *10⁴ <= nums[i] <= 4* 10⁴

## Intuition

The first thing I noticed is that every term is a square. Squaring removes the original sign of each number, so positive and negative values behave the same way.

That means the only decision that matters is which squared values sit on the plus positions (even indices) and which sit on the minus positions (odd indices).

To maximize the final sum I should put the largest squares on the plus spots and the smallest squares on the minus spots. Sorting the squares in descending order makes this assignment straightforward.

## Approach

1. Create a new list that stores the square of every number in nums.
2. Sort that list from largest to smallest.
3. Count how many plus positions exist. This is always (n + 1) // 2.
4. Walk through the sorted list once. Add the first (n + 1) // 2 values and subtract the remaining values.
5. Return the accumulated result.

This works because the largest available squares receive the positive signs and the smallest ones receive the negative signs.

## Data Structures Used

- Array / Vector / List of long integers: stores the squared values. I chose a separate array so the original nums stays unchanged and the sort becomes simple.
- Sorting algorithm (built-in): used to order the squares from largest to smallest in O(n log n) time.

No extra maps, sets, or trees are required. A plain array is enough.

## Operations & Behavior Summary

- Convert every element to its square.
- Sort the squares in descending order.
- Calculate the number of positive positions: (length + 1) divided by 2.
- Initialize answer to zero.
- Loop through the sorted squares:
  - If the current index is less than the positive-count, add the value.
  - Otherwise subtract the value.
- Return the final answer.

The whole process is one conversion pass, one sort, and one linear scan.

## Complexity

| Type              | Complexity   | Explanation                                                                 |
|-------------------|--------------|-----------------------------------------------------------------------------|
| Time Complexity   | O(n log n)   | Sorting the squares dominates. The conversion and the final scan are both O(n). n is the length of nums. |
| Space Complexity  | O(n)         | An extra array of size n is created to hold the squares. Sorting may use additional temporary space depending on the language implementation. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long maxAlternatingSum(vector<int>& nums) {
        int n = nums.size();
        vector<long long> sq(n);
        for (int i = 0; i < n; i++) {
            sq[i] = 1LL * nums[i] * nums[i];
        }
        sort(sq.rbegin(), sq.rend());
        long long ans = 0;
        int pos = (n + 1) / 2;
        for (int i = 0; i < n; i++) {
            if (i < pos) ans += sq[i];
            else ans -= sq[i];
        }
        return ans;
    }
};
```

### Java

```java
class Solution {
    public long maxAlternatingSum(int[] nums) {
        int n = nums.length;
        long[] sq = new long[n];
        for (int i = 0; i < n; i++) {
            sq[i] = (long) nums[i] * nums[i];
        }
        Arrays.sort(sq);
        long ans = 0;
        int pos = (n + 1) / 2;
        for (int i = 0; i < n; i++) {
            if (i < pos) ans += sq[n - 1 - i];
            else ans -= sq[n - 1 - i];
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
var maxAlternatingSum = function(nums) {
    const n = nums.length;
    const sq = nums.map(x => x * x);
    sq.sort((a, b) => b - a);
    let ans = 0;
    const pos = Math.floor((n + 1) / 2);
    for (let i = 0; i < n; i++) {
        if (i < pos) ans += sq[i];
        else ans -= sq[i];
    }
    return ans;
};
```

### TypeScript

```typescript
function maxAlternatingSum(nums: number[]): number {
    const n = nums.length;
    const sq = nums.map(x => x * x);
    sq.sort((a, b) => b - a);
    let ans = 0;
    const pos = Math.floor((n + 1) / 2);
    for (let i = 0; i < n; i++) {
        if (i < pos) ans += sq[i];
        else ans -= sq[i];
    }
    return ans;
};
```

### Python3

```python
class Solution:
    def maxAlternatingSum(self, nums: List[int]) -> int:
        n = len(nums)
        sq = [x * x for x in nums]
        sq.sort(reverse=True)
        ans = 0
        pos = (n + 1) // 2
        for i in range(n):
            if i < pos:
                ans += sq[i]
            else:
                ans -= sq[i]
        return ans
```

### Go

```go
func maxAlternatingSum(nums []int) int64 {
    n := len(nums)
    sq := make([]int64, n)
    for i, v := range nums {
        sq[i] = int64(v) * int64(v)
    }
    sort.Slice(sq, func(i, j int) bool {
        return sq[i] > sq[j]
    })
    var ans int64
    pos := (n + 1) / 2
    for i := 0; i < n; i++ {
        if i < pos {
            ans += sq[i]
        } else {
            ans -= sq[i]
        }
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

I start by reading the length of the input array. Then I build a new list that contains only the squared values. Squaring every number once removes any worry about signs and prepares the data for sorting.

Next I sort that list in descending order. In C++ I can reverse-sort directly. In Java I sort ascending and then read from the back. In JavaScript, TypeScript, Python and Go I sort with a descending comparator. After this step the largest square sits at index 0 and the smallest sits at the end.

I calculate how many positions will receive a plus sign: (n + 1) // 2. When n is odd there is one extra plus position; when n is even the counts are equal.

I then walk through the sorted list once. For every index that is smaller than the plus-count I add the square to the answer. For every remaining index I subtract it. Because the list is already ordered, the biggest squares automatically get the plus signs and the smallest get the minus signs.

Finally I return the answer. The same logic works in every language; only the syntax for sorting and type casting differs. Edge cases such as a single-element array or an array full of zeros are handled naturally because the sort and the single pass still produce the correct result.

## Examples

Example 1  
Input: nums = [1, 2, 3]  
Output: 12  

Squares become [1, 4, 9]. Sorted descending: [9, 4, 1].  
Two plus positions: 9 + 4 - 1 = 12.

Example 2  
Input: nums = [1, -1, 2, -2, 3, -3]  
Output: 16  

Squares become [1, 1, 4, 4, 9, 9]. Sorted descending: [9, 9, 4, 4, 1, 1].  
Three plus positions: 9 + 9 + 4 - 4 - 1 - 1 = 16.

Example 3  
Input: nums = [5]  
Output: 25  

Only one element. It sits on the single plus position, so the answer is simply 25.

## How to Use / Run Locally

C++  

1. Save the code in a file named main.cpp.  
2. Compile with: g++ -std=c++17 main.cpp -o main  
3. Run with: ./main  

Java  

1. Save the code in a file named Solution.java.  
2. Compile with: javac Solution.java  
3. Run with: java Solution  

JavaScript  

1. Save the code in a file named solution.js.  
2. Run with: node solution.js  

TypeScript  

1. Save the code in a file named solution.ts.  
2. Compile with: tsc solution.ts  
3. Run the generated JavaScript with: node solution.js  

Python3  

1. Save the code in a file named solution.py.  
2. Run with: python3 solution.py  

Go  

1. Save the code in a file named main.go.  
2. Run with: go run main.go  

In each case you will need to add a small main function or test harness that creates an example array and prints the result of the function call.

## Notes & Optimizations

- Because we only care about squares, negative numbers never need special handling.
- The solution works for both odd and even lengths without any extra branching.
- An in-place approach is possible: overwrite the original array with squares and sort it. That reduces the extra space to the temporary space used by the sort itself.
- For n = 1 the algorithm correctly returns the single squared value.
- The maximum possible absolute value of a square is 1.6 × 10⁹, so the final answer can reach roughly 1.6 × 10¹⁴. Always use 64-bit integers (long long, long, int64) to avoid overflow.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
