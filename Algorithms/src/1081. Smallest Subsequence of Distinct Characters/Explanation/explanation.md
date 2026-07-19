# 1081. Smallest Subsequence of Distinct Characters

A clean and optimized solution for **LeetCode 1081 - Smallest Subsequence of Distinct Characters** using a **Greedy Algorithm** and **Monotonic Stack**. This repository explains the intuition, approach, complexity analysis, and provides multi-language solutions for C++, Java, JavaScript, Python, and Go.

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

## Problem Summary

The goal is to return the **lexicographically smallest subsequence** that contains every distinct character from the given string exactly once.

Unlike removing duplicate characters randomly, the order of characters must still follow the original string because we are building a subsequence, not rearranging characters.

The input is a lowercase English string, and the output should contain every unique character exactly once while being the smallest possible in dictionary order.

This is a classic **Greedy + Monotonic Stack** problem and is one of the most popular interview questions related to stacks.

## Constraints

| Constraint | Value |
|------------|-------|
| String Length | 1 ≤ s.length ≤ 1000 |
| Characters | Lowercase English letters only |

## Intuition

The first thing I noticed was that every unique character must appear exactly once in the final answer.

Whenever I see a smaller character, I would like to place it earlier because that makes the final string smaller in lexicographical order.

However, I cannot simply remove larger characters whenever I want. If a larger character never appears again later in the string, removing it would make it impossible to include every distinct character.

That observation naturally leads to a greedy strategy. I only remove a character if I know I can safely add it back later.

A stack is perfect for this because I only need to compare the current character with the most recently chosen one.

## Approach

I first count how many times every character appears.

While scanning the string from left to right, I keep a stack that stores the current best answer.

For every character:

1. Decrease its remaining frequency.
2. Skip it if it already exists in the stack.
3. Compare it with the top of the stack.
4. While the top character is larger and still appears later, remove it.
5. Push the current character into the stack.

At the end, the stack itself becomes the final answer.

This greedy approach always produces the smallest valid subsequence.

## Data Structures Used

| Data Structure | Purpose |
|---------------|---------|
| Stack | Stores the current best subsequence while processing the string |
| Frequency Array | Tracks how many times each character will still appear |
| Visited Array | Prevents duplicate characters from entering the stack |

## Operations & Behavior Summary

The algorithm performs the following operations:

- Count the frequency of every character.
- Create an empty stack.
- Track which characters are already inside the stack.
- Process every character from left to right.
- Skip duplicate characters already included.
- Remove larger characters if they appear again later.
- Push the current character into the stack.
- Return the stack as the final answer.

Every character is pushed once and popped at most once, making the algorithm very efficient.

## Complexity

