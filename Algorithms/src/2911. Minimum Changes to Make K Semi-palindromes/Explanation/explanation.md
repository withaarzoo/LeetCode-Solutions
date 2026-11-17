# Minimum Changes to Make K Semi-palindromes - LeetCode 2911 Solution

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

You are given a string `s` and an integer `k`. Your task is to split `s` into exactly `k` non-empty substrings. For every substring you must make it a semi-palindrome by changing the fewest letters possible. The goal is to return the smallest total number of letter changes needed across all `k` pieces.

A semi-palindrome is a string that can be split into groups based on a repeating pattern of length `d` (where `d` is a proper divisor of the string length). Each of those groups must itself form a regular palindrome.

The problem asks for the minimum letter changes required to achieve such a partition.

## Constraints

- `2 <= s.length <= 200`
- `1 <= k <= s.length / 2`
- `s` contains only lowercase English letters

## Intuition

The length of the string is at most 200, so checking every possible substring is realistic. Once I know the cheapest way to turn any substring into a semi-palindrome, the original problem becomes a classic “split into exactly k pieces” dynamic-programming task. That observation is the starting point of the whole solution.

## Approach

First I build a cost table. For every substring `s[i..j]` I try every valid divisor `d` of its length. For each divisor I examine the character groups that share the same residue modulo `d` and count how many positions must change so that each group becomes a true palindrome. I keep the smallest count among all valid `d` and store it as the cost of that substring.

After the cost table is ready I run a second dynamic-programming stage. `dp[i][p]` stores the minimum total cost of turning the first `i` characters into exactly `p` semi-palindromes. The recurrence simply tries every possible place where the last piece can start and adds the precomputed cost of that piece to the best answer for the prefix.

The final answer sits in `dp[n][k]`.

## Data Structures Used

- Two-dimensional integer array `cost[n][n]` – stores the minimum changes needed for every substring.
- Two-dimensional integer array `dp[n+1][k+1]` – stores the best partition costs for prefixes of the string.
- Simple integer variables for length, divisors, and temporary costs. No fancy containers are required because the size limits are small.

## Operations & Behavior Summary

1. Initialise the cost table with a large “impossible” value.
2. For every substring of length at least 2:
   - Enumerate every proper divisor `d`.
   - For each residue class, walk the corresponding characters and count mismatches that break the palindrome property.
   - Record the cheapest divisor cost for that substring.
3. Initialise the DP table: zero cost for zero characters and zero pieces; everything else starts as impossible.
4. For increasing string prefixes and increasing piece counts, try every legal previous cut point and keep the minimum total cost.
5. Return the entry that corresponds to the whole string and exactly `k` pieces.

## Complexity

| Metric            | Value          | Explanation |
|-------------------|----------------|-------------|
| Time Complexity   | O(n³ + k·n²)  | Building the cost table dominates with O(n³) work; the partition DP adds O(k·n²). |
| Space Complexity  | O(n²)         | Both the cost table and the DP table use quadratic space. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int minimumChanges(string s, int k) {
        int n = s.size();
        const int INF = 1e9;
        vector<vector<int>> cost(n, vector<int>(n, INF));
        for (int i = 0; i < n; ++i) {
            for (int j = i + 1; j < n; ++j) {
                int L = j - i + 1;
                int best = INF;
                for (int d = 1; d < L; ++d) {
                    if (L % d != 0) continue;
                    int cur = 0;
                    int m = L / d;
                    for (int r = 0; r < d; ++r) {
                        for (int t = 0; t < m / 2; ++t) {
                            if (s[i + r + t * d] != s[i + r + (m - 1 - t) * d])
                                ++cur;
                        }
                    }
                    best = min(best, cur);
                }
                cost[i][j] = best;
            }
        }
        vector<vector<int>> dp(n + 1, vector<int>(k + 1, INF));
        dp[0][0] = 0;
        for (int i = 1; i <= n; ++i) {
            for (int p = 1; p <= k; ++p) {
                for (int prev = 0; prev <= i - 2; ++prev) {
                    if (dp[prev][p - 1] != INF && cost[prev][i - 1] != INF)
                        dp[i][p] = min(dp[i][p], dp[prev][p - 1] + cost[prev][i - 1]);
                }
            }
        }
        return dp[n][k];
    }
};
```

### Java
```java
class Solution {
    public int minimumChanges(String s, int k) {
        int n = s.length();
        final int INF = 1000000000;
        int[][] cost = new int[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) cost[i][j] = INF;
        }
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                int L = j - i + 1;
                int best = INF;
                for (int d = 1; d < L; d++) {
                    if (L % d != 0) continue;
                    int cur = 0;
                    int m = L / d;
                    for (int r = 0; r < d; r++) {
                        for (int t = 0; t < m / 2; t++) {
                            if (s.charAt(i + r + t * d) != s.charAt(i + r + (m - 1 - t) * d))
                                cur++;
                        }
                    }
                    best = Math.min(best, cur);
                }
                cost[i][j] = best;
            }
        }
        int[][] dp = new int[n + 1][k + 1];
        for (int i = 0; i <= n; i++) {
            for (int p = 0; p <= k; p++) dp[i][p] = INF;
        }
        dp[0][0] = 0;
        for (int i = 1; i <= n; i++) {
            for (int p = 1; p <= k; p++) {
                for (int prev = 0; prev <= i - 2; prev++) {
                    if (dp[prev][p - 1] != INF && cost[prev][i - 1] != INF)
                        dp[i][p] = Math.min(dp[i][p], dp[prev][p - 1] + cost[prev][i - 1]);
                }
            }
        }
        return dp[n][k];
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @param {number} k
 * @return {number}
 */
