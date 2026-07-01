# 1358. Number of Substrings Containing All Three Characters

A beginner-friendly solution for LeetCode 1358 using the Sliding Window technique. This repository explains the intuition, approach, complexity analysis, and provides multi-language implementations in C++, Java, JavaScript, Python3, and Go.

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
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

Given a string that contains only the characters `a`, `b`, and `c`, the goal is to count how many substrings contain at least one occurrence of all three characters.

A substring is a continuous part of the original string.

Instead of checking every possible substring one by one, the challenge is to find an efficient algorithm that works even when the string becomes very large.

This problem is a classic example of the **Sliding Window** technique and is commonly asked in coding interviews because it tests both observation skills and optimization.

---

## Constraints

| Constraint | Value |
|------------|-------|
| String Length | `3 <= s.length <= 5 × 10^4` |
| Characters | Only `a`, `b`, and `c` |

---

## Intuition

My first thought was to generate every possible substring and check whether it contains all three characters.

That approach is simple, but it becomes too slow because there are almost `n²` substrings.

Then I noticed something important.

Once a window already contains `a`, `b`, and `c`, extending that window further to the right will never remove those characters. That means every larger substring starting from the same left position will also be valid.

This observation makes the Sliding Window approach a perfect fit.

---

## Approach

I maintain a window using two pointers.

First, I expand the right pointer one character at a time.

While expanding, I keep track of how many times `a`, `b`, and `c` appear inside the current window.

As soon as the window contains all three characters, I know every substring ending from the current right position to the end of the string is also valid.

So instead of counting them one by one, I add them together in a single operation.

Then I move the left pointer forward to search for the next valid window.

Since each pointer moves only forward, the entire algorithm runs in linear time.

---

## Data Structures Used

| Data Structure | Purpose |
|---------------|---------|
| Frequency Array | Stores how many times `a`, `b`, and `c` appear inside the current sliding window. |
| Two Pointers | Represent the left and right boundaries of the current window. |
| Integer Variables | Store the answer and window positions. |

No extra data structures that grow with the input size are required.

---

## Operations & Behavior Summary

1. Create a frequency array for the three characters.
2. Initialize both window pointers.
3. Expand the right pointer across the string.
4. Update the frequency of the current character.
5. Check whether the current window contains all three characters.
6. If it does:
   - Count every valid substring that starts at the current left position.
   - Shrink the window by moving the left pointer.
