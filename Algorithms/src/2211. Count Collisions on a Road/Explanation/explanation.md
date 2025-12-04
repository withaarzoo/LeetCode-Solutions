# 2211. Count Collisions on a Road

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

I am given a string `directions` of length `n`.
Each character represents a car on an infinite road:

* `'L'` → car is moving left
* `'R'` → car is moving right
* `'S'` → car is stationary (already stopped)

All cars move with the **same speed**.

Collisions happen with these rules:

1. When a left-moving car and a right-moving car collide (opposite directions), the number of collisions increases by **2**, and both cars become stationary.
2. When a moving car hits a stationary car, the number of collisions increases by **1**, and the moving car becomes stationary.
3. After a collision, cars involved **do not move anymore**.

I need to return the **total number of collisions** that will happen on the road.

---

## Constraints

* `1 ≤ directions.length ≤ 10^5`
* `directions[i]` is one of: `'L'`, `'R'`, `'S'`

So I must design an `O(n)` or close-to-linear solution. Anything slower (like simulating step-by-step movement) will be too slow.

---

## Intuition

When I first saw the problem, my brain started simulating cars:

* Car going right `R`, another car coming from right side as `L`, or a stationary `S` in between…
* I tried to imagine them colliding step by step.

But then I realized this would be messy and inefficient.

So I asked myself:

> “Which cars can **never** collide?”

I quickly noticed:

* Any car on the **far left** that is going `'L'` will just run away to the left. There is no car to its left → no collision.
* Any car on the **far right** that is going `'R'` will just run away to the right. There is no car to its right → no collision.

So I can safely **ignore**:

* all **leading `'L'`** from the beginning of the string
* all **trailing `'R'`** from the end of the string

Then I focused only on the **middle part** after trimming.

Inside this middle part, I realized something beautiful:

> Every car that is still moving (`'L'` or `'R'`) is guaranteed to collide **exactly once** and become stationary.

So in this middle segment:

* Every `'L'` or `'R'` will contribute **1 collision**.
* `'S'` cars don’t add new collisions (they are already stopped).

So the answer becomes simply:

> **Count of non-`'S'` cars (i.e., `'L'` or `'R'`) in the middle substring after trimming the outer safe cars.**

---

## Approach

1. Let `directions` be the input string, and `n` be its length.

2. I use two pointers:

   * `i` starting from the **left** (index 0),
   * `j` starting from the **right** (index `n-1`).

3. Move `i` rightwards while `directions[i] == 'L'`.

   * These are cars at the far left moving left, so they never collide.

4. Move `j` leftwards while `directions[j] == 'R'`.

   * These are cars at the far right moving right, so they also never collide.

5. Now I focus on the substring from `i` to `j` (inclusive).
   This is the **active collision zone**.

6. I scan from `k = i` to `k = j`:

   * For each character:

     * If it is `'L'` or `'R'`, I add 1 to `collisions`.
     * If it is `'S'`, I do nothing.

7. Return the total `collisions`.

Why is this correct?

* Inside this middle zone:

  * A right-moving car `R` will eventually hit something to its right (`L` or `S`).
  * A left-moving car `L` will eventually hit something to its left (`R` or `S`).
  * After colliding once, they stop and never move again.
* So each moving car (`L` or `R`) will cause **one** collision.
* Stationary cars do not cause new collisions; they just get hit.

---

## Data Structures Used

* I only use a few **integer variables** and **indexes**:

  * `i` and `j` (two pointers)
  * a loop index `k`
  * an integer `collisions` to store the count

No extra arrays, no stacks, no queues.

---

## Operations & Behavior Summary

* **Trimming phase**:

  * Skip continuous `'L'` from the start.
  * Skip continuous `'R'` from the end.

* **Counting phase**:

  * For the remaining middle part, count every `'L'` and `'R'`.

Each character is processed at most a constant number of times.

---

## Complexity

