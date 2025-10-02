# 3100. Water Bottles II — README

---

## Problem Title

**3100. Water Bottles II**

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
* [Step-by-step Detailed Explanation (per language)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

We are given two integers: `numBottles` (initial full water bottles I have) and `numExchange` (the number of empty bottles required to exchange for **one** full bottle).

Rules:

* I can drink any number of full bottles turning them into empty bottles.
* I may exchange `numExchange` empty bottles for **one** full bottle. **After each exchange, `numExchange` increases by 1** (so the next exchange requires more empties).
* I cannot perform multiple exchanges with the same `numExchange` value simultaneously — every time I exchange, `numExchange` increases.
* Return the **maximum** number of bottles I can drink.

(This is the variant where the exchange threshold increments after each exchange.)

---

## Constraints

* `1 <= numBottles <= 100`
* `1 <= numExchange <= 100`

---

## Intuition

I thought: every time I drink a full bottle, it becomes an empty bottle which may help for future exchanges. Because the exchange requirement increases after every exchange, I cannot batch multiple exchanges at the same threshold. So drinking everything I currently have first gives me the maximum empties to try to do one-by-one exchanges (increasing the threshold each time). Repeat until no more exchanges are possible and there are no full bottles left.

---

## Approach

1. Keep counters:

   * `full` = current full bottles I have (start = `numBottles`)
   * `empty` = current empty bottles I have (start = 0)
   * `ans` = total bottles drunk (start = 0)
   * `curEx` = current exchange requirement (start = `numExchange`)
2. Loop while `full > 0`:

   * Drink all `full` bottles: `ans += full`, `empty += full`, `full = 0`.
   * Now attempt exchanges **one at a time**: while `empty >= curEx`:

     * `empty -= curEx`
     * `full += 1`
     * `curEx += 1` (increase threshold immediately after each exchange)
3. Repeat until no full bottle remains and no further exchanges possible.
4. `ans` is the final answer.

This greedy simulation follows the rules directly and is easy to reason about.

---

## Data Structures Used

* Primitive integer counters only (`full`, `empty`, `ans`, `curEx`).
* No arrays, lists, sets, or other containers are needed.

---

## Operations & Behavior Summary

* **Drink operation**: convert all `full` to `empty` and increase `ans`.
* **Exchange operation**: while `empty >= curEx`, exchange `curEx` empties for **one** full bottle and increment `curEx` by 1.
* Repeat drinking + exchanging until no more actions can be performed.

---

## Complexity

* **Time Complexity:** `O(A)` where `A` = total number of operations that produce a bottle drunk or an exchange. More intuitively, `O(final_bottles_drunk)` which is ≤ some small constant given the problem constraints (practically `O(numBottles + number_of_exchanges)`). With constraints ≤100 this is extremely fast.
* **Space Complexity:** `O(1)` — only a fixed number of integer variables are used.

---

## Multi-language Solutions

### C++

```c++
/*
 * Solution (C++) - Greedy simulation
 * Function signature matches the requested style.
 */
class Solution {
public:
    int maxBottlesDrunk(int numBottles, int numExchange) {
        int full = numBottles;    // current full bottles
        int empty = 0;            // current empty bottles
        int ans = 0;              // total drunk
        int curEx = numExchange;  // current exchange requirement (increases after every exchange)

        while (full > 0) {
            // Drink all full bottles now
            ans += full;
            empty += full;
            full = 0;

            // Exchange empties for one full bottle at a time,
            // because curEx increases after each exchange.
            while (empty >= curEx) {
                empty -= curEx;  // spend empties
                full += 1;       // get one full bottle to drink later
                curEx += 1;      // exchange threshold increases
            }
        }
        return ans;
    }
};
```

### Java

```java
// Solution (Java) - Greedy simulation
class Solution {
    public int maxBottlesDrunk(int numBottles, int numExchange) {
        int full = numBottles;
        int empty = 0;
        int ans = 0;
        int curEx = numExchange;

        while (full > 0) {
            // Drink all available full bottles
            ans += full;
            empty += full;
            full = 0;

            // Exchange one-by-one while possible
            while (empty >= curEx) {
                empty -= curEx; // use empties
                full += 1;      // receive one full bottle
                curEx += 1;     // increase exchange requirement
            }
        }
        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * Solution (JavaScript) - Greedy simulation
 * @param {number} numBottles
 * @param {number} numExchange
 * @return {number}
 */
var maxBottlesDrunk = function(numBottles, numExchange) {
    let full = numBottles;
    let empty = 0;
    let ans = 0;
    let curEx = numExchange;

    while (full > 0) {
        // Drink all full bottles
        ans += full;
        empty += full;
        full = 0;

        // Exchange empties for one full bottle at a time
        while (empty >= curEx) {
            empty -= curEx;
            full += 1;
            curEx += 1;
        }
    }
    return ans;
};
```

### Python3

```python
# Solution (Python3) - Greedy simulation
class Solution:
    def maxBottlesDrunk(self, numBottles: int, numExchange: int) -> int:
        full = numBottles
        empty = 0
        ans = 0
        curEx = numExchange

        while full > 0:
            # Drink all the full bottles
            ans += full
            empty += full
            full = 0

            # Exchange empties one-by-one
            while empty >= curEx:
                empty -= curEx
                full += 1
                curEx += 1

        return ans
```

### Go

```go
// Solution (Go) - Greedy simulation
package main

func maxBottlesDrunk(numBottles int, numExchange int) int {
    full := numBottles
    empty := 0
    ans := 0
    curEx := numExchange

    for full > 0 {
        // Drink all full bottles
        ans += full
        empty += full
        full = 0

        // Exchange one-by-one while possible
        for empty >= curEx {
            empty -= curEx
            full += 1
            curEx += 1
        }
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Below I explain the core logic once — it is identical in all languages. For clarity I reference the Python version lines (the same operations occur in other languages):

1. **Initialize variables**

   ```python
   full = numBottles
   empty = 0
   ans = 0
   curEx = numExchange
   ```

   * `full`: how many full bottles I currently have to drink.
   * `empty`: how many empty bottles I currently have (will be used for exchanges).
   * `ans`: total bottles I have drunk so far.
   * `curEx`: current exchange threshold (increases after each change).

2. **Main loop** — `while full > 0`:

   * As long as I have any full bottles, I should drink them immediately because drinking creates empties that may allow exchanges.

3. **Drink all full bottles**

   ```python
   ans += full
   empty += full
   full = 0
   ```

   * I drink all `full` bottles and convert them into empties. This increases `ans` and `empty`.

4. **Exchange loop (one-by-one)**

   ```python
   while empty >= curEx:
       empty -= curEx
       full += 1
       curEx += 1
   ```

   * While I have at least `curEx` empties, I perform exactly **one** exchange at a time:

     * Spend `curEx` empties to gain one full bottle.
     * Immediately increase `curEx` (next exchange will need more empties).
   * This ensures I respect the rule that `numExchange` increments after every exchange and prevents batching multiple exchanges at the same threshold.

5. **Repeat**

   * After exchange(s), any new `full` bottles will be drunk in the next outer loop iteration. Continue until no more actions are possible.

6. **Return answer**

   * The variable `ans` holds the total number of bottles drunk.

**Important subtlety:** Because the exchange threshold grows after **each** exchange, we cannot simply do `full += empty / curEx` in one step — we must simulate exchanges one-by-one (or carefully simulate the threshold changes). The one-by-one simulation is simple, correct, and fast for given constraints.

---

## Examples

**Example 1**

* Input: `numBottles = 13`, `numExchange = 6`
* Execution trace (compressed):

  1. Drink 13 → `ans=13`, `empty=13`.
  2. Exchange 6 → get 1 full → `empty=7`, `curEx=7`, `full=1`.
  3. Exchange 7 → get 1 full → `empty=0`, `curEx=8`, `full=2`.
  4. Drink 2 → `ans=15`, `empty=2`.
  5. `empty(2) < curEx(8)` → stop.
* Output: `15`

**Example 2**

* Input: `numBottles = 10`, `numExchange = 3`
* Execution trace (compressed):

  1. Drink 10 → `ans=10`, `empty=10`.
  2. Exchange 3 → get 1 full → `empty=7`, `curEx=4`.
  3. Exchange 4 → get 1 full → `empty=3`, `curEx=5`.
  4. Drink 2 → `ans=12`, `empty=5`.
  5. Exchange 5 → get 1 full → `empty=0`, `curEx=6`.
  6. Drink 1 → `ans=13`.
  7. Stop.
* Output: `13`

---

## How to use / Run locally

**C++**

1. Create `Solution.cpp` containing the `Solution` class above and add a `main` to call it:

```cpp
#include <iostream>
int main() {
    Solution s;
    std::cout << s.maxBottlesDrunk(13, 6) << std::endl; // prints 15
    return 0;
}
```

2. Compile and run:

```bash
g++ -std=c++17 Solution.cpp -O2 -o solution
./solution
```

**Java**

1. Create `Solution.java` with the `Solution` class and a `main`:

```java
public class Solution {
    // include method maxBottlesDrunk here...
    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.maxBottlesDrunk(13, 6)); // prints 15
    }
}
```

2. Compile and run:

```bash
javac Solution.java
java Solution
```

**JavaScript (Node.js)**

1. Create `solution.js` with the function and test:

```javascript
function maxBottlesDrunk(numBottles, numExchange) {
    // include function body
}
console.log(maxBottlesDrunk(13, 6)); // prints 15
```

2. Run:

```bash
node solution.js
```

**Python3**

1. Create `solution.py`:

```python
class Solution:
    def maxBottlesDrunk(self, numBottles, numExchange):
        # function body

if __name__ == "__main__":
    s = Solution()
    print(s.maxBottlesDrunk(13, 6))  # prints 15
```

2. Run:

```bash
python3 solution.py
```

**Go**

1. Create `solution.go` with the function and `main`:

```go
package main
import "fmt"
func main() {
    fmt.Println(maxBottlesDrunk(13, 6)) // prints 15
}
```

2. Run:

```bash
go run solution.go
```

---

## Notes & Optimizations

* Because `curEx` changes after every exchange, we must simulate exchanges one-by-one (or emulate the same behavior using math). The one-by-one simulation is easy and very fast for the given constraints (`<= 100`).
* The algorithm is `O(final_bottles_drunk)` in time and `O(1)` in space.
* Edge cases:

  * If `numExchange` is very large compared to `numBottles`, we may never be able to exchange — algorithm handles it naturally.
  * If `numExchange == 1`, the algorithm will exchange greedily and still behave correctly (after every exchange `curEx` increments, preventing infinite loops).
* For theoretical large inputs (not in current constraints), we would still simulate until no exchange possible; if performance is a concern for much larger bounds, a math-based approach can be derived but must carefully handle the incrementing threshold.

---

## Author

**[Md. Aarzoo Islam](https://bento.me/withaarzoo)**
