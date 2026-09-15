---

# 2472. Maximum Number of Non-overlapping Palindrome Substrings

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

You are given a string `s` and a positive integer `k`. The task is to select the largest possible set of non-overlapping substrings from `s` such that every selected substring has length at least `k` and is a palindrome.

A substring is any contiguous sequence of characters. Non-overlapping means the selected pieces do not share any character positions.

Return the maximum number of such substrings you can pick.

This is a classic dynamic programming problem on strings that appears in coding interviews and LeetCode contests. It tests your ability to combine palindrome checking with optimal substructure for non-overlapping selections.

## Constraints

- 1 <= k <= s.length <= 2000
- `s` consists of lowercase English letters only

Because the length of the string can reach 2000, any solution slower than O(n²) will time out.

## Intuition

The first thing that stood out is that the order of decisions matters. Once I decide to take a valid palindrome ending at a certain index, the next one must start after it. That means the best answer for any prefix of the string only depends on earlier prefixes.

This observation points straight to dynamic programming. If I can quickly answer “is the substring from j to i a palindrome?”, then I can try every possible last piece and keep the maximum count.

Pre-computing all palindromic substrings in O(n²) time is acceptable under the given constraints and makes the later DP clean and fast.

## Approach

I build a 2-D boolean table that tells me in constant time whether any substring is a palindrome. I fill the table by length: every single character is a palindrome, every pair of equal characters is a palindrome, and for longer lengths I check the two ends plus the already-computed inner part.

After the table is ready I create a 1-D DP array where `dp[i]` stores the maximum number of valid non-overlapping pieces I can form using only the first `i` characters of the string.

For every position `i` I have two choices:

1. Do not end a new piece at index `i-1`. Then the answer is simply `dp[i-1]`.
2. End a new piece at index `i-1`. I look at every possible start index `j` such that the length is at least `k` and the substring is a palindrome. If both conditions hold, I can take `dp[j] + 1`.

I keep the better of the two choices. At the end `dp[n]` is the answer for the whole string.

## Data Structures Used

- A 2-D boolean matrix `isPal` of size n × n.  
  It stores whether every possible substring is a palindrome. I need constant-time lookups later, so the quadratic space is justified.

- A 1-D integer array `dp` of size n + 1.  
  `dp[i]` holds the best answer for the prefix of length i. This is the classic optimal-substructure table for the non-overlapping selection problem.

No other data structures are required. The solution stays within the memory limits for n = 2000.

## Operations & Behavior Summary

1. Allocate and initialize the isPal matrix.  
2. Mark every length-1 substring as a true.  
3. Mark every length-2 substring whose two characters are equal as true.  
4. For lengths 3 to n, fill the remaining cells using the recurrence: ends match and the inner substring is already known to be a palindrome.  
5. Allocate the dp array and set dp[0] = 0.  
6. For each right endpoint i from 1 to n:  
   - Start with the skip option (dp[i] = dp[i-1]).  
   - Try every feasible left endpoint j and update with dp[j] + 1 when a valid palindrome is found.  
7. Return dp[n].

## Complexity

| Metric            | Value   | Explanation |
|-------------------|---------|-------------|
| Time Complexity   | O(n²)   | Building the isPal table takes O(n²). The DP then examines O(n) possible left endpoints for each of the n right endpoints, also O(n²). |
| Space Complexity  | O(n²)   | The isPal matrix dominates. The dp array uses only O(n) extra space. |

