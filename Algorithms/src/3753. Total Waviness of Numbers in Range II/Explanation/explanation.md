# 3753. Total Waviness of Numbers in Range II

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

Given two integers `num1` and `num2`, we need to calculate the total waviness of every number inside the inclusive range `[num1, num2]`.

A digit is considered:

* A peak if it is strictly greater than both neighboring digits.
* A valley if it is strictly smaller than both neighboring digits.

The first and last digits can never be peaks or valleys because they do not have two neighbors.

The waviness of a number is the total count of all peaks and valleys present in that number.

The goal is to return the sum of waviness values for every number in the given range.

Since the constraints go up to `10^15`, checking every number one by one is far too slow. A more advanced approach such as Digit DP is required.

Relevant SEO keywords naturally associated with this problem include:

* LeetCode 3753 Solution
* Total Waviness of Numbers in Range II
* Digit DP
* Dynamic Programming
* Range Query Problems
* Competitive Programming
* DSA Interview Problems
* Peak and Valley Digits

## Constraints

| Constraint  | Value        |
| ----------- | ------------ |
| num1        | 1 ≤ num1     |
| num2        | num1 ≤ num2  |
| Upper Limit | num2 ≤ 10^15 |

## Intuition

My first thought was to directly calculate the waviness of every number in the range.

That works for very small inputs, but it immediately becomes impossible when the range can contain trillions of numbers.

I noticed that whether a digit becomes a peak or valley depends only on its immediate neighbors.

That observation naturally leads to Digit DP.

Instead of generating every number, I can build valid numbers digit by digit and count waviness contributions while constructing them.

This allows me to efficiently count the total waviness for all numbers up to a given limit without visiting each number individually.

## Approach

I use Digit DP to calculate the total waviness from `0` to `N`.

The solution works in the following steps:

1. Convert the upper bound into a string.
2. Build numbers digit by digit.
3. Keep track of:

   * Current position
   * Whether the current prefix is still restricted by the upper bound
   * Whether the number has started
   * Last digit
   * Second last digit
4. Whenever a new digit is added, check whether the previous digit becomes a peak or valley.
5. Accumulate the waviness contribution.
6. Memoize states to avoid repeated calculations.
7. Create a helper function `solve(N)` that returns total waviness from `0` to `N`.
8. Use range subtraction:

   `solve(num2) - solve(num1 - 1)`

This standard Digit DP technique converts a range problem into two prefix calculations.

## Data Structures Used

### Memoization Table

Stores already computed DP states.

Why?

Because many digit combinations repeat during recursion.

### String

The upper bound number is converted into a string so individual digits can be processed easily.

### Recursive DP State

The recursion stores:

* Current position
* Tight flag
* Started flag
* Last digit
* Second last digit

These values completely describe the current state.

## Operations & Behavior Summary

The algorithm performs the following operations:

1. Read the upper bound digit by digit.
2. Build all valid numbers through recursion.
3. Handle leading zeros separately.
4. Track the previous two digits.
5. Whenever three consecutive digits are available:

   * Check peak condition.
   * Check valley condition.
6. Add waviness contribution.
7. Continue building remaining digits.
8. Store computed states in DP.
9. Return total waviness for all valid numbers.
10. Subtract prefix results to obtain the final range answer.

## Complexity

| Metric           | Complexity                  |
| ---------------- | --------------------------- |
| Time Complexity  | O(L × 11 × 11 × 2 × 2 × 10) |
| Space Complexity | O(L × 11 × 11 × 2 × 2)      |

Where:

* `L` is the number of digits in the upper bound.
* Maximum `L` is only around 16 because `10^15` contains very few digits.

This makes the solution extremely efficient.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    struct Node {
        long long cnt;
        long long wav;
    };

    string s;
    Node dp[20][2][11][11];
    bool vis[20][2][11][11];

    Node dfs(int pos, int started, int last, int secondLast, bool tight) {
        if (pos == (int)s.size()) {
            return {1, 0};
        }

        if (!tight && vis[pos][started][last][secondLast]) {
            return dp[pos][started][last][secondLast];
        }

        int limit = tight ? s[pos] - '0' : 9;

        Node res = {0, 0};

        for (int d = 0; d <= limit; d++) {
            bool ntight = tight && (d == limit);

            if (!started && d == 0) {
                Node nxt = dfs(pos + 1, 0, 10, 10, ntight);

                res.cnt += nxt.cnt;
                res.wav += nxt.wav;
            } else {
                long long add = 0;

                if (started && secondLast != 10) {
                    if ((last > secondLast && last > d) ||
                        (last < secondLast && last < d)) {
                        add = 1;
                    }
                }

                int nSecondLast = started ? last : 10;
                int nLast = d;

                Node nxt = dfs(pos + 1, 1, nLast, nSecondLast, ntight);

                res.cnt += nxt.cnt;
                res.wav += nxt.wav + add * nxt.cnt;
            }
        }

        if (!tight) {
            vis[pos][started][last][secondLast] = true;
            dp[pos][started][last][secondLast] = res;
        }

        return res;
    }

    long long solve(long long n) {
        if (n < 0) return 0;

        s = to_string(n);
        memset(vis, 0, sizeof(vis));

        return dfs(0, 0, 10, 10, true).wav;
    }

    long long totalWaviness(long long num1, long long num2) {
        return solve(num2) - solve(num1 - 1);
    }
};
```

### Java

```java
class Solution {

