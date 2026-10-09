# 1541. Minimum Insertions to Balance a Parentheses String

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

You are given a string `s` that contains only the characters `'('` and `')'`.  

The string is considered balanced only when every left parenthesis `'('` has a matching pair of two consecutive right parentheses `'))'` that come after it. In other words, we treat `'('` as an opening bracket and `'))'` as its closing bracket.  

You are allowed to insert the characters `'('` or `')'` at any positions in the string. Your task is to find the minimum number of insertions required to make the string balanced.  

Input is a single string `s`. Output is a single integer representing the fewest insertions needed.

## Constraints

- 1 <= s.length <= 10^5
- s consists of '(' and ')' only

## Intuition

When I first read the problem, I noticed that a normal parentheses balance check would not work. Here every opening bracket needs two closing brackets that sit next to each other.  

A single `')'` is incomplete, and a pair of `'))'` without a previous unmatched `'('` also needs fixing.  

So the natural idea is to walk through the string once, keep a running count of unmatched openings, and decide on the spot whether a closing character needs an extra insertion or can close an existing opening. This greedy one-pass idea avoids building any new string and still guarantees the minimum number of changes.

## Approach

I keep two simple counters: one that tracks how many `'('` are still waiting for their pair of `'))'`, and another that records the total insertions I have made so far.  

I scan the string from left to right.  

- When I see a `'('`, I just increase the waiting-open count by one.  
- When I see a `')'`, I look at the next character. If it is also a `')'`, I treat the two characters as a complete closing pair and skip ahead. If it is not, I insert one extra `')'` in my mind and count that insertion.  
- After I have a complete pair (real or inserted), I check the waiting-open count. If it is zero, this pair has no opening to close, so I insert a `'('` as well. Otherwise I simply decrease the waiting-open count.  

When the scan finishes, every remaining unmatched `'('` still needs two right parentheses, so I add twice that number to the answer.  

The whole process finishes in a single pass and uses only a constant amount of extra memory.

## Data Structures Used

No complex data structures are required.  

I only use a few integer variables:  
- a counter for the number of unmatched openings,  
- a counter for the total insertions,  
- and the current index while scanning the string.  

These simple variables are enough because the problem can be solved greedily without remembering the exact positions of earlier brackets.

## Operations & Behavior Summary

1. Initialize insertion count and unmatched-open count to zero.  
2. Walk through every character of the string.  
3. On `'('`: increase the unmatched-open count.  
4. On `')'`:  
   - If the next character is also `')'`, consume both characters as a valid pair.  
   - Otherwise, count one insertion for the missing second `')'`.  
   - Then either decrease the unmatched-open count or count one insertion for a missing `'('`.  
5. After the loop, add twice the remaining unmatched-open count (each still needs two `')'`).  
6. Return the final insertion count.

## Complexity

| Complexity       | Value | Explanation |
|------------------|-------|-------------|
| Time Complexity  | O(n)  | n is the length of the input string. Every character is examined a constant number of times. |
| Space Complexity | O(1)  | Only a handful of integer variables are used. No extra arrays, stacks, or maps are allocated. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int minInsertions(string s) {
        int ans = 0, x = 0;
        int n = s.length();
        for (int i = 0; i < n; ++i) {
            if (s[i] == '(') {
                ++x;
            } else {
                if (i < n - 1 && s[i + 1] == ')') {
                    ++i;
                } else {
                    ++ans;
                }
                if (x == 0) {
                    ++ans;
                } else {
                    --x;
                }
            }
        }
        ans += x << 1;
        return ans;
    }
};
```

### Java
```java
class Solution {
    public int minInsertions(String s) {
        int ans = 0, x = 0;
        int n = s.length();
        for (int i = 0; i < n; ++i) {
            if (s.charAt(i) == '(') {
                ++x;
            } else {
                if (i < n - 1 && s.charAt(i + 1) == ')') {
                    ++i;
                } else {
                    ++ans;
                }
                if (x == 0) {
                    ++ans;
                } else {
                    --x;
                }
            }
        }
        ans += x << 1;
        return ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @return {number}
 */
var minInsertions = function(s) {
    let ans = 0, x = 0;
    const n = s.length;
    for (let i = 0; i < n; ++i) {
        if (s[i] === '(') {
            ++x;
        } else {
            if (i < n - 1 && s[i + 1] === ')') {
                ++i;
            } else {
                ++ans;
            }
            if (x === 0) {
                ++ans;
            } else {
                --x;
            }
        }
    }
    ans += x << 1;
    return ans;
};
```

### TypeScript
```typescript
function minInsertions(s: string): number {
    let ans = 0, x = 0;
    const n = s.length;
    for (let i = 0; i < n; ++i) {
        if (s[i] === '(') {
            ++x;
        } else {
            if (i < n - 1 && s[i + 1] === ')') {
                ++i;
            } else {
                ++ans;
            }
            if (x === 0) {
                ++ans;
            } else {
                --x;
            }
        }
    }
    ans += x << 1;
    return ans;
};
```

### Python3
```python
class Solution:
    def minInsertions(self, s: str) -> int:
        ans = x = 0
        i, n = 0, len(s)
        while i < n:
            if s[i] == '(':
                x += 1
            else:
                if i < n - 1 and s[i + 1] == ')':
                    i += 1
                else:
                    ans += 1
                if x == 0:
                    ans += 1
                else:
                    x -= 1
            i += 1
        ans += x << 1
        return ans
```

### Go
```go
func minInsertions(s string) int {
    ans, x, n := 0, 0, len(s)
    for i := 0; i < n; i++ {
        if s[i] == '(' {
            x++
        } else {
            if i < n-1 && s[i+1] == ')' {
                i++
            } else {
                ans++
            }
            if x == 0 {
                ans++
            } else {
                x--
            }
        }
    }
    ans += x << 1
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages; only the syntax changes.  

I begin by setting two integers to zero: one for the answer (insertions) and one for the number of openings that still need closing. I also store the length of the string so I never recompute it.  

Inside the main loop I examine the character at the current index.  

If the character is an opening parenthesis, I simply increment the unmatched-open counter. That records one more bracket that will later need two closing brackets.  

If the character is a closing parenthesis, I first look one step ahead. When the next character is also a closing parenthesis, I advance the index so both characters are consumed together. When the next character is not a closing parenthesis (or I am already at the end), I increase the insertion counter by one; this models the insertion of the missing second closing parenthesis.  

After a complete pair has been formed, I look at the unmatched-open counter. If it is still zero, the pair has nothing to close, so I insert an opening parenthesis and increase the answer. Otherwise I decrease the unmatched-open counter, meaning one earlier opening has now been satisfied.  

When the loop ends, any value left in the unmatched-open counter represents openings that never received their two closing brackets. Each of them requires two insertions, so I add twice that value to the answer.  

Edge cases are handled naturally:  
- A string of only openings produces 2 × length insertions.  
- A string of only closings forces insertions for missing openings and for incomplete pairs.  
- Already-balanced strings such as `"())"` leave both counters at zero and return 0.  

Because every language follows the same control flow, the same reasoning applies whether you are writing C++, Java, JavaScript, TypeScript, Python, or Go.

## Examples

**Example 1**  
Input: `s = "(()))"`  
Output: `1`  

Trace:  
- First `'('` → unmatched opens = 1  
- Second `'('` → unmatched opens = 2  
- Next two `')'` form a pair → unmatched opens = 1  
- Last `')'` is alone → insert one `')'`, then close the remaining open → unmatched opens = 0  
- Final answer = 1  

**Example 2**  
Input: `s = "())"`  
Output: `0`  

Trace:  
- `'('` → unmatched opens = 1  
- Next two `')'` form a pair → unmatched opens = 0  
- Nothing left, answer stays 0  

**Example 3**  
Input: `s = "))())("`  
Output: `3`  

Trace:  
- First two `')'` form a pair with no open → insert one `'('`  
- Next two `')'` form a pair with no open → insert another `'('`  
- Last `'('` remains unmatched → needs two `')'`  
- Total insertions = 3  

## How to Use / Run Locally

**C++**  
1. Copy the solution into a file named `main.cpp`.  
2. Compile: `g++ -std=c++17 main.cpp -o main`  
3. Run: `./main` (you may add a simple driver that reads a string and prints the result).

**Java**  
1. Place the code inside a class file `Solution.java`.  
2. Compile: `javac Solution.java`  
3. Run with a small test harness that creates an instance and calls `minInsertions`.

**JavaScript**  
1. Save the function in a file `solution.js`.  
2. Run with Node: `node solution.js` (add a few console.log test cases at the bottom).

**TypeScript**  
1. Save the function in `solution.ts`.  
2. Compile: `tsc solution.ts`  
3. Run the generated JavaScript with Node.

**Python3**  
1. Save the class in `solution.py`.  
2. Run: `python3 solution.py` (add a few print statements that call the method).

**Go**  
1. Save the function in `main.go`.  
2. Run: `go run main.go` (add a main function that tests a few strings).

## Notes & Optimizations

The algorithm already runs in linear time and constant space, which is optimal for the given constraints (n up to 10^5).  

A stack-based solution is possible but uses extra memory and is not necessary. The pure counter approach is both simpler and faster in practice.  

Watch out for the final multiplication by two: forgetting it is a common off-by-one source of wrong answers. Also make sure the look-ahead does not go past the end of the string; every language version above guards against that.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)