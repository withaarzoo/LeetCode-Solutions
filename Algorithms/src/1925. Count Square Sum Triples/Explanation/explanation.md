# 1925. Count Square Sum Triples

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

A **square triple** `(a, b, c)` is a triple of integers such that:

```text
a^2 + b^2 = c^2
```

Given an integer `n`, we need to return the **number of square triples** such that:

```text
1 <= a, b, c <= n
```

Note: Triples are **ordered**. So `(3, 4, 5)` and `(4, 3, 5)` are counted as **two different** triples.

---

## Constraints

* `1 <= n <= 250`

The upper bound is small, so we can afford an `O(n²)` or even `O(n³)` algorithm, but we still prefer something clean and efficient.

---

## Intuition

I thought of this problem as counting **Pythagorean triples** within a limited range:

* We need integers `a`, `b`, and `c` between `1` and `n`.
* They must satisfy `a² + b² = c²`.

First idea:

* Use three loops: try every `a`, `b`, `c` and check the equation.
* That is `O(n³)`. For `n = 250`, this is still okay but not elegant.

Then I realized:

* For every pair `(a, b)`, the value `c²` is fixed as `a² + b²`.
* So instead of looping over `c`, I can **calculate** `c`:

  * `c = sqrt(a² + b²)`.

So the improved idea:

* Use only **two loops** (`a` and `b`).
* Compute `sum = a² + b²`.
* Take `c = floor(sqrt(sum))`.
* If `c <= n` and `c² == sum`, then we found a valid triple `(a, b, c)`.

This gives us an efficient and simple `O(n²)` solution.

---

## Approach

Step-by-step, my approach is:

1. Initialize `count = 0` to store the total number of valid triples.

2. Loop over all possible values of `a` from `1` to `n`.

3. Inside that, loop over all possible values of `b` from `1` to `n`.

4. For each pair `(a, b)`:

   * Compute `sumSquares = a * a + b * b`.
   * This is supposed to be `c²`.

5. Compute `c` as the **integer square root** of `sumSquares`:

   * In most languages: `c = floor(sqrt(sumSquares))`.
   * In Python, I can use `math.isqrt(sumSquares)` which is exact.

6. Check if:

   * `c <= n` (because `c` must be within the allowed range).
   * `c * c == sumSquares` (to ensure `sumSquares` is a perfect square).

7. If both conditions are true:

   * We have a valid triple `(a, b, c)`.
   * Increment `count`.

8. After the loops end, return `count`.

Because we iterate over `(a, b)` in order, we automatically count ordered triples like `(3, 4, 5)` and `(4, 3, 5)` separately.

---

## Data Structures Used

* Only simple integer variables:

  * `count` for the result.
  * Loop counters `a`, `b`.
  * `sumSquares` to store `a² + b²`.
  * `c` for the integer square root.

No arrays, no maps, no extra data structures.
So the solution is very memory-light.

---

## Operations & Behavior Summary

For every pair `(a, b)`:

1. Compute squares and add:

   * `a² + b²`.

2. Find integer square root:

   * `c = floor(sqrt(a² + b²))` (or `isqrt` in Python).

3. Validate:

   * If `c` is in range (`1 <= c <= n`)
   * And it exactly satisfies `c² = a² + b²`

4. If valid, increase counter.

At the end, the function returns how many such triples exist.

---

## Complexity

* **Time Complexity:**

  * Two loops: `a` in `[1, n]` and `b` in `[1, n]`.
  * Total pairs `(a, b)` = `n * n = n²`.
  * For each pair, we do constant-time work (multiplications + sqrt).
  * So the time complexity is:

    ```text
    O(n²)
    ```

* **Space Complexity:**

  * We only use a few integer variables.
  * No extra collections.
  * So the space complexity is:

    ```text
    O(1)   // constant extra memory
    ```

---

## Multi-language Solutions

### C++

```c++
#include <cmath>

class Solution {
public:
    int countTriples(int n) {
        int count = 0;

        // Try all possible pairs (a, b)
        for (int a = 1; a <= n; ++a) {
            for (int b = 1; b <= n; ++b) {
                int sumSquares = a * a + b * b;    // this should be c^2

                int c = static_cast<int>(std::sqrt(sumSquares)); // integer square root

                // Check if c is in range and forms a perfect square
                if (c <= n && c * c == sumSquares) {
                    count++; // (a, b, c) is a valid square triple
                }
            }
        }

        return count;
    }
};
```

---

### Java

```java
class Solution {
    public int countTriples(int n) {
        int count = 0;

        // Try all possible pairs (a, b)
        for (int a = 1; a <= n; a++) {
            for (int b = 1; b <= n; b++) {
                int sumSquares = a * a + b * b;   // this should be c^2

                int c = (int) Math.sqrt(sumSquares); // integer square root

                // Check if c is in range and forms a perfect square
                if (c <= n && c * c == sumSquares) {
                    count++; // (a, b, c) is a valid square triple
                }
            }
        }

        return count;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number} n
 * @return {number}
 */
var countTriples = function(n) {
    let count = 0;

    // Try all possible pairs (a, b)
    for (let a = 1; a <= n; a++) {
        for (let b = 1; b <= n; b++) {
            const sumSquares = a * a + b * b; // this should be c^2

            const c = Math.floor(Math.sqrt(sumSquares)); // integer square root

            // Check if c is in range and forms a perfect square
            if (c <= n && c * c === sumSquares) {
                count++; // (a, b, c) is a valid square triple
            }
        }
    }

    return count;
};
```

---

### Python3

```python
import math

class Solution:
    def countTriples(self, n: int) -> int:
        count = 0

        # Try all possible pairs (a, b)
        for a in range(1, n + 1):
            for b in range(1, n + 1):
                sum_squares = a * a + b * b   # this should be c^2

                # Exact integer square root (no floating-point issues)
                c = math.isqrt(sum_squares)

                # Check if c is in range and forms a perfect square
                if c <= n and c * c == sum_squares:
                    count += 1  # (a, b, c) is a valid square triple

        return count
```

---

### Go

```go
package main

import "math"

func countTriples(n int) int {
    count := 0

    // Try all possible pairs (a, b)
    for a := 1; a <= n; a++ {
        for b := 1; b <= n; b++ {
            sumSquares := a*a + b*b // this should be c^2

            c := int(math.Sqrt(float64(sumSquares))) // integer square root

            // Check if c is in range and forms a perfect square
            if c <= n && c*c == sumSquares {
                count++ // (a, b, c) is a valid square triple
            }
        }
    }

    return count
}
```

> For LeetCode in Go, just keep the function (without `package main`) as required by their template.

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Below is the **core logic** that is shared in all languages:

1. **Initialize answer:**

   ```text
   count = 0
   ```

   I use this to keep track of how many valid triples I find.

2. **Two nested loops for `a` and `b`:**

   * For C++ / Java:

     ```cpp
     for (int a = 1; a <= n; ++a) {
         for (int b = 1; b <= n; ++b) {
             ...
         }
     }
     ```

   * For JavaScript / Python / Go it’s syntactically different, but logically the same.

   Here I am going over **every ordered pair** `(a, b)` such that `1 <= a, b <= n`.

3. **Compute `a² + b²`:**

   ```text
   sumSquares = a * a + b * b
   ```

   This is the value that should equal `c²` if a triple exists.

4. **Compute `c` as integer square root:**

   * C++: `int c = (int)std::sqrt(sumSquares);`
   * Java: `int c = (int)Math.sqrt(sumSquares);`
   * JS: `const c = Math.floor(Math.sqrt(sumSquares));`
   * Python: `c = math.isqrt(sum_squares)` (clean and precise)
   * Go: `c := int(math.Sqrt(float64(sumSquares)))`

   This step finds the integer candidate `c` such that `c²` is **closest below or equal** to `sumSquares`.

5. **Check the constraints for `c`:**

   ```text
   if c <= n and c * c == sumSquares:
       count++
   ```

   * `c * c == sumSquares` ensures `sumSquares` is a **perfect square**.
   * `c <= n` ensures `c` is not outside our allowed range.
   * If this passes, `(a, b, c)` is a valid square triple.

6. **Return the final result:**

   After the nested loops complete, we simply return:

   ```text
   return count
   ```

---

## Examples

### Example 1

**Input:**

```text
n = 5
```

Valid triples:

* `(3, 4, 5)`
* `(4, 3, 5)`

**Output:**

```text
2
```

---

### Example 2

**Input:**

```text
n = 10
```

Valid triples:

* `(3, 4, 5)`
* `(4, 3, 5)`
* `(6, 8, 10)`
* `(8, 6, 10)`

**Output:**

```text
4
```

---

## How to use / Run locally

### C++

```bash
g++ -std=c++17 -O2 main.cpp -o main
./main
```

Where `main.cpp` contains:

```c++
#include <bits/stdc++.h>
using namespace std;

// Paste the Solution class here and call it from main()
int main() {
    Solution sol;
    int n;
    cin >> n;
    cout << sol.countTriples(n) << endl;
    return 0;
}
```

---

### Java

```bash
javac Main.java
java Main
```

`Main.java`:

```java
public class Main {
    public static void main(String[] args) {
        java.util.Scanner sc = new java.util.Scanner(System.in);
        int n = sc.nextInt();
        Solution sol = new Solution();
        System.out.println(sol.countTriples(n));
    }
}

// Paste the Solution class below
```

---

### JavaScript (Node.js)

```bash
node main.js
```

`main.js`:

```javascript
const fs = require("fs");

// Paste the countTriples function here

const input = fs.readFileSync(0, "utf8").trim();
const n = Number(input);
console.log(countTriples(n));
```

---

### Python3

```bash
python3 main.py
```

`main.py`:

```python
from sys import stdin
from math import isqrt

# Paste the Solution class here

if __name__ == "__main__":
    n = int(stdin.readline().strip())
    sol = Solution()
    print(sol.countTriples(n))
```

---

### Go

```bash
go run main.go
```

`main.go`:

```go
package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
)

// Paste countTriples function here

func main() {
    in := bufio.NewReader(os.Stdin)
    var n int
    fmt.Fscan(in, &n)
    fmt.Println(countTriples(n))
}
```

---

## Notes & Optimizations

* **Why not 3 loops?**
  We could loop `a`, `b`, and `c`, but that is `O(n³)`. Since we can derive `c` from `a` and `b`, we save one loop and bring it down to `O(n²)`.

* **Using integer square root:**

  * In Python, `math.isqrt` is perfect because it avoids floating-point issues entirely.
  * In other languages, converting `sqrt` to `int` is safe here because:

    * We always verify with `c * c == sumSquares`.

* **Ordered triples:**
  We intentionally keep both `(a, b, c)` and `(b, a, c)`.
  If the problem wanted unordered triples, we would restrict loops (e.g., `b` starting from `a`) or handle duplicates.

* **Constraint-friendly:**
  `n <= 250` means `n² = 62,500` iterations, which is extremely fast, even with multiple test runs.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
