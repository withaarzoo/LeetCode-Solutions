````markdown
# Minimum Operations to Make Array Sum Divisible by K

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
- [How to use / Run locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

You are given an integer array `nums` and an integer `k`.

In **one operation** you can:
- choose any index `i`
- and replace `nums[i]` with `nums[i] - 1`.

Each operation decreases the **array sum** by exactly `1`.

You can do this operation any number of times.

Your task is to return the **minimum number of operations** required to make the **sum of the array divisible by `k`**.

---

## Constraints

From the problem statement:

- `1 <= nums.length <= 1000`
- `1 <= nums[i] <= 1000`
- `1 <= k <= 100`

So the array is not very large, but we still want an `O(n)` solution.

---

## Intuition

When I read the operation carefully, I noticed something important:

- Every operation **only reduces the sum by 1**.
- It doesn’t matter **which index** I choose; the total sum always goes down by exactly `1`.

So if the original sum is `S` and I do `x` operations, the new sum becomes:

\[
\text{new\_sum} = S - x
\]

I want:

\[
(S - x) \bmod k = 0
\]

This means:

\[
S \bmod k = x \bmod k
\]

Let me call:

\[
r = S \bmod k
\]

Then `x % k` must be equal to `r`.  
The **smallest non-negative** `x` with this property is simply:

\[
x = r
\]

So I realized the answer is just:

> **Minimum operations = (sum of array) % k**

That’s it. No DP, no greedy, just pure math and remainder logic.  
Simple sa observation, but very powerful 😄.

---

## Approach

Step-by-step, my approach is:

1. **Compute the sum** of all elements in `nums`.  
   Let this sum be `S`.

2. **Compute the remainder** of `S` when divided by `k`:
   ```text
   r = S % k
````

3. If `r == 0`, the sum is already divisible by `k`.
   → I don’t need any operations, so answer is `0`.

4. Otherwise, I need to reduce the sum by exactly `r`.

   * Because after `r` operations, new sum will be `S - r`
   * And `S - r` will be divisible by `k`.

5. So I simply **return `r`**.

Why is it always possible to do `r` operations?

* Each operation reduces the sum by 1.
* I can apply operations on any element multiple times.
* Maximum I can reduce is turning everything to zero (`S` operations).
* Since `0 <= r < k` and `S` is at least `1`, we can always do `r` operations.

So the formula is always valid.

---

## Data Structures Used

* I only use:

  * A **few integer variables** (`sum`, `remainder`) to calculate the result.

No extra arrays, sets, maps, or complex data structures.

So memory usage is constant.

---

## Operations & Behavior Summary

* **Input:** Array `nums`, integer `k`.
* **Operation allowed:**

  * pick any index `i`
  * do `nums[i] = nums[i] - 1`
  * total sum decreases by `1`.
* **Goal:**

  * Make the **final array sum divisible by `k`**.
* **Behavior of algorithm:**

  * Count the total sum.
  * Compute `sum % k`.
  * Return this remainder as the minimum number of operations.

---

## Complexity

* **Time Complexity: `O(n)`**

  * I go through the array **once** to compute the sum.
  * Here, `n` = `nums.length`.

* **Space Complexity: `O(1)`**

  * I use only a few variables (`sum`, `remainder`).
  * No additional data structure that grows with `n`.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minOperations(vector<int>& nums, int k) {
        long long sum = 0;
        
        // Step 1: Calculate total sum of the array
        for (int x : nums) {
            sum += x;
        }
        
        // Step 2: The minimum number of operations required
        // is simply the remainder when sum is divided by k.
        int remainder = sum % k;
        
        // Step 3: Return remainder (0 means already divisible)
        return remainder;
    }
};
```

---

### Java

```java
class Solution {
    public int minOperations(int[] nums, int k) {
        long sum = 0;
        
        // Step 1: Calculate total sum of the array
        for (int x : nums) {
            sum += x;
        }
        
        // Step 2: Minimum operations is sum % k
        int remainder = (int)(sum % k);
        
        // Step 3: Return the remainder
        return remainder;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var minOperations = function(nums, k) {
    let sum = 0;
    
    // Step 1: Calculate total sum of the array
    for (const x of nums) {
        sum += x;
    }
    
    // Step 2: Minimum operations is sum % k
    const remainder = sum % k;
    
    // Step 3: Return remainder
    return remainder;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def minOperations(self, nums: List[int], k: int) -> int:
        # Step 1: Calculate total sum of the array
        total_sum = sum(nums)
        
        # Step 2: Minimum operations is total_sum % k
        remainder = total_sum % k
        
        # Step 3: Return remainder
        return remainder
```

---

### Go

```go
package main

func minOperations(nums []int, k int) int {
    var sum int64 = 0
    
    // Step 1: Calculate total sum of the array
    for _, x := range nums {
        sum += int64(x)
    }
    
    // Step 2: Minimum operations equals sum % k
    remainder := int(sum % int64(k))
    
    // Step 3: Return remainder
    return remainder
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Here I explain the logic using the **same steps** for all languages.
Only syntax is different, but idea is exactly the same.

### 1. Compute the sum

**Example in C++:**

```cpp
long long sum = 0;
for (int x : nums) {
    sum += x;
}
```

* I start with `sum = 0`.
* I loop over each element `x` in `nums`.
* I keep adding each `x` to `sum`.
* I use `long long` to avoid overflow when numbers are big.

Exactly same thing happens in other languages:

* **Java:**

  ```java
  long sum = 0;
  for (int x : nums) {
      sum += x;
  }
  ```

* **JavaScript:**

  ```javascript
  let sum = 0;
  for (const x of nums) {
      sum += x;
  }
  ```

* **Python:**

  ```python
  total_sum = sum(nums)
  ```

* **Go:**

  ```go
  var sum int64 = 0
  for _, x := range nums {
      sum += int64(x)
  }
  ```

---

### 2. Take modulo with `k`

**C++:**

```cpp
int remainder = sum % k;
```

* This gives me `remainder = S % k`.
* If `remainder == 0`, sum is already divisible by `k`.
* If `remainder > 0`, I need exactly that many operations.

Other languages:

* **Java:** `int remainder = (int)(sum % k);`
* **JavaScript:** `const remainder = sum % k;`
* **Python:** `remainder = total_sum % k`
* **Go:** `remainder := int(sum % int64(k))`

---

### 3. Return the remainder

Finally:

```cpp
return remainder;
```

Same for all languages:

* Just return that remainder as the **minimum number of operations**.

Why?

* Each operation reduces sum by 1.
* After `remainder` operations, new sum becomes `S - remainder`.
* `(S - remainder) % k = 0` by our modulo equation.

So this is both **correct** and **minimal**.

---

## Examples

### Example 1

**Input:**

```text
nums = [3, 9, 7], k = 5
```

* `sum = 3 + 9 + 7 = 19`
* `remainder = 19 % 5 = 4`
* Minimum operations = `4`

One possible sequence:

* Decrease `9` four times → `9 → 5`
* New array: `[3, 5, 7]`
* New sum: `3 + 5 + 7 = 15`
* `15 % 5 = 0` ✅

---

### Example 2

**Input:**

```text
nums = [4, 1, 3], k = 4
```

* `sum = 4 + 1 + 3 = 8`
* `remainder = 8 % 4 = 0`
* Minimum operations = `0`

Array sum is already divisible by `4`, so no changes needed.

---

### Example 3

**Input:**

```text
nums = [3, 2], k = 6
```

* `sum = 3 + 2 = 5`
* `remainder = 5 % 6 = 5`
* Minimum operations = `5`

We need to reduce the sum by `5` to make it `0`, which is divisible by `6`.

One possible way:

* Decrease `3` three times → `3 → 0`
* Decrease `2` two times → `2 → 0`
* New array = `[0, 0]`, sum = `0`, divisible by `6`.

---

## How to use / Run locally

### C++

```bash
g++ -std=c++17 -O2 main.cpp -o main
./main
```

Inside `main.cpp`:

```cpp
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int minOperations(vector<int>& nums, int k) {
        long long sum = 0;
        for (int x : nums) sum += x;
        return sum % k;
    }
};

