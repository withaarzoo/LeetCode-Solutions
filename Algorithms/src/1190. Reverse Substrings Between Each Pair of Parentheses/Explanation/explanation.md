# 1190. Reverse Substrings Between Each Pair of Parentheses – LeetCode Solution

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

You are given a string that contains only lowercase English letters and parentheses. The parentheses are always balanced. Your task is to reverse the substring that sits between every pair of matching parentheses. You must start with the innermost pairs and work your way outward. After all reversals are done, the final string must not contain any parentheses at all.

This is a classic string manipulation problem that appears in interviews and weekly contests. It tests how well you can handle nested structures and efficient traversal without repeatedly reversing large pieces of text.

## Constraints

- 1 <= s.length <= 2000
- s contains only lowercase English letters and the characters '(' and ')'
- All parentheses in s are balanced

## Intuition

When I first looked at the problem, I noticed that the parentheses act like boundaries that tell me exactly which parts of the string need to be flipped. The deepest pairs must be handled first, so a simple left-to-right reverse would not work.  

I realized I could pre-compute the matching partner of every parenthesis. Once I know where each opening parenthesis pairs with its closing one, I can jump from one to the other and change direction. Changing direction is what actually produces the reversed order, so I never need to reverse substrings by hand.

## Approach

I walk through the string once and use a stack to remember the positions of every opening parenthesis. Whenever I meet a closing parenthesis, I pop the matching opening index and record both directions in a pair array.  

After that preparation step I start at the left end of the string with a forward direction. I move one character at a time. If the current character is a letter I add it to the result. If it is a parenthesis I jump straight to its partner and flip the direction. Because the direction changes every time I cross a pair, the letters come out in the exact order the problem asks for. I stop when the index walks past either end of the string.

## Data Structures Used

- An array (or vector) of size n that stores the matching index of every parenthesis. This lets me jump in constant time.
- A stack that temporarily holds the indices of opening parentheses while I build the matching pairs. The stack guarantees that the most recent unmatched opening parenthesis is always on top.

These two structures together give me everything I need for an efficient linear pass.

## Operations & Behavior Summary

1. Create a pair array of length n and an empty stack.  
2. Scan the string from left to right.  
   - Push the index of every '(' onto the stack.  
   - When a ')' appears, pop the top index, store the mutual links in the pair array.  
3. Initialize an empty result, set the current index to 0 and the direction to +1.  
4. While the index stays inside the string:  
   - If the character is a parenthesis, jump to its partner and reverse the direction.  
   - Otherwise append the letter to the result.  
   - Add the current direction to the index.  
5. Return the finished result string.

## Complexity

| Metric            | Value | Explanation |
|-------------------|-------|-------------|
| Time Complexity   | O(n)  | One linear pass builds the matching pairs; a second linear pass builds the answer. Every character is visited a constant number of times. |
| Space Complexity  | O(n)  | The pair array and the temporary stack both grow linearly with the length of the input string. The result string also needs O(n) space in the worst case. |

## Multi-language Solutions

### C++
```cpp
class Solution { 
public: 
    string reverseParentheses(string s) { 
        int n = s.size();
        vector<int> pair(n);
        stack<int> st;
        for (int i = 0; i < n; ++i) {
            if (s[i] == '(') st.push(i);
            else if (s[i] == ')') {
                int j = st.top(); st.pop();
                pair[i] = j;
                pair[j] = i;
            }
        }
        string res;
        int i = 0, dir = 1;
        while (i >= 0 && i < n) {
            if (s[i] == '(' || s[i] == ')') {
                i = pair[i];
                dir = -dir;
            } else {
                res += s[i];
            }
            i += dir;
        }
        return res;
    } 
};
```

### Java
```java
class Solution { 
    public String reverseParentheses(String s) { 
        int n = s.length();
        int[] pair = new int[n];
        Deque<Integer> st = new ArrayDeque<>();
        for (int i = 0; i < n; ++i) {
            if (s.charAt(i) == '(') st.push(i);
            else if (s.charAt(i) == ')') {
                int j = st.pop();
                pair[i] = j;
                pair[j] = i;
            }
        }
        StringBuilder res = new StringBuilder();
        int i = 0, dir = 1;
        while (i >= 0 && i < n) {
            if (s.charAt(i) == '(' || s.charAt(i) == ')') {
                i = pair[i];
                dir = -dir;
            } else {
                res.append(s.charAt(i));
            }
            i += dir;
        }
        return res.toString();
    } 
}
```

### JavaScript
```javascript
/** 
 * @param {string} s 
 * @return {string} 
 */ 
var reverseParentheses = function(s) { 
    const n = s.length;
    const pair = new Array(n);
    const st = [];
    for (let i = 0; i < n; ++i) {
        if (s[i] === '(') st.push(i);
        else if (s[i] === ')') {
            const j = st.pop();
            pair[i] = j;
            pair[j] = i;
        }
    }
    let res = '';
    let i = 0, dir = 1;
    while (i >= 0 && i < n) {
        if (s[i] === '(' || s[i] === ')') {
            i = pair[i];
            dir = -dir;
        } else {
            res += s[i];
        }
        i += dir;
    }
    return res;
};
```

### TypeScript
```typescript
function reverseParentheses(s: string): string {
    const n = s.length;
    const pair: number[] = new Array(n);
    const st: number[] = [];
    for (let i = 0; i < n; ++i) {
        if (s[i] === '(') st.push(i);
        else if (s[i] === ')') {
            const j = st.pop()!;
            pair[i] = j;
            pair[j] = i;
        }
    }
    let res = '';
    let i = 0, dir = 1;
    while (i >= 0 && i < n) {
        if (s[i] === '(' || s[i] === ')') {
            i = pair[i];
            dir = -dir;
        } else {
            res += s[i];
        }
        i += dir;
    }
    return res;
};
```

### Python3
```python
class Solution:
    def reverseParentheses(self, s: str) -> str:
        n = len(s)
        pair = [0] * n
        st = []
        for i in range(n):
            if s[i] == '(':
                st.append(i)
            elif s[i] == ')':
                j = st.pop()
                pair[i] = j
                pair[j] = i
        res = []
        i, dir = 0, 1
        while 0 <= i < n:
            if s[i] in '()':
                i = pair[i]
                dir = -dir
            else:
                res.append(s[i])
            i += dir
        return ''.join(res)
```

### Go
```go
func reverseParentheses(s string) string {
    n := len(s)
    pair := make([]int, n)
    st := []int{}
    for i := 0; i < n; i++ {
        if s[i] == '(' {
            st = append(st, i)
        } else if s[i] == ')' {
            j := st[len(st)-1]
            st = st[:len(st)-1]
            pair[i] = j
            pair[j] = i
        }
    }
    var res []byte
    i, dir := 0, 1
    for i >= 0 && i < n {
        if s[i] == '(' || s[i] == ')' {
            i = pair[i]
            dir = -dir
        } else {
            res = append(res, s[i])
        }
        i += dir
    }
    return string(res)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core logic is identical across all six languages; only the syntax changes.

First I allocate a pair array of the same length as the input. I also create an empty stack that will hold integer indices.  

I loop from 0 to n-1. When I see an opening parenthesis I push its index. When I see a closing parenthesis I pop the most recent opening index. Those two positions are partners, so I write each one into the pair array pointing at the other. After this loop every parenthesis knows exactly where its match lives.

Next I prepare an empty result container and two variables: the current index (starting at 0) and a direction flag (starting at +1).  

I keep moving while the index stays inside the string bounds. At every step I examine the character under the current index.  

- If it is a parenthesis I replace the index with the value stored in the pair array and flip the sign of the direction. That single jump moves me past the whole pair and reverses the travel direction, which is what produces the reversed letter order.  
- If it is an ordinary letter I simply append it to the result.  

Then I add the current direction to the index. Because the direction keeps flipping inside nested pairs, the letters are collected from the inside out exactly as the problem requires.  

When the index finally walks off either end of the string the loop ends and I return the result. Edge cases such as a string with no parentheses or a single pair surrounding the whole string are handled automatically by the same logic.

## Examples

**Example 1**  
Input: `s = "(abcd)"`  
Output: `"dcba"`  

Trace:  
- pair[0] = 5, pair[5] = 0  
- Start at 0, hit '(', jump to 5, direction becomes -1  
- Move left: collect d, c, b, a  
- Reach index -1 and stop  

**Example 2**  
Input: `s = "(u(love)i)"`  
Output: `"iloveu"`  

Trace:  
- Matching pairs: (0,8) and (2,7)  
- Direction flips twice; the inner “love” is collected right-to-left, then the outer letters are collected left-to-right after the second flip, producing “iloveu”.  

**Example 3**  
Input: `s = "(ed(et(oc))el)"`  
Output: `"leetcode"`  

Trace:  
- Three nested pairs are recorded.  
- Each jump reverses direction, so the letters are visited in the order that yields “leetcode” after all flips.

## How to Use / Run Locally

**C++**  
Save the code in a file named `main.cpp`. Compile with  
`g++ -std=c++17 main.cpp -o main`  
Run with  
`./main`

**Java**  
Save the code in a file named `Solution.java`. Compile with  
`javac Solution.java`  
Run with  
`java Solution`

**JavaScript**  
Save the code in a file named `solution.js`. Run with  
`node solution.js`

**TypeScript**  
Save the code in a file named `solution.ts`. Compile with  
`tsc solution.ts`  
Run the generated JavaScript with  
`node solution.js`

**Python3**  
Save the code in a file named `solution.py`. Run with  
`python3 solution.py`

**Go**  
Save the code in a file named `main.go`. Run with  
`go run main.go`

In every language you will need to add a small driver that reads a test string and prints the returned result. The solution function itself is ready to drop into any online judge or local test harness.

## Notes & Optimizations

The algorithm already runs in linear time, which is optimal for this problem. An alternative approach that repeatedly finds and reverses innermost pairs also works, but its worst-case time is quadratic. For the given constraint of n ≤ 2000 both methods are fast enough; the linear version is simply cleaner and scales better.

Because the parentheses are guaranteed to be balanced, no extra validation is required. Empty strings and strings that contain only letters are handled correctly by the same code path.

If you need to support multiple queries on the same string you can pre-compute the pair array once and reuse it, but that is outside the scope of the original problem.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)