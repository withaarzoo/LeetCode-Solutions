# Maximum Number of Jumps to Reach the Last Index

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

This problem asks us to find the maximum number of jumps needed to reach the last index of an array.

We start from index `0`, and from any index `i`, we are allowed to jump to another index `j` only if:

* `i < j`
* The difference between the values stays within the target range

The condition looks like this:

```text
-target <= nums[j] - nums[i] <= target
```

The goal is not to reach the end using the fewest jumps.
Instead, we need the largest possible number of valid jumps.

If reaching the last index is impossible, we return `-1`.

This is a classic Dynamic Programming problem because every position depends on previously reachable positions.

Popular SEO keywords naturally related to this problem include:

* LeetCode 2770 solution
* Maximum Number of Jumps to Reach the Last Index
* Dynamic Programming array problem
* DP jump game problem
* Maximum jumps DP solution
* Competitive programming DP tutorial
* Array traversal optimization

---

## Constraints

| Constraint                 | Value                   |
| -------------------------- | ----------------------- |
| `2 <= nums.length <= 1000` | Array size              |
| `-10^9 <= nums[i] <= 10^9` | Array values            |
| `0 <= target <= 2 * 10^9`  | Allowed jump difference |

---

## Intuition

The first thing I noticed was that every index can connect to multiple future indices.

That immediately made me think of a graph-like structure where:

* each index acts like a node
* every valid jump acts like an edge

Since the problem asks for the maximum jumps, I realized I should keep track of the best answer possible for every index.

So I defined:

```text
dp[i] = maximum jumps needed to reach index i
```

If an index cannot be reached, I store `-1`.

From there, the solution becomes straightforward:

* try all possible jumps
* check whether the jump is valid
* update the best answer for the destination index

---

## Approach

I used Dynamic Programming with two nested loops.

Step-by-step strategy:

1. Create a DP array of size `n`
2. Fill every value with `-1`
3. Set `dp[0] = 0` because the starting position needs zero jumps
4. Traverse every index `i`
5. From each index, try jumping to every future index `j`
6. Check whether the jump condition is valid
7. If valid, update the maximum jumps for `j`
8. Return the value stored at the last index

The solution works because every future state depends on already computed states.

---

## Data Structures Used

| Data Structure        | Purpose                                        |
| --------------------- | ---------------------------------------------- |
| Array / Vector / List | Stores the DP values                           |
| Integer Variables     | Used for indices, differences, and jump counts |

I only needed one DP array because the problem does not require storing actual paths.

---

## Operations & Behavior Summary

The algorithm works in these stages:

1. Initialize all indices as unreachable
2. Mark index `0` as reachable
3. Visit every index one by one
4. Try every possible forward jump
5. Check whether the difference condition is satisfied
6. Update the destination index with the best jump count
7. Continue until all possibilities are checked
8. Return the final DP value

This is essentially a bottom-up Dynamic Programming solution.

---

## Complexity

| Type             | Complexity | Explanation                         |
| ---------------- | ---------- | ----------------------------------- |
| Time Complexity  | `O(n²)`    | Every pair `(i, j)` is checked once |
| Space Complexity | `O(n)`     | One DP array is used                |

Where:

