# 3302. Find the Lexicographically Smallest Valid Sequence

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

LeetCode 3302, **Find the Lexicographically Smallest Valid Sequence**, gives us two strings, `word1` and `word2`.

I need to select exactly `word2.length` indices from `word1`. The selected indices must be in increasing order.

When I take the characters from `word1` using those indices, the resulting string must be **almost equal** to `word2`. This means I can change at most one character in the selected string to make it exactly equal to `word2`.

The goal is not to find the lexicographically smallest resulting string. I need to find the **lexicographically smallest array of indices**.

If no valid sequence of indices exists, I return an empty array.

For example:

```text
word1 = "v b c c a"
word2 = "a b c"
```

A valid answer can be:

```text
[1, 2, 4]
```

The selected characters are `"acc"`, which can become `"abc"` by changing one character.

The main challenge is deciding when I can safely use the one allowed mismatch while still completing the rest of `word2`.

## Constraints

* `1 <= word2.length <= word1.length <= 3 * 10^5`
* `word1` contains only lowercase English letters.
* `word2` contains only lowercase English letters.
* The answer must contain exactly `word2.length` indices.
* The indices must be strictly increasing.
* The selected characters can differ from `word2` at most once.

## Intuition

My first thought is to scan `word1` from left to right and always pick the earliest possible index.

That makes sense because the answer itself is an array of indices. If I can start the answer with index `2` instead of `3`, then `[2, ...]` is always lexicographically smaller than `[3, ...]`.

The problem is that I cannot blindly choose every character.

I am allowed only one mismatch.

So when `word1[i]` does not match the current character of `word2`, I need to know whether using the mismatch at index `i` will still leave enough room to match all remaining characters exactly.

To answer that efficiently, I first scan from right to left.

During this backward scan, I find positions that can match the suffixes of `word2`. Then, while scanning from left to right, I can quickly check whether using the current index as the one mismatch is safe.

This gives me a greedy solution with linear time complexity.

## Approach

I use two passes over `word1`.

### 1. Build suffix matching information

I start from the end of both strings.

I try to match `word2` from right to left using characters from `word1`.

For every position `j` in `word2`, I store a position in `word1` where the suffix starting at `j` can be matched.

I call this array `last`.

The important idea is that `last[j + 1]` tells me where I can start matching the remaining part of `word2` after the current character.

### 2. Greedily build the answer

Now I scan `word1` from left to right.

For every position `i`:

* If `word1[i] == word2[j]`, I take index `i`.
* Otherwise, I try to use the one allowed mismatch.
* I use that mismatch only if the remaining suffix of `word2` can still be matched.

The safety condition is:

```text
j is the last character
OR
i < last[j + 1]
```

If this condition is true, I can safely use `i` as the mismatching position.

### 3. Check whether the complete sequence was built

After the scan, I check whether all characters of `word2` were matched.

If yes, I return the answer.

Otherwise, I return an empty array.

Because I always scan from left to right and take the earliest safe index, the resulting index sequence is lexicographically smallest.

## Data Structures Used

### `last` array

I use an integer array of size `word2.length`.

`last[j]` stores a position in `word1` that can be used while matching the suffix of `word2` starting at index `j`.

This lets me check whether the remaining characters can still be matched after using the one mismatch.

### Answer array

I store the selected indices in an array or list.

The answer always contains exactly `word2.length` indices when a valid sequence exists.

No large dynamic programming table is required.

## Operations & Behavior Summary

The algorithm can be viewed as the following plain-English pseudocode:

1. Create a `last` array for `word2`.
2. Start at the end of both strings.
3. Match characters from right to left.
4. Store the matching positions in `last`.
5. Start scanning `word1` from the beginning.
6. Keep track of the current character needed from `word2`.
7. If the characters match, select the current index.
8. If they do not match, check whether the one allowed mismatch can be used safely.
9. If the mismatch is safe, select the current index and mark the mismatch as used.
10. Continue until all characters of `word2` are selected or `word1` ends.
11. Return the selected indices if their count equals `word2.length`.
12. Otherwise, return an empty array.

