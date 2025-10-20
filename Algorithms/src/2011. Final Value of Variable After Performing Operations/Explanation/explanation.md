# Final Value of Variable After Performing Operations (LeetCode 2011)

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

We are given an array of strings `operations`. Each element is one of the following four strings: `"++X"`, `"X++"`, `"--X"`, or `"X--"`. The variable `X` starts at `0`. Each `"++X"` or `"X++"` operation increments `X` by `1`. Each `"--X"` or `"X--"` operation decrements `X` by `1`. Return the final value of `X` after performing all operations.

---

## Constraints

* `1 <= operations.length <= 100`
* Each `operations[i]` is one of: `"++X"`, `"X++"`, `"--X"`, or `"X--"`.

---

## Intuition

I thought this is straightforward: each operation either adds `1` or subtracts `1` from `X`. So I can iterate over the operations once, detect whether it's an increment or decrement, and update a single integer. No need for extra arrays or maps. Each operation string has fixed small length (3), so string checks are constant-time.

---

## Approach

1. Initialize `X = 0`.
2. Loop through each string `op` in `operations`.
3. If `op` indicates increment (contains `'+'`), do `X += 1`.
4. Otherwise do `X -= 1`.
5. Return `X`.

This is a single-pass O(n) solution with O(1) extra space.

---

## Data Structures Used

* A single integer variable `X` to store the current value.
* Input array/string list `operations` (provided by caller).

No additional data structures are needed.

---

## Operations & Behavior Summary

* `"++X"` or `"X++"` → increment `X` by 1.
* `"--X"` or `"X--"` → decrement `X` by 1.
* We treat each operation independently and apply changes sequentially.

---

## Complexity

* **Time Complexity:** `O(n)` — where `n` is the number of operations in the input array. Each operation is checked in constant time (strings are length 3).
* **Space Complexity:** `O(1)` — only a single integer `X` is used in addition to input.

---

## Multi-language Solutions

### C++

```c++
#include <vector>
#include <string>
using namespace std;

class Solution {
public:
    int finalValueAfterOperations(vector<string>& operations) {
        int X = 0; // initial value
        for (const string &op : operations) {
            // If the operation contains '+', it's an increment; otherwise it's a decrement.
            if (op.find('+') != string::npos) X++;
            else X--;
        }
        return X;
    }
};
```

### Java

```java
class Solution {
    public int finalValueAfterOperations(String[] operations) {
        int X = 0; // initial value
        for (String op : operations) {
            // check if operation includes '+'
            if (op.indexOf('+') != -1) X++;
            else X--;
        }
        return X;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string[]} operations
 * @return {number}
 */
var finalValueAfterOperations = function(operations) {
    let X = 0; // initial value
    for (let op of operations) {
        // if op contains '+' => increment, else decrement
        if (op.indexOf('+') !== -1) X++;
        else X--;
    }
    return X;
};
```

### Python3

```python
from typing import List

class Solution:
    def finalValueAfterOperations(self, operations: List[str]) -> int:
        X = 0  # initial value
        for op in operations:
            if '+' in op:  # increment operation
                X += 1
            else:           # decrement operation
                X -= 1
        return X
```

### Go

```go
package main

func finalValueAfterOperations(operations []string) int {
    X := 0 // initial value
    for _, op := range operations {
        // if '+' appears it's increment, otherwise decrement
        if containsPlus(op) {
            X++
        } else {
            X--
        }
    }
    return X
}

// helper checks for '+' in a short string
func containsPlus(s string) bool {
    for i := 0; i < len(s); i++ {
        if s[i] == '+' {
            return true
        }
    }
    return false
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Below I explain the key logic and each important line in a simple, step-by-step way.

### Common logic (applies to all languages)

1. Start with `X = 0` because problem states the variable starts at zero.
2. For every operation (string) in the input:

   * Check whether the string represents increment or decrement.
   * I choose to check for `'+'` character. If `'+'` exists, it must be `"++X"` or `"X++"` → increment.
   * Otherwise it is `"--X"` or `"X--"` → decrement.
3. Return the final `X`.

### C++ notes

* `op.find('+') != string::npos` returns true if `'+'` is present.
* Use `const string&` in loop to avoid copying strings.

### Java notes

* `op.indexOf('+') != -1` checks presence of `'+'`.
* Strings are immutable; using enhanced for-loop is clean and efficient.

### JavaScript notes

* `op.indexOf('+') !== -1` or `op.includes('+')` can be used. `includes` is fine too for modern JS.
* Use `let X = 0` and `for...of` for readability.

### Python3 notes

* `if '+' in op:` is concise and readable.
* `List[str]` typing is optional but helpful.

### Go notes

* Strings are byte arrays; loop bytes to check `s[i] == '+'`.
* Keep helper `containsPlus` for clarity and avoid importing `strings` package for this tiny check (both are fine).

---

## Examples

**Example 1**

```
Input: operations = ["--X","X++","X++"]
Process:
 X=0
 "--X" -> X = -1
 "X++" -> X = 0
 "X++" -> X = 1
Output: 1
```

**Example 2**

```
Input: operations = ["++X","++X","X++"]
Output: 3
Explanation: all are increments -> final X = 3
```

**Example 3**

```
Input: operations = ["X++","++X","--X","X--"]
Process yields X = 0
Output: 0
```

---

## How to use / Run locally

### C++

1. Copy the `class Solution` into your LeetCode C++ template or a local `.cpp` file with a `main()` wrapper for custom testing.
2. Compile with: `g++ -std=c++17 solution.cpp -o solution`
3. Run: `./solution` (if you add a `main` to call the method and print output).

### Java

1. Place the `Solution` class in a `Solution.java` file.
2. If running locally, add a `main` method to test the function.
3. Compile: `javac Solution.java`
4. Run: `java Solution`

### JavaScript

1. Paste the function into Node or a browser console.
2. To run with Node, create `solution.js` and include a test harness that calls `finalValueAfterOperations`.
3. Run: `node solution.js`

### Python3

1. Save the `Solution` class to `solution.py` and add a test harness:

```python
if __name__ == "__main__":
    ops = ["--X","X++","X++"]
    print(Solution().finalValueAfterOperations(ops))
```

2. Run: `python3 solution.py`

### Go

1. Create `main.go` and wrap the function in a `main` that calls `finalValueAfterOperations` and prints results.
2. Build and run:

```
go run main.go
```

---

## Notes & Optimizations

* Because each operation string is length 3, checking for `'+'` or `' -'` is constant time — no extra micro-optimizations are required.
* Alternatively, checking `op[1]` (middle character) would also work: for `"++X"` and `"--X"` middle char is `+`/`-`. But since `"X++"` and `"X--"` have middle char also `+`/`-`, `op[1]` is sufficient and slightly faster (direct index instead of search). Example micro-optimization:

  * `if (op[1] == '+') X++; else X--;`
* This optimization is safe here because every string has length 3 and valid format.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
