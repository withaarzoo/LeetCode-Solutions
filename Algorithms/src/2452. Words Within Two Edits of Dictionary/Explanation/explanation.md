# 2452. Words Within Two Edits of Dictionary

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

We are given two string arrays:

* `queries`
* `dictionary`

All words have the same length.

For each word in `queries`, we need to check whether it can become any word in `dictionary` using at most two character changes.

If yes, we include that query word in the final answer.

The answer should preserve the same order as the original `queries` array.

---

## Constraints

```text
1 <= queries.length, dictionary.length <= 100
1 <= queries[i].length == dictionary[j].length <= 100
All words contain only lowercase English letters.
```

---

## Intuition

I thought about comparing every query word with every dictionary word.

Since all words have the same length, I can simply count how many positions have different characters.

* If the difference count is `0`, `1`, or `2`, then the query word can be transformed into that dictionary word in at most two edits.
* If the difference becomes greater than `2`, then there is no need to continue checking that pair.

So, for every query word, I try all dictionary words until I find one that matches within two edits.

---

## Approach

1. Create an empty answer list.
2. Loop through every word in `queries`.
3. For each query word:

   * Compare it with every word in `dictionary`.
   * Count how many characters are different.
   * If the difference count is more than `2`, stop checking that pair early.
4. If any dictionary word differs by at most `2` characters:

   * Add the query word to the result.
   * Move to the next query word.
5. Return the final result.

---

## Data Structures Used

* Array / List for storing the final result.
* Nested loops for comparing query words and dictionary words.
* Integer variable for counting character differences.

---

## Operations & Behavior Summary

| Operation         | Description                                              |
| ----------------- | -------------------------------------------------------- |
| Compare words     | Compare each query word with every dictionary word       |
| Count differences | Count character mismatches between two words             |
| Early stopping    | Stop checking once mismatch count becomes greater than 2 |
| Add result        | Add query word if a valid dictionary match is found      |

---

## Complexity

* Time Complexity: `O(q * d * n)`

  * `q` = number of query words
  * `d` = number of dictionary words
  * `n` = length of each word

* Space Complexity: `O(1)`

  * Ignoring the output array, only a few variables are used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<string> twoEditWords(vector<string>& queries, vector<string>& dictionary) {
        vector<string> result;

        for (string query : queries) {
            for (string word : dictionary) {
                int diff = 0;

                for (int i = 0; i < query.size(); i++) {
                    if (query[i] != word[i]) {
                        diff++;
                    }

                    if (diff > 2) {
                        break;
                    }
                }

                if (diff <= 2) {
                    result.push_back(query);
                    break;
                }
            }
        }

        return result;
    }
};
```

### Java

```java
class Solution {
    public List<String> twoEditWords(String[] queries, String[] dictionary) {
        List<String> result = new ArrayList<>();

        for (String query : queries) {
            for (String word : dictionary) {
                int diff = 0;

                for (int i = 0; i < query.length(); i++) {
                    if (query.charAt(i) != word.charAt(i)) {
                        diff++;
                    }

                    if (diff > 2) {
                        break;
                    }
                }

                if (diff <= 2) {
                    result.add(query);
                    break;
                }
            }
        }

        return result;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string[]} queries
 * @param {string[]} dictionary
 * @return {string[]}
 */
var twoEditWords = function(queries, dictionary) {
    const result = [];

    for (const query of queries) {
        for (const word of dictionary) {
            let diff = 0;

            for (let i = 0; i < query.length; i++) {
                if (query[i] !== word[i]) {
                    diff++;
                }

                if (diff > 2) {
                    break;
                }
            }

            if (diff <= 2) {
                result.push(query);
                break;
            }
        }
    }

    return result;
};
```

### Python3

```python
class Solution:
    def twoEditWords(self, queries: List[str], dictionary: List[str]) -> List[str]:
        result = []

        for query in queries:
            for word in dictionary:
                diff = 0

                for i in range(len(query)):
                    if query[i] != word[i]:
                        diff += 1

                    if diff > 2:
                        break

                if diff <= 2:
                    result.append(query)
                    break

        return result
```

### Go

```go
func twoEditWords(queries []string, dictionary []string) []string {
    result := []string{}

    for _, query := range queries {
        for _, word := range dictionary {
            diff := 0

            for i := 0; i < len(query); i++ {
                if query[i] != word[i] {
                    diff++
                }

                if diff > 2 {
                    break
                }
            }

            if diff <= 2 {
                result = append(result, query)
                break
            }
        }
    }

    return result
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### Step 1: Create Result Array

We first create an empty array or list called `result`.

This will store all query words that can match any dictionary word within two edits.

### Step 2: Loop Through Each Query Word

We go through every word in `queries` one by one.

```text
query = "word"
```

### Step 3: Compare With Every Dictionary Word

For each query word, we compare it with all dictionary words.

```text
dictionary = ["wood", "joke", "moat"]
```

### Step 4: Count Character Differences

We compare characters at the same positions.

Example:

```text
word = "word"
dictionaryWord = "wood"
```

Character comparison:

```text
w == w
o == o
r != o
d == d
```

Only one character is different.

So:

```text
diff = 1
```

### Step 5: Stop Early If Difference Becomes Greater Than 2

If at any point `diff > 2`, then we already know this dictionary word cannot match.

So we stop early and move to the next dictionary word.

This small optimization helps reduce unnecessary comparisons.

### Step 6: Add Valid Query Word

If `diff <= 2`, then we add the query word into the result array.

```text
result.push_back(query)
```

Then we stop checking further dictionary words for that query.

### Step 7: Return Final Answer

After checking all query words, we return the result array.

---

## Examples

### Example 1

```text
Input:
queries = ["word", "note", "ants", "wood"]
dictionary = ["wood", "joke", "moat"]

Output:
["word", "note", "wood"]
```

Explanation:

* `word` differs from `wood` by 1 character.
* `note` differs from `joke` by 2 characters.
* `ants` differs by more than 2 characters from all dictionary words.
* `wood` exactly matches `wood`.

### Example 2

```text
Input:
queries = ["yes"]
dictionary = ["not"]

Output:
[]
```

Explanation:

* `yes` differs from `not` by more than 2 characters.
* So it cannot be included.

---

## How to use / Run locally

### C++

```bash
g++ main.cpp -o main
./main
```

### Java

```bash
javac Solution.java
java Solution
```

### JavaScript

```bash
node solution.js
```

### Python3

```bash
python solution.py
```

### Go

```bash
go run solution.go
```

---

## Notes & Optimizations

* All words have the same length, which makes character comparison simple.
* We stop early when the difference count becomes greater than `2`.
* This avoids unnecessary comparisons.
* The brute-force approach is already fast enough because constraints are small.
* Maximum possible operations are manageable within the problem limits.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
