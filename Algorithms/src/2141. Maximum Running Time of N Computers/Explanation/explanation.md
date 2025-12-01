* # 2141. Maximum Running Time of N Computers

---

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
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

We have:

* `n` computers.
* An array `batteries`, where `batteries[i]` is the number of minutes the `i`-th battery can power **one** computer.

We can:

* Insert at most **one battery per computer** at the start.
* At any later time, we can **instantly swap** batteries between computers as many times as we want.
* Batteries cannot be recharged.

Goal:

> Find the **maximum number of minutes** we can run **all `n` computers simultaneously** using the given batteries.

Return that maximum running time.

---

## Constraints

* `1 <= n <= batteries.length <= 10^5`
* `1 <= batteries[i] <= 10^9`

These constraints force us to use an algorithm roughly `O(m log something)` where `m = batteries.length`.
Anything like simulating minute-by-minute or heavy DP will be too slow.

---

## Intuition

I asked myself a simple question:

> “If I want every computer to run for exactly `T` minutes, what must be true?”

Each of the `n` computers needs `T` minutes → total required power = `n * T`.

Because I can **freely swap** batteries:

* From a battery with capacity `b`, I never need more than `T` minutes for this target.
* So each battery can contribute at most `min(b, T)` minutes toward those `T` minutes.

So the **total usable minutes for target time `T`** is:

[
\text{usable}(T) = \sum_i \min(batteries[i], T)
]

For `T` to be possible:

[
\sum_i \min(batteries[i], T) \ge n \cdot T
]

Also, if we can run for `T` minutes, we can surely run for any smaller time `T' < T`.
So the feasibility of `T` is **monotonic** → perfect case for **binary search on T**.

---

## Approach

1. **Total Energy and Max Bound**

   * Compute `total = sum(batteries)`.
   * Even if we distribute power perfectly, we cannot run each computer for more than `total / n` minutes.
   * So answer `T` lies in range `[0, total / n]`.

2. **Binary Search on Time**

   * Set `low = 0`, `high = total / n`.
   * While `low < high`:

     * Take `mid = (low + high + 1) / 2` (upper mid).
     * Check if running all computers for `mid` minutes is possible.

3. **Feasibility Check for Time `mid`**

   * Traverse `batteries`.
   * Maintain `usable = 0`.
   * For each battery `b`:

     * Add `min(b, mid)` to `usable`.
   * If at any point `usable >= mid * n`, we can stop early.
   * If after the loop `usable >= mid * n`, then time `mid` is **feasible**.

4. **Adjust Search Range**

   * If `mid` is feasible → `low = mid` (try longer time).
   * Else → `high = mid - 1` (time too long).

5. **Answer**

   * When loop ends, `low == high` and it is the maximum feasible running time.

We must use 64-bit integers (like `long long`, `long`, `int64`) because sums can go up to about `10^14`.

---

## Data Structures Used

* Just the given array `batteries`.
* A few scalar variables:

  * `total` – sum of all battery times.
  * `low`, `high`, `mid` – for binary search.
  * `usable` – sum of `min(b[i], mid)` while checking feasibility.

No extra arrays or complex data structures are required.

---

## Operations & Behavior Summary

* **Summation**: One pass over `batteries` to compute the total capacity.
* **Binary Search**:

  * Repeatedly pick a candidate time `mid`.
  * For each `mid`, scan `batteries` and add `min(b[i], mid)` to `usable`.
  * Decide if `mid` is feasible and adjust search range.

This process moves us to the maximum `T` where the feasibility condition holds.

---

## Complexity

Let `m = batteries.length`.

* **Time Complexity:** `O(m * log (total / n))`

  * Each feasibility check scans all `m` batteries.
  * We do about `log2(total / n)` checks due to binary search.
* **Space Complexity:** `O(1)`

  * Besides a few integer variables, we don’t use any additional data structures.

---

## Multi-language Solutions

### C++

```cpp
#include <vector>
#include <numeric>
using namespace std;

class Solution {
public:
    long long maxRunTime(int n, vector<int>& batteries) {
        // Total sum of all battery capacities
        long long total = 0;
        for (int b : batteries) total += b;

        // Upper bound on the maximum time per computer
        long long low = 0;
        long long high = total / n;

        // Binary search for the maximum feasible time
        while (low < high) {
            long long mid = low + (high - low + 1) / 2; // upper mid

            long long usable = 0;
            for (int b : batteries) {
                // Each battery contributes at most 'mid' minutes
                usable += min<long long>(b, mid);
                if (usable >= mid * n) break; // early exit if enough
            }

            if (usable >= mid * n) {
                // We can run all n computers for 'mid' minutes
                low = mid;  // try longer time
            } else {
                // 'mid' is too large
                high = mid - 1;
            }
        }

        return low;
    }
};
```

---

### Java

