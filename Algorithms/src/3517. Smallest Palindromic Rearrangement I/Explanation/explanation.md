# 3517. Smallest Palindromic Rearrangement I | LeetCode Solution

A beginner-friendly explanation and optimized solution for **LeetCode 3517 - Smallest Palindromic Rearrangement I**.

This repository explains the intuition, algorithm, complexity analysis, and provides implementations in **C++**, **Java**, **JavaScript**, **Python3**, and **Go**. The solution uses a simple **greedy approach with frequency counting** to construct the **lexicographically smallest palindrome** in linear time.

---

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

---

## Problem Summary

You are given a string that is already guaranteed to be a palindrome.

Your task is to rearrange its characters so that the resulting string is still a palindrome, but this time it must also be the **lexicographically smallest** possible palindrome.

Since every rearrangement must remain a palindrome, the left half completely determines the right half. That observation makes the problem much simpler than it first appears.

### Input

- A lowercase English string `s`
- The string is already a valid palindrome

### Output

- Return the lexicographically smallest palindromic rearrangement of `s`

---

## Constraints

| Constraint | Value |
| ------------ | ------- |
| `1 <= s.length <= 10^5` | Yes |
| Characters | Lowercase English letters |
| Input Guarantee | `s` is already a palindrome |

---

## Intuition

The first thing I noticed was that I didn't actually need to rearrange the whole string.

A palindrome is made of two mirror-image halves. Once the left half is fixed, the right half is automatically determined because it must be its reverse.

That means the real problem is simply building the smallest possible left half.

To make it as small as possible, I place smaller letters before larger ones. Every character contributes half of its occurrences to the left side, while an odd-frequency character naturally becomes the middle of the palindrome.

This greedy idea immediately gives the smallest answer.

---

## Approach

I solve the problem in a few simple steps.

1. Count the frequency of every character.
2. Build the left half by taking `frequency / 2` copies of each letter from `'a'` to `'z'`.
3. If a character has an odd frequency, store it as the middle character.
4. Reverse the left half to create the right half.
5. Join the three parts together.

The left half is built in sorted order, so the complete palindrome is automatically the lexicographically smallest one.

---

## Data Structures Used

| Data Structure | Purpose |
| --------------- | --------- |
| Frequency Array | Stores how many times each letter appears |
| String / StringBuilder | Builds the left half, middle, and final answer |
| Reversed String | Creates the mirrored right half |

The frequency array has only 26 entries because the string contains only lowercase English letters.

---

## Operations & Behavior Summary

The algorithm performs the following operations.

1. Read every character once.
2. Count the frequency of each letter.
3. Visit letters from `'a'` to `'z'`.
4. Add half of every frequency to the left side.
5. Save the odd-frequency character as the center.
6. Reverse the left side.
7. Combine left + middle + reversed left.
8. Return the answer.

Because the letters are processed in alphabetical order, the smallest characters always appear as early as possible.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | **O(n)** | The string is scanned once, and the answer is built once. |
| Space Complexity | **O(n)** | The output string requires linear space. The frequency array uses constant extra space. |