* `n` = length of the `nums` array

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maximumJumps(vector<int>& nums, int target) {
        
        int n = nums.size();
        
        // dp[i] = maximum jumps needed to reach index i
        vector<int> dp(n, -1);
        
        // Starting index needs 0 jumps
        dp[0] = 0;

        // Try every starting index
        for (int i = 0; i < n; i++) {
            
            // If current index is unreachable, skip it
            if (dp[i] == -1) continue;

            // Try jumping to every next index
            for (int j = i + 1; j < n; j++) {
                
                // Difference between values
                long long diff = 1LL * nums[j] - nums[i];

                // Check whether jump is valid
                if (diff >= -target && diff <= target) {
                    
                    // Update maximum jumps for index j
                    dp[j] = max(dp[j], dp[i] + 1);
                }
            }
        }

        // Final answer
        return dp[n - 1];
    }
};
```

### Java

```java
class Solution {
    public int maximumJumps(int[] nums, int target) {
        
        int n = nums.length;

        // dp[i] = maximum jumps needed to reach index i
        int[] dp = new int[n];

        // Fill with -1 meaning unreachable
        Arrays.fill(dp, -1);

        // Starting index
        dp[0] = 0;

        // Try every current index
        for (int i = 0; i < n; i++) {

            // Skip unreachable indices
            if (dp[i] == -1) continue;

            // Try every next index
            for (int j = i + 1; j < n; j++) {

                // Calculate difference
                long diff = (long) nums[j] - nums[i];

                // Check valid jump
                if (diff >= -target && diff <= target) {

                    // Update maximum jumps
                    dp[j] = Math.max(dp[j], dp[i] + 1);
                }
            }
        }

        // Return answer
        return dp[n - 1];
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var maximumJumps = function(nums, target) {
    
    const n = nums.length;

    // dp[i] = maximum jumps needed to reach index i
    const dp = new Array(n).fill(-1);

    // Starting position
    dp[0] = 0;

    // Try every current index
    for (let i = 0; i < n; i++) {

        // Skip unreachable indices
        if (dp[i] === -1) continue;

        // Try every next index
        for (let j = i + 1; j < n; j++) {

            // Calculate difference
            const diff = nums[j] - nums[i];

            // Check valid jump
            if (diff >= -target && diff <= target) {

                // Update maximum jumps
                dp[j] = Math.max(dp[j], dp[i] + 1);
            }
        }
    }

    // Final answer
    return dp[n - 1];
};
```

### Python3

```python
class Solution:
    def maximumJumps(self, nums: List[int], target: int) -> int:
        
        n = len(nums)

        # dp[i] = maximum jumps needed to reach index i
        dp = [-1] * n

        # Starting index requires 0 jumps
        dp[0] = 0

        # Try every current index
        for i in range(n):

            # Skip unreachable indices
            if dp[i] == -1:
                continue

            # Try every next index
            for j in range(i + 1, n):

                # Difference between values
                diff = nums[j] - nums[i]

                # Check whether jump is valid
                if -target <= diff <= target:

                    # Update maximum jumps
                    dp[j] = max(dp[j], dp[i] + 1)

        # Final answer
        return dp[n - 1]
```

### Go

```go
func maximumJumps(nums []int, target int) int {
    
    n := len(nums)

    // dp[i] = maximum jumps needed to reach index i
    dp := make([]int, n)

    // Initialize all values as -1
    for i := 0; i < n; i++ {
        dp[i] = -1
    }

    // Starting index
    dp[0] = 0

    // Try every current index
    for i := 0; i < n; i++ {

        // Skip unreachable indices
        if dp[i] == -1 {
            continue
        }

        // Try every next index
        for j := i + 1; j < n; j++ {

            // Difference between values
            diff := nums[j] - nums[i]

            // Check valid jump
            if diff >= -target && diff <= target {

                // Update maximum jumps
                if dp[i]+1 > dp[j] {
                    dp[j] = dp[i] + 1
                }
            }
        }
    }

    // Final answer
    return dp[n-1]
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic stays exactly the same in all five languages.

The only differences are syntax and how arrays are initialized.

### Step 1 — Create the DP Array

I create a DP array where:

```text
dp[i] = maximum jumps needed to reach index i
```

Initially, every value is set to `-1`.

That means the index is currently unreachable.

---

### Step 2 — Mark the Starting Position

The first index is already reachable because we start there.

So:

```text
dp[0] = 0
```

This means zero jumps are needed to stand at index `0`.

---

### Step 3 — Traverse Every Index

Now I visit every index one by one.

For each index:

* if it is unreachable, skip it
* otherwise try all future jumps

This prevents unnecessary work.

---

### Step 4 — Try All Future Jumps

For every index `i`, I check all indices `j > i`.

This follows the rule that jumps can only move forward.

---

### Step 5 — Validate the Jump

For every possible jump:

```text
nums[j] - nums[i]
```

must stay inside:

```text
[-target, target]
```

If the condition fails, that jump is ignored.

---

### Step 6 — Update the Best Answer

If the jump works:

```text
dp[j] = max(dp[j], dp[i] + 1)
```

Why `+1`?

Because we are taking one extra jump from `i` to `j`.

Why `max`?

Because the problem asks for the maximum number of jumps.

---

### Step 7 — Return the Last Value

At the end:

* if the last index is reachable, return its DP value
* otherwise return `-1`

---

### Language-specific Notes

#### C++

* Usually uses `vector<int>`
* `long long` is safer for subtraction because values can be large

#### Java

* Uses `Arrays.fill()` for initialization
* `long` helps avoid overflow during subtraction

#### JavaScript

* Arrays are dynamic
* Number type handles large integers safely for this problem

#### Python3

* Python integers automatically handle large values
* The cleanest implementation among all versions

#### Go

* Slices are used instead of vectors
* Manual initialization is needed for `-1`

---

## Examples

### Example 1

Input:

```text
nums = [1,3,6,4,1,2]
target = 2
```

Output:

```text
3
```

Explanation:

Possible path:

```text
0 -> 1 -> 3 -> 5
```

Total jumps = `3`

---

### Example 2

Input:

```text
nums = [1,3,6,4,1,2]
target = 3
```

Output:

```text
5
```

Explanation:

Possible path:

```text
0 -> 1 -> 2 -> 3 -> 4 -> 5
```

Every jump satisfies the condition.

---

### Example 3

Input:

```text
nums = [1,3,6,4,1,2]
target = 0
```

Output:

```text
-1
```

Explanation:

No valid jump sequence can reach the last index.

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

Run using Node.js:

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

Some important things I kept in mind:

* The problem size is only `1000`, so `O(n²)` works comfortably
* Using DP is much simpler than trying graph traversal approaches
* The solution only stores counts, not actual paths
* Overflow handling matters in C++ and Java because values can reach `10^9`

Possible alternative approaches:

* DFS + Memoization
* Graph traversal interpretation

But the bottom-up DP solution is cleaner and easier to debug.

Edge cases worth testing:

* No valid jumps
* Single valid path
* Multiple possible paths
* Very large positive and negative values
* Target equal to `0`

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
