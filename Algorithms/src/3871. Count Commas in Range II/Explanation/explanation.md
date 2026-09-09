# Count Commas in Range II — LeetCode 3871

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

In LeetCode 3871, **Count Commas in Range II**, I am given an integer `n`.

I need to count the total number of commas that appear when writing every integer from `1` to `n` using standard number formatting.

A comma is added after every three digits from the right.

For example:

```text
1,000     -> 1 comma
10,000    -> 1 comma
999,999   -> 1 comma
1,000,000 -> 2 commas
```

The goal is to return the total number of commas used for all numbers from `1` through `n`.

For example, when `n = 1002`:

```text
1,000
1,001
1,002
```

Each of these numbers contains one comma, so the answer is `3`.

This problem looks like a counting problem, but the large constraint means I cannot simply visit every number one by one.

## Constraints

| Constraint | Value             |
| ---------- | ----------------- |
| `n`        | `1 <= n <= 10^15` |

Because `n` can be extremely large, an `O(n)` solution is not practical.

## Intuition

My first thought was to loop from `1` to `n`, format every number, and count its commas.

That works logically, but it becomes far too slow when `n` is close to `10^15`.

The important observation is that numbers with the same number of digits use the same number of commas.

For example:

```text
1 - 999
    0 commas

1,000 - 999,999
    1 comma

1,000,000 - 999,999,999
    2 commas

1,000,000,000 - 999,999,999,999
    3 commas
```

So instead of checking every number, I can count an entire range at once.

That reduces the problem to only a few groups.

## Approach

I start with the first number that needs a comma:

```text
start = 1000
commas = 1
answer = 0
```

For each group, I find its ending value.

The first group is:

```text
1,000 -------------------- 999,999
          1 comma
```

The next group is:

```text
1,000,000 ---------------- 999,999,999
             2 commas
```

The next starting point is always `1000` times the previous starting point.

For every group, I calculate:

```text
number of values = end - start + 1
group contribution = number of values * commas
```

Then I add that contribution to the answer.

For `n = 1002`:

```text
start = 1000
end   = 1002

count = 1002 - 1000 + 1
      = 3

contribution = 3 * 1
             = 3
```

After processing that group, the next start becomes `1,000,000`, which is already greater than `1002`, so the algorithm stops.

The important part is that I never process numbers individually. I process complete comma groups.

## Data Structures Used

No special data structure is needed.

I only use a few integer variables:

| Variable | Purpose                                       |
| -------- | --------------------------------------------- |
| `start`  | Beginning of the current comma group          |
| `end`    | End of the current group                      |
| `commas` | Number of commas in every number in the group |
| `count`  | Number of integers in the current group       |
| `answer` | Total commas counted so far                   |

This keeps the solution simple and uses constant extra space.

## Operations & Behavior Summary

The algorithm works like this:

1. Start at `1000`, because numbers below it contain no commas.
2. Set the comma count to `1`.
3. Find the end of the current group.
4. Count how many numbers are inside the group.
5. Multiply that count by the number of commas per number.
6. Add the result to the total answer.
7. Multiply `start` by `1000` to move to the next comma group.
8. Increase the comma count by `1`.
9. Repeat until `start > n`.
10. Return the final answer.

In plain pseudocode:

```text
answer = 0
start = 1000
commas = 1

while start <= n:
    end = min(n, start * 1000 - 1)
    count = end - start + 1
    answer += count * commas

    start *= 1000
    commas += 1

return answer
```

## Complexity

| Complexity       |         Cost | Explanation                                                                             |
| ---------------- | -----------: | --------------------------------------------------------------------------------------- |
| Time Complexity  | `O(log₁₀ n)` | `start` is multiplied by `1000` each iteration, so only the digit groups are processed. |
| Space Complexity |       `O(1)` | I only use a few integer variables and no extra data structure.                         |

