# Unique 3-Digit Even Numbers

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

In this problem, I am given an array of digits and need to find how many distinct 3-digit even numbers I can form using those digits.

There are a few important rules:

* The number must have exactly 3 digits.
* The first digit cannot be `0`.
* The last digit must be even.
* Each copy of a digit can only be used once for a number.
* The answer should contain distinct numbers only.

For example, with:

```text
digits = [1, 2, 3, 4]
```

I can form numbers such as `124`, `132`, `214`, `234`, and so on.

The final answer is the number of unique 3-digit even numbers that can be created.

## Constraints

| Constraint                 | Meaning                                              |
| -------------------------- | ---------------------------------------------------- |
| `3 <= digits.length <= 10` | The input contains at least 3 and at most 10 digits. |
| `0 <= digits[i] <= 9`      | Every element is a single digit from `0` to `9`.     |

## Intuition

My first thought was to try every possible 3-digit number and check whether I can build it from the given digits.

Since there are only 10 possible digits, the search space is very small.

For a valid 3-digit even number, I have three positions:

```text
Hundreds     Tens       Ones
   1          2          3
   |          |          |
  1-9        0-9       0,2,4,6,8
```

The first position cannot contain `0`, while the last position must contain an even digit.

The main issue is duplicate digits. If the input contains two copies of `2`, I can use `2` twice in a number, but I cannot use it three times.

So I use a frequency array to store how many copies of each digit are available.

## Approach

I use a frequency array of size `10`.

For every digit from `0` to `9`, I store how many times it appears in the input.

Then I try every possible combination of three digit values:

1. Choose the first digit from `1` to `9`.
2. Choose the second digit from `0` to `9`.
3. Choose the third digit from `0, 2, 4, 6, 8`.
4. Check whether the selected digits are available in the required quantities.
5. If they are available, count this number.
6. Continue with the next combination.

For example:

```text
digits = [0, 2, 2]

Frequency:

+-------+---+---+---+---+---+
| Digit | 0 | 1 | 2 | 3 | ... |
+-------+---+---+---+---+---+
| Count | 1 | 0 | 2 | 0 | ... |
+-------+---+---+---+---+---+
```

Now:

```text
202  -> needs two 2s and one 0 -> valid
220  -> needs two 2s and one 0 -> valid
222  -> needs three 2s         -> invalid
```

Therefore, the answer is `2`.

Because I loop over digit values rather than positions in the input array, the same number is never counted multiple times.

## Data Structures Used

### Frequency Array

I use an integer array of size `10`.

```text
freq[0] -> number of zeroes
freq[1] -> number of ones
freq[2] -> number of twos
...
freq[9] -> number of nines
```

This lets me quickly check whether I have enough copies of a digit.

Since there are always exactly 10 possible digits, this data structure uses constant space.

## Operations & Behavior Summary

The algorithm works in three main stages.

### 1. Build the frequency table

I scan the input array once and increase the frequency of every digit.

```text
digits = [0, 2, 2]

freq[0] = 1
freq[2] = 2
```

### 2. Generate valid digit combinations

I use three loops:

```text
First digit   -> 1 to 9
Second digit  -> 0 to 9
Third digit   -> 0, 2, 4, 6, 8
```

This already takes care of the two important number rules:

```text
First digit != 0
Last digit  is even
```

### 3. Check digit availability

For every candidate, I check how many copies of each selected digit are required.

If the same digit appears multiple times in the candidate, I make sure the input contains enough copies.

If everything is available, I increase the answer.

## Complexity

| Complexity       |   Cost | Explanation                                                                      |
| ---------------- | -----: | -------------------------------------------------------------------------------- |
| Time Complexity  | `O(n)` | I scan `n` input digits and check only `9 × 10 × 5 = 450` possible combinations. |
| Space Complexity | `O(1)` | The frequency array always contains only 10 elements.                            |

Here, `n` represents the length of the input `digits` array.

