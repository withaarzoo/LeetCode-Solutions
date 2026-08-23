# 1927. Sum Game - LeetCode Solution

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
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

This repository contains an optimized solution for LeetCode problem 1927, Sum Game.

Alice and Bob take turns replacing `?` characters in an even-length string with digits from `0` to `9`. Alice starts first.

The string is divided into two equal halves. After every `?` has been replaced, Bob wins if the sum of digits in the left half is equal to the sum of digits in the right half. Alice wins if the two sums are different.

The goal is to determine whether Alice can win when both players make optimal moves.

The input is a string containing digits and `?` characters. The output is a boolean value:

* `true` if Alice can force a win
* `false` if Bob can force the two sums to become equal

The main challenge is avoiding game simulation. The number of possible moves can grow quickly, but this Sum Game problem can be reduced to a simple mathematical condition.

## Constraints

| Constraint                    | Description                                                           |
| ----------------------------- | --------------------------------------------------------------------- |
| `2 <= num.length <= 10^5`     | The input string can contain up to 100,000 characters.                |
| `num.length` is even          | The string can always be divided into two equal halves.               |
| `num` contains digits and `?` | Every character is either a digit from `0` to `9` or a question mark. |

## Intuition

My first thought was to look at the game move by move. But that would mean considering many possible digits for every `?`, which is not practical.

I noticed that the exact position of a `?` inside the same half does not matter. Only two things matter for each half:

* The sum of the digits already present
* The number of `?` characters

So I only need to calculate the current sum difference and compare it with the effect that the remaining question marks can have.

A `?` can be replaced with any value from `0` to `9`, so one unmatched question mark can create a maximum adjustment of `9`.

After simplifying the game logic, I found that Bob can force both sums to become equal only in one exact situation:

`2 × (leftSum - rightSum) = 9 × (rightQuestion - leftQuestion)`

If this equality is not possible, Alice can force the final sums to be different.

This turns the problem into a linear scan instead of a game simulation.

## Approach

I use the following steps:

1. Divide the string into a left half and a right half.
2. Calculate the sum of known digits in the left half.
3. Calculate the sum of known digits in the right half.
4. Count the number of `?` characters in both halves.
5. Compare the current sum difference with the difference created by the question marks.
6. Check whether Bob can satisfy the exact equality condition.
7. If Bob can force equal sums, return `false`.
8. Otherwise, Alice can force different sums, so return `true`.

The final condition is:

`2 * (leftSum - rightSum) != 9 * (rightQuestion - leftQuestion)`

If this expression is true, Alice wins.

I use multiplication instead of division so the entire calculation stays in integers and there are no fraction or floating-point issues.

## Data Structures Used

No complex data structures are needed for this Sum Game solution.

| Variable Type     | Purpose                                           |
| ----------------- | ------------------------------------------------- |
| Integer variables | Store the digit sum of each half                  |
| Integer variables | Count `?` characters in each half                 |
| Input string      | Provides the digits and question marks to process |

I do not use arrays, hash maps, stacks, queues, or recursion because the answer depends only on four values: two sums and two question mark counts.

## Operations & Behavior Summary

The algorithm works like this:

1. Find the middle index of the string.
2. Start scanning from left to right.
3. If the current character is in the left half:

   * Add its value to the left sum if it is a digit.
   * Increase the left question mark count if it is `?`.
4. If the current character is in the right half:

   * Add its value to the right sum if it is a digit.
   * Increase the right question mark count if it is `?`.
5. After scanning the entire string, compare:

   * The difference between the two known sums
   * The possible balancing effect created by the difference in `?` counts
6. If Bob can make the final sums exactly equal, Alice loses.
7. Otherwise, Alice has a winning strategy.

This is the complete logic behind the optimized LeetCode 1927 Sum Game solution.

## Complexity

