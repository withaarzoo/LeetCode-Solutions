# Climbing Stairs II - LeetCode 3693 Solution

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

You are climbing a staircase that has n + 1 steps, numbered from 0 to n. You are given a 1-indexed integer array costs of length n, where costs[i] is the cost of step i.

From any step i you can jump only to step i + 1, i + 2, or i + 3. The cost of jumping from step i to step j is defined as costs[j] + (j - i) squared.

You start at step 0 with cost 0. Your task is to return the minimum total cost required to reach step n.

This is a classic minimum cost climbing stairs problem that needs an efficient dynamic programming solution because n can be as large as 100000.

## Constraints

- 1 <= n == costs.length <= 10^5
- 1 <= costs[i] <= 10^4

## Intuition

The first thing that stands out is the limited jump length. You can only move one, two, or three steps at a time. That immediately suggests that the cheapest way to reach any step only depends on the three steps before it.

The extra cost of (j - i) squared also means longer jumps become expensive quickly. So a pure greedy choice of always taking the cheapest next step would fail. I needed a way to remember the best cost to every previous step and reuse those answers. That is the classic signal for dynamic programming on a linear path.

## Approach

I created an array dp where dp[j] stores the minimum cost needed to reach step j.

I set dp[0] = 0 because the climb starts there for free.

Then for every step j from 1 to n I looked at the three possible previous positions i = j-1, j-2 and j-3 (only if they exist). For each of those i values I calculated:

dp[i] + costs[j-1] + (j - i) * (j - i)

I kept the smallest of those three candidates and stored it in dp[j].

After the loop finished, dp[n] held the answer.

This bottom-up order guarantees that every needed previous value is already computed, and the whole process runs in linear time.

## Data Structures Used

- A one-dimensional array (or vector / list) of size n + 1.  
  This stores the minimum cost to reach every step. I chose a simple array because the states form a straight line and each state only needs the three previous values. No extra graph or priority queue is required.

## Operations & Behavior Summary

1. Create a dp array of size n + 1 and fill it with a very large number.
2. Set dp[0] = 0.
3. For j = 1 to n:
   - Try the three possible previous steps (j-1, j-2, j-3).
   - For each valid previous step i, compute the jump cost and update dp[j] with the minimum.
4. Return the value stored in dp[n].

The algorithm never looks more than three steps backward, so the work per step stays constant.

## Complexity

| Complexity Type | Value | Explanation |
|-----------------|-------|-------------|
| Time Complexity | O(n) | We process each of the n steps once and do a constant amount of work (three candidates) for every step. |
| Space Complexity | O(n) | We keep an array of size n + 1 to store the minimum cost for every step. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int climbStairs(int n, vector<int>& costs) {
        vector<long long> dp(n + 1, LLONG_MAX);
        dp[0] = 0;
        for (int j = 1; j <= n; ++j) {
            for (int k = 1; k <= 3; ++k) {
                int i = j - k;
                if (i >= 0) {
                    long long jump = (long long)(j - i) * (j - i);
                    dp[j] = min(dp[j], dp[i] + costs[j - 1] + jump);
                }
            }
        }
        return (int)dp[n];
    }
};
```

### Java
```java
class Solution {
    public int climbStairs(int n, int[] costs) {
        long[] dp = new long[n + 1];
        Arrays.fill(dp, Long.MAX_VALUE);
        dp[0] = 0;
        for (int j = 1; j <= n; j++) {
            for (int k = 1; k <= 3; k++) {
                int i = j - k;
                if (i >= 0) {
                    long jump = (long)(j - i) * (j - i);
                    dp[j] = Math.min(dp[j], dp[i] + costs[j - 1] + jump);
                }
            }
        }
        return (int)dp[n];
    }
}
```

### JavaScript
```javascript
/**
 * @param {number} n
 * @param {number[]} costs
 * @return {number}
 */
var climbStairs = function(n, costs) {
    const dp = new Array(n + 1).fill(Number.MAX_SAFE_INTEGER);
    dp[0] = 0;
    for (let j = 1; j <= n; j++) {
        for (let k = 1; k <= 3; k++) {
            const i = j - k;
            if (i >= 0) {
                const jump = (j - i) * (j - i);
                dp[j] = Math.min(dp[j], dp[i] + costs[j - 1] + jump);
            }
        }
    }
    return dp[n];
};
```

### TypeScript
```typescript
function climbStairs(n: number, costs: number[]): number {
    const dp: number[] = new Array(n + 1).fill(Number.MAX_SAFE_INTEGER);
    dp[0] = 0;
    for (let j = 1; j <= n; j++) {
        for (let k = 1; k <= 3; k++) {
            const i = j - k;
            if (i >= 0) {
                const jump = (j - i) * (j - i);
                dp[j] = Math.min(dp[j], dp[i] + costs[j - 1] + jump);
            }
        }
    }
    return dp[n];
};
```

### Python3
```python
class Solution:
    def climbStairs(self, n: int, costs: List[int]) -> int:
        dp = [float('inf')] * (n + 1)
        dp[0] = 0
        for j in range(1, n + 1):
            for k in range(1, 4):
                i = j - k
                if i >= 0:
                    jump = (j - i) * (j - i)
                    dp[j] = min(dp[j], dp[i] + costs[j - 1] + jump)
        return dp[n]
