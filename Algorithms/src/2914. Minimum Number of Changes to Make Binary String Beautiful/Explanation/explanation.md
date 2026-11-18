# Minimum Number of Changes to Make Binary String Beautiful - LeetCode 2914 Solution

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
  - [TypeScript](#typescript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

You are given a binary string `s` of even length. The goal is to make this string beautiful by changing the fewest characters possible.

A string is called beautiful if you can split it into one or more substrings where:
- Every substring has even length
- Every substring contains only 0s or only 1s

You can flip any character from 0 to 1 or from 1 to 0. Return the minimum number of such flips needed.

This is a classic string manipulation problem that appears in coding interviews and competitive programming contests. The key challenge is finding an efficient way to decide which characters to change without checking every possible partition.

## Constraints

- 2 <= s.length <= 10^5
- s.length is always even
- s[i] is either '0' or '1'

## Intuition

The first useful observation is that any valid even-length block of identical characters can be broken down into pairs of size 2. Because of this, the problem reduces to making every consecutive pair of characters the same.

If two characters in a pair are already equal, no change is needed for that pair. If they differ, exactly one change is enough to make them match. Counting these mismatched pairs gives the minimum number of changes required to make the entire binary string beautiful.

This pair-wise view avoids the need to explore larger partitions and keeps the solution simple and optimal.

## Approach

Walk through the string two characters at a time. For each pair, check whether the two characters are the same. Whenever they differ, increase a counter by one. At the end, the counter holds the answer.

There is no need to decide what value each pair should become. The cost is always zero when the pair already matches and one when it does not. This works for every pair independently, so the total is simply the number of differing pairs.

## Data Structures Used

No extra data structures are required. A single integer variable tracks the number of changes. The input string is processed in place using only index access, which keeps both time and memory usage minimal.

## Operations & Behavior Summary

1. Initialize a counter to zero.
2. Start at index 0 and move forward by two positions each step.
3. Compare the character at the current index with the character immediately after it.
4. If the two characters are different, add one to the counter.
5. Continue until the end of the string is reached.
6. Return the final value of the counter.

## Complexity

| Type              | Complexity | Explanation                                                                 |
|-------------------|------------|-----------------------------------------------------------------------------|
| Time Complexity   | O(n)       | The algorithm examines each character exactly once while stepping by two. n is the length of the input string. |
| Space Complexity  | O(1)       | Only a constant amount of extra memory is used for the counter and loop variables. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int minChanges(string s) {
        int changes = 0;
        for (int i = 0; i < s.length(); i += 2) {
            if (s[i] != s[i + 1]) {
                changes++;
            }
        }
        return changes;
    }
};
```

### Java
```java
class Solution {
    public int minChanges(String s) {
        int changes = 0;
        for (int i = 0; i < s.length(); i += 2) {
            if (s.charAt(i) != s.charAt(i + 1)) {
                changes++;
            }
        }
        return changes;
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @return {number}
 */
var minChanges = function(s) {
    let changes = 0;
    for (let i = 0; i < s.length; i += 2) {
        if (s[i] !== s[i + 1]) {
            changes++;
        }
    }
    return changes;
};
```

### TypeScript
```typescript
function minChanges(s: string): number {
    let changes = 0;
    for (let i = 0; i < s.length; i += 2) {
        if (s[i] !== s[i + 1]) {
            changes++;
        }
    }
    return changes;
};
```

### Python3
```python
class Solution:
    def minChanges(self, s: str) -> int:
        changes = 0
        for i in range(0, len(s), 2):
            if s[i] != s[i + 1]:
                changes += 1
        return changes
```

### Go
```go
func minChanges(s string) int {
    changes := 0
    for i := 0; i < len(s); i += 2 {
        if s[i] != s[i+1] {
            changes++
        }
    }
    return changes
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core logic is identical across all six languages. A counter starts at zero. A loop begins at the first character and advances by two positions on every iteration. Inside the loop the current character is compared with the next one. When they do not match, the counter is increased by one. After the loop finishes, the counter is returned.

Because the string length is guaranteed to be even, the loop never tries to read past the end of the string. Edge cases such as a string of all identical characters produce a zero answer, which is correct. Strings where every pair differs produce an answer equal to half the length, which is also correct.

The only small differences between languages are syntactic:
- C++ and Java use traditional for-loops with index arithmetic.
- JavaScript and TypeScript use the same loop style but with `let` and string indexing.
- Python uses `range` with a step of 2.
- Go uses a classic for-loop with the same step-2 pattern.

None of these differences change the underlying algorithm or its complexity.

## Examples

**Example 1**  
Input: s = "1001"  
Output: 2  

Trace:  
- Pair "10" differs → 1 change  
- Pair "01" differs → 1 change  
- Total = 2  

One possible result after changes is "1100", which can be split into "11" and "00".

**Example 2**  
Input: s = "10"  
Output: 1  

Trace:  
- Only one pair "10" differs → 1 change  
- Result becomes "11" or "00", both beautiful.

**Example 3**  
Input: s = "0000"  
Output: 0  

Trace:  
- Pair "00" matches → 0  
- Pair "00" matches → 0  
- Total = 0  
The string is already beautiful.

## How to Use / Run Locally

**C++**  
1. Copy the code into a file named `main.cpp`.  
2. Compile with `g++ -std=c++17 main.cpp -o main`.  
3. Run with `./main`. You may need to add a small main function that reads a string and prints the result.

**Java**  
1. Place the code inside a class named `Solution` in a file `Solution.java`.  
2. Compile with `javac Solution.java`.  
3. Write a small driver class that creates an instance, calls `minChanges`, and prints the answer.

**JavaScript**  
1. Save the function in a file `solution.js`.  
2. Run with Node.js: `node solution.js`. Add a few console.log calls with sample inputs to test.

**TypeScript**  
1. Save the function in `solution.ts`.  
2. Compile with `tsc solution.ts` (requires TypeScript installed).  
3. Run the resulting JavaScript file with Node.js.

**Python3**  
1. Save the class in `solution.py`.  
2. Run with `python3 solution.py`. Add a few test calls at the bottom of the file.

**Go**  
1. Save the function in `main.go` inside a package main.  
2. Add a main function that calls `minChanges` with sample strings.  
3. Run with `go run main.go`.

## Notes & Optimizations

The pair-wise approach is already optimal. No further asymptotic improvement is possible because every character must be examined at least once in the worst case.

An alternative way of thinking is to force the entire string into alternating blocks of "00" and "11", but that would never require fewer changes than the simple pair count. The current solution already achieves the theoretical minimum.

The algorithm handles the full constraint range (up to 10^5 characters) comfortably because it runs in linear time and constant extra space.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
