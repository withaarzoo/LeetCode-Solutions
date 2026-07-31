# 3016. Minimum Number of Pushes to Type Word II

A greedy solution to **LeetCode 3016 - Minimum Number of Pushes to Type Word II** with an optimized approach. This repository explains the intuition, algorithm, complexity analysis, and provides implementations in **C++, Java, JavaScript, Python3, and Go**.

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
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

In this problem, we are given a word containing lowercase English letters. We are allowed to remap the letters to the phone keypad keys numbered from **2 to 9** in any way we like.

Each key can hold any number of letters, but every letter must belong to exactly one key.

Typing a letter requires pressing its key based on its position on that key.

- The first letter on a key needs **1 push**.
- The second letter needs **2 pushes**.
- The third letter needs **3 pushes**, and so on.

The goal is to find the **minimum total number of key presses** needed to type the given word after choosing the best possible mapping.

Since we can freely assign letters to keys, the challenge becomes finding the most efficient arrangement.

---

## Constraints

| Constraint | Value |
|------------|-------|
| `1 <= word.length <= 10^5` | Maximum length of the input string |
| `word` | Contains only lowercase English letters |

---

## Intuition

The first thing I noticed was that the actual letters do not matter. Only how often they appear matters.

If one letter appears hundreds of times while another appears only once, it makes sense to give the cheaper position to the more frequent letter.

There are only **8 usable keys**, so only **8 letters** can cost one push each.

The next **8 letters** must cost two pushes.

After that, the next **8 letters** cost three pushes, and the remaining letters cost four pushes.

Because there are only 26 lowercase letters, these four levels are enough for every possible input.

This immediately suggests a greedy solution.

---

## Approach

I solved the problem in the following steps.

1. Count the frequency of every lowercase letter.
2. Sort the frequencies in descending order.
3. Start assigning push costs from the highest frequency.
4. The first 8 frequencies get a cost of 1.
5. The next 8 frequencies get a cost of 2.
6. The next 8 frequencies get a cost of 3.
7. The remaining frequencies get a cost of 4.
8. Multiply every frequency by its assigned cost and add everything together.

This greedy assignment guarantees the minimum total number of pushes.

---

## Data Structures Used

| Data Structure | Purpose |
| --------------- | --------- |
| Frequency Array | Stores how many times each letter appears |
| Sorting | Orders frequencies from highest to lowest so the cheapest positions go to the most frequent letters |

Only a fixed-size array of 26 elements is needed, so memory usage stays constant.

---

## Operations & Behavior Summary

The algorithm performs the following operations.

1. Read the input string.
2. Count the occurrence of every letter.
3. Sort all frequencies in decreasing order.
4. Ignore letters whose frequency is zero.
5. Assign push costs according to their sorted position.
6. Multiply frequency × push cost.
7. Add every contribution to the final answer.
8. Return the minimum number of pushes.

---

## Complexity

| Type | Complexity | Explanation |
|------|------------|-------------|
| Time Complexity | **O(n)** | Counting takes O(n), while sorting only 26 values is constant time. |
| Space Complexity | **O(1)** | Only a frequency array of size 26 is used. |

