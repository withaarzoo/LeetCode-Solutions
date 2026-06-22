# 1189. Maximum Number of Balloons

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

Given a string `text`, the goal is to determine how many times the word `"balloon"` can be formed using the characters available in the string.

Each character from the input string can only be used once.

To solve this problem, I need to count how many times the required characters appear and then determine which character becomes the limiting factor when forming the word `"balloon"`.

This is a classic LeetCode string counting problem that uses frequency counting and greedy observation.

### Input

A string containing lowercase English letters.

### Output

An integer representing the maximum number of times the word `"balloon"` can be formed.

---

## Constraints

| Constraint                                   | Value |
| -------------------------------------------- | ----- |
| `1 <= text.length <= 10^4`                   | Valid |
| `text` consists of lowercase English letters | Valid |

---

## Intuition

The first thing I noticed is that the target word is always the same: `"balloon"`.

Instead of trying to build the word repeatedly, I can simply count how many times each required character appears.

The word `"balloon"` contains:

* `b` → 1 time
* `a` → 1 time
* `l` → 2 times
* `o` → 2 times
* `n` → 1 time

This means that even if I have many copies of some letters, the answer is limited by the character that runs out first.

For example, if I only have one `'n'`, then I can only build one `"balloon"` regardless of how many other letters exist.

That observation leads directly to a frequency-counting solution.

---

## Approach

1. Count the frequency of every character in the input string.
2. Extract the counts of the characters needed for `"balloon"`.
3. Since `'l'` and `'o'` are required twice, divide their counts by 2.
4. Find the minimum among:

   * count of `b`
   * count of `a`
   * count of `l / 2`
   * count of `o / 2`
   * count of `n`
5. Return that minimum value.

The smallest value tells us how many complete copies of `"balloon"` can be formed.

---

## Data Structures Used

### Frequency Array

A frequency array of size 26 is used to count occurrences of each lowercase English letter.

Why I chose it:

* Constant-time updates
* Constant-time lookups
* Simple implementation
* Optimal for lowercase alphabet problems

---

## Operations & Behavior Summary

The algorithm performs the following steps:

1. Traverse the input string once.
2. Count occurrences of every character.
3. Read counts for `b`, `a`, `l`, `o`, and `n`.
4. Adjust counts of `l` and `o` because each is required twice.
5. Find the minimum valid count.
6. Return the result.

In simple terms:

* Count letters.
* Check how many complete `"balloon"` words each letter can support.
* Return the smallest value.

---

## Complexity

| Metric           | Complexity | Explanation                                                           |
| ---------------- | ---------- | --------------------------------------------------------------------- |
| Time Complexity  | O(n)       | I scan the input string once, where `n` is the length of `text`.      |
| Space Complexity | O(1)       | The frequency array always stores 26 values regardless of input size. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxNumberOfBalloons(string text) {
        // Store frequency of all lowercase letters
        vector<int> freq(26, 0);

        // Count each character
        for (char ch : text) {
            freq[ch - 'a']++;
        }

        // Find the limiting character count
        return min({
            freq['b' - 'a'],        // Need 1 'b'
            freq['a' - 'a'],        // Need 1 'a'
            freq['l' - 'a'] / 2,    // Need 2 'l'
            freq['o' - 'a'] / 2,    // Need 2 'o'
            freq['n' - 'a']         // Need 1 'n'
        });
    }
};
```

### Java

```java
class Solution {
    public int maxNumberOfBalloons(String text) {
        // Store frequency of all lowercase letters
        int[] freq = new int[26];

        // Count each character
        for (char ch : text.toCharArray()) {
            freq[ch - 'a']++;
        }

        // Return the smallest possible complete balloon count
        return Math.min(
            Math.min(freq['b' - 'a'], freq['a' - 'a']),
            Math.min(
                Math.min(freq['l' - 'a'] / 2, freq['o' - 'a'] / 2),
                freq['n' - 'a']
            )
        );
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} text
 * @return {number}
 */
var maxNumberOfBalloons = function(text) {
    // Store frequency of all lowercase letters
    const freq = new Array(26).fill(0);

    // Count each character
    for (const ch of text) {
        freq[ch.charCodeAt(0) - 97]++;
    }

    // Return the limiting character count
    return Math.min(
        freq['b'.charCodeAt(0) - 97],      // Need 1 'b'
        freq['a'.charCodeAt(0) - 97],      // Need 1 'a'
        Math.floor(freq['l'.charCodeAt(0) - 97] / 2), // Need 2 'l'
        Math.floor(freq['o'.charCodeAt(0) - 97] / 2), // Need 2 'o'
        freq['n'.charCodeAt(0) - 97]       // Need 1 'n'
    );
};
```

### Python3

```python
class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        # Store frequency of all lowercase letters
        freq = [0] * 26

