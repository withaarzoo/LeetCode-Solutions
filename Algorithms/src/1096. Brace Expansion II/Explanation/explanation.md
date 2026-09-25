# LeetCode 1096. Brace Expansion II – Complete Solution in C++, Java, Python, JavaScript, TypeScript & Go

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

Brace Expansion II is a hard-level LeetCode problem that asks you to expand a given expression into every possible word it can produce. The expression follows a simple grammar. A single lowercase letter stands for itself. A comma-separated list inside braces means “choose any one of these options.” Writing two expressions next to each other means you must concatenate every word from the first expression with every word from the second. The final answer must be the sorted list of unique words that the whole expression can generate.

You are given one string called `expression`. You return a list of strings that contains every distinct word the expression represents, sorted in lexicographical order.

This problem is a classic exercise in recursive parsing, set union, and Cartesian product of strings. It appears frequently in coding interviews that test your ability to handle nested structures and combinatorial generation without producing duplicates.

## Constraints

- 1 ≤ expression.length ≤ 60
- expression consists only of the characters `{`, `}`, `,` and lowercase English letters
- The given expression is always valid according to the grammar described in the problem

## Intuition

When I first read the problem I noticed three clear operations: single letters, unions created by commas, and concatenations created by placing expressions next to each other. Nested braces made it obvious that a recursive approach would be natural. I realized I could walk through the string once with a shared index and let a recursive helper return the set of strings produced by the current sub-expression. Using a set automatically removes duplicates, which matches the problem requirement that each word appears only once in the final answer.

## Approach

I keep a single integer index that points to the current character of the expression. A recursive function called `parse` does the real work.  

Inside `parse` I maintain two sets. One set (called `cur`) holds the strings built so far for the current alternative. The other set (called `res`) collects finished alternatives.  

- When I see a letter I turn it into a one-element set and multiply it into `cur` (Cartesian product of strings).  
- When I see an opening brace I recursively parse everything inside the matching closing brace and multiply the resulting set into `cur`.  
- When I see a comma I dump the current alternative into `res`, reset `cur` to the empty string, and continue.  

At the end of the current scope I add whatever is left in `cur` into `res` and return that set. After the outermost call finishes I simply convert the set into a sorted list and return it.

This strategy processes the expression from left to right, respects nesting, and never generates duplicate words.

## Data Structures Used

- A mutable integer index that tracks the current position in the expression string.  
- Sets of strings to store intermediate and final results. Sets give automatic uniqueness and make the union operation cheap.  
- Temporary sets used for the Cartesian product of two existing sets.  

No extra stacks or queues are required because recursion itself manages the nesting of braces.

## Operations & Behavior Summary

1. Start at index 0 and call the recursive parse function.  
2. While the current character is not a closing brace (or the end of the string):  
   - If the character is a letter, form a singleton set and multiply it into the current working set.  
   - If the character is `{`, advance the index, recursively parse the content, advance past the matching `}`, and multiply the returned set into the current working set.  
   - If the character is a comma, save the current working set into the result set, reset the working set to the empty string, and continue.  
3. After the loop, save the last working set into the result set.  
4. Return the result set to the caller.  
5. At the top level convert the final set into a sorted list and return it.

## Complexity

| Complexity       | Value          | Explanation |
|------------------|----------------|-------------|
| Time Complexity  | O(n · s)       | n is the length of the expression (≤ 60). s is the number of distinct words produced. Each character is examined a constant number of times; the cost of set unions and Cartesian products is proportional to the size of the sets. |
| Space Complexity | O(s)           | The dominant space is the sets that hold intermediate and final strings. Recursion depth is at most O(n) because of nested braces, which is still linear in the input size. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    vector<string> braceExpansionII(string expression) {
        int i = 0;
        set<string> res = parse(expression, i);
        return vector<string>(res.begin(), res.end());
    }
