# 3751. Total Waviness of Numbers in Range I

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

The **Total Waviness of Numbers in Range I** problem asks us to calculate the total waviness of every number inside a given inclusive range `[num1, num2]`.

A digit is considered a:

* Peak if it is strictly greater than both neighboring digits.
* Valley if it is strictly smaller than both neighboring digits.

The first and last digits can never be peaks or valleys because they do not have two neighbors.

For every number in the range, we count all peaks and valleys and add them together. The final answer is the sum of waviness values for all numbers in the range.

This is a classic array and string simulation problem that focuses on digit comparison and traversal.

---

## Constraints

| Constraint                  | Value       |
| --------------------------- | ----------- |
| `1 <= num1 <= num2 <= 10^5` | Valid range |

---

## Intuition

When I first looked at the constraints, I noticed that the maximum value is only `100000`.

That means the total number of numbers I may need to process is at most around one hundred thousand, which is completely manageable.

Instead of searching for a complicated mathematical pattern, I can simply examine every number in the range.

For each number:

1. Convert it into digits.
2. Check every middle digit.
3. Count peaks and valleys.
4. Add the result to the final answer.

Since each number contains very few digits, this straightforward approach is already efficient enough.

---

## Approach

I solve the problem using direct simulation.

Step 1:

Loop through every number from `num1` to `num2`.

Step 2:

Convert the current number into a string.

Step 3:

If the number contains fewer than three digits, skip it because it cannot have a peak or valley.

Step 4:

Check every middle digit.

For each position:

* Compare with the left neighbor.
* Compare with the right neighbor.

Step 5:

If the digit is larger than both neighbors, count one peak.

Step 6:

If the digit is smaller than both neighbors, count one valley.

Step 7:

Add the waviness of the current number to the global answer.

Step 8:

Return the final sum.

---

## Data Structures Used

### String

I convert each number into a string so that accessing neighboring digits becomes simple.

Why I use it:

* Easy digit access
* Cleaner code
* No need for repeated division and modulo operations

### Integer Variables

Used for:

* Current number
* Loop indices
* Final answer
* Waviness count

No advanced data structures are needed for this problem.

---

## Operations & Behavior Summary

The algorithm performs the following operations:

1. Visit every number in the range.
2. Convert the number into digit form.
3. Ignore numbers with fewer than three digits.
4. Scan all middle positions.
5. Check peak conditions.
6. Check valley conditions.
7. Increase waviness whenever a valid pattern appears.
8. Accumulate the result.
9. Return the total waviness.

---

## Complexity

| Metric           | Complexity               | Explanation                                             |
| ---------------- | ------------------------ | ------------------------------------------------------- |
| Time Complexity  | O((num2 - num1 + 1) × d) | Every number is checked once and each digit is examined |
| Space Complexity | O(d)                     | String representation of the current number             |

Where:

* `d` = number of digits in the current number
* `d ≤ 6` because `num2 ≤ 10^5`

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int totalWaviness(int num1, int num2) {
        int answer = 0;

        // Check every number in the given range
        for (int num = num1; num <= num2; num++) {
            string s = to_string(num);

            // Numbers with fewer than 3 digits cannot have peaks or valleys
            if (s.size() < 3) {
                continue;
            }

            // Check every middle digit
            for (int i = 1; i < (int)s.size() - 1; i++) {
                // Peak condition
                if (s[i] > s[i - 1] && s[i] > s[i + 1]) {
                    answer++;
                }
                // Valley condition
                else if (s[i] < s[i - 1] && s[i] < s[i + 1]) {
                    answer++;
                }
            }
        }

        return answer;
    }
};
```

### Java

```java
class Solution {
    public int totalWaviness(int num1, int num2) {
        int answer = 0;

        // Check every number in the range
        for (int num = num1; num <= num2; num++) {
            String s = String.valueOf(num);

            // Numbers with fewer than 3 digits have waviness 0
            if (s.length() < 3) {
                continue;
            }

            // Check every middle digit
            for (int i = 1; i < s.length() - 1; i++) {
                // Peak condition
                if (s.charAt(i) > s.charAt(i - 1) &&
                    s.charAt(i) > s.charAt(i + 1)) {
                    answer++;
                }
                // Valley condition
                else if (s.charAt(i) < s.charAt(i - 1) &&
                         s.charAt(i) < s.charAt(i + 1)) {
                    answer++;
                }
            }
        }

        return answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} num1
 * @param {number} num2
 * @return {number}
 */
var totalWaviness = function(num1, num2) {
    let answer = 0;

    // Check every number in the range
    for (let num = num1; num <= num2; num++) {
        const s = String(num);

        // Numbers with fewer than 3 digits have no peaks or valleys
        if (s.length < 3) {
            continue;
        }

        // Check every middle digit
        for (let i = 1; i < s.length - 1; i++) {
            // Peak condition
            if (s[i] > s[i - 1] && s[i] > s[i + 1]) {
                answer++;
            }
            // Valley condition
            else if (s[i] < s[i - 1] && s[i] < s[i + 1]) {
                answer++;
            }
        }
    }

    return answer;
};
```

### Python3

```python
class Solution:
    def totalWaviness(self, num1: int, num2: int) -> int:
        answer = 0

        # Check every number in the range
        for num in range(num1, num2 + 1):
            s = str(num)

            # Numbers with fewer than 3 digits have waviness 0
            if len(s) < 3:
                continue

            # Check every middle digit
            for i in range(1, len(s) - 1):
                # Peak condition
                if s[i] > s[i - 1] and s[i] > s[i + 1]:
                    answer += 1

                # Valley condition
                elif s[i] < s[i - 1] and s[i] < s[i + 1]:
                    answer += 1

        return answer
```

### Go

```go
func totalWaviness(num1 int, num2 int) int {
    answer := 0

    // Check every number in the range
    for num := num1; num <= num2; num++ {
        s := strconv.Itoa(num)

        // Numbers with fewer than 3 digits have waviness 0
        if len(s) < 3 {
            continue
        }

        // Check every middle digit
        for i := 1; i < len(s)-1; i++ {
            // Peak condition
            if s[i] > s[i-1] && s[i] > s[i+1] {
                answer++
            } else if s[i] < s[i-1] && s[i] < s[i+1] {
                // Valley condition
                answer++
            }
        }
    }

    return answer
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same in all five languages.

### Step 1: Iterate Through the Range

I process every number between `num1` and `num2`.

Since the range size is small enough, checking each number individually is perfectly acceptable.

### Step 2: Convert Number into Digits

Rather than extracting digits manually using division and modulo operations, I convert the number into a string.

This makes neighboring comparisons much easier.

### Step 3: Skip Short Numbers

A number with fewer than three digits cannot contain a middle digit.

Without a middle digit, there is no possible peak or valley.

So these numbers contribute zero waviness.

### Step 4: Check Every Middle Digit

I only examine indices from:

```text
1 to length - 2
```

The first and last digits are never valid candidates.

### Step 5: Detect Peaks

A peak occurs when:

```text
current > left
current > right
```

Both conditions must be true.

### Step 6: Detect Valleys

A valley occurs when:

```text
current < left
current < right
```

Again, both conditions must be true.

### Step 7: Count Waviness

Whenever a peak or valley is found, I increase the count.

Multiple peaks and valleys can exist in the same number.

### Step 8: Add to Final Answer

The waviness of each number gets added to the overall result.

After all numbers are processed, the answer is returned.

---

## Examples

### Example 1

Input

```text
num1 = 120
num2 = 130
```

Output

```text
3
```

Trace

```text
120 -> peak at 2 -> 1
121 -> peak at 2 -> 1
130 -> peak at 3 -> 1
```

Total

```text
1 + 1 + 1 = 3
```

---

### Example 2

Input

```text
num1 = 198
num2 = 202
```

Output

```text
3
```

Trace

```text
198 -> peak at 9 -> 1
201 -> valley at 0 -> 1
202 -> valley at 0 -> 1
```

Total

```text
3
```

---

### Example 3

Input

```text
num1 = 4848
num2 = 4848
```

Output

```text
2
```

Trace

```text
4 8 4 8
  ^
Peak

4 8 4 8
    ^
Valley
```

Total waviness

```text
2
```

---

## How to Use / Run Locally

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

Build

```bash
go build solution.go
```

---

## Notes & Optimizations

* The straightforward simulation approach is already fast enough for the given constraints.
* Every number contains at most six digits, so digit processing is extremely cheap.
* No extra arrays, maps, stacks, queues, or trees are needed.
* Converting numbers into strings keeps the implementation simple and easy to understand.
* A digit equal to one of its neighbors is neither a peak nor a valley because the comparison must be strictly greater or strictly smaller.
* This solution is optimal for the constraints provided in the problem.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
