# 3120. Count the Number of Special Characters I

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

This LeetCode problem asks us to count how many characters in a string are considered "special".

A character is called special when:

* its lowercase version exists in the string
* and its uppercase version also exists in the same string

For example:

* `'a'` and `'A'` together make `'a'` a special character
* `'b'` without `'B'` does not count

The input is a string containing only English uppercase and lowercase letters.

The output should be the total number of special characters present in the string.

This is a simple string and hash set problem that focuses on character lookup and efficient checking.

---

## Constraints

| Constraint                                                   | Value                 |
| ------------------------------------------------------------ | --------------------- |
| `1 <= word.length <= 50`                                     | String length range   |
| `word` contains only uppercase and lowercase English letters | Valid characters only |

---

## Intuition

When I first looked at this problem, I noticed that I only care whether a character exists or not.

The number of times a letter appears does not matter.

So instead of comparing every character with every other character, I thought it would be much easier to store all characters once and then simply check whether both lowercase and uppercase versions are present.

Since there are only 26 English letters, checking all possible characters becomes very fast and clean.

This immediately suggested using a hash set.

---

## Approach

First, I store every character from the string inside a set.

The set helps me quickly check whether a character exists.

Then I loop from `'a'` to `'z'`.

For every lowercase letter:

1. Check whether the lowercase character exists
2. Check whether the uppercase version also exists
3. If both are present, increase the answer

Finally, return the total count.

This approach avoids unnecessary nested loops and keeps the solution simple.

---

## Data Structures Used

| Data Structure  | Why I Used It                                     |
| --------------- | ------------------------------------------------- |
| Hash Set / Set  | Fast lookup for checking whether characters exist |
| Integer Counter | Stores the total number of special characters     |

The set is the most important part of this solution because it gives near constant-time lookup.

---

## Operations & Behavior Summary

Here is what the algorithm does step by step:

1. Read the input string
2. Store all characters inside a set
3. Start looping through letters from `'a'` to `'z'`
4. For each letter:

   * check lowercase existence
   * check uppercase existence
5. If both exist:

   * increase the answer count
6. Return the final answer

The algorithm only scans the string once and then performs 26 checks.

---

## Complexity

| Type             | Complexity | Explanation                                                                                                |
| ---------------- | ---------- | ---------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n)`     | `n` is the length of the string. Building the set takes linear time. Checking 26 letters is constant work. |
| Space Complexity | `O(n)`     | Extra space is used for the set storing characters from the string.                                        |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int numberOfSpecialChars(string word) {
        
        // Store all characters present in the string
        unordered_set<char> st(word.begin(), word.end());

        // Variable to store final answer
        int count = 0;

        // Check every lowercase letter from 'a' to 'z'
        for(char ch = 'a'; ch <= 'z'; ch++) {

            // If both lowercase and uppercase exist,
            // then this character is special
            if(st.count(ch) && st.count(ch - 'a' + 'A')) {
                count++;
            }
        }

        // Return total special characters
        return count;
    }
};
```

### Java

```java
class Solution {
    public int numberOfSpecialChars(String word) {
        
        // Store all characters inside a HashSet
        HashSet<Character> set = new HashSet<>();

        // Add every character from the string
        for(char ch : word.toCharArray()) {
            set.add(ch);
        }

        // Variable to store answer
        int count = 0;

        // Check all lowercase English letters
        for(char ch = 'a'; ch <= 'z'; ch++) {

            // Check if both lowercase and uppercase exist
            if(set.contains(ch) && set.contains((char)(ch - 'a' + 'A'))) {
                count++;
            }
        }

        // Return final count
        return count;
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
    
    // Store all characters inside a Set
    const set = new Set(word);

    // Variable to store answer
    let count = 0;

    // Loop through all lowercase letters
    for(let i = 0; i < 26; i++) {

        // Current lowercase character
        let lower = String.fromCharCode(97 + i);

        // Corresponding uppercase character
        let upper = String.fromCharCode(65 + i);

        // If both exist, increase answer
        if(set.has(lower) && set.has(upper)) {
            count++;
        }
    }

    // Return total special characters
    return count;
};
```

