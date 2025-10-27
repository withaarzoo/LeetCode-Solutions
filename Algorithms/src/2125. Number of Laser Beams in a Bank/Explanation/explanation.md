# 2125. Number of Laser Beams in a Bank

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

We are given a 0-indexed binary string array `bank` that represents the floor plan of a bank. Each string in `bank` is a row. A character `'1'` denotes a security device in that cell and `'0'` means empty.

A laser beam exists between two devices if:

1. They are in two different rows `r1` and `r2` where `r1 < r2`.
2. Every row strictly between `r1` and `r2` has **no devices** (`'0'` only).

Return the total number of laser beams in the bank.

---

## Constraints

* `m == bank.length` (number of rows)
* `n == bank[i].length` (number of columns)
* `1 <= m, n <= 500`
* `bank[i][j]` is either `'0'` or `'1'`.

---

## Intuition

I thought about what actually creates a laser beam: two security devices form a beam only if they sit in different rows and **every row between them is empty**. That means:

* Devices in the same row never form beams with each other.
* If a row contains any device, it blocks beams between rows above and below it unless those rows are directly adjacent in the sequence of non-empty rows.

So I realized I only need to count the number of devices in each row and then pair each non-empty row with the **next** non-empty row. For each such pair, every device in the upper row connects to every device in the lower row, giving `count_upper * count_lower` beams.

---

## Approach

1. Process rows from top to bottom.
2. For each row, count the number of `'1'` characters (devices).
3. Keep track of the device count of the **previous non-empty row** (call it `prev`).
4. If the current row has `cnt` devices and `prev > 0`, then add `prev * cnt` to the answer.
5. If current row is non-empty, update `prev = cnt`. If the row is empty, skip it without changing `prev`.
6. Continue until all rows are processed. The sum accumulated is the final answer.

This is a single-pass solution (we only visit each character once for counting) and uses constant extra space.

---

## Data Structures Used

* Primitive counters (integers/longs).
* No extra arrays, lists, sets, or maps are required beyond reading the input.

---

## Operations & Behavior Summary

* Count `'1'`s in each row: `O(n)` per row where `n` is number of columns in that row.
* Maintain and update two variables:

  * `prev` — device count of previous non-empty row.
  * `ans` — accumulated number of beams.
* Multiply counts of consecutive non-empty rows and add to `ans`.

---

## Complexity

* **Time Complexity:** `O(m * n)` where `m` = number of rows (`bank.length`) and `n` = number of columns (length of each string). We examine each character once to count devices per row.
* **Space Complexity:** `O(1)` extra space (only constant number of integer variables). The input itself occupies `O(m*n)` but we don't allocate additional structures.

---

## Multi-language Solutions

### C++

```c++
class Solution {
public:
    int numberOfBeams(vector<string>& bank) {
        long long ans = 0;           // total beams
        long long prev = 0;          // device count in previous non-empty row
        for (auto &row : bank) {
            long long cnt = 0;
            for (char ch : row) if (ch == '1') ++cnt; // count devices
            if (cnt > 0) {
                ans += prev * cnt; // beams between prev non-empty and current
                prev = cnt;
            }
        }
        return (int)ans;
    }
};
```

### Java

```java
class Solution {
    public int numberOfBeams(String[] bank) {
        long ans = 0;     // total beams
        long prev = 0;    // devices in previous non-empty row
        for (String row : bank) {
            long cnt = 0;
            for (int i = 0; i < row.length(); ++i) {
                if (row.charAt(i) == '1') cnt++;
            }
            if (cnt > 0) {
                ans += prev * cnt;
                prev = cnt;
            }
        }
        return (int) ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string[]} bank
 * @return {number}
 */
var numberOfBeams = function(bank) {
    let ans = 0;
    let prev = 0;
    for (const row of bank) {
        let cnt = 0;
        for (let i = 0; i < row.length; ++i) if (row[i] === '1') cnt++;
        if (cnt > 0) {
            ans += prev * cnt;
            prev = cnt;
        }
    }
    return ans;
};
```