| Complexity       | Value  | Explanation                                                                  |
| ---------------- | ------ | ---------------------------------------------------------------------------- |
| Time Complexity  | `O(n)` | `n` is the length of `num`. I scan every character exactly once.             |
| Space Complexity | `O(1)` | I use only a fixed number of integer variables and no extra data structures. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool sumGame(string num) {
        // Find the middle because the string always has even length.
        int mid = num.size() / 2;

        // Store the sums of known digits in both halves.
        int leftSum = 0, rightSum = 0;

        // Count how many '?' characters exist in both halves.
        int leftQuestion = 0, rightQuestion = 0;

        // Scan every character once and update the correct half.
        for (int i = 0; i < num.size(); i++) {
            if (i < mid) {
                // This character belongs to the left half.
                if (num[i] == '?')
                    leftQuestion++;
                else
                    leftSum += num[i] - '0';
            } else {
                // This character belongs to the right half.
                if (num[i] == '?')
                    rightQuestion++;
                else
                    rightSum += num[i] - '0';
            }
        }

        // Bob can force equality only when this exact condition is true.
        // If equality is impossible, Alice can force the sums to be different.
        return 2 * (leftSum - rightSum) !=
               9 * (rightQuestion - leftQuestion);
    }
};
```

### Java

```java
class Solution {
    public boolean sumGame(String num) {
        // Find the starting index of the right half.
        int mid = num.length() / 2;

        // Store the sums of known digits in both halves.
        int leftSum = 0;
        int rightSum = 0;

        // Count '?' characters in both halves.
        int leftQuestion = 0;
        int rightQuestion = 0;

        // Visit every character and update its corresponding half.
        for (int i = 0; i < num.length(); i++) {
            if (i < mid) {
                // This character is in the left half.
                if (num.charAt(i) == '?') {
                    leftQuestion++;
                } else {
                    leftSum += num.charAt(i) - '0';
                }
            } else {
                // This character is in the right half.
                if (num.charAt(i) == '?') {
                    rightQuestion++;
                } else {
                    rightSum += num.charAt(i) - '0';
                }
            }
        }

        // Bob wins only when he can make both final sums exactly equal.
        // Otherwise, Alice can force the sums to stay different.
        return 2 * (leftSum - rightSum) !=
               9 * (rightQuestion - leftQuestion);
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} num
 * @return {boolean}
 */
var sumGame = function(num) {
    // Find where the right half starts.
    const mid = num.length / 2;

    // Store sums of known digits and counts of '?' characters.
    let leftSum = 0;
    let rightSum = 0;
    let leftQuestion = 0;
    let rightQuestion = 0;

    // Scan every character once.
    for (let i = 0; i < num.length; i++) {
        if (i < mid) {
            // Update information for the left half.
            if (num[i] === '?') {
                leftQuestion++;
            } else {
                leftSum += Number(num[i]);
            }
        } else {
            // Update information for the right half.
            if (num[i] === '?') {
                rightQuestion++;
            } else {
                rightSum += Number(num[i]);
            }
        }
    }

    // If Bob cannot satisfy this equality, Alice can force a win.
    return 2 * (leftSum - rightSum) !==
           9 * (rightQuestion - leftQuestion);
};
```

### Python3

```python
class Solution:
    def sumGame(self, num: str) -> bool:
        # Find the index where the right half begins.
        mid = len(num) // 2

        # Store the sums of known digits in both halves.
        left_sum = 0
        right_sum = 0

        # Count how many '?' characters are in both halves.
        left_question = 0
        right_question = 0

        # Scan every character and update the correct half.
        for i, char in enumerate(num):
            if i < mid:
                # This character belongs to the left half.
                if char == '?':
                    left_question += 1
                else:
                    left_sum += int(char)
            else:
                # This character belongs to the right half.
                if char == '?':
                    right_question += 1
                else:
                    right_sum += int(char)