```

### Go
```go
func climbStairs(n int, costs []int) int {
    dp := make([]int64, n+1)
    for i := range dp {
        dp[i] = 1e18
    }
    dp[0] = 0
    for j := 1; j <= n; j++ {
        for k := 1; k <= 3; k++ {
            i := j - k
            if i >= 0 {
                jump := int64(j-i) * int64(j-i)
                if dp[i]+int64(costs[j-1])+jump < dp[j] {
                    dp[j] = dp[i] + int64(costs[j-1]) + jump
                }
            }
        }
    }
    return int(dp[n])
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

Across all six languages the logic is identical. Only the syntax for arrays, loops and integer types changes.

I first allocate an array of size n + 1 and initialize every cell to a very large value so that the first real cost I write will always be smaller. I then force dp[0] to zero.

The outer loop walks from step 1 all the way to step n. Inside that loop a small inner loop tries the three legal jump lengths. For each candidate previous index i I check that i is non-negative, compute the squared distance, add the cost of the landing step, and keep the minimum.

Because the jumps are short, the intermediate sums can grow beyond 32-bit integers when n is large. That is why the C++, Java and Go versions use 64-bit integers for the dp array and only cast the final answer back to a normal int. JavaScript and TypeScript rely on the fact that Number can safely hold integers up to 2^53, which is enough for this problem. Python handles big integers automatically.

Edge cases are handled naturally: when j is 1 or 2 the invalid previous indices are simply skipped by the i >= 0 check. When n equals 1 the algorithm correctly evaluates the single possible jump from 0 to 1.

## Examples

**Example 1**  
Input: n = 4, costs = [1, 2, 3, 4]  
Output: 13  

Trace:  
- Reach 1 from 0: cost = 1 + 1^2 = 2  
- Reach 2 from 1: cost = 2 + 1^2 = 3, total so far 5  
- Reach 4 from 2: cost = 4 + 2^2 = 8, total 13  
Other paths cost more, so 13 is optimal.

**Example 2**  
Input: n = 4, costs = [5, 1, 6, 2]  
Output: 11  

Trace:  
- Jump directly 0 → 2: cost = 1 + 2^2 = 5  
- Jump 2 → 4: cost = 2 + 2^2 = 6, total 11  
This beats any path that lands on the expensive early steps.

**Example 3**  
Input: n = 3, costs = [9, 8, 3]  
Output: 12  

Trace:  
- Single jump 0 → 3: cost = 3 + 3^2 = 12  
Any path that stops at step 1 or 2 ends up more expensive.

## How to Use / Run Locally

**C++**  
1. Save the code in a file named main.cpp.  
2. Compile with: g++ -std=c++17 main.cpp -o main  
3. Run: ./main  

**Java**  
1. Save the code in a file named Solution.java.  
2. Compile with: javac Solution.java  
3. Run with a small driver that creates a Solution object and calls climbStairs.

**JavaScript**  
1. Save the code in a file named climbStairs.js.  
2. Run with Node: node climbStairs.js (add a small test call at the bottom).

**TypeScript**  
1. Save the code in a file named climbStairs.ts.  
2. Compile with: tsc climbStairs.ts  
3. Run the generated JavaScript with Node.

**Python3**  
1. Save the code in a file named solution.py.  
2. Run: python3 solution.py (add a short test at the bottom).

**Go**  
1. Save the code in a file named main.go.  
2. Run: go run main.go  

In every language you will need a tiny main / driver function that supplies sample inputs and prints the returned value. The core climbStairs function itself stays exactly as shown in the empty blocks above.

## Notes & Optimizations

Because each state only depends on the three previous states, the O(n) space can be reduced to O(1) by keeping only three rolling variables. The version shown here keeps the full array for clarity; both versions pass the time limits.

When costs[i] are large and n is close to 1e5, intermediate values can exceed 2^31. Using 64-bit integers (or Python’s unlimited integers) prevents overflow.

An alternative top-down memoized recursion works but needs extra care with recursion depth on large n. The bottom-up loop is safer and faster in practice.

The same pattern appears in many “jump game with costs” and “minimum cost climbing stairs” problems, so mastering this solution helps with a whole family of dynamic programming questions.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
