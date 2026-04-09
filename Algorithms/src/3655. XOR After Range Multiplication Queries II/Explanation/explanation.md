# 3655. XOR After Range Multiplication Queries II

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

## Problem Summary

You are given an integer array `nums` and a list of queries.

Each query is represented as:

```text
[l, r, k, v]
```

For every query:

* Start from index `l`
* Keep jumping by `k`
* Multiply every visited element by `v`
* Apply modulo `1e9 + 7`

After processing all queries, return the XOR of all elements.

---

## Constraints

```text
1 <= n == nums.length <= 10^5
1 <= nums[i] <= 10^9
1 <= q == queries.length <= 10^5
queries[i] = [li, ri, ki, vi]
0 <= li <= ri < n
1 <= ki <= n
1 <= vi <= 10^5
```

---

## Intuition

At first, I thought about simulating every query directly.

For every query, I can start from `l`, keep jumping by `k`, and multiply each valid position.

But this is too slow.

If `k = 1`, a single query may touch almost every element.
If there are many such queries, the total time becomes too large.

So I split queries into two groups:

1. Large `k`
2. Small `k`

For large `k`, the number of visited positions is already small, so direct simulation is fast enough.

For small `k`, many queries may share the same step size. So instead of processing each query separately, I group them by `k` and use a difference-array style approach.

---

## Approach

1. Let:

```text
limit = sqrt(n)
```

1. For every query:

   * If `k >= limit`, process it directly.
   * Otherwise, store it inside a map grouped by `k`.

2. For every small `k` group:

   * Create a `diff` array initialized with `1`
   * For query `[l, r, k, v]`

     * Multiply `diff[l]` by `v`
     * Find the next position after the last affected index
     * Multiply that next position by inverse of `v`

3. Propagate values:

```text
diff[i] *= diff[i-k]
```

1. Multiply final value into `nums[i]`

2. XOR all numbers and return answer.

---

## Data Structures Used

* Array / Vector for storing `nums`
* HashMap / Dictionary for grouping queries by `k`
* Difference array for storing multiplier effects

---

## Operations & Behavior Summary

### Large `k`

If `k` is large, then the number of affected indices is small.

Example:

```text
k = 500
```

Then only a few positions are visited.

So direct simulation is efficient.

### Small `k`

If `k` is small, many positions are affected.

Instead of processing every query separately, I:

* Group all queries with same `k`
* Use a multiplicative difference array
* Use modular inverse to stop the multiplier effect

---

## Complexity

* Time Complexity: `O(q * sqrt(n) + n * sqrt(n))`
* Space Complexity: `O(n)`

Where:

* `n` = size of array
* `q` = number of queries

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    static const int MOD = 1e9 + 7;

    long long modPow(long long base, long long exp) {
        long long result = 1;

        while (exp > 0) {
            if (exp & 1) {
                result = (result * base) % MOD;
            }

            base = (base * base) % MOD;
            exp >>= 1;
        }

        return result;
    }

    long long modInverse(long long x) {
        return modPow(x, MOD - 2);
    }

    int xorAfterQueries(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        int limit = sqrt(n) + 1;

        unordered_map<int, vector<vector<int>>> smallQueries;

        for (auto &q : queries) {
            int l = q[0];
            int r = q[1];
            int k = q[2];
            int v = q[3];

            if (k >= limit) {
                for (int i = l; i <= r; i += k) {
                    nums[i] = (1LL * nums[i] * v) % MOD;
                }
            } else {
                smallQueries[k].push_back(q);
            }
        }

        for (auto &[k, group] : smallQueries) {
            vector<long long> diff(n, 1);

            for (auto &q : group) {
                int l = q[0];
                int r = q[1];
                int v = q[3];

                diff[l] = (diff[l] * v) % MOD;

                int steps = (r - l) / k;
                int nextPos = l + (steps + 1) * k;

                if (nextPos < n) {
                    diff[nextPos] = (diff[nextPos] * modInverse(v)) % MOD;
                }
            }

            for (int i = 0; i < n; i++) {
                if (i >= k) {
                    diff[i] = (diff[i] * diff[i - k]) % MOD;
                }

                nums[i] = (1LL * nums[i] * diff[i]) % MOD;
            }
        }

        int answer = 0;

        for (int num : nums) {
            answer ^= num;
        }

        return answer;
    }
};
```

### Java

```java
class Solution {
    private static final int MOD = 1_000_000_007;