        # Bob can force equality only in this exact situation.
        # Otherwise, Alice can always make the final sums different.
        return 2 * (left_sum - right_sum) != \
               9 * (right_question - left_question)
```

### Go

```go
func sumGame(num string) bool {
 // Find where the right half starts.
 mid := len(num) / 2

 // Store sums of known digits in both halves.
 leftSum, rightSum := 0, 0

 // Count '?' characters in both halves.
 leftQuestion, rightQuestion := 0, 0

 // Scan every character once.
 for i := 0; i < len(num); i++ {
  if i < mid {
   // Update information for the left half.
   if num[i] == '?' {
    leftQuestion++
   } else {
    leftSum += int(num[i] - '0')
   }
  } else {
   // Update information for the right half.
   if num[i] == '?' {
    rightQuestion++
   } else {
    rightSum += int(num[i] - '0')
   }
  }
 }

 // If Bob cannot make this equality true, Alice wins.
 return 2*(leftSum-rightSum) != 9*(rightQuestion-leftQuestion)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The same algorithm is used in C++, Java, JavaScript, Python3, and Go. The syntax changes between languages, but the logic and complexity remain the same.

First, I find the middle of the string. Since the problem guarantees an even-length string, the middle index cleanly separates the left and right halves.

For example:

`num = "5023"`

The two halves are:

* Left half: `"50"`
* Right half: `"23"`

Next, I keep track of four values.

`leftSum` stores the total of all known digits in the left half.

`rightSum` stores the total of all known digits in the right half.

`leftQuestion` stores the number of `?` characters in the left half.

`rightQuestion` stores the number of `?` characters in the right half.

I then scan the string one character at a time.

If a character belongs to the left half, I update either `leftSum` or `leftQuestion`.

If it belongs to the right half, I update either `rightSum` or `rightQuestion`.

When the scan is finished, I have reduced the entire game to four numbers.

The most important part is the final comparison.

Bob wants the final sums to be equal. Alice wants them to be different.

A question mark can be assigned a digit from `0` through `9`. Because of optimal play, Bob can force equality only when the current sum difference and the question mark difference satisfy this exact equation:

`2 * (leftSum - rightSum) = 9 * (rightQuestion - leftQuestion)`

The multiplication by `2` avoids division and keeps the calculation completely integer-based.

If the equality holds, Bob can force equal sums. Since the function asks whether Alice wins, the answer is `false`.

If the equality does not hold, Bob cannot guarantee equality. Alice can choose values that keep the two final sums different, so the answer is `true`.

The language-specific behavior is straightforward:

### C++

I use `string` for the input and integer variables for sums and counts. Character digits are converted using subtraction from `'0'`.

### Java

I use `String` and access characters with `charAt()`. A digit character is converted to its numeric value by subtracting `'0'`.

### JavaScript

I use a string and numeric variables. Digit characters can be converted using `Number()`.

### Python3

I iterate through the string and convert digit characters using `int()`. Python integers handle all required values safely.

### Go

I scan the string using indexes. Since the input contains only ASCII digits and `?`, byte indexing works correctly. Digit bytes are converted by subtracting `'0'`.

## Examples

### Example 1

**Input:**

`num = "5023"`

**Expected Output:**

`false`

**Trace:**

* Left half: `"50"`
* Right half: `"23"`
* Left sum: `5 + 0 = 5`
* Right sum: `2 + 3 = 5`
* There are no `?` characters

The two sums are already equal, and Alice has no move available.

Bob wins, so the answer is `false`.

### Example 2

**Input:**

`num = "25??"`

**Expected Output:**

`true`

**Trace:**

* Left half: `"25"`
* Right half: `"??"`
* Left sum: `7`
* Right sum: `0`
* Left question marks: `0`
* Right question marks: `2`

The mathematical condition for Bob to force equality is not satisfied.

Alice can choose a value that prevents the two final sums from becoming equal.

The answer is `true`.

### Example 3

**Input:**

`num = "??"`

**Expected Output:**

`false`

**Trace:**

* Left sum: `0`
* Right sum: `0`
* Left question marks: `1`
* Right question marks: `1`

Both sides have equal starting sums and equal numbers of question marks.

Bob can always respond in a way that keeps the final sums balanced.

The answer is `false`.

## How to Use / Run Locally

Create a file for the language you want to test and paste the corresponding solution into it.

### C++

Save the file as `main.cpp`.

Compile:

```bash
g++ -std=c++17 main.cpp -o main
```

Run:

```bash
./main
```

On Windows:

```bash
main.exe
```

### Java

Save the file as `Solution.java`.

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

If you are testing the code locally, you may need to add your own `main` method because the LeetCode solution format only contains the `Solution` class.

### JavaScript

Save the file as `solution.js`.

Run:

```bash
node solution.js
```

For local testing, add your own test calls below the solution function.

### Python3

Save the file as `solution.py`.

Run:

```bash
python3 solution.py
```

On some Windows systems:

```bash
python solution.py
```

You may need to add a small test section to call the `Solution` class locally.

### Go

Save the file as `main.go`.

Run:

```bash
go run main.go
```

For a standalone executable:

```bash
go build main.go
```

Then run the generated executable.

## Notes & Optimizations

The biggest optimization is avoiding simulation.

Trying every possible digit for every `?` would create a huge number of possible game states. With up to `10^5` characters, that approach is impossible.

This solution only needs one pass through the string.

The positions of question marks inside the same half are not important. Only their count matters because every question mark has the same allowed digit range.

Using the mathematical condition also avoids recursion, memoization, minimax, and dynamic programming.

An important edge case is when there are no question marks. In that case, the result depends only on whether the two existing sums are already equal.

Another important case is when both halves have the same number of `?` characters. Then Bob can balance the game only when the known digit sums already match.

The solution is optimal with `O(n)` time and `O(1)` extra space, which makes it suitable for the maximum constraint of `10^5`.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
