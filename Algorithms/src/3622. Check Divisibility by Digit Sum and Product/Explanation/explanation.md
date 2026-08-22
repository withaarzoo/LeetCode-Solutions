# 3622. Check Divisibility by Digit Sum and Product - LeetCode Solution

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

In LeetCode 3622, **Check Divisibility by Digit Sum and Product**, I am given a positive integer `n`.

I need to calculate two values from its digits:

1. The sum of all digits.
2. The product of all digits.

Then I add these two values together.

The final task is to check whether the original number `n` is divisible by that result. If it is divisible, the answer is `true`. Otherwise, the answer is `false`.

For example, if `n = 99`:

* Digit sum = `9 + 9 = 18`
* Digit product = `9 × 9 = 81`
* Total = `18 + 81 = 99`

Since `99` is divisible by `99`, the answer is `true`.

This is a simple digit manipulation problem where I can process every digit directly using basic arithmetic.

## Constraints

| Constraint           | Value  |
| -------------------- | ------ |
| Minimum value of `n` | `1`    |
| Maximum value of `n` | `10^6` |

## Intuition

The first thing I noticed was that I only need information about the digits of the number.

I do not need to sort anything, search through an array, or use any complicated algorithm. I just need to visit every digit once and calculate its contribution to both the digit sum and digit product.

For each digit, I can do two things at the same time:

* Add it to the running sum.
* Multiply it into the running product.

After processing all digits, I add the final sum and product. Then I check whether the original number is divisible by that value.

This makes a direct digit-by-digit solution the most natural approach.

## Approach

I use a temporary value to process the digits while keeping the original number safe for the final divisibility check.

The solution works like this:

1. Save the original value of `n`.
2. Initialize `digitSum` as `0`.
3. Initialize `digitProduct` as `1`.
4. Extract the last digit using modulo `10`.
5. Add that digit to `digitSum`.
6. Multiply that digit into `digitProduct`.
7. Remove the last digit using integer division by `10`.
8. Repeat until no digits are left.
9. Add `digitSum` and `digitProduct`.
10. Check whether the original number is divisible by this final value.
11. Return `true` if the remainder is `0`; otherwise, return `false`.

This approach processes every digit exactly once.

## Data Structures Used

No additional data structure is needed for this solution.

I only use a few integer variables:

* `original` stores the original value of the number.
* `digitSum` stores the sum of all digits.
* `digitProduct` stores the product of all digits.
* `digit` stores the current digit being processed.

Since there is no need to store all digits separately, an array, list, stack, or string would only add unnecessary work.

## Operations & Behavior Summary

The algorithm follows this simple flow:

1. Keep a copy of the original number.
2. Start the digit sum at `0`.
3. Start the digit product at `1`.
4. Take the last digit of the current number.
5. Update both the sum and product using that digit.
6. Remove the processed digit.
7. Continue until the number becomes `0`.
8. Add the digit sum and digit product.
9. Divide the original number by this value.
10. If the remainder is `0`, return `true`; otherwise, return `false`.

This gives a clean and efficient solution for the Check Divisibility by Digit Sum and Product problem.

## Complexity

| Complexity       | Value      | Explanation                                                                                    |
| ---------------- | ---------- | ---------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(log n)` | I process each digit of `n` once. A number with value `n` has approximately `log₁₀(n)` digits. |
| Space Complexity | `O(1)`     | I only use a fixed number of variables and no extra data structures.                           |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool checkDivisibility(int n) {
        int original = n; // I save the original number because n will be changed while extracting digits.
        int digitSum = 0; // I start the sum at 0 because digits will be added to it.
        int digitProduct = 1; // I start the product at 1 because it is the neutral value for multiplication.

        while (n > 0) { // I keep processing until every digit has been removed.
            int digit = n % 10; // I extract the last digit of the current number.
            digitSum += digit; // I add the current digit to the total digit sum.
            digitProduct *= digit; // I multiply the current digit into the total digit product.
            n /= 10; // I remove the last digit using integer division.
        }

        int divisor = digitSum + digitProduct; // I calculate the required sum of digit sum and digit product.
        return original % divisor == 0; // I return true only when the original number is divisible by this value.
    }
};
```