Since `n <= 10^15`, the number of iterations is very small.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long countCommas(long long n) {
        // Start at 1000 because numbers below 1000 contain no commas.
        long long start = 1000;

        // Every number from 1000 to 999999 contains exactly one comma.
        long long commas = 1;

        // Store the total number of commas found across all numbers.
        long long answer = 0;

        // Process one comma group at a time.
        while (start <= n) {
            // The group normally ends just before start * 1000.
            // If that exceeds n, the group ends at n instead.
            long long end = (start > n / 1000)
                ? n
                : start * 1000 - 1;

            // Count how many numbers are present in this group.
            long long count = end - start + 1;

            // Every number in this group has the same number of commas,
            // so I can add their total contribution at once.
            answer += count * commas;

            // Move to the next comma group.
            start *= 1000;

            // The next group has one additional comma.
            ++commas;
        }

        // Return the total number of commas used.
        return answer;
    }
};
```

### Java

```java
class Solution {
    public long countCommas(long n) {
        // Start at 1000 because numbers smaller than 1000 have no commas.
        long start = 1000;

        // Numbers in the first comma group contain one comma.
        long commas = 1;

        // This variable stores the total number of commas.
        long answer = 0;

        // Process each comma group until the start goes beyond n.
        while (start <= n) {
            // If start * 1000 would go beyond n, use n as the group end.
            // The division check also avoids unnecessary large multiplication.
            long end = (start > n / 1000)
                ? n
                : start * 1000 - 1;

            // Calculate how many numbers belong to this group.
            long count = end - start + 1;

            // Add the contribution of this whole group.
            answer += count * commas;

            // Move to the next group, where the comma count increases by one.
            start *= 1000;
            ++commas;
        }

        // Return the total comma count.
        return answer;
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
    // Start at 1000 because every smaller number has zero commas.
    let start = 1000;

    // Numbers in this first group have one comma.
    let commas = 1;

    // Store the total number of commas.
    let answer = 0;

    // Process each group while its starting value is inside the range.
    while (start <= n) {
        // Check before multiplying so I do not needlessly create a value
        // larger than n and keep all useful calculations within the safe range.
        let end;

        if (start > n / 1000) {
            // The current group reaches only up to n.
            end = n;
        } else {
            // Otherwise, the group ends immediately before start * 1000.
            end = start * 1000 - 1;
        }

        // Count how many integers are inside the current group.
        const count = end - start + 1;

        // Every number in this group uses the same number of commas.
        answer += count * commas;

        // Move to the next group of numbers.
        start *= 1000;

        // The next group has one additional comma per number.
        commas++;
    }

    // Return the total number of commas.
    return answer;
};
```

### Python3

```python
class Solution:
    def countCommas(self, n: int) -> int:
        # Start at 1000 because numbers below 1000 contain no commas.
        start = 1000

        # Every number from 1000 onward in the first group has one comma.
        commas = 1

        # Store the total number of commas across the whole range.
        answer = 0

        # Process one comma group at a time.
        while start <= n:
            # If the next group boundary is beyond n, stop the group at n.
            if start > n // 1000:
                end = n
            else:
                # Otherwise, the group ends just before start * 1000.
                end = start * 1000 - 1

            # Count the numbers inside this group.
            count = end - start + 1

            # Add the number of commas contributed by this entire group.
            answer += count * commas

            # Move to the next group, which starts at a power of 1000.
            start *= 1000

            # The next group needs one more comma per number.
            commas += 1

