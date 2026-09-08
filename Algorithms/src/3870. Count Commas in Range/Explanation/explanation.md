# 3870. Count Commas in Range

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

The problem asks us to find the total number of commas used when writing every integer from `1` to `n` using standard number formatting.

A comma is added after every three digits from the right.

For example:

```text
999       -> 999
1,000     -> 1 comma
12,345    -> 1 comma
1,000,000 -> 2 commas
```

Given an integer `n`, I need to return the total number of commas used for all numbers from `1` through `n`.

For example, when `n = 1002`, the numbers `1000`, `1001`, and `1002` each contain one comma.

So the answer is:

```text
3
```

This is a simple counting problem, but the key is to avoid checking every number one by one.

## Constraints

| Constraint       | Description                                                 |
| ---------------- | ----------------------------------------------------------- |
| `1 <= n <= 10^5` | `n` is a positive integer                                   |
| Range            | All integers from `1` to `n` are considered                 |
| Formatting       | A comma is inserted after every three digits from the right |

## Intuition

My first thought was to format every number from `1` to `n` and count its commas. That would work, but it is unnecessary.

I noticed that commas start appearing at specific number ranges:

```text
1 to 999
    -> 0 commas

1,000 to 999,999
    -> 1 comma

1,000,000 to 999,999,999
    -> 2 commas
```

So I can count entire ranges instead of processing every number.

For example, if `n = 1002`:

```text
1 ... 999 | 1000 | 1001 | 1002
  no comma |  1  |  1   |  1
```

There are exactly `3` numbers with commas.

The same idea works for larger values. Every time I reach another power of `1000`, each number from that point onward gets one additional comma.

## Approach

I start with `1000`, because this is the first number that contains a comma.

For every threshold, I count how many numbers from that threshold through `n` exist.

The number of values in this range is:

```text
n - threshold + 1
```

I add that to the answer because every number in this range gets one additional comma.

Then I multiply the threshold by `1000` and repeat.

The thresholds look like this:

```text
1000
  |
  | × 1000
  v
1,000,000
  |
  | × 1000
  v
1,000,000,000
```

For example, with `n = 1,002,005`:

```text
Threshold        Numbers counted
-----------------------------------------
1,000            1,002,005 - 1,000 + 1
1,000,000        1,002,005 - 1,000,000 + 1
```

The first threshold counts the first comma in every number from `1,000` onward.

The second threshold counts the extra comma present in every number from `1,000,000` onward.

This means a number such as `1,002,005` is counted twice, which is correct because it contains two commas.

## Data Structures Used

No data structure is required for this solution.

I only use:

* `answer` to store the total number of commas.
* `threshold` to represent the current comma range.

Both are simple integer variables, so the solution uses constant extra space.

## Operations & Behavior Summary

The algorithm works like this:

1. Set `answer = 0`.
2. Set `threshold = 1000`.
3. While `threshold <= n`:

   * Count the numbers from `threshold` to `n`.
   * Add that count to `answer`.
   * Multiply `threshold` by `1000`.
4. Return `answer`.

In simple pseudocode:

```text
answer = 0
threshold = 1000

while threshold <= n:
    answer += n - threshold + 1
    threshold *= 1000

return answer
```

The important observation is that I am counting comma positions rather than formatting individual numbers.

## Complexity

| Complexity       | Analysis        |
| ---------------- | --------------- |
| Time Complexity  | `O(log₁₀₀₀(n))` |
| Space Complexity | `O(1)`          |

The time complexity is `O(log₁₀₀₀(n))` because the threshold is multiplied by `1000` after every iteration.

Here, `n` is the given upper limit. Since the threshold grows very quickly, only a few iterations are needed.