* **Time Complexity:** `O(n)`

  * `n` is the length of the string `directions`.
  * I do one pass from left to right (for leading `L`), one pass from right to left (for trailing `R`), and one pass over the remaining subarray. Overall, still linear.

* **Space Complexity:** `O(1)`

  * I use only a few integer variables.
  * No extra data structures depend on `n`.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int countCollisions(string directions) {
        int n = directions.size();
        int i = 0, j = n - 1;
        
        // Skip all leading 'L' cars (they move left forever, no collision)
        while (i < n && directions[i] == 'L') {
            i++;
        }
        
        // Skip all trailing 'R' cars (they move right forever, no collision)
        while (j >= 0 && directions[j] == 'R') {
            j--;
        }
        
        int collisions = 0;
        // In the middle part, every 'L' or 'R' will collide exactly once
        for (int k = i; k <= j; k++) {
            if (directions[k] != 'S') {
                collisions++;
            }
        }
        
        return collisions;
    }
};
```

---

### Java

```java
class Solution {
    public int countCollisions(String directions) {
        int n = directions.length();
        int i = 0, j = n - 1;
        
        // Skip all leading 'L' cars (safe, no collision)
        while (i < n && directions.charAt(i) == 'L') {
            i++;
        }
        
        // Skip all trailing 'R' cars (safe, no collision)
        while (j >= 0 && directions.charAt(j) == 'R') {
            j--;
        }
        
        int collisions = 0;
        // Count all moving cars ('L' or 'R') in the middle part
        for (int k = i; k <= j; k++) {
            if (directions.charAt(k) != 'S') {
                collisions++;
            }
        }
        
        return collisions;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {string} directions
 * @return {number}
 */
var countCollisions = function(directions) {
    const n = directions.length;
    let i = 0, j = n - 1;
    
    // Skip leading 'L' cars - they never collide
    while (i < n && directions[i] === 'L') {
        i++;
    }
    
    // Skip trailing 'R' cars - they never collide
    while (j >= 0 && directions[j] === 'R') {
        j--;
    }
    
    let collisions = 0;
    // Every non-'S' in the middle will collide exactly once
    for (let k = i; k <= j; k++) {
        if (directions[k] !== 'S') {
            collisions++;
        }
    }
    
    return collisions;
};
```

---

### Python3

```python
class Solution:
    def countCollisions(self, directions: str) -> int:
        n = len(directions)
        i, j = 0, n - 1
        
        # Skip leading 'L' cars (safe)
        while i < n and directions[i] == 'L':
            i += 1
        
        # Skip trailing 'R' cars (safe)
        while j >= 0 and directions[j] == 'R':
            j -= 1
        
        collisions = 0
        # Count all moving cars ('L' or 'R') in the remaining middle part
        for k in range(i, j + 1):
            if directions[k] != 'S':
                collisions += 1
        
        return collisions
```

---

### Go

```go
package main

func countCollisions(directions string) int {
 n := len(directions)
 i, j := 0, n-1

 // Skip leading 'L' cars (no collision)
 for i < n && directions[i] == 'L' {
  i++
 }

 // Skip trailing 'R' cars (no collision)
 for j >= 0 && directions[j] == 'R' {
  j--
 }

 collisions := 0
 // Every non-'S' in this middle segment will collide exactly once
 for k := i; k <= j; k++ {
  if directions[k] != 'S' {
   collisions++
  }
 }

 return collisions
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Here I’ll explain the logic using the C++ version, but the same idea is used in all languages.

### 1. Define variables

```cpp
int n = directions.size();
int i = 0, j = n - 1;
```

* `n` = length of the input string.
* `i` starts from the **left**, `j` from the **right**.

---

### 2. Skip leading `'L'` cars

```cpp
while (i < n && directions[i] == 'L') {
    i++;
}
```

* While I am inside the string and the current car is `'L'`, I move `i` to the right.
* These cars are on the extreme left, going further left.
  No one is to their left, so they never collide.

All implementations do the same with slightly different syntax.

---

### 3. Skip trailing `'R'` cars

```cpp
while (j >= 0 && directions[j] == 'R') {
    j--;
}
```

* While I am inside the string and the current car at `j` is `'R'`, I move `j` to the left.
* These cars are on the extreme right, going further right.
  No one is to their right, so they never collide.

Again, same idea in Java, JS, Python, Go.

---

### 4. Initialize collision counter

```cpp
int collisions = 0;
```

I will store the total number of collisions here.

---

### 5. Count non-`'S'` cars in the middle segment

```cpp
for (int k = i; k <= j; k++) {
    if (directions[k] != 'S') {
        collisions++;
    }
}
```

* I loop from `i` to `j` (inclusive). This is the “collision zone”.
* If the car at position `k` is:

  * `'R'` → it will eventually collide once.
  * `'L'` → it will eventually collide once.
  * `'S'` → already stopped, so no new collision directly from this car.

So for every character that is **not `'S'`**, I increase `collisions` by 1.

This is exactly the same logic in:

* Java: using `charAt(k)` and `!= 'S'`
* JavaScript: `directions[k] !== 'S'`
* Python: `directions[k] != 'S'`
* Go: `directions[k] != 'S'`

---

### 6. Return the answer

```cpp
return collisions;
```

Finally, I return the total number of collisions.

---

## Examples

### Example 1

**Input:**

```text
directions = "RLRSLL"
```

**Process idea (conceptually):**

* No leading `'L'` at the start.
* No trailing `'R'` at the end.
* Middle part is the whole string.

Cars:

* `R` (0) and `L` (1) collide → 2 collisions, both stop.
* Stationary or stopped cars get hit by further cars in sequence.

**Output:**

```text
5
```

---

### Example 2

**Input:**

```text
directions = "LLRR"
```

* Leading `'L'`s: `"LL"` → skipped.
* Trailing `'R'`s: `"RR"` → skipped.
* Middle part is empty.

No car ever collides.

**Output:**

```text
0
```

---

## How to use / Run locally

### C++

1. Save the C++ solution in a file, e.g. `solution.cpp`.
2. Add a `main` function or use the LeetCode environment. Example for local test:

```cpp
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int countCollisions(string directions) {
        int n = directions.size();
        int i = 0, j = n - 1;
        while (i < n && directions[i] == 'L') i++;
        while (j >= 0 && directions[j] == 'R') j--;
        int collisions = 0;
        for (int k = i; k <= j; k++) {
            if (directions[k] != 'S') collisions++;
        }
        return collisions;
    }
};

int main() {
    Solution sol;
    string s = "RLRSLL";
    cout << sol.countCollisions(s) << endl;
    return 0;
}
```

3. Compile & run:

```bash
g++ -std=c++17 solution.cpp -o solution
./solution
```

---

### Java

1. Save as `Solution.java`.

```bash
javac Solution.java
java Solution
```

(LeetCode will provide the `main` method environment automatically.)

---

### JavaScript (Node.js)

1. Save as `solution.js`.
2. To test locally:

```bash
node solution.js
```

(You can write a small driver at the bottom to call `countCollisions`.)

---

### Python3

1. Save as `solution.py`.
2. Example driver:

```python
if __name__ == "__main__":
    sol = Solution()
    print(sol.countCollisions("RLRSLL"))
```

3. Run:

```bash
python3 solution.py
```

---

### Go

1. Save as `main.go`. Ensure you include `package main` and a `main()` function.

```bash
go run main.go
```

---

## Notes & Optimizations

* I **do not** simulate every time step of movement.
  Simulation would be `O(n * steps)` and not needed.
* I use a pure **counting + trimming** approach:

  * Trim safe edges (`L...` on left, `...R` on right).
  * Count moving cars (`L` or `R`) in the middle.
* This solution is:

  * Simple to implement
  * Easy to reason about
  * Optimal in both time (`O(n)`) and space (`O(1)`)

If later I want to extend this, I can still simulate or visualize collisions, but for counting only, this counting trick is enough.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
