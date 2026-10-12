# Apply Operations on Array to Maximize Sum of Squares | LeetCode 2897 Bit Manipulation Greedy Solution

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

You are given a 0-indexed integer array `nums` and a positive integer `k`.  

You can perform the following operation any number of times: pick two different indices `i` and `j`, then at the same time set `nums[i]` to (`nums[i]` AND `nums[j]`) and `nums[j]` to (`nums[i]` OR `nums[j]`).  

After you finish performing operations, you pick exactly `k` numbers from the final array and compute the sum of their squares. Your goal is to make that sum as large as possible. Because the answer can get huge, return it modulo `10^9 + 7`.

In short, the problem asks you to rearrange bits among the numbers (using the given AND/OR operation) so that when you take the `k` largest values you can form and square them, their total is maximized.

## Constraints

- `1 <= k <= nums.length <= 10^5`
- `1 <= nums[i] <= 10^9`

## Intuition

I looked at the operation and saw that it basically moves bits around between two numbers without creating or destroying any bits. The total number of times each bit appears across the whole array stays the same no matter how many times I apply the operation. So the bits are fixed; I just get to decide which numbers end up holding them.

To make the sum of squares as large as possible I want the biggest numbers I can build, because squaring rewards larger values more. That means I should pile the highest bits onto as few numbers as possible and then just pick the top `k` of them.

## Approach

I first count how many numbers have each bit set. There are only 30 bits that matter because the numbers are at most `10^9`.

Then I create an array of `n` zeros. I walk from the highest bit down to the lowest. For every bit, I take the count I recorded and add that bit value to the first “count” positions in my array. Because I always start from the front, the left side of the array keeps getting the high bits first and ends up holding the largest possible numbers.

After all bits are placed I just take the first `k` values, square them, add them up, and take the result modulo `10^9+7`.

## Data Structures Used

- A fixed-size array (or vector) of length 30 that stores the count of how many times each bit is set across the entire input. Size 30 is enough because `nums[i] <= 10^9`.
- A second array of length `n` that holds the newly constructed numbers. We fill this array greedily from the left so the largest possible values sit at the front.

These two structures keep the solution simple and efficient while letting us freely redistribute every bit.

## Operations & Behavior Summary

1. Scan every number and, for each of the 30 bit positions, count how many times that bit is turned on.
2. Create a new array filled with zeros.
3. Starting from bit 29 down to bit 0, look at the count for that bit and turn the bit on in the first `count` slots of the new array.
4. After every bit has been placed, the left side of the array contains the largest numbers we can form.
5. Take the first `k` numbers, square each one, add the squares together, and reduce the total modulo `10^9+7`.
6. Return that final value.

## Complexity

