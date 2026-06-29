# 1967. Number of Strings That Appear as Substrings in Word

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

This problem asks us to count how many strings from an array of patterns appear as a substring in a given word.

A substring means a continuous part of a string. So the task is simple: for each pattern, check whether it exists inside `word`, and count it if it does.

The input gives:

* an array of strings called `patterns`
* one string called `word`

The output is:

* the number of patterns that are found inside `word` as substrings

This is a classic string matching problem, and the clean solution is to check each pattern one by one.

## Constraints

| Constraint             |                            Value |
| ---------------------- | -------------------------------: |
| Number of patterns     |    `1 <= patterns.length <= 100` |
| Length of each pattern | `1 <= patterns[i].length <= 100` |
| Length of word         |        `1 <= word.length <= 100` |
| Character set          |   Lowercase English letters only |

These limits are small, which makes a direct substring search easy and efficient.

## Intuition

My first thought was that I do not need anything fancy here.

The input size is small, and I only need to know whether each pattern appears inside the word. That means I can simply scan through the array and check every string with a normal substring search.

This works well because the problem is not asking for the longest match, the earliest match, or all match positions. It only asks for the count of matching patterns. So a straightforward pattern matching solution is enough.

## Approach

I go through every string in `patterns`.

For each pattern, I check whether it is present in `word`.

If it is present, I increase the answer by one.

I do this for all patterns, even if some patterns are repeated, because each entry in the array counts separately.

That is the whole idea. No extra sorting. No advanced preprocessing. Just direct substring checking.

## Data Structures Used

* **Array / list of strings**: to store the patterns.
* **Integer counter**: to store the final answer.

No hash set, trie, or prefix table is needed here because the constraints are small and the direct approach is already good enough.

## Operations & Behavior Summary

1. Start with answer = 0.
2. Read each pattern from the array.
3. Check whether the pattern exists inside `word`.
4. If it exists, add 1 to the answer.
5. After checking all patterns, return the answer.

That is the full behavior of the algorithm in plain English.

## Complexity

| Metric           |     Complexity | Explanation                                                                                                                                                            |
| ---------------- | -------------: | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n * m * k)` | `n` is the number of patterns, `m` is the length of `word`, and `k` is the average length of a pattern. Each pattern is checked against `word` using substring search. |
| Space Complexity |         `O(1)` | Only one counter is used. No extra data structure grows with input size.                                                                                               |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int numOfStrings(vector<string>& patterns, string word) {
        // Store the number of matching patterns
        int count = 0;

        // Check every pattern one by one
        for (string &pattern : patterns) {
            // If the pattern exists inside word, increase the answer
            if (word.find(pattern) != string::npos) {
                count++;
            }
        }

        // Return the total number of matching patterns
        return count;
    }
};
```

### Java

```java
class Solution {
    public int numOfStrings(String[] patterns, String word) {

        // Store the number of matching patterns
        int count = 0;

        // Check every pattern
        for (String pattern : patterns) {

            // indexOf() returns -1 if the substring does not exist
            if (word.indexOf(pattern) != -1) {
                count++;
            }
        }

        // Return the final answer
        return count;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string[]} patterns
 * @param {string} word
 * @return {number}
 */
var numOfStrings = function(patterns, word) {

    // Store the number of matching patterns
    let count = 0;

    // Check every pattern
    for (const pattern of patterns) {

        // includes() returns true if pattern is a substring
        if (word.includes(pattern)) {
            count++;
        }
    }

    // Return the total count
    return count;
};
```

### Python3

```python
class Solution:
    def numOfStrings(self, patterns: List[str], word: str) -> int:

        # Store the number of matching patterns
        count = 0

        # Check every pattern
        for pattern in patterns:

            # If pattern exists inside word, increase the answer
            if pattern in word:
                count += 1

        # Return the total number of matches
        return count
```

### Go

```go
func numOfStrings(patterns []string, word string) int {

    // Store the number of matching patterns
    count := 0

    // Check every pattern
    for _, pattern := range patterns {

        // strings.Contains() checks whether pattern exists in word
        if strings.Contains(word, pattern) {
            count++
        }
    }

    // Return the final answer
    return count
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in every language, so the reasoning stays consistent.

I start by creating a variable to store the answer. This value begins at zero because I have not checked any pattern yet.

Then I loop through the `patterns` array one item at a time. I do this because every pattern must be checked independently.

For each pattern, I search inside `word`. If the pattern is found anywhere in `word`, I count it. If not, I move on to the next one.

I do not stop early after one match because there may be more patterns that also appear in the word.

I also do not remove duplicates. If the same pattern appears twice in the input and both are substrings of `word`, both should be counted. That is why a simple loop is the correct choice.

At the end, I return the total count.

This logic is easy to understand and easy to trust because it follows the problem statement directly.

## Examples

### Example 1

**Input:**
`patterns = ["a", "abc", "bc", "d"]`
`word = "abc"`

**Output:**
`3`

**Trace:**

* `"a"` appears in `"abc"`
* `"abc"` appears in `"abc"`
* `"bc"` appears in `"abc"`
* `"d"` does not appear in `"abc"`

So the final answer is `3`.

### Example 2

**Input:**
`patterns = ["a", "b", "c"]`
`word = "aaaaabbbbb"`

**Output:**
`2`

**Trace:**

* `"a"` appears in the word
* `"b"` appears in the word
* `"c"` does not appear in the word

So the final answer is `2`.

### Example 3

**Input:**
`patterns = ["a", "a", "a"]`
`word = "ab"`

**Output:**
`3`

**Trace:**
Each `"a"` in the array is checked separately. Since `"a"` appears in `"ab"`, all three entries are counted.

## How to Use / Run Locally

For **C++**, save the solution in a `.cpp` file, compile it with a standard C++ compiler, and run the executable.

For **Java**, place the solution in a `.java` file, compile it with `javac`, and run it with `java`.

For **JavaScript**, save the code in a `.js` file and run it with Node.js.

For **Python3**, save the code in a `.py` file and run it with Python 3.

For **Go**, save the code in a `.go` file and run it with the Go toolchain.

Each version follows the same logic, so switching between languages is mostly about syntax, not strategy.

## Notes & Optimizations

This problem looks small, but it is a good example of choosing the right level of solution.

A direct substring search is the best fit here because:

* the input size is small
* the logic is simple
* the code stays readable
* the performance is already more than enough

A trie or advanced string algorithm would be unnecessary overhead for these constraints.

One thing to remember is that duplicates in `patterns` are counted separately. That is an important detail and often where mistakes happen.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
