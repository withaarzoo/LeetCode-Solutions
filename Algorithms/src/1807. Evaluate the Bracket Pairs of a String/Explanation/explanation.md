# 1807. Evaluate the Bracket Pairs of a String – LeetCode Solution

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

You are given a string that contains some bracket pairs. Each pair looks like `(key)` and the key inside is never empty. You also get a list of known key-value pairs called knowledge.

Your job is to go through the string and replace every `(key)` with the matching value from knowledge. If a key is missing from knowledge, you replace the whole pair with a single question mark `?`.

There are no nested brackets, and every opening bracket has a matching closing one. Keys from knowledge never repeat.

The final output is the new string after all replacements are done.

This is a classic string substitution problem that appears often in competitive programming and coding interviews under the name Evaluate the Bracket Pairs of a String.

## Constraints

- 1 <= s.length <= 10^5
- 0 <= knowledge.length <= 10^5
- knowledge[i].length == 2
- 1 <= key_i.length, value_i.length <= 10
- s consists of lowercase English letters and the characters '(' and ')'
- Every '(' has a matching ')'
- Keys inside brackets are never empty
- No nested brackets exist
- Keys and values contain only lowercase English letters
- Every key in knowledge is unique

## Intuition

The first thing I noticed is that the brackets never nest and every open has a close. That means I can safely walk the string from left to right and handle one pair at a time.

Looking up each key inside a long list would be too slow when the string is large. So the natural next step is to put every known pair into a hash map. After that, a single pass over the string is enough to build the answer.

## Approach

I first load the entire knowledge list into a hash map so any key can be found in constant time.

Then I create an empty result string (or a string builder in languages that need one).

I walk through the input character by character:

- If the current character is not '(', I simply copy it into the result.
- If I see '(', I keep moving until I find the matching ')'. Everything between them is the key.
- I look the key up in the map. If it exists I append its value; otherwise I append "?".
- After the closing parenthesis I continue from the next character.

Because the brackets never nest, this single left-to-right scan finishes the whole job.

## Data Structures Used

- Hash Map (unordered_map / HashMap / Map / dict): stores every key-value pair from knowledge. Chosen because lookups become O(1) on average, which is required for the large input sizes.
- String Builder / String Accumulation: used to build the final answer efficiently instead of creating many temporary strings.

## Operations & Behavior Summary

1. Build a map from the knowledge list.
2. Initialize an empty result.
3. Scan the string with an index.
4. When a normal letter appears, append it.
5. When '(' appears, collect the key until the matching ')'.
6. Replace the whole pair with the mapped value or with "?".
7. Move past the closing parenthesis and continue.
8. Return the finished result string.

## Complexity

