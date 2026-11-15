# Number of Sub-arrays With Odd Sum - LeetCode 1524 Solution

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

We are given an array of positive integers. The task is to count how many contiguous subarrays have an odd sum. Because the count can grow very large, the final answer must be returned modulo 10^9 + 7.

In simple terms, we need an efficient way to find every possible subarray whose elements add up to an odd number and then report that count under the given modulo.

This is a classic prefix-sum parity problem that appears often in coding interviews and competitive programming contests.

## Constraints

- 1 <= arr.length <= 10^5
- 1 <= arr[i] <= 100

The large upper bound on array length forces us to find a linear-time solution. Any approach slower than O(n) will time out.

## Intuition

The first thing that stands out is that we only care about whether a sum is odd or even. The actual value of the sum does not matter.

A subarray sum is odd exactly when the number of odd elements inside it is odd. Looking at prefix sums makes this easy to track. If two prefix sums have different parities, the subarray lying between them must have an odd sum.

So instead of checking every possible subarray, we can keep a running count of how many even-prefix and odd-prefix sums we have seen so far. That single observation turns an O(n^2) idea into a clean O(n) solution.

## Approach

We maintain three pieces of information while walking through the array once:

1. The parity of the prefix sum up to the current index (0 for even, 1 for odd).
2. How many even prefixes we have already seen.
3. How many odd prefixes we have already seen.

We start with one even prefix (the empty prefix before any elements, whose sum is 0).  

For every new element we update the current parity.  
- If the new parity is odd, every earlier even prefix forms a valid odd-sum subarray ending at the current position. We add that count to the answer and then increase the odd-prefix counter.  
- If the new parity is even, we do the opposite: add the current odd-prefix count and increase the even-prefix counter.

All additions are taken modulo 10^9 + 7 so the answer never overflows.

This gives us the exact number of subarrays with odd sum in a single linear pass.

## Data Structures Used

Only a handful of integer variables are required:

- A running parity variable that stores the current prefix sum modulo 2.
- Two counters that track the number of even and odd prefixes seen so far.
- An answer variable that accumulates the final count.

No arrays, maps, or extra data structures are needed. The solution uses constant extra space.

## Operations & Behavior Summary

1. Initialize even-prefix count to 1 and odd-prefix count to 0.  
2. Initialize the answer and the running parity to 0.  
3. For each number in the array:  
   - Add it to the running parity (modulo 2).  
   - If the parity is now odd, add the current even-prefix count to the answer, then increment the odd-prefix count.  
   - If the parity is now even, add the current odd-prefix count to the answer, then increment the even-prefix count.  
4. After the loop finishes, return the answer modulo 10^9 + 7.

## Complexity