Where **n** is the length of the input string.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minimumPushes(string word) {
        // Store the frequency of every lowercase letter
        vector<int> freq(26, 0);

        // Count how many times each letter appears
        for (char ch : word) {
            freq[ch - 'a']++;
        }

        // Sort frequencies from largest to smallest
        sort(freq.begin(), freq.end(), greater<int>());

        int ans = 0;

        // Assign push cost according to the position
        for (int i = 0; i < 26; i++) {
            // Skip letters that do not appear
            if (freq[i] == 0) break;

            // Every 8 letters increase the push count by 1
            int pushes = (i / 8) + 1;

            // Add the total contribution of this letter
            ans += freq[i] * pushes;
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int minimumPushes(String word) {
        // Store the frequency of every lowercase letter
        int[] freq = new int[26];

        // Count each letter
        for (char ch : word.toCharArray()) {
            freq[ch - 'a']++;
        }

        // Sort frequencies in increasing order
        Arrays.sort(freq);

        int ans = 0;
        int index = 0;

        // Traverse from the largest frequency to the smallest
        for (int i = 25; i >= 0; i--) {
            // Ignore unused letters
            if (freq[i] == 0) break;

            // Every 8 letters need one extra push
            int pushes = (index / 8) + 1;

            // Add this letter's contribution
            ans += freq[i] * pushes;

            index++;
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
var minimumPushes = function(word) {
    // Store the frequency of every letter
    const freq = new Array(26).fill(0);

    // Count each character
    for (const ch of word) {
        freq[ch.charCodeAt(0) - 97]++;
    }

    // Sort from largest to smallest
    freq.sort((a, b) => b - a);

    let ans = 0;

    // Assign push cost based on sorted position
    for (let i = 0; i < 26; i++) {
        // Stop when no more letters exist
        if (freq[i] === 0) break;

        // Every group of 8 letters gets one more push
        const pushes = Math.floor(i / 8) + 1;

        // Add contribution
        ans += freq[i] * pushes;
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def minimumPushes(self, word: str) -> int:
        # Store the frequency of every lowercase letter
        freq = [0] * 26

        # Count each character
        for ch in word:
            freq[ord(ch) - ord('a')] += 1

        # Sort frequencies from largest to smallest
        freq.sort(reverse=True)

        ans = 0

        # Assign push cost according to position
        for i in range(26):
            # Ignore letters that never appear
            if freq[i] == 0:
                break

            # Every 8 letters increase the push count
            pushes = (i // 8) + 1

            # Add contribution
            ans += freq[i] * pushes

        return ans
```

### Go

```go
func minimumPushes(word string) int {
 // Store the frequency of every lowercase letter
 freq := make([]int, 26)

 // Count each character
 for _, ch := range word {
  freq[ch-'a']++
 }

 // Sort frequencies from largest to smallest
 sort.Slice(freq, func(i, j int) bool {
  return freq[i] > freq[j]
 })

 ans := 0

 // Assign push cost based on sorted position
 for i := 0; i < 26; i++ {
  // Stop once unused letters are reached
  if freq[i] == 0 {
   break
  }

  // Every 8 letters require one extra push
  pushes := (i / 8) + 1

  // Add this letter's total contribution
  ans += freq[i] * pushes
 }

 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is exactly the same in all five languages. Only the syntax changes.

### Step 1 — Count Letter Frequencies

The first job is counting how many times every lowercase letter appears.

Instead of storing every character separately, I keep a frequency array of size 26.

For example,

```
word = "aaabbccd"
```

becomes

```
a = 3
b = 2
c = 2
d = 1
```

Everything else stays zero.

This takes one pass through the string.

---

### Step 2 — Sort the Frequencies

Now I sort the frequency array from largest to smallest.

After sorting,

```
3 2 2 1 0 0 0 ...
```

The most common letters automatically move to the front.

This is exactly what I want because cheaper positions should always be assigned to the letters that appear the most.

---

### Step 3 — Assign Push Costs

The push cost depends only on the position inside the sorted list.

```
Positions 0–7   → 1 push
Positions 8–15  → 2 pushes
Positions 16–23 → 3 pushes
Positions 24–25 → 4 pushes
```

Since every group contains exactly eight letters, the push count increases after every eighth position.

---

### Step 4 — Compute the Answer

For every non-zero frequency,

```
Contribution = Frequency × Push Cost
```

The contributions are added together.

That final sum is the minimum number of pushes required.

---

### Why This Greedy Strategy Works

Suppose two letters have frequencies

```
20
5
```

If the letter appearing 20 times receives a cheaper position than the letter appearing 5 times, the total cost is smaller.

Swapping them would only increase the answer.

Because of this, sorting the frequencies and assigning cheaper positions first is always the optimal strategy.

---

### Language Differences

**C++**

Uses `vector<int>` and `sort()` with `greater<int>()`.

**Java**

Uses an integer array and `Arrays.sort()`. Since Java sorts in ascending order, the frequencies are processed from the end.

**JavaScript**

Uses an array and a custom comparator for descending sorting.

**Python3**

Uses a list of size 26 and `sort(reverse=True)`.

**Go**

Uses a slice and `sort.Slice()` with a custom comparison function.

Although the syntax changes, every implementation follows exactly the same greedy algorithm.

---

## Examples

### Example 1

**Input**

```
word = "abcde"
```

**Output**

```
5
```

**Explanation**

Every character appears once.

The first five letters can all be placed in one-push positions.

```
1 + 1 + 1 + 1 + 1 = 5
```

---

### Example 2

**Input**

```
word = "xyzxyzxyzxyz"
```

**Output**

```
12
```

**Explanation**

Each of the three letters appears four times.

All three letters can occupy one-push positions.

```
4 × 1 + 4 × 1 + 4 × 1 = 12
```

---

### Example 3

**Input**

```
word = "aabbccddeeffgghhiiiiii"
```

**Output**

```
24
```

**Explanation**

The most frequent letter receives the cheapest position.

The remaining letters are assigned according to their frequencies, producing the smallest possible total cost.

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

Run

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

- The greedy approach is the optimal solution.
- Only letter frequencies matter. The actual characters never affect the answer.
- Since there are only 26 lowercase letters, sorting is effectively constant time.
- The algorithm easily handles the maximum input size of `10^5`.
- Memory usage stays constant because only a frequency array of size 26 is used.
- No hash map is required since the alphabet size is fixed.
- This is one of the cleanest greedy problems on LeetCode because the optimal assignment naturally follows from sorting the frequencies.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
