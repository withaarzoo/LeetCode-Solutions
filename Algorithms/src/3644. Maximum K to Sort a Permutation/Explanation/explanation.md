# Maximum K to Sort a Permutation - LeetCode 3644 Solution

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

You are given an integer array `nums` of length `n`. This array is a permutation of the numbers from `0` to `n-1`.

You are allowed to swap any two elements at indices `i` and `j` only when the bitwise AND of those two values equals a fixed non-negative integer `k`. In other words, you can swap `nums[i]` and `nums[j]` only if `nums[i] & nums[j] == k`.

Your task is to find the largest possible value of `k` such that you can sort the entire array into non-decreasing order using any number of these restricted swaps. If the array is already sorted, the answer is `0`.

This problem combines permutation sorting with bit manipulation and is a good practice question for competitive programming and DSA interviews.

## Constraints

- `1 <= n == nums.length <= 10^5`
- `0 <= nums[i] <= n - 1`
- `nums` is a permutation of integers from `0` to `n - 1`

## Intuition

When I first read the problem I noticed that any number already sitting in its correct position never needs to move. That means only the misplaced numbers matter.

Those misplaced numbers have to be rearranged among themselves under the AND constraint. The largest `k` that can still work must be a value whose bits are present in every misplaced number. Taking the bitwise AND of all misplaced values gives exactly that candidate. Anything larger would contain a bit that at least one misplaced number does not have, so it could never appear as the result of an AND between two of them.

This observation leads to a simple linear scan that produces the correct maximum `k`.

## Approach

I walk through the array once from left to right.  
For every index `i` I check whether `nums[i]` equals `i`.  
If it does not, the value is misplaced and I include it in a running bitwise AND.  
If the whole array is already sorted, the running AND stays empty and I return `0`.  
Otherwise the final AND value is the largest possible `k`.

No extra arrays or maps are required. One pass is enough.

## Data Structures Used

- A single integer variable that holds the running bitwise AND of all misplaced values.  
  I chose a plain integer because the operation is associative and commutative, so I never need to store the individual values.

No other data structures are used. This keeps both time and space optimal for the given constraints.

## Operations & Behavior Summary

1. Initialize an answer variable to a sentinel value that means “no misplaced element seen yet”.
2. Loop over every index `i` from `0` to `n-1`.
3. If `nums[i]` is not equal to `i`, update the running AND with that value.
4. After the loop finishes:
   - If the sentinel is still present, the array was already sorted → return `0`.
   - Otherwise return the computed AND value.

That is the entire algorithm.

## Complexity

| Type              | Complexity | Explanation                                                                 |
|-------------------|------------|-----------------------------------------------------------------------------|
| Time Complexity   | O(n)       | One linear scan over the array of length n. Each element is examined once. |
| Space Complexity  | O(1)       | Only a constant number of integer variables are used. No extra arrays.     |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int sortPermutation(vector<int>& nums) {
        int ans = -1;
        for (int i = 0; i < (int)nums.size(); i++) {
            if (nums[i] != i) {
                if (ans == -1) ans = nums[i];
                else ans &= nums[i];
            }
        }
        return ans == -1 ? 0 : ans;
    }
};
```

### Java
```java
class Solution {
    public int sortPermutation(int[] nums) {
        int ans = -1;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != i) {
                if (ans == -1) ans = nums[i];
                else ans &= nums[i];
            }
        }
        return ans == -1 ? 0 : ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var sortPermutation = function(nums) {
    let ans = -1;
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] !== i) {
            if (ans === -1) ans = nums[i];
            else ans &= nums[i];
        }
    }
    return ans === -1 ? 0 : ans;
};
```

### TypeScript
```typescript
function sortPermutation(nums: number[]): number {
    let ans = -1;
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] !== i) {
            if (ans === -1) ans = nums[i];
            else ans &= nums[i];
        }
    }
    return ans === -1 ? 0 : ans;
};
```

### Python3
```python
class Solution:
    def sortPermutation(self, nums: List[int]) -> int:
        ans = -1
        for i in range(len(nums)):
            if nums[i] != i:
                if ans == -1:
                    ans = nums[i]
                else:
                    ans &= nums[i]
        return 0 if ans == -1 else ans