### Python3

```python
class Solution:
    def numberOfSpecialChars(self, word: str) -> int:
        
        # Store all characters from the string
        st = set(word)

        # Variable to store answer
        count = 0

        # Check every lowercase letter
        for i in range(26):

            # Current lowercase character
            lower = chr(ord('a') + i)

            # Corresponding uppercase character
            upper = chr(ord('A') + i)

            # If both exist, it is a special character
            if lower in st and upper in st:
                count += 1

        # Return final answer
        return count
```

### Go

```go
func numberOfSpecialChars(word string) int {
    
    // Map used like a set to store characters
    st := make(map[rune]bool)

    // Store every character from the string
    for _, ch := range word {
        st[ch] = true
    }

    // Variable to store answer
    count := 0

    // Check all lowercase English letters
    for ch := 'a'; ch <= 'z'; ch++ {

        // Find corresponding uppercase character
        upper := ch - 'a' + 'A'

        // If both exist, increase answer
        if st[ch] && st[upper] {
            count++
        }
    }

    // Return final count
    return count
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic stays exactly the same in all five languages.

The only difference is the syntax used for sets and loops.

### Step 1 — Store Characters

The first thing I do is store every character from the string into a set.

Why?

Because I only care whether a character exists.

Example:

```text
word = "aaAbBc"
```

After storing inside a set:

```text
{a, A, b, B, c}
```

Duplicate characters automatically disappear, which is perfect for this problem.

---

### Step 2 — Loop Through English Letters

Instead of scanning the string repeatedly, I simply check all 26 lowercase letters.

For every character from `'a'` to `'z'`:

* I check whether lowercase exists
* I check whether uppercase exists

This keeps the solution clean and predictable.

---

### Step 3 — Convert Lowercase to Uppercase

For every lowercase letter:

* `'a'` becomes `'A'`
* `'b'` becomes `'B'`
* and so on

Different languages handle this conversion differently:

* C++ and Java use character arithmetic
* JavaScript uses ASCII conversion functions
* Python uses `chr()` and `ord()`
* Go uses rune arithmetic

But the idea is identical everywhere.

---

### Step 4 — Count Special Characters

If both lowercase and uppercase versions exist:

* increase the answer counter

Example:

```text
a exists
A exists
```

So:

```text
count++
```

This continues for all 26 letters.

---

### Step 5 — Return Final Answer

After checking every letter:

* return the total count

That final value is the number of special characters in the string.

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

```text
a and A exist
b and B exist
c and C exist
```

Total special characters = 3

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

```text
No uppercase characters exist
```

So no character is special.

---

### Example 3

Input:

```text
word = "abBCab"
```

Output:

```text
1
```

Explanation:

```text
Only b and B exist together
```

So the answer is 1.

---

## How to Use / Run Locally

### C++

Compile:

```bash
g++ main.cpp -o main
```

Run:

```bash
./main
```

---

### Java

Compile:

```bash
javac Main.java
```

Run:

```bash
java Main
```

---

### JavaScript

Run using Node.js:

```bash
node main.js
```

---

### Python3

Run:

```bash
python main.py
```

or

```bash
python3 main.py
```

---

### Go

Run:

```bash
go run main.go
```

---

## Notes & Optimizations

A brute force solution could compare every character with every other character.

That would take `O(n²)` time.

Using a hash set improves the lookup speed and makes the solution much more efficient.

Some alternative approaches:

* Using two boolean arrays of size 26
* Using bit manipulation
* Using frequency arrays

But the hash set solution is probably the easiest to read and explain.

Edge cases worth testing:

* String with only lowercase letters
* String with only uppercase letters
* Empty combinations like `"aA"`
* Repeated characters like `"aaaaAAAA"`

Because the problem size is small, almost any optimized approach works well, but the set-based solution keeps the code short and beginner-friendly.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