| Complexity       | Value          | Explanation                                                                 |
|------------------|----------------|-----------------------------------------------------------------------------|
| Time Complexity  | O(30 × n)      | We scan the array once for each of the 30 bit positions, so the work is linear in the size of the input times a tiny constant. |
| Space Complexity | O(n)           | We need an array of size n to store the rebuilt numbers. The bit-count array of size 30 is constant and does not grow with n. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int maxSum(vector<int>& nums, int k) {
        const int MOD = 1e9 + 7;
        int n = nums.size();
        vector<int> cnt(30, 0);
        for (int x : nums) {
            for (int b = 0; b < 30; ++b) {
                if (x & (1 << b)) cnt[b]++;
            }
        }
        vector<long long> a(n, 0);
        for (int b = 29; b >= 0; --b) {
            int c = cnt[b];
            for (int i = 0; i < c; ++i) {
                a[i] |= (1LL << b);
            }
        }
        long long ans = 0;
        for (int i = 0; i < k; ++i) {
            ans = (ans + a[i] * a[i]) % MOD;
        }
        return (int)ans;
    }
};
```

### Java

```java
class Solution {
    public int maxSum(List<Integer> nums, int k) {
        final int MOD = 1000000007;
        int n = nums.size();
        int[] cnt = new int[30];
        for (int x : nums) {
            for (int b = 0; b < 30; b++) {
                if ((x & (1 << b)) != 0) cnt[b]++;
            }
        }
        long[] a = new long[n];
        for (int b = 29; b >= 0; b--) {
            int c = cnt[b];
            for (int i = 0; i < c; i++) {
                a[i] |= (1L << b);
            }
        }
        long ans = 0;
        for (int i = 0; i < k; i++) {
            ans = (ans + a[i] * a[i]) % MOD;
        }
        return (int) ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxSum = function(nums, k) {
    const MOD = 1000000007;
    const n = nums.length;
    const cnt = new Array(30).fill(0);
    for (const x of nums) {
        for (let b = 0; b < 30; b++) {
            if (x & (1 << b)) cnt[b]++;
        }
    }
    const a = new Array(n).fill(0);
    for (let b = 29; b >= 0; b--) {
        const c = cnt[b];
        for (let i = 0; i < c; i++) {
            a[i] |= (1 << b);
        }
    }
    let ans = 0;
    for (let i = 0; i < k; i++) {
        ans = (ans + a[i] * a[i]) % MOD;
    }
    return ans;
};
```

### TypeScript

```typescript
function maxSum(nums: number[], k: number): number {
    const MOD = 1000000007;
    const n = nums.length;
    const cnt = new Array(30).fill(0);
    for (const x of nums) {
        for (let b = 0; b < 30; b++) {
            if (x & (1 << b)) cnt[b]++;
        }
    }
    const a = new Array(n).fill(0);
    for (let b = 29; b >= 0; b--) {
        const c = cnt[b];
        for (let i = 0; i < c; i++) {
            a[i] |= (1 << b);
        }
    }
    let ans = 0;
    for (let i = 0; i < k; i++) {
        ans = (ans + a[i] * a[i]) % MOD;
    }
    return ans;
};
```

### Python3

```python
class Solution:
    def maxSum(self, nums: List[int], k: int) -> int:
        MOD = 10**9 + 7
        n = len(nums)
        cnt = [0] * 30
        for x in nums:
            for b in range(30):
                if x & (1 << b):
                    cnt[b] += 1
        a = [0] * n
        for b in range(29, -1, -1):
            c = cnt[b]
            for i in range(c):
                a[i] |= (1 << b)
        ans = 0
        for i in range(k):
            ans = (ans + a[i] * a[i]) % MOD
        return ans
```

### Go

```go
func maxSum(nums []int, k int) int {
    const MOD = 1000000007
    n := len(nums)
    cnt := make([]int, 30)
    for _, x := range nums {
        for b := 0; b < 30; b++ {
            if x&(1<<b) != 0 {
                cnt[b]++
            }
        }
    }
    a := make([]int64, n)
    for b := 29; b >= 0; b-- {
        c := cnt[b]
        for i := 0; i < c; i++ {
            a[i] |= 1 << b
        }
    }
    var ans int64
    for i := 0; i < k; i++ {
        ans = (ans + a[i]*a[i]) % MOD
    }
    return int(ans)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical in every language; only the syntax changes.

I start by making a small array called `cnt` of size 30. Each slot will hold how many numbers currently have that bit turned on.

I walk through every number in the input. For each number I check every bit from 0 to 29. If the bit is set I add one to the matching slot in `cnt`. After this loop finishes I know exactly how many of each bit I am allowed to use.

Next I create another array `a` of size `n` and fill it with zeros. This array will become the numbers I build.

I now loop over the bits from the highest (29) down to the lowest (0). For the current bit I look up how many copies of it I have (the value in `cnt`). I then take the first that-many positions in `a` and turn that bit on in each of them by doing a bitwise OR with the corresponding power of two. Because I always fill from the left side, the numbers at the front keep receiving the higher bits and therefore grow larger than the numbers farther to the right.

When every bit has been placed, the left part of `a` contains the largest numbers I can form. I only need the first `k` of them. I square each one, add the squares together, and take the running total modulo `10^9+7` so the answer stays inside the required range.

That final sum is what I return.

Edge cases are handled naturally: if `k` equals `n` we simply use every rebuilt number; if some bits have a count larger than `n` the extra bits are never placed (they cannot exist); numbers that become zero after redistribution are simply ignored when we pick only the top `k`.

## Examples

**Example 1**

Input: `nums = [2,6,5,8]`, `k = 2`  
Output: `261`

Trace:  

- Bit counts (bits 0-3 for simplicity): bit 1 appears 3 times, bit 2 appears 3 times, bit 3 appears 2 times.  
- We place the highest bits first into the leftmost positions, producing roughly the numbers 15 and 6 (and two zeros).  
- `15² + 6² = 225 + 36 = 261`.

**Example 2**

Input: `nums = [4,5,4,7]`, `k = 3`  
Output: `90`

Trace:  

- The bits already form reasonably large numbers. After redistribution we can obtain 7, 5 and 4.  
- `7² + 5² + 4² = 49 + 25 + 16 = 90`. No operation is even required.

**Example 3 (edge case)**

Input: `nums = [1,1,1]`, `k = 1`  
Output: `1`  
All bits can be piled onto a single number, but the only available bit is the 0-th bit, so the largest value we can form is still 1.

## How to Use / Run Locally

**C++**  

```bash
g++ -std=c++17 -O2 solution.cpp -o solution
./solution
```

**Java**  

```bash
javac Solution.java
java Solution
```

**JavaScript**  

```bash
node solution.js
```

**TypeScript**  

```bash
tsc solution.ts
node solution.js
```

**Python3**  

```bash
python3 solution.py
```

**Go**  

```bash
go run solution.go
```

Each file should contain a simple `main` that reads the input array and `k`, calls the `maxSum` function, and prints the result. You can also paste the function directly into an online judge.

## Notes & Optimizations

- Because we only care about the top `k` numbers, we could stop filling the array after the first `k` positions. The version shown keeps the full array for clarity; the asymptotic complexity stays the same.
- An alternative implementation builds the `k` numbers one by one: for each of the `k` slots it grabs every remaining bit that is still available. Both approaches produce the same answer.
- Watch out for overflow when squaring. Always use 64-bit integers (`long long` / `long` / `int64`) before taking the modulo.
- When `k` is close to `n` the extra zeros at the end do not affect the answer, so the extra work is harmless.
- The same bit-count idea can be reused for many other problems that ask you to redistribute bits under an AND/OR or XOR constraint.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
