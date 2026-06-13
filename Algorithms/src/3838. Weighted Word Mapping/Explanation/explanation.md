# 3838. Weighted Word Mapping

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

LeetCode 3838: Weighted Word Mapping is a simple string and array simulation problem.

We are given:

* An array of lowercase English words.
* A weight array of length 26 where each index represents the weight of a lowercase letter.

For every word:

1. Calculate the total weight of all characters.
2. Take the total weight modulo 26.
3. Convert the result into a character using reverse alphabetical order:

   * `0 → z`
   * `1 → y`
   * ...
   * `25 → a`
4. Append the mapped character to the answer.

The goal is to return the final string formed by concatenating all mapped characters in order.

This problem mainly tests string processing, array indexing, simulation, and character mapping.

---

## Constraints

| Constraint                                    | Value                 |
| --------------------------------------------- | --------------------- |
| `1 <= words.length <= 100`                    | Number of words       |
| `1 <= words[i].length <= 10`                  | Length of each word   |
| `weights.length == 26`                        | Fixed alphabet size   |
| `1 <= weights[i] <= 100`                      | Weight of each letter |
| `words[i]` contains lowercase English letters | Valid input format    |

---

## Intuition

The first thing I noticed was that every word contributes exactly one character to the final answer.

To generate that character, I only need the total weight of the word.

Since every lowercase letter already has a predefined weight, I can simply:

* Visit every character in the word.
* Add its corresponding weight.
* Compute modulo 26.
* Convert the result into a reverse alphabetical character.

Because the alphabet size is fixed at 26, the mapping becomes very straightforward.

Instead of building a lookup table, I can directly calculate the character using simple character arithmetic.

---

## Approach

1. Create an empty result string.
2. Process every word one by one.
3. Calculate the total weight of the current word.
4. Use the character index:

   * `'a' → 0`
   * `'b' → 1`
   * ...
   * `'z' → 25`
5. Add the corresponding weight from the weights array.
6. Compute:

   * `totalWeight % 26`
7. Convert the result to a reverse alphabet character.
8. Append the character to the answer.
9. Return the final string after all words are processed.

This gives an efficient and clean solution.

---

## Data Structures Used

### String

Used to store the final answer.

### Integer Variables

Used for:

* Tracking word weight sums
* Storing modulo results

### Array

The `weights` array provides O(1) access to every character weight.

No additional complex data structures are required.

---

## Operations & Behavior Summary

The algorithm performs the following operations:

1. Start with an empty answer.
2. Read each word.
3. Traverse every character inside the word.
4. Convert the character into its alphabet index.
5. Add the corresponding weight.
6. Compute modulo 26.
7. Convert the modulo value into a reverse alphabet character.
8. Add the character to the answer.
9. Continue until all words are processed.
10. Return the generated string.

This is a pure simulation approach with constant-time character mapping.

---

## Complexity

| Metric           | Complexity | Explanation                                      |
| ---------------- | ---------- | ------------------------------------------------ |
| Time Complexity  | `O(n × k)` | `n` = number of words, `k` = average word length |
| Space Complexity | `O(n)`     | Output string stores one character per word      |

Since every character is visited exactly once, this is already optimal.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string mapWordWeights(vector<string>& words, vector<int>& weights) {
        string result;
        
        // Process each word independently
        for (const string& word : words) {
            int sumWeight = 0;
            
            // Add the weight of every character
            for (char ch : word) {
                sumWeight += weights[ch - 'a'];
            }
            
            // Reduce the weight into range [0, 25]
            int value = sumWeight % 26;
            
            // Reverse mapping:
            // 0 -> z, 1 -> y, ..., 25 -> a
            result.push_back(char('z' - value));
        }
        
        return result;
    }
};
```

### Java

```java
class Solution {
    public String mapWordWeights(String[] words, int[] weights) {
        StringBuilder result = new StringBuilder();
        
        // Process every word
        for (String word : words) {
            int sumWeight = 0;
            
            // Add weights of all characters
            for (char ch : word.toCharArray()) {
                sumWeight += weights[ch - 'a'];
            }
            
            // Take modulo 26
            int value = sumWeight % 26;
            
            // Convert to reverse alphabetical character
            result.append((char)('z' - value));
        }
        
        return result.toString();
    }
}
```

### JavaScript

```javascript
/**
 * @param {string[]} words
 * @param {number[]} weights
 * @return {string}
 */