### Java

```java
class Solution {
    public boolean checkDivisibility(int n) {
        int original = n; // I save the original number because n will change while extracting digits.
        int digitSum = 0; // I start the digit sum at 0 for repeated addition.
        int digitProduct = 1; // I start the digit product at 1 for repeated multiplication.

        while (n > 0) { // I continue until all digits have been processed.
            int digit = n % 10; // I get the last digit of the current number.
            digitSum += digit; // I add the digit to the total sum.
            digitProduct *= digit; // I multiply the digit into the total product.
            n /= 10; // I remove the last digit from the number.
        }

        int divisor = digitSum + digitProduct; // I calculate the value that must divide the original number.
        return original % divisor == 0; // I return whether the remainder is 0.
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @return {boolean}
 */
var checkDivisibility = function(n) {
    const original = n; // I save the original number because n will be reduced digit by digit.
    let digitSum = 0; // I initialize the digit sum for repeated addition.
    let digitProduct = 1; // I initialize the digit product for repeated multiplication.

    while (n > 0) { // I continue until there are no digits left to process.
        const digit = n % 10; // I extract the last digit.
        digitSum += digit; // I add the current digit to the digit sum.
        digitProduct *= digit; // I multiply the current digit into the digit product.
        n = Math.floor(n / 10); // I remove the last digit and keep only the remaining whole-number part.
    }

    const divisor = digitSum + digitProduct; // I calculate the required divisor.
    return original % divisor === 0; // I return true only if the original number divides exactly.
};
```

### Python3

```python
class Solution:
    def checkDivisibility(self, n: int) -> bool:
        original = n  # I save the original value because n will change while extracting digits.
        digit_sum = 0  # I start the digit sum at 0 for addition.
        digit_product = 1  # I start the digit product at 1 for multiplication.

        while n > 0:  # I continue until every digit has been processed.
            digit = n % 10  # I extract the last digit of the current number.
            digit_sum += digit  # I add the digit to the total digit sum.
            digit_product *= digit  # I multiply the digit into the total digit product.
            n //= 10  # I remove the last digit using integer division.

        divisor = digit_sum + digit_product  # I calculate the value that must divide the original number.
        return original % divisor == 0  # I return whether the original number is exactly divisible.
```

### Go