var minimumChanges = function(s, k) {
    const n = s.length;
    const INF = 1e9;
    const cost = Array.from({length: n}, () => Array(n).fill(INF));
    for (let i = 0; i < n; i++) {
        for (let j = i + 1; j < n; j++) {
            const L = j - i + 1;
            let best = INF;
            for (let d = 1; d < L; d++) {
                if (L % d !== 0) continue;
                let cur = 0;
                const m = L / d;
                for (let r = 0; r < d; r++) {
                    for (let t = 0; t < Math.floor(m / 2); t++) {
                        if (s[i + r + t * d] !== s[i + r + (m - 1 - t) * d])
                            cur++;
                    }
                }
                best = Math.min(best, cur);
            }
            cost[i][j] = best;
        }
    }
    const dp = Array.from({length: n + 1}, () => Array(k + 1).fill(INF));
    dp[0][0] = 0;
    for (let i = 1; i <= n; i++) {
        for (let p = 1; p <= k; p++) {
            for (let prev = 0; prev <= i - 2; prev++) {
                if (dp[prev][p - 1] !== INF && cost[prev][i - 1] !== INF)
                    dp[i][p] = Math.min(dp[i][p], dp[prev][p - 1] + cost[prev][i - 1]);
            }
        }
    }
    return dp[n][k];
};
```

### TypeScript
```typescript
function minimumChanges(s: string, k: number): number {
    const n = s.length;
    const INF = 1e9;
    const cost: number[][] = Array.from({length: n}, () => Array(n).fill(INF));
    for (let i = 0; i < n; i++) {
        for (let j = i + 1; j < n; j++) {
            const L = j - i + 1;
            let best = INF;
            for (let d = 1; d < L; d++) {
                if (L % d !== 0) continue;
                let cur = 0;
                const m = L / d;
                for (let r = 0; r < d; r++) {
                    for (let t = 0; t < Math.floor(m / 2); t++) {
                        if (s[i + r + t * d] !== s[i + r + (m - 1 - t) * d])
                            cur++;
                    }
                }
                best = Math.min(best, cur);
            }
            cost[i][j] = best;
        }
    }
    const dp: number[][] = Array.from({length: n + 1}, () => Array(k + 1).fill(INF));
    dp[0][0] = 0;
    for (let i = 1; i <= n; i++) {
        for (let p = 1; p <= k; p++) {
            for (let prev = 0; prev <= i - 2; prev++) {
                if (dp[prev][p - 1] !== INF && cost[prev][i - 1] !== INF)
                    dp[i][p] = Math.min(dp[i][p], dp[prev][p - 1] + cost[prev][i - 1]);
            }
        }
    }
    return dp[n][k];
};
```

### Python3
```python
class Solution:
    def minimumChanges(self, s: str, k: int) -> int:
        n = len(s)
        INF = 10**9
        cost = [[INF] * n for _ in range(n)]
        for i in range(n):
            for j in range(i + 1, n):
                L = j - i + 1
                best = INF
                for d in range(1, L):
                    if L % d != 0:
                        continue
                    cur = 0
                    m = L // d
                    for r in range(d):
                        for t in range(m // 2):
                            if s[i + r + t * d] != s[i + r + (m - 1 - t) * d]:
                                cur += 1
                    best = min(best, cur)
                cost[i][j] = best
        dp = [[INF] * (k + 1) for _ in range(n + 1)]
        dp[0][0] = 0
        for i in range(1, n + 1):
            for p in range(1, k + 1):
                for prev in range(i - 1):
                    if prev <= i - 2 and dp[prev][p - 1] != INF and cost[prev][i - 1] != INF:
                        dp[i][p] = min(dp[i][p], dp[prev][p - 1] + cost[prev][i - 1])
        return dp[n][k]
```

### Go
```go
func minimumChanges(s string, k int) int {
    n := len(s)
    const INF = 1000000000
    cost := make([][]int, n)
    for i := range cost {
        cost[i] = make([]int, n)
        for j := range cost[i] {
            cost[i][j] = INF
        }
    }
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            L := j - i + 1
            best := INF
            for d := 1; d < L; d++ {
                if L%d != 0 {
                    continue
                }
                cur := 0
                m := L / d
                for r := 0; r < d; r++ {
                    for t := 0; t < m/2; t++ {
                        if s[i+r+t*d] != s[i+r+(m-1-t)*d] {
                            cur++
                        }
                    }
                }
                if cur < best {
                    best = cur
                }
            }
            cost[i][j] = best
        }
    }
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, k+1)
        for p := range dp[i] {
            dp[i][p] = INF
        }
    }
    dp[0][0] = 0
    for i := 1; i <= n; i++ {
        for p := 1; p <= k; p++ {
            for prev := 0; prev <= i-2; prev++ {
                if dp[prev][p-1] != INF && cost[prev][i-1] != INF {
                    if dp[prev][p-1]+cost[prev][i-1] < dp[i][p] {
                        dp[i][p] = dp[prev][p-1] + cost[prev][i-1]
                    }
                }
            }
        }
    }
    return dp[n][k]
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages; only the syntax changes.

