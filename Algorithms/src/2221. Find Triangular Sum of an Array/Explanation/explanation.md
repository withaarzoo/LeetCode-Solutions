# Find Triangular Sum of an Array — README

**Problem Title:** 2221. Find Triangular Sum of an Array

---

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
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

We are given a 0-indexed integer array `nums` where each `nums[i]` is a digit in range `[0,9]`. The *triangular sum* is obtained by repeatedly replacing the array with a new array of pairwise sums modulo 10 until only one element remains. Return that final single value.

Concretely:

1. If `nums` has length `n == 1`, return `nums[0]`.
2. Otherwise, produce a new array `newNums` of length `n-1` where `newNums[i] = (nums[i] + nums[i+1]) % 10` for `0 <= i < n-1`.
3. Replace `nums` with `newNums` and repeat until one element remains.

---

## Constraints

* `1 <= nums.length <= 1000`
* `0 <= nums[i] <= 9`

---

## Intuition

I thought about the transformation and saw it always reduces the array length by one using pairwise sums modulo 10. Instead of allocating a new array every step, I realized I can compute the new values directly in the front portion of the same array (i.e., in-place). This saves memory and keeps code simple.

---

## Approach

1. Start with the full array length `len = n`.
2. For `len` from `n` down to `2`, update `nums[i] = (nums[i] + nums[i+1]) % 10` for all `i` from `0` to `len-2`. This computes the next-level array in-place at `nums[0..len-2]`.
3. After completing the loop, `nums[0]` is the triangular sum — return it.

This effectively simulates the process but avoids extra arrays and is easy to reason about.

---

## Data Structures Used

* **Array (in-place modification)** — we overwrite the beginning of the input array to represent the reduced array at each step.

---

## Operations & Behavior Summary

* **Pairwise addition modulo 10**: `(a + b) % 10` for each adjacent pair.
* **In-place overwrite**: Write the new value at the smaller index (`i`) while still being able to read `nums[i+1]` because it hasn't been overwritten yet in the same pass.
* **Iterative shrink**: Reduce effective length by 1 each iteration.

---

## Complexity

* **Time Complexity:** `O(n^2)` where `n = nums.length`.

  * Explanation: In the first pass we perform `n-1` operations, then `n-2`, ..., `1`. Sum = `n(n-1)/2 = O(n^2)`.
* **Space Complexity:** `O(1)` extra space.

  * Explanation: We modify the input array in-place and do not use additional arrays proportional to `n`.

---

## Multi-language Solutions

> All solutions implement the in-place iterative approach described above.

---

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

/*
 * Triangular Sum (in-place method)
 * Time: O(n^2), Space: O(1) extra
 */
class Solution {
public:
    int triangularSum(vector<int>& nums) {
        int n = nums.size();
        // Reduce the effective length from n to 1
        for (int len = n; len > 1; --len) {
            // Compute new values for indices 0 .. len-2 in-place
            for (int i = 0; i < len - 1; ++i) {
                nums[i] = (nums[i] + nums[i + 1]) % 10;
            }
        }
        return nums[0];
    }
};

// Example main to run locally
int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    int n;
    if (!(cin >> n)) return 0;
    vector<int> nums(n);
    for (int i = 0; i < n; ++i) cin >> nums[i];
    Solution sol;
    cout << sol.triangularSum(nums) << "\n";
    return 0;
}
```

---

### Java

```java
import java.util.*;

public class Solution {
    /**
     * Triangular Sum (in-place)
     * Time: O(n^2), Space: O(1) extra
     */
    public int triangularSum(int[] nums) {
        int n = nums.length;
        for (int len = n; len > 1; --len) {
            for (int i = 0; i < len - 1; ++i) {
                nums[i] = (nums[i] + nums[i + 1]) % 10;
            }
        }
        return nums[0];
    }

    // Example main to run locally
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        if (!sc.hasNextInt()) return;
        int n = sc.nextInt();
        int[] nums = new int[n];
        for (int i = 0; i < n; ++i) nums[i] = sc.nextInt();
        Solution s = new Solution();
        System.out.println(s.triangularSum(nums));
        sc.close();
    }
}
```

---

### JavaScript

```javascript
/**
 * Triangular Sum (in-place)
 * Time: O(n^2), Space: O(1) extra
 *
 * Example usage:
 *   node triangularSum.js 5 1 2 3 4 5
 * where first arg is n and next n args are array elements.
 */

function triangularSum(nums) {
    for (let len = nums.length; len > 1; --len) {
        for (let i = 0; i < len - 1; ++i) {
            nums[i] = (nums[i] + nums[i + 1]) % 10;
        }
    }
    return nums[0];
}

// If run as script, read arguments
if (require.main === module) {
    const args = process.argv.slice(2).map(Number);
    if (args.length === 0) process.exit(0);
    const n = args[0];
    const nums = args.slice(1, 1 + n);
    console.log(triangularSum(nums));
}

module.exports = triangularSum;
```

---

### Python3

```python
# Triangular Sum (in-place)
# Time: O(n^2), Space: O(1) extra

from typing import List
import sys

class Solution:
    def triangularSum(self, nums: List[int]) -> int:
        n = len(nums)
        for length in range(n, 1, -1):
            for i in range(length - 1):
                nums[i] = (nums[i] + nums[i + 1]) % 10
        return nums[0]

