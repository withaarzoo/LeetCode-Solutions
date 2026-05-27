# 3121. Count the Number of Special Characters II

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

---

## Problem Summary

In this LeetCode string problem, we need to count how many characters are considered "special".

A character is called special when:

* The lowercase version exists in the string
* The uppercase version also exists
* Every lowercase occurrence appears before the first uppercase occurrence

The input is a single string containing uppercase and lowercase English letters.

The output is an integer representing the total number of special characters.

This is a good string processing and indexing problem because it tests how efficiently we track character positions inside a string.

---

## Constraints

| Constraint                              | Value                   |
| --------------------------------------- | ----------------------- |
| `1 <= word.length <= 2 * 10^5`          | Large input size        |
| `word` consists only of English letters | Uppercase and lowercase |

---

## Intuition

When I first looked at this DSA problem, I realized I did not need to store every occurrence of every character.

I only cared about two things:

* The last position of each lowercase letter
* The first position of each uppercase letter

Why?

Because if the last lowercase position is still before the first uppercase position, then all lowercase letters naturally appear before uppercase.

That immediately turns the problem into a simple indexing problem instead of a complicated string traversal problem.

This approach is fast, clean, and works perfectly for competitive programming constraints.

---

## Approach

First, I create two arrays of size 26.

* One array stores the last occurrence of lowercase letters
* Another array stores the first occurrence of uppercase letters

Then I traverse the string once.

For every character:

* If it is lowercase:

  * update its latest position
* If it is uppercase:

  * store its first position only

After processing the string, I loop through all 26 letters.

For every letter:

* check if lowercase exists
* check if uppercase exists
* verify that:

```text
last lowercase index < first uppercase index
```

If this condition is true, that character is special.

Finally, I return the total count.

This solution is efficient because the string is scanned only once.

---

## Data Structures Used

| Data Structure   | Purpose                         |
| ---------------- | ------------------------------- |
| Array of size 26 | Store last lowercase positions  |
| Array of size 26 | Store first uppercase positions |
| Integer variable | Store final count               |

I used arrays because the English alphabet size is fixed at 26, which makes lookup operations extremely fast.

---

## Operations & Behavior Summary

1. Initialize two arrays with `-1`
2. Traverse the string from left to right
3. Track:

   * latest lowercase occurrence
   * earliest uppercase occurrence
4. Loop through all letters from `a` to `z`
5. Check whether lowercase and uppercase both exist
6. Verify ordering condition
7. Count valid special characters
8. Return answer

This is a classic linear-time string algorithm often used in coding interview problems and competitive programming solutions.

---

## Complexity