## Complexity

| Metric           | Complexity | Explanation                                                                                                             |
| ---------------- | ---------- | ----------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n + m)` | `n` is the length of `word1` and `m` is the length of `word2`. I scan the strings from right to left and left to right. |
| Space Complexity | `O(m)`     | I use the `last` array of size `m`, along with the answer array containing at most `m` indices.                         |

Since `word1.length` can be as large as `3 * 10^5`, a linear solution is important.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> validSequence(string word1, string word2) {
        int n = word1.size();
        int m = word2.size();

        vector<int> last(m, -1);

        int i = n - 1;
        int j = m - 1;

        while (i >= 0 && j >= 0) {
            if (word1[i] == word2[j]) {
                last[j] = i;
                --j;
            }

            --i;
        }

        vector<int> ans;
        ans.reserve(m);

        bool canSkip = true;
        j = 0;

        for (i = 0; i < n && j < m; ++i) {
            if (word1[i] == word2[j]) {
                ans.push_back(i);
                ++j;
            }
            else if (canSkip &&
                     (j == m - 1 || i < last[j + 1])) {
                canSkip = false;
                ans.push_back(i);
                ++j;
            }
        }

        if (j == m) {
            return ans;
        }

        return {};
    }
};
```

### Java

```java
class Solution {
    public int[] validSequence(String word1, String word2) {
        int n = word1.length();
        int m = word2.length();

        int[] last = new int[m];

        java.util.Arrays.fill(last, -1);

        int i = n - 1;
        int j = m - 1;

        while (i >= 0 && j >= 0) {
            if (word1.charAt(i) == word2.charAt(j)) {
                last[j] = i;
                --j;
            }

            --i;
        }

        int[] ans = new int[m];
        int size = 0;

        boolean canSkip = true;
        j = 0;

        for (i = 0; i < n && j < m; ++i) {
            if (word1.charAt(i) == word2.charAt(j)) {
                ans[size++] = i;
                ++j;
            }
            else if (canSkip &&
                     (j == m - 1 || i < last[j + 1])) {
                canSkip = false;
                ans[size++] = i;
                ++j;
            }
        }

        if (j == m) {
            return ans;
        }

        return new int[0];
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} word1
 * @param {string} word2
 * @return {number[]}
 */
var validSequence = function(word1, word2) {
    const n = word1.length;
    const m = word2.length;

    const last = new Array(m).fill(-1);

    let i = n - 1;
    let j = m - 1;

    while (i >= 0 && j >= 0) {
        if (word1[i] === word2[j]) {
            last[j] = i;
            --j;
        }

        --i;
    }

    const ans = [];
    let canSkip = true;
    j = 0;

    for (i = 0; i < n && j < m; ++i) {
        if (word1[i] === word2[j]) {
            ans.push(i);
            ++j;
        }
        else if (canSkip && (j === m - 1 || i < last[j + 1])) {
            canSkip = false;
            ans.push(i);
            ++j;
        }
    }

    if (j === m) {
        return ans;
    }

    return [];
};
```

### Python3

```python
from typing import List

class Solution:
    def validSequence(self, word1: str, word2: str) -> List[int]:
        n = len(word1)
        m = len(word2)

        last = [-1] * m

        i = n - 1
        j = m - 1

        while i >= 0 and j >= 0:
            if word1[i] == word2[j]:
                last[j] = i
                j -= 1

            i -= 1

        ans = []
        can_skip = True
        j = 0

        for i in range(n):
            if j == m:
                break

            if word1[i] == word2[j]:
                ans.append(i)
                j += 1

            elif can_skip and (j == m - 1 or i < last[j + 1]):
                can_skip = False
                ans.append(i)
                j += 1

        if j == m:
            return ans

        return []
```