int main() {
    vector<int> nums = {3, 9, 7};
    int k = 5;
    Solution sol;
    cout << sol.minOperations(nums, k) << endl; // Output: 4
    return 0;
}
```

---

### Java

```bash
javac Main.java
java Main
```

`Main.java`:

```java
public class Main {
    public static void main(String[] args) {
        int[] nums = {3, 9, 7};
        int k = 5;
        Solution sol = new Solution();
        System.out.println(sol.minOperations(nums, k)); // Output: 4
    }
}

class Solution {
    public int minOperations(int[] nums, int k) {
        long sum = 0;
        for (int x : nums) sum += x;
        return (int)(sum % k);
    }
}
```

---

### JavaScript (Node.js)

```bash
node main.js
```

`main.js`:

```javascript
var minOperations = function(nums, k) {
    let sum = 0;
    for (const x of nums) sum += x;
    return sum % k;
};

console.log(minOperations([3, 9, 7], 5)); // Output: 4
```

---

### Python3

```bash
python3 main.py
```

`main.py`:

```python
from typing import List

class Solution:
    def minOperations(self, nums: List[int], k: int) -> int:
        return sum(nums) % k

if __name__ == "__main__":
    sol = Solution()
    print(sol.minOperations([3, 9, 7], 5))  # Output: 4
```

---

### Go

```bash
go run main.go
```

`main.go`:

```go
package main

import "fmt"

func minOperations(nums []int, k int) int {
    var sum int64 = 0
    for _, x := range nums {
        sum += int64(x)
    }
    return int(sum % int64(k))
}

func main() {
    nums := []int{3, 9, 7}
    k := 5
    fmt.Println(minOperations(nums, k)) // Output: 4
}
```

---

## Notes & Optimizations

* The key optimization here is **realizing the math**:

  * Instead of simulating operations (which would be very slow),
  * We directly compute the number of operations using `sum % k`.
* This turns the problem into a **simple `O(n)` scan**.
* Space is **constant**, so it’s very memory efficient.
* This approach is optimal for the given constraints and will easily pass all test cases.

If `k` or `nums[i]` were extremely large (like `1e18`), I would be a bit more careful with using 64-bit integers everywhere to avoid overflow — but here we are safe with `long long` / `int64`.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
