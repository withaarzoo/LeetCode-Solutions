# XOR After Range Multiplication Queries I

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
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

You are given:

* An integer array `nums`
* A list of queries where each query is `[li, ri, ki, vi]`

For every query:

1. Start from index `li`
2. Keep jumping by `ki`
3. Multiply every visited element by `vi`
4. Take modulo `1e9 + 7`

After processing all queries, return the XOR of all elements in the final array.

---

## Constraints

```text
1 <= n == nums.length <= 10^3
1 <= nums[i] <= 10^9
1 <= queries.length <= 10^3
queries[i] = [li, ri, ki, vi]
0 <= li <= ri < n
1 <= ki <= n
1 <= vi <= 10^5
```

---

## Intuition

I thought about directly simulating every query.

For each query `[li, ri, ki, vi]`, I can:

* Start from `li`
* Jump by `ki`
* Continue until `ri`
* Multiply every visited element by `vi`

The constraints are small enough.

Even if every query touches almost all elements, the total number of operations is still manageable.

So instead of using a complicated optimization, I can simply follow the problem statement exactly.

---

## Approach

1. Loop through every query.
2. Extract `l`, `r`, `k`, and `v`.
3. Visit indices:

```text
l, l + k, l + 2k, ... <= r
```

1. Multiply the current value by `v`.
2. Apply modulo `1e9 + 7`.
3. After all queries are processed, XOR all elements of the final array.
4. Return the XOR result.

---

## Data Structures Used

* Array / Vector / List for storing `nums`
* 2D Array / Vector for storing queries
* Integer variable for final XOR answer

No extra complex data structure is required.

---

## Operations & Behavior Summary

For every query:

```text
Index Pattern: li, li + ki, li + 2*ki, ... <= ri
```

For every visited index:

```text
nums[index] = (nums[index] * vi) % MOD
```

At the end:

```text
answer = nums[0] ^ nums[1] ^ nums[2] ^ ...
```

---

## Complexity

* Time Complexity: `O(q * n)`

  * `q` = number of queries
  * `n` = size of array
  * In worst case, every query may touch nearly all elements.

* Space Complexity: `O(1)`

  * Only a few extra variables are used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int xorAfterQueries(vector<int>& nums, vector<vector<int>>& queries) {
        const long long MOD = 1e9 + 7;

        for (auto& q : queries) {
            int l = q[0];
            int r = q[1];
            int k = q[2];
            int v = q[3];

            for (int i = l; i <= r; i += k) {
                nums[i] = (1LL * nums[i] * v) % MOD;
            }
        }

        int ans = 0;
        for (int num : nums) {
            ans ^= num;
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int xorAfterQueries(int[] nums, int[][] queries) {
        long MOD = 1000000007L;

        for (int[] q : queries) {
            int l = q[0];
            int r = q[1];
            int k = q[2];
            int v = q[3];

            for (int i = l; i <= r; i += k) {
                nums[i] = (int)((1L * nums[i] * v) % MOD);
            }
        }

        int ans = 0;
        for (int num : nums) {
            ans ^= num;
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number[][]} queries
 * @return {number}
 */
var xorAfterQueries = function(nums, queries) {
    const MOD = 1000000007n;

    for (const [l, r, k, v] of queries) {
        for (let i = l; i <= r; i += k) {
            nums[i] = Number((BigInt(nums[i]) * BigInt(v)) % MOD);
        }
    }

    let ans = 0;
    for (const num of nums) {
        ans ^= num;
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def xorAfterQueries(self, nums: List[int], queries: List[List[int]]) -> int:
        MOD = 10**9 + 7

        for l, r, k, v in queries:
            for i in range(l, r + 1, k):
                nums[i] = (nums[i] * v) % MOD

        ans = 0
        for num in nums:
            ans ^= num

        return ans
```

### Go

```go
func xorAfterQueries(nums []int, queries [][]int) int {
    const MOD int64 = 1000000007

    for _, q := range queries {
        l, r, k, v := q[0], q[1], q[2], q[3]

        for i := l; i <= r; i += k {
            nums[i] = int((int64(nums[i]) * int64(v)) % MOD)
        }
    }

    ans := 0
    for _, num := range nums {
        ans ^= num
    }

    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Suppose:

```text
nums = [2, 3, 1, 5, 4]
queries = [[1, 4, 2, 3], [0, 2, 1, 2]]
```

### Query 1: `[1, 4, 2, 3]`

This means:

* Start at index `1`
* End at index `4`
* Jump by `2`
* Multiply by `3`

Visited indices:

```text
1, 3
```

Update:

```text
nums[1] = 3 * 3 = 9
nums[3] = 5 * 3 = 15
```

Array becomes:

```text
[2, 9, 1, 15, 4]
```

### Query 2: `[0, 2, 1, 2]`

This means:

* Start at index `0`
* End at index `2`
* Jump by `1`
* Multiply by `2`

Visited indices:

```text
0, 1, 2
```

Update:

```text
nums[0] = 2 * 2 = 4
nums[1] = 9 * 2 = 18
nums[2] = 1 * 2 = 2
```

Array becomes:

```text
[4, 18, 2, 15, 4]
```

Now compute XOR:

```text
4 ^ 18 ^ 2 ^ 15 ^ 4 = 31
```

Final answer:

```text
31
```

---

## Examples

### Example 1

```text
Input:
nums = [1,1,1]
queries = [[0,2,1,4]]

Output:
4
```

Explanation:

```text
Array becomes [4,4,4]
4 ^ 4 ^ 4 = 4
```

### Example 2

```text
Input:
nums = [2,3,1,5,4]
queries = [[1,4,2,3],[0,2,1,2]]

Output:
31
```

---

## How to use / Run locally

### C++

```bash
g++ main.cpp -o main
./main
```

### Java

```bash
javac Solution.java
java Solution
```

### Python3

```bash
python solution.py
```

### JavaScript

```bash
node solution.js
```

### Go

```bash
go run solution.go
```

---

## Notes & Optimizations

* Since constraints are small, brute force simulation is enough.
* No advanced data structure is needed.
* Using modulo after every multiplication prevents overflow.
* In JavaScript, `BigInt` is used because multiplication may exceed safe integer range.
* The solution is simple, clean, and easy to understand.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