        # Return the final total.
        return answer
```

### Go

```go
func countCommas(n int64) int64 {
 // Start at 1000 because numbers smaller than 1000 have no commas.
 start := int64(1000)

 // The first comma group contains one comma per number.
 commas := int64(1)

 // Store the total number of commas.
 var answer int64

 // Process every comma group that starts within the range.
 for start <= n {
  var end int64

  // If the next group boundary is beyond n, the current group ends at n.
  if start > n/1000 {
   end = n
  } else {
   // Otherwise, the current group ends immediately before start * 1000.
   end = start*1000 - 1
  }

  // Count how many numbers are in the current group.
  count := end - start + 1

  // Every number in the group has the same comma count.
  answer += count * commas

  // Move to the next group of numbers.
  start *= 1000

  // The next group has one additional comma.
  commas++
 }

 // Return the total number of commas.
 return answer
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five languages, so the main difference is how each language handles integer types and syntax.

### C++

I use `long long` because the input can reach `10^15`, which is much larger than a normal 32-bit integer.

I initialize `start` with `1000` because every number below `1000` contributes zero commas.

Then I keep a `commas` variable starting at `1`.

Inside the loop, I calculate the end of the current group. Normally, that is:

```text
start * 1000 - 1
```

But when the input ends in the middle of a group, I only need to process up to `n`.

The implementation also checks `n / 1000` before multiplying. This avoids unnecessary overflow-style arithmetic and keeps the boundary calculation safe.

After finding the range, I calculate:

```text
count = end - start + 1
```

and add:

```text
count * commas
```

to the answer.

Finally, I multiply `start` by `1000` and increment `commas` to move to the next group.

### Java

The Java solution uses `long` for the same reason as C++: the input can be as large as `10^15`.

The loop follows the exact same grouped counting idea.

The important part is that I do not convert numbers into strings. Converting and formatting every number would create an unnecessary `O(n)` approach.

Instead, one iteration represents an entire group of numbers that all have the same comma count.

The condition used to calculate `end` also avoids multiplying a value when the next group would already be outside the required range.

### JavaScript

The JavaScript version uses the normal `number` type.

The largest input, `10^15`, is still within JavaScript's safe integer range, so the required arithmetic can be performed with `number`.

I still avoid blindly multiplying `start` by `1000` while calculating the group boundary. I first check whether the next boundary is relevant.

The rest of the logic stays unchanged:

```text
find group
    ↓
count numbers
    ↓
multiply by commas
    ↓
add to answer
    ↓
move to next group
```

This makes the JavaScript solution just as efficient in terms of algorithmic complexity.

### Python3

Python integers can grow beyond fixed-width integer limits, so there is no overflow concern for this problem.

I still keep the grouped approach because the main challenge is time complexity, not integer capacity.

The variables have straightforward meanings:

```text
start  -> current group beginning
commas -> commas per number
answer -> total commas
```

For each group, I calculate the number of values and multiply it by the common comma count.

Python makes this solution quite short because the arithmetic directly matches the idea.

### Go

The Go solution uses `int64` because `n` can be as large as `10^15`.

I use the same group boundaries as the other solutions.

The loop processes:

```text
1000 - 999999
1000000 - 999999999
1000000000 - ...
```

rather than individual values.

Using `int64` also makes the intended numeric range explicit and avoids depending on the size of the platform's `int` type.

## Examples

### Example 1

Input:

```text
n = 1002
```

Numbers with commas:

```text
1,000
1,001
1,002
```

There are `3` numbers, and each has `1` comma.

```text
3 × 1 = 3
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

All numbers have fewer than four digits.

```text
1 - 998
   0 commas
```

So:

```text
998 × 0 = 0
```

Output:

```text
0
```

### Example 3

Input:

```text
n = 1,000,000
```

The groups are:

```text
1 - 999
    0 commas

1,000 - 999,999
    1 comma

1,000,000
    2 commas
```

The first group contributes nothing.

The second group contains:

```text
999,999 - 1,000 + 1
= 999,000
```

numbers.

So its contribution is:

```text
999,000 × 1
= 999,000
```

The last number contributes:

```text
1 × 2
= 2
```

Therefore:

```text
999,000 + 2
= 999,002
```

Output:

```text
999002
```

## How to Use / Run Locally

The repository contains the same algorithm implemented in C++, Java, JavaScript, Python3, and Go.

### C++

Save the solution as `main.cpp`.

Compile it with:

```bash
g++ -std=c++17 main.cpp -o main
```

Run it with:

```bash
./main
```

### Java

Save the solution as `Solution.java`.

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

On LeetCode, the `Solution` class is provided by the platform, so the method can be submitted directly.

### JavaScript

Save the solution as `main.js`.

Run it with Node.js:

```bash
node main.js
```

You need a recent Node.js installation.

### Python3

Save the solution as `solution.py`.

Run:

```bash
python3 solution.py
```

Python 3.8+ is recommended.

### Go

Save the solution as `main.go`.

Run:

```bash
go run main.go
```

Or build it first:

```bash
go build main.go
```

## Notes & Optimizations

The biggest optimization is avoiding an `O(n)` loop.

A direct approach would try to examine every number:

```text
1, 2, 3, 4, ..., n
```

That is impossible for values close to `10^15`.

Instead, I group numbers by how many commas they contain.

The useful boundaries are powers of `1000`:

```text
1,000
1,000,000
1,000,000,000
1,000,000,000,000
1,000,000,000,000,000
```

This means the algorithm only needs a handful of iterations.

Another important edge case is when `n < 1000`.

For example:

```text
n = 998
```

The loop never starts because `start = 1000` is already greater than `n`, so the answer correctly remains `0`.

I also avoid converting numbers to strings. String formatting would be unnecessary overhead because I already know the comma count from the digit group.

The result is an `O(log₁₀ n)` time and `O(1)` space solution, which is easily fast enough for the given constraint.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