    private long modPow(long base, long exp) {
        long result = 1;

        while (exp > 0) {
            if ((exp & 1) == 1) {
                result = (result * base) % MOD;
            }

            base = (base * base) % MOD;
            exp >>= 1;
        }

        return result;
    }

    private long modInverse(long x) {
        return modPow(x, MOD - 2);
    }

    public int xorAfterQueries(int[] nums, int[][] queries) {
        int n = nums.length;
        int limit = (int)Math.sqrt(n) + 1;

        Map<Integer, List<int[]>> smallQueries = new HashMap<>();

        for (int[] q : queries) {
            int l = q[0];
            int r = q[1];
            int k = q[2];
            int v = q[3];

            if (k >= limit) {
                for (int i = l; i <= r; i += k) {
                    nums[i] = (int)((1L * nums[i] * v) % MOD);
                }
            } else {
                smallQueries.computeIfAbsent(k, x -> new ArrayList<>()).add(q);
            }
        }

        for (Map.Entry<Integer, List<int[]>> entry : smallQueries.entrySet()) {
            int k = entry.getKey();
            List<int[]> group = entry.getValue();

            long[] diff = new long[n];
            Arrays.fill(diff, 1L);

            for (int[] q : group) {
                int l = q[0];
                int r = q[1];
                int v = q[3];

                diff[l] = (diff[l] * v) % MOD;

                int steps = (r - l) / k;
                int nextPos = l + (steps + 1) * k;

                if (nextPos < n) {
                    diff[nextPos] = (diff[nextPos] * modInverse(v)) % MOD;
                }
            }

            for (int i = 0; i < n; i++) {
                if (i >= k) {
                    diff[i] = (diff[i] * diff[i - k]) % MOD;
                }

                nums[i] = (int)((1L * nums[i] * diff[i]) % MOD);
            }
        }

        int answer = 0;

        for (int num : nums) {
            answer ^= num;
        }

        return answer;
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
    const MOD = 1000000007;
    const n = nums.length;
    const limit = Math.floor(Math.sqrt(n)) + 1;

    function modPow(base, exp) {
        let result = 1n;
        let b = BigInt(base);
        let e = BigInt(exp);
        const mod = BigInt(MOD);

        while (e > 0n) {
            if (e & 1n) {
                result = (result * b) % mod;
            }

            b = (b * b) % mod;
            e >>= 1n;
        }

        return result;
    }

    function modInverse(x) {
        return modPow(x, MOD - 2);
    }

    const smallQueries = new Map();

    for (const [l, r, k, v] of queries) {
        if (k >= limit) {
            for (let i = l; i <= r; i += k) {
                nums[i] = Number((BigInt(nums[i]) * BigInt(v)) % BigInt(MOD));
            }
        } else {
            if (!smallQueries.has(k)) {
                smallQueries.set(k, []);
            }

            smallQueries.get(k).push([l, r, v]);
        }
    }

    for (const [k, group] of smallQueries.entries()) {
        const diff = Array(n).fill(1n);

        for (const [l, r, v] of group) {
            diff[l] = (diff[l] * BigInt(v)) % BigInt(MOD);

            const steps = Math.floor((r - l) / k);
            const nextPos = l + (steps + 1) * k;

            if (nextPos < n) {
                diff[nextPos] =
                    (diff[nextPos] * modInverse(v)) % BigInt(MOD);
            }
        }

        for (let i = 0; i < n; i++) {
            if (i >= k) {
                diff[i] = (diff[i] * diff[i - k]) % BigInt(MOD);
            }

            nums[i] = Number((BigInt(nums[i]) * diff[i]) % BigInt(MOD));
        }
    }

    let answer = 0;

    for (const num of nums) {
        answer ^= num;
    }

    return answer;
};
```

### Python3

```python
class Solution:
    def xorAfterQueries(self, nums: List[int], queries: List[List[int]]) -> int:
        MOD = 10**9 + 7
        n = len(nums)
        limit = int(n ** 0.5) + 1

        def mod_pow(base, exp):
            result = 1

            while exp > 0:
                if exp & 1:
                    result = (result * base) % MOD

                base = (base * base) % MOD
                exp >>= 1

            return result