Although the algorithm checks 450 combinations, `450` is a fixed constant, so the overall time complexity is `O(n)`.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int totalNumbers(vector<int>& digits) {
        // I store how many times each digit appears.
        int freq[10] = {};

        // I build the frequency table so I can handle duplicate digits correctly.
        for (int digit : digits) {
            freq[digit]++;
        }

        // I store the total number of distinct valid 3-digit even numbers.
        int answer = 0;

        // I choose the hundreds digit from 1 to 9 because leading zero is not allowed.
        for (int first = 1; first <= 9; first++) {
            // I choose the tens digit from 0 to 9.
            for (int second = 0; second <= 9; second++) {
                // I choose only even digits for the ones position.
                for (int third = 0; third <= 8; third += 2) {
                    // I check whether all three required digits are available.
                    if (freq[first] == 0 || freq[second] == 0 || freq[third] == 0) {
                        // If any required digit is missing, this number cannot be formed.
                        continue;
                    }

                    // If the same digit is used in multiple positions, I need enough copies.
                    if (first == second && second == third && freq[first] < 3) {
                        // Three equal digits require at least three copies.
                        continue;
                    }

                    // If the first two digits are equal, I need at least two copies.
                    if (first == second && freq[first] < 2) {
                        // There are not enough copies of the first digit.
                        continue;
                    }

                    // If the first and third digits are equal, I need at least two copies.
                    if (first == third && freq[first] < 2) {
                        // There are not enough copies of the first digit.
                        continue;
                    }

                    // If the second and third digits are equal, I need at least two copies.
                    if (second == third && freq[second] < 2) {
                        // There are not enough copies of the second digit.
                        continue;
                    }

                    // All three digits are available, so this distinct number is valid.
                    answer++;
                }
            }
        }

        // I return the number of distinct valid 3-digit even numbers.
        return answer;
    }
};
```

### Java

```java
class Solution {
    public int totalNumbers(int[] digits) {
        // I store how many times each digit appears.
        int[] freq = new int[10];

        // I build the frequency table to correctly handle duplicate digits.
        for (int digit : digits) {
            freq[digit]++;
        }

        // I store the number of distinct valid 3-digit even numbers.
        int answer = 0;

        // I choose the hundreds digit from 1 to 9 because zero cannot be the first digit.
        for (int first = 1; first <= 9; first++) {
            // I choose the tens digit from 0 to 9.
            for (int second = 0; second <= 9; second++) {
                // I choose only even digits for the last position.
                for (int third = 0; third <= 8; third += 2) {
                    // I skip the number if any required digit does not exist.
                    if (freq[first] == 0 || freq[second] == 0 || freq[third] == 0) {
                        continue;
                    }

                    // Three equal digits need three copies of that digit.
                    if (first == second && second == third && freq[first] < 3) {
                        continue;
                    }

                    // The first and second positions need two copies when they are equal.
                    if (first == second && freq[first] < 2) {
                        continue;
                    }

                    // The first and third positions need two copies when they are equal.
                    if (first == third && freq[first] < 2) {
                        continue;
                    }

                    // The second and third positions need two copies when they are equal.
                    if (second == third && freq[second] < 2) {
                        continue;
                    }

                    // This combination forms one distinct valid number.
                    answer++;
                }
            }
        }

        // I return the final count.
        return answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} digits
 * @return {number}
 */
var totalNumbers = function(digits) {
    // I store how many times each digit appears.
    const freq = Array(10).fill(0);

    // I build the frequency table to handle duplicate digits correctly.
    for (const digit of digits) {
        freq[digit]++;
    }

    // I store the number of distinct valid 3-digit even numbers.
    let answer = 0;

    // I choose the hundreds digit from 1 to 9 because a leading zero is not allowed.
    for (let first = 1; first <= 9; first++) {
        // I choose the tens digit from 0 to 9.
        for (let second = 0; second <= 9; second++) {
            // I choose only even digits for the ones position.
            for (let third = 0; third <= 8; third += 2) {
                // I skip this combination if any required digit is unavailable.
                if (freq[first] === 0 || freq[second] === 0 || freq[third] === 0) {
                    continue;
                }

                // Three equal digits require three copies of the same digit.
                if (first === second && second === third && freq[first] < 3) {
                    continue;
                }

                // Equal first and second digits require two copies.
                if (first === second && freq[first] < 2) {
                    continue;
                }

                // Equal first and third digits require two copies.
                if (first === third && freq[first] < 2) {
                    continue;
                }

                // Equal second and third digits require two copies.
                if (second === third && freq[second] < 2) {
                    continue;
                }

                // This combination represents one distinct valid number.
                answer++;
            }
        }
    }

    // I return the total number of valid distinct numbers.
    return answer;
};
```

### Python3

```python
class Solution:
    def totalNumbers(self, digits: List[int]) -> int:
        # I store how many times each digit appears.
        freq = [0] * 10