### Go

```go
func validSequence(word1 string, word2 string) []int {
 n := len(word1)
 m := len(word2)

 last := make([]int, m)

 for i := 0; i < m; i++ {
  last[i] = -1
 }

 i := n - 1
 j := m - 1

 for i >= 0 && j >= 0 {
  if word1[i] == word2[j] {
   last[j] = i
   j--
  }

  i--
 }

 ans := make([]int, 0, m)
 canSkip := true
 j = 0

 for i := 0; i < n && j < m; i++ {
  if word1[i] == word2[j] {
   ans = append(ans, i)
   j++
  } else if canSkip && (j == m-1 || i < last[j+1]) {
   canSkip = false
   ans = append(ans, i)
   j++
  }
 }

 if j == m {
  return ans
 }

 return []int{}
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core algorithm is the same in all five languages. Only the syntax and standard library operations change.

### Step 1: Store the string lengths

I first get the lengths of `word1` and `word2`.

Let:

```text
n = word1.length
m = word2.length
```

I need `m` indices in the final answer.

### Step 2: Create the suffix information array

I create an array called `last` with `m` positions.

Initially, every position represents that I have not found a matching position yet.

The purpose of this array is not to store the complete solution. It only tells me where the remaining suffix can be matched.

This is enough for the greedy decision later.

### Step 3: Scan from right to left

I set one pointer to the end of `word1` and another to the end of `word2`.

Whenever the characters match, I store the current position in `last`.

Then I move the `word2` pointer backward.

I continue until I reach the beginning of either string.

This creates the information I need about the suffixes.

### Step 4: Start the greedy scan

After building `last`, I reset the `word2` pointer to `0`.

Now I scan `word1` from left to right.

This direction is important.

Since I want the lexicographically smallest array of indices, I want to choose the smallest possible index at every position.

### Step 5: Handle an exact match

If:

```text
word1[i] == word2[j]
```

I select `i`.

There is no reason to skip this index.

Taking an earlier matching index gives me more space for the remaining characters and also makes the answer lexicographically smaller.

Then I move to the next character of `word2`.

### Step 6: Handle a mismatch

If the characters are different, I have two possibilities.

I can ignore this index and continue searching.

Or I can use this index as my one allowed mismatch.

I choose the second option only if it is safe.

For the last character of `word2`, it is automatically safe because there are no characters remaining.

For other positions, I check whether:

```text
i < last[j + 1]
```

This means the suffix beginning at `word2[j + 1]` can still be matched after the current index.

If this condition is false, using the mismatch now could make the rest of the sequence impossible.

So I skip the current index.

### Step 7: Mark the mismatch as used

Once I use a mismatch, I cannot use another one.

I keep a boolean flag such as `canSkip` to remember this.

After using the mismatch, I set it to false.

From that point onward, only exact character matches are accepted.

### Step 8: Check the final result

When the scan ends, I check whether I successfully selected `m` indices.

If I did, the answer is valid.

If I did not, there is no valid sequence that satisfies the problem conditions.

I return an empty array in that case.

### C++ behavior

The C++ implementation uses `vector<int>` for the suffix positions and the result.

String characters can be accessed directly using `word1[i]` and `word2[j]`.

The resulting solution works comfortably within the constraints because it uses linear scans.

### Java behavior

The Java implementation uses `int[]` for the suffix information and the final result.

Characters are accessed using `charAt()`.

An integer counter is used while filling the answer array so that the final result contains exactly the required number of indices.

### JavaScript behavior

The JavaScript implementation uses normal arrays.

`word1[i]` and `word2[j]` directly access individual characters.

The answer is built using `push()`.

The algorithm does not use expensive nested loops, so it remains linear.

### Python3 behavior

The Python implementation uses a normal list for `last` and another list for the answer.

Python strings support direct indexing, so `word1[i]` and `word2[j]` are enough to compare characters.

The solution avoids nested scanning and therefore remains `O(n + m)`.

### Go behavior

The Go implementation uses slices for the `last` array and the answer.

Because the input contains lowercase English letters, indexing the strings gives the required byte values directly.

The answer slice is created with enough capacity for `word2.length` elements, avoiding unnecessary reallocations.

## Examples

### Example 1

**Input**

```text
word1 = "v b c c a"
word2 = "a b c"
```

Using normal string notation:

```text
word1 = "vbcca"
word2 = "abc"
```

**Output**

```text
[0, 1, 2]
```

The selected characters are:

```text
v b c
```

I compare them with:

```text
a b c
```

Only the first character is different, so I can change `v` to `a`.

The answer starts at index `0`, which makes it lexicographically smaller than any valid answer starting at a later index.

### Example 2

**Input**

```text
word1 = "bacdc"
word2 = "abc"
```

**Output**

```text
[1, 2, 4]
```

I first check index `0`.

```text
word1[0] = b
word2[0] = a
```

They are different.

Although I have one mismatch available, using it immediately would prevent me from matching the remaining characters correctly.

So I skip index `0`.

At index `1`:

```text
word1[1] = a
word2[0] = a
```

I take index `1`.

At index `2`:

```text
word1[2] = c
word2[1] = b
```

This is a safe mismatch because the remaining `c` can still be matched at index `4`.

So I take index `2`.

Finally, index `4` matches `c`.

The result is:

```text
[1, 2, 4]
```

### Example 3

**Input**

```text
word1 = "aaaaaa"
word2 = "aaabc"
```

**Output**

```text
[]
```

I need:

```text
a a a b c
```

from `word1`, but `word1` contains only `a`.

I can change at most one character.

There is no way to create both the required `b` and `c`.

Therefore, no valid sequence exists, so the answer is an empty array.

## How to Use / Run Locally

The solution can be tested locally in any of the five supported languages.

### C++

Save the solution in a file such as:

```text
solution.cpp
```

Compile it with:

```bash
g++ -std=c++17 -O2 solution.cpp -o solution
```

Then run:

```bash
./solution
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