```go
func checkDivisibility(n int) bool {
 original := n // I save the original number because n will be changed while extracting digits.
 digitSum := 0 // I start the digit sum at 0 for repeated addition.
 digitProduct := 1 // I start the digit product at 1 for repeated multiplication.

 for n > 0 { // I keep processing digits until the number becomes 0.
  digit := n % 10 // I extract the last digit of the current number.
  digitSum += digit // I add the current digit to the digit sum.
  digitProduct *= digit // I multiply the current digit into the digit product.
  n /= 10 // I remove the last digit using integer division.
 }

 divisor := digitSum + digitProduct // I calculate the required divisor from the sum and product.
 return original%divisor == 0 // I return true only when the original number has no remainder.
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five programming languages. Only the syntax for declaring variables, defining functions, and performing integer division changes slightly.

I first save the original value of `n`.

This is important because I need to repeatedly divide the number by `10` while extracting its digits. After all digits are processed, the working value becomes `0`. Without saving the original value, I would not have the number needed for the final divisibility check.

Next, I initialize the digit sum with `0`.

This is the correct starting value because each digit will be added to it. Starting with any other value would change the final sum.

I initialize the digit product with `1`.

I do this because `1` is the neutral value for multiplication. If I started the product with `0`, every multiplication would remain `0`, which would give the wrong result.

Then I repeatedly extract the last digit.

Modulo `10` gives me the last digit of a positive integer. For example:

* `99 % 10 = 9`
* `23 % 10 = 3`
* `456 % 10 = 6`

Once I have the digit, I update both calculations.

I add it to the digit sum and multiply it into the digit product. Processing both values in the same loop avoids making two separate passes through the number.

After that, I remove the last digit using integer division by `10`.

For example:

* `99 / 10` becomes `9`
* `9 / 10` becomes `0`

The loop continues until every digit has been processed.

Finally, I add the digit sum and digit product together. This gives the divisor required by the problem.

I then check whether the original number leaves a remainder of `0` when divided by this value.

If the remainder is `0`, the number is divisible, so I return `true`.

Otherwise, I return `false`.

### Language-specific behavior

In C++, Java, and Go, integer division with integer variables automatically removes the decimal part.

In Python, floor division is used to make the integer division behavior explicit.

In JavaScript, numbers are stored as floating-point values, so the result is converted back to an integer when removing digits.

The core algorithm remains exactly the same in every language.

## Examples

### Example 1

**Input:** `n = 99`

**Expected Output:** `true`

**Trace:**

* Digits are `9` and `9`
* Digit sum = `9 + 9 = 18`
* Digit product = `9 × 9 = 81`
* Final divisor = `18 + 81 = 99`
* `99 % 99 = 0`

The remainder is `0`, so the answer is `true`.

### Example 2

**Input:** `n = 23`

**Expected Output:** `false`

**Trace:**

* Digits are `2` and `3`
* Digit sum = `2 + 3 = 5`
* Digit product = `2 × 3 = 6`
* Final divisor = `5 + 6 = 11`
* `23 % 11 = 1`

The remainder is not `0`, so the answer is `false`.

### Example 3

**Input:** `n = 22`

**Expected Output:** `false`

**Trace:**

* Digit sum = `2 + 2 = 4`
* Digit product = `2 × 2 = 4`
* Final divisor = `4 + 4 = 8`
* `22 % 8 = 6`

The remainder is not `0`, so the answer is `false`.

## How to Use / Run Locally

Before running any version, paste the solution code into a file using the correct file extension.

### C++

Save the solution in a file such as:

```text
solution.cpp
```

Compile it:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it:

```bash
./solution
```

On Windows:

```bash
solution.exe
```

### Java

Save the solution in:

```text
Solution.java
```

Compile it:

```bash
javac Solution.java
```

Run it:

```bash
java Solution
```

### JavaScript

Save the solution in:

```text
solution.js
```

Make sure Node.js is installed, then run:

```bash
node solution.js
```

### Python3

Save the solution in:

```text
solution.py
```

Run it with Python 3:

```bash
python3 solution.py
```

On some systems, the command may be:

```bash
python solution.py
```

### Go

Save the solution in:

```text
solution.go
```

Run it directly:

```bash
go run solution.go
```

Or build an executable first:

```bash
go build solution.go
```

Then run the generated executable.

For the LeetCode platform, I only need to paste the required solution method or function into the editor. LeetCode handles the input and test cases automatically.

## Notes & Optimizations

The main edge case to think about is a number containing the digit `0`.

If any digit is `0`, the complete digit product becomes `0`. The final divisor is then simply the digit sum.

For example, with `n = 101`:

* Digit sum = `1 + 0 + 1 = 2`
* Digit product = `1 × 0 × 1 = 0`
* Final divisor = `2`

The same arithmetic-based approach still works correctly.

I do not convert the number into a string because direct digit extraction using modulo and division is enough. A string-based solution would also work, but it would use extra space and is unnecessary here.

There is also no need for a more advanced optimization because the maximum input size is small. Still, the arithmetic solution is already optimal for this problem because every digit must be examined at least once.

This makes the solution simple, efficient, and easy to implement in C++, Java, JavaScript, Python, and Go.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
