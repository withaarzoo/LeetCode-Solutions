# 1291. Sequential Digits - LeetCode Solution

A beginner-friendly and optimized solution for **LeetCode 1291 - Sequential Digits**. This repository explains the intuition, approach, complexity analysis, and provides multi-language implementations in C++, Java, JavaScript, Python, and Go. If you're preparing for coding interviews or practicing Data Structures and Algorithms (DSA), this problem is a great example of working with number generation, string manipulation, and simple enumeration techniques.

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

The goal is to find every **sequential digit number** within a given range.

A sequential digit number is a number where every digit is exactly one greater than the previous digit.

For example:

- `123`
- `4567`
- `6789`

These are valid because each digit increases by one.

Numbers like `124`, `135`, or `332` are not sequential.

The input contains two integers:

- `low` — the beginning of the range
- `high` — the end of the range

The output should be a sorted list containing every sequential digit number that falls within the given range.

This problem is commonly asked in coding interviews because it tests observation skills more than complicated algorithms.

---

## Constraints

| Constraint | Value |
|------------|-------|
| Minimum value of `low` | 10 |
| Maximum value of `high` | 10<sup>9</sup> |
| Condition | `10 ≤ low ≤ high ≤ 10⁹` |

---

## Intuition

When I first looked at the problem, I noticed something interesting.

Every valid sequential digit number already exists inside the fixed sequence:

```text
123456789
```

For example:

- `123`
- `234`
- `3456`
- `6789`

They are simply continuous parts of the same string.

That means I don't need recursion, backtracking, or checking every number in the range.

Instead, I can generate every possible substring of `"123456789"` with valid lengths and keep only those numbers that lie between `low` and `high`.

Since there are only nine digits, the total number of possible sequential numbers is very small.

---

## Approach

I solved the problem using a simple generation technique.

First, I stored the string:

```text
123456789
```

Then I calculated the number of digits in both `low` and `high`.

After that, I generated every possible substring whose length falls within that range.

For each substring:

1. Convert it into an integer.
2. Check whether it lies between `low` and `high`.
3. If it does, add it to the answer.

Since I generate numbers from left to right and from smaller lengths to larger lengths, the final result is already sorted.

No extra sorting step is required.

---

## Data Structures Used

| Data Structure | Purpose |
|---------------|---------|
| String | Stores `"123456789"` and helps generate sequential numbers easily. |
| List / Vector / Array | Stores all valid sequential digit numbers that fall within the given range. |

No advanced data structures such as stacks, queues, trees, or hash maps are needed for this problem.

---

## Operations & Behavior Summary

The algorithm works in the following order:

1. Store the fixed digit sequence `"123456789"`.
2. Count the digits in `low`.
3. Count the digits in `high`.
4. Generate every substring whose length is between those two digit counts.
5. Convert each substring into an integer.
6. Check whether the generated number lies inside the required range.
7. Add valid numbers to the answer list.
8. Return the answer.

Because there are only a handful of possible sequential digit numbers, this approach finishes almost instantly.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | **O(1)** | There are only a fixed number of possible sequential digit numbers, so the work never grows with the input size. |
| Space Complexity | **O(1)** | Only the output list is used. No extra memory proportional to the input is required. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> sequentialDigits(int low, int high) {
        // String containing all consecutive digits
        string digits = "123456789";

        // Result array
        vector<int> ans;

        // Number of digits in low and high
        int minLen = to_string(low).size();
        int maxLen = to_string(high).size();

        // Try every possible length
        for (int len = minLen; len <= maxLen; len++) {

            // Generate every substring of current length
            for (int start = 0; start + len <= 9; start++) {

                // Convert substring into an integer
                int num = stoi(digits.substr(start, len));

                // Keep only numbers inside the range
                if (num >= low && num <= high)
                    ans.push_back(num);
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public List<Integer> sequentialDigits(int low, int high) {

        // String containing all consecutive digits
        String digits = "123456789";

        // Result list
        List<Integer> ans = new ArrayList<>();

        // Number of digits in low and high
        int minLen = String.valueOf(low).length();
        int maxLen = String.valueOf(high).length();

        // Try every possible length
        for (int len = minLen; len <= maxLen; len++) {

            // Generate every substring of current length
            for (int start = 0; start + len <= 9; start++) {

                // Convert substring into an integer
                int num = Integer.parseInt(digits.substring(start, start + len));

                // Keep only numbers inside the range
                if (num >= low && num <= high) {
                    ans.add(num);
                }
            }
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} low
 * @param {number} high
 * @return {number[]}
 */
var sequentialDigits = function(low, high) {

    // String containing all consecutive digits
    const digits = "123456789";

    // Result array
    const ans = [];

    // Number of digits in low and high
    const minLen = low.toString().length;
    const maxLen = high.toString().length;

    // Try every possible length
    for (let len = minLen; len <= maxLen; len++) {

        // Generate every substring of current length
        for (let start = 0; start + len <= 9; start++) {

            // Convert substring into a number
            const num = Number(digits.substring(start, start + len));

            // Keep only numbers inside the range
            if (num >= low && num <= high) {
                ans.push(num);
            }
        }
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def sequentialDigits(self, low: int, high: int) -> List[int]:

        # String containing all consecutive digits
        digits = "123456789"

        # Result list
        ans = []

        # Number of digits in low and high
        min_len = len(str(low))
        max_len = len(str(high))

        # Try every possible length
        for length in range(min_len, max_len + 1):

            # Generate every substring of current length
            for start in range(10 - length):

                # Convert substring into an integer
                num = int(digits[start:start + length])

                # Keep only numbers inside the range
                if low <= num <= high:
                    ans.append(num)

        return ans
```

### Go

```go
func sequentialDigits(low int, high int) []int {

 // String containing all consecutive digits
 digits := "123456789"

 // Result slice
 ans := []int{}

 // Number of digits in low and high
 minLen := len([]byte(fmt.Sprintf("%d", low)))
 maxLen := len([]byte(fmt.Sprintf("%d", high)))

 // Try every possible length
 for length := minLen; length <= maxLen; length++ {

  // Generate every substring of current length
  for start := 0; start+length <= 9; start++ {

   // Convert substring into an integer
   num, _ := strconv.Atoi(digits[start : start+length])

   // Keep only numbers inside the range
   if num >= low && num <= high {
    ans = append(ans, num)
   }
  }
 }

 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The overall logic is exactly the same in every programming language.

Only the syntax changes.

The first step is creating the fixed string:

```text
123456789
```

This string already contains every possible sequential digit number.

Next, the algorithm determines how many digits the answer can have.

For example, if:

```text
low = 100
high = 13000
```

then only numbers with 3, 4, or 5 digits need to be generated.

There is no reason to create numbers with two digits or six digits because they can never be part of the answer.

For every valid length, the algorithm slides across the string.

For example, when the current length is four, the generated numbers become:

```text
1234
2345
3456
4567
5678
6789
```

Each substring is converted into an integer.

After conversion, the algorithm checks whether the number lies inside the required range.

If it does, the number is stored.

Otherwise, it is ignored.

Because the substrings are generated from left to right, the numbers naturally appear in increasing order.

This removes the need for sorting later.

The same idea is implemented in C++, Java, JavaScript, Python3, and Go.

Only functions like substring extraction and string-to-integer conversion have different names in each language.

---

## Examples

### Example 1

**Input**

```text
low = 100
high = 300
```

**Output**

```text
[123, 234]
```

**Trace**

- Generate every 3-digit sequential number.
- `123` is valid.
- `234` is valid.
- `345` is greater than `300`.
- Final answer becomes:

```text
[123, 234]
```

---

### Example 2

**Input**

```text
low = 1000
high = 13000
```

**Output**

```text
[1234,2345,3456,4567,5678,6789,12345]
```

**Trace**

Generate every 4-digit sequential number.

```text
1234
2345
3456
4567
5678
6789
```

Generate every 5-digit sequential number.

```text
12345
23456
34567
...
```

Only `12345` is inside the range.

---

### Example 3

**Input**

```text
low = 50
high = 150
```

**Output**

```text
[56,67,78,89,123]
```

**Trace**

The algorithm generates every 2-digit and 3-digit sequential number.

Only numbers inside the given range are stored.

---

## How to Use / Run Locally

Clone the repository:

```bash
git clone https://github.com/your-username/sequential-digits.git
```

Move into the project folder:

```bash
cd sequential-digits
```

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

---

## Notes & Optimizations

- The search space is extremely small because only sequential numbers made from the digits `1` through `9` are possible.
- The algorithm never checks every integer between `low` and `high`, making it much faster than brute force.
- The generated numbers are already sorted, so no sorting step is required.
- This solution avoids recursion, backtracking, and complex data structures.
- Another possible approach is using Breadth-First Search (BFS) to build sequential numbers digit by digit, but using the fixed string `"123456789"` is shorter, easier to understand, and equally efficient.
- Since the number of possible sequential digit numbers is fixed, the running time remains constant regardless of the input values.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