Then run:

```bash
java Solution
```

### JavaScript

Save the solution as:

```text
solution.js
```

Run it with Node.js:

```bash
node solution.js
```

### Python3

Save the solution as:

```text
solution.py
```

Run:

```bash
python3 solution.py
```

### Go

Save the solution as:

```text
solution.go
```

Run:

```bash
go run solution.go
```

For local testing, I would add a small driver program for the language I am using and call the `validSequence` function with the sample inputs.

On LeetCode, I only need to submit the required `Solution` class or function because the platform provides the input handling and test runner.

## Notes & Optimizations

The most important detail is that the answer is judged by the **indices**, not by the string created from those indices.

That is why the left-to-right greedy scan is necessary.

I also need to be careful with the one allowed mismatch. I cannot simply take the first mismatching character I see.

The mismatch is valid only when the remaining suffix of `word2` can still be matched exactly.

The right-to-left `last` array makes this check possible in constant time.

A brute-force solution would try many possible index sequences and would be far too slow for `word1.length` up to `3 * 10^5`.

A dynamic programming solution with a large two-dimensional state would also use too much memory.

The suffix matching array reduces the required extra information to `O(m)`.

Another useful observation is that I do not need to construct the resulting string at all. The problem only asks for the index sequence, so I directly build and return the indices.

The edge cases I need to consider include:

* `word1` and `word2` have the same length.
* The entire match requires exactly one mismatch.
* The entire match is already exact.
* The mismatch occurs at the first character.
* The mismatch occurs at the last character.
* No valid sequence exists.
* There are many repeated characters, so choosing the earliest index matters.

The final approach is a greedy left-to-right construction supported by a right-to-left suffix matching pass. It runs in `O(n + m)` time and uses `O(m)` extra space.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/?utm_source=chatgpt.com)
