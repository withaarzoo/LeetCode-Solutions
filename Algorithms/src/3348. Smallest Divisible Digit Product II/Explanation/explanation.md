# Smallest Divisible Digit Product II

Competitive programming solution for **LeetCode 3348 - Smallest Divisible Digit Product II** using **greedy construction, prime factorization, and dynamic programming**.

This problem is a strong example of how digit DP-style thinking, number theory, and greedy validation work together to build the smallest valid zero-free number. The core idea is to factorize `t`, check whether the target is even possible, and then build the answer from left to right while preserving the original prefix as much as possible.

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)

  * [C++](#c)
  * [Java](#java)
  * [JavaScript](#javascript)
  * [Python3](#python3)
  * [Go](#go)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

You are given:

* a string `num`, which represents a positive integer
* an integer `t`

You need to return the **smallest zero-free number** that is **greater than or equal to `num`** and whose **digit product is divisible by `t`**. If no such number exists, return `"-1"`.

A zero-free number means that none of its digits is `0`. The challenge is to keep the number as small as possible while still making the digit product divisible by `t`. The uploaded notes explain that the solution first checks whether `t` contains only the prime factors `2, 3, 5, 7`, and then greedily constructs the answer with help from a small DP table.

## Constraints

* `2 <= num.length <= 2 × 10^5`
* `num` contains only digits from `'0'` to `'9'`
* `num` does not contain leading zeros
* `1 <= t <= 10^14`

## Intuition

The key observation is simple: digits `1` to `9` can only contribute the prime factors `2`, `3`, `5`, and `7`. So if `t` contains any other prime factor, the answer is impossible right away. This is why the first step is to factorize `t` completely and reject unsupported values early.

The second observation is about building the smallest valid number. We should preserve the left part of `num` as much as possible. If we need to change a digit, we should do it as far to the right as possible. After that, we fill the remaining positions with the smallest digits that still keep the product valid. A tiny DP table helps us quickly check whether the remaining positions are enough to cover the needed factors of `2` and `3`.

## Approach

1. Factorize `t` into counts of `2`, `3`, `5`, and `7`.
2. If `t` still has a remainder greater than `1`, return `"-1"`.
3. Precompute a small DP table that tells us the minimum digits needed to cover any required amount of `2`s and `3`s.
4. Precompute how much each digit from `0` to `9` contributes to the prime factors.
5. Check whether `num` itself is already valid and zero-free.
6. If not, scan from right to left and try increasing one digit.
7. For each candidate replacement digit, check whether the remaining suffix can still satisfy the factor requirements.
8. If yes, build the suffix greedily using the smallest possible digits.
9. If no position works, build a longer valid string from scratch.

This is exactly the greedy-plus-validation pattern used in the uploaded explanation.

## Data Structures Used

* **String**: to store the input number and final answer.
* **Small 2D DP array**: to store the minimum number of digits needed for remaining `2` and `3` factor requirements.
* **Factor lookup arrays**: to quickly know how many `2`s, `3`s, `5`s, and `7`s each digit contributes.
* **Integer counters**: to track the remaining required factor counts while building the answer.

## Operations & Behavior Summary

1. Break `t` into prime factors.
2. Reject impossible cases early.
3. Build the DP table once.
4. Read the original number and see whether it is already valid.
5. If not, move from right to left and try to improve one digit.
6. For each change, verify whether the suffix can still fit the remaining requirements.
7. If it fits, fill the suffix greedily with the smallest digits.
8. If no same-length answer exists, create the smallest valid longer number.

## Complexity

| Metric | Complexity | Explanation                                                                         |
| ------ | ---------- | ----------------------------------------------------------------------------------- |
| Time   | `O(n)`     | The number is scanned a constant number of times, and the DP table is fixed-size.   |
| Space  | `O(n)`     | The final answer string needs linear space; auxiliary structures are constant-size. |

The uploaded notes also describe the same complexity result: linear time overall and linear output space, with the DP table treated as constant size.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string smallestNumber(string num, long long t) {
        int req2 = 0, req3 = 0, req5 = 0, req7 = 0;
        long long temp = t;
        // Divide out all 2s, 3s, 5s, and 7s from the target
        while (temp % 2 == 0) { temp /= 2; req2++; }
        while (temp % 3 == 0) { temp /= 3; req3++; }
        while (temp % 5 == 0) { temp /= 5; req5++; }
        while (temp % 7 == 0) { temp /= 7; req7++; }
        // If the remaining value is > 1, it has invalid prime factors
        if (temp > 1) return "-1";

        // dp[i][j] stores the minimum digits to get AT LEAST i twos and j threes
        int dp[60][40];
        for (int i = 0; i < 60; ++i) {
            for (int j = 0; j < 40; ++j) {
                dp[i][j] = 1e9;
            }
        }
        dp[0][0] = 0;
        
        // Transitions for digits that provide factors of 2 and 3: 2, 3, 4, 6, 8, 9
        int trans[6][2] = {{1, 0}, {0, 1}, {2, 0}, {1, 1}, {3, 0}, {0, 2}};
        for (int i = 0; i < 60; ++i) {
            for (int j = 0; j < 40; ++j) {
                if (dp[i][j] == 1e9) continue;
                for (auto& tr : trans) {
                    // Cap the states to avoid out-of-bounds, since we only need "at least"
                    int ni = min(59, i + tr[0]);
                    int nj = min(39, j + tr[1]);
                    dp[ni][nj] = min(dp[ni][nj], dp[i][j] + 1);
                }
            }
        }
        // Backward propagation to ensure dp[i][j] reflects "at least" i and j factors
        for (int i = 59; i >= 0; --i) {
            for (int j = 39; j >= 0; --j) {
                if (i < 59) dp[i][j] = min(dp[i][j], dp[i + 1][j]);
                if (j < 39) dp[i][j] = min(dp[i][j], dp[i][j + 1]);
            }
        }

        // Precompute the factors provided by each digit 0-9
        int F2[] = {0, 0, 1, 0, 2, 0, 1, 0, 3, 0};
        int F3[] = {0, 0, 0, 1, 0, 0, 1, 0, 0, 2};
        int F5[] = {0, 0, 0, 0, 0, 1, 0, 0, 0, 0};
        int F7[] = {0, 0, 0, 0, 0, 0, 0, 1, 0, 0};

        int n = num.length();
        bool has_zero = false;
        int first_zero = n;
        for (int i = 0; i < n; ++i) {
            if (num[i] == '0') {
                has_zero = true;
                first_zero = i;
                break;
            }
        }

        // Check if the input number itself perfectly satisfies the requirements
        if (!has_zero) {
            int r2 = req2, r3 = req3, r5 = req5, r7 = req7;
            for (char c : num) {
                int d = c - '0';
                r2 = max(0, r2 - F2[d]);
                r3 = max(0, r3 - F3[d]);
                r5 = max(0, r5 - F5[d]);
                r7 = max(0, r7 - F7[d]);
            }
            if (r2 == 0 && r3 == 0 && r5 == 0 && r7 == 0) return num;
        }

        // We can only keep prefixes that occur before the first '0'
        int limit = min(n - 1, first_zero);
        int p2 = 0, p3 = 0, p5 = 0, p7 = 0;
        for (int i = 0; i < limit; ++i) {
            int d = num[i] - '0';
            p2 += F2[d];
            p3 += F3[d];
            p5 += F5[d];
            p7 += F7[d];
        }

        // Scan backwards to find the best rightmost prefix change
        for (int i = limit; i >= 0; --i) {
            int start_d = (num[i] - '0') + 1;
            // Try all possible strictly greater replacements for the current digit
            for (int d = start_d; d <= 9; ++d) {
                int n2 = max(0, req2 - p2 - F2[d]);
                int n3 = max(0, req3 - p3 - F3[d]);
                int n5 = max(0, req5 - p5 - F5[d]);
                int n7 = max(0, req7 - p7 - F7[d]);
                int L = n - 1 - i;
                
                // If remaining length L can fit all required factors, we lock it in
                if (n7 + n5 + dp[n2][n3] <= L) {
                    string ans = num.substr(0, i) + to_string(d);
                    int rem2 = n2, rem3 = n3, rem5 = n5, rem7 = n7;
                    // Build suffix greedily with smallest valid characters
                    for (int pos = 0; pos < L; ++pos) {
                        for (int x = 1; x <= 9; ++x) {
                            int nn2 = max(0, rem2 - F2[x]);
                            int nn3 = max(0, rem3 - F3[x]);
                            int nn5 = max(0, rem5 - F5[x]);
                            int nn7 = max(0, rem7 - F7[x]);
                            if (nn7 + nn5 + dp[nn2][nn3] <= L - 1 - pos) {
                                ans += to_string(x);
                                rem2 = nn2; rem3 = nn3; rem5 = nn5; rem7 = nn7;
                                break;
                            }
                        }
                    }
                    return ans;
                }
            }
            // Strip the contribution of the digit we are passing backwards over
            if (i > 0) {
                int d = num[i - 1] - '0';
                p2 -= F2[d];
                p3 -= F3[d];
                p5 -= F5[d];
                p7 -= F7[d];
            }
        }

        // If no matching prefix exists, construct an entirely new string length M
        int min_len_needed = req7 + req5 + dp[req2][req3];
        int M = max(n + 1, min_len_needed);
        string ans = "";
        int rem2 = req2, rem3 = req3, rem5 = req5, rem7 = req7;
        
        // Loop over the new length M and place the smallest valid digits left to right
        for (int pos = 0; pos < M; ++pos) {
            for (int x = 1; x <= 9; ++x) {
                int nn2 = max(0, rem2 - F2[x]);
                int nn3 = max(0, rem3 - F3[x]);
                int nn5 = max(0, rem5 - F5[x]);
                int nn7 = max(0, rem7 - F7[x]);
                if (nn7 + nn5 + dp[nn2][nn3] <= M - 1 - pos) {
                    ans += to_string(x);
                    rem2 = nn2; rem3 = nn3; rem5 = nn5; rem7 = nn7;
                    break;
                }
            }
        }
        return ans;
    }
};
```

### Java

```java
class Solution {
    public String smallestNumber(String num, long t) {
        int req2 = 0, req3 = 0, req5 = 0, req7 = 0;
        long temp = t;
        // Strip out allowed prime factors
        while (temp % 2 == 0) { temp /= 2; req2++; }
        while (temp % 3 == 0) { temp /= 3; req3++; }
        while (temp % 5 == 0) { temp /= 5; req5++; }
        while (temp % 7 == 0) { temp /= 7; req7++; }
        // Detect unsupported prime factors immediately
        if (temp > 1) return "-1";

        int[][] dp = new int[60][40];
        for (int i = 0; i < 60; ++i) {
            for (int j = 0; j < 40; ++j) {
                dp[i][j] = 1000000000;
            }
        }
        dp[0][0] = 0;
        
        // Apply digit transitions that generate 2s and 3s
        int[][] trans = {{1, 0}, {0, 1}, {2, 0}, {1, 1}, {3, 0}, {0, 2}};
        for (int i = 0; i < 60; ++i) {
            for (int j = 0; j < 40; ++j) {
                if (dp[i][j] == 1000000000) continue;
                for (int[] tr : trans) {
                    int ni = Math.min(59, i + tr[0]);
                    int nj = Math.min(39, j + tr[1]);
                    dp[ni][nj] = Math.min(dp[ni][nj], dp[i][j] + 1);
                }
            }
        }
        
        // Propagate backwards so looking up dp[i][j] covers needing AT LEAST i and j
        for (int i = 59; i >= 0; --i) {
            for (int j = 39; j >= 0; --j) {
                if (i < 59) dp[i][j] = Math.min(dp[i][j], dp[i + 1][j]);
                if (j < 39) dp[i][j] = Math.min(dp[i][j], dp[i][j + 1]);
            }
        }

        int[] F2 = {0, 0, 1, 0, 2, 0, 1, 0, 3, 0};
        int[] F3 = {0, 0, 0, 1, 0, 0, 1, 0, 0, 2};
        int[] F5 = {0, 0, 0, 0, 0, 1, 0, 0, 0, 0};
        int[] F7 = {0, 0, 0, 0, 0, 0, 0, 1, 0, 0};

        int n = num.length();
        boolean hasZero = false;
        int firstZero = n;
        for (int i = 0; i < n; ++i) {
            if (num.charAt(i) == '0') {
                hasZero = true;
                firstZero = i;
                break;
            }
        }

        // Ensure the input naturally fulfills the requirements without any changes
        if (!hasZero) {
            int r2 = req2, r3 = req3, r5 = req5, r7 = req7;
            for (int i = 0; i < n; i++) {
                int d = num.charAt(i) - '0';
                r2 = Math.max(0, r2 - F2[d]);
                r3 = Math.max(0, r3 - F3[d]);
                r5 = Math.max(0, r5 - F5[d]);
                r7 = Math.max(0, r7 - F7[d]);
            }
            if (r2 == 0 && r3 == 0 && r5 == 0 && r7 == 0) return num;
        }

        // Limit the scope to indices before a '0' occurs
        int limit = Math.min(n - 1, firstZero);
        int p2 = 0, p3 = 0, p5 = 0, p7 = 0;
        for (int i = 0; i < limit; ++i) {
            int d = num.charAt(i) - '0';
            p2 += F2[d];
            p3 += F3[d];
            p5 += F5[d];
            p7 += F7[d];
        }

        // Step backwards to modify the rightmost possible viable digit
        for (int i = limit; i >= 0; --i) {
            int startD = (num.charAt(i) - '0') + 1;
            for (int d = startD; d <= 9; ++d) {
                int n2 = Math.max(0, req2 - p2 - F2[d]);
                int n3 = Math.max(0, req3 - p3 - F3[d]);
                int n5 = Math.max(0, req5 - p5 - F5[d]);
                int n7 = Math.max(0, req7 - p7 - F7[d]);
                int L = n - 1 - i;
                
                // If factors fit, finalize string dynamically
                if (n7 + n5 + dp[n2][n3] <= L) {
                    StringBuilder ans = new StringBuilder(num.substring(0, i));
                    ans.append(d);
                    int rem2 = n2, rem3 = n3, rem5 = n5, rem7 = n7;
                    // Choose lexicographically smallest valid digit iteratively
                    for (int pos = 0; pos < L; ++pos) {
                        for (int x = 1; x <= 9; ++x) {
                            int nn2 = Math.max(0, rem2 - F2[x]);
                            int nn3 = Math.max(0, rem3 - F3[x]);
                            int nn5 = Math.max(0, rem5 - F5[x]);
                            int nn7 = Math.max(0, rem7 - F7[x]);
                            if (nn7 + nn5 + dp[nn2][nn3] <= L - 1 - pos) {
                                ans.append(x);
                                rem2 = nn2; rem3 = nn3; rem5 = nn5; rem7 = nn7;
                                break;
                            }
                        }
                    }
                    return ans.toString();
                }
            }
            if (i > 0) {
                int d = num.charAt(i - 1) - '0';
                p2 -= F2[d];
                p3 -= F3[d];
                p5 -= F5[d];
                p7 -= F7[d];
            }
        }

        // Start from scratch if no modified prefix fits the target constraint
        int minLenNeeded = req7 + req5 + dp[req2][req3];
        int M = Math.max(n + 1, minLenNeeded);
        StringBuilder ans = new StringBuilder();
        int rem2 = req2, rem3 = req3, rem5 = req5, rem7 = req7;
        
        for (int pos = 0; pos < M; ++pos) {
            for (int x = 1; x <= 9; ++x) {
                int nn2 = Math.max(0, rem2 - F2[x]);
                int nn3 = Math.max(0, rem3 - F3[x]);
                int nn5 = Math.max(0, rem5 - F5[x]);
                int nn7 = Math.max(0, rem7 - F7[x]);
                if (nn7 + nn5 + dp[nn2][nn3] <= M - 1 - pos) {
                    ans.append(x);
                    rem2 = nn2; rem3 = nn3; rem5 = nn5; rem7 = nn7;
                    break;
                }
            }
        }
        return ans.toString();
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} num
 * @param {number} t
 * @return {string}
 */
var smallestNumber = function(num, t) {
    let req2 = 0, req3 = 0, req5 = 0, req7 = 0;
    let temp = t;
    // Map initial factorization needed from the number
    while (temp % 2 === 0) { temp /= 2; req2++; }
    while (temp % 3 === 0) { temp /= 3; req3++; }
    while (temp % 5 === 0) { temp /= 5; req5++; }
    while (temp % 7 === 0) { temp /= 7; req7++; }
    // Abort if target cannot be formed using digits 1-9
    if (temp > 1) return "-1";

    const dp = Array.from({ length: 60 }, () => Array(40).fill(1e9));
    dp[0][0] = 0;
    
    // Test base digit additions to build forward map of dependencies
    const trans = [[1, 0], [0, 1], [2, 0], [1, 1], [3, 0], [0, 2]];
    for (let i = 0; i < 60; ++i) {
        for (let j = 0; j < 40; ++j) {
            if (dp[i][j] === 1e9) continue;
            for (const tr of trans) {
                const ni = Math.min(59, i + tr[0]);
                const nj = Math.min(39, j + tr[1]);
                dp[ni][nj] = Math.min(dp[ni][nj], dp[i][j] + 1);
            }
        }
    }
    
    // Resolve states iteratively to handle 'at least' thresholding properly
    for (let i = 59; i >= 0; --i) {
        for (let j = 39; j >= 0; --j) {
            if (i < 59) dp[i][j] = Math.min(dp[i][j], dp[i + 1][j]);
            if (j < 39) dp[i][j] = Math.min(dp[i][j], dp[i][j + 1]);
        }
    }

    const F2 = [0, 0, 1, 0, 2, 0, 1, 0, 3, 0];
    const F3 = [0, 0, 0, 1, 0, 0, 1, 0, 0, 2];
    const F5 = [0, 0, 0, 0, 0, 1, 0, 0, 0, 0];
    const F7 = [0, 0, 0, 0, 0, 0, 0, 1, 0, 0];

    const n = num.length;
    let hasZero = false;
    let firstZero = n;
    for (let i = 0; i < n; ++i) {
        if (num[i] === '0') {
            hasZero = true;
            firstZero = i;
            break;
        }
    }

    // Verify whether zero-free original num evaluates as sufficient
    if (!hasZero) {
        let r2 = req2, r3 = req3, r5 = req5, r7 = req7;
        for (let char of num) {
            let d = Number(char);
            r2 = Math.max(0, r2 - F2[d]);
            r3 = Math.max(0, r3 - F3[d]);
            r5 = Math.max(0, r5 - F5[d]);
            r7 = Math.max(0, r7 - F7[d]);
        }
        if (r2 === 0 && r3 === 0 && r5 === 0 && r7 === 0) return num;
    }

    // Capture intact factor counts strictly ahead of zeros
    const limit = Math.min(n - 1, firstZero);
    let p2 = 0, p3 = 0, p5 = 0, p7 = 0;
    for (let i = 0; i < limit; ++i) {
        let d = Number(num[i]);
        p2 += F2[d];
        p3 += F3[d];
        p5 += F5[d];
        p7 += F7[d];
    }

    // Rewind loop finding longest matching prefix segment safely incrementable
    for (let i = limit; i >= 0; --i) {
        let startD = Number(num[i]) + 1;
        for (let d = startD; d <= 9; ++d) {
            let n2 = Math.max(0, req2 - p2 - F2[d]);
            let n3 = Math.max(0, req3 - p3 - F3[d]);
            let n5 = Math.max(0, req5 - p5 - F5[d]);
            let n7 = Math.max(0, req7 - p7 - F7[d]);
            let L = n - 1 - i;
            
            // Assemble output string cleanly checking remaining capacity DP maps
            if (n7 + n5 + dp[n2][n3] <= L) {
                let ans = num.substring(0, i) + d.toString();
                let rem2 = n2, rem3 = n3, rem5 = n5, rem7 = n7;
                for (let pos = 0; pos < L; ++pos) {
                    for (let x = 1; x <= 9; ++x) {
                        let nn2 = Math.max(0, rem2 - F2[x]);
                        let nn3 = Math.max(0, rem3 - F3[x]);
                        let nn5 = Math.max(0, rem5 - F5[x]);
                        let nn7 = Math.max(0, rem7 - F7[x]);
                        if (nn7 + nn5 + dp[nn2][nn3] <= L - 1 - pos) {
                            ans += x.toString();
                            rem2 = nn2; rem3 = nn3; rem5 = nn5; rem7 = nn7;
                            break;
                        }
                    }
                }
                return ans;
            }
        }
        if (i > 0) {
            let d = Number(num[i - 1]);
            p2 -= F2[d];
            p3 -= F3[d];
            p5 -= F5[d];
            p7 -= F7[d];
        }
    }

    // Force build longer string assuming original length bounds are exhausted
    let minLenNeeded = req7 + req5 + dp[req2][req3];
    let M = Math.max(n + 1, minLenNeeded);
    let ans = "";
    let rem2 = req2, rem3 = req3, rem5 = req5, rem7 = req7;
    
    for (let pos = 0; pos < M; ++pos) {
        for (let x = 1; x <= 9; ++x) {
            let nn2 = Math.max(0, rem2 - F2[x]);
            let nn3 = Math.max(0, rem3 - F3[x]);
            let nn5 = Math.max(0, rem5 - F5[x]);
            let nn7 = Math.max(0, rem7 - F7[x]);
            if (nn7 + nn5 + dp[nn2][nn3] <= M - 1 - pos) {
                ans += x.toString();
                rem2 = nn2; rem3 = nn3; rem5 = nn5; rem7 = nn7;
                break;
            }
        }
    }
    return ans;
};
```

### Python3

```python
class Solution:
    def smallestNumber(self, num: str, t: int) -> str:
        req2 = req3 = req5 = req7 = 0
        temp = t
        # Remove primary factors to detect impossible structures natively
        while temp % 2 == 0:
            temp //= 2
            req2 += 1
        while temp % 3 == 0:
            temp //= 3
            req3 += 1
        while temp % 5 == 0:
            temp //= 5
            req5 += 1
        while temp % 7 == 0:
            temp //= 7
            req7 += 1
        if temp > 1: return "-1"

        dp = [[float('inf')] * 40 for _ in range(60)]
        dp[0][0] = 0
        
        # Track shortest path generating required factors explicitly via base digits
        trans = [(1, 0), (0, 1), (2, 0), (1, 1), (3, 0), (0, 2)]
        for i in range(60):
            for j in range(40):
                if dp[i][j] == float('inf'):
                    continue
                for d2, d3 in trans:
                    ni = min(59, i + d2)
                    nj = min(39, j + d3)
                    dp[ni][nj] = min(dp[ni][nj], dp[i][j] + 1)
                    
        # Retroactive updates enforce property solving "minimum digits for AT LEAST i, j factors"
        for i in range(59, -1, -1):
            for j in range(39, -1, -1):
                if i < 59:
                    dp[i][j] = min(dp[i][j], dp[i + 1][j])
                if j < 39:
                    dp[i][j] = min(dp[i][j], dp[i][j + 1])

        F2 = [0, 0, 1, 0, 2, 0, 1, 0, 3, 0]
        F3 = [0, 0, 0, 1, 0, 0, 1, 0, 0, 2]
        F5 = [0, 0, 0, 0, 0, 1, 0, 0, 0, 0]
        F7 = [0, 0, 0, 0, 0, 0, 0, 1, 0, 0]

        n = len(num)
        has_zero = False
        first_zero = n
        for idx, char in enumerate(num):
            if char == '0':
                has_zero = True
                first_zero = idx
                break

        # If zero-free initially, confirm requirement checks directly
        if not has_zero:
            r2, r3, r5, r7 = req2, req3, req5, req7
            for char in num:
                d = int(char)
                r2 = max(0, r2 - F2[d])
                r3 = max(0, r3 - F3[d])
                r5 = max(0, r5 - F5[d])
                r7 = max(0, r7 - F7[d])
            if r2 == 0 and r3 == 0 and r5 == 0 and r7 == 0:
                return num

        # Cap iteration cleanly evaluating prefixes
        limit = min(n - 1, first_zero)
        p2 = p3 = p5 = p7 = 0
        for i in range(limit):
            d = int(num[i])
            p2 += F2[d]
            p3 += F3[d]
            p5 += F5[d]
            p7 += F7[d]

        # Scan inverse array trying incrementally superior rightmost numbers safely
        for i in range(limit, -1, -1):
            start_d = int(num[i]) + 1
            for d in range(start_d, 10):
                n2 = max(0, req2 - p2 - F2[d])
                n3 = max(0, req3 - p3 - F3[d])
                n5 = max(0, req5 - p5 - F5[d])
                n7 = max(0, req7 - p7 - F7[d])
                L = n - 1 - i
                
                # Check spatial fit mapping remainder components seamlessly
                if n7 + n5 + dp[n2][n3] <= L:
                    ans_list = list(num[:i]) + [str(d)]
                    rem2, rem3, rem5, rem7 = n2, n3, n5, n7
                    # Insert minimum variables dynamically left-to-right filling constraints
                    for pos in range(L):
                        for x in range(1, 10):
                            nn2 = max(0, rem2 - F2[x])
                            nn3 = max(0, rem3 - F3[x])
                            nn5 = max(0, rem5 - F5[x])
                            nn7 = max(0, rem7 - F7[x])
                            if nn7 + nn5 + dp[nn2][nn3] <= L - 1 - pos:
                                ans_list.append(str(x))
                                rem2, rem3, rem5, rem7 = nn2, nn3, nn5, nn7
                                break
                    return "".join(ans_list)
            
            if i > 0:
                d = int(num[i - 1])
                p2 -= F2[d]
                p3 -= F3[d]
                p5 -= F5[d]
                p7 -= F7[d]

        # Push dimension length when base string fundamentally lacks sufficient scale
        min_len_needed = req7 + req5 + dp[req2][req3]
        M = max(n + 1, min_len_needed)
        ans_list = []
        rem2, rem3, rem5, rem7 = req2, req3, req5, req7
        
        for pos in range(M):
            for x in range(1, 10):
                nn2 = max(0, rem2 - F2[x])
                nn3 = max(0, rem3 - F3[x])
                nn5 = max(0, rem5 - F5[x])
                nn7 = max(0, rem7 - F7[x])
                if nn7 + nn5 + dp[nn2][nn3] <= M - 1 - pos:
                    ans_list.append(str(x))
                    rem2, rem3, rem5, rem7 = nn2, nn3, nn5, nn7
                    break
        return "".join(ans_list)
```

### Go

```go
func smallestNumber(num string, t int64) string {
    var req2, req3, req5, req7 int
    temp := t
    // Decode factors structurally avoiding manual string checks prematurely
    for temp%2 == 0 { temp /= 2; req2++ }
    for temp%3 == 0 { temp /= 3; req3++ }
    for temp%5 == 0 { temp /= 5; req5++ }
    for temp%7 == 0 { temp /= 7; req7++ }
    // Abandon checks if unrelated factors appear internally
    if temp > 1 {
        return "-1"
    }

    var dp [60][40]int
    for i := 0; i < 60; i++ {
        for j := 0; j < 40; j++ {
            dp[i][j] = 1000000000
        }
    }
    dp[0][0] = 0
    
    // Evaluate transition variables actively generating combinations
    trans := [6][2]int{{1, 0}, {0, 1}, {2, 0}, {1, 1}, {3, 0}, {0, 2}}
    for i := 0; i < 60; i++ {
        for j := 0; j < 40; j++ {
            if dp[i][j] == 1000000000 {
                continue
            }
            for _, tr := range trans {
                ni := i + tr[0]
                if ni > 59 { ni = 59 }
                nj := j + tr[1]
                if nj > 39 { nj = 39 }
                if dp[i][j]+1 < dp[ni][nj] {
                    dp[ni][nj] = dp[i][j] + 1
                }
            }
        }
    }
    
    // Smooth states forcing condition values safely representing thresholds
    for i := 59; i >= 0; i-- {
        for j := 39; j >= 0; j-- {
            if i < 59 && dp[i+1][j] < dp[i][j] {
                dp[i][j] = dp[i+1][j]
            }
            if j < 39 && dp[i][j+1] < dp[i][j] {
                dp[i][j] = dp[i][j+1]
            }
        }
    }

    F2 := []int{0, 0, 1, 0, 2, 0, 1, 0, 3, 0}
    F3 := []int{0, 0, 0, 1, 0, 0, 1, 0, 0, 2}
    F5 := []int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0}
    F7 := []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}

    n := len(num)
    hasZero := false
    firstZero := n
    for i := 0; i < n; i++ {
        if num[i] == '0' {
            hasZero = true
            firstZero = i
            break
        }
    }

    maxF := func(a, b int) int { if a > b { return a }; return b }
    minF := func(a, b int) int { if a < b { return a }; return b }

    // Examine input layout skipping logic checks if target validates strictly
    if !hasZero {
        r2, r3, r5, r7 := req2, req3, req5, req7
        for i := 0; i < n; i++ {
            d := int(num[i] - '0')
            r2 = maxF(0, r2-F2[d])
            r3 = maxF(0, r3-F3[d])
            r5 = maxF(0, r5-F5[d])
            r7 = maxF(0, r7-F7[d])
        }
        if r2 == 0 && r3 == 0 && r5 == 0 && r7 == 0 {
            return num
        }
    }

    // Capture prefix lengths retaining existing components accurately
    limit := minF(n-1, firstZero)
    var p2, p3, p5, p7 int
    for i := 0; i < limit; i++ {
        d := int(num[i] - '0')
        p2 += F2[d]
        p3 += F3[d]
        p5 += F5[d]
        p7 += F7[d]
    }

    // Traverse downward checking edits strictly maintaining factor logic
    for i := limit; i >= 0; i-- {
        startD := int(num[i]-'0') + 1
        for d := startD; d <= 9; d++ {
            n2 := maxF(0, req2-p2-F2[d])
            n3 := maxF(0, req3-p3-F3[d])
            n5 := maxF(0, req5-p5-F5[d])
            n7 := maxF(0, req7-p7-F7[d])
            L := n - 1 - i

            // Implement string builder executing smallest sequence safely
            if n7+n5+dp[n2][n3] <= L {
                ans := []byte(num[:i])
                ans = append(ans, byte(d+'0'))
                rem2, rem3, rem5, rem7 := n2, n3, n5, n7
                
                for pos := 0; pos < L; pos++ {
                    for x := 1; x <= 9; x++ {
                        nn2 := maxF(0, rem2-F2[x])
                        nn3 := maxF(0, rem3-F3[x])
                        nn5 := maxF(0, rem5-F5[x])
                        nn7 := maxF(0, rem7-F7[x])
                        if nn7+nn5+dp[nn2][nn3] <= L-1-pos {
                            ans = append(ans, byte(x+'0'))
                            rem2, rem3, rem5, rem7 = nn2, nn3, nn5, nn7
                            break
                        }
                    }
                }
                return string(ans)
            }
        }
        if i > 0 {
            d := int(num[i-1] - '0')
            p2 -= F2[d]
            p3 -= F3[d]
            p5 -= F5[d]
            p7 -= F7[d]
        }
    }

    // Create string sequence expanding boundary lengths dynamically
    minLenNeeded := req7 + req5 + dp[req2][req3]
    M := maxF(n+1, minLenNeeded)
    var ans []byte
    rem2, rem3, rem5, rem7 := req2, req3, req5, req7
    
    for pos := 0; pos < M; pos++ {
        for x := 1; x <= 9; x++ {
            nn2 := maxF(0, rem2-F2[x])
            nn3 := maxF(0, rem3-F3[x])
            nn5 := maxF(0, rem5-F5[x])
            nn7 := maxF(0, rem7-F7[x])
            if nn7+nn5+dp[nn2][nn3] <= M-1-pos {
                ans = append(ans, byte(x+'0'))
                rem2, rem3, rem5, rem7 = nn2, nn3, nn5, nn7
                break
            }
        }
    }
    return string(ans)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five languages.

First, we split `t` into powers of `2`, `3`, `5`, and `7`. If anything remains after that, the problem has no solution.

Next, we build a DP table for the `2` and `3` requirements. This table is small and fixed, so it is fast to compute and reuse. The notes explain that this table is used to check whether the remaining suffix length is enough to satisfy the remaining factors.

Then we check whether `num` is already a valid zero-free answer. If it is, we return it directly.

If not, we move backward and try to increase one digit. The rightmost possible change is always better because it keeps the prefix as small as possible. After choosing a larger digit, we test whether the suffix still has enough room for the remaining factor requirements.

If the suffix is feasible, we fill it greedily from left to right. At each position, we try digits from `1` to `9` and pick the smallest digit that still leaves a valid remainder.

If no same-length answer can be formed, we increase the total length and build a fresh string from scratch using the same greedy check.

That is the complete reasoning behind the solutions in all five languages. The uploaded walkthrough describes the same right-to-left replacement strategy and the same greedy suffix construction.

## Examples

### Example 1

**Input**

```text
num = "1234"
t = 256
```

**Output**

```text
1488
```

**Trace**

* `256 = 2^8`
* The number must be zero-free and its digit product must cover eight `2`s.
* The algorithm keeps the prefix as much as possible.
* It changes the smallest right-side part that allows the remainder to stay valid.
* The greedy suffix becomes `1488`.

This matches the example shown in the problem notes.

### Example 2

**Input**

```text
num = "12355"
t = 50
```

**Output**

```text
12355
```

**Trace**

* `50 = 2 × 5^2`
* The original number is already zero-free.
* Its digit product already has enough factors.
* So the answer is the same string.

This is also one of the examples shown in the notes.

### Example 3

**Input**

```text
num = "11111"
t = 26
```

**Output**

```text
-1
```

**Trace**

* `26 = 2 × 13`
* The factor `13` cannot be produced by digits `1` to `9`
* So no valid answer exists.

This impossibility check is exactly the kind of early rejection described in the explanation.

## How to Use / Run Locally

### C++

```bash
g++ solution.cpp -o solution
./solution
```

### Java

```bash
javac Solution.java
java Solution
```

### JavaScript

```bash
node solution.js
```

### Python3

```bash
python3 solution.py
```

### Go

```bash
go run solution.go
```

## Notes & Optimizations

* Reject invalid `t` early if it contains any prime factor greater than `7`.
* Use the DP table only for `2` and `3`; `5` and `7` can be handled directly because they each require specific digits.
* Try to change digits from right to left to keep the number as small as possible.
* Build the suffix greedily from left to right.
* If the original length cannot work, construct a longer valid number.
* This solution is efficient enough for very large inputs because all heavy work is reduced to constant-size DP and linear scans.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