Where `n` is the length of the given string.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string smallestPalindrome(string s) {
        // Store frequency of every lowercase letter
        vector<int> freq(26, 0);

        // Count character frequencies
        for (char c : s)
            freq[c - 'a']++;

        string left = "";
        string middle = "";

        // Build the left half and find the middle character
        for (int i = 0; i < 26; i++) {
            // Add half of the occurrences to the left side
            left.append(freq[i] / 2, char('a' + i));

            // If frequency is odd, this character stays in the center
            if (freq[i] % 2)
                middle = char('a' + i);
        }

        // Right half is simply the reverse of the left half
        string right = left;
        reverse(right.begin(), right.end());

        // Combine all three parts
        return left + middle + right;
    }
};
```

### Java

```java
class Solution {
    public String smallestPalindrome(String s) {

        // Store frequency of every lowercase letter
        int[] freq = new int[26];

        // Count character frequencies
        for (char c : s.toCharArray())
            freq[c - 'a']++;

        StringBuilder left = new StringBuilder();
        String middle = "";

        // Build the left half and find the middle character
        for (int i = 0; i < 26; i++) {

            // Add half of the occurrences to the left side
            for (int j = 0; j < freq[i] / 2; j++)
                left.append((char) ('a' + i));

            // Odd frequency character becomes the center
            if (freq[i] % 2 == 1)
                middle = String.valueOf((char) ('a' + i));
        }

        // Right half is the reverse of the left half
        String right = new StringBuilder(left).reverse().toString();

        // Return the complete palindrome
        return left.toString() + middle + right;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @return {string}
 */
var smallestPalindrome = function(s) {

    // Store frequency of every lowercase letter
    const freq = new Array(26).fill(0);

    // Count character frequencies
    for (const ch of s)
        freq[ch.charCodeAt(0) - 97]++;

    let left = "";
    let middle = "";

    // Build the left half and find the middle character
    for (let i = 0; i < 26; i++) {

        // Add half of the occurrences to the left side
        left += String.fromCharCode(97 + i).repeat(Math.floor(freq[i] / 2));

        // Odd frequency character becomes the center
        if (freq[i] % 2)
            middle = String.fromCharCode(97 + i);
    }

    // Right half is the reverse of the left half
    const right = left.split("").reverse().join("");

    // Return the complete palindrome
    return left + middle + right;
};
```

### Python3

```python
class Solution:
    def smallestPalindrome(self, s: str) -> str:

        # Store frequency of every lowercase letter
        freq = [0] * 26

        # Count character frequencies
        for ch in s:
            freq[ord(ch) - ord('a')] += 1

        left = []
        middle = ""

        # Build the left half and find the middle character
        for i in range(26):

            # Add half of the occurrences to the left side
            left.append(chr(ord('a') + i) * (freq[i] // 2))

            # Odd frequency character becomes the center
            if freq[i] % 2:
                middle = chr(ord('a') + i)

        # Convert list into string
        left = "".join(left)

        # Right half is the reverse of the left half
        right = left[::-1]

        # Return the complete palindrome
        return left + middle + right
```

### Go

```go
func smallestPalindrome(s string) string {

 // Store frequency of every lowercase letter
 freq := make([]int, 26)

 // Count character frequencies
 for _, ch := range s {
  freq[ch-'a']++
 }

 left := make([]byte, 0)
 middle := byte(0)

 // Build the left half and find the middle character
 for i := 0; i < 26; i++ {

  // Add half of the occurrences to the left side
  for j := 0; j < freq[i]/2; j++ {
   left = append(left, byte('a'+i))
  }

  // Odd frequency character becomes the center
  if freq[i]%2 == 1 {
   middle = byte('a' + i)
  }
 }

 // Right half is the reverse of the left half
 right := make([]byte, len(left))
 for i := 0; i < len(left); i++ {
  right[i] = left[len(left)-1-i]
 }

 // Build the final answer
 if middle != 0 {
  return string(left) + string(middle) + string(right)
 }

 return string(left) + string(right)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is exactly the same in all five languages. Only the syntax changes.

### Step 1: Count every character

The algorithm starts by counting how many times every lowercase letter appears.

Instead of sorting the string, a frequency array is enough because there are only 26 possible characters.

This makes counting very fast.

---

### Step 2: Build the left half

Next, the algorithm visits every letter from `'a'` to `'z'`.

For each character, it places `frequency / 2` copies into the left half.

For example:

- `a` appears 6 times → 3 go to the left
- `b` appears 4 times → 2 go to the left
- `c` appears 2 times → 1 goes to the left

Placing smaller letters first guarantees the smallest lexicographical order.

---

### Step 3: Find the middle character

If the palindrome has an odd length, exactly one character will have an odd frequency.

That character becomes the middle of the palindrome.

Since the input is already guaranteed to be a palindrome, there can never be more than one odd-frequency character.

---

### Step 4: Build the right half

The right half is simply the reverse of the left half.

This works because every palindrome is perfectly symmetric.

There is no need to calculate the right side separately.

---

### Step 5: Construct the final answer

Finally, the algorithm joins everything together.

```
Left + Middle + Reverse(Left)
```

This always produces a valid palindrome.

Because the left half was built in alphabetical order, the complete string is also the lexicographically smallest palindrome.

The behavior is identical in C++, Java, JavaScript, Python3, and Go. Only the language syntax used to build strings and reverse them is different.

---

## Examples

### Example 1

**Input**

```text
s = "babab"
```

**Output**

```text
abbba
```

### Trace

```
Frequency

a -> 2
b -> 3

Left = "ab"
Middle = "b"
Right = "ba"

Result = "abbba"
```

---

### Example 2

**Input**

```text
s = "daccad"
```

**Output**

```text
acddca
```

### Trace

```
Frequency

a -> 2
c -> 2
d -> 2

Left = "acd"
Right = "dca"

Result = "acddca"
```

---

### Example 3

**Input**

```text
s = "z"
```

**Output**

```text
z
```

### Trace

```
Only one character exists.

It is already the smallest possible palindrome.
```

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone https://github.com/your-username/your-repository.git
```

Move into the project folder.

```bash
cd your-repository
```

### Compile and Run C++

Compile

```bash
g++ solution.cpp -o solution
```

Run

```bash
./solution
```

---

### Compile and Run Java

Compile

```bash
javac Solution.java
```

Run

```bash
java Solution
```

---

### Run JavaScript

```bash
node solution.js
```

---

### Run Python3

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Run Go

```bash
go run solution.go
```

---

## Notes & Optimizations

- The input is already guaranteed to be a palindrome, so there is no need to validate it.
- A frequency array is faster than sorting because the alphabet size is fixed.
- The algorithm processes each character only once, making it suitable even for the maximum input size.
- The greedy construction works because the left half completely determines the final palindrome.
- Sorting the entire string would also work, but counting frequencies is simpler and more efficient.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
