# 1674. Minimum Moves to Make Array Complementary

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

LeetCode 1674 — Minimum Moves to Make Array Complementary is a classic array and prefix sum problem.

We are given:

* An integer array `nums`
* An integer `limit`

The array length is always even.

For every index `i`, the pair:

```text
nums[i] + nums[n - 1 - i]
```

must become the same value across the entire array.

In one move, we can replace any number with another value between `1` and `limit`.

The goal is to find the minimum number of moves required to make the array complementary.

This problem looks difficult at first because every pair can behave differently depending on the target sum. A brute-force solution becomes too slow, so we need an optimized approach using a difference array and prefix sum technique.

Relevant SEO keywords naturally connected to this problem include:

* Minimum Moves to Make Array Complementary
* LeetCode 1674 solution
* Prefix Sum algorithm
* Difference Array technique
* Competitive Programming
* DSA array problem
* Range update optimization
* Greedy + Prefix Sum approach

---

## Constraints

| Constraint                      | Value         |
| ------------------------------- | ------------- |
| `n == nums.length`              | Array size    |
| `2 <= n <= 10^5`                | Array length  |
| `1 <= nums[i] <= limit <= 10^5` | Element range |
| `n` is even                     | Guaranteed    |

---

## Intuition

When I first looked at this problem, I noticed that the array is checked in mirrored pairs.

That means:

```text
nums[i]
nums[n - 1 - i]
```

always belong together.

For every pair, there are only three possibilities:

* The target sum already exists → `0 moves`
* The target sum can be reached by changing one number → `1 move`
* Otherwise → `2 moves`

The biggest observation is that one pair can support a whole range of sums using only one move.

Instead of recalculating every target sum separately, I realized I could mark ranges efficiently using a difference array.

That turns a slow brute-force solution into an optimized linear solution.

---

## Approach

First, I process the array using mirrored pairs.

For every pair:

```text
(a, b)
```

I keep:

```text
a <= b
```

Now I analyze what sums are possible.

### Case 1 — Zero moves

If the target sum is:

```text
a + b
```

then no operation is needed.

---

### Case 2 — One move

With one replacement, I can create every sum in this range:

```text
[a + 1, b + limit]
```

This is because:

* I can lower one number to `1`
* Or increase one number to `limit`

---

### Case 3 — Two moves

Any sum outside that range requires changing both numbers.

---

After understanding this behavior, I use a difference array to efficiently update ranges instead of checking every target sum manually.

Finally, I build prefix sums and find the minimum possible moves among all target sums.

---

## Data Structures Used

| Data Structure   | Why It Was Used                          |
| ---------------- | ---------------------------------------- |
| Array            | To store the input numbers               |
| Difference Array | To apply fast range updates              |
| Prefix Sum       | To compute final move counts efficiently |

The difference array is the key optimization in this problem.

Without it, the solution would become too slow for large inputs.

---

## Operations & Behavior Summary

1. Split the array into mirrored pairs.
2. For each pair:

   * Mark the range where only 1 move is needed.
   * Mark the exact sum where 0 moves are needed.
3. Assume every pair initially needs 2 moves.
4. Apply all updates using a difference array.
5. Build prefix sums over all possible target sums.
6. Track the minimum moves found.

This avoids checking every pair against every possible sum individually.

---

## Complexity

| Type             | Complexity     | Explanation                                                          |
| ---------------- | -------------- | -------------------------------------------------------------------- |
| Time Complexity  | `O(n + limit)` | Each pair is processed once, then all possible sums are scanned once |
| Space Complexity | `O(limit)`     | Extra difference array is used for range updates                     |

Where:

* `n` = size of the array
* `limit` = maximum allowed value in the array

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minMoves(vector<int>& nums, int limit) {
        
        int n = nums.size();

        // Difference array to track move changes
        vector<int> diff(2 * limit + 2, 0);

        // Process every pair
        for (int i = 0; i < n / 2; i++) {

            int a = min(nums[i], nums[n - 1 - i]);
            int b = max(nums[i], nums[n - 1 - i]);

            // By default every sum needs 2 moves
            // We subtract moves where improvement is possible

            // For sums in [a+1, b+limit], only 1 move is needed
            diff[a + 1] -= 1;
            diff[b + limit + 1] += 1;

            // For exact sum a+b, 0 moves are needed
            diff[a + b] -= 1;
            diff[a + b + 1] += 1;
        }

        int pairs = n / 2;

        // Initially every pair contributes 2 moves
        int current = pairs * 2;

        int answer = INT_MAX;

        // Build prefix sums to compute moves for every target sum
        for (int sum = 2; sum <= 2 * limit; sum++) {

            current += diff[sum];

            answer = min(answer, current);
        }

        return answer;
    }
};
```

### Java

```java
class Solution {
    public int minMoves(int[] nums, int limit) {
        
        int n = nums.length;

        // Difference array
        int[] diff = new int[2 * limit + 2];

        // Process all pairs
        for (int i = 0; i < n / 2; i++) {

            int a = Math.min(nums[i], nums[n - 1 - i]);
            int b = Math.max(nums[i], nums[n - 1 - i]);

            // 1 move range
            diff[a + 1] -= 1;
            diff[b + limit + 1] += 1;

            // Exact sum needing 0 moves
            diff[a + b] -= 1;
            diff[a + b + 1] += 1;
        }

        int pairs = n / 2;

        // Initially assume 2 moves for every pair
        int current = pairs * 2;

        int answer = Integer.MAX_VALUE;

        // Prefix sum traversal
        for (int sum = 2; sum <= 2 * limit; sum++) {

            current += diff[sum];

            answer = Math.min(answer, current);
        }

        return answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} limit
 * @return {number}
 */
var minMoves = function(nums, limit) {
    
    const n = nums.length;

    // Difference array
    const diff = new Array(2 * limit + 2).fill(0);

    // Process all pairs
    for (let i = 0; i < n / 2; i++) {

        let a = Math.min(nums[i], nums[n - 1 - i]);
        let b = Math.max(nums[i], nums[n - 1 - i]);

        // Range where only 1 move is needed
        diff[a + 1] -= 1;
        diff[b + limit + 1] += 1;

        // Exact sum where 0 moves are needed
        diff[a + b] -= 1;
        diff[a + b + 1] += 1;
    }

    const pairs = Math.floor(n / 2);

    // Initially assume every pair needs 2 moves
    let current = pairs * 2;

    let answer = Number.MAX_SAFE_INTEGER;

    // Build prefix sums
    for (let sum = 2; sum <= 2 * limit; sum++) {

        current += diff[sum];

        answer = Math.min(answer, current);
    }

    return answer;
};
```

### Python3

```python
class Solution:
    def minMoves(self, nums: List[int], limit: int) -> int:
        
