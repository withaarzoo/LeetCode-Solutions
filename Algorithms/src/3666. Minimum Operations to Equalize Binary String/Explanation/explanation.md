# 3666. Minimum Operations to Equalize Binary String

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

---

## Problem Summary

You are given a binary string `s` and an integer `k`.

In one operation, you must choose exactly `k` different indices and flip them:

* '0' becomes '1'
* '1' becomes '0'

Return the minimum number of operations required to make the entire string equal to '1'.
If it is not possible, return -1.

---

## Constraints

* 1 <= s.length <= 10^5
* s[i] is either '0' or '1'
* 1 <= k <= s.length

---

## Intuition

When I first looked at the problem, I realized that every operation flips exactly `k` characters.

That means in each operation:

* `k` positions change
* `n - k` positions remain unchanged

So instead of thinking position by position, I started thinking globally.

Let:

* `n = length of string`
* `zero = number of '0'`
* `one = n - zero`

Every '0' must be flipped odd number of times.
Every '1' must be flipped even number of times.

So the real question becomes:
Can I distribute flips across operations so that all zeros turn into ones and ones stay valid?

Then I noticed something important:
The answer depends on whether total operations `m` is odd or even.

So I handled two cases separately:

* Odd number of operations
* Even number of operations

And I computed the minimum valid `m` for both.

---

## Approach

Let:

* `n = s.length`
* `zero = count of '0'`
* `one = n - zero`
* `base = n - k`

### Step 1: Edge Case

If `zero == 0`, return 0.

If `n == k`:

* Each operation flips entire string
* If all zeros → return 1
* If already all ones → return 0
* Otherwise → impossible

---

### Step 2: Try Odd Number of Operations

Minimum operations required must satisfy:

1. Enough flips for zeros:
   m >= ceil(zero / k)

2. Enough non-flips for ones:
   m >= ceil(one / base)

Take maximum of both.
Then make `m` odd.

Also parity condition must match:

(k % 2) == (zero % 2)

---

### Step 3: Try Even Number of Operations

Conditions:

1. m >= ceil(zero / k)
2. m >= ceil(zero / base)

Then make `m` even.

Also:

zero must be even

---

### Step 4: Take Minimum Valid Answer

Check both cases.
Return smallest valid value.
If none valid → return -1.

---

## Data Structures Used

* Integer counters
* Basic arithmetic operations

No extra data structures are required.

---

## Operations & Behavior Summary

Each operation:

* Flips exactly `k` bits
* Leaves `n - k` bits unchanged

After `m` operations:

* Total flips = m × k
* Total non-flips = m × (n - k)

We ensure distribution satisfies:

* Zeros flipped odd times
* Ones flipped even times

---

## Complexity

Time Complexity: O(n)

* We scan the string once to count zeros.

Space Complexity: O(1)