### Python3

```python3
from typing import List

class Solution:
    def numberOfBeams(self, bank: List[str]) -> int:
        ans = 0
        prev = 0
        for row in bank:
            cnt = row.count('1')
            if cnt > 0:
                ans += prev * cnt
                prev = cnt
        return ans
```

### Go

```go
package main

func numberOfBeams(bank []string) int {
    ans := 0
    prev := 0
    for _, row := range bank {
        cnt := 0
        for i := 0; i < len(row); i++ {
            if row[i] == '1' {
                cnt++
            }
        }
        if cnt > 0 {
            ans += prev * cnt
            prev = cnt
        }
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll walk through the algorithm and key lines in the code as if I'm teaching a friend.

### 1) The core idea

* We only care about rows that contain at least one device. Call those rows "non-empty".
* If row A (above) is non-empty with `p` devices and the next non-empty row B (below) has `c` devices, then every device in A pairs with every device in B because the rows between them are empty. That's `p * c` beams.
* We should only pair non-empty rows that are consecutive in the sequence of non-empty rows. A non-empty row blocks pairing between earlier and later rows.

### 2) Variables we maintain

* `ans`: the running total for number of beams.
* `prev`: the device count for the last non-empty row we encountered while scanning downwards.

### 3) Row processing (loop)

For each row:

* Count how many `'1'` characters exist (`cnt`). This is done via `row.count('1')` in Python or by iterating characters in other languages.
* If `cnt == 0` (empty row), do nothing — keep `prev` as-is.
* If `cnt > 0` (non-empty row):

  * `ans += prev * cnt` (this pairs current row with previous non-empty row)
  * `prev = cnt` (update `prev` to current row's device count for the next pairing)

### 4) Example walkthrough (short)

Given rows with counts `[3, 0, 2, 4]`:

* Start: `prev = 0`, `ans = 0`.
* Row1 has 3 -> `ans += 0 * 3 = 0`; `prev = 3`.
* Row2 has 0 -> skip.
* Row3 has 2 -> `ans += 3 * 2 = 6`; `prev = 2`.
* Row4 has 4 -> `ans += 2 * 4 = 8`; `prev = 4`.
* Total `ans = 14`.

This matches the rule: beams only occur between consecutive non-empty rows (Row1↔Row3 and Row3↔Row4 here).

### 5) Correctness & edge cases

* If all rows are empty or there's only one non-empty row, `ans` remains 0 because no two devices are on different rows satisfying the conditions.
* Multiplication of counts is safe with 32-bit ints for the constraints given, but we used 64-bit accumulators in some languages for safety.

---

## Examples

1. **Example 1**

```
Input: bank = ["011001","000000","010100","001000"]
Output: 8
```

Explanation: Counting devices per row gives `[3, 0, 2, 1]` (example interpretation). Pair consecutive non-empty rows: `3*2 + 2*1 = 6 + 2 = 8`.

2. **Example 2**

```
Input: bank = ["000","111","000"]
Output: 0
```

Explanation: There's only one non-empty row (`111`) so no beams exist.

---

## How to use / Run locally

1. Clone or copy this README into your problem repository.
2. Place the appropriate language file (`Solution.*`) into the local project or run in the platform's editor.

For local testing with small harnesses:

* **Python**: create a small runner script that instantiates `Solution()` and calls `numberOfBeams` with test arrays.
* **JavaScript (Node.js)**: create a .js file and call the function with sample input, then `console.log` the result.
* **C++/Java/Go**: write a `main` function or small test driver that reads hard-coded inputs and prints results.

---

## Notes & Optimizations

* This solution is already optimal in time and space for the problem constraints (`O(m*n)` time, `O(1)` extra space).
* If you need to micro-optimize, counting `1`'s can be done with bit tricks only if input representation changes. For string input, `row.count('1')` or a simple loop is the fastest and clearest.
* Use 64-bit accumulators (`long long`/`long`) in languages where integer overflow is a concern, though with given constraints typical 32-bit ints will still be safe.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