private:
    set<string> parse(const string& expr, int& i) {
        set<string> res;
        set<string> cur{""};
        while (i < expr.size() && expr[i] != '}') {
            if (expr[i] == '{') {
                ++i;
                set<string> next = parse(expr, i);
                ++i;
                cur = product(cur, next);
            } else if (expr[i] == ',') {
                res.insert(cur.begin(), cur.end());
                cur = {""};
                ++i;
            } else {
                set<string> next{string(1, expr[i])};
                ++i;
                cur = product(cur, next);
            }
        }
        res.insert(cur.begin(), cur.end());
        return res;
    }
    set<string> product(const set<string>& a, const set<string>& b) {
        set<string> res;
        for (const string& x : a)
            for (const string& y : b)
                res.insert(x + y);
        return res;
    }
};
```

### Java
```java
class Solution {
    private int i;
    private String expr;
    public List<String> braceExpansionII(String expression) {
        this.expr = expression;
        this.i = 0;
        Set<String> res = parse();
        List<String> ans = new ArrayList<>(res);
        Collections.sort(ans);
        return ans;
    }
    private Set<String> parse() {
        Set<String> res = new TreeSet<>();
        Set<String> cur = new TreeSet<>();
        cur.add("");
        while (i < expr.length() && expr.charAt(i) != '}') {
            if (expr.charAt(i) == '{') {
                i++;
                Set<String> next = parse();
                i++;
                cur = product(cur, next);
            } else if (expr.charAt(i) == ',') {
                res.addAll(cur);
                cur = new TreeSet<>();
                cur.add("");
                i++;
            } else {
                Set<String> next = new TreeSet<>();
                next.add(String.valueOf(expr.charAt(i)));
                i++;
                cur = product(cur, next);
            }
        }
        res.addAll(cur);
        return res;
    }
    private Set<String> product(Set<String> a, Set<String> b) {
        Set<String> res = new TreeSet<>();
        for (String x : a)
            for (String y : b)
                res.add(x + y);
        return res;
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} expression
 * @return {string[]}
 */
var braceExpansionII = function(expression) {
    let i = 0;
    const parse = () => {
        const res = new Set();
        let cur = new Set([""]);
        while (i < expression.length && expression[i] !== '}') {
            if (expression[i] === '{') {
                i++;
                const next = parse();
                i++;
                cur = product(cur, next);
            } else if (expression[i] === ',') {
                for (const s of cur) res.add(s);
                cur = new Set([""]);
                i++;
            } else {
                const next = new Set([expression[i]]);
                i++;
                cur = product(cur, next);
            }
        }
        for (const s of cur) res.add(s);
        return res;
    };
    const product = (a, b) => {
        const res = new Set();
        for (const x of a)
            for (const y of b)
                res.add(x + y);
        return res;
    };
    const result = parse();
    return Array.from(result).sort();
};
```

### TypeScript
```typescript
function braceExpansionII(expression: string): string[] {
    let i = 0;
    const parse = (): Set<string> => {
        const res = new Set<string>();
        let cur = new Set<string>([""]);
        while (i < expression.length && expression[i] !== '}') {
            if (expression[i] === '{') {
                i++;
                const next = parse();
                i++;
                cur = product(cur, next);
            } else if (expression[i] === ',') {
                for (const s of cur) res.add(s);
                cur = new Set<string>([""]);
                i++;
            } else {
                const next = new Set<string>([expression[i]]);
                i++;
                cur = product(cur, next);
            }
        }
        for (const s of cur) res.add(s);
        return res;
    };
    const product = (a: Set<string>, b: Set<string>): Set<string> => {
        const res = new Set<string>();
        for (const x of a)
            for (const y of b)
                res.add(x + y);
        return res;
    };
    const result = parse();
    return Array.from(result).sort();
}
```

### Python3
```python
class Solution:
    def braceExpansionII(self, expression: str) -> list[str]:
        self.i = 0
        self.expr = expression
        res = self.parse()
        return sorted(res)
    def parse(self):
        res = set()
        cur = {""}
        while self.i < len(self.expr) and self.expr[self.i] != '}':
            if self.expr[self.i] == '{':
                self.i += 1
                nxt = self.parse()
                self.i += 1
                cur = self.product(cur, nxt)
            elif self.expr[self.i] == ',':
                res |= cur
                cur = {""}
                self.i += 1
            else:
                nxt = {self.expr[self.i]}
                self.i += 1
                cur = self.product(cur, nxt)
        res |= cur
        return res
    def product(self, a, b):
        return {x + y for x in a for y in b}
```

### Go
```go
func braceExpansionII(expression string) []string {
    i := 0
    var parse func() map[string]struct{}
    product := func(a, b map[string]struct{}) map[string]struct{} {
        res := make(map[string]struct{})
        for x := range a {
            for y := range b {
                res[x+y] = struct{}{}
            }
        }
        return res
    }
    parse = func() map[string]struct{} {
        res := make(map[string]struct{})
        cur := map[string]struct{}{"": {}}
        for i < len(expression) && expression[i] != '}' {
            if expression[i] == '{' {
                i++
                next := parse()
                i++
                cur = product(cur, next)
            } else if expression[i] == ',' {
                for s := range cur {
                    res[s] = struct{}{}
                }
                cur = map[string]struct{}{"": {}}
                i++
            } else {
                next := map[string]struct{}{string(expression[i]): {}}
                i++
                cur = product(cur, next)
            }
        }
        for s := range cur {
            res[s] = struct{}{}
        }
        return res
    }
    result := parse()
    ans := make([]string, 0, len(result))
    for s := range result {
        ans = append(ans, s)
    }
    sort.Strings(ans)
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

All six implementations follow the same high-level design, so the reasoning is identical across languages. Differences appear only in syntax and the concrete set type each language provides.

I start by declaring an integer index that will be shared by every recursive call. The main public function simply initializes this index to zero, calls the recursive helper, converts the returned set into a sorted list, and returns that list.

Inside the helper I create two sets. The first set (I call it `cur`) starts with a single empty string. That empty string is the identity element for concatenation; it lets me multiply the first real piece without writing special-case code. The second set (I call it `res`) will collect every finished alternative.

I then enter a loop that continues as long as I have not reached the end of the string and have not seen a closing brace that belongs to an outer call.

- When the current character is a lowercase letter I create a one-element set containing that letter, advance the index, and replace `cur` by the Cartesian product of the old `cur` and the new singleton.  
- When the current character is `{` I advance past it, recursively call the same helper (which will stop at the matching `}`), advance past that closing brace, and again replace `cur` by the product of the old `cur` and the set returned by the recursive call.  
- When the current character is a comma I know one alternative is finished. I insert every string currently in `cur` into `res`, reset `cur` back to a set containing only the empty string, and move the index forward.

After the loop finishes I still have the last alternative sitting in `cur`, so I insert those strings into `res` as well. The helper then returns `res`.

Because every intermediate collection is a set, duplicate words disappear automatically. The outermost call simply sorts the final set and hands the sorted list back to the caller.

Edge cases are handled naturally: a single letter produces a one-element list, nested braces are processed by the recursive call, and consecutive commas never appear because the expression is guaranteed to be valid.

## Examples

**Example 1**  
Input: `"{a,b}{c,{d,e}}"`  
Expected output: `["ac","ad","ae","bc","bd","be"]`  

Trace:  
- The first brace group produces the set `{"a","b"}`.  
- The second brace group produces the set `{"c","d","e"}`.  
- Their Cartesian product yields the six strings shown above.  
- Sorting leaves the order unchanged because the strings are already in lexicographical order.

**Example 2**  
Input: `"{{a,z},a{b,c},{ab,z}}"`  
Expected output: `["a","ab","ac","z"]`  

Trace:  
- The outermost brace contains three alternatives.  
- First alternative expands to `{"a","z"}`.  
- Second alternative expands to `{"ab","ac"}`.  
- Third alternative expands to `{"ab","z"}`.  
- The union of these three sets, after removing duplicates, is `{"a","ab","ac","z"}`.  
- Sorting produces the listed order.

**Example 3**  
Input: `"a{b,c}d"`  
Expected output: `["abd","acd"]`  

Trace:  
- The letter `a` multiplies with the set `{"b","c"}` giving `{"ab","ac"}`.  
- That result multiplies with the letter `d` giving the final two strings.

## How to Use / Run Locally

**C++**  
1. Copy the C++ code into a file named `main.cpp`.  
2. Compile with `g++ -std=c++17 main.cpp -o brace`.  
3. Run with `./brace`.  
4. You can hard-code a test expression inside `main` or read it from standard input.

**Java**  
1. Place the Java code inside a file named `Solution.java`.  
2. Compile with `javac Solution.java`.  
3. Run with `java Solution`.  
4. Add a `main` method that creates an instance and calls `braceExpansionII` with a sample string.

**JavaScript**  
1. Save the code as `braceExpansionII.js`.  
2. Run with Node.js: `node braceExpansionII.js`.  
3. Call the function with a test expression and print the returned array.

**TypeScript**  
1. Save the code as `braceExpansionII.ts`.  
2. Compile with `tsc braceExpansionII.ts`.  
3. Run the generated JavaScript file with `node braceExpansionII.js`.

**Python3**  
1. Save the code as `brace_expansion_ii.py`.  
2. Run with `python3 brace_expansion_ii.py`.  
3. Instantiate the class and call the method with any valid expression.

**Go**  
1. Save the code as `brace_expansion_ii.go`.  
2. Run with `go run brace_expansion_ii.go`.  
3. Add a `main` function that prints the result of `braceExpansionII` for a test case.

## Notes & Optimizations

The length constraint of 60 characters guarantees that even the largest possible set of words stays manageable. Using sets from the beginning avoids an extra post-processing step that would otherwise be needed to remove duplicates.  

An alternative iterative solution that uses an explicit stack is possible, but the recursive version is shorter, easier to reason about, and still efficient under the given constraints.  

If the problem ever allowed much longer expressions, one could switch to a bottom-up dynamic-programming approach that builds sets for every substring, but that would be overkill for the current limits.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)