# 3754. Concatenate Non-Zero Digits and Multiply by Sum I

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

In this LeetCode problem, we are given a non-negative integer `n`.

The goal is to create a new integer `x` by taking all the non-zero digits from `n` and keeping them in their original order. Every zero digit is skipped.

After building `x`, we calculate the sum of its digits. Finally, we return the product of `x` and the digit sum.

For example, if `n = 10203004`, the non-zero digits are `1`, `2`, `3`, and `4`.

So:

* `x = 1234`
* `sum = 1 + 2 + 3 + 4 = 10`
* Answer = `1234 × 10 = 12340`

This problem mainly tests digit manipulation, number construction, place value handling, and basic math operations.

## Constraints

| Constraint                         | Description                                       |
| ---------------------------------- | ------------------------------------------------- |
| `0 <= n <= 10^9`                   | The input is a non-negative integer               |
| Digits must stay in original order | Non-zero digits cannot be reversed                |
| Zero digits are ignored            | Only digits from `1` to `9` are used to build `x` |

## Intuition

The first thing I noticed is that I need two values while processing the number.

I need the new number formed from all non-zero digits, and I also need the sum of those digits.

The important part is keeping the digits in their original order.

A common way to extract digits is using `% 10`, but that reads the number from right to left. I could reverse the result later, but that adds extra work.

Instead, I can process the digits from left to right by first finding the highest place value of the number.

For example, with `10203004`, the highest place value is `10000000`. Dividing by this value gives me the first digit.

Once I have a digit, I can ignore it if it is zero. If it is non-zero, I append it to the new number using:

`x = x * 10 + digit`

At the same time, I add the digit to the sum.

This gives a direct and efficient solution without using strings, arrays, or any extra data structure.

## Approach

I solve the problem in a few simple steps.

1. Start with `x = 0` to store the number made from non-zero digits.
2. Start with `sum = 0` to store the sum of those digits.
3. Find the highest power of `10` needed to reach the first digit of `n`.
4. Read each digit from left to right.
5. If the current digit is zero, skip it.
6. If the current digit is non-zero:

   * Append it to `x`.
   * Add it to `sum`.
7. Move to the next smaller place value.
8. Return `x * sum`.

For `n = 10203004`, the non-zero digits are processed like this:

`1 -> 12 -> 123 -> 1234`

At the same time, the digit sum becomes:

`1 -> 3 -> 6 -> 10`

The final result is:

`1234 * 10 = 12340`

## Data Structures Used

This solution does not use any extra data structure.

Only a few variables are needed:

* `x` stores the integer formed by concatenating all non-zero digits.
* `sum` stores the sum of the non-zero digits.
* `divisor` stores the current decimal place value.
* `digit` stores the digit currently being processed.

Because no array, string, stack, queue, map, or set is used, the extra space stays constant.

## Operations & Behavior Summary

The algorithm works in two main stages.

First, it finds the highest decimal place value of the input number.

For example:

* `7` starts with divisor `1`
* `1000` starts with divisor `1000`
* `10203004` starts with divisor `10000000`

Second, it scans the number from left to right.

For each digit:

1. Extract the current digit using division.
2. Remove that digit from the remaining number.
3. Check whether the digit is zero.
4. Skip zero digits.
5. Append every non-zero digit to `x`.
6. Add every non-zero digit to `sum`.
7. Reduce the divisor by a factor of `10`.

After all digits are processed, multiply `x` by `sum` and return the result.

## Complexity

| Complexity       | Value  | Explanation                                                                                    |
| ---------------- | ------ | ---------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(k)` | `k` is the number of digits in `n`. The algorithm scans the digits a constant number of times. |
| Space Complexity | `O(1)` | Only a fixed number of variables are used. No extra data structure grows with the input size.  |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long sumAndMultiply(int n) {
        // This will store the number made from all non-zero digits.
        long long x = 0;
        
        // This will store the sum of all non-zero digits.
        long long sum = 0;
        
        // Find the highest place value so I can read digits left to right.
        int divisor = 1;
        while (n / divisor >= 10) {
            divisor *= 10;
        }
        
        // Process every digit from left to right.
        while (divisor > 0) {
            // Get the digit at the current place value.
            int digit = n / divisor;
            
            // Remove the current digit from n.
            n %= divisor;
            
            // Only non-zero digits are added to x and sum.
            if (digit != 0) {
                // Shift x left by one decimal place and append the digit.
                x = x * 10 + digit;
                
                // Add the same digit to the digit sum.
                sum += digit;
            }
            
            // Move to the next smaller place value.
            divisor /= 10;
        }
        
        // Multiply the concatenated number by its digit sum.
        return x * sum;
    }
};
```

### Java

```java
class Solution {
    public long sumAndMultiply(int n) {
        // This stores the number formed by all non-zero digits.
        long x = 0;
        
        // This stores the sum of all non-zero digits.
        long sum = 0;
        
        // Find the highest place value to read digits from left to right.
        int divisor = 1;
        while (n / divisor >= 10) {
            divisor *= 10;
        }
        
        // Process every digit from left to right.
        while (divisor > 0) {
            // Extract the digit at the current place value.
            int digit = n / divisor;
            
            // Remove the extracted digit from n.
            n %= divisor;
            
            // Ignore zero because it should not be part of x.
            if (digit != 0) {
                // Append the digit to the end of x.
                x = x * 10 + digit;
                
                // Add the digit to the sum.
                sum += digit;
            }
            
            // Move to the next digit.
            divisor /= 10;
        }
        
        // Return the required product.
        return x * sum;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @return {number}
 */
var sumAndMultiply = function(n) {
    // This stores the number formed by all non-zero digits.
    let x = 0;
    
    // This stores the sum of all non-zero digits.
    let sum = 0;
    
    // Find the highest place value to read digits left to right.
    let divisor = 1;
    while (Math.floor(n / divisor) >= 10) {
        divisor *= 10;
    }
    
    // Process every digit from left to right.
    while (divisor > 0) {
        // Extract the digit at the current place value.
        const digit = Math.floor(n / divisor);
        
        // Remove the current digit from n.
        n %= divisor;
        
        // Ignore zero digits completely.
        if (digit !== 0) {
            // Append the current digit to x.
            x = x * 10 + digit;
            
            // Add the current digit to the sum.
            sum += digit;
        }
        
        // Move to the next smaller place value.
        divisor = Math.floor(divisor / 10);
    }
    
    // Return the concatenated number multiplied by its digit sum.
    return x * sum;
};
```

### Python3

```python
class Solution:
    def sumAndMultiply(self, n: int) -> int:
        # This stores the number formed by all non-zero digits.
        x = 0
        
        # This stores the sum of all non-zero digits.
        digit_sum = 0
        
        # Find the highest place value to read digits left to right.
        divisor = 1
        while n // divisor >= 10:
            divisor *= 10
        
        # Process every digit from left to right.
        while divisor > 0:
            # Extract the digit at the current place value.
            digit = n // divisor
            
            # Remove the current digit from n.
            n %= divisor
            
            # Ignore zero digits because they should not be part of x.
            if digit != 0:
                # Append the current digit to x.
                x = x * 10 + digit
                
                # Add the current digit to the sum.
                digit_sum += digit
            
            # Move to the next smaller place value.
            divisor //= 10
        
        # Return the required product.
        return x * digit_sum
```

### Go

```go
func sumAndMultiply(n int) int64 {
    // This stores the number formed by all non-zero digits.
    var x int64 = 0

    // This stores the sum of all non-zero digits.
    var sum int64 = 0

    // Find the highest place value to read digits left to right.
    divisor := 1
    for n/divisor >= 10 {
        divisor *= 10
    }

    // Process every digit from left to right.
    for divisor > 0 {
        // Extract the digit at the current place value.
        digit := n / divisor

        // Remove the current digit from n.
        n %= divisor

        // Ignore zero digits completely.
        if digit != 0 {
            // Append the current digit to x.
            x = x*10 + int64(digit)

            // Add the current digit to the sum.
            sum += int64(digit)
        }

        // Move to the next smaller place value.
        divisor /= 10
    }

    // Return the concatenated number multiplied by its digit sum.
    return x * sum
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The main logic is the same in all five programming languages.

The solution starts by creating two variables.

The first variable stores the new integer made from non-zero digits. It starts at `0` because no digit has been added yet.

The second variable stores the sum of the non-zero digits. It also starts at `0`.

Next, the algorithm finds the highest decimal place value.

Suppose the input is:

`n = 10203004`

The divisor grows like this:

`1 -> 10 -> 100 -> 1000 -> 10000 -> 100000 -> 1000000 -> 10000000`

Now the algorithm can read the first digit using integer division.

The first calculation gives:

`10203004 / 10000000 = 1`

The digit is `1`, so it is added to both values:

`x = 1`

`sum = 1`

The first digit is then removed from the remaining number.

The next digit is `0`. Since zero should not be included in the new number, nothing is added to `x` or `sum`.

The next non-zero digit is `2`.

To append `2` to the current value of `x`, the algorithm uses:

`x = x * 10 + 2`

Since `x` was `1`, it becomes:

`1 * 10 + 2 = 12`

The digit sum also changes:

`sum = 1 + 2 = 3`

The same process continues for every remaining digit.

When the digit `3` is found:

`x = 12 * 10 + 3 = 123`

`sum = 3 + 3 = 6`

When the digit `4` is found:

`x = 123 * 10 + 4 = 1234`

`sum = 6 + 4 = 10`

After all digits are processed, the result is:

`1234 * 10 = 12340`

In C++, Java, and Go, a wider integer type is used for the result because multiplication can produce a value larger than a standard 32-bit integer.

In JavaScript, numbers are stored using the `Number` type. The values in this problem stay within the safe integer range.

In Python3, integers automatically grow when needed, so no special larger integer type is required.

The input `n = 0` is also handled correctly. No non-zero digit is added, so both `x` and `sum` remain `0`. The final result is `0`.

## Examples

### Example 1

**Input:**

`n = 10203004`

**Expected Output:**

`12340`

**Trace:**

The non-zero digits are:

`1, 2, 3, 4`

They form:

`x = 1234`

Their sum is:

`1 + 2 + 3 + 4 = 10`

Final calculation:

`1234 * 10 = 12340`

### Example 2

**Input:**

`n = 1000`

**Expected Output:**

`1`

**Trace:**

The only non-zero digit is `1`.

So:

`x = 1`

`sum = 1`

Final calculation:

`1 * 1 = 1`

### Example 3

**Input:**

`n = 0`

**Expected Output:**

`0`

**Trace:**

There are no non-zero digits.

So:

`x = 0`

`sum = 0`

Final calculation:

`0 * 0 = 0`

## How to Use / Run Locally

Before running the solution locally, place the required function inside a complete program with sample input and output code.

### C++

Save the file as `main.cpp`.

Compile it with:

```bash
g++ -std=c++17 main.cpp -o main
```

Run it with:

```bash
./main
```

On Windows, run:

```bash
main.exe
```

### Java

Save the file as `Solution.java`.

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

### JavaScript

Save the file as `solution.js`.

Run it using Node.js:

```bash
node solution.js
```

### Python3

Save the file as `solution.py`.

Run it with:

```bash
python3 solution.py
```

On some systems, the command may be:

```bash
python solution.py
```

### Go

Save the file as `main.go`.

Run it directly with:

```bash
go run main.go
```

Or build the executable first:

```bash
go build main.go
```

Then run the generated executable.

## Notes & Optimizations

The main edge case is `n = 0`. The solution handles it naturally because both the concatenated number and digit sum remain zero.

Zeros at the beginning, middle, or end of the number do not cause any problem. Every zero is simply ignored.

A string-based solution is another valid approach. The number can be converted to a string, and each character can be checked one by one. That version may feel easier for beginners, but it uses extra space for the string.

Another possible method is to extract digits from right to left using `% 10`. However, this reverses the processing order, so extra work is needed to restore the original order.

The left-to-right place value approach avoids both problems. It keeps the original digit order and uses constant extra space.

For the given constraints, this solution is already optimal. Every digit must be checked at least once, so the `O(k)` time complexity cannot be improved in a meaningful way.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