* Only a few integer variables are used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minOperations(string s, int k) {
        int n = s.length();
        int zero = 0;

        for (char c : s)
            if (c == '0') zero++;

        if (zero == 0) return 0;

        if (n == k) {
            if (zero == n) return 1;
            return -1;
        }

        int one = n - zero;
        int base = n - k;

        int ans = INT_MAX;

        if ((k % 2) == (zero % 2)) {
            long long m = max(
                (zero + k - 1) / k,
                (one + base - 1) / base
            );
            if (m % 2 == 0) m++;
            ans = min(ans, (int)m);
        }

        if (zero % 2 == 0) {
            long long m = max(
                (zero + k - 1) / k,
                (zero + base - 1) / base
            );
            if (m % 2 == 1) m++;
            ans = min(ans, (int)m);
        }

        return ans == INT_MAX ? -1 : ans;
    }
};
```

### Java

```java
class Solution {
    public int minOperations(String s, int k) {
        int n = s.length();
        int zero = 0;

        for (char c : s.toCharArray())
            if (c == '0') zero++;

        if (zero == 0) return 0;

        if (n == k) {
            if (zero == n) return 1;
            return -1;
        }

        int one = n - zero;
        int base = n - k;

        long ans = Long.MAX_VALUE;

        if ((k % 2) == (zero % 2)) {
            long m = Math.max(
                (zero + k - 1) / k,
                (one + base - 1) / base
            );
            if (m % 2 == 0) m++;
            ans = Math.min(ans, m);
        }

        if (zero % 2 == 0) {
            long m = Math.max(
                (zero + k - 1) / k,
                (zero + base - 1) / base
            );
            if (m % 2 == 1) m++;
            ans = Math.min(ans, m);
        }

        return ans == Long.MAX_VALUE ? -1 : (int)ans;
    }
}
```

### JavaScript

```javascript
var minOperations = function(s, k) {
    const n = s.length;
    let zero = 0;

    for (let c of s)
        if (c === '0') zero++;

    if (zero === 0) return 0;

    if (n === k) {
        if (zero === n) return 1;
        return -1;
    }

    const one = n - zero;
    const base = n - k;

    let ans = Infinity;

    if ((k % 2) === (zero % 2)) {
        let m = Math.max(
            Math.ceil(zero / k),
            Math.ceil(one / base)
        );
        if (m % 2 === 0) m++;
        ans = Math.min(ans, m);
    }

    if (zero % 2 === 0) {
        let m = Math.max(
            Math.ceil(zero / k),
            Math.ceil(zero / base)
        );
        if (m % 2 === 1) m++;
        ans = Math.min(ans, m);
    }

    return ans === Infinity ? -1 : ans;
};
```

### Python3

```python
class Solution:
    def minOperations(self, s: str, k: int) -> int:
        n = len(s)
        zero = s.count('0')

        if zero == 0:
            return 0

        if n == k:
            if zero == n:
                return 1
            return -1

        one = n - zero
        base = n - k

        ans = float('inf')

        if (k % 2) == (zero % 2):
            m = max(
                (zero + k - 1) // k,
                (one + base - 1) // base
            )
            if m % 2 == 0:
                m += 1
            ans = min(ans, m)

        if zero % 2 == 0:
            m = max(
                (zero + k - 1) // k,
                (zero + base - 1) // base
            )
            if m % 2 == 1:
                m += 1
            ans = min(ans, m)

        return -1 if ans == float('inf') else ans
```

### Go

```go
func minOperations(s string, k int) int {
    n := len(s)
    zero := 0

    for i := 0; i < n; i++ {
        if s[i] == '0' {
            zero++
        }
    }

    if zero == 0 {
        return 0
    }

    if n == k {
        if zero == n {
            return 1
        }
        return -1
    }

    one := n - zero
    base := n - k

    ans := int(^uint(0) >> 1)

    if (k%2) == (zero%2) {
        m := max(
            (zero+k-1)/k,
            (one+base-1)/base,
        )
        if m%2 == 0 {
            m++
        }
        if m < ans {
            ans = m
        }
    }

    if zero%2 == 0 {
        m := max(
            (zero+k-1)/k,
            (zero+base-1)/base,
        )
        if m%2 == 1 {
            m++
        }
        if m < ans {
            ans = m
        }
    }

    if ans == int(^uint(0)>>1) {
        return -1
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. Count number of zeros.
2. If zero is 0, return 0.
3. If n equals k, handle full flip special case.
4. Compute base = n - k.
5. Try odd operations case.
6. Try even operations case.
7. Adjust parity accordingly.
8. Return minimum valid answer.

---

## Examples

Example 1:
Input: s = "110", k = 1
Output: 1

Example 2:
Input: s = "0101", k = 3
Output: 2

Example 3:
Input: s = "101", k = 2
Output: -1

---

## How to use / Run locally

1. Copy the desired language code.
2. Paste into LeetCode editor.
3. Run with custom test cases.

For local run:

* C++: g++ file.cpp
* Java: javac Solution.java
* Python: python file.py
* Go: go run file.go

---

## Notes & Optimizations

* Only one pass through string is required.
* No extra memory used.
* Handles large constraints efficiently.
* Carefully manage parity logic.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
