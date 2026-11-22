# Check if All A's Appears Before All B's – LeetCode 2124 Solution

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

You are given a string `s` that contains only the characters `'a'` and `'b'`. The task is to check whether every `'a'` appears before every `'b'` in the string. If this condition holds, return `true`. Otherwise return `false`.

In simple terms, once a `'b'` shows up, no `'a'` is allowed to appear after it. Strings made of only `'a'`s or only `'b'`s are valid. This is a classic string order checking problem often seen in beginner DSA practice and LeetCode easy-level string questions.

## Constraints

- `1 <= s.length <= 100`
- `s[i]` is either `'a'` or `'b'`

## Intuition

When I first read the problem, I noticed that the relative order of all `'a'`s and `'b'`s is the only thing that matters. The moment I see a `'b'`, any later `'a'` immediately breaks the rule. So the natural idea is to walk through the string once while remembering whether I have already seen a `'b'`. That single piece of information is enough to decide the answer.

## Approach

I keep a boolean flag that starts as `false`. As I scan each character:

- If the current character is `'b'`, I set the flag to `true`.
- If the current character is `'a'` and the flag is already `true`, I know a `'b'` came earlier, so I return `false` at once.
- If I finish the whole string without hitting that situation, I return `true`.

This single-pass check covers every valid and invalid case without needing extra storage or multiple scans.

## Data Structures Used

- A single boolean variable (`seenB` or equivalent). It records whether a `'b'` has appeared so far. No arrays, maps, or stacks are required because the problem only cares about the relative order of two characters.

## Operations & Behavior Summary

1. Initialize a flag to `false` (no `'b'` seen yet).
2. Iterate over every character of the input string.
3. When a `'b'` is found, flip the flag to `true`.
4. When an `'a'` is found after the flag is `true`, immediately reject the string.
5. If the loop ends normally, accept the string.

The algorithm never looks ahead or stores positions; it only reacts to the current character and the remembered state.

## Complexity

| Complexity       | Value | Explanation |
|------------------|-------|-------------|
| Time Complexity  | O(n)  | We examine each of the n characters in the string exactly once. |
| Space Complexity | O(1)  | Only a constant amount of extra memory (one boolean) is used. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    bool checkString(string s) {
        bool seenB = false;
        for (char c : s) {
            if (c == 'b') {
                seenB = true;
            } else if (seenB) {
                return false;
            }
        }
        return true;
    }
};
```

### Java
```java
class Solution {
    public boolean checkString(String s) {
        boolean seenB = false;
        for (char c : s.toCharArray()) {
            if (c == 'b') {
                seenB = true;
            } else if (seenB) {
                return false;
            }
        }
        return true;
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @return {boolean}
 */
var checkString = function(s) {
    let seenB = false;
    for (let c of s) {
        if (c === 'b') {
            seenB = true;
        } else if (seenB) {
            return false;
        }
    }
    return true;
};
```

### TypeScript
```typescript
function checkString(s: string): boolean {
    let seenB = false;
    for (const c of s) {
        if (c === 'b') {
            seenB = true;
        } else if (seenB) {
            return false;
        }
    }
    return true;
};
```

### Python3
```python
class Solution:
    def checkString(self, s: str) -> bool:
        seenB = False
        for c in s:
            if c == 'b':
                seenB = True
            elif seenB:
                return False
        return True
```

### Go
```go
func checkString(s string) bool {
    seenB := false
    for _, c := range s {
        if c == 'b' {
            seenB = true
        } else if seenB {
            return false
        }
    }
    return true
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages, so the reasoning stays the same.

I declare a boolean that starts false. This boolean is the only memory I need.

Then I walk the string character by character.

- When the character is `'b'`, I mark that a `'b'` has been seen. From this point forward the mark stays set.
- When the character is `'a'`, I look at the mark. If the mark is already set, a `'b'` appeared earlier and the order is wrong, so I return false immediately.
- If the mark is still unset, the `'a'` is still in the allowed region, so I simply continue.

After the loop finishes, every `'a'` must have appeared before any `'b'`, or there were no `'a'`s, or there were no `'b'`s. In all those cases I return true.

Edge cases fall out naturally:
- All `'a'`s → flag never becomes true → true.
- All `'b'`s → flag becomes true but no later `'a'` appears → true.
- Empty string is impossible because of the length constraint.
- A single character is always true.

No language-specific tricks are needed; the same flag-and-scan pattern works cleanly in C++, Java, JavaScript, TypeScript, Python, and Go.

## Examples

**Example 1**  
Input: `s = "aaabbb"`  
Output: `true`  
Trace: The first three characters are `'a'`, flag stays false. Then three `'b'`s set the flag, but no `'a'` follows. Result is true.

**Example 2**  
Input: `s = "abab"`  
Output: `false`  
Trace: First `'a'` (flag still false). Then `'b'` sets the flag. Next `'a'` finds the flag already true → return false.

**Example 3**  
Input: `s = "bbb"`  
Output: `true`  
Trace: Every character is `'b'`, so the flag becomes true, yet no `'a'` ever appears. Result is true.

## How to Use / Run Locally

**C++**  
Save the code in a file named `solution.cpp`. Compile with `g++ solution.cpp -o solution` and run `./solution`. You can add a small `main` that reads a string and prints the result for quick testing.

**Java**  
Save the class in `Solution.java`. Compile with `javac Solution.java` and run `java Solution`. Add a `main` method that creates an instance and calls `checkString` with sample inputs.

**JavaScript**  
Save the function in `solution.js`. Run with Node: `node solution.js`. You can call the function directly at the bottom of the file with test strings and log the results.

**TypeScript**  
Save the function in `solution.ts`. Compile with `tsc solution.ts` then run the generated JavaScript, or use `ts-node solution.ts` if you have it installed.

**Python3**  
Save the class in `solution.py`. Run with `python3 solution.py`. Add a few print statements under an `if __name__ == "__main__":` block to test the method.

**Go**  
Save the function in `solution.go`. Run with `go run solution.go`. You can add a `main` function that prints the results of several test cases.

## Notes & Optimizations

The solution already runs in linear time and constant space, which is optimal for this problem. An alternative is to search for the substring `"ba"`; if it exists the answer is false. That approach is equally correct but may create a temporary string or perform an extra scan depending on the language’s string library. The flag method avoids any extra allocation and stays easy to understand for beginners learning string traversal and state tracking.

Watch out for the case where the string ends with `'a'` after a `'b'` has already appeared — that is the classic failure pattern the flag catches instantly.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
