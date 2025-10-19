# 1625. Lexicographically Smallest String After Applying Operations

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

You are given a string `s` of even length consisting of digits `'0'` to `'9'`, and two integers `a` and `b`. You can apply any number of operations in any order on `s`:

1. Add `a` to every digit at odd indices (0-indexed). Digits wrap around modulo 10.
2. Rotate the string to the right by `b` positions.

Return the lexicographically smallest string you can obtain by applying these operations any number of times.

---

## Constraints

* `2 <= s.length <= 100`
* `s.length` is even
* `s` consists of digits from `'0'` to `'9'`
* `1 <= a <= 9`
* `1 <= b <= s.length - 1`

---

## Intuition

I thought: there are only two deterministic operations that transform the string, and every result is a finite-length string of digits. This forms a finite graph of states (strings). Each state can go to at most two others (add-operation result and rotate-operation result). So I can search all reachable states starting from `s` and track the lexicographically smallest one. BFS (or DFS) with a `visited` set will explore every reachable string without infinite loops. Since `s.length` ≤ 100 and digits are limited, this is feasible and reliable.

---

## Approach

1. Treat each unique string as a node in a graph.
2. Use BFS from the initial string `s`. Maintain a `visited` set to avoid revisiting the same string.
3. For each string `cur` popped from the queue:

   * Compare with current best answer and update if smaller.
   * Compute the string after adding `a` to all odd indices.
   * Compute the string after rotating right by `b`.
   * If any new string is unseen, add it to `visited` and to the queue.
4. After BFS finishes, return the smallest seen string.

This guarantees the lexicographically smallest reachable string because we enumerate all reachable states.

---

## Data Structures Used

* `Queue` (BFS) — to explore states level by level (ArrayDeque / deque / queue).
* `Set` (`visited`) — to avoid processing a state twice and to prevent infinite loops.
* Strings / character arrays — to construct transformed states.

---

## Operations & Behavior Summary

* **Add Operation**: For every odd index `i` (1, 3, 5, ...), replace digit `d` with `(d + a) % 10`.
* **Rotate Operation**: Right rotation by `b`: `rotated = s[n-b:] + s[:n-b]`.

Both operations run in `O(n)` time where `n` is the string length.

---

## Complexity

