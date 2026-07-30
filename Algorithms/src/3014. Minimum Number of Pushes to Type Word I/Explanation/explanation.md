# 3014. Minimum Number of Pushes to Type Word I

A simple greedy solution for LeetCode 3014 that finds the minimum number of key presses required to type a word after optimally remapping a telephone keypad.

---

## Table of Contents

- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
  - [C++](#c)
  - [Java](#java)
  - [JavaScript](#javascript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

In this LeetCode greedy problem, I am given a string containing only distinct lowercase English letters.

A telephone keypad has 8 usable keys (2 through 9). I can remap the letters onto these keys in any way I want. The goal is to arrange the letters so that typing the given word requires the fewest total key presses.

Each key press depends on the letter's position on its assigned key:

- First letter on a key requires 1 push.
- Second letter requires 2 pushes.
- Third letter requires 3 pushes.
- And so on.

My task is to return the minimum total number of pushes needed to type the entire word.

This problem is a great example of a Greedy Algorithm because I always want to place letters in the cheapest available positions first.

---

## Constraints

| Constraint | Value |
| ------------ | ------- |
| Word Length | 1 ≤ word.length ≤ 26 |
| Characters | Lowercase English letters |
| Letter Uniqueness | Every letter is distinct |

---

## Intuition

The first thing I noticed was that every letter appears exactly once.

Since every character has the same importance, I do not need to worry about letter frequency. I only need to assign letters to the cheapest possible positions on the keypad.

There are only eight usable keys.

That means:

- The first 8 letters can all cost 1 push.
- The next 8 letters must cost 2 pushes.
- The next 8 letters cost 3 pushes.
- Any remaining letters cost 4 pushes.

Once I realized this pattern, the solution became much simpler. The answer depends only on how many letters exist, not which letters they are.

---

## Approach

I solve the problem using a simple greedy strategy.

1. Find the length of the word.
2. Visit every character one by one.
3. Every group of eight letters shares the same typing cost.
4. Calculate the cost using:

   `index / 8 + 1`

5. Add the cost to the final answer.
6. Return the total number of pushes.

Since every letter is distinct, the actual character values never affect the answer.

---

## Data Structures Used

| Data Structure | Purpose |
| --------------- | --------- |
| Integer Variable | Stores the total number of pushes |
| Loop Counter | Tracks the current letter index |

No arrays, hash maps, stacks, queues, or other data structures are required for this solution.

---

## Operations & Behavior Summary

The algorithm performs the following steps:

1. Start the answer at zero.
2. Traverse every letter in the word.
3. Determine which group of eight the current letter belongs to.
4. Compute its typing cost.
5. Add that cost to the total.
6. Continue until every letter has been processed.
7. Return the minimum total number of pushes.

Since every letter is unique, this greedy assignment is always optimal.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | O(n) | I visit each character exactly once. |
| Space Complexity | O(1) | Only a few integer variables are used. No extra memory grows with input size. |

Where:

- **n** = number of characters in the given word.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minimumPushes(string word) {
        // Store the final minimum number of pushes
        int pushes = 0;

        // Visit every distinct letter
        for (int i = 0; i < word.size(); i++) {

            // Every group of 8 letters increases the push count by 1
            // i = 0..7   -> cost = 1
            // i = 8..15  -> cost = 2
            // i = 16..23 -> cost = 3
            // i = 24..25 -> cost = 4
            pushes += (i / 8) + 1;
        }

        // Return the minimum total pushes
        return pushes;
    }
};
```

### Java

```java
class Solution {
    public int minimumPushes(String word) {

        // Store the final answer
        int pushes = 0;

        // Traverse every character
        for (int i = 0; i < word.length(); i++) {

            // Every block of 8 letters has the same cost
            pushes += (i / 8) + 1;
        }

        // Return the minimum number of pushes
        return pushes;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} word
 * @return {number}
 */
var minimumPushes = function(word) {

    // Store the answer
    let pushes = 0;

    // Visit every character
    for (let i = 0; i < word.length; i++) {

        // Cost increases after every 8 letters
        pushes += Math.floor(i / 8) + 1;
    }

    // Return the minimum pushes
    return pushes;
};
```

### Python3

```python
class Solution:
    def minimumPushes(self, word: str) -> int:

        # Store the final answer
        pushes = 0

        # Visit every character
        for i in range(len(word)):

            # Every group of 8 letters has the same typing cost
            pushes += (i // 8) + 1

        # Return the minimum number of pushes
        return pushes
```

### Go

```go
func minimumPushes(word string) int {

    // Store the final answer
    pushes := 0

    // Traverse every character
    for i := 0; i < len(word); i++ {

        // Every block of 8 letters increases the cost by 1
        pushes += (i / 8) + 1
    }

    // Return the minimum number of pushes
    return pushes
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical in all five programming languages. Only the syntax changes.

First, I create a variable to store the total number of pushes.

Next, I iterate through every character in the word.

Interestingly, I never need the actual character itself. The only thing that matters is its position in the traversal.

For every letter, I calculate its typing cost.

The first eight letters receive a cost of one.

The next eight letters receive a cost of two.

The next eight letters receive a cost of three.

The remaining letters receive a cost of four.

This happens naturally because integer division groups every eight consecutive indices together.

After calculating the cost, I immediately add it to the answer.

Once every character has been processed, I return the accumulated total.

Since all letters appear exactly once, this greedy strategy always produces the minimum possible number of pushes.

The implementation remains identical in C++, Java, JavaScript, Python3, and Go. Only language syntax differs.

---

## Examples

### Example 1

**Input**

```text
word = "abcde"
```

**Output**

```text
5
```

**Explanation**

There are only five letters.

Each one can be placed in the first position of a different key.

Cost:

- a → 1
- b → 1
- c → 1
- d → 1
- e → 1

Total = 5

---

### Example 2

**Input**

```text
word = "xycdefghij"
```

**Output**

```text
12
```

**Explanation**

There are ten letters.

- First 8 letters cost 1 each.
- Remaining 2 letters cost 2 each.

Total:

8 × 1 + 2 × 2 = 12

---

### Example 3

**Input**

```text
word = "abcdefghijklmnop"
```

**Output**

```text
24
```

**Explanation**

There are sixteen letters.

- First 8 letters → cost 1 each
- Next 8 letters → cost 2 each

Total:

8 × 1 + 8 × 2 = 24

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

Run with Node.js:

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

---

## Notes & Optimizations

- Since every letter is distinct, I never need to count frequencies.
- The actual letter values do not matter.
- No sorting is required.
- No additional memory is needed.
- The solution works in linear time.
- This is the optimal solution because every cheapest keypad position is filled before moving to a more expensive one.
- A more advanced version of this problem introduces repeated letters, where frequency-based greedy sorting becomes necessary. This easier version does not require that extra step.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