    static class Node {
        long cnt;
        long wav;

        Node(long cnt, long wav) {
            this.cnt = cnt;
            this.wav = wav;
        }
    }

    String s;
    Node[][][][] dp = new Node[20][2][11][11];
    boolean[][][][] vis = new boolean[20][2][11][11];

    private Node dfs(int pos, int started, int last,
                     int secondLast, boolean tight) {

        if (pos == s.length()) {
            return new Node(1, 0);
        }

        if (!tight && vis[pos][started][last][secondLast]) {
            return dp[pos][started][last][secondLast];
        }

        int limit = tight ? s.charAt(pos) - '0' : 9;

        long totalCnt = 0;
        long totalWav = 0;

        for (int d = 0; d <= limit; d++) {
            boolean ntight = tight && (d == limit);

            if (started == 0 && d == 0) {
                Node nxt = dfs(pos + 1, 0, 10, 10, ntight);

                totalCnt += nxt.cnt;
                totalWav += nxt.wav;
            } else {
                long add = 0;

                if (started == 1 && secondLast != 10) {
                    if ((last > secondLast && last > d) ||
                        (last < secondLast && last < d)) {
                        add = 1;
                    }
                }

                int nSecondLast = (started == 1) ? last : 10;

                Node nxt = dfs(pos + 1, 1, d, nSecondLast, ntight);

                totalCnt += nxt.cnt;
                totalWav += nxt.wav + add * nxt.cnt;
            }
        }

        Node res = new Node(totalCnt, totalWav);

        if (!tight) {
            vis[pos][started][last][secondLast] = true;
            dp[pos][started][last][secondLast] = res;
        }

        return res;
    }

    private long solve(long n) {
        if (n < 0) return 0;

        s = Long.toString(n);

        vis = new boolean[20][2][11][11];
        dp = new Node[20][2][11][11];

        return dfs(0, 0, 10, 10, true).wav;
    }

    public long totalWaviness(long num1, long num2) {
        return solve(num2) - solve(num1 - 1);
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} num1
 * @param {number} num2
 * @return {number}
 */
var totalWaviness = function(num1, num2) {

    function solve(n) {
        if (n < 0) return 0;

        const s = String(n);
        const memo = new Map();

        function dfs(pos, started, last, secondLast, tight) {
            if (pos === s.length) {
                return [1n, 0n];
            }

            const key = `${pos}|${started}|${last}|${secondLast}`;

            if (!tight && memo.has(key)) {
                return memo.get(key);
            }

            const limit = tight ? Number(s[pos]) : 9;

            let cnt = 0n;
            let wav = 0n;

            for (let d = 0; d <= limit; d++) {
                const ntight = tight && d === limit;

                if (!started && d === 0) {
                    const [c, w] = dfs(pos + 1, false, 10, 10, ntight);

                    cnt += c;
                    wav += w;
                } else {
                    let add = 0n;

                    if (started && secondLast !== 10) {
                        if (
                            (last > secondLast && last > d) ||
                            (last < secondLast && last < d)
                        ) {
                            add = 1n;
                        }
                    }

                    const nSecondLast = started ? last : 10;

                    const [c, w] = dfs(
                        pos + 1,
                        true,
                        d,
                        nSecondLast,
                        ntight
                    );

                    cnt += c;
                    wav += w + add * c;
                }
            }

            const res = [cnt, wav];

            if (!tight) {
                memo.set(key, res);
            }

            return res;
        }

        return Number(dfs(0, false, 10, 10, true)[1]);
    }

    return solve(num2) - solve(num1 - 1);
};
```

### Python3

```python
class Solution:
    def totalWaviness(self, num1: int, num2: int) -> int:

        from functools import lru_cache

        def solve(n: int) -> int:
            if n < 0:
                return 0

            s = str(n)

            @lru_cache(None)
            def dfs(pos, started, last, second_last, tight):
                if pos == len(s):
                    return (1, 0)

                limit = int(s[pos]) if tight else 9

                total_cnt = 0
                total_wav = 0

                for d in range(limit + 1):
                    ntight = tight and d == limit

                    if not started and d == 0:
                        cnt, wav = dfs(
                            pos + 1,
                            False,
                            10,
                            10,
                            ntight
                        )

                        total_cnt += cnt
                        total_wav += wav

                    else:
                        add = 0

                        if started and second_last != 10:
                            if (
                                (last > second_last and last > d)
                                or
                                (last < second_last and last < d)
                            ):
                                add = 1

                        n_second_last = last if started else 10