| Aspect            | Complexity | Explanation |
|-------------------|------------|-------------|
| Time Complexity   | O(n)       | We examine each of the n elements exactly once and perform a constant amount of work per element. |
| Space Complexity  | O(1)       | Only a fixed number of integer variables are used, independent of the size of the input array. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int numOfSubarrays(vector<int>& arr) {
        const int MOD = 1e9 + 7;
        int even = 1, odd = 0;
        long long ans = 0;
        int sum = 0;
        for (int x : arr) {
            sum = (sum + x) % 2;
            if (sum == 1) {
                ans = (ans + even) % MOD;
                odd++;
            } else {
                ans = (ans + odd) % MOD;
                even++;
            }
        }
        return ans;
    }
};
```

### Java
```java
class Solution {
    public int numOfSubarrays(int[] arr) {
        final int MOD = 1000000007;
        int even = 1, odd = 0;
        long ans = 0;
        int sum = 0;
        for (int x : arr) {
            sum = (sum + x) % 2;
            if (sum == 1) {
                ans = (ans + even) % MOD;
                odd++;
            } else {
                ans = (ans + odd) % MOD;
                even++;
            }
        }
        return (int) ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} arr
 * @return {number}
 */
var numOfSubarrays = function(arr) {
    const MOD = 1e9 + 7;
    let even = 1, odd = 0;
    let ans = 0;
    let sum = 0;
    for (const x of arr) {
        sum = (sum + x) % 2;
        if (sum === 1) {
            ans = (ans + even) % MOD;
            odd++;
        } else {
            ans = (ans + odd) % MOD;
            even++;
        }
    }
    return ans;
};
```

### TypeScript
```typescript
function numOfSubarrays(arr: number[]): number {
    const MOD = 1e9 + 7;
    let even = 1, odd = 0;
    let ans = 0;
    let sum = 0;
    for (const x of arr) {
        sum = (sum + x) % 2;
        if (sum === 1) {
            ans = (ans + even) % MOD;
            odd++;
        } else {
            ans = (ans + odd) % MOD;
            even++;
        }
    }
    return ans;
};
```

### Python3
```python
class Solution:
    def numOfSubarrays(self, arr: list[int]) -> int:
        MOD = 10**9 + 7
        even = 1
        odd = 0
        ans = 0
        s = 0
        for x in arr:
            s = (s + x) % 2
            if s == 1:
                ans = (ans + even) % MOD
                odd += 1
            else:
                ans = (ans + odd) % MOD
                even += 1
        return ans
```

### Go
```go
func numOfSubarrays(arr []int) int {
    const MOD = 1000000007
    even, odd := 1, 0
    ans := 0
    sum := 0
    for _, x := range arr {
        sum = (sum + x) % 2
        if sum == 1 {
            ans = (ans + even) % MOD
            odd++
        } else {
            ans = (ans + odd) % MOD
            even++
        }
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages; only the syntax changes.

We begin by declaring the modulus constant so every addition stays within the required range.  

Next we create two counters. The even counter starts at 1 because the imaginary prefix before the first element has sum 0, which is even. The odd counter starts at 0.  

A variable keeps the parity of the sum seen so far. Another variable accumulates the answer.  

Inside the single loop we update the parity by adding the current array element and taking modulo 2.  

When the parity becomes odd we know that every previously recorded even prefix produces an odd-sum subarray ending at the current index. We therefore add the even counter to the answer and then raise the odd counter by one.  

When the parity stays even we add the odd counter instead and raise the even counter.  

Because the same reasoning applies no matter which language we use, the only differences are how we write the modulo operation, how we declare variables, and how we iterate over the array. The underlying decision process never changes.

Edge-case handling is automatic: an empty array is impossible under the given constraints, and a single-element array is correctly handled by the initial even-prefix count of 1.

## Examples

**Example 1**  
Input: arr = [1, 3, 5]  
Output: 4  

Trace:  
- After 1 the prefix is odd → add 1 (the initial even count) → answer = 1, odd = 1  
- After 3 the prefix is even → add 1 → answer = 2, even = 2  
- After 5 the prefix is odd → add 2 → answer = 4, odd = 2  
The four subarrays with odd sum are [1], [1,3,5], [3], [5].

**Example 2**  
Input: arr = [2, 4, 6]  
Output: 0  

Trace:  
Every prefix stays even, so we always add the odd counter which remains 0. No odd-sum subarray exists.

**Example 3**  
Input: arr = [1, 2, 3, 4, 5, 6, 7]  
Output: 16  

The algorithm correctly counts all 16 subarrays whose sums are odd by maintaining the two running parity counters.

## How to Use / Run Locally

**C++**  
1. Copy the C++ solution into a file named `main.cpp`.  
2. Compile with `g++ -std=c++17 main.cpp -o main`.  
3. Run with `./main`. You may need to add a small driver that reads the array and prints the result.

**Java**  
1. Place the code inside a file named `Solution.java`.  
2. Compile with `javac Solution.java`.  
3. Run with `java Solution` after adding a main method that feeds test cases.

**JavaScript**  
1. Save the function in a file named `solution.js`.  
2. Run with Node.js: `node solution.js`.  
3. Call the function with sample arrays and log the returned value.

**TypeScript**  
1. Save the function in a file named `solution.ts`.  
2. Compile with `tsc solution.ts` (or run directly with `ts-node`).  
3. Execute the generated JavaScript or use ts-node to see the output.

**Python3**  
1. Put the class in a file named `solution.py`.  
2. Run with `python3 solution.py`.  
3. Instantiate the class and call the method with test arrays.

**Go**  
1. Save the function in a file named `main.go`.  
2. Run with `go run main.go`.  
3. Add a main function that prepares the input slice and prints the result.

In every language the core function expects a single integer array and returns an integer. You only need a thin wrapper to supply test data.

## Notes & Optimizations

The solution already runs in optimal linear time and constant extra space.  

One common alternative is to build a full prefix-sum array and then count pairs of different parities. That approach is also O(n) but uses O(n) extra memory, which is unnecessary here.  

Because every element is at most 100, the actual numeric sum never overflows a 64-bit integer even before the modulo, so we can safely keep only the parity.  

If the problem constraints ever grow beyond 10^5, the same algorithm still works; the only change required would be a larger data type for the answer counter in languages that do not promote automatically.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
