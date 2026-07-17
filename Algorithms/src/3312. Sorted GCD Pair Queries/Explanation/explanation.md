# 3312. Sorted GCD Pair Queries

A beginner-friendly and optimized solution for **LeetCode 3312 - Sorted GCD Pair Queries**. This repository explains the idea behind the algorithm, walks through the approach step by step, and provides solutions in **C++, Java, JavaScript, Python3, and Go**. The solution uses **number theory**, **frequency counting**, **inclusion-exclusion**, **prefix sums**, and **binary search** to answer GCD pair queries efficiently.

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

You are given an integer array `nums` and another array called `queries`.

For every possible pair `(i, j)` where `i < j`, calculate the **Greatest Common Divisor (GCD)** of the two numbers. Collect all these GCD values into a new array and sort it in ascending order.

Each query represents an index inside this sorted GCD array. Your task is to return the value stored at that index.

Since the number of pairs can be extremely large, generating and sorting the entire array is not practical. The goal is to answer every query efficiently using an optimized algorithm.

---

## Constraints

| Constraint | Value |
|------------|-------|
| `2 <= nums.length <= 100000` | Number of elements |
| `1 <= nums[i] <= 50000` | Value of each element |
| `1 <= queries.length <= 100000` | Number of queries |
| `0 <= queries[i] < n * (n - 1) / 2` | Valid query index |

---

## Intuition

The first thing I noticed was that creating every possible pair would be far too expensive. With up to `100000` numbers, the total number of pairs becomes enormous.

Instead of generating every GCD, I started thinking from the opposite direction.

What if I count how many pairs have GCD exactly `1`, how many have GCD exactly `2`, how many have GCD exactly `3`, and so on?

If I already know these counts, then I also know what the sorted GCD array looks like without actually building it.

After that, answering a query becomes nothing more than finding which GCD value contains the requested position.

---

## Approach

1. Find the maximum value present in the array.
2. Count how many times every number appears.
3. For every possible divisor, count how many numbers are divisible by it.
4. Use the combination formula to calculate how many pairs are divisible by that divisor.
5. Remove pairs that actually belong to larger GCD values using inclusion-exclusion.
6. Store the number of pairs having each exact GCD.
7. Build a prefix sum over these counts.
8. Use binary search for every query to locate the corresponding GCD value.

This avoids generating every pair while still answering every query efficiently.

---

## Data Structures Used

| Data Structure | Purpose |
|---------------|---------|
| Frequency Array | Counts how many times each value appears |
| Integer Arrays | Stores the exact number of pairs for every possible GCD |
| Prefix Sum Array | Helps answer queries without building the sorted GCD array |
| Binary Search | Quickly finds the GCD corresponding to a query index |

---

## Operations & Behavior Summary

The algorithm works in several stages.

First, it counts the frequency of every number.

Next, for each possible divisor, it counts how many array elements are divisible by that divisor.

Using those counts, it calculates how many pairs could have that divisor as a common divisor.

Those counts still include larger GCD values, so inclusion-exclusion removes those duplicates to obtain the exact number of pairs for every GCD.

Once every exact count is known, a prefix sum is built.