The space complexity is `O(1)` because I only use a few integer variables. No array, hash map, string, or other extra data structure is needed.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int countCommas(int n) {
        long long answer = 0; // Stores the total number of commas.
        long long threshold = 1000; // First number that contains a comma.

        while (threshold <= n) {
            // Every number from threshold through n has one comma
            // corresponding to this comma group.
            answer += n - threshold + 1;

            // Move to the next comma level: 1,000 -> 1,000,000 -> 1,000,000,000.
            threshold *= 1000;
        }

        // The problem's answer fits in an integer for the given constraints.
        return static_cast<int>(answer);
    }
};
```

### Java

```java
class Solution {
    public int countCommas(int n) {
        long answer = 0; // Stores the total number of commas.
        long threshold = 1000; // First number that contains a comma.

        while (threshold <= n) {
            // Every number from threshold through n contributes
            // one comma for this comma group.
            answer += n - threshold + 1;

            // Move to the next comma level.
            threshold *= 1000;
        }

        // The answer fits in an int for the given constraints.
        return (int) answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @return {number}
 */
var countCommas = function(n) {
    let answer = 0; // Stores the total number of commas.
    let threshold = 1000; // First number that contains a comma.

    while (threshold <= n) {
        // Every number from threshold through n contributes
        // one comma for this comma group.
        answer += n - threshold + 1;

        // Move to the next comma level.
        threshold *= 1000;
    }

    // Return the total number of commas.
    return answer;
};
```

### Python3

```python
class Solution:
    def countCommas(self, n: int) -> int:
        answer = 0  # Stores the total number of commas.
        threshold = 1000  # First number that contains a comma.

        while threshold <= n:
            # Every number from threshold through n contributes
            # one comma for this comma group.
            answer += n - threshold + 1

            # Move to the next comma level.
            threshold *= 1000

        # Return the total number of commas.
        return answer
```

### Go

```go
func countCommas(n int) int {
 answer := 0       // Stores the total number of commas.
 threshold := 1000 // First number that contains a comma.

 for threshold <= n {
  // Every number from threshold through n contributes
  // one comma for this comma group.
  answer += n - threshold + 1

  // Move to the next comma level.
  threshold *= 1000
 }

 // Return the total number of commas.
 return answer
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five languages. Only the syntax changes.

### C++

I start with an integer variable called `answer` and set it to `0`.

This variable stores the total number of commas found so far.

Next, I create `threshold` and set it to `1000`.

I use `1000` because all numbers below `1000` have fewer than four digits and therefore contain no commas.

Then I run a loop while `threshold` is less than or equal to `n`.

Inside the loop, I calculate:

```text
n - threshold + 1
```

This tells me how many numbers are between `threshold` and `n`, including both endpoints.

I add this value to `answer`.

After that, I multiply `threshold` by `1000`.

This moves the calculation to the next comma level.

I use `long long` for the intermediate calculation to keep the arithmetic safe.

Finally, I return the answer as an integer because the answer fits within the required range.

### Java

The Java solution follows the same mathematical approach.

I keep the total in a `long` variable called `answer`.

The threshold starts at `1000`.

Inside the loop, I calculate how many numbers are at least as large as the current threshold:

```text
n - threshold + 1
```

That entire group contributes one additional comma.

Then I multiply the threshold by `1000` to move to the next group.

Once the threshold becomes larger than `n`, there are no more comma levels to process.

The final answer is returned as an `int`.

### JavaScript

In JavaScript, I use two variables:

```text
answer
threshold
```

`answer` stores the running total, while `threshold` starts at `1000`.

The loop checks whether the current threshold is still inside the range.

For each valid threshold, I add:

```text
n - threshold + 1
```

Then I multiply the threshold by `1000`.

JavaScript's `Number` type is enough for the constraints in this problem.

At the end, `answer` contains the total number of commas.

### Python3

The Python solution uses the same two variables.

I initialize:

```text
answer = 0
threshold = 1000
```

The loop continues while the threshold is not greater than `n`.

For every threshold, I count all numbers from that threshold to `n`.

Then I increase the threshold by multiplying it by `1000`.

Python integers automatically handle large integer values, so I do not need to choose a separate integer type for overflow protection.

Finally, I return `answer`.

### Go

In Go, I again use `answer` and `threshold`.

The threshold starts at `1000`.

The `for` loop continues while the threshold is less than or equal to `n`.

For every valid threshold, I calculate:

```text
n - threshold + 1
```

and add it to the total.

Then I multiply the threshold by `1000` to move to the next comma level.

Once the threshold becomes larger than `n`, the loop ends and I return the answer.

## Examples

### Example 1

Input:

```text
n = 1002
```

The numbers containing commas are:

```text
1000 -> 1,000
1001 -> 1,001
1002 -> 1,002
```

There are three such numbers.

```text
1000 ---------------- 1002
  |                     |
  +--- 3 numbers -------+
```

Output:

```text
3
```

### Example 2

Input:

```text
n = 998
```

Every number from `1` to `998` has fewer than four digits.

```text
1 ---------------- 998
      0 commas
```

So no comma is used.

Output:

```text
0
```

### Example 3

Input:

```text
n = 1,000,000
```

First, every number from `1,000` to `1,000,000` contributes one comma.

```text
1,000 ---------------- 1,000,000
       1 comma each
```

Then `1,000,000` contains a second comma:

```text
1,000,000
    ^    ^
    two commas
```

So the algorithm counts the first comma using the `1000` threshold and the second comma using the `1,000,000` threshold.

The result is:

```text
1,000,000 - 1,000 + 1 = 999,001
1,000,000 - 1,000,000 + 1 = 1

Total = 999,002
```

Output:

```text
999002
```

## How to Use / Run Locally

The code on this README is intended to match the LeetCode `Solution` format. To run it locally, I can place the function inside a small program with a `main` function and test it with different values of `n`.

### C++

Save the solution in a `.cpp` file.

Compile it with:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it with:

```bash
./solution
```

### Java

Save the solution in a Java file.

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

For local testing, I would add a `main` method and call `countCommas()` with different inputs.

### JavaScript

Save the solution in a `.js` file.

Run it with Node.js:

```bash
node solution.js
```

I can add `console.log()` statements to test values such as `998`, `1002`, and `1000000`.

### Python3

Save the solution in a `.py` file.

Run it with:

```bash
python3 solution.py
```

I can create an instance of the `Solution` class and call `countCommas()` for local testing.

### Go

Save the solution in a `.go` file.

Run it with:

```bash
go run solution.go
```

For local testing, I can add a `main` function and print the result for different test cases.

## Notes & Optimizations

The most important optimization is avoiding a loop from `1` to `n`.

A direct solution could format every number and count its commas, but that would take `O(n)` time and perform unnecessary string operations.

Instead, I only process the comma thresholds:

```text
1000
1000000
1000000000
...
```

Each threshold represents one additional comma.

There are a few important edge cases:

* If `n < 1000`, the answer is `0`.
* If `n = 1000`, the answer is `1`.
* If `n = 999`, the answer is `0`.
* If `n = 1,000,000`, the number `1,000,000` contributes two commas.
* The `+1` in `n - threshold + 1` is necessary because the range includes both `threshold` and `n`.

The main tradeoff is that this solution relies on recognizing the pattern in standard number formatting. Once that pattern is clear, there is no need for string conversion or individual number processing.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
