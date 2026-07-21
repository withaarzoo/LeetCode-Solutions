# Maximize Active Section with Trade I - LeetCode 3499 Solution

A clean and optimized solution for **LeetCode 3499 - Maximize Active Section with Trade I**. This repository explains the intuition, approach, complexity analysis, and implementation strategy using an efficient linear-time algorithm. The solution is available in **C++**, **Java**, **JavaScript**, **Python**, and **Go**, making it useful for interview preparation, competitive programming, and DSA practice.

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
- [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

You are given a binary string where:

- `1` represents an active section.
- `0` represents an inactive section.

You can perform **at most one trade**.

A trade always happens in two steps.

First, you must choose a contiguous block of `1`s that is completely surrounded by `0`s and convert it into `0`s.

After that, choose a contiguous block of `0`s that is surrounded by `1`s and convert it into `1`s.

Your goal is to maximize the total number of active sections after performing the best possible trade.

The challenge is finding the optimal trade without checking every possible combination, since the string length can be as large as `100,000`.

---

## Constraints

| Constraint | Value |
|------------|-------|
| `1 <= n <= 10^5` | Length of the binary string |
| `s[i]` | Either `'0'` or `'1'` |

---

## Intuition

The first thing I noticed was that the trade always follows the same pattern.

When I remove a valid block of `1`s, its neighboring `0` blocks immediately become connected into one larger block. That new merged block is then converted back into `1`s.

After writing down a few examples on paper, I realized something interesting.

The size of the removed `1` block does not actually affect the final gain.

The only thing that matters is the combined size of the left and right neighboring `0` blocks.

That means instead of simulating every trade, I only need to find the internal `1` block whose neighboring `0` blocks have the largest total length.

This observation reduces the problem to a simple linear scan.

---

## Approach

I start by counting the total number of `1`s in the original string because that is my initial answer.

Next, I add a virtual `1` at both ends of the string.

This makes every edge case behave like a normal middle case and avoids writing extra conditions.

Then I split the entire string into consecutive runs of identical characters.

Each run stores:

- Whether it is a `0` block or a `1` block.
- The length of that block.

Now I iterate through every run.

Whenever I find a `1` block that has a `0` block on both sides, it is a valid candidate.

For every valid candidate, I calculate:

```
gain = leftZeroLength + rightZeroLength
```

I keep the maximum gain found during the scan.

Finally, I return:

```
originalOnes + maximumGain
```

If no valid trade exists, the maximum gain remains zero, so the original count is already the answer.

---

## Data Structures Used

### String

The input string is scanned once to count active sections and build the augmented string.

### Run-Length Encoding (RLE)

Instead of processing individual characters repeatedly, I group consecutive equal characters into runs.

Each run stores:

- Character (`0` or `1`)
- Length of the block

This makes it very easy to access neighboring blocks in constant time.

### Array / Vector / List

The runs are stored inside a dynamic array.

Every language uses its natural container:

- C++ → `vector`
- Java → `ArrayList`
- JavaScript → `Array`
- Python → `list`
- Go → `slice`

---

## Operations & Behavior Summary

The algorithm performs these steps:

1. Count the number of active sections.
2. Add virtual `1`s to both ends.
3. Convert the string into consecutive runs.
4. Visit every `1` block.
5. Ignore boundary `1` blocks that are not surrounded by `0`s.
6. Compute the gain using the neighboring `0` blocks.
7. Keep the largest gain.
8. Add the gain to the original number of active sections.
9. Return the final answer.

Every character is processed only a constant number of times, making the solution highly efficient.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | **O(n)** | Every character is visited only once while building and scanning the runs. |
| Space Complexity | **O(n)** | Extra space is used for the augmented string and run-length encoding structure. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxActiveSectionsAfterTrade(string s) {
        // Count active sections in the original string
        int ones = 0;
        for (char c : s)
            if (c == '1')
                ones++;

        // Add virtual '1' on both ends
        string t = "1" + s + "1";

        vector<pair<char, int>> runs;

        // Build run-length encoding
        for (char c : t) {
            if (runs.empty() || runs.back().first != c)
                runs.push_back({c, 1});
            else
                runs.back().second++;
        }

        int best = 0;

        // Check every internal 1-block
        for (int i = 1; i + 1 < (int)runs.size(); i++) {
            if (runs[i].first == '1' &&
                runs[i - 1].first == '0' &&
                runs[i + 1].first == '0') {

                // Gain equals left zero length + right zero length
                best = max(best, runs[i - 1].second + runs[i + 1].second);
            }
        }

        return ones + best;
    }
};
```

### Java

```java
class Solution {
    public int maxActiveSectionsAfterTrade(String s) {
        // Count active sections
        int ones = 0;
        for (char c : s.toCharArray()) {
            if (c == '1') ones++;
        }

        // Add virtual boundaries
        String t = "1" + s + "1";

        ArrayList<Character> type = new ArrayList<>();
        ArrayList<Integer> len = new ArrayList<>();

        // Run-length encoding
        for (char c : t.toCharArray()) {
            if (type.isEmpty() || type.get(type.size() - 1) != c) {
                type.add(c);
                len.add(1);
            } else {
                len.set(len.size() - 1, len.get(len.size() - 1) + 1);
            }
        }

        int best = 0;

        // Check every valid 1-block
        for (int i = 1; i + 1 < type.size(); i++) {
            if (type.get(i) == '1' &&
                type.get(i - 1) == '0' &&
                type.get(i + 1) == '0') {

                best = Math.max(best, len.get(i - 1) + len.get(i + 1));
            }
        }

        return ones + best;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @return {number}
 */
var maxActiveSectionsAfterTrade = function(s) {
    // Count original active sections
    let ones = 0;
    for (const ch of s) {
        if (ch === '1') ones++;
    }

    // Add virtual boundaries
    const t = "1" + s + "1";

    const runs = [];

    // Run-length encoding
    for (const ch of t) {
        if (runs.length === 0 || runs[runs.length - 1][0] !== ch) {
            runs.push([ch, 1]);
        } else {
            runs[runs.length - 1][1]++;
        }
    }

    let best = 0;

    // Try every internal 1-block
    for (let i = 1; i + 1 < runs.length; i++) {
        if (
            runs[i][0] === '1' &&
            runs[i - 1][0] === '0' &&
            runs[i + 1][0] === '0'
        ) {
            best = Math.max(best, runs[i - 1][1] + runs[i + 1][1]);
        }
    }

    return ones + best;
};
```

### Python3

```python
class Solution:
    def maxActiveSectionsAfterTrade(self, s: str) -> int:
        # Count active sections in the original string
        ones = s.count("1")

        # Add virtual boundaries
        t = "1" + s + "1"

        runs = []

        # Run-length encoding
        for ch in t:
            if not runs or runs[-1][0] != ch:
                runs.append([ch, 1])
            else:
                runs[-1][1] += 1

        best = 0

        # Check every internal 1-block
        for i in range(1, len(runs) - 1):
            if (
                runs[i][0] == "1"
                and runs[i - 1][0] == "0"
                and runs[i + 1][0] == "0"
            ):
                best = max(best, runs[i - 1][1] + runs[i + 1][1])

        return ones + best
```

### Go

```go
func maxActiveSectionsAfterTrade(s string) int {
 // Count original active sections
 ones := 0
 for _, ch := range s {
  if ch == '1' {
   ones++
  }
 }

 // Add virtual boundaries
 t := "1" + s + "1"

 type Run struct {
  ch  byte
  len int
 }

 runs := []Run{}

 // Run-length encoding
 for i := 0; i < len(t); i++ {
  if len(runs) == 0 || runs[len(runs)-1].ch != t[i] {
   runs = append(runs, Run{t[i], 1})
  } else {
   runs[len(runs)-1].len++
  }
 }

 best := 0

 // Check every valid 1-block
 for i := 1; i+1 < len(runs); i++ {
  if runs[i].ch == '1' &&
   runs[i-1].ch == '0' &&
   runs[i+1].ch == '0' {

   gain := runs[i-1].len + runs[i+1].len
   if gain > best {
    best = gain
   }
  }
 }

 return ones + best
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The algorithm behaves exactly the same in every programming language. Only the syntax changes.

First, the solution counts all active sections in the original string. This gives the starting answer before any trade is performed.

Next, two virtual `1`s are added to the beginning and end of the string. These extra characters simplify handling edge cases because every possible trade can now be treated in the same way.

The algorithm then creates consecutive runs of identical characters.

For example,

```
1100011100
```

becomes

```
11
000
111
00
```

Instead of checking every character, the algorithm only works with these larger blocks.

After the runs are built, each `1` block is examined.

If the block has a `0` block immediately before it and another `0` block immediately after it, then it is a valid trade candidate.

The potential increase is simply the combined lengths of those two neighboring `0` blocks.

The algorithm remembers the largest increase found during the scan.

Finally, that increase is added to the original number of active sections.

If no valid `1` block exists, no trade is possible, so the original count is returned unchanged.

This strategy avoids expensive simulations and guarantees the optimal answer in linear time.

---

## Examples

### Example 1

**Input**

```
s = "010"
```

**Output**

```
1
```

**Trace**

There is no `1` block surrounded by `0`s.

No trade can be performed.

The answer remains `1`.

---

### Example 2

**Input**

```
s = "0100"
```

**Output**

```
4
```

**Trace**

The middle `1` block joins the surrounding `0` blocks after removal.

That merged block is converted back into `1`s.

The final string becomes all active sections.

---

### Example 3

**Input**

```
s = "1000100"
```

**Output**

```
7
```

**Trace**

The internal `1` block produces the largest neighboring zero merge.

After the trade, every position becomes active.

The final answer is `7`.

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone https://github.com/your-username/your-repository.git
```

Move into the project directory.

```bash
cd your-repository
```

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

---

## Notes & Optimizations

- The algorithm works in linear time, making it suitable for the largest input size.
- Run-Length Encoding keeps the implementation clean and avoids repeated work.
- The virtual boundary characters eliminate complicated edge-case handling.
- No brute-force simulation is required.
- Every possible valid trade is considered exactly once.
- This approach is commonly used in competitive programming when consecutive segments matter more than individual characters.
- The same idea can be adapted to many binary string optimization problems involving contiguous groups.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