* **Time Complexity:** `O(M * n)`, where `n` is the length of `s` and `M` is the number of distinct reachable strings. For each distinct string we perform two `O(n)` transformations. In practice `M` is limited by constraints and operation structure.
* **Space Complexity:** `O(M * n)` for storing visited strings and queue.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    string findLexSmallestString(string s, int a, int b) {
        int n = s.size();
        unordered_set<string> seen;
        queue<string> q;
        seen.insert(s);
        q.push(s);
        string ans = s;

        while (!q.empty()) {
            string cur = q.front(); q.pop();
            if (cur < ans) ans = cur;

            // Operation 1: add 'a' to odd indices
            string addOp = cur;
            for (int i = 1; i < n; i += 2) {
                int d = (addOp[i] - '0' + a) % 10;
                addOp[i] = char('0' + d);
            }
            if (!seen.count(addOp)) {
                seen.insert(addOp);
                q.push(addOp);
            }

            // Operation 2: rotate right by b
            string rotOp = cur.substr(n - b) + cur.substr(0, n - b);
            if (!seen.count(rotOp)) {
                seen.insert(rotOp);
                q.push(rotOp);
            }
        }
        return ans;
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {
    public String findLexSmallestString(String s, int a, int b) {
        int n = s.length();
        Set<String> seen = new HashSet<>();
        Queue<String> q = new ArrayDeque<>();
        seen.add(s);
        q.add(s);
        String ans = s;

        while (!q.isEmpty()) {
            String cur = q.poll();
            if (cur.compareTo(ans) < 0) ans = cur;

            // Operation 1: add a to odd indices
            char[] arr = cur.toCharArray();
            for (int i = 1; i < n; i += 2) {
                int d = (arr[i] - '0' + a) % 10;
                arr[i] = (char)('0' + d);
            }
            String addOp = new String(arr);
            if (!seen.contains(addOp)) {
                seen.add(addOp);
                q.add(addOp);
            }

            // Operation 2: rotate right by b
            String rotOp = cur.substring(n - b) + cur.substring(0, n - b);
            if (!seen.contains(rotOp)) {
                seen.add(rotOp);
                q.add(rotOp);
            }
        }
        return ans;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {string} s
 * @param {number} a
 * @param {number} b
 * @return {string}
 */
var findLexSmallestString = function(s, a, b) {
    const n = s.length;
    const seen = new Set();
    const q = [];
    seen.add(s);
    q.push(s);
    let ans = s;

    while (q.length) {
        const cur = q.shift();
        if (cur < ans) ans = cur;

        // add a to odd indices
        let ch = cur.split('');
        for (let i = 1; i < n; i += 2) {
            ch[i] = String((parseInt(ch[i]) + a) % 10);
        }
        const addOp = ch.join('');
        if (!seen.has(addOp)) {
            seen.add(addOp);
            q.push(addOp);
        }

        // rotate right by b
        const rotOp = cur.slice(n - b) + cur.slice(0, n - b);
        if (!seen.has(rotOp)) {
            seen.add(rotOp);
            q.push(rotOp);
        }
    }

    return ans;
};
```

---

### Python3

```python
from collections import deque

class Solution:
    def findLexSmallestString(self, s: str, a: int, b: int) -> str:
        n = len(s)
        seen = set([s])
        q = deque([s])
        ans = s

        while q:
            cur = q.popleft()
            if cur < ans:
                ans = cur

            # Operation 1: add a to odd indices
            arr = list(cur)
            for i in range(1, n, 2):
                arr[i] = str((int(arr[i]) + a) % 10)
            addOp = ''.join(arr)
            if addOp not in seen:
                seen.add(addOp)
                q.append(addOp)

            # Operation 2: rotate right by b
            rotOp = cur[-b:] + cur[:-b]
            if rotOp not in seen:
                seen.add(rotOp)
                q.append(rotOp)

        return ans
```

---

### Go

```go
package main

import (
 "container/list"
)

func findLexSmallestString(s string, a int, b int) string {
 n := len(s)
 seen := make(map[string]bool)
 q := list.New()
 seen[s] = true
 q.PushBack(s)
 ans := s

 for q.Len() > 0 {
  front := q.Remove(q.Front()).(string)
  if front < ans {
   ans = front
  }

  // add a to odd indices
  addBytes := []byte(front)
  for i := 1; i < n; i += 2 {
   d := int(addBytes[i]-'0')
   d = (d + a) % 10
   addBytes[i] = byte('0' + d)
  }
  addOp := string(addBytes)
  if !seen[addOp] {
   seen[addOp] = true
   q.PushBack(addOp)
  }

  // rotate right by b
  rotOp := front[n-b:] + front[:n-b]
  if !seen[rotOp] {
   seen[rotOp] = true
   q.PushBack(rotOp)
  }
 }

 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the logic and the important lines in small chunks (I’ll use the Python version as the main narrative, but the same maps directly to other languages).

### Initialization

```python
n = len(s)
seen = set([s])
q = deque([s])
ans = s
```

* `n` stores the length of `s`.
* `seen` keeps strings we've already processed to avoid cycles.
* `q` is a BFS queue initialized with the starting string.
* `ans` stores the best (smallest) string found so far.

Equivalent operations in other languages:

* C++: `unordered_set<string> seen; queue<string> q;`
* Java: `Set<String> seen = new HashSet<>(); Queue<String> q = new ArrayDeque<>();`
* JS: `const seen = new Set(); const q = [];`
* Go: `seen := make(map[string]bool); q := list.New()`

### BFS Loop

```python
while q:
    cur = q.popleft()
    if cur < ans:
        ans = cur
```

* Pop the next state (`cur`) to explore.
* Update `ans` if `cur` is lexicographically smaller.

### Operation 1 — Add `a` to odd indices

```python
arr = list(cur)
for i in range(1, n, 2):
    arr[i] = str((int(arr[i]) + a) % 10)
addOp = ''.join(arr)
if addOp not in seen:
    seen.add(addOp)
    q.append(addOp)
```

* Convert the string to a mutable list of characters.
* For every odd index (1,3,5,...), compute `(digit + a) % 10` and update.
* Create the new string `addOp`.
* If unseen, add to `seen` and queue.
* Same logic applies in other languages (C++ loop, Java char array, JS split/join, Go []byte).

### Operation 2 — Rotate right by `b`

```python
rotOp = cur[-b:] + cur[:-b]
if rotOp not in seen:
    seen.add(rotOp)
    q.append(rotOp)
```

* Right rotation by `b` moves the last `b` characters to the front.
* If unseen, add to `seen` and queue.

### Termination

* The BFS finishes when the queue is empty.
* `ans` has the lexicographically smallest reachable string.

---

## Examples

**Example 1**

```
Input: s = "5525", a = 9, b = 2
Output: "2050"
Explanation: One possible sequence:
"5525" -> rotate -> "2555" -> add -> "2454" -> add -> "2353" -> rotate -> "5323" -> add -> "5222" -> add -> "5121" -> rotate -> "2151" -> add -> "2050"
No lexicographically smaller string possible.
```

**Example 2**

```
Input: s = "74", a = 5, b = 1
Output: "24"
```

**Example 3**

```
Input: s = "0011", a = 4, b = 2
Output: "0011"
```

---

## How to use / Run locally

### C++

* Compile and run in a LeetCode-style environment or paste the `Solution` class in your local file and call the function from a `main()` wrapper for testing.
* Example (g++): `g++ -std=c++17 solution.cpp && ./a.out`

### Java

* Place `Solution` class in a file `Solution.java`. Use a `main` to call `findLexSmallestString`.
* Example: `javac Solution.java && java Solution`

### JavaScript

* Use Node.js. Put the function in a file (e.g., `solution.js`), export or call directly with test cases.
* Example: `node solution.js`

### Python3

* Put the `Solution` class in `solution.py`. Create an instance and call `findLexSmallestString`.
* Example: `python3 solution.py`

### Go

* Put the `findLexSmallestString` function in a file `main.go`. Add a `main()` that calls it and prints the result.
* Example: `go run main.go`

---

## Notes & Optimizations

* BFS / DFS both work. BFS provides clear structure and avoids deep recursion limits.
* `visited` prevents infinite loops and repeated processing.
* Two transformations per string — each costs `O(n)` time. Storage cost depends on number of reachable strings.
* For some inputs, you can reason about cycles analytically (e.g., repeated adding on odd indices has period ≤ 10), but a general BFS is simple and robust for constraints.
* Minor optimization: rotating by `b` repeated may create only `gcd(n, b)` distinct rotations. But BFS + visited naturally exploits this.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