7. Continue until the right pointer reaches the end.
8. Return the total number of valid substrings.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | `O(n)` | Both pointers move across the string only once. |
| Space Complexity | `O(1)` | Only a fixed-size frequency array of three elements is used. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int numberOfSubstrings(string s) {
        // Store the frequency of 'a', 'b', and 'c' inside the current window
        vector<int> freq(3, 0);

        int left = 0;
        int ans = 0;
        int n = s.size();

        // Expand the window one character at a time
        for (int right = 0; right < n; right++) {

            // Add the current character into the window
            freq[s[right] - 'a']++;

            // Keep shrinking while the window contains all three characters
            while (freq[0] > 0 && freq[1] > 0 && freq[2] > 0) {

                // Every substring ending from 'right' to the last index is valid
                ans += (n - right);

                // Remove the leftmost character before shrinking the window
                freq[s[left] - 'a']--;

                // Move the left pointer forward
                left++;
            }
        }

        // Return the total number of valid substrings
        return ans;
    }
};
```

### Java

```java
class Solution {
    public int numberOfSubstrings(String s) {

        // Store the frequency of 'a', 'b', and 'c'
        int[] freq = new int[3];

        int left = 0;
        int ans = 0;
        int n = s.length();

        // Expand the window
        for (int right = 0; right < n; right++) {

            // Include the current character
            freq[s.charAt(right) - 'a']++;

            // Shrink while all three characters exist
            while (freq[0] > 0 && freq[1] > 0 && freq[2] > 0) {

                // Count every possible ending position
                ans += (n - right);

                // Remove the leftmost character
                freq[s.charAt(left) - 'a']--;

                // Move left forward
                left++;
            }
        }

        // Return the final answer
        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @return {number}
 */
var numberOfSubstrings = function(s) {

    // Store the frequency of 'a', 'b', and 'c'
    const freq = [0, 0, 0];

    let left = 0;
    let ans = 0;
    const n = s.length;

    // Expand the window
    for (let right = 0; right < n; right++) {

        // Add the current character
        freq[s.charCodeAt(right) - 97]++;

        // Shrink while all characters are present
        while (freq[0] > 0 && freq[1] > 0 && freq[2] > 0) {

            // Every larger ending index is also valid
            ans += (n - right);

            // Remove the leftmost character
            freq[s.charCodeAt(left) - 97]--;

            // Move left forward
            left++;
        }
    }

    // Return the answer
    return ans;
};
```

### Python3

```python
class Solution:
    def numberOfSubstrings(self, s: str) -> int:

        # Store the frequency of 'a', 'b', and 'c'
        freq = [0] * 3

        left = 0
        ans = 0
        n = len(s)

        # Expand the window
        for right in range(n):

            # Add the current character
            freq[ord(s[right]) - ord('a')] += 1

            # Shrink while all three characters exist
            while freq[0] > 0 and freq[1] > 0 and freq[2] > 0:

                # Count every valid substring starting at 'left'
                ans += n - right

                # Remove the leftmost character
                freq[ord(s[left]) - ord('a')] -= 1

                # Move left forward
                left += 1

        # Return the total count
        return ans
```

### Go

```go
func numberOfSubstrings(s string) int {

 // Store the frequency of 'a', 'b', and 'c'
 freq := make([]int, 3)

 left := 0
 ans := 0
 n := len(s)

 // Expand the window
 for right := 0; right < n; right++ {

  // Add the current character
  freq[s[right]-'a']++

  // Shrink while all three characters exist
  for freq[0] > 0 && freq[1] > 0 && freq[2] > 0 {

   // Every larger ending position is also valid
   ans += n - right

   // Remove the leftmost character
   freq[s[left]-'a']--

   // Move the left pointer
   left++
  }
 }

 // Return the final answer
 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is exactly the same in every language. Only the syntax changes.

First, a frequency array is created to count how many times each of the three characters appears inside the current window.

Next, two pointers are initialized.

The left pointer marks where the current window starts.

The right pointer keeps expanding the window by moving one character at a time.

Whenever a new character enters the window, its frequency is increased.

After updating the frequency, the algorithm checks whether the window contains at least one `a`, one `b`, and one `c`.

If the condition is satisfied, the current window is already valid.

Instead of checking every possible ending separately, the algorithm immediately counts every substring that starts at the current left position and ends anywhere from the current right index to the last character.

After counting those substrings, the algorithm removes the leftmost character from the window.

The left pointer is then moved forward.

If the window still contains all three characters, the same counting process continues.

Otherwise, the window becomes invalid and the algorithm resumes expanding the right pointer.

Every character enters the window once and leaves the window once.

That is why the overall running time remains linear.

This behavior is identical in C++, Java, JavaScript, Python3, and Go. The only differences are language syntax and array access.

---

## Examples

### Example 1

**Input**

```text
s = "abcabc"
```

**Output**

```text
10
```

**Trace**

- First valid window is `"abc"`
- Count every possible extension.
- Continue shrinking and expanding.
- Final answer becomes **10**.

---

### Example 2

**Input**

```text
s = "aaacb"
```

**Output**

```text
3
```

**Trace**

The valid substrings are:

- `aaacb`
- `aacb`
- `acb`

Total = **3**

---

### Example 3

**Input**

```text
s = "abc"
```

**Output**

```text
1
```

**Trace**

The entire string already contains all three characters, so there is exactly one valid substring.

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone https://github.com/your-username/your-repository.git
```

Move into the project directory.

```bash
cd your-repository
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

---

### Java

Compile

```bash
javac Solution.java
```

Run

```bash
java Solution
```

---

### JavaScript

Run using Node.js

```bash
node solution.js
```

---

### Python3

Run

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Run

```bash
go run solution.go
```

---

## Notes & Optimizations

- The brute-force solution takes `O(n²)` time and is too slow for the maximum input size.
- The Sliding Window technique reduces the complexity to `O(n)`.
- Since there are only three possible characters, a fixed-size frequency array is faster than using a hash map.
- Every character is processed at most twice, once when entering the window and once when leaving it.
- This problem is a great example of counting substrings efficiently without generating every substring explicitly.
- The same Sliding Window pattern is useful for many interview problems involving substrings, character frequencies, and minimum or maximum windows.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