        def mod_inverse(x):
            return mod_pow(x, MOD - 2)

        small_queries = {}

        for l, r, k, v in queries:
            if k >= limit:
                i = l
                while i <= r:
                    nums[i] = (nums[i] * v) % MOD
                    i += k
            else:
                if k not in small_queries:
                    small_queries[k] = []

                small_queries[k].append((l, r, v))

        for k, group in small_queries.items():
            diff = [1] * n

            for l, r, v in group:
                diff[l] = (diff[l] * v) % MOD

                steps = (r - l) // k
                next_pos = l + (steps + 1) * k

                if next_pos < n:
                    diff[next_pos] = (diff[next_pos] * mod_inverse(v)) % MOD

            for i in range(n):
                if i >= k:
                    diff[i] = (diff[i] * diff[i - k]) % MOD

                nums[i] = (nums[i] * diff[i]) % MOD

        answer = 0

        for num in nums:
            answer ^= num

        return answer
```

### Go

```go
func xorAfterQueries(nums []int, queries [][]int) int {
 const MOD int64 = 1_000_000_007

 n := len(nums)
 limit := int(math.Sqrt(float64(n))) + 1

 modPow := func(base, exp int64) int64 {
  result := int64(1)

  for exp > 0 {
   if exp&1 == 1 {
    result = (result * base) % MOD
   }

   base = (base * base) % MOD
   exp >>= 1
  }

  return result
 }

 // Modular inverse
 var modInverse func(int64) int64
 modInverse = func(x int64) int64 {
  return modPow(x, MOD-2)
 }

 smallQueries := make(map[int][][]int)

 for _, q := range queries {
  l, r, k, v := q[0], q[1], q[2], q[3]

  if k >= limit {
   for i := l; i <= r; i += k {
    nums[i] = int((int64(nums[i]) * int64(v)) % MOD)
   }
  } else {
   smallQueries[k] = append(smallQueries[k], q)
  }
 }

 for k, group := range smallQueries {
  diff := make([]int64, n)

  for i := 0; i < n; i++ {
   diff[i] = 1
  }

  for _, q := range group {
   l, r, v := q[0], q[1], q[3]

   diff[l] = (diff[l] * int64(v)) % MOD

   steps := (r - l) / k
   nextPos := l + (steps+1)*k

   if nextPos < n {
    diff[nextPos] = (diff[nextPos] * modInverse(int64(v))) % MOD
   }
  }

  for i := 0; i < n; i++ {
   if i >= k {
    diff[i] = (diff[i] * diff[i-k]) % MOD
   }

   nums[i] = int((int64(nums[i]) * diff[i]) % MOD)
  }
 }

 answer := 0

 for _, num := range nums {
  answer ^= num
 }

 return answer
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Suppose:

```text
nums = [2, 3, 1, 5, 4]
queries = [[1,4,2,3], [0,2,1,2]]
```

### First Query

```text
[1,4,2,3]
```

Meaning:

* Start from index `1`
* Jump by `2`
* Multiply by `3`

Affected positions:

```text
1, 3
```

Updated array:

```text
[2, 9, 1, 15, 4]
```

### Second Query

```text
[0,2,1,2]
```

Affected positions:

```text
0, 1, 2
```

Updated array:

```text
[4, 18, 2, 15, 4]
```

### Final XOR

```text
4 ^ 18 ^ 2 ^ 15 ^ 4 = 31
```

### Why Difference Array Works

Suppose:

```text
query = [1, 5, 2, 3]
```

Affected positions:

```text
1, 3, 5
```

We do:

```text
diff[1] *= 3
```

And after the last valid position, we stop the effect:

```text
diff[nextPosition] *= inverse(3)
```

Then we propagate:

```text
diff[i] *= diff[i-k]
```

This keeps the multiplier active only on indices that are part of the same jump chain.

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

```bash
g++ -std=c++17 solution.cpp -o solution
./solution
```

```bash
javac Solution.java
java Solution
```

```bash
python solution.py
```

```bash
node solution.js
```

```bash
go run solution.go
```

---

## Notes & Optimizations

* Splitting queries into large `k` and small `k` is the main optimization.
* Direct simulation is used only when it is cheap.
* Difference array plus modular inverse makes grouped updates efficient.
* Fast exponentiation is used for modular inverse.
* Final XOR is done in one pass.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
