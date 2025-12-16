# 3573. Best Time to Buy and Sell Stock V

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given an array `prices`, where `prices[i]` represents the stock price on day `i`.
I am also given an integer `k`, which is the **maximum number of transactions** I can make.

Each transaction can be:

* **Normal transaction**: buy first, then sell later
* **Short selling transaction**: sell first, then buy back later

Rules:

* I must finish one transaction before starting another.
* I cannot buy and sell on the same day.
* A transaction is counted **only when it is completed**.

My task is to return the **maximum total profit** I can make using **at most `k` transactions**.

---

## Constraints

* `2 ≤ prices.length ≤ 1000`
* `1 ≤ prices[i] ≤ 10^9`
* `1 ≤ k ≤ prices.length / 2`

---

## Intuition

When I first analyzed the problem, I realized this is **not a normal stock problem**.

Here, I can make profit in **both rising and falling markets** because **short selling is allowed**.

So at any day, I can be in one of three states:

1. I am free (no position)
2. I am holding a **long position** (I bought earlier)
3. I am holding a **short position** (I sold earlier)

The most important rule I understood was:

> A transaction is counted **only when I close a position**, not when I open it.

This made it clear that I need a **state-based Dynamic Programming solution**.

---

## Approach

I solved this problem using **Dynamic Programming with state tracking**.

I defined my DP state as:

```
dp[i][state][t]
```

Where:

* `i` = current day
* `state`:

  * `0` → no position
  * `1` → holding long
  * `2` → holding short
* `t` = number of transactions already completed

### Key rules I enforced

* I can only close a position if `t < k`
* I **must not end** with an open position
* Ending with a long or short position is treated as invalid

The final answer is stored in:

```
dp[0][0][0]
```

---

## Data Structures Used

* 3D Dynamic Programming table
* Arrays / slices (language specific)
* Recursion + memoization (C++ / Java)
* Bottom-up DP (Python, JS, Go)

---

## Operations & Behavior Summary

* Buy → moves from state `0` to `1`
* Sell → moves from state `1` to `0` and increases transaction count
* Short sell → moves from state `0` to `2`
* Buy back → moves from state `2` to `0` and increases transaction count
* Skip → stay in the same state

All invalid states are blocked using a very large negative value.

---

## Complexity

* **Time Complexity:** `O(n × k)`

  * `n` = number of days
  * `k` = maximum transactions

* **Space Complexity:** `O(n × k × 3)`

  * 3 states per day per transaction count

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long dp[1005][505][3];
    int n, K;
    vector<int> p;

    long long solve(int i, int t, int s) {
        if (i == n) return (s == 0 ? 0 : LLONG_MIN / 4);
        long long &res = dp[i][t][s];
        if (res != -1) return res;

        res = solve(i + 1, t, s);

        if (s == 0) {
            res = max(res, solve(i + 1, t, 1) - p[i]);
            res = max(res, solve(i + 1, t, 2) + p[i]);
        } else if (s == 1 && t < K) {
            res = max(res, solve(i + 1, t + 1, 0) + p[i]);
        } else if (s == 2 && t < K) {
            res = max(res, solve(i + 1, t + 1, 0) - p[i]);
        }
        return res;
    }

    long long maximumProfit(vector<int>& prices, int k) {
        p = prices;
        n = prices.size();
        K = k;
        memset(dp, -1, sizeof(dp));
        return solve(0, 0, 0);
    }
};
```

---

### Java

```java
class Solution {
    int n, K;
    int[] p;
    Long[][][] dp;
    final long NEG = (long)-1e18;

    long dfs(int i, int t, int s) {
        if (i == n) return s == 0 ? 0 : NEG;
        if (dp[i][t][s] != null) return dp[i][t][s];

        long res = dfs(i + 1, t, s);

        if (s == 0) {
            res = Math.max(res, dfs(i + 1, t, 1) - p[i]);
            res = Math.max(res, dfs(i + 1, t, 2) + p[i]);
        } else if (s == 1 && t < K) {
            res = Math.max(res, dfs(i + 1, t + 1, 0) + p[i]);
        } else if (s == 2 && t < K) {
            res = Math.max(res, dfs(i + 1, t + 1, 0) - p[i]);
        }

        return dp[i][t][s] = res;
    }