| Type | Complexity | Explanation |
|------|------------|-------------|
| Time Complexity | O(n) | Every character is pushed and popped at most one time. |
| Space Complexity | O(1) | Only fixed-size arrays and a stack containing at most 26 lowercase letters are used. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string smallestSubsequence(string s) {
        // Store how many times each character still appears
        vector<int> freq(26, 0);
        for (char c : s)
            freq[c - 'a']++;

        // Track whether a character is already inside the stack
        vector<bool> inStack(26, false);

        // This string works as our stack
        string st;

        for (char c : s) {
            // One occurrence of this character has now been processed
            freq[c - 'a']--;

            // Skip duplicate characters already included
            if (inStack[c - 'a'])
                continue;

            // Remove larger characters if they appear again later
            while (!st.empty() &&
                   st.back() > c &&
                   freq[st.back() - 'a'] > 0) {

                // Mark removed character as no longer inside the stack
                inStack[st.back() - 'a'] = false;
                st.pop_back();
            }

            // Add current character to the stack
            st.push_back(c);
            inStack[c - 'a'] = true;
        }

        // The stack already contains the final answer
        return st;
    }
};
```

### Java

```java
class Solution {
    public String smallestSubsequence(String s) {

        // Count remaining occurrences of every character
        int[] freq = new int[26];
        for (char c : s.toCharArray()) {
            freq[c - 'a']++;
        }

        // Check whether a character is already inside the stack
        boolean[] inStack = new boolean[26];

        // StringBuilder is used as a stack
        StringBuilder stack = new StringBuilder();

        for (char c : s.toCharArray()) {

            // Current occurrence has been processed
            freq[c - 'a']--;

            // Skip duplicate characters
            if (inStack[c - 'a']) {
                continue;
            }

            // Remove larger characters that will appear again
            while (stack.length() > 0 &&
                   stack.charAt(stack.length() - 1) > c &&
                   freq[stack.charAt(stack.length() - 1) - 'a'] > 0) {

                inStack[stack.charAt(stack.length() - 1) - 'a'] = false;
                stack.deleteCharAt(stack.length() - 1);
            }

            // Add current character
            stack.append(c);
            inStack[c - 'a'] = true;
        }

        return stack.toString();
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @return {string}
 */
var smallestSubsequence = function(s) {

    // Count remaining occurrences
    const freq = new Array(26).fill(0);
    for (const ch of s) {
        freq[ch.charCodeAt(0) - 97]++;
    }

    // Track characters already inside the stack
    const inStack = new Array(26).fill(false);

    // Array works as a stack
    const stack = [];

    for (const ch of s) {

        // Current occurrence has been processed
        freq[ch.charCodeAt(0) - 97]--;

        // Skip duplicates
        if (inStack[ch.charCodeAt(0) - 97]) {
            continue;
        }

        // Remove larger characters if they appear again later
        while (
            stack.length > 0 &&
            stack[stack.length - 1] > ch &&
            freq[stack[stack.length - 1].charCodeAt(0) - 97] > 0
        ) {
            inStack[stack.pop().charCodeAt(0) - 97] = false;
        }

        // Push current character
        stack.push(ch);
        inStack[ch.charCodeAt(0) - 97] = true;
    }

    // Join stack into the answer
    return stack.join("");
};
```

### Python3

```python
class Solution:
    def smallestSubsequence(self, s: str) -> str:

        # Count remaining occurrences of every character
        freq = [0] * 26
        for ch in s:
            freq[ord(ch) - ord('a')] += 1

        # Track characters already inside the stack
        in_stack = [False] * 26

        # Stack to build the answer
        stack = []

        for ch in s:

            # Current occurrence has been processed
            freq[ord(ch) - ord('a')] -= 1

            # Skip duplicate characters
            if in_stack[ord(ch) - ord('a')]:
                continue

            # Remove larger characters that appear again later
            while (
                stack and
                stack[-1] > ch and
                freq[ord(stack[-1]) - ord('a')] > 0
            ):
                in_stack[ord(stack.pop()) - ord('a')] = False

            # Add current character
            stack.append(ch)
            in_stack[ord(ch) - ord('a')] = True

        # Convert stack into a string
        return "".join(stack)
```

### Go

```go
func smallestSubsequence(s string) string {

 // Count remaining occurrences of every character
 freq := make([]int, 26)
 for _, ch := range s {
  freq[ch-'a']++
 }

 // Track whether a character is already inside the stack
 inStack := make([]bool, 26)

 // Stack to build the answer
 stack := make([]byte, 0)

 for i := 0; i < len(s); i++ {

  ch := s[i]

  // Current occurrence has been processed
  freq[ch-'a']--

  // Skip duplicate characters
  if inStack[ch-'a'] {
   continue
  }

  // Remove larger characters if they appear again later
  for len(stack) > 0 &&
   stack[len(stack)-1] > ch &&
   freq[stack[len(stack)-1]-'a'] > 0 {

   inStack[stack[len(stack)-1]-'a'] = false
   stack = stack[:len(stack)-1]
  }

  // Push current character
  stack = append(stack, ch)
  inStack[ch-'a'] = true
 }

 // Convert stack into the final answer
 return string(stack)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical in every language. Only the syntax changes.

First, count how many times every character appears in the string.

This frequency information tells us whether a character can safely be removed from the current answer. If it appears again later, removing it now is completely safe.

Next, create a stack that stores the current subsequence.

Also maintain a visited array. This prevents duplicate characters from being added to the stack because every distinct character must appear exactly once.

Now process the string from left to right.

For every character:

- Reduce its remaining frequency because the current occurrence has now been processed.
- If the character is already present inside the stack, ignore it.
- Otherwise, compare it with the top of the stack.

If the stack's top character is larger than the current character and that larger character still appears later, remove it.

This makes room for the smaller character to appear earlier, which improves the lexicographical order.

Repeat this until the top of the stack is no longer removable.

Finally, insert the current character into the stack and mark it as visited.

Once every character has been processed, the stack already contains the smallest subsequence with all distinct characters.

Since each character enters the stack once and leaves at most once, the algorithm remains linear.

## Examples

### Example 1

**Input**

```text
s = "bcabc"
```

**Output**

```text
"abc"
```

**Trace**

- Read `b`
- Read `c`
- Read `a`
- Remove `c`
- Remove `b`
- Add `a`
- Add `b`
- Add `c`

Final answer:

```text
abc
```

---

### Example 2

**Input**

```text
s = "cbacdcbc"
```

**Output**

```text
"acdb"
```

**Trace**

- Add `c`
- Replace it when smaller characters can safely move ahead
- Keep characters that will not appear again
- Continue updating the stack greedily

Final answer:

```text
acdb
```

---

### Example 3

**Input**

```text
s = "cdadabcc"
```

**Output**

```text
"adbc"
```

**Trace**

The algorithm keeps replacing larger removable characters with smaller ones while preserving every distinct character exactly once.

Final answer:

```text
adbc
```

## How to Use / Run Locally

Clone the repository.

```bash
git clone <repository-url>
cd <repository-folder>
```

### C++

Compile

```bash
g++ solution.cpp -o solution
```

Run

```bash
./solution
```

### Java

Compile

```bash
javac Solution.java
```

Run

```bash
java Solution
```

### JavaScript

Run

```bash
node solution.js
```

### Python3

Run

```bash
python solution.py
```

or

```bash
python3 solution.py
```

### Go

Run

```bash
go run solution.go
```

## Notes & Optimizations

- This problem is the same as LeetCode 316 (Remove Duplicate Letters).
- A greedy algorithm works because every decision is made only when it is guaranteed to improve the final answer.
- The stack never stores duplicate characters.
- Characters that never appear again are never removed.
- Since only lowercase English letters are allowed, the extra memory remains constant.
- A brute-force approach would be far too slow because it would generate many possible subsequences.
- The greedy monotonic stack solution is the standard interview solution and achieves the best possible time complexity.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