Both bounds are optimal for the constraints; anything slower fails on the largest test cases.

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int maxPalindromes(string s, int k) {
        int n = s.size();
        // isPal[i][j] == true  iff  s[i..j] is a palindrome
        vector<vector<bool>> isPal(n, vector<bool>(n, false));
        
        // every single character is a palindrome
        for (int i = 0; i < n; ++i)
            isPal[i][i] = true;
        
        // every pair of identical characters is a palindrome
        for (int i = 0; i + 1 < n; ++i)
            if (s[i] == s[i + 1])
                isPal[i][i + 1] = true;
        
        // longer lengths: ends equal and inner part already known to be palindrome
        for (int len = 3; len <= n; ++len)
            for (int i = 0; i + len - 1 < n; ++i) {
                int j = i + len - 1;
                if (s[i] == s[j] && isPal[i + 1][j - 1])
                    isPal[i][j] = true;
            }
        
        // dp[i] = maximum number of valid pieces inside s[0..i-1]
        vector<int> dp(n + 1, 0);
        for (int i = 1; i <= n; ++i) {
            dp[i] = dp[i - 1];                 // skip the last character
            // try every possible start of a piece that ends at i-1
            for (int j = 0; j <= i - k; ++j)
                if (isPal[j][i - 1])
                    dp[i] = max(dp[i], dp[j] + 1);
        }
        return dp[n];
    }
};
```

### Java
```java
class Solution {
    public int maxPalindromes(String s, int k) {
        int n = s.length();
        // isPal[i][j] == true  iff  s[i..j] is a palindrome
        boolean[][] isPal = new boolean[n][n];
        
        // every single character is a palindrome
        for (int i = 0; i < n; ++i)
            isPal[i][i] = true;
        
        // every pair of identical characters is a palindrome
        for (int i = 0; i + 1 < n; ++i)
            if (s.charAt(i) == s.charAt(i + 1))
                isPal[i][i + 1] = true;
        
        // longer lengths: ends equal and inner part already known to be palindrome
        for (int len = 3; len <= n; ++len)
            for (int i = 0; i + len - 1 < n; ++i) {
                int j = i + len - 1;
                if (s.charAt(i) == s.charAt(j) && isPal[i + 1][j - 1])
                    isPal[i][j] = true;
            }
        
        // dp[i] = maximum number of valid pieces inside s[0..i-1]
        int[] dp = new int[n + 1];
        for (int i = 1; i <= n; ++i) {
            dp[i] = dp[i - 1];                 // skip the last character
            // try every possible start of a piece that ends at i-1
            for (int j = 0; j <= i - k; ++j)
                if (isPal[j][i - 1])
                    dp[i] = Math.max(dp[i], dp[j] + 1);
        }
        return dp[n];
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
var maxPalindromes = function(s, k) {
    const n = s.length;
    // isPal[i][j] == true  iff  s[i..j] is a palindrome
    const isPal = Array.from({length: n}, () => Array(n).fill(false));
    
    // every single character is a palindrome
    for (let i = 0; i < n; ++i)
        isPal[i][i] = true;
    
    // every pair of identical characters is a palindrome
    for (let i = 0; i + 1 < n; ++i)
        if (s[i] === s[i + 1])
            isPal[i][i + 1] = true;
    
    // longer lengths: ends equal and inner part already known to be palindrome
    for (let len = 3; len <= n; ++len)
        for (let i = 0; i + len - 1 < n; ++i) {
            const j = i + len - 1;
            if (s[i] === s[j] && isPal[i + 1][j - 1])
                isPal[i][j] = true;
        }
    
    // dp[i] = maximum number of valid pieces inside s[0..i-1]
    const dp = new Array(n + 1).fill(0);
    for (let i = 1; i <= n; ++i) {
        dp[i] = dp[i - 1];                 // skip the last character
        // try every possible start of a piece that ends at i-1
        for (let j = 0; j <= i - k; ++j)
            if (isPal[j][i - 1])
                dp[i] = Math.max(dp[i], dp[j] + 1);
    }
    return dp[n];
};
```

### TypeScript
```typescript
function maxPalindromes(s: string, k: number): number {
    const n = s.length;
    // isPal[i][j] == true  iff  s[i..j] is a palindrome
    const isPal: boolean[][] = Array.from({length: n}, () => Array(n).fill(false));
    
    // every single character is a palindrome
    for (let i = 0; i < n; ++i)
        isPal[i][i] = true;
    
    // every pair of identical characters is a palindrome
    for (let i = 0; i + 1 < n; ++i)
        if (s[i] === s[i + 1])
            isPal[i][i + 1] = true;
    
    // longer lengths: ends equal and inner part already known to be palindrome
    for (let len = 3; len <= n; ++len)
        for (let i = 0; i + len - 1 < n; ++i) {
            const j = i + len - 1;
            if (s[i] === s[j] && isPal[i + 1][j - 1])
                isPal[i][j] = true;
        }
    
    // dp[i] = maximum number of valid pieces inside s[0..i-1]
    const dp: number[] = new Array(n + 1).fill(0);
    for (let i = 1; i <= n; ++i) {
        dp[i] = dp[i - 1];                 // skip the last character
        // try every possible start of a piece that ends at i-1
        for (let j = 0; j <= i - k; ++j)
            if (isPal[j][i - 1])
                dp[i] = Math.max(dp[i], dp[j] + 1);
    }
    return dp[n];
}
```

### Python3
```python
class Solution:
    def maxPalindromes(self, s: str, k: int) -> int:
        n = len(s)
        # isPal[i][j] == True  iff  s[i..j] is a palindrome
        isPal = [[False] * n for _ in range(n)]
        
        # every single character is a palindrome
        for i in range(n):
            isPal[i][i] = True
        
        # every pair of identical characters is a palindrome
        for i in range(n - 1):
            if s[i] == s[i + 1]:
                isPal[i][i + 1] = True
        
        # longer lengths: ends equal and inner part already known to be palindrome
        for length in range(3, n + 1):
            for i in range(n - length + 1):
                j = i + length - 1
                if s[i] == s[j] and isPal[i + 1][j - 1]:
                    isPal[i][j] = True
        
        # dp[i] = maximum number of valid pieces inside s[0..i-1]
        dp = [0] * (n + 1)
        for i in range(1, n + 1):
            dp[i] = dp[i - 1]                 # skip the last character
            # try every possible start of a piece that ends at i-1
            for j in range(i - k + 1):
                if isPal[j][i - 1]:
                    dp[i] = max(dp[i], dp[j] + 1)
        return dp[n]
```

### Go
```go
func maxPalindromes(s string, k int) int {
    n := len(s)
    // isPal[i][j] == true  iff  s[i..j] is a palindrome
    isPal := make([][]bool, n)
    for i := range isPal {
        isPal[i] = make([]bool, n)
    }
    
    // every single character is a palindrome
    for i := 0; i < n; i++ {
        isPal[i][i] = true
    }
    
    // every pair of identical characters is a palindrome
    for i := 0; i+1 < n; i++ {
        if s[i] == s[i+1] {
            isPal[i][i+1] = true
        }
    }
    
    // longer lengths: ends equal and inner part already known to be palindrome
    for length := 3; length <= n; length++ {
        for i := 0; i+length-1 < n; i++ {
            j := i + length - 1
            if s[i] == s[j] && isPal[i+1][j-1] {
                isPal[i][j] = true
            }
        }
    }
    
    // dp[i] = maximum number of valid pieces inside s[0..i-1]
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = dp[i-1] // skip the last character
        // try every possible start of a piece that ends at i-1
        for j := 0; j <= i-k; j++ {
            if isPal[j][i-1] {
                if dp[j]+1 > dp[i] {
                    dp[i] = dp[j] + 1
                }
            }
        }
    }
    return dp[n]
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages; only syntax differs.

First I create the boolean matrix. In every language this is a straightforward nested array of size n by n, initialized to false.

I set the main diagonal to true because a single character is always a palindrome. Then I walk consecutive pairs and set those cells to true when the two characters match.

The heart of the pre-computation is the length loop. I start from length 3 and grow upward. For a candidate interval [left, right] I only need to test whether the characters at the two ends are equal and whether the already-filled inner cell is true. Because shorter lengths have already been processed, the inner value is always available. This classic bottom-up filling guarantees correctness and runs in quadratic time.

After the matrix is ready I allocate the dp array of size n+1 and leave dp[0] at zero. For every right endpoint i I first copy the previous value; that corresponds to the decision “I do not finish a new piece here.” Then I iterate over every possible start index that would give length at least k. When the matrix reports a true value I consider taking one extra piece after the optimal solution of the earlier prefix. I keep the maximum.

When the loop finishes, the last cell of dp contains the answer. The same sequence of decisions works in C++, Java, JavaScript, TypeScript, Python and Go; only the way arrays and loops are written changes.

Edge cases are handled automatically: if k is larger than n the inner loop never runs and the answer stays zero; if the whole string is one big palindrome the DP will correctly report 1.

## Examples

**Example 1**

Input: s = "abaccdbbd", k = 3  
Output: 2  

Trace:  
- The substring “aba” (indices 0-2) is a palindrome of length 3.  
- The substring “dbbd” (indices 5-8) is a palindrome of length 4.  
- They do not overlap, so the maximum count is 2.  
- The DP reaches the value 2 when the second piece is considered.

**Example 2**

Input: s = "adbcda", k = 2  
Output: 0  

Trace:  
- No substring of length 2 or more is a palindrome.  
- Therefore every possible update is skipped and the answer remains 0.

**Example 3**

Input: s = "aaa", k = 1  
Output: 3  

Trace:  
- Every single character is a valid piece.  
- The DP can pick all three of them because they never overlap.

## How to Use / Run Locally

**C++**  
Save the code in a file named `main.cpp`. Compile with  
`g++ -std=c++17 main.cpp -o main`  
Run with `./main` and supply the input through standard input or hard-code a test case inside main.

**Java**  
Save the code in `Solution.java`. Compile with `javac Solution.java` and run with `java Solution`. Add a main method that creates a Solution object and calls the method with sample strings.

**JavaScript**  
Save the code in `maxPalindromes.js`. Run with Node.js:  
`node maxPalindromes.js`  
You can call the function directly from the terminal or from another script.

**TypeScript**  
Save the code in `maxPalindromes.ts`. Compile with `tsc maxPalindromes.ts` then run the generated JavaScript file with Node.

**Python3**  
Save the code in `solution.py`. Run with  
`python3 solution.py`  
Add a few print statements at the bottom to test the examples.

**Go**  
Save the code in `main.go`. Run with  
`go run main.go`  
or build an executable with `go build` and execute the binary.

In every language you can paste the official sample inputs into a small driver function to verify the output before submitting to LeetCode.

## Notes & Optimizations

The quadratic space of the isPal matrix is the main tradeoff. For n = 2000 it uses a few megabytes, which is fine on LeetCode. If memory ever becomes a problem one could expand around centers on the fly, but that would make the DP slower by a logarithmic factor or force more complicated bookkeeping.

A pure greedy approach that always takes the shortest possible palindrome can fail on some inputs, so the DP is safer.

The same framework works for many other non-overlapping substring problems: replace the palindrome check with any other property that can be pre-computed in O(n²) and the rest of the code stays unchanged.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)