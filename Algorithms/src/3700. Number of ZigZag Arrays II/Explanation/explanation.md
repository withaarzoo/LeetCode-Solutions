# 3700. Number of ZigZag Arrays II

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

---

## Problem Summary

**Number of ZigZag Arrays II** is a Dynamic Programming and Matrix Exponentiation problem where we need to count the total number of valid ZigZag arrays of length `n`.

Each element of the array must:

* Stay within the range `[l, r]`
* Be different from its adjacent element
* Never create three consecutive elements that are strictly increasing
* Never create three consecutive elements that are strictly decreasing

Since the length of the array can be extremely large (`10^9`), a normal DP solution is far too slow. The challenge is to find a way to represent transitions efficiently and compute the answer using fast matrix exponentiation.

The final answer must be returned modulo `10^9 + 7`.

This problem combines:

* Dynamic Programming
* State Compression
* Matrix Exponentiation
* Graph Transitions
* Modular Arithmetic

---

## Constraints

| Constraint         | Value             |
| ------------------ | ----------------- |
| `3 <= n <= 10^9`   | Array length      |
| `1 <= l < r <= 75` | Value range       |
| Answer             | Modulo `10^9 + 7` |

---

## Intuition

The first thing I noticed was that the condition about three consecutive increasing or decreasing elements completely controls the shape of the array.

If I ever have:

```text
1 < 3 < 5
```

then I create a strictly increasing sequence of length three, which is invalid.

Similarly:

```text
5 > 3 > 1
```

is also invalid.

That means after moving upward once, the next move must go downward.

And after moving downward once, the next move must go upward.

So instead of thinking about the whole array, I only need to remember:

1. The current value
2. Whether the next comparison should go up or down

That creates a very small state space because the value range is at most 75.

Since `n` can be as large as `10^9`, matrix exponentiation becomes the natural choice.

---

## Approach

1. Let:

```text
m = r - l + 1
```

1. Create two states for every value:

   * `up[x]`
   * `down[x]`

2. `up[x]` means:

   * Current value is `x`
   * Next value must be larger

3. `down[x]` means:

   * Current value is `x`
   * Next value must be smaller

4. Build a transition matrix.

5. Add transitions:

   * `up[x] → down[y]` whenever `y > x`
   * `down[x] → up[y]` whenever `y < x`

6. Create an initial state vector containing all valid starting states.

7. Compute:

```text
T^(n - 1)
```

using binary exponentiation.

1. Apply the resulting matrix to the starting vector.

2. Sum all valid ending states.

3. Return the answer modulo `10^9 + 7`.

---

## Data Structures Used

### Matrix

Used to represent state transitions.

Each cell stores how one state can move to another state.

### Vector

Used to represent the current count of ways to reach each state.

### 2D Arrays

Used internally for matrix multiplication and exponentiation.

---

## Operations & Behavior Summary

1. Calculate the number of possible values.
2. Create two states for every value.
3. Build all valid transitions.
4. Store transitions inside a matrix.
5. Create an identity matrix.
6. Raise the transition matrix to power `n - 1`.
7. Multiply transitions using fast exponentiation.
8. Compute the final state counts.
9. Sum all reachable states.
10. Return the answer modulo `10^9 + 7`.

---

## Complexity

| Type             | Complexity          | Explanation                                                     |
| ---------------- | ------------------- | --------------------------------------------------------------- |
| Time Complexity  | `O((2m)^3 × log n)` | Matrix multiplication is performed during binary exponentiation |
| Space Complexity | `O((2m)^2)`         | Required for storing transition matrices                        |

Where:

```text
m = r - l + 1
```

and

```text
m ≤ 75
```

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    static const long long MOD = 1000000007LL;

    // Multiply two matrices
    vector<vector<long long>> multiply(
        const vector<vector<long long>>& A,
        const vector<vector<long long>>& B
    ) {
        int sz = A.size();

        vector<vector<long long>> C(sz, vector<long long>(sz, 0));

        for (int i = 0; i < sz; i++) {
            for (int k = 0; k < sz; k++) {
                if (A[i][k] == 0) continue;

                long long cur = A[i][k];

                for (int j = 0; j < sz; j++) {
                    if (B[k][j] == 0) continue;

                    C[i][j] = (C[i][j] + cur * B[k][j]) % MOD;
                }
            }
        }

        return C;
    }

    int zigZagArrays(int n, int l, int r) {
        int m = r - l + 1;
        int sz = 2 * m;

        vector<vector<long long>> T(sz, vector<long long>(sz, 0));

        // State layout:
        // [0 ... m-1]       => up[x]
        // [m ... 2m-1]      => down[x]

        for (int x = 0; x < m; x++) {

            // up[x] -> down[y] for y > x
            for (int y = x + 1; y < m; y++) {
                T[x][m + y] = 1;
            }

            // down[x] -> up[y] for y < x
            for (int y = 0; y < x; y++) {
                T[m + x][y] = 1;
            }
        }

        // Identity matrix
        vector<vector<long long>> result(sz, vector<long long>(sz, 0));
        for (int i = 0; i < sz; i++) {
            result[i][i] = 1;
        }

        long long power = n - 1;

        // Fast exponentiation
        while (power > 0) {
            if (power & 1) {
                result = multiply(result, T);
            }

            T = multiply(T, T);
            power >>= 1;
        }

        // Initial vector is all ones
        vector<long long> initial(sz, 1);

        long long answer = 0;

        // Sum of all entries in result * initial
        for (int i = 0; i < sz; i++) {
            long long rowSum = 0;

            for (int j = 0; j < sz; j++) {
                rowSum = (rowSum + result[i][j]) % MOD;
            }

            answer = (answer + rowSum) % MOD;
        }

        return (int)answer;
    }
};
```

### Java

```java
class Solution {
    private static final long MOD = 1_000_000_007L;