Finally, every query is answered using binary search on the prefix array without constructing the huge sorted GCD array.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time | `O(M log M + Q log M)` | `M` is the maximum value in `nums` and `Q` is the number of queries. The divisor iteration uses harmonic complexity, while each query uses binary search. |
| Space | `O(M)` | Extra arrays are used for frequencies, exact GCD counts, and prefix sums. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> gcdValues(vector<int>& nums, vector<long long>& queries) {
        // Find the largest value because every helper array
        // only needs to be built up to this value.
        int mx = *max_element(nums.begin(), nums.end());

        // freq[x] = how many times value x appears.
        vector<int> freq(mx + 1, 0);
        for (int x : nums) freq[x]++;

        // exact[g] = number of pairs whose GCD is exactly g.
        vector<long long> exact(mx + 1, 0);

        // Process from largest divisor to smallest.
        for (int g = mx; g >= 1; g--) {

            // Count numbers divisible by g.
            long long cnt = 0;
            for (int m = g; m <= mx; m += g)
                cnt += freq[m];

            // Total pairs whose GCD is a multiple of g.
            long long pairs = cnt * (cnt - 1) / 2;

            // Remove pairs already assigned to larger GCDs.
            for (int m = g * 2; m <= mx; m += g)
                pairs -= exact[m];

            exact[g] = pairs;
        }

        // prefix[g] = number of pairs with GCD <= g.
        vector<long long> prefix(mx + 1, 0);
        for (int g = 1; g <= mx; g++)
            prefix[g] = prefix[g - 1] + exact[g];

        vector<int> ans;

        for (long long q : queries) {
            // First GCD whose prefix count is greater than q.
            int g = lower_bound(prefix.begin() + 1, prefix.end(), q + 1) - prefix.begin();
            ans.push_back(g);
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int[] gcdValues(int[] nums, long[] queries) {

        // Find maximum value.
        int mx = 0;
        for (int x : nums) mx = Math.max(mx, x);

        // Frequency of every value.
        int[] freq = new int[mx + 1];
        for (int x : nums) freq[x]++;

        // exact[g] = pairs having GCD exactly g.
        long[] exact = new long[mx + 1];

        // Process divisors from large to small.
        for (int g = mx; g >= 1; g--) {

            // Count numbers divisible by g.
            long cnt = 0;
            for (int m = g; m <= mx; m += g)
                cnt += freq[m];

            // Total pairs with GCD multiple of g.
            long pairs = cnt * (cnt - 1) / 2;

            // Remove larger exact GCD counts.
            for (int m = g * 2; m <= mx; m += g)
                pairs -= exact[m];

            exact[g] = pairs;
        }

        // Prefix sums.
        long[] prefix = new long[mx + 1];
        for (int g = 1; g <= mx; g++)
            prefix[g] = prefix[g - 1] + exact[g];

        int[] ans = new int[queries.length];

        for (int i = 0; i < queries.length; i++) {

            int l = 1, r = mx;

            // Binary search answer.
            while (l < r) {
                int mid = (l + r) / 2;

                if (prefix[mid] >= queries[i] + 1)
                    r = mid;
                else
                    l = mid + 1;
            }

            ans[i] = l;
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number[]} queries
 * @return {number[]}
 */
var gcdValues = function(nums, queries) {

    // Find maximum value.
    let mx = 0;
    for (const x of nums) mx = Math.max(mx, x);

    // Frequency array.
    const freq = new Array(mx + 1).fill(0);
    for (const x of nums) freq[x]++;

    // exact[g] = pairs with GCD exactly g.
    const exact = new Array(mx + 1).fill(0);

    // Compute exact counts.
    for (let g = mx; g >= 1; g--) {

        let cnt = 0;

        // Count divisible numbers.
        for (let m = g; m <= mx; m += g)
            cnt += freq[m];

        let pairs = cnt * (cnt - 1) / 2;

        // Remove larger GCD contributions.
        for (let m = g * 2; m <= mx; m += g)
            pairs -= exact[m];

        exact[g] = pairs;
    }

    // Prefix sums.
    const prefix = new Array(mx + 1).fill(0);
    for (let g = 1; g <= mx; g++)
        prefix[g] = prefix[g - 1] + exact[g];

    const ans = [];

    for (const q of queries) {

        let l = 1, r = mx;

        // Binary search.
        while (l < r) {
            const mid = (l + r) >> 1;

            if (prefix[mid] >= q + 1)
                r = mid;
            else
                l = mid + 1;
        }

        ans.push(l);
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def gcdValues(self, nums: List[int], queries: List[int]) -> List[int]:

        # Maximum value in the array.
        mx = max(nums)

        # Frequency of each value.
        freq = [0] * (mx + 1)
        for x in nums:
            freq[x] += 1

        # exact[g] = pairs having GCD exactly g.
        exact = [0] * (mx + 1)

        # Process from largest divisor.
        for g in range(mx, 0, -1):

            # Count divisible numbers.
            cnt = 0
            for m in range(g, mx + 1, g):
                cnt += freq[m]

            # Total pairs with GCD multiple of g.
            pairs = cnt * (cnt - 1) // 2

            # Remove already computed larger GCDs.
            for m in range(g * 2, mx + 1, g):
                pairs -= exact[m]

            exact[g] = pairs

        # Prefix sums.
        prefix = [0] * (mx + 1)
        for g in range(1, mx + 1):
            prefix[g] = prefix[g - 1] + exact[g]

        ans = []

        for q in queries:

            # Binary search.
            l, r = 1, mx

            while l < r:
                mid = (l + r) // 2

                if prefix[mid] >= q + 1:
                    r = mid
                else:
                    l = mid + 1

            ans.append(l)

        return ans
```

### Go

```go
func gcdValues(nums []int, queries []int64) []int {

 // Find maximum value.
 mx := 0
 for _, x := range nums {
  if x > mx {
   mx = x
  }
 }

 // Frequency array.
 freq := make([]int, mx+1)
 for _, x := range nums {
  freq[x]++
 }

 // exact[g] = pairs with GCD exactly g.
 exact := make([]int64, mx+1)

 // Process from large divisor to small.
 for g := mx; g >= 1; g-- {

  var cnt int64 = 0

  // Count divisible numbers.
  for m := g; m <= mx; m += g {
   cnt += int64(freq[m])
  }

  // Total pairs whose GCD is multiple of g.
  pairs := cnt * (cnt - 1) / 2

  // Remove larger exact GCDs.
  for m := g * 2; m <= mx; m += g {
   pairs -= exact[m]
  }

  exact[g] = pairs
 }

 // Prefix sums.
 prefix := make([]int64, mx+1)
 for g := 1; g <= mx; g++ {
  prefix[g] = prefix[g-1] + exact[g]
 }

 ans := make([]int, len(queries))

 for i, q := range queries {

  l, r := 1, mx

  // Binary search answer.
  for l < r {
   mid := (l + r) / 2

   if prefix[mid] >= q+1 {
    r = mid
   } else {
    l = mid + 1
   }
  }

  ans[i] = l
 }

 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical across all five languages. Only the syntax changes.

The solution starts by finding the largest value in the input array. This tells us the largest possible GCD we ever need to consider.

Next, a frequency array is built so that every number knows how many times it appears.

After that, every possible divisor is processed from the largest value down to `1`.

For each divisor, all of its multiples are visited.

This tells us exactly how many numbers are divisible by the current divisor.

Using the combination formula,

`count × (count - 1) / 2`

we calculate how many pairs are divisible by that divisor.

However, those pairs may actually have a larger GCD.

To fix this, we subtract all pair counts that were already assigned to larger multiples.

Because we process from large to small, those larger answers are already known.

The remaining value represents the number of pairs whose GCD is exactly equal to the current divisor.

Once every divisor has been processed, a prefix sum array is built.

Each prefix value tells us how many pairs have GCD less than or equal to that value.

Instead of building the massive sorted GCD array, we simply binary search inside the prefix sums.

The first prefix that becomes larger than the query index tells us the answer immediately.

This makes every query extremely fast.

---

## Examples

### Example 1

**Input**

```text
nums = [2,3,4]
queries = [0,2,2]
```

**Output**

```text
[1,2,2]
```

**Explanation**

The GCD values of every pair are:

```text
[1,2,1]
```

After sorting:

```text
[1,1,2]
```

The requested indices return:

```text
[1,2,2]
```

---

### Example 2

**Input**

```text
nums = [4,4,2,1]
queries = [5,3,1,0]
```

**Output**

```text
[4,2,1,1]
```

**Explanation**

The sorted GCD values become:

```text
[1,1,1,2,2,4]
```

Each query simply points to one position in this sorted list.

---

### Example 3

**Input**

```text
nums = [2,2]
queries = [0,0]
```

**Output**

```text
[2,2]
```

**Explanation**

Only one pair exists.

Its GCD is `2`, so every query returns `2`.

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone https://github.com/your-username/your-repository.git
cd your-repository
```

### C++

Compile

```bash
g++ solution.cpp -O2 -std=c++17
```

Run

```bash
./a.out
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

---

## Notes & Optimizations

- Generating every pair is impossible for the largest test cases.
- Frequency counting avoids repeatedly scanning the input array.
- Processing divisors from largest to smallest allows inclusion-exclusion to work naturally.
- Prefix sums eliminate the need to build the sorted GCD array.
- Binary search answers each query in logarithmic time.
- This approach is well suited for competitive programming because it handles very large inputs efficiently while keeping memory usage low.
- The same algorithm works across all supported programming languages without changing the core logic.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
