# Generate Parentheses | LeetCode 22 Solution

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

The Generate Parentheses problem asks you to create every possible string made of well-formed parentheses when you are given n pairs.  

A well-formed string means every opening parenthesis has a matching closing one and they are properly nested. You never close more than you have opened at any point.  

You receive a single integer n. Your function must return a list of all valid strings that can be formed with exactly n pairs of parentheses.

## Constraints

- 1 <= n <= 8

Because n is tiny, even a solution that generates every Catalan-number sized set of strings stays fast enough.

## Intuition

When I first saw the Generate Parentheses problem I noticed two simple facts.  

First, any valid string must contain exactly n opening and n closing parentheses.  

Second, while I build the string I can never place a closing parenthesis if I already have more closes than opens so far.  

Those two observations immediately point to a recursive construction: keep track of how many opens and closes I still need, and only add a character when the balance stays legal. That is the classic backtracking idea for bracket sequences.

## Approach

I start with an empty path and two counters: remaining opens and remaining closes, both set to n.  

At every step I look at the two possible moves:  

- If I still have opens left, I add an opening parenthesis, decrease the open count, and continue.  
- If the remaining closes are greater than the remaining opens, I am allowed to add a closing parenthesis, decrease the close count, and continue.  

Whenever both counters reach zero the current path is a complete valid string, so I store it.  

After each recursive call I remove the last character so the path is ready for the next choice. This simple back-and-forth explores every legal combination without ever generating an invalid one.

## Data Structures Used

- A result list that stores every finished valid string.  
- A temporary path (string, string builder, or character list) that holds the parentheses being built right now.  
- Two integer counters that track how many opens and closes are still needed.  

No extra fancy structures are required. The recursion itself keeps the current state.

## Operations & Behavior Summary

1. Initialize an empty result list and an empty current path.  
2. Call a helper with remaining opens = n and remaining closes = n.  
3. If both remaining counts are zero, convert the path to a string and add it to the result.  
4. Otherwise try to place an opening parenthesis when opens remain.  
5. Try to place a closing parenthesis only when closes > opens.  
6. After each recursive call, undo the last placement so the next branch starts clean.  
7. When the helper returns, the result list contains every well-formed combination.

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time | O(4^n / sqrt(n)) | The number of valid strings is the nth Catalan number, which grows like 4^n / n^{1.5}. Each string has length 2n, so total work follows that growth. |
| Space | O(n) | Recursion depth is at most 2n and the temporary path uses O(n) space. The output list is required by the problem and is not counted as extra space. |

## Multi-language Solutions

### C++
```cpp
class Solution { 
public: 
    vector<string> generateParenthesis(int n) { 
        vector<string> ans;
        string cur;
        function<void(int,int)> dfs = [&](int open, int close) {
            if (open == 0 && close == 0) {
                ans.push_back(cur);
                return;
            }
            if (open > 0) {
                cur.push_back('(');
                dfs(open - 1, close);
                cur.pop_back();
            }
            if (close > open) {
                cur.push_back(')');
                dfs(open, close - 1);
                cur.pop_back();
            }
        };
        dfs(n, n);
        return ans;
    } 
};
```

### Java
```java
class Solution { 
    public List<String> generateParenthesis(int n) { 
        List<String> ans = new ArrayList<>();
        StringBuilder cur = new StringBuilder();
        dfs(n, n, cur, ans);
        return ans;
    }
    private void dfs(int open, int close, StringBuilder cur, List<String> ans) {
        if (open == 0 && close == 0) {
            ans.add(cur.toString());
            return;
        }
        if (open > 0) {
            cur.append('(');
            dfs(open - 1, close, cur, ans);
            cur.deleteCharAt(cur.length() - 1);
        }
        if (close > open) {
            cur.append(')');
            dfs(open, close - 1, cur, ans);
            cur.deleteCharAt(cur.length() - 1);
        }
    }
}
```

### JavaScript
```javascript
/** 
 * @param {number} n 
 * @return {string[]} 
 */ 
var generateParenthesis = function(n) { 
    const ans = [];
    const cur = [];
    function dfs(open, close) {
        if (open === 0 && close === 0) {
            ans.push(cur.join(''));
            return;
        }
        if (open > 0) {
            cur.push('(');
            dfs(open - 1, close);
            cur.pop();
        }
        if (close > open) {
            cur.push(')');
            dfs(open, close - 1);
            cur.pop();
        }
    }
    dfs(n, n);
    return ans;
};
```

### TypeScript
```typescript
function generateParenthesis(n: number): string[] { 
    const ans: string[] = [];
    const cur: string[] = [];
    function dfs(open: number, close: number): void {
        if (open === 0 && close === 0) {
            ans.push(cur.join(''));
            return;
        }
        if (open > 0) {
            cur.push('(');
            dfs(open - 1, close);
            cur.pop();
        }
        if (close > open) {
            cur.push(')');
            dfs(open, close - 1);
            cur.pop();
        }
    }
    dfs(n, n);
    return ans;
};
```

### Python3
```python
class Solution: 
    def generateParenthesis(self, n: int) -> list[str]: 
        ans = []
        cur = []
        def dfs(open, close):
            if open == 0 and close == 0:
                ans.append(''.join(cur))
                return
            if open > 0:
                cur.append('(')
                dfs(open - 1, close)
                cur.pop()
            if close > open:
                cur.append(')')
                dfs(open, close - 1)
                cur.pop()
        dfs(n, n)
        return ans
```

### Go
```go
func generateParenthesis(n int) []string { 
    ans := []string{}
    cur := []byte{}
    var dfs func(open, close int)
    dfs = func(open, close int) {
        if open == 0 && close == 0 {
            ans = append(ans, string(cur))
            return
        }
        if open > 0 {
            cur = append(cur, '(')
            dfs(open-1, close)
            cur = cur[:len(cur)-1]
        }
        if close > open {
            cur = append(cur, ')')
            dfs(open, close-1)
            cur = cur[:len(cur)-1]
        }
    }
    dfs(n, n)
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

All six implementations follow the same backtracking logic; only the syntax for building and undoing the current path differs.

I keep a shared result container and a mutable path. The recursive helper receives the two remaining counts.  

When both counts hit zero the path is complete, so I copy it into the result.  

If opens remain I append '(', recurse with opens decreased by one, then remove the character.  

If closes exceed opens I append ')', recurse with closes decreased by one, then remove the character.  

The order of the two checks is important: I always try the open branch first, then the close branch. Because I only close when the balance allows it, every finished string is guaranteed to be well-formed.  

In C++ I use a std::string and push_back / pop_back.  
In Java I use a StringBuilder and delete the last character.  
In JavaScript and TypeScript I keep an array of characters and join only when a complete string is ready.  
In Python I use a list of characters and join at the end of a valid path.  
In Go I use a byte slice and slice it back to the previous length after the recursive call.  

Edge cases are handled automatically. When n equals 1 the only path is "()". When n equals 0 (although the constraint starts at 1) the helper would simply return an empty list. No extra base-case code is needed beyond the zero-count check.

## Examples

**Example 1**  
Input: n = 3  
Expected output: ["((()))","(()())","(())()","()(())","()()()"]  

Trace:  
Start with opens = 3, closes = 3.  
I keep adding opens until the first three characters are "(((".  
Then I can only add closes, producing "((()))".  
Backtracking lets me place a close earlier, for example after two opens, producing "(()())", and so on. Every legal interleaving appears exactly once.

**Example 2**  
Input: n = 1  
Expected output: ["()"]  

Trace:  
I must place the single open first (the only legal move).  
Then I place the single close.  
The path becomes "()" and both counters reach zero, so the string is collected.

**Example 3**  
Input: n = 2  
Expected output: ["(())","()()"]  

Trace:  
One path places both opens first, then both closes: "(())".  
The other path places open, close, open, close: "()()".  
Any attempt to close before the matching open is rejected by the balance check.

## How to Use / Run Locally

Clone the repository and navigate to the language folder you want to try.

**C++**  
Compile with any recent g++:  
`g++ -std=c++17 solution.cpp -o solution`  
Run: `./solution`  
(You may need a small main that reads n and prints the returned vector.)

**Java**  
Compile: `javac Solution.java`  
Run: `java Solution`  
(Add a main method that calls generateParenthesis and prints the list.)

**JavaScript**  
Run with Node: `node solution.js`  
(Add a console.log of the function call at the bottom.)

**TypeScript**  
Compile first: `tsc solution.ts`  
Then run the generated JavaScript, or use ts-node: `ts-node solution.ts`

**Python3**  
Simply execute: `python3 solution.py`  
(Add a short test block under the class.)

**Go**  
Place the function in a package main file and run: `go run solution.go`  
(Add a main that prints the result for a chosen n.)

In every language the core function expects an integer n and returns the list of strings. You only need a tiny driver to feed it a value and display the answer.

## Notes & Optimizations

Because n never exceeds 8 the Catalan-number growth stays tiny, so the straightforward backtracking is already optimal for the given limits.  

An alternative dynamic-programming approach builds longer strings from shorter ones, but it uses more memory and does not improve asymptotic time.  

You can generate the strings in lexicographical order by always trying the open branch before the close branch, which the current code already does.  

If the problem ever asked for the count of valid strings instead of the strings themselves, a pure Catalan-number formula would be faster, but here we must produce the actual sequences.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)