# 1518. Water Bottles

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

You start with `numBottles` full water bottles. After drinking a full bottle you get one empty bottle. You can exchange `numExchange` empty bottles for one full bottle from the market. Return the **maximum number of bottles** you can drink.

Example: If `numBottles = 9` and `numExchange = 3`, you can drink 13 bottles in total.

---

## Constraints

* `1 <= numBottles <= 100`
* `2 <= numExchange <= 100`

(Constraints are small here so both simulation and formula approaches are safe.)

---

## Intuition

I thought about the process like this: I drink all full bottles I have and collect empty bottles. When the number of empty bottles reaches `numExchange`, I trade them for full bottles, drink those, get more empties, and repeat. This repeats until I don’t have enough empties to exchange.

There is also a math shortcut: for `k = numExchange` and `k > 1`, the total is `numBottles + (numBottles - 1) / (k - 1)` (integer division). But I prefer to show the simple simulation first because it's easy to understand and maps directly to the real process.

---

## Approach

1. Set `total = numBottles` — the bottles I drink initially.
2. Set `empties = numBottles` — because I get an empty bottle for every bottle I drank.
3. While `empties >= numExchange`:

   * Exchange `empties / numExchange` empties for `newFull` bottles.
   * Drink those `newFull` bottles: `total += newFull`.
   * Update `empties = newFull + (empties % numExchange)`. (Leftover empties + empties from newly drunk bottles.)
4. When loop stops, `total` is the answer.

I prefer this simulation because it's simple, intuitive and O(1) extra space. The math formula gives O(1) time too and is good as an optimized variant.

---

## Data Structures Used

* Primitive integers only (`int` / `long` depending on language).
  No arrays, stacks, or other containers are needed.

---

## Operations & Behavior Summary

* **Drink**: reduces full count implicitly (we track drinks in `total`) and increases `empties`.
* **Exchange**: `empties / numExchange` empties → 1 full bottle each.
* Repeat drink & exchange until `empties < numExchange`.

---

## Complexity

**Simulation approach**

* **Time Complexity:** O(T) where T is number of exchange iterations (in practice ≤ total bottles drunk). With constraints T is small; this is effectively O(total_bottles_drunk).
* **Space Complexity:** O(1) — only a few integer variables.

**Formula approach**

* **Time Complexity:** O(1) — constant arithmetic.
* **Space Complexity:** O(1).

---

## Multi-language Solutions

> All implementations below use the simulation approach (clear and easy to follow). I also mention the O(1) formula variant in the Notes & Optimizations section.

---

### C++

```c++
/*
 * Simulation solution for LeetCode 1518. Water Bottles
 * Time: O(T) (practically small), Space: O(1)
 */
class Solution {
public:
    int numWaterBottles(int numBottles, int numExchange) {
        // total bottles I've drunk so far
        int total = numBottles;
        // number of empty bottles I currently have
        int empties = numBottles;

        // while I have enough empty bottles to exchange for at least one full bottle
        while (empties >= numExchange) {
            int newFull = empties / numExchange;           // how many new full bottles I get
            total += newFull;                             // drink them immediately
            empties = newFull + (empties % numExchange);  // new empties = newly drunk + leftover empties
        }

        return total;
    }
};
```

---

### Java

```java
// Simulation solution for LeetCode 1518. Water Bottles
// Time: O(T), Space: O(1)
class Solution {
    public int numWaterBottles(int numBottles, int numExchange) {
        int total = numBottles;  // bottles drunk initially
        int empties = numBottles; // empties after initial drinks

        while (empties >= numExchange) {
            int newFull = empties / numExchange;    // exchange empties for full bottles
            total += newFull;                       // drink them
            empties = newFull + (empties % numExchange); // update empties
        }

        return total;
    }
}
```

---

### JavaScript

```javascript
/**
 * Simulation solution for LeetCode 1518. Water Bottles
 * @param {number} numBottles
 * @param {number} numExchange
 * @return {number}
 */
var numWaterBottles = function(numBottles, numExchange) {
    let total = numBottles; // drunk so far
    let empties = numBottles; // empties after drinking initial bottles

    while (empties >= numExchange) {
        const newFull = Math.floor(empties / numExchange); // bottles obtained by exchange
        total += newFull;                                  // drink them
        empties = newFull + (empties % numExchange);       // update empties
    }

    return total;
};
```

---

### Python3

```python
# Simulation solution for LeetCode 1518. Water Bottles
# Time: O(T), Space: O(1)

class Solution:
    def numWaterBottles(self, numBottles: int, numExchange: int) -> int:
        total = numBottles       # bottles I've already drunk
        empties = numBottles     # empty bottles after drinking initial ones

        while empties >= numExchange:
            new_full = empties // numExchange   # full bottles I can get this round
            total += new_full                   # drink them
            empties = new_full + (empties % numExchange)  # update empties

        return total
```

---

### Go

```go
// Simulation solution for LeetCode 1518. Water Bottles
// Time: O(T), Space: O(1)
package main

func numWaterBottles(numBottles int, numExchange int) int {
    total := numBottles
    empties := numBottles

    for empties >= numExchange {
        newFull := empties / numExchange
        total += newFull
        empties = newFull + empties%numExchange
    }

    return total
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the algorithm step-by-step and map it to the code. The logic is identical across languages; variable names in code match the explanation.

### Core idea (simple):

1. Drink all bottles you have at start. That gives you `numBottles` empties.
2. While empties >= numExchange:

   * Exchange empties for new full bottles: `newFull = empties / numExchange`
   * Drink those new full bottles — add `newFull` to `total`.
   * After drinking them, those become empty bottles, and some empties may remain from the exchange step:
     `empties = newFull + (empties % numExchange)`
3. Repeat until you cannot exchange anymore (`empties < numExchange`).
4. `total` is the final answer.

### Line-by-line mapping (using the Python code as the reference):

```python
total = numBottles
empties = numBottles
```

* I start with `numBottles` full bottles, so I can drink `numBottles` immediately. After drinking them I have `numBottles` empties.

```python
while empties >= numExchange:
```

* I can only exchange if I have at least `numExchange` empty bottles.

```python
    new_full = empties // numExchange
```

* Integer division tells me how many full bottles I can get in this exchange step.

```python
    total += new_full
```

* I drink the new bottles immediately, so I add that to the running total.

```python
    empties = new_full + (empties % numExchange)
```

* After drinking `new_full` bottles, those become empty bottles themselves.
* `(empties % numExchange)` are the leftover empties that weren't used in the exchange.
* So the next-round empties is leftover empties plus the newly created empties.

```python
return total
```

* When loop finishes, no more exchanges are possible. `total` is the maximum number of bottles I could drink.

### Important mapping notes to other languages

* C++/Java: use `/` for integer division with ints. Same formulas apply.
* JavaScript: use `Math.floor(empties / numExchange)` for integer division.
* Go: use `/` and `%` with ints as shown.
* There is no need for arrays or extra storage; all state fits in a few integers.

---

## Examples

**Example 1**

```
Input: numBottles = 9, numExchange = 3
Output: 13
Explanation:
Drink 9 initially -> 9 empties
Exchange 9 empties for 3 full -> drink them (total 12), empties: 3
Exchange 3 empties for 1 full -> drink it (total 13), empties: 1
No more exchanges possible.
```

**Example 2**

```
Input: numBottles = 15, numExchange = 4
Output: 19
Explanation:
Drink 15 -> 15 empties
Exchange 15 // 4 = 3 -> drink them (total 18), empties = 3 + (15 % 4 = 3) = 6
Exchange 6 // 4 = 1 -> drink it (total 19), empties = 1 + (6 % 4 = 2) = 3
Stop.
```

---

## How to use / Run locally

Below are minimal instructions to compile/run examples locally. For each language create a file (examples shown), and if needed add a small `main()` / test harness.

### C++

Save as `solution.cpp` (add a simple `main()` to test):

```bash
g++ -std=c++17 solution.cpp -O2 -o solution
./solution
```

### Java

Save as `Solution.java` (add `public static void main(String[] args)` for a quick run):

```bash
javac Solution.java
java Solution
```

### JavaScript (Node.js)

Save as `solution.js`:

```bash
node solution.js
```

### Python3

Save as `solution.py`:

```bash
python3 solution.py
```

### Go

Save as `main.go` (use `package main` and a `main()` to test):

```bash
go run main.go
```

(For LeetCode, paste the class/function directly into the online editor — you do not need the above harnesses.)

---

## Notes & Optimizations

* **Mathematical optimization (O(1) formula):** There's a direct formula for this problem:

  ```
  total = numBottles + (numBottles - 1) // (numExchange - 1)
  ```

  This works when `numExchange >= 2`. It comes from observing the repeated exchange process and counting how many extra bottles each original bottle effectively produces, but be careful to use integer division properly. This formula gives the same answer as the simulation and runs in constant time.

* **Edge cases:** Given constraints (`numExchange >= 2`), the formula is safe. If someone supplied `numExchange <= 1` (invalid for the problem), the exchange process would be infinite — but the constraints prevent that.

* **Why simulation is fine here:** Constraints are small; simulation is simple, easy to understand, bug-resistant, and uses constant extra space.

---

## Author

[Md. Aarzoo Islam](https://bento.me/withaarzoo)