        # Count each character
        for ch in text:
            freq[ord(ch) - ord('a')] += 1

        # Return the smallest possible complete balloon count
        return min(
            freq[ord('b') - ord('a')],      # Need 1 'b'
            freq[ord('a') - ord('a')],      # Need 1 'a'
            freq[ord('l') - ord('a')] // 2, # Need 2 'l'
            freq[ord('o') - ord('a')] // 2, # Need 2 'o'
            freq[ord('n') - ord('a')]       # Need 1 'n'
        )
```

### Go

```go
func maxNumberOfBalloons(text string) int {
    // Store frequency of all lowercase letters
    freq := make([]int, 26)

    // Count each character
    for _, ch := range text {
        freq[ch-'a']++
    }

    // Start with count of 'b'
    ans := freq['b'-'a']

    // Update answer with the minimum possible value
    if freq['a'-'a'] < ans {
        ans = freq['a'-'a']
    }

    if freq['l'-'a']/2 < ans {
        ans = freq['l'-'a'] / 2
    }

    if freq['o'-'a']/2 < ans {
        ans = freq['o'-'a'] / 2
    }

    if freq['n'-'a'] < ans {
        ans = freq['n'-'a']
    }

    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same in all five languages.

### Step 1: Create a Frequency Counter

I first create a structure that stores how many times each letter appears.

Since only lowercase English letters are present, I can use a fixed-size array of length 26.

---

### Step 2: Count Every Character

I iterate through the string once.

For each character:

* Find its alphabet position.
* Increase its frequency count.

After this pass, I know exactly how many times every letter appears.

---

### Step 3: Check Required Balloon Characters

The word `"balloon"` requires:

| Character | Required Count |
| --------- | -------------- |
| b         | 1              |
| a         | 1              |
| l         | 2              |
| o         | 2              |
| n         | 1              |

These are the only characters that matter.

All other letters can be ignored.

---

### Step 4: Handle Repeated Letters

The tricky part is that:

* `l` appears twice
* `o` appears twice

If I have:

```text
l = 5
```

I cannot build five balloons.

I can only build:

```text
5 / 2 = 2
```

complete balloons.

The same idea applies to `o`.

---

### Step 5: Find the Limiting Character

Every balloon needs all required letters.

Suppose:

```text
b = 7
a = 4
l = 8
o = 6
n = 2
```

Possible balloons:

```text
b → 7
a → 4
l → 4
o → 3
n → 2
```

The smallest value is:

```text
2
```

So only two complete copies of `"balloon"` can be formed.

---

### Step 6: Return the Answer

The minimum value among all required character counts is the final answer.

That value represents the maximum number of valid `"balloon"` words that can be constructed.

---

## Examples

### Example 1

**Input**

```text
text = "nlaebolko"
```

**Output**

```text
1
```

**Explanation**

Character counts are enough to form exactly one `"balloon"`.

---

### Example 2

**Input**

```text
text = "loonbalxballpoon"
```

**Output**

```text
2
```

**Explanation**

All required letters exist in sufficient quantity to build two copies of `"balloon"`.

---

### Example 3

**Input**

```text
text = "leetcode"
```

**Output**

```text
0
```

**Explanation**

Several required characters are missing, so no complete `"balloon"` can be formed.

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
javac Solution.java
```

Run:

```bash
java Solution
```

---

### JavaScript

Run:

```bash
node solution.js
```

---

### Python3

Run:

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Run:

```bash
go run solution.go
```

Build:

```bash
go build solution.go
```

---

## Notes & Optimizations

* This is a frequency counting problem.
* A hash map can also be used, but a fixed-size array is slightly faster.
* The solution only needs one pass through the string.
* No sorting is required.
* No extra processing is required after counting frequencies.
* This is already the optimal solution for the given constraints.
* The key observation is remembering that `l` and `o` are needed twice in the target word.

### Common Mistakes

1. Forgetting that `l` appears twice.
2. Forgetting that `o` appears twice.
3. Taking the maximum count instead of the minimum count.
4. Trying to repeatedly construct the word instead of using frequency counts.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