# Example run: provide n and the array via stdin
if __name__ == "__main__":
    data = sys.stdin.read().strip().split()
    if not data:
        sys.exit(0)
    it = iter(data)
    n = int(next(it))
    nums = [int(next(it)) for _ in range(n)]
    s = Solution()
    print(s.triangularSum(nums))
```

---

### Go

```go
package main

import (
 "bufio"
 "fmt"
 "os"
)

/*
Triangular Sum (in-place)
Time: O(n^2), Space: O(1) extra
*/
func triangularSum(nums []int) int {
 n := len(nums)
 for length := n; length > 1; length-- {
  for i := 0; i < length-1; i++ {
   nums[i] = (nums[i] + nums[i+1]) % 10
  }
 }
 return nums[0]
}

func main() {
 in := bufio.NewReader(os.Stdin)
 var n int
 if _, err := fmt.Fscan(in, &n); err != nil {
  return
 }
 nums := make([]int, n)
 for i := 0; i < n; i++ {
  fmt.Fscan(in, &nums[i])
 }
 fmt.Println(triangularSum(nums))
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Below I'll explain the common algorithm (same across all languages). I'll use a simple pseudo-line-block and then map it to language-specific lines.

### Core idea (pseudo)

```bash
n = nums.length
for len from n down to 2:
    for i from 0 to len-2:
        nums[i] = (nums[i] + nums[i+1]) % 10
return nums[0]
```

#### Why this works (intuition)

* Each iteration produces the next row of the triangular formation: each new element is the sum modulo 10 of two adjacent numbers from the previous row.
* Overwriting `nums[i]` with the new value is safe because we only read `nums[i+1]` which still holds the previous-row value (we always move `i` from left to right).
* After finishing the inner loop for a given `len`, the first `len-1` positions contain the newly formed row. We then treat `len-1` as the current length for the next outer iteration.

#### Walkthrough with an example `nums = [1, 2, 3, 4, 5]`

* `len = 5` → compute:

  * `nums[0] = (1+2)%10 = 3`
  * `nums[1] = (2+3)%10 = 5`
  * `nums[2] = (3+4)%10 = 7`
  * `nums[3] = (4+5)%10 = 9`
    Now `nums` begins with `[3,5,7,9,...]` and length is effectively 4.
* `len = 4` → compute:

  * `nums[0] = (3+5)%10 = 8`
  * `nums[1] = (5+7)%10 = 2`
  * `nums[2] = (7+9)%10 = 6`
    Now front is `[8,2,6,...]`.
* `len = 3` → `nums[0] = (8+2)%10 = 0`, `nums[1] = (2+6)%10 = 8` → `[0,8,...]`.
* `len = 2` → `nums[0] = (0+8)%10 = 8`.
* Return `8`.

Each language's provided code follows this same flow. The language-specific main / I/O wrappers are only for running locally.

---

## Examples

**Example 1**

```bash
Input: nums = [1,2,3,4,5]
Output: 8
Explanation: The triangular reduction yields final value 8 (see the step-by-step above).
```

**Example 2**

```bash
Input: nums = [5]
Output: 5
Explanation: Only one element, triangular sum is that value.
```

---

## How to use / Run locally

Assume you have saved the language-specific source into files:

### C++

1. Save as `triangular.cpp`.
2. Compile and run:

```bash
g++ -std=c++17 -O2 triangular.cpp -o triangular
# Provide input: first n, then n numbers
echo "5 1 2 3 4 5" | ./triangular
# Output: 8
```

### Java

1. Save as `Solution.java`.
2. Compile and run:

```bash
javac Solution.java
# Provide input: first n, then n numbers
echo "5 1 2 3 4 5" | java Solution
# Output: 8
```

### JavaScript (Node.js)

1. Save as `triangularSum.js`.
2. Run:

```bash
# Example passing args: n followed by n numbers
node triangularSum.js 5 1 2 3 4 5
# Output: 8
```

Or provide via stdin by modifying script to read from stdin.

### Python3

1. Save as `triangular.py`.
2. Run:

```bash
echo "5 1 2 3 4 5" | python3 triangular.py
# Output: 8
```

### Go

1. Save as `triangular.go`.
2. Run:

```bash
go run triangular.go <<EOF
5 1 2 3 4 5
EOF
# Output: 8
```

---

## Notes & Optimizations

* The in-place simulation is simple, safe, and efficient for the constraints (`n <= 1000`). It uses constant extra memory and `O(n^2)` time.
* There is a combinatorial identity: the final result equals `sum_{i=0..n-1} C(n-1, i) * nums[i] (mod 10)`. Computing binomial coefficients modulo 10 robustly is tricky because modulo 10 is not prime — multiplicative inverses don't always exist for factorials containing factors 2 or 5. You could do combinatorics modulo 2 and 5 and combine via CRT, but that's overkill for this constraint.
* If `n` were large (e.g., 10^5), we would need an entirely different strategy (e.g., combinatorial preprocessing with careful modular arithmetic or transforms). But for `n <= 1000`, `O(n^2)` is perfectly fine.
* If you want to preserve the original `nums` array, clone it first (costs `O(n)` extra memory).

---

## Author

[Md. Aarzoo Islam](https://bento.me/withaarzoo)
