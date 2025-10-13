# 2273 - Find Resultant Array After Removing Anagrams

---

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
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

We are given a 0-indexed array of lowercase strings `words`. We repeatedly perform the following: choose an index `i` with `0 < i < words.length` such that `words[i]` and `words[i-1]` are anagrams, and delete `words[i]`. Keep doing this until no more deletions are possible. Return the resulting `words` array.

Important property: the final array is unique regardless of the order we choose deletions. So a single left-to-right pass comparing each word with the last *kept* word suffices.

---

## Constraints

* `1 <= words.length <= 100`
* `1 <= words[i].length <= 10`
* `words[i]` consists of lowercase English letters.

Because word length ≤ 10, sorting a word is cheap (worst-case 10 log 10).

---

## Intuition

I thought: two strings are anagrams if they contain the same letters with the same counts. So I can convert each word into a canonical **signature** which is identical for all anagrams. If I walk left-to-right and keep only the first word of any group of adjacent anagrams (i.e., compare each word’s signature to the signature of the previously kept word), I’ll obtain the final array directly.

A signature can be:

* The sorted characters of the word (easy and efficient for small word length).
* Or a 26-length frequency signature (deterministic O(m) creation), useful if words were long.

---

## Approach

1. Initialize empty `result` list and `prevSig` as empty string.
2. For each word `w` in `words`:

   * Build `sig` = signature of `w` (sorted characters or frequency string).
   * If `sig != prevSig`:

     * Append `w` to `result`.
     * Set `prevSig = sig`.
   * Else (sig equals prevSig) → skip `w` (it's an anagram of last kept word).
3. Return `result`.

This is a single-pass greedy scan that keeps one representative per run of adjacent anagrams.

---

## Data Structures Used

* Result vector/list to store kept words.
* A `string` to store the signature of the last kept word.
* Temporary buffer (array or sorted string) to create signatures.

---

## Operations & Behavior Summary

* Per-word operation: create signature (sort or count letters), compare with `prevSig`.
* Append word to result only if not anagram of previous kept word.
* Behavior: collapse any run of adjacent anagrams into the first element of that run.

---

## Complexity

* **Time Complexity:**

  * Using sorted-character signature: `O(n * m log m)` where `n` is number of words and `m` is maximum word length (here `m ≤ 10`, so sorting cost is negligible).
  * Using frequency signature: `O(n * m)` (best asymptotic if `m` large).
* **Space Complexity:**

  * `O(n * m)` for output (worst-case all words kept).
  * Additional `O(m)` temporary space for signature creation.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    vector<string> removeAnagrams(vector<string>& words) {
        vector<string> result;
        string prevSig = ""; // signature of last kept word

        for (auto &w : words) {
            string sig = w;
            sort(sig.begin(), sig.end()); // signature: sorted characters
            if (sig != prevSig) {
                result.push_back(w);      // keep this word
                prevSig = move(sig);      // update prevSig
            }
        }
        return result;
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {
    public List<String> removeAnagrams(String[] words) {
        List<String> result = new ArrayList<>();
        String prevSig = "";

        for (String w : words) {
            char[] arr = w.toCharArray();
            Arrays.sort(arr);               // signature: sorted chars
            String sig = new String(arr);
            if (!sig.equals(prevSig)) {
                result.add(w);
                prevSig = sig;
            }
        }
        return result;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {string[]} words
 * @return {string[]}
 */
var removeAnagrams = function(words) {
    const result = [];
    let prevSig = "";

    for (const w of words) {
        const sig = w.split('').sort().join(''); // signature by sorted characters
        if (sig !== prevSig) {
            result.push(w);
            prevSig = sig;
        }
    }
    return result;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def removeAnagrams(self, words: List[str]) -> List[str]:
        result = []
        prev_sig = ""
        for w in words:
            sig = ''.join(sorted(w))   # signature by sorted characters
            if sig != prev_sig:
                result.append(w)
                prev_sig = sig
        return result
```

---

### Go

```go
package main

import (
	"sort"
)

// helper: returns sorted-character signature of s
func sortedSig(s string) string {
    b := []byte(s)
    sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
    return string(b)
}

func removeAnagrams(words []string) []string {
    var result []string
    prevSig := ""
    for _, w := range words {
        sig := sortedSig(w)
        if sig != prevSig {
            result = append(result, w)
            prevSig = sig
        }
    }
    return result
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Below I explain the core algorithm line-by-line in a language-agnostic manner, and then add small language-specific notes.

### Core algorithm (common steps)

1. `result = []` — prepare an empty container to store final kept words.
2. `prevSig = ""` — no previous signature at start (empty).
3. For each word `w` in `words`:

   * Create `sig`:

     * Sorted-signature: `sig = sorted characters of w` (e.g., `"baba"` → `"aabb"`).
     * Frequency-signature (alternative): encode counts of 26 letters into a string like `"1#0#2#..."`.
   * `if sig != prevSig`:

     * Append `w` to `result` because it is not an anagram of the previously kept word.
     * Set `prevSig = sig`.
   * Otherwise skip `w`.
4. Return `result`.

This simple comparison uses the fact that we only remove a word if it is an anagram of the previous kept word; runs of adjacent anagrams collapse to their first element.

---

### C++ specific notes

* `move(sig)` after pushing `w` into result avoids copying `sig` when assigning to `prevSig`. `sort(sig.begin(), sig.end())` is in-place and fast for small strings.

### Java specific notes

* `char[] arr = w.toCharArray(); Arrays.sort(arr); String sig = new String(arr)` — converting to `char[]` and sorting is straightforward. String equality uses `.equals()`.

### JavaScript specific notes

* `w.split('').sort().join('')` creates a signature quickly. Arrays in JS sort lexicographically by default and for single-letter elements it is correct.

### Python specific notes

* `sig = ''.join(sorted(w))` is concise and readable. Python sorted returns a list of chars which we join.

### Go specific notes

* Convert string to byte slice `[]byte(s)` then use `sort.Slice` to sort bytes in-place; convert back to string.

---

## Examples

1. Example 1
   Input: `["abba","baba","bbaaa","cd","cd"]`
   (Note: example adjusted for clarity)
   Actually valid sample from problem:
   Input: `["abba","baba","bbaa","cd","cd"]`
   Process (left-to-right):

   * `"abba"` signature `"aabb"` → keep.
   * `"baba"` signature `"aabb"` = prevSig → skip.
   * `"bbaa"` signature `"aabb"` = prevSig → skip.
   * `"cd"` signature `"cd"` != prevSig `"aabb"` → keep.
   * Next `"cd"` signature `"cd"` = prevSig → skip.
     Output: `["abba","cd"]`

2. Example 2
   Input: `["a","b","c","d","e"]`
   No adjacent anagrams → Output: `["a","b","c","d","e"]`

---

## How to use / Run locally

### C++

* Put `Solution` class into a file, e.g., `solution.cpp`. Build a `main()` wrapper to call `removeAnagrams` and print results for testing. Compile with:

```bash
g++ -std=c++17 solution.cpp -O2 -o solution
./solution
```

### Java

* Put class `Solution` in `Solution.java`. Create a `main` method to test. Compile & run:

```bash
javac Solution.java
java Solution
```

### JavaScript (Node)

* Save the `removeAnagrams` function in `solution.js` and add test harness code that calls it and prints `console.log(...)`.

```bash
node solution.js
```

### Python3

* Save the class in `solution.py`. Add a test harness (construct `Solution` and call `removeAnagrams`) and run:

```bash
python3 solution.py
```

### Go

* Put functions in `main.go`, include a `main()` to test `removeAnagrams`, then:

```bash
go run main.go
```

---

## Notes & Optimizations

* For these constraints (`m ≤ 10`), sorting each word is simple, readable, and very fast.
* If words were much longer, consider building a 26-integer frequency signature and encoding it (string with separators) to avoid the `m log m` factor.
* Memory-wise, only one additional signature string is stored (besides the result list). If needed, we can avoid creating temporary strings when using fixed-size frequency arrays by using immutable keys or reusable buffers.
* The algorithm runs in a single pass (left-to-right) and yields the unique final array as the problem guarantees.

---

## Author

[Md. Aarzoo Islam](https://bento.me/withaarzoo)