var mapWordWeights = function(words, weights) {
    let result = "";
    
    // Process every word
    for (const word of words) {
        let sumWeight = 0;
        
        // Add character weights
        for (const ch of word) {
            sumWeight += weights[ch.charCodeAt(0) - 97];
        }
        
        // Reduce into range [0, 25]
        const value = sumWeight % 26;
        
        // Reverse alphabet mapping
        result += String.fromCharCode('z'.charCodeAt(0) - value);
    }
    
    return result;
};
```

### Python3

```python
class Solution:
    def mapWordWeights(self, words: List[str], weights: List[int]) -> str:
        result = []
        
        # Process each word
        for word in words:
            sum_weight = 0
            
            # Add weights of all characters
            for ch in word:
                sum_weight += weights[ord(ch) - ord('a')]
            
            # Take modulo 26
            value = sum_weight % 26
            
            # Reverse alphabetical mapping
            result.append(chr(ord('z') - value))
        
        return "".join(result)
```

### Go

```go
func mapWordWeights(words []string, weights []int) string {
 var result []byte

 // Process each word
 for _, word := range words {
  sumWeight := 0

  // Add weights of all characters
  for _, ch := range word {
   sumWeight += weights[ch-'a']
  }

  // Reduce into range [0, 25]
  value := sumWeight % 26

  // Reverse alphabet mapping:
  // 0 -> z, 1 -> y, ..., 25 -> a
  result = append(result, byte('z'-value))
 }

 return string(result)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains identical across all five languages.

### Step 1: Initialize the Result

Create an empty string (or string builder equivalent) that will store the final answer.

Since each word contributes one character, the result size will eventually equal the number of words.

---

### Step 2: Process Each Word

Loop through every word in the input array.

Each word is handled independently.

No information from previous words is needed.

---

### Step 3: Calculate Word Weight

For the current word:

* Visit every character.
* Convert it into its alphabet index.
* Use that index to access the weight array.

Example:

```text
a → index 0
b → index 1
c → index 2
```

Add all weights together.

This gives the total weight of the word.

---

### Step 4: Apply Modulo 26

The problem requires:

```text
weight % 26
```

This guarantees a value between:

```text
0 and 25
```

which perfectly matches the alphabet size.

---

### Step 5: Reverse Alphabet Mapping

The problem uses reverse alphabetical order.

Mapping:

```text
0 → z
1 → y
2 → x
...
25 → a
```

Instead of creating a lookup table, direct character arithmetic can be used.

This reduces memory usage and keeps the implementation simple.

---

### Step 6: Append Character

Add the mapped character to the result.

Since each word generates exactly one character, the answer grows by one position after every iteration.

---

### Step 7: Return Final String

After all words are processed, return the completed string.

This contains the mapped character for every word in the same order as the input.

---

## Examples

### Example 1

#### Input

```text
words = ["abcd","def","xyz"]

weights = [5,3,12,14,1,2,3,2,10,6,6,9,7,8,7,10,8,9,6,9,9,8,3,7,7,2]
```

#### Output

```text
"rij"
```

#### Explanation

```text
abcd → 34 → 34 % 26 = 8 → r
def  → 17 → 17 % 26 = 17 → i
xyz  → 16 → 16 % 26 = 16 → j
```

Final answer:

```text
"rij"
```

---

### Example 2

#### Input

```text
words = ["a","b","c"]
weights = [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
```

#### Output

```text
"yyy"
```

#### Explanation

```text
1 % 26 = 1
1 → y
```

Each word maps to `y`.

---

### Example 3

#### Input

```text
words = ["abcd"]
```

#### Output

```text
"g"
```

#### Explanation

```text
Weight = 19
19 % 26 = 19
19 → g
```

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

* The alphabet size is fixed at 26.
* Weight lookup is always O(1).
* No sorting is required.
* No hashing is required.
* No extra preprocessing is needed.
* Character arithmetic is faster and cleaner than creating a reverse lookup table.
* The solution already achieves optimal time complexity because every character must be examined at least once.
* This problem is a good example of string simulation and array indexing techniques commonly seen in coding interviews and competitive programming contests.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