```

### Go
```go
func sortPermutation(nums []int) int {
    ans := -1
    for i, v := range nums {
        if v != i {
            if ans == -1 {
                ans = v
            } else {
                ans &= v
            }
        }
    }
    if ans == -1 {
        return 0
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

I start by creating a variable that will hold the running AND. I set it to a special value (usually `-1`) so I can later detect whether any misplaced number was ever found.

Then I walk the array with a simple index loop. At each position I compare the value with its index. When they differ I know the number is out of place.

- The first time I meet a misplaced number I simply copy it into the answer variable.  
- Every later misplaced number is combined with the current answer using the bitwise AND operator. This gradually keeps only the bits that are common to all misplaced values.

When the loop ends I check the answer variable. If it is still the original sentinel, the array contained no misplaced numbers, so the answer is `0`. Otherwise the variable already contains the maximum valid `k`.

The same logic appears in every language. The only differences are syntax details such as how arrays are declared, how the bitwise AND is written, and how a missing value is represented. The underlying reasoning stays identical across C++, Java, JavaScript, TypeScript, Python3 and Go.

Edge cases are handled naturally:
- An already sorted array produces the sentinel and returns `0`.
- A completely reversed array produces the AND of every element, which is often `0`.
- Arrays with only a few misplaced pairs correctly capture the AND of just those pairs.

## Examples

**Example 1**  
Input: `nums = [0, 3, 2, 1]`  
Misplaced values: `3` and `1`  
Running AND: `3 & 1 = 1`  
Output: `1`  
Trace: Position 0 and 2 are correct. The only needed swap is between indices 1 and 3, and their AND equals 1, so the maximum k is 1.

**Example 2**  
Input: `nums = [0, 1, 3, 2]`  
Misplaced values: `3` and `2`  
Running AND: `3 & 2 = 2`  
Output: `2`  
Trace: Positions 0 and 1 are already correct. Swapping the last two elements is allowed under k = 2 and produces a sorted array.

**Example 3**  
Input: `nums = [3, 2, 1, 0]`  
Misplaced values: `3, 2, 1, 0`  
Running AND: `3 & 2 & 1 & 0 = 0`  
Output: `0`  
Trace: Every element is out of place. The only k that is present in all of them is 0.

## How to Use / Run Locally

**C++**  
1. Copy the C++ solution into a file named `main.cpp`.  
2. Compile with `g++ -std=c++17 main.cpp -o main`.  
3. Run with `./main` and feed the input according to the problem format.

**Java**  
1. Place the code inside a class named `Solution` in a file `Solution.java`.  
2. Compile with `javac Solution.java`.  
3. Run a small driver that creates an instance and calls `sortPermutation`.

**JavaScript**  
1. Save the function in a file `solution.js`.  
2. Run with Node.js: `node solution.js`.  
3. You can also paste the function into any online JS playground.

**TypeScript**  
1. Save the function in a file `solution.ts`.  
2. Compile with `tsc solution.ts` then run the generated JavaScript, or use `ts-node solution.ts`.

**Python3**  
1. Save the class in a file `solution.py`.  
2. Run with `python3 solution.py`.  
3. Create a small test harness that instantiates `Solution` and calls the method.

**Go**  
1. Place the function in a file `main.go` inside a proper package.  
2. Run with `go run main.go`.  
3. Supply input through standard input or hard-code a few test cases.

## Notes & Optimizations

The solution already runs in linear time and constant extra space, which is optimal for the given constraints (`n ≤ 10^5`).

An alternative approach could build the set of all misplaced values first and then compute their AND, but that would use O(n) extra memory for no gain in asymptotic speed.

Because the numbers form a permutation, each value appears exactly once. This fact is not required by the algorithm, yet it guarantees that the answer is always well-defined.

If the problem later asked for the sequence of swaps rather than only the maximum k, the same AND value would still be the starting point; you would then need to model the allowed moves as a graph and find the connected components.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)