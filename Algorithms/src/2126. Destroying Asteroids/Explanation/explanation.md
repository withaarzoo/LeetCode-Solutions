# 2126. Destroying Asteroids

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

LeetCode 2126 - Destroying Asteroids is a greedy algorithm problem where we are given the initial mass of a planet and an array representing asteroid masses.

The planet can collide with asteroids in any order.

If the planet's mass is greater than or equal to an asteroid's mass, the asteroid gets destroyed and the planet gains that asteroid's mass.

If the asteroid is larger than the current planet mass, the planet gets destroyed immediately.

The goal is to determine whether it is possible to destroy all asteroids by choosing the best collision order.

### Input

* An integer `mass`
* An integer array `asteroids`

### Output

* `true` if all asteroids can be destroyed
* `false` otherwise

This problem is a classic greedy sorting problem often asked in coding interviews and competitive programming contests.

---

## Constraints

| Constraint                      | Value               |
| ------------------------------- | ------------------- |
| `1 <= mass <= 10^5`             | Initial planet mass |
| `1 <= asteroids.length <= 10^5` | Number of asteroids |
| `1 <= asteroids[i] <= 10^5`     | Asteroid mass       |

---

## Intuition

My first observation was that the order of collisions is completely flexible.

Since destroying an asteroid increases the planet's mass, it makes sense to destroy the smallest asteroids first. Smaller asteroids are easier to consume and help the planet grow faster.

Once the planet becomes larger, handling bigger asteroids becomes easier.

This naturally suggests sorting the asteroid masses and processing them from smallest to largest.

If I cannot destroy the smallest remaining asteroid, then I definitely cannot destroy any larger asteroid either.

That observation leads directly to a greedy solution.

---

## Approach

1. Sort the asteroid array in ascending order.
2. Store the current planet mass using a larger integer type because the mass can grow significantly.
3. Traverse the sorted array.
4. For every asteroid:

   * Check whether the current planet mass is at least the asteroid mass.
   * If not, return `false`.
   * Otherwise destroy the asteroid and add its mass to the planet.
5. If all asteroids are processed successfully, return `true`.

This approach guarantees that the planet grows as early as possible and always has the best chance of surviving future collisions.

---

## Data Structures Used

### Array

The asteroid masses are stored in an array.

### Sorting

The array is sorted so that smaller asteroids are processed before larger ones.

### Integer Variable

A larger integer type (`long long`, `long`, or `int64`) is used to safely store the growing planet mass.

No additional data structures such as stacks, queues, heaps, or hash maps are required.

---

## Operations & Behavior Summary

The algorithm performs the following operations:

1. Receive the initial planet mass.
2. Sort all asteroid masses.
3. Start with the current planet mass.
4. Visit each asteroid in sorted order.
5. Compare the planet mass with the asteroid mass.
6. If the planet is smaller:

   * Return `false`.
7. Otherwise:

   * Destroy the asteroid.
   * Increase the planet mass.
8. Continue until every asteroid is processed.
9. Return `true` if all asteroids are destroyed.

---

## Complexity

| Metric           | Complexity   | Explanation                         |
| ---------------- | ------------ | ----------------------------------- |
| Time Complexity  | `O(n log n)` | Sorting dominates the runtime       |
| Space Complexity | `O(1)`       | Only a few extra variables are used |

Where:

* `n` = number of asteroids

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool asteroidsDestroyed(int mass, vector<int>& asteroids) {
        // Sort asteroids from smallest to largest
        sort(asteroids.begin(), asteroids.end());

        // Use long long because mass can become very large
        long long currentMass = mass;

        // Try destroying asteroids one by one
        for (int asteroid : asteroids) {

            // If planet is too small, it gets destroyed
            if (currentMass < asteroid) {
                return false;
            }

            // Gain the asteroid's mass
            currentMass += asteroid;
        }

        // All asteroids were destroyed
        return true;
    }
};
```

### Java

```java
class Solution {
    public boolean asteroidsDestroyed(int mass, int[] asteroids) {

        // Sort asteroids from smallest to largest
        Arrays.sort(asteroids);

        // Use long because mass can exceed int range
        long currentMass = mass;

        // Process each asteroid
        for (int asteroid : asteroids) {

            // Planet cannot destroy this asteroid
            if (currentMass < asteroid) {
                return false;
            }

            // Gain asteroid mass
            currentMass += asteroid;
        }

        // Successfully destroyed all asteroids
        return true;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} mass
 * @param {number[]} asteroids
 * @return {boolean}
 */
var asteroidsDestroyed = function(mass, asteroids) {

    // Sort asteroids in ascending order
    asteroids.sort((a, b) => a - b);

    // Store current planet mass
    let currentMass = mass;

    // Process every asteroid
    for (const asteroid of asteroids) {

        // Cannot destroy this asteroid
        if (currentMass < asteroid) {
            return false;
        }

        // Gain its mass
        currentMass += asteroid;
    }

    // All asteroids destroyed
    return true;
};
```

### Python3

```python
class Solution:
    def asteroidsDestroyed(self, mass: int, asteroids: List[int]) -> bool:

        # Sort asteroids from smallest to largest
        asteroids.sort()

        # Current planet mass
        current_mass = mass

        # Try destroying each asteroid
        for asteroid in asteroids:

            # Planet is too small
            if current_mass < asteroid:
                return False

            # Gain asteroid mass
            current_mass += asteroid

        # All asteroids destroyed
        return True
```

### Go

```go
func asteroidsDestroyed(mass int, asteroids []int) bool {

    // Sort asteroids from smallest to largest
    sort.Ints(asteroids)

    // Use int64 because mass can grow significantly
    currentMass := int64(mass)

    // Process each asteroid
    for _, asteroid := range asteroids {

        // Planet cannot destroy this asteroid
        if currentMass < int64(asteroid) {
            return false
        }

        // Gain asteroid mass
        currentMass += int64(asteroid)
    }

    // All asteroids destroyed
    return true
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same across all five languages.

### Step 1: Sort the Asteroids

The first thing the algorithm does is sort the asteroid masses.

Why?

Because smaller asteroids are easier to destroy.

Destroying them early increases the planet's mass and makes future collisions easier.

### Step 2: Store Current Mass

The current planet mass is stored in a variable.

As asteroids are destroyed, this value keeps increasing.

A larger integer type is used because the accumulated mass may become much larger than the original input values.

### Step 3: Process Each Asteroid

The algorithm iterates through the sorted asteroid list.

For every asteroid:

* Compare planet mass and asteroid mass.
* If the planet is too small, the process stops immediately.

At this point there is no valid way to continue because every remaining asteroid is at least as large as the current one.

### Step 4: Gain Mass

If the planet successfully destroys an asteroid:

```text
new_mass = current_mass + asteroid_mass
```

This growth is the key reason the greedy strategy works.

Each successful collision makes future collisions easier.

### Step 5: Finish the Traversal

If every asteroid is processed successfully, the planet survives all collisions.

The answer becomes:

```text
true
```

Otherwise:

```text
false
```

---

## Examples

### Example 1

**Input**

```text
mass = 10
asteroids = [3,9,19,5,21]
```

**Output**

```text
true
```

**Trace**

```text
Sort -> [3,5,9,19,21]

10 -> 13
13 -> 18
18 -> 27
27 -> 46
46 -> 67
```

All asteroids are destroyed.

---

### Example 2

**Input**

```text
mass = 5
asteroids = [4,9,23,4]
```

**Output**

```text
false
```

**Trace**

```text
Sort -> [4,4,9,23]

5 -> 9
9 -> 13
13 -> 22
```

The next asteroid is `23`.

Since:

```text
22 < 23
```

the planet gets destroyed.

---

### Example 3

**Input**

```text
mass = 1
asteroids = [1]
```

**Output**

```text
true
```

**Trace**

```text
1 >= 1
1 + 1 = 2
```

The only asteroid is destroyed successfully.

---

## How to Use / Run Locally

### C++

Compile:

```bash
g++ solution.cpp -o solution
```

Run:

```bash
./solution
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

or

```bash
python3 solution.py
```

---

### Go

Run:

```bash
go run solution.go
```

Build:

```bash
go build solution.go
```

---

## Notes & Optimizations

* Sorting is the key idea behind this solution.
* Trying all possible collision orders would be extremely expensive.
* The greedy approach avoids unnecessary work.
* Always use a larger integer type because the planet mass grows after every successful collision.
* If the current smallest asteroid cannot be destroyed, no other ordering can help.
* This makes the greedy strategy both correct and optimal.
* The solution is efficient enough for the maximum constraint size of `100000` asteroids.

### Alternative Idea

A brute-force approach could try different collision orders, but it would quickly become impossible due to the enormous number of permutations.

Sorting provides the optimal solution with excellent performance.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