I begin by allocating an `n × n` cost matrix and fill every cell with a large sentinel value that means “this substring cannot be made into a semi-palindrome”.

Next I iterate over every left index `i` and every right index `j > i`. The length of the current substring is `L = j - i + 1`. For every candidate divisor `d` from 1 to `L-1` that actually divides `L`, I process the `d` independent residue groups. Inside each group the characters form a sequence of length `m = L / d`. I only need to look at the first half of that sequence: whenever the character at position `t` differs from the character at position `m-1-t`, one change is required. Summing those differences over all groups gives the cost of the chosen `d`. I keep the minimum cost among all valid `d` and write it into `cost[i][j]`.

After the cost table is complete I allocate a DP table of size `(n+1) × (k+1)`. The base case is `dp[0][0] = 0`. For every prefix length `i` from 1 to `n` and every piece count `p` from 1 to `k`, I try every earlier cut point `prev` that leaves a final piece of length at least 2. If both the prefix up to `prev` with `p-1` pieces and the final piece itself are feasible, I update `dp[i][p]` with the sum of their costs.

When the loops finish, `dp[n][k]` contains the answer. Edge cases such as length-2 substrings (only divisor `d = 1`) and the impossibility of length-1 semi-palindromes are handled automatically by the loops and the sentinel values.

## Examples

**Example 1**  
Input: `s = "abcac"`, `k = 2`  
Output: `1`  
Trace: One optimal split is `"ab"` + `"cac"`. `"cac"` is already a semi-palindrome. Changing the single letter of `"ab"` into `"aa"` costs 1, giving a total of 1.

**Example 2**  
Input: `s = "abcdef"`, `k = 2`  
Output: `2`  
Trace: Split into `"abc"` and `"def"`. Each three-letter string needs one change to become a semi-palindrome, for a total of 2.

**Example 3**  
Input: `s = "aabbaa"`, `k = 3`  
Output: `0`  
Trace: Split into `"aa"`, `"bb"`, `"aa"`. All three pieces are already semi-palindromes, so no changes are required.

## How to Use / Run Locally

**C++**  
1. Save the solution in a file named `main.cpp`.  
2. Compile: `g++ -std=c++17 main.cpp -o main`  
3. Run: `./main` (or add a small driver that reads input and prints the result).

**Java**  
1. Save the class in `Solution.java`.  
2. Compile: `javac Solution.java`  
3. Run with a main method that creates an instance and calls `minimumChanges`.

**JavaScript**  
1. Save the function in `solution.js`.  
2. Run with Node: `node solution.js` (add a few console.log calls for testing).

**TypeScript**  
1. Save the function in `solution.ts`.  
2. Compile: `tsc solution.ts`  
3. Run the generated JavaScript with Node.

**Python3**  
1. Save the class in `solution.py`.  
2. Run: `python3 solution.py` (add a few print statements to test the method).

**Go**  
1. Save the function in `main.go` inside a package main.  
2. Run: `go run main.go`.

In every language you can hard-code the sample inputs or read them from standard input according to your preference.

## Notes & Optimizations

- Length-1 strings can never be semi-palindromes, which is why the DP only considers pieces of length at least 2. The constraint `k ≤ n/2` guarantees a feasible answer always exists.
- The cost-table construction is O(n³) in the worst case, but because the number of divisors of any length is tiny the practical running time is much faster.
- An alternative approach that tries to reuse palindrome-cost computations across different divisors is possible, yet the extra bookkeeping is not worth the modest gain for n ≤ 200.
- Memory can be reduced by computing costs on the fly inside the DP, but the clearer two-stage version is easier to verify and still comfortably fits within typical limits.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