        n = len(nums)

        # Difference array
        diff = [0] * (2 * limit + 2)

        # Process every pair
        for i in range(n // 2):

            a = min(nums[i], nums[n - 1 - i])
            b = max(nums[i], nums[n - 1 - i])

            # Range where only 1 move is needed
            diff[a + 1] -= 1
            diff[b + limit + 1] += 1

            # Exact sum where 0 moves are needed
            diff[a + b] -= 1
            diff[a + b + 1] += 1

        pairs = n // 2

        # Initially every pair needs 2 moves
        current = pairs * 2

        answer = float('inf')

        # Prefix sum traversal
        for target_sum in range(2, 2 * limit + 1):

            current += diff[target_sum]

            answer = min(answer, current)

        return answer
```

### Go

```go
func minMoves(nums []int, limit int) int {
    
    n := len(nums)

    // Difference array
    diff := make([]int, 2*limit+2)

    // Process every pair
    for i := 0; i < n/2; i++ {

        a := nums[i]
        b := nums[n-1-i]

        // Keep a <= b
        if a > b {
            a, b = b, a
        }

        // Range where 1 move is enough
        diff[a+1] -= 1
        diff[b+limit+1] += 1

        // Exact sum where 0 moves are needed
        diff[a+b] -= 1
        diff[a+b+1] += 1
    }

    pairs := n / 2

    // Initially assume 2 moves for every pair
    current := pairs * 2

    answer := current

    // Prefix sum traversal
    for sum := 2; sum <= 2*limit; sum++ {

        current += diff[sum]

        if current < answer {
            answer = current
        }
    }

    return answer
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic stays exactly the same in all five languages.

Only syntax changes.

### Step 1 — Process mirrored pairs

I iterate from:

```text
0 → n/2
```

because each element is paired with:

```text
n - 1 - i
```

This guarantees that every pair is processed exactly once.

---

### Step 2 — Sort the pair logically

For every pair:

```text
(a, b)
```

I keep:

```text
a <= b
```

This makes range calculations easier and avoids unnecessary conditions later.

---

### Step 3 — Assume worst-case moves

Initially, I assume every target sum requires:

```text
2 moves
```

for every pair.

This gives a safe starting point.

Later, I reduce the move count where better cases exist.

---

### Step 4 — Handle one-move range

A pair can achieve every sum inside:

```text
[a + 1, b + limit]
```

with only one move.

Instead of updating all values one by one, I use a difference array.

This allows range updates in constant time.

That is the main optimization in this LeetCode hard problem.

---

### Step 5 — Handle exact zero-move sum

The exact sum:

```text
a + b
```

already exists.

So this target sum needs one fewer move again.

That changes:

```text
1 move → 0 moves
```

for that exact position.

---

### Step 6 — Build prefix sums

After all updates are stored in the difference array, I compute prefix sums.

This converts range updates into actual move counts for every target sum.

---

### Step 7 — Find minimum answer

While building prefix sums, I continuously track:

```text
minimum moves
```

across all possible sums.

That final minimum becomes the answer.

---

## Examples

### Example 1

Input:

```text
nums = [1,2,4,3]
limit = 4
```

Output:

```text
1
```

Explanation:

Pairs are:

```text
(1,3)
(2,4)
```

Changing `4 → 2` makes both pairs sum to `4`.

Only one move is needed.

---

### Example 2

Input:

```text
nums = [1,2,2,1]
limit = 2
```

Output:

```text
2
```

Explanation:

Pairs:

```text
(1,1)
(2,2)
```

To make all sums equal, two changes are required.

---

### Example 3

Input:

```text
nums = [1,2,1,2]
limit = 2
```

Output:

```text
0
```

Explanation:

All mirrored pairs already produce the same sum.

No moves are needed.

---

## How to Use / Run Locally

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

---

### Go

Run:

```bash
go run solution.go
```

---

## Notes & Optimizations

* A brute-force solution checking every sum against every pair would be too slow.
* The difference array reduces range updates from `O(limit)` to `O(1)`.
* Prefix sums help reconstruct the final move counts efficiently.
* This is one of the most important optimization patterns in competitive programming.
* Edge cases mostly involve:

  * very small arrays
  * already complementary arrays
  * arrays where every pair needs two changes

An alternative approach could use direct counting, but it would not be as efficient as the prefix sum optimization used here.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