                        cnt, wav = dfs(
                            pos + 1,
                            True,
                            d,
                            n_second_last,
                            ntight
                        )

                        total_cnt += cnt
                        total_wav += wav + add * cnt

                return (total_cnt, total_wav)

            return dfs(0, False, 10, 10, True)[1]

        return solve(num2) - solve(num1 - 1)
```

### Go

```go
func totalWaviness(num1 int64, num2 int64) int64 {

 type Node struct {
  cnt int64
  wav int64
 }

 var solve func(int64) int64

 solve = func(n int64) int64 {
  if n < 0 {
   return 0
  }

  s := []byte(fmt.Sprintf("%d", n))

  type State struct {
   pos        int
   started    int
   last       int
   secondLast int
  }

  memo := map[State]Node{}

  var dfs func(int, int, int, int, bool) Node

  dfs = func(pos, started, last, secondLast int, tight bool) Node {
   if pos == len(s) {
    return Node{1, 0}
   }

   st := State{pos, started, last, secondLast}

   if !tight {
    if val, ok := memo[st]; ok {
     return val
    }
   }

   limit := 9
   if tight {
    limit = int(s[pos] - '0')
   }

   res := Node{}

   for d := 0; d <= limit; d++ {
    ntight := tight && d == limit

    if started == 0 && d == 0 {
     nxt := dfs(pos+1, 0, 10, 10, ntight)

     res.cnt += nxt.cnt
     res.wav += nxt.wav
    } else {
     var add int64 = 0

     if started == 1 && secondLast != 10 {
      if (last > secondLast && last > d) ||
       (last < secondLast && last < d) {
       add = 1
      }
     }

     nSecondLast := 10
     if started == 1 {
      nSecondLast = last
     }

     nxt := dfs(pos+1, 1, d, nSecondLast, ntight)

     res.cnt += nxt.cnt
     res.wav += nxt.wav + add*nxt.cnt
    }
   }

   if !tight {
    memo[st] = res
   }

   return res
  }

  return dfs(0, 0, 10, 10, true).wav
 }

 return solve(num2) - solve(num1-1)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains identical across all five implementations.

Only syntax differs.

### Step 1: Create a Prefix Solver

The first helper function computes total waviness from `0` to `N`.

This simplifies the range calculation later.

### Step 2: Convert Number to Digits

The upper bound is converted into a string.

This makes it easy to process one digit at a time.

### Step 3: Define DP State

Each state stores:

* Current digit position
* Tight restriction
* Started status
* Last digit
* Second last digit

These values uniquely determine future possibilities.

### Step 4: Handle Leading Zeros

Before the number starts, zeros are ignored.

This prevents leading zeros from affecting waviness calculations.

### Step 5: Try Every Possible Digit

For every position:

* Choose a digit from `0` to `9`
* Respect the upper bound if the state is tight

### Step 6: Detect Peaks and Valleys

Once three digits are available:

`a b c`

The middle digit `b` becomes:

Peak if:

`b > a` and `b > c`

Valley if:

`b < a` and `b < c`

Whenever either condition is true, waviness increases by one.

### Step 7: Accumulate Results

Each recursive call returns:

* Number count
* Total waviness

These values are merged into the current state.

### Step 8: Memoization

Non-tight states are cached.

This avoids recalculating identical states repeatedly.

### Step 9: Range Query

Finally:

`Answer = solve(num2) - solve(num1 - 1)`

This gives the total waviness inside the required range.

## Examples

### Example 1

Input

```text
num1 = 120
num2 = 130
```

Output

```text
3
```

Explanation

* 120 → middle digit 2 is a peak
* 121 → middle digit 2 is a peak
* 130 → middle digit 3 is a peak

Total = 3

---

### Example 2

Input

```text
num1 = 198
num2 = 202
```

Output

```text
3
```

Explanation

* 198 → 9 is a peak
* 201 → 0 is a valley
* 202 → 0 is a valley

Total = 3

---

### Example 3

Input

```text
num1 = 4848
num2 = 4848
```

Output

```text
2
```

Explanation

* Second digit 8 is a peak
* Third digit 4 is a valley

Total waviness = 2

## How to Use / Run Locally

### C++

Compile

```bash
g++ solution.cpp -O2 -std=c++17
```

Run

```bash
./a.out
```

### Java

Compile

```bash
javac Solution.java
```

Run

```bash
java Solution
```

### JavaScript

Run

```bash
node solution.js
```

### Python3

Run

```bash
python solution.py
```

### Go

Run

```bash
go run solution.go
```

## Notes & Optimizations

* Brute force is completely infeasible for the given constraints.
* Digit DP is the intended solution.
* Only the previous two digits are required to determine future waviness contributions.
* Memoization dramatically reduces repeated calculations.
* The solution scales efficiently even near the maximum constraint.
* Leading zeros must be handled carefully to avoid counting invalid peaks or valleys.
* Using range subtraction is much faster than evaluating every number individually.
* The approach is suitable for competitive programming contests and advanced Dynamic Programming interviews.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
