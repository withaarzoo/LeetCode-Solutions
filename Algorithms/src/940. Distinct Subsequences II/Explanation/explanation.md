# 940. Distinct Subsequences II — LeetCode Solution

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)

  * [C++](#c)
  * [Java](#java)
  * [JavaScript](#javascript)
  * [Python3](#python3)
  * [Go](#go)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

LeetCode 940, **Distinct Subsequences II**, asks us to find the number of distinct non-empty subsequences that can be formed from a string `s`.

A subsequence is created by deleting zero or more characters while keeping the relative order of the remaining characters.

For example, from `"abc"`, we can create:

```text
"a", "b", "c", "ab", "ac", "bc", "abc"
```

So the answer is `7`.

The main difficulty is handling duplicate subsequences when the same character appears multiple times.

For example, `"aba"` can create `"a"` using either occurrence of `'a'`, but `"a"` should only be counted once.

The answer can also become very large, so the result must be returned modulo:

```text
10^9 + 7
```

This repository explains an efficient **dynamic programming solution for LeetCode 940** with implementations in C++, Java, JavaScript, Python3, and Go.

## Constraints

| Constraint    | Value                     |
| ------------- | ------------------------- |
| String length | `1 <= s.length <= 2000`   |
| Characters    | Lowercase English letters |
| Modulo        | `10^9 + 7`                |

## Intuition

My first thought was that every new character gives me two choices for every existing subsequence:

```text
Do not take the character
        OR
Take the character
```

So if I currently have `dp` distinct subsequences, adding a character seems to give:

```text
2 * dp
```

But this counts duplicates when the character has appeared before.

For example, after processing `"ab"`, I have:

```text
""
"a"
"b"
"ab"
```

If I add another `'a'`, I can create:

```text
"a"
"aa"
"ba"
"aba"
```

But `"a"` already existed.

So I need a way to remember how many subsequences existed before the previous occurrence of each character.

That leads to the key idea:

```text
new dp = 2 * old dp - previous contribution of this character
```

I store that previous contribution in an array of size `26`.

## Approach

I start with:

```text
dp = 1
```

The `1` represents the empty subsequence.

For every character `c` in the string, I do three things.

First, I calculate the new number of subsequences:

```text
new dp = 2 * dp - last[c]
```

The `2 * dp` comes from keeping the old subsequences and creating new ones by appending the current character.

Second, if this character appeared before, `last[c]` removes the subsequences that would be counted again.

Third, I update:

```text
last[c] = old dp
```

This makes the value ready for the next occurrence of the same character.

For `"aba"`:

```text
Character     Old dp     Removed     New dp
--------------------------------------------
'a'              1          0           2
'b'              2          0           4
'a'              4          1           7
```

After processing the entire string, `dp` still includes the empty subsequence.

So I return:

```text
dp - 1
```

For `"aba"`:

```text
7 - 1 = 6
```

The six non-empty distinct subsequences are:

```text
"a", "b", "ab", "aa", "ba", "aba"
```

## Data Structures Used

### `dp`

I use one integer variable to store the current number of distinct subsequences, including the empty subsequence.

### `last[26]`

I use an array of 26 values because the string contains only lowercase English letters.

Each position represents one character:

```text
'a' -> 0
'b' -> 1
'c' -> 2
...
'z' -> 25
```

`last[c]` stores the value of `dp` from before the previous occurrence of character `c`.

No set or list of actual subsequences is needed. This keeps the solution efficient.

## Operations & Behavior Summary

The algorithm can be viewed as the following pseudocode:

```text
Start with:
dp = 1

For every character c:

    Save the current dp

    Double dp
    This represents keeping every old subsequence
    and also creating a version with c appended

    Subtract last[c]
    This removes duplicates caused by the previous occurrence of c

    Store the old dp in last[c]

Return dp - 1
```

The subtraction is the important part that changes this from counting all subsequences to counting only distinct subsequences.

## Complexity

| Complexity       | Value  | Explanation                                                          |
| ---------------- | ------ | -------------------------------------------------------------------- |
| Time Complexity  | `O(n)` | I process each of the `n` characters exactly once.                   |
| Space Complexity | `O(1)` | I use an array of only 26 elements, regardless of the string length. |

Here, `n` is the length of the input string.

Because the alphabet contains only 26 lowercase English letters, the extra space is constant.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int distinctSubseqII(string s) {
        const long long MOD = 1000000007; // I use this modulo because the answer can be very large.
        
        long long dp = 1; // I start with 1 because the empty subsequence is initially the only subsequence.
        vector<long long> last(26, 0); // last[c] stores dp from before the previous occurrence of character c.
        
        for (char c : s) { // I process every character once.
            int index = c - 'a'; // I convert the character into an index from 0 to 25.
            
            long long oldDp = dp; // I save the old count because last[index] must be updated with this value.
            
            dp = (2 * dp - last[index] + MOD) % MOD; // I double the subsequences and remove duplicates made by c.
            
            last[index] = oldDp; // I remember the old count for the next occurrence of this character.
        }
        
        return (dp - 1 + MOD) % MOD; // I remove the empty subsequence because the problem asks for non-empty ones.
    }
};
```

### Java

```java
class Solution {
    public int distinctSubseqII(String s) {
        final long MOD = 1000000007L; // I use this modulo because the answer can be very large.
        
        long dp = 1; // I start with the empty subsequence as the only subsequence.
        long[] last = new long[26]; // last[c] stores dp from before the previous occurrence of character c.
        
        for (char c : s.toCharArray()) { // I process every character exactly once.
            int index = c - 'a'; // I convert the character into an index from 0 to 25.
            
            long oldDp = dp; // I save the old count before changing dp.
            
            dp = (2 * dp - last[index] + MOD) % MOD; // I add both choices and remove duplicate subsequences.
            
            last[index] = oldDp; // I store the old count for the next occurrence of this character.
        }
        
        return (int)((dp - 1 + MOD) % MOD); // I remove the empty subsequence from the final answer.
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @return {number}
 */
var distinctSubseqII = function(s) {
    const MOD = 1000000007; // I use this modulo because the answer can be very large.
    
    let dp = 1; // I start with the empty subsequence.
    const last = new Array(26).fill(0); // last[c] stores dp from before the previous occurrence of c.
    
    for (const ch of s) { // I process every character exactly once.
        const index = ch.charCodeAt(0) - 97; // I convert the lowercase character into an index from 0 to 25.
        
        const oldDp = dp; // I save the old count because it is needed for last[index].
        
        dp = (2 * dp - last[index] + MOD) % MOD; // I double the count and remove duplicate subsequences.
        
        last[index] = oldDp; // I remember the old count for future occurrences of this character.
    }
    
    return (dp - 1 + MOD) % MOD; // I remove the empty subsequence from the answer.
};
```

### Python3

```python
class Solution:
    def distinctSubseqII(self, s: str) -> int:
        MOD = 1000000007  # I use this modulo because the answer can be very large.
        
        dp = 1  # I start with the empty subsequence as the only subsequence.
        last = [0] * 26  # last[c] stores dp from before the previous occurrence of character c.
        
        for ch in s:  # I process every character exactly once.
            index = ord(ch) - ord('a')  # I convert the lowercase character into an index from 0 to 25.
            
            old_dp = dp  # I save the old count before changing dp.
            
            dp = (2 * dp - last[index] + MOD) % MOD  # I double the count and remove duplicate subsequences.
            
            last[index] = old_dp  # I store the old count for the next occurrence of this character.
        
        return (dp - 1 + MOD) % MOD  # I remove the empty subsequence from the final answer.
```

### Go

```go
func distinctSubseqII(s string) int {
 const MOD int64 = 1000000007 // I use this modulo because the answer can be very large.

 var dp int64 = 1 // I start with the empty subsequence as the only subsequence.
 last := make([]int64, 26) // last[c] stores dp from before the previous occurrence of c.

 for i := 0; i < len(s); i++ { // I process every character exactly once.
  index := int(s[i] - 'a') // I convert the lowercase character into an index from 0 to 25.

  oldDp := dp // I save the old count before changing dp.

  dp = (2*dp - last[index] + MOD) % MOD // I double the count and remove duplicate subsequences.

  last[index] = oldDp // I remember the old count for future occurrences of this character.
 }

 return int((dp - 1 + MOD) % MOD) // I remove the empty subsequence from the final answer.
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five languages, so the important part is understanding the state we maintain.

### 1. Initialize the modulo

I use:

```text
MOD = 1000000007
```

The number of subsequences can grow extremely quickly. Taking the modulo after every operation prevents the numbers from becoming too large.

The modulo is applied throughout the calculation instead of waiting until the end.

### 2. Start `dp` with 1

I initialize:

```text
dp = 1
```

This may look slightly unusual because the problem asks for non-empty subsequences.

I include the empty subsequence during the calculation because it makes the transition much cleaner.

Initially, there is exactly one subsequence:

```text
""
```

At the end, I remove it.

### 3. Create the `last` array

I create an array with 26 positions:

```text
last[0 ... 25]
```

Every value starts at `0`.

When I process a character for the first time, there is no previous occurrence, so there is nothing to subtract.

### 4. Convert the character into an index

For a lowercase character, I convert it to a number from `0` to `25`.

For example:

```text
'a' -> 0
'b' -> 1
'z' -> 25
```

This lets me access the correct value in `last`.

The exact syntax is different in each language, but the idea stays the same.

### 5. Save the old `dp`

Before changing `dp`, I save its current value.

I need this because the old value is what gets stored in `last[c]`.

For example:

```text
old dp = 4
```

After calculating the new value, I store:

```text
last[c] = 4
```

### 6. Double the number of subsequences

Suppose I currently have:

```text
dp = 4
```

When I process a new character, every existing subsequence has two possibilities:

```text
Existing subsequence
       /       \
   skip c     take c
```

So I initially get:

```text
4 * 2 = 8
```

This is where:

```text
2 * dp
```

comes from.

### 7. Remove duplicates

If the current character appeared before, some of these 8 subsequences may already exist.

The value stored in:

```text
last[c]
```

tells me exactly how many duplicate subsequences were introduced by this repeated character.

So I calculate:

```text
dp = 2 * dp - last[c]
```

I also add `MOD` before taking the modulo so that the intermediate value stays non-negative.

### 8. Update the character's previous contribution

After calculating the new `dp`, I store the old value:

```text
last[c] = old dp
```

This is important because the next occurrence of the same character needs this information.

For `"aaa"`, the state looks like:

```text
Character     Old dp     last[a] used     New dp
-------------------------------------------------
'a'              1             0             2
'a'              2             1             3
'a'              3             2             4
```

At the end:

```text
dp = 4
```

Removing the empty subsequence gives:

```text
4 - 1 = 3
```

The distinct non-empty subsequences are:

```text
"a"
"aa"
"aaa"
```

### 9. Why the algorithm works in all five languages

C++, Java, JavaScript, Python3, and Go all implement the same state transition:

```text
dp = 2 * dp - last[c]
```

The main differences are only language-specific details such as arrays, character-to-index conversion, integer types, and modulo operations.

For C++ and Java, I use a wide integer type during calculations because the multiplication can temporarily exceed a normal 32-bit integer.

For JavaScript, the values stay safely within the range needed by this calculation.

For Python3, integers automatically support large values.

For Go, I use `int64` for the arithmetic before converting the final result to `int`.

### 10. Remove the empty subsequence

The DP deliberately counts the empty subsequence.

The problem does not want it, so the final operation is:

```text
answer = dp - 1
```

The modulo is applied here as well to keep the result valid.

## Examples

### Example 1

Input:

```text
s = "abc"
```

Processing:

```text
Start: dp = 1

'a':
dp = 2

'b':
dp = 4

'c':
dp = 8
```

There are no repeated characters, so nothing needs to be removed.

The final answer is:

```text
8 - 1 = 7
```

Output:

```text
7
```

The distinct non-empty subsequences are:

```text
"a", "b", "c", "ab", "ac", "bc", "abc"
```

### Example 2

Input:

```text
s = "aba"
```

State changes:

```text
Start:
dp = 1

After 'a':
dp = 2
last[a] = 1

After 'b':
dp = 4
last[b] = 2

After second 'a':
dp = 2 * 4 - 1
dp = 7
```

The `1` removes the duplicate subsequence `"a"`.

Finally:

```text
7 - 1 = 6
```

Output:

```text
6
```

### Example 3

Input:

```text
s = "aaa"
```

State changes:

```text
Start:
dp = 1

First 'a':
dp = 2

Second 'a':
dp = 3

Third 'a':
dp = 4
```

After removing the empty subsequence:

```text
4 - 1 = 3
```

Output:

```text
3
```

The distinct subsequences are:

```text
"a", "aa", "aaa"
```

## How to Use / Run Locally

The repository contains the same algorithm implemented in five programming languages.

### C++

Save the solution in a file such as:

```text
solution.cpp
```

Compile it with:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it with:

```bash
./solution
```

On Windows, you can run:

```bash
solution.exe
```

### Java

Save the solution in:

```text
Solution.java
```

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

When using the code directly on LeetCode, the `Solution` class can be submitted without adding your own `main` method.

### JavaScript

Save the solution in:

```text
solution.js
```

Run it using Node.js:

```bash
node solution.js
```

You can add a small test case at the bottom of the file when testing locally.

### Python3

Save the solution in:

```text
solution.py
```

Run it with:

```bash
python3 solution.py
```

On some Windows installations, the command may be:

```bash
python solution.py
```

### Go

Save the solution in:

```text
solution.go
```

Run it with:

```bash
go run solution.go
```

For a compiled executable, use:

```bash
go build solution.go
```

## Notes & Optimizations

The most important optimization is avoiding the generation of actual subsequences.

A brute-force solution could generate every possible subsequence, but there can be up to `2^n` subsequences. With `n` up to `2000`, that approach is not practical.

I also do not use a `set` to store every generated string. Although a set can help remove duplicates, storing all subsequences would require far too much memory.

Instead, I only keep:

```text
dp
```

and:

```text
last[26]
```

This reduces the solution to `O(n)` time and `O(1)` extra space.

The empty subsequence is counted internally because it makes the DP transition simple. I remove it only once at the end.

Another important detail is that I update `last[c]` after calculating the new `dp`. If I updated it before the calculation, I would lose the previous value that is needed to remove duplicates.

The solution works especially well here because the input contains only lowercase English letters. If the character set were much larger, I could replace the fixed 26-element array with another suitable mapping structure.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)

---
