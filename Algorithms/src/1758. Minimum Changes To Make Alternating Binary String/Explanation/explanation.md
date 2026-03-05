# Problem Title

1. Minimum Changes To Make Alternating Binary String

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

## Problem Summary

You are given a binary string `s` consisting only of characters `0` and `1`.

In one operation, you can change any `0` to `1` or `1` to `0`.

A string is called alternating if no two adjacent characters are equal.

Examples of alternating strings:

0101
1010

Your task is to return the minimum number of operations required to make the given string alternating.

## Constraints

1 <= s.length <= 10^4

s[i] is either '0' or '1'

## Intuition

When I first looked at the problem, I realized that an alternating binary string can only have two possible patterns.

Pattern 1 starts with 0:

010101...

Pattern 2 starts with 1:

101010...

So instead of trying many combinations, I only need to compare the given string with these two patterns.

While scanning the string, I count how many characters do not match each pattern. Every mismatch means I must flip that character.

Finally, I return the minimum number of changes between the two patterns.

## Approach

1. Initialize two counters.

   * One for pattern starting with 0.
   * One for pattern starting with 1.

2. Traverse the string once.

3. For each index:

   * If index is even

     * Pattern1 expects '0'
     * Pattern2 expects '1'
   * If index is odd

     * Pattern1 expects '1'
     * Pattern2 expects '0'

4. If the current character does not match the expected character, increment the mismatch counter.

5. After scanning the entire string, return the minimum of the two counters.

## Data Structures Used

No special data structures are required.

We only use a few integer variables to count mismatches.

## Operations & Behavior Summary

* Traverse the string once
* Compare characters with expected alternating pattern
* Count mismatches
* Return minimum mismatch count

## Complexity

Time Complexity

O(n)

Where n is the length of the string. The string is scanned exactly once.

Space Complexity

O(1)

Only constant extra variables are used.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minOperations(string s) {
        int startWith0 = 0;
        int startWith1 = 0;

        for(int i = 0; i < s.length(); i++){

            char expected0 = (i % 2 == 0) ? '0' : '1';
            char expected1 = (i % 2 == 0) ? '1' : '0';

            if(s[i] != expected0) startWith0++;
            if(s[i] != expected1) startWith1++;
        }

        return min(startWith0, startWith1);
    }
};
```

### Java

```java
class Solution {
    public int minOperations(String s) {

        int startWith0 = 0;
        int startWith1 = 0;

        for(int i = 0; i < s.length(); i++){

            char expected0 = (i % 2 == 0) ? '0' : '1';
            char expected1 = (i % 2 == 0) ? '1' : '0';

            if(s.charAt(i) != expected0) startWith0++;
            if(s.charAt(i) != expected1) startWith1++;
        }

        return Math.min(startWith0, startWith1);
    }
}
```

### JavaScript

```javascript
var minOperations = function(s) {

    let startWith0 = 0;
    let startWith1 = 0;

    for(let i = 0; i < s.length; i++){

        let expected0 = (i % 2 === 0) ? '0' : '1';
        let expected1 = (i % 2 === 0) ? '1' : '0';

        if(s[i] !== expected0) startWith0++;
        if(s[i] !== expected1) startWith1++;
    }

    return Math.min(startWith0, startWith1);
};
```

### Python3

```python
class Solution:
    def minOperations(self, s: str) -> int:

        startWith0 = 0
        startWith1 = 0

        for i in range(len(s)):

            expected0 = '0' if i % 2 == 0 else '1'
            expected1 = '1' if i % 2 == 0 else '0'

            if s[i] != expected0:
                startWith0 += 1

            if s[i] != expected1:
                startWith1 += 1

        return min(startWith0, startWith1)
```

### Go

```go
func minOperations(s string) int {

    startWith0 := 0
    startWith1 := 0

    for i := 0; i < len(s); i++ {

        var expected0 byte
        var expected1 byte

        if i%2 == 0 {
            expected0 = '0'
            expected1 = '1'
        } else {
            expected0 = '1'
            expected1 = '0'
        }

        if s[i] != expected0 {
            startWith0++
        }

        if s[i] != expected1 {
            startWith1++
        }
    }

    if startWith0 < startWith1 {
        return startWith0
    }

    return startWith1
}
```

## Step-by-step Detailed Explanation

Step 1

Initialize two counters that will store mismatch counts for both possible alternating patterns.

startWith0 counts operations when the pattern begins with 0.

startWith1 counts operations when the pattern begins with 1.

Step 2

Traverse the entire string using a loop.

For every index, we determine what character should appear according to both patterns.

Step 3

If the index is even:

Pattern starting with 0 expects 0
Pattern starting with 1 expects 1

If the index is odd:

Pattern starting with 0 expects 1
Pattern starting with 1 expects 0

Step 4

Compare the actual character with the expected one.

If they do not match, it means we must flip that character, so we increase the mismatch counter.

Step 5

After finishing the loop, we compare both mismatch counts.

The minimum of the two is the answer because it represents the smallest number of flips needed to make the string alternating.

## Examples

Example 1

Input

s = "0100"

Output

1

Explanation

Changing the last character produces "0101", which is alternating.

Example 2

Input

s = "10"

Output

0

Explanation

The string is already alternating.

Example 3

Input

s = "1111"

Output

2

Explanation

Two flips are required to convert the string to either "0101" or "1010".

## How to use / Run locally

Clone the repository

```bash
git clone https://github.com/your-username/leetcode-solutions.git
```

Navigate to the project directory

```bash
cd leetcode-solutions
```

Compile and run depending on the language.

Example for C++

```bash
g++ solution.cpp
./a.out
```

Example for Python

```bash
python solution.py
```

## Notes & Optimizations

Only one pass through the string is required.

No extra memory structures are needed.

This makes the algorithm efficient even for the maximum constraint of 10^4 characters.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