    // Multiply two matrices
    private long[][] multiply(long[][] A, long[][] B) {
        int sz = A.length;

        long[][] C = new long[sz][sz];

        for (int i = 0; i < sz; i++) {
            for (int k = 0; k < sz; k++) {
                if (A[i][k] == 0) continue;

                long cur = A[i][k];

                for (int j = 0; j < sz; j++) {
                    if (B[k][j] == 0) continue;

                    C[i][j] = (C[i][j] + cur * B[k][j]) % MOD;
                }
            }
        }

        return C;
    }

    public int zigZagArrays(int n, int l, int r) {
        int m = r - l + 1;
        int sz = 2 * m;

        long[][] T = new long[sz][sz];

        for (int x = 0; x < m; x++) {

            for (int y = x + 1; y < m; y++) {
                T[x][m + y] = 1;
            }

            for (int y = 0; y < x; y++) {
                T[m + x][y] = 1;
            }
        }

        long[][] result = new long[sz][sz];
        for (int i = 0; i < sz; i++) {
            result[i][i] = 1;
        }

        long power = n - 1;

        while (power > 0) {
            if ((power & 1) == 1) {
                result = multiply(result, T);
            }

            T = multiply(T, T);
            power >>= 1;
        }

        long answer = 0;

        for (int i = 0; i < sz; i++) {
            long rowSum = 0;

            for (int j = 0; j < sz; j++) {
                rowSum = (rowSum + result[i][j]) % MOD;
            }

            answer = (answer + rowSum) % MOD;
        }

        return (int) answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @param {number} l
 * @param {number} r
 * @return {number}
 */
var zigZagArrays = function(n, l, r) {
    const MOD = 1000000007n;

    const m = r - l + 1;
    const sz = 2 * m;

    // Matrix multiplication
    const multiply = (A, B) => {
        const C = Array.from(
            { length: sz },
            () => Array(sz).fill(0n)
        );

        for (let i = 0; i < sz; i++) {
            for (let k = 0; k < sz; k++) {
                if (A[i][k] === 0n) continue;

                const cur = A[i][k];

                for (let j = 0; j < sz; j++) {
                    if (B[k][j] === 0n) continue;

                    C[i][j] = (C[i][j] + cur * B[k][j]) % MOD;
                }
            }
        }

        return C;
    };

    let T = Array.from(
        { length: sz },
        () => Array(sz).fill(0n)
    );

    for (let x = 0; x < m; x++) {

        // up[x] -> down[y]
        for (let y = x + 1; y < m; y++) {
            T[x][m + y] = 1n;
        }

        // down[x] -> up[y]
        for (let y = 0; y < x; y++) {
            T[m + x][y] = 1n;
        }
    }

    let result = Array.from(
        { length: sz },
        (_, i) =>
            Array.from(
                { length: sz },
                (_, j) => (i === j ? 1n : 0n)
            )
    );

    let power = BigInt(n - 1);

    while (power > 0n) {
        if (power & 1n) {
            result = multiply(result, T);
        }

        T = multiply(T, T);
        power >>= 1n;
    }

    let answer = 0n;

    for (let i = 0; i < sz; i++) {
        let rowSum = 0n;

        for (let j = 0; j < sz; j++) {
            rowSum = (rowSum + result[i][j]) % MOD;
        }

        answer = (answer + rowSum) % MOD;
    }

    return Number(answer);
};
```

### Python3

```python
class Solution:
    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        MOD = 1000000007

        m = r - l + 1
        sz = 2 * m

        # Matrix multiplication
        def multiply(A, B):
            C = [[0] * sz for _ in range(sz)]

            for i in range(sz):
                for k in range(sz):
                    if A[i][k] == 0:
                        continue

                    cur = A[i][k]

                    for j in range(sz):
                        if B[k][j] == 0:
                            continue

                        C[i][j] = (C[i][j] + cur * B[k][j]) % MOD

            return C

        T = [[0] * sz for _ in range(sz)]

        for x in range(m):

            # up[x] -> down[y]
            for y in range(x + 1, m):
                T[x][m + y] = 1

            # down[x] -> up[y]
            for y in range(x):
                T[m + x][y] = 1

        result = [[0] * sz for _ in range(sz)]
        for i in range(sz):
            result[i][i] = 1

        power = n - 1

        while power:
            if power & 1:
                result = multiply(result, T)

            T = multiply(T, T)
            power >>= 1

        answer = 0

        for i in range(sz):
            row_sum = sum(result[i]) % MOD
            answer = (answer + row_sum) % MOD

        return answer
```

### Go

```go
func zigZagArrays(n int, l int, r int) int {
 const MOD int64 = 1000000007

 m := r - l + 1
 sz := 2 * m

 // Matrix multiplication
 multiply := func(A, B [][]int64) [][]int64 {
  C := make([][]int64, sz)

  for i := 0; i < sz; i++ {
   C[i] = make([]int64, sz)

   for k := 0; k < sz; k++ {
    if A[i][k] == 0 {
     continue
    }

    cur := A[i][k]

    for j := 0; j < sz; j++ {
     if B[k][j] == 0 {
      continue
     }

     C[i][j] = (C[i][j] + cur*B[k][j]) % MOD
    }
   }
  }

  return C
 }

 T := make([][]int64, sz)

 for i := 0; i < sz; i++ {
  T[i] = make([]int64, sz)
 }

 for x := 0; x < m; x++ {

  // up[x] -> down[y]
  for y := x + 1; y < m; y++ {
   T[x][m+y] = 1
  }

  // down[x] -> up[y]
  for y := 0; y < x; y++ {
   T[m+x][y] = 1
  }
 }

 result := make([][]int64, sz)

 for i := 0; i < sz; i++ {
  result[i] = make([]int64, sz)
  result[i][i] = 1
 }

 power := n - 1

 for power > 0 {
  if power&1 == 1 {
   result = multiply(result, T)
  }

  T = multiply(T, T)
  power >>= 1
 }

 var answer int64 = 0

 for i := 0; i < sz; i++ {
  var rowSum int64 = 0

  for j := 0; j < sz; j++ {
   rowSum = (rowSum + result[i][j]) % MOD
  }

  answer = (answer + rowSum) % MOD
 }

 return int(answer)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same in all five languages.

### Step 1: Create States

For every possible value, create two states.

```text
up[value]
down[value]
```

This allows us to remember which direction must be chosen next.

---

### Step 2: Build Transitions

For every state:

If the current state is:

```text
up[x]
```

the next value must be larger.

Therefore:

```text
up[x] → down[y]
```

for every:

```text
y > x
```

After moving upward, the direction flips.

---

If the current state is:

```text
down[x]
```

the next value must be smaller.

Therefore:

```text
down[x] → up[y]
```

for every:

```text
y < x
```

Again the direction flips.

---

### Step 3: Create Transition Matrix

All valid transitions are stored inside a square matrix.

```text
matrix[from][to]
```

A value of 1 means the transition is allowed.

A value of 0 means it is not.

---

### Step 4: Initialize Starting States

Every value can start as:

```text
up[value]
```

or

```text
down[value]
```

because the second element has not been chosen yet.

So all states initially contain one valid way.

---

### Step 5: Use Matrix Exponentiation

Instead of applying transitions one by one:

```text
T × T × T × ...
```

we repeatedly square the matrix.

```text
T
T²
T⁴
T⁸
T¹⁶
...
```

This reduces the number of operations dramatically.

---

### Step 6: Apply the Final Matrix

After computing:

```text
T^(n - 1)
```

we apply it to the starting vector.

This gives the total number of ways to build arrays of length `n`.

---

### Step 7: Sum Every State

Any valid ending state contributes to the answer.

So we sum all state counts.

```text
answer = sum(all states)
```

and return:

```text
answer % (10^9 + 7)
```

---

## Examples

### Example 1

**Input**

```text
n = 3
l = 4
r = 5
```

**Output**

```text
2
```

**Valid Arrays**

```text
[4,5,4]
[5,4,5]
```

Only two alternating patterns exist.

---

### Example 2

**Input**

```text
n = 3
l = 1
r = 3
```

**Output**

```text
10
```

**Explanation**

Examples include:

```text
[1,2,1]
[1,3,1]
[2,1,2]
[3,1,3]
```

and several others that satisfy the ZigZag property.

---

### Example 3

**Input**

```text
n = 4
l = 1
r = 2
```

**Output**

```text
2
```

**Valid Arrays**

```text
[1,2,1,2]
[2,1,2,1]
```

The direction alternates perfectly.

---

## How to Use / Run Locally

### C++

Compile:

```bash
g++ solution.cpp -O2 -std=c++17
```

Run:

```bash
./a.out
```

---

### Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

---

### JavaScript

Run:

```bash
node solution.js
```

---

### Python3

Run:

```bash
python solution.py
```

---

### Go

Run:

```bash
go run solution.go
```

---

## Notes & Optimizations

* A normal Dynamic Programming solution will not work because `n` can be as large as `10^9`.
* The value range is very small (`≤ 75`), making state compression practical.
* Matrix exponentiation reduces the dependence on `n` from linear to logarithmic.
* Binary exponentiation is the key optimization that makes the solution pass.
* Modular arithmetic is required throughout all calculations.
* The transition graph remains fixed regardless of `n`, making matrix methods ideal.
* The solution efficiently handles the largest possible input constraints.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