    public long maximumProfit(int[] prices, int k) {
        p = prices;
        n = prices.length;
        K = k;
        dp = new Long[n + 1][k + 1][3];
        return dfs(0, 0, 0);
    }
}
```

---

### JavaScript

```javascript
var maximumProfit = function(prices, k) {
    const n = prices.length;
    const NEG = -1e18;
    const dp = Array.from({ length: n + 1 }, () =>
        Array.from({ length: 3 }, () => Array(k + 1).fill(NEG))
    );

    for (let t = 0; t <= k; t++) dp[n][0][t] = 0;

    for (let i = n - 1; i >= 0; i--) {
        for (let t = 0; t <= k; t++) {
            dp[i][0][t] = Math.max(
                dp[i + 1][0][t],
                -prices[i] + dp[i + 1][1][t],
                prices[i] + dp[i + 1][2][t]
            );

            if (t < k) {
                dp[i][1][t] = Math.max(
                    dp[i + 1][1][t],
                    prices[i] + dp[i + 1][0][t + 1]
                );
                dp[i][2][t] = Math.max(
                    dp[i + 1][2][t],
                    -prices[i] + dp[i + 1][0][t + 1]
                );
            }
        }
    }
    return dp[0][0][0];
};
```

---

### Python3

```python
class Solution:
    def maximumProfit(self, prices, k):
        n = len(prices)
        NEG = -10**18
        dp = [[[NEG]*(k+1) for _ in range(3)] for _ in range(n+1)]

        for t in range(k+1):
            dp[n][0][t] = 0

        for i in range(n-1, -1, -1):
            for t in range(k+1):
                dp[i][0][t] = max(
                    dp[i+1][0][t],
                    -prices[i] + dp[i+1][1][t],
                    prices[i] + dp[i+1][2][t]
                )
                if t < k:
                    dp[i][1][t] = max(dp[i+1][1][t], prices[i] + dp[i+1][0][t+1])
                    dp[i][2][t] = max(dp[i+1][2][t], -prices[i] + dp[i+1][0][t+1])

        return dp[0][0][0]
```

---

### Go

```go
func maximumProfit(prices []int, k int) int64 {
    n := len(prices)
    const NEG int64 = -1e18

    dp := make([][][]int64, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([][]int64, 3)
        for s := 0; s < 3; s++ {
            dp[i][s] = make([]int64, k+1)
            for t := 0; t <= k; t++ {
                dp[i][s][t] = NEG
            }
        }
    }

    for t := 0; t <= k; t++ {
        dp[n][0][t] = 0
    }

    for i := n - 1; i >= 0; i-- {
        for t := 0; t <= k; t++ {
            dp[i][0][t] = max3(
                dp[i+1][0][t],
                -int64(prices[i]) + dp[i+1][1][t],
                int64(prices[i]) + dp[i+1][2][t],
            )
            if t < k {
                dp[i][1][t] = max(dp[i+1][1][t], int64(prices[i]) + dp[i+1][0][t+1])
                dp[i][2][t] = max(dp[i+1][2][t], -int64(prices[i]) + dp[i+1][0][t+1])
            }
        }
    }
    return dp[0][0][0]
}

func max(a, b int64) int64 {
    if a > b { return a }
    return b
}

func max3(a, b, c int64) int64 {
    return max(a, max(b, c))
}
```

---

## Step-by-step Detailed Explanation

1. I iterate day by day from the end.
2. For each day, I try all three states.
3. I test all valid actions for that state.
4. I block invalid states using a negative infinity value.
5. I store the best profit for each `(day, state, transaction)`.

---

## Examples

**Input**

```
prices = [12,16,19,8,1,19,13,9]
k = 3
```

**Output**

```
36
```

---

## How to use / Run locally

1. Clone the repository
2. Choose your language folder
3. Compile / run using standard compiler or interpreter
4. Submit the same logic to LeetCode

---

## Notes & Optimizations

* This is a **state-machine DP**
* Greedy approaches fail here
* Space can be optimized further to `O(k)`
* Core idea applies to many advanced stock problems

---

## Author

* **Md Aarzoo Islam**

  * [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