        # I build the frequency table so duplicate digits are handled correctly.
        for digit in digits:
            freq[digit] += 1

        # I store the total number of distinct valid 3-digit even numbers.
        answer = 0

        # I choose the hundreds digit from 1 to 9 because zero cannot be the first digit.
        for first in range(1, 10):
            # I choose the tens digit from 0 to 9.
            for second in range(10):
                # I choose only even digits for the ones position.
                for third in range(0, 10, 2):
                    # I skip this combination if any required digit is unavailable.
                    if freq[first] == 0 or freq[second] == 0 or freq[third] == 0:
                        continue

                    # Three equal digits require three copies of that digit.
                    if first == second == third and freq[first] < 3:
                        continue

                    # Equal first and second digits require two copies.
                    if first == second and freq[first] < 2:
                        continue

                    # Equal first and third digits require two copies.
                    if first == third and freq[first] < 2:
                        continue

                    # Equal second and third digits require two copies.
                    if second == third and freq[second] < 2:
                        continue

                    # This combination forms one distinct valid number.
                    answer += 1

        # I return the final count.
        return answer
```

### Go

```go
func totalNumbers(digits []int) int {
 // I store how many times each digit appears.
 freq := make([]int, 10)

 // I build the frequency table to handle duplicate digits correctly.
 for _, digit := range digits {
  freq[digit]++
 }

 // I store the number of distinct valid 3-digit even numbers.
 answer := 0

 // I choose the hundreds digit from 1 to 9 because zero cannot be the first digit.
 for first := 1; first <= 9; first++ {
  // I choose the tens digit from 0 to 9.
  for second := 0; second <= 9; second++ {
   // I choose only even digits for the ones position.
   for third := 0; third <= 8; third += 2 {
    // I skip this combination if any required digit is unavailable.
    if freq[first] == 0 || freq[second] == 0 || freq[third] == 0 {
     continue
    }

    // Three equal digits require three copies of that digit.
    if first == second && second == third && freq[first] < 3 {
     continue
    }

    // Equal first and second digits require two copies.
    if first == second && freq[first] < 2 {
     continue
    }

    // Equal first and third digits require two copies.
    if first == third && freq[first] < 2 {
     continue
    }

    // Equal second and third digits require two copies.
    if second == third && freq[second] < 2 {
     continue
    }

    // This combination represents one distinct valid number.
    answer++
   }
  }
 }

 // I return the total number of valid numbers.
 return answer
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five languages. Only the syntax for creating arrays, loops, and accessing values changes.

### C++

I first create an array with 10 positions. Each position represents one digit from `0` to `9`.

I then scan the input vector and increase the corresponding frequency.

After that, I use three nested loops.

The outer loop starts from `1` because the hundreds digit cannot be zero.

The middle loop checks every digit from `0` to `9` because there is no restriction on the tens digit.

The inner loop checks only even digits. This means I never need to create an odd number and reject it later.

For every combination, I check whether the selected digits are available.

If all three selected digits are different, one copy of each is enough.

If two positions contain the same digit, I need at least two copies.

If all three positions contain the same digit, I need at least three copies.

Once the frequency checks pass, I increase the answer.

### Java

The Java solution follows the same logic using an integer array of size `10`.

The enhanced `for` loop is used to build the frequency table.

The three nested loops then explore all possible hundreds, tens, and ones digits.

Java arrays give direct access to the frequency of every digit, so checking availability takes constant time.

The important part is to check repeated digits before counting a candidate. This prevents numbers such as `222` from being counted when the input contains only two copies of `2`.

### JavaScript

In JavaScript, I create an array containing ten zeroes.

I use each digit as an index:

```text
freq[digit]
```

For example, if the current digit is `4`, I increase `freq[4]`.

The nested loops work exactly like the other implementations.

The first loop avoids `0`, and the last loop moves through only even digits.

JavaScript's array indexing makes the frequency lookup constant time, so the overall algorithm remains efficient.

### Python3

In Python3, I create a list with ten zeroes.

I then loop through `digits` and update the corresponding position.

The three loops use Python's `range()`:

```text
1 ... 9       -> hundreds digit
0 ... 9       -> tens digit
0,2,4,6,8     -> ones digit
```

For the last position, `range(0, 10, 2)` naturally produces only even digits.

The repeated-digit checks make sure the candidate does not require more copies of a digit than are available.

### Go

In Go, I create an integer slice with length `10`.

I update the frequency using:

```text
freq[digit]++
```

Then I use three `for` loops to generate the possible digit combinations.

Go does not have a separate `while` keyword, but its `for` loop handles this type of iteration directly.

The frequency checks work the same way as in the other implementations.

Since the input contains only single digits, the frequency slice is always small and the solution uses constant extra space.

## Examples

### Example 1

Input:

```text
digits = [1, 2, 3, 4]
```

The possible valid numbers include:

```text
124  132  142  214
234  312  314  324
342  412  432
```

There are `12` distinct 3-digit even numbers.

Output:

```text
12
```

The last digit is always even, and no digit is reused within a number.

### Example 2

Input:

```text
digits = [0, 2, 2]
```

Frequency:

```text
0 -> 1
2 -> 2
```

Possible numbers:

```text
202 -> valid
220 -> valid
222 -> invalid
```

`222` is invalid because only two copies of `2` are available.

Output:

```text
2
```

### Example 3

Input:

```text
digits = [6, 6, 6]
```

Only one distinct number can be formed:

```text
666
```

All three copies of `6` are available, so it is valid.

Output:

```text
1
```

## How to Use / Run Locally

The code blocks above are intentionally empty so the actual solution can be added separately.

For local testing, I can place the solution inside a small program with the required input and output handling.

### C++

Save the completed solution as:

```text
solution.cpp
```

Compile it with:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it with:

```bash
./solution
```

### Java

Save the completed solution as:

```text
Solution.java
```

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

### JavaScript

Save the completed solution as:

```text
solution.js
```

Run it with:

```bash
node solution.js
```

### Python3

Save the completed solution as:

```text
solution.py
```

Run it with:

```bash
python3 solution.py
```

### Go

Save the completed solution as:

```text
solution.go
```

Run it with:

```bash
go run solution.go
```

For LeetCode, I only need to paste the corresponding `Solution` class or function into the problem's editor because LeetCode provides the input and calls the required method automatically.

## Notes & Optimizations

The input size is very small, so brute-force enumeration is a good fit for this problem.

I do not need to generate permutations of the input array. Doing that would create duplicate work when the input contains repeated digits.

For example:

```text
[2, 2, 0]
```

contains two identical `2`s. Generating permutations and then removing duplicates is unnecessary.

The frequency-array approach is cleaner because I work directly with digit values and their available counts.

There are only:

```text
9 × 10 × 5 = 450
```

possible combinations to check.

Some important edge cases are:

* All digits are the same, such as `[6, 6, 6]`.
* There are repeated digits, such as `[0, 2, 2]`.
* There is no even digit, such as `[1, 3, 5]`.
* Zero is available, but it cannot be used as the first digit.
* A candidate needs more copies of a digit than the input contains.

The key idea is to enforce the number rules while generating candidates instead of generating invalid numbers first and filtering them afterward.

## Author

[Md Aarzoo Islam]

[Code with Aarzoo on Instagram](https://www.instagram.com/codewithaarzoo.in/?utm_source=chatgpt.com)