```java
class Solution {
    public long maxRunTime(int n, int[] batteries) {
        long total = 0L;
        for (int b : batteries) total += b;

        long low = 0L;
        long high = total / n; // maximum possible time per computer

        while (low < high) {
            long mid = low + (high - low + 1) / 2; // upper mid

            long usable = 0L;
            for (int b : batteries) {
                // Each battery contributes at most 'mid' minutes
                usable += Math.min((long) b, mid);
                if (usable >= mid * n) break;
            }

            if (usable >= mid * n) {
                // 'mid' minutes is feasible
                low = mid;
            } else {
                // 'mid' minutes is not feasible
                high = mid - 1;
            }
        }

        return low;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number} n
 * @param {number[]} batteries
 * @return {number}
 */
var maxRunTime = function(n, batteries) {
    // Sum all battery capacities
    let total = 0;
    for (const b of batteries) total += b;

    let low = 0;
    let high = Math.floor(total / n); // upper bound on answer

    while (low < high) {
        const mid = Math.floor((low + high + 1) / 2); // upper mid

        let usable = 0;
        for (const b of batteries) {
            // Each battery contributes at most 'mid'
            usable += Math.min(b, mid);
            if (usable >= mid * n) break;
        }

        if (usable >= mid * n) {
            // 'mid' is feasible
            low = mid;
        } else {
            // 'mid' is too large
            high = mid - 1;
        }
    }

    return low;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def maxRunTime(self, n: int, batteries: List[int]) -> int:
        # Total sum of battery capacities
        total = sum(batteries)

        # Upper bound: perfect distribution
        low, high = 0, total // n

        # Binary search on time
        while low < high:
            mid = (low + high + 1) // 2  # upper mid

            usable = 0
            for b in batteries:
                # Each battery contributes at most 'mid' minutes
                usable += min(b, mid)
                if usable >= mid * n:
                    break

            if usable >= mid * n:
                # mid minutes is possible
                low = mid
            else:
                # mid minutes is not possible
                high = mid - 1

        return low
```

---

### Go

```go
package main

func maxRunTime(n int, batteries []int) int64 {
    // Sum all battery capacities
    var total int64 = 0
    for _, b := range batteries {
        total += int64(b)
    }

    var low int64 = 0
    var high int64 = total / int64(n) // upper bound

    for low < high {
        mid := low + (high-low+1)/2 // upper mid

        var usable int64 = 0
        for _, b := range batteries {
            if int64(b) < mid {
                usable += int64(b)
            } else {
                usable += mid
            }
            if usable >= mid*int64(n) {
                break
            }
        }

        if usable >= mid*int64(n) {
            low = mid
        } else {
            high = mid - 1
        }
    }

    return low
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Since all languages follow the **same logic**, I’ll explain it once and you can map it to any code version.

### 1. Sum of all batteries

Example (Python, similar in others):

```python
total = sum(batteries)
```

* I add all capacities together.
* This tells me the **total energy** available.
* No matter what, total running time for all computers combined can’t exceed this.

### 2. Setting the binary search range

```python
low, high = 0, total // n
```

* `low` = definitely possible time. `0` is always possible.
* `high` = maximum possible time if energy is perfectly shared: `total / n`.
* Our answer lies somewhere in this range.

### 3. Binary search loop

```python
while low < high:
    mid = (low + high + 1) // 2
    ...
```

* I choose `mid` as the **candidate time**.
* I use the **upper mid** formula `(low + high + 1) // 2` to avoid infinite loops when `low` and `high` are close.
* Now I just need to check: “Can we run all computers for `mid` minutes?”

### 4. Computing usable energy for time `mid`

```python
usable = 0
for b in batteries:
    usable += min(b, mid)
    if usable >= mid * n:
        break
```

* For each battery `b`:

  * It can help for at most `mid` minutes in our `mid`-minute plan.
  * So contribution is `min(b, mid)`.
* I keep adding these contributions to `usable`.
* As soon as `usable` reaches `mid * n`, I know we have enough power, so I break early.

### 5. Checking feasibility of `mid`

```python
if usable >= mid * n:
    low = mid
else:
    high = mid - 1
```

* Required total power for `mid` minutes is `mid * n`.
* If `usable >= mid * n`:

  * It means our current batteries can support `mid` minutes.
  * So I move `low` up to `mid` and try for a bigger time next.
* Otherwise:

  * We can’t reach `mid` minutes.
  * So I reduce `high` to `mid - 1` to try a smaller time.

### 6. Final answer

When `low == high`, the binary search finishes:

* This `low` is the **maximum** `T` for which we passed the feasibility test.
* I return `low`.

The same logic is implemented in:

* C++: using `long long` and `min<long long>`.
* Java: using `long` and `Math.min`.
* JavaScript: using `number` with careful integer arithmetic.
* Python: using Python’s built-in big integers.
* Go: using `int64`.

---

## Examples

### Example 1

```text
Input: n = 2, batteries = [3,3,3]
Output: 4
```

Explanation (high-level):

* Total energy = 9.
* `total / n = 9 / 2 = 4` (integer division).
* We find that 4 minutes is feasible but 5 is not.
* So the maximum time is 4.

### Example 2

```text
Input: n = 2, batteries = [1,1,1,1]
Output: 2
```

Explanation:

* Total energy = 4.
* `total / n = 4 / 2 = 2`.
* We can run both computers for at most 2 minutes simultaneously.

---

## How to use / Run locally

1. **Clone your repo** (assuming this README is in that repo):

```bash
git clone <your-repo-url>
cd <your-repo-folder>
```

2. **Compile & Run**

* **C++** (with `solution.cpp`):

```bash
g++ -std=c++17 solution.cpp -o solution
./solution
```

* **Java** (with `Solution.java`):

```bash
javac Solution.java
java Solution
```

* **JavaScript** (Node.js):

```bash
node solution.js
```

* **Python3**:

```bash
python3 solution.py
```

* **Go**:

```bash
go run solution.go
```

3. Plug the test values (like examples) into a small driver `main` function or use online judges / LeetCode directly.

---

## Notes & Optimizations

* We **don’t** need to sort the batteries.
  Sorting is unnecessary because our feasibility check only depends on `min(b[i], mid)` and summation.
* We use **64-bit integers** to avoid overflow for:

  * `total`
  * `mid * n`
  * `usable`
* The early break (`if usable >= mid * n: break`) saves time for big inputs, especially when candidate time `mid` is small.
* Binary search range is tight:

  * Lower bound = 0
  * Upper bound = `total / n`
    This keeps the number of iterations small (`~log2(1e9)` at worst).

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