| Type             | Complexity | Explanation                                                              |
| ---------------- | ---------- | ------------------------------------------------------------------------ |
| Time Complexity  | `O(n)`     | The string is traversed once where `n` is the length of the input string |
| Space Complexity | `O(1)`     | Only fixed-size arrays of length 26 are used                             |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int numberOfSpecialChars(string word) {
        
        // Store last occurrence of lowercase letters
        vector<int> lower(26, -1);

        // Store first occurrence of uppercase letters
        vector<int> upper(26, -1);

        // Traverse the string
        for (int i = 0; i < word.size(); i++) {

            char ch = word[i];

            // If lowercase letter
            if (ch >= 'a' && ch <= 'z') {

                // Update last occurrence
                lower[ch - 'a'] = i;
            }
            else {

                // Convert uppercase letter to index
                int idx = ch - 'A';

                // Store only first occurrence
                if (upper[idx] == -1) {
                    upper[idx] = i;
                }
            }
        }

        int ans = 0;

        // Check all 26 letters
        for (int i = 0; i < 26; i++) {

            // Both lowercase and uppercase must exist
            if (lower[i] != -1 && upper[i] != -1) {

                // All lowercase must come before uppercase
                if (lower[i] < upper[i]) {
                    ans++;
                }
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int numberOfSpecialChars(String word) {
        
        // Store last occurrence of lowercase letters
        int[] lower = new int[26];

        // Store first occurrence of uppercase letters
        int[] upper = new int[26];

        // Initialize arrays with -1
        Arrays.fill(lower, -1);
        Arrays.fill(upper, -1);

        // Traverse the string
        for (int i = 0; i < word.length(); i++) {

            char ch = word.charAt(i);

            // If lowercase letter
            if (ch >= 'a' && ch <= 'z') {

                // Update last occurrence
                lower[ch - 'a'] = i;
            }
            else {

                int idx = ch - 'A';

                // Store only first occurrence
                if (upper[idx] == -1) {
                    upper[idx] = i;
                }
            }
        }

        int ans = 0;

        // Check all letters
        for (int i = 0; i < 26; i++) {

            // Both lowercase and uppercase must exist
            if (lower[i] != -1 && upper[i] != -1) {

                // Lowercase must appear before uppercase
                if (lower[i] < upper[i]) {
                    ans++;
                }
            }
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} word
 * @return {number}
 */
var numberOfSpecialChars = function(word) {
    
    // Store last occurrence of lowercase letters
    const lower = new Array(26).fill(-1);

    // Store first occurrence of uppercase letters
    const upper = new Array(26).fill(-1);

    // Traverse the string
    for (let i = 0; i < word.length; i++) {

        const ch = word[i];

        // If lowercase letter
        if (ch >= 'a' && ch <= 'z') {

            // Update last occurrence
            lower[ch.charCodeAt(0) - 97] = i;
        }
        else {

            const idx = ch.charCodeAt(0) - 65;

            // Store only first occurrence
            if (upper[idx] === -1) {
                upper[idx] = i;
            }
        }
    }

    let ans = 0;

    // Check all letters
    for (let i = 0; i < 26; i++) {

        // Both lowercase and uppercase must exist
        if (lower[i] !== -1 && upper[i] !== -1) {

            // Lowercase must come before uppercase
            if (lower[i] < upper[i]) {
                ans++;
            }
        }
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def numberOfSpecialChars(self, word: str) -> int:
        
        # Store last occurrence of lowercase letters
        lower = [-1] * 26

        # Store first occurrence of uppercase letters
        upper = [-1] * 26

        # Traverse the string
        for i, ch in enumerate(word):

            # If lowercase letter
            if 'a' <= ch <= 'z':

                # Update last occurrence
                lower[ord(ch) - ord('a')] = i

            else:
                idx = ord(ch) - ord('A')

                # Store only first occurrence
                if upper[idx] == -1:
                    upper[idx] = i

        ans = 0

        # Check all 26 letters
        for i in range(26):

            # Both lowercase and uppercase must exist
            if lower[i] != -1 and upper[i] != -1:

                # Lowercase must come before uppercase
                if lower[i] < upper[i]:
                    ans += 1

        return ans
```

### Go

```go
func numberOfSpecialChars(word string) int {
    
    // Store last occurrence of lowercase letters
    lower := make([]int, 26)

    // Store first occurrence of uppercase letters
    upper := make([]int, 26)

    // Initialize arrays with -1
    for i := 0; i < 26; i++ {
        lower[i] = -1
        upper[i] = -1
    }

    // Traverse the string
    for i, ch := range word {

        // If lowercase letter
        if ch >= 'a' && ch <= 'z' {

            // Update last occurrence
            lower[ch-'a'] = i
        } else {

            idx := ch - 'A'

            // Store only first occurrence
            if upper[idx] == -1 {
                upper[idx] = i
            }
        }
    }

    ans := 0

    // Check all letters
    for i := 0; i < 26; i++ {

        // Both lowercase and uppercase must exist
        if lower[i] != -1 && upper[i] != -1 {

            // Lowercase must come before uppercase
            if lower[i] < upper[i] {
                ans++
            }
        }
    }

    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic stays exactly the same in all five programming languages.

The only difference is syntax.

### Step 1: Store character positions

The algorithm starts by creating two arrays.

One array tracks lowercase letters.

Another tracks uppercase letters.

Each position represents one character:

| Index | Character |
| ----- | --------- |
| 0     | a / A     |
| 1     | b / B     |
| ...   | ...       |
| 25    | z / Z     |

Initially, every value is `-1`.

That means the character has not appeared yet.

---

### Step 2: Traverse the string

Next, the algorithm scans the string from left to right.

For lowercase letters:

* update the latest index
* overwrite previous positions

For uppercase letters:

* store only the first occurrence
* ignore future uppercase appearances

This matters because the problem specifically says:

> every lowercase occurrence must appear before the first uppercase occurrence

So the first uppercase index becomes extremely important.

---

### Step 3: Compare positions

After scanning the entire string, the algorithm checks all 26 letters.

For every letter:

* lowercase must exist
* uppercase must exist

Then it verifies:

```text
last lowercase index < first uppercase index
```

If true, that character is counted as special.

---

### Why this works

Suppose:

```text
aaAbcBC
```

For character `a`:

* last lowercase `a` = index 1
* first uppercase `A` = index 2

Since:

```text
1 < 2
```

all lowercase `a` letters appear before uppercase `A`.

So `a` is special.

---

### Important edge case

Consider:

```text
AbBCab
```

For character `a`:

* lowercase `a` appears after uppercase `A`

That immediately breaks the condition.

So `a` is not special.

The same logic applies to every character.

---

### Why arrays are better here

Some people may use hash maps or sets.

But arrays are faster and simpler because:

* only 26 English letters exist
* constant-time lookup is guaranteed
* memory usage stays very small

That makes this solution ideal for LeetCode medium string problems.

---

## Examples

### Example 1

Input:

```text
word = "aaAbcBC"
```

Output:

```text
3
```

Explanation:

* `a` is valid
* `b` is valid
* `c` is valid

All lowercase letters appear before uppercase letters.

---

### Example 2

Input:

```text
word = "abc"
```

Output:

```text
0
```

Explanation:

No uppercase letters exist.

So no character can be special.

---

### Example 3

Input:

```text
word = "AbBCab"
```

Output:

```text
0
```

Explanation:

Lowercase letters appear after uppercase letters.

That breaks the required condition.

---

## How to Use / Run Locally

### C++

Compile:

```bash
g++ solution.cpp -o solution
```

Run:

```bash
./solution
```

---

### Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

---

### JavaScript

Run using Node.js:

```bash
node solution.js
```

---

### Python3

Run:

```bash
python solution.py
```

---

### Go

Run:

```bash
go run solution.go
```

---

## Notes & Optimizations

* This solution is already optimal for the given constraints.
* Since only 26 letters exist, arrays are faster than maps.
* The algorithm avoids nested loops completely.
* The solution handles large strings efficiently.
* Edge cases like missing uppercase or lowercase letters are naturally handled.
* Another possible approach is using bitmasking, but it makes the solution harder to read without improving complexity much.

This is a clean and interview-friendly string indexing solution.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