| Type              | Value     | Explanation                                                                 |
|-------------------|-----------|-----------------------------------------------------------------------------|
| Time Complexity   | O(n + m)  | n is the length of s. m is the total size of all strings in knowledge. One pass over each. |
| Space Complexity  | O(m)      | The hash map stores the knowledge pairs. The result needs O(n) space for the output. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    string evaluate(string s, vector<vector<string>>& knowledge) {
        unordered_map<string, string> mp;
        for (auto& p : knowledge) mp[p[0]] = p[1];
        string res;
        int n = s.size();
        for (int i = 0; i < n; ) {
            if (s[i] == '(') {
                int j = i + 1;
                while (s[j] != ')') ++j;
                string key = s.substr(i + 1, j - i - 1);
                res += mp.count(key) ? mp[key] : "?";
                i = j + 1;
            } else {
                res += s[i++];
            }
        }
        return res;
    }
};
```

### Java
```java
class Solution {
    public String evaluate(String s, List<List<String>> knowledge) {
        Map<String, String> mp = new HashMap<>();
        for (List<String> p : knowledge) mp.put(p.get(0), p.get(1));
        StringBuilder res = new StringBuilder();
        int n = s.length();
        for (int i = 0; i < n; ) {
            if (s.charAt(i) == '(') {
                int j = i + 1;
                while (s.charAt(j) != ')') ++j;
                String key = s.substring(i + 1, j);
                res.append(mp.getOrDefault(key, "?"));
                i = j + 1;
            } else {
                res.append(s.charAt(i++));
            }
        }
        return res.toString();
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @param {string[][]} knowledge
 * @return {string}
 */
var evaluate = function(s, knowledge) {
    const mp = new Map();
    for (const [k, v] of knowledge) mp.set(k, v);
    let res = "";
    const n = s.length;
    for (let i = 0; i < n; ) {
        if (s[i] === '(') {
            let j = i + 1;
            while (s[j] !== ')') ++j;
            const key = s.substring(i + 1, j);
            res += mp.has(key) ? mp.get(key) : "?";
            i = j + 1;
        } else {
            res += s[i++];
        }
    }
    return res;
};
```

### TypeScript
```typescript
function evaluate(s: string, knowledge: string[][]): string {
    const mp = new Map<string, string>();
    for (const [k, v] of knowledge) mp.set(k, v);
    let res = "";
    const n = s.length;
    for (let i = 0; i < n; ) {
        if (s[i] === '(') {
            let j = i + 1;
            while (s[j] !== ')') ++j;
            const key = s.substring(i + 1, j);
            res += mp.has(key) ? mp.get(key)! : "?";
            i = j + 1;
        } else {
            res += s[i++];
        }
    }
    return res;
};
```

### Python3
```python
class Solution:
    def evaluate(self, s: str, knowledge: list[list[str]]) -> str:
        mp = {k: v for k, v in knowledge}
        res = []
        n = len(s)
        i = 0
        while i < n:
            if s[i] == '(':
                j = i + 1
                while s[j] != ')':
                    j += 1
                key = s[i + 1:j]
                res.append(mp.get(key, "?"))
                i = j + 1
            else:
                res.append(s[i])
                i += 1
        return "".join(res)
```

### Go
```go
func evaluate(s string, knowledge [][]string) string {
    mp := make(map[string]string, len(knowledge))
    for _, p := range knowledge {
        mp[p[0]] = p[1]
    }
    var res strings.Builder
    n := len(s)
    for i := 0; i < n; {
        if s[i] == '(' {
            j := i + 1
            for s[j] != ')' {
                j++
            }
            key := s[i+1 : j]
            if v, ok := mp[key]; ok {
                res.WriteString(v)
            } else {
                res.WriteByte('?')
            }
            i = j + 1
        } else {
            res.WriteByte(s[i])
            i++
        }
    }
    return res.String()
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core logic is identical across all six languages. Only the syntax for maps and string building changes.

I start by creating a hash map and inserting every pair from knowledge. This takes linear time relative to the knowledge size and gives me instant lookups later.

Next I prepare a result container. In C++ and Go I use a string that grows efficiently. In Java I use StringBuilder. In JavaScript, TypeScript and Python I collect pieces and join them at the end.

I keep a single index that walks through the input string.  
Whenever the character under the index is not an opening parenthesis I append it and move forward by one.

When I hit '(', I know a key begins at the next position. I advance a second index until I land on the matching ')'. The problem guarantees this always succeeds and that no nesting exists, so the second index never goes out of bounds.

I extract the substring that sits strictly between the two parentheses. That substring is the key.  
I query the map. If the key is present I append the stored value; otherwise I append a single question mark.

Finally I set the main index to the position right after the closing parenthesis so the next iteration continues cleanly.

This process repeats until the entire string has been processed. The result now contains the fully evaluated string and is ready to be returned.

Edge cases such as an empty knowledge list, keys that never appear, or strings with many repeated keys are all handled automatically by the map lookup and the single-pass scan.

## Examples

**Example 1**  
Input: s = "(name)is(age)yearsold", knowledge = [["name","bob"],["age","two"]]  
Output: "bobistwoyearsold"  

Trace:  
- Map becomes {name → bob, age → two}  
- "(name)" is replaced by "bob"  
- "is" is copied as-is  
- "(age)" is replaced by "two"  
- "yearsold" is copied as-is  

**Example 2**  
Input: s = "hi(name)", knowledge = [["a","b"]]  
Output: "hi?"  

Trace:  
- Map contains only "a"  
- "(name)" is not found, so it becomes "?"  

**Example 3**  
Input: s = "(a)(a)(a)aaa", knowledge = [["a","yes"]]  
Output: "yesyesyesaaa"  

Trace:  
- Every "(a)" is replaced by "yes"  
- The three plain letters "aaa" stay unchanged  

## How to Use / Run Locally

**C++**  
1. Save the code in a file named `main.cpp`.  
2. Compile: `g++ -std=c++17 main.cpp -o main`  
3. Run: `./main`  

**Java**  
1. Save the code in a file named `Solution.java`.  
2. Compile: `javac Solution.java`  
3. Run: `java Solution`  

**JavaScript**  
1. Save the code in a file named `evaluate.js`.  
2. Run: `node evaluate.js`  

**TypeScript**  
1. Save the code in a file named `evaluate.ts`.  
2. Compile: `tsc evaluate.ts`  
3. Run the generated JavaScript with `node evaluate.js`  

**Python3**  
1. Save the code in a file named `evaluate.py`.  
2. Run: `python3 evaluate.py`  

**Go**  
1. Save the code in a file named `main.go`.  
2. Run: `go run main.go`  

In each case you will need to add a small main function or test harness that creates sample inputs and prints the result.

## Notes & Optimizations

The solution already runs in linear time, which is optimal for this problem.  

Because every key is unique in knowledge and keys are short (at most 10 characters), the hash map never becomes a bottleneck.  

An alternative approach would be to use repeated string replacement, but that would be slower for large inputs and harder to control. The single-pass method is both cleaner and faster.  

Watch out for the empty-knowledge case and for strings that contain no brackets at all; both are handled correctly by the same logic.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)