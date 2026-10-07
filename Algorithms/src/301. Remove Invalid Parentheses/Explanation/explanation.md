# LeetCode 301. Remove Invalid Parentheses Solution - Backtracking Approach

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

LeetCode 301. Remove Invalid Parentheses asks you to take a string that contains letters and parentheses. Your job is to delete the smallest possible number of parentheses so the remaining string becomes valid. A string is valid when every opening parenthesis has a matching closing one and they are properly ordered.

You must return every unique valid string that can be formed after the minimum number of deletions. The order of the strings in the answer does not matter.

This is a classic hard-level string and backtracking problem that appears frequently in coding interviews. The goal is not only to make the string valid, but to explore all different ways of reaching a valid result with the fewest removals.

## Constraints

- 1 <= s.length <= 25
- s consists of lowercase English letters and parentheses '(' and ')'
- There will be at most 20 parentheses in s

These tight limits make an exponential search feasible.

## Intuition

When I first looked at the problem, I realized that not every parenthesis needs to be examined the same way. Some opening and closing parentheses are clearly extra. If I can count exactly how many left parentheses and how many right parentheses must be removed, I can limit the search to only those deletion counts.

Because the string is short, trying every legal combination of keep-or-delete decisions is acceptable. The key insight is to never allow the balance of parentheses to go negative during the search and to stop only when the exact number of required deletions has been used.

## Approach

I solve the remove invalid parentheses problem with a two-phase method.

First I scan the string once and count the minimum number of left and right parentheses that must be deleted. I keep a running counter for unmatched opening parentheses. Every time I meet a closing parenthesis without a match, I increase the right-deletion count. Any leftover opening parentheses at the end become the left-deletion count.

Then I start a depth-first search. At each character I decide whether to keep it or delete it, but only when a deletion is still allowed. Letters are always kept. Opening parentheses can be kept (increasing the current open count) or deleted. Closing parentheses can be kept only when there is a matching open parenthesis, or deleted if deletions remain.

I build the current candidate string as I go. When I reach the end of the input and both deletion counters are zero and the open count is zero, I store the finished string in a set. The set automatically removes duplicates. At the end I return every unique string that was collected.

This guarantees that every returned string is valid and was obtained with the minimum number of removals.

## Data Structures Used

- A set (hash set) to store the final unique valid strings and avoid duplicates.
- A mutable string or character list that acts as the current path while building candidates during recursion.
- Simple integer counters for remaining left deletions, remaining right deletions, and the current open balance.

These structures keep the code clean and the memory usage low given the small input size.

## Operations & Behavior Summary

1. Walk through the input string and compute the exact number of left and right parentheses that must be removed.
2. Start a recursive search from index 0 with the pre-computed deletion limits and an open count of zero.
3. At each position:
   - If the character is a letter, always append it and continue.
   - If the character is '(', optionally delete it (if left deletions remain) or keep it and increase the open count.
   - If the character is ')', optionally delete it (if right deletions remain) or keep it only when the open count is positive.
4. When the end of the string is reached, accept the candidate only if both deletion counters and the open count are zero.
5. Collect every accepted candidate in a set and finally return the set as a list.

## Complexity

| Type              | Complexity | Explanation |
|-------------------|------------|-------------|
| Time Complexity   | O(2^n)     | n is the length of the string (at most 25). Each parenthesis offers a keep-or-delete choice, but the deletion limits and balance checks prune many branches early. |
| Space Complexity  | O(n)       | Recursion depth is at most n. Extra space is also used for the path and the set of results, both of which stay small under the given constraints. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    vector<string> removeInvalidParentheses(string s) {
        int left = 0, right = 0;
        for (char c : s) {
            if (c == '(') left++;
            else if (c == ')') {
                if (left > 0) left--;
                else right++;
            }
        }
        unordered_set<string> res;
        string path;
        dfs(s, 0, left, right, 0, path, res);
        return vector<string>(res.begin(), res.end());
    }
private:
    void dfs(const string& s, int i, int left, int right, int open, string& path, unordered_set<string>& res) {
        if (i == s.size()) {
            if (left == 0 && right == 0 && open == 0) res.insert(path);
            return;
        }
        char c = s[i];
        if (c != '(' && c != ')') {
            path.push_back(c);
            dfs(s, i + 1, left, right, open, path, res);
            path.pop_back();
            return;
        }
        if (c == '(') {
            if (left > 0) {
                dfs(s, i + 1, left - 1, right, open, path, res);
            }
            path.push_back(c);
            dfs(s, i + 1, left, right, open + 1, path, res);
            path.pop_back();
        } else {
            if (right > 0) {
                dfs(s, i + 1, left, right - 1, open, path, res);
            }
            if (open > 0) {
                path.push_back(c);
                dfs(s, i + 1, left, right, open - 1, path, res);
                path.pop_back();
            }
        }
    }
};
```

### Java
```java
class Solution {
    public List<String> removeInvalidParentheses(String s) {
        int left = 0, right = 0;
        for (char c : s.toCharArray()) {
            if (c == '(') left++;
            else if (c == ')') {
                if (left > 0) left--;
                else right++;
            }
        }
        Set<String> res = new HashSet<>();
        StringBuilder path = new StringBuilder();
        dfs(s, 0, left, right, 0, path, res);
        return new ArrayList<>(res);
    }
    private void dfs(String s, int i, int left, int right, int open, StringBuilder path, Set<String> res) {
        if (i == s.length()) {
            if (left == 0 && right == 0 && open == 0) res.add(path.toString());
            return;
        }
        char c = s.charAt(i);
        if (c != '(' && c != ')') {
            path.append(c);
            dfs(s, i + 1, left, right, open, path, res);
            path.deleteCharAt(path.length() - 1);
            return;
        }
        if (c == '(') {
            if (left > 0) {
                dfs(s, i + 1, left - 1, right, open, path, res);
            }
            path.append(c);
            dfs(s, i + 1, left, right, open + 1, path, res);
            path.deleteCharAt(path.length() - 1);
        } else {
            if (right > 0) {
                dfs(s, i + 1, left, right - 1, open, path, res);
            }
            if (open > 0) {
                path.append(c);
                dfs(s, i + 1, left, right, open - 1, path, res);
                path.deleteCharAt(path.length() - 1);
            }
        }
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @return {string[]}
 */
var removeInvalidParentheses = function(s) {
    let left = 0, right = 0;
    for (let c of s) {
        if (c === '(') left++;
        else if (c === ')') {
            if (left > 0) left--;
            else right++;
        }
    }
    const res = new Set();
    const path = [];
    function dfs(i, leftRem, rightRem, open) {
        if (i === s.length) {
            if (leftRem === 0 && rightRem === 0 && open === 0) res.add(path.join(''));
            return;
        }
        const c = s[i];
        if (c !== '(' && c !== ')') {
            path.push(c);
            dfs(i + 1, leftRem, rightRem, open);
            path.pop();
            return;
        }
        if (c === '(') {
            if (leftRem > 0) {
                dfs(i + 1, leftRem - 1, rightRem, open);
            }
            path.push(c);
            dfs(i + 1, leftRem, rightRem, open + 1);
            path.pop();
        } else {
            if (rightRem > 0) {
                dfs(i + 1, leftRem, rightRem - 1, open);
            }
            if (open > 0) {
                path.push(c);
                dfs(i + 1, leftRem, rightRem, open - 1);
                path.pop();
            }
        }
    }
    dfs(0, left, right, 0);
    return Array.from(res);
};
```

### TypeScript
```typescript
function removeInvalidParentheses(s: string): string[] {
    let left = 0, right = 0;
    for (const c of s) {
        if (c === '(') left++;
        else if (c === ')') {
            if (left > 0) left--;
            else right++;
        }
    }
    const res = new Set<string>();
    const path: string[] = [];
    function dfs(i: number, leftRem: number, rightRem: number, open: number): void {
        if (i === s.length) {
            if (leftRem === 0 && rightRem === 0 && open === 0) res.add(path.join(''));
            return;
        }
        const c = s[i];
        if (c !== '(' && c !== ')') {
            path.push(c);
            dfs(i + 1, leftRem, rightRem, open);
            path.pop();
            return;
        }
        if (c === '(') {
            if (leftRem > 0) {
                dfs(i + 1, leftRem - 1, rightRem, open);
            }
            path.push(c);
            dfs(i + 1, leftRem, rightRem, open + 1);
            path.pop();
        } else {
            if (rightRem > 0) {
                dfs(i + 1, leftRem, rightRem - 1, open);
            }
            if (open > 0) {
                path.push(c);
                dfs(i + 1, leftRem, rightRem, open - 1);
                path.pop();
            }
        }
    }
    dfs(0, left, right, 0);
    return Array.from(res);
}
```

### Python3
```python
class Solution:
    def removeInvalidParentheses(self, s: str) -> list[str]:
        left = right = 0
        for c in s:
            if c == '(':
                left += 1
            elif c == ')':
                if left > 0:
                    left -= 1
                else:
                    right += 1
        res = set()
        path = []
        def dfs(i, left_rem, right_rem, open_cnt):
            if i == len(s):
                if left_rem == 0 and right_rem == 0 and open_cnt == 0:
                    res.add(''.join(path))
                return
            c = s[i]
            if c != '(' and c != ')':
                path.append(c)
                dfs(i + 1, left_rem, right_rem, open_cnt)
                path.pop()
                return
            if c == '(':
                if left_rem > 0:
                    dfs(i + 1, left_rem - 1, right_rem, open_cnt)
                path.append(c)
                dfs(i + 1, left_rem, right_rem, open_cnt + 1)
                path.pop()
            else:
                if right_rem > 0:
                    dfs(i + 1, left_rem, right_rem - 1, open_cnt)
                if open_cnt > 0:
                    path.append(c)
                    dfs(i + 1, left_rem, right_rem, open_cnt - 1)
                    path.pop()
        dfs(0, left, right, 0)
        return list(res)
```

### Go
```go
func removeInvalidParentheses(s string) []string {
    left, right := 0, 0
    for _, c := range s {
        if c == '(' {
            left++
        } else if c == ')' {
            if left > 0 {
                left--
            } else {
                right++
            }
        }
    }
    res := make(map[string]struct{})
    path := make([]byte, 0, len(s))
    var dfs func(i, leftRem, rightRem, open int)
    dfs = func(i, leftRem, rightRem, open int) {
        if i == len(s) {
            if leftRem == 0 && rightRem == 0 && open == 0 {
                res[string(path)] = struct{}{}
            }
            return
        }
        c := s[i]
        if c != '(' && c != ')' {
            path = append(path, c)
            dfs(i+1, leftRem, rightRem, open)
            path = path[:len(path)-1]
            return
        }
        if c == '(' {
            if leftRem > 0 {
                dfs(i+1, leftRem-1, rightRem, open)
            }
            path = append(path, c)
            dfs(i+1, leftRem, rightRem, open+1)
            path = path[:len(path)-1]
        } else {
            if rightRem > 0 {
                dfs(i+1, leftRem, rightRem-1, open)
            }
            if open > 0 {
                path = append(path, c)
                dfs(i+1, leftRem, rightRem, open-1)
                path = path[:len(path)-1]
            }
        }
    }
    dfs(0, left, right, 0)
    ans := make([]string, 0, len(res))
    for k := range res {
        ans = append(ans, k)
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages. I first compute the minimum left and right deletions with a single linear scan. This step is crucial because it turns an unrestricted search into a tightly bounded one.

In the recursive function I pass the current index, the remaining left deletions, the remaining right deletions, and the current open count. I also maintain a path that holds the characters kept so far.

When the index reaches the end of the string I check three conditions at once: left deletions used up, right deletions used up, and open count back to zero. Only then is the path added to the result set.

For ordinary letters the decision is forced: keep the letter, recurse, then backtrack by removing it from the path.

For an opening parenthesis I always try the keep option (append and increase open). I also try the delete option when left deletions are still available. After each recursive call I restore the path so other branches stay clean.

For a closing parenthesis the keep option is allowed only when open is greater than zero. The delete option is allowed when right deletions remain. Again the path is restored after each call.

Because every recursive step either advances the index or reduces a deletion counter, the search cannot loop forever. The set guarantees that identical strings produced by different deletion sequences appear only once in the final answer.

Edge cases such as an empty result string, strings with only letters, or strings that need no deletions are handled naturally by the same logic.

## Examples

**Example 1**

Input: s = "()())()"

Output: ["(())()","()()()"]

Trace: The scan finds one extra right parenthesis. The search explores ways to delete exactly one ')'. Both "(())()" and "()()()" use one deletion and stay balanced, so both are kept.

**Example 2**

Input: s = "(a)())()"

Output: ["(a())()","(a)()()"]

Trace: Same extra right parenthesis. The letter 'a' is always kept. The two different places to delete the extra ')' produce the two distinct valid strings.

**Example 3**

Input: s = ")("

Output: [""]

Trace: One extra left and one extra right must be deleted. The only way to use exactly those deletions is to remove both parentheses, leaving the empty string.

## How to Use / Run Locally

**C++**  
Copy the code into a file named `main.cpp`. Compile with `g++ -std=c++17 main.cpp -o main` and run `./main`. Add a simple driver that calls the function and prints the result.

**Java**  
Place the code inside a class named `Solution` in a file `Solution.java`. Compile with `javac Solution.java` and run with a small main method that creates an instance and prints the returned list.

**JavaScript**  
Save the function in a file `solution.js`. Run it with Node.js: `node solution.js`. Add a few console.log statements to test the examples.

**TypeScript**  
Save the function in `solution.ts`. Compile with `tsc solution.ts` then run the generated JavaScript, or use `ts-node solution.ts` directly.

**Python3**  
Save the class in `solution.py`. Run with `python3 solution.py`. Add a short test block under `if __name__ == "__main__":` to print the results.

**Go**  
Place the function in a file `solution.go` inside a package main. Add a main function that calls it and prints the slice. Build and run with `go run solution.go`.

## Notes & Optimizations

The pre-computation of the exact number of left and right deletions is the most important optimization. Without it the search would explore many paths that delete more or fewer parentheses than necessary.

Using a set to store results removes the need for manual duplicate checking. Because the maximum number of parentheses is 20, the number of unique valid strings stays manageable.

An alternative BFS approach that removes one parenthesis at a time and stops at the first valid level also works and guarantees minimum removals. The backtracking version shown here is usually faster in practice because it never explores levels beyond the minimum deletion count.

The same framework can be adapted for related problems such as generating all valid parentheses or checking balanced strings with different types of brackets.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)