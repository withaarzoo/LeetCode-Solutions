# 3116. Kth Smallest Amount With Single Denomination Combination | LeetCode Solution

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

The problem gives an array called `coins`, where every value represents a coin denomination, and an integer `k`.

There are unlimited coins of every denomination. However, there is one important rule: different denominations cannot be mixed together.

That means if a denomination is `3`, I can make:

`3, 6, 9, 12, ...`

If another denomination is `5`, I can make:

`5, 10, 15, 20, ...`

The goal is to find the `k`th smallest distinct amount that can be made using any one denomination.

The main challenge is that `k` can be very large, so generating all possible amounts and sorting them is not efficient. The optimized solution uses binary search, inclusion-exclusion, GCD, and LCM to count valid amounts without generating them one by one.

## Constraints

| Constraint                   | Value                                       |
| ---------------------------- | ------------------------------------------- |
| Number of coin denominations | `1 <= coins.length <= 15`                   |
| Coin value                   | `1 <= coins[i] <= 25`                       |
| Value of `k`                 | `1 <= k <= 2 * 10^9`                        |
| Coin values                  | `coins` contains pairwise distinct integers |

These constraints are the main reason a direct simulation is not practical. With `k` as large as `2 * 10^9`, the solution needs to count valid amounts mathematically instead of generating them.

## Intuition

My first thought was to look at what each denomination can produce.

For a coin `c`, the possible amounts are simply its positive multiples:

`c, 2c, 3c, 4c, ...`

So the problem becomes finding the `k`th smallest positive number that is divisible by at least one denomination in `coins`.

The direct approach would be to generate multiples of every coin and merge them. But this quickly becomes impossible when `k` is very large.

The important observation is that I do not actually need to generate the amounts. If I choose any number `x`, I can ask:

How many valid amounts are less than or equal to `x`?

For one coin `c`, the answer is:

`floor(x / c)`

When multiple coins are involved, some numbers can be divisible by more than one denomination. This creates duplicate counting.

That is where the inclusion-exclusion principle helps. I can count multiples of each denomination, subtract overlaps, add intersections of three denominations, and continue this pattern for every subset.

Once I can calculate how many valid amounts exist up to `x`, I can use binary search to find the smallest `x` whose count is at least `k`.

That `x` is the answer.

## Approach

I use the following steps.

1. Sort the `coins` array so smaller denominations are processed first.

2. Remove redundant denominations.

   * If a previously kept coin divides the current coin, the current coin does not create any new amounts.
   * For example, if `2` is already present, `4` is redundant because every multiple of `4` is also a multiple of `2`.

3. Generate every non-empty subset of the remaining useful coins.

4. For each subset, calculate its LCM.

   * Every number divisible by all coins in that subset must be divisible by their LCM.

5. Use the inclusion-exclusion principle.

   * Odd-sized subsets are added.
   * Even-sized subsets are subtracted.

6. Create a counting function `count(x)`.

   * It returns how many distinct valid amounts are less than or equal to `x`.

7. Binary search for the smallest value `x` where:

   `count(x) >= k`

8. Return that value as the `k`th smallest amount.

The binary search upper bound can safely be:

`smallestCoin * k`

because the first `k` multiples of the smallest coin alone already give `k` valid amounts.

## Data Structures Used

### Sorted Array

I sort the input array to make redundant coin removal easier.

A smaller coin is checked before a larger one, so I can quickly determine whether the larger coin's multiples are already covered.

### Dynamic Array or List

I store only useful coin denominations after removing redundant ones.

This can significantly reduce the number of subsets that need to be processed.

### LCM Array

I precompute the LCM for every non-empty subset.

This avoids calculating the same LCM again during every binary search iteration.

### Sign Array

For each subset, I store whether it should be added or subtracted during inclusion-exclusion.

* Odd number of coins: add
* Even number of coins: subtract

No complex data structures such as heaps, trees, graphs, or hash maps are required.

## Operations & Behavior Summary

The algorithm works like this:

1. Sort all coin denominations.
2. Remove denominations whose multiples are already completely covered by a smaller denomination.
3. Generate every non-empty subset of the remaining coins.
4. Calculate the LCM for each subset.
5. Mark each subset with an inclusion-exclusion sign.
6. Start binary search between `1` and `smallestCoin * k`.
7. For each middle value:

   * Count multiples of every subset LCM.
   * Add counts for odd-sized subsets.
   * Subtract counts for even-sized subsets.
8. If the total count is at least `k`, search the left half.
9. Otherwise, search the right half.
10. When binary search finishes, the remaining value is the `k`th smallest valid amount.

## Complexity

| Complexity       | Value                           | Explanation                                                                                                                                                                                                                                 |
| ---------------- | ------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(m * 2^m + 2^m * log(k * c))` | `m` is the number of useful coin denominations after removing redundant coins, and `c` is the smallest coin value. The first part precomputes subset LCMs, while the second part comes from binary search and inclusion-exclusion counting. |
| Space Complexity | `O(2^m)`                        | I store the LCM and sign for every non-empty subset of the useful denominations.                                                                                                                                                            |

Since `coins.length <= 15`, the number of possible subsets remains manageable.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long findKthSmallest(vector<int>& coins, int k) {
        // Sort coins so smaller denominations are processed first.
        sort(coins.begin(), coins.end());

        // Store only useful coins because some denominations can be redundant.
        vector<long long> useful;

        for (int coin : coins) {
            bool redundant = false;

            // If an already kept coin divides this coin, all multiples of this
            // coin are already covered by that smaller coin.
            for (long long prev : useful) {
                if (coin % prev == 0) {
                    redundant = true;
                    break;
                }
            }

            // Keep the coin only if it adds new possible amounts.
            if (!redundant) {
                useful.push_back(coin);
            }
        }

        // The kth multiple of the smallest coin is always a valid upper bound.
        long long high = useful[0] * 1LL * k;
        long long low = 1;

        int m = useful.size();
        int totalMasks = 1 << m;

        // lcms[mask] stores the LCM of all coins selected by this mask.
        vector<long long> lcms(totalMasks, 1);

        // signs[mask] is +1 for odd subset sizes and -1 for even subset sizes.
        vector<int> signs(totalMasks, 1);

        // Precompute the LCM and inclusion-exclusion sign for every subset.
        for (int mask = 1; mask < totalMasks; ++mask) {
            long long currentLCM = 1;
            int bits = 0;

            for (int i = 0; i < m; ++i) {
                // Include useful[i] when its bit is set.
                if (mask & (1 << i)) {
                    long long g = std::gcd(currentLCM, useful[i]);

                    // Divide before multiplying to reduce the risk of overflow.
                    currentLCM = currentLCM / g;

                    // If the LCM becomes larger than high, its contribution
                    // will always be zero, so I cap it at high + 1.
                    if (currentLCM > high / useful[i]) {
                        currentLCM = high + 1;
                        break;
                    }

                    currentLCM *= useful[i];
                    ++bits;
                }
            }

            // Save the subset LCM for future count() calls.
            lcms[mask] = currentLCM;

            // Odd subsets are added and even subsets are subtracted.
            signs[mask] = (bits % 2 == 1) ? 1 : -1;
        }

        // Count how many valid amounts are less than or equal to x.
        auto count = [&](long long x) {
            long long result = 0;

            for (int mask = 1; mask < totalMasks; ++mask) {
                // An LCM larger than x cannot divide any number up to x.
                if (lcms[mask] <= x) {
                    result += signs[mask] * (x / lcms[mask]);
                }
            }

            return result;
        };

        // Find the smallest value that contains at least k valid amounts.
        while (low < high) {
            long long mid = low + (high - low) / 2;

            // If mid already contains k or more valid amounts,
            // the answer can be mid or something smaller.
            if (count(mid) >= k) {
                high = mid;
            } else {
                // Otherwise, I need to search to the right.
                low = mid + 1;
            }
        }

        // low is the kth smallest valid amount.
        return low;
    }
};
```

### Java

```java
import java.util.*;

class Solution {
    public long findKthSmallest(int[] coins, int k) {
        // Sort coins so smaller denominations are processed first.
        Arrays.sort(coins);

        // Store only denominations that add new valid multiples.
        List<Long> usefulList = new ArrayList<>();

        for (int coin : coins) {
            boolean redundant = false;

            // Check whether a smaller kept coin already covers this coin.
            for (long prev : usefulList) {
                if (coin % prev == 0) {
                    redundant = true;
                    break;
                }
            }

            // Keep only non-redundant coins.
            if (!redundant) {
                usefulList.add((long) coin);
            }
        }

        int m = usefulList.size();

        // Convert the list into an array for faster repeated access.
        long[] useful = new long[m];
        for (int i = 0; i < m; i++) {
            useful[i] = usefulList.get(i);
        }

        // The kth multiple of the smallest coin is always a valid upper bound.
        long low = 1;
        long high = useful[0] * k;

        int totalMasks = 1 << m;

        // Store the LCM for every subset.
        long[] lcms = new long[totalMasks];

        // Store +1 for odd subsets and -1 for even subsets.
        int[] signs = new int[totalMasks];

        // Precompute LCM values and inclusion-exclusion signs.
        for (int mask = 1; mask < totalMasks; mask++) {
            long currentLCM = 1;
            int bits = 0;

            for (int i = 0; i < m; i++) {
                // Process this coin only when its bit belongs to the subset.
                if ((mask & (1 << i)) != 0) {
                    long g = gcd(currentLCM, useful[i]);

                    // Divide first to keep multiplication smaller.
                    currentLCM /= g;

                    // Cap the LCM when it becomes too large to matter.
                    if (currentLCM > high / useful[i]) {
                        currentLCM = high + 1;
                        break;
                    }

                    currentLCM *= useful[i];
                    bits++;
                }
            }

            // Save the final LCM for this subset.
            lcms[mask] = currentLCM;

            // Apply inclusion-exclusion based on subset size.
            signs[mask] = (bits % 2 == 1) ? 1 : -1;
        }

        // Binary search for the smallest value with at least k valid amounts.
        while (low < high) {
            long mid = low + (high - low) / 2;
            long count = 0;

            // Count numbers up to mid using inclusion-exclusion.
            for (int mask = 1; mask < totalMasks; mask++) {
                // This subset contributes only when its LCM is not larger than mid.
                if (lcms[mask] <= mid) {
                    count += signs[mask] * (mid / lcms[mask]);
                }
            }

            // Move left when mid already contains at least k valid amounts.
            if (count >= k) {
                high = mid;
            } else {
                // Otherwise, move right.
                low = mid + 1;
            }
        }

        // low is the kth smallest valid amount.
        return low;
    }

    private long gcd(long a, long b) {
        // Use the Euclidean algorithm to calculate the greatest common divisor.
        while (b != 0) {
            long temp = a % b;
            a = b;
            b = temp;
        }

        return a;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} coins
 * @param {number} k
 * @return {number}
 */
var findKthSmallest = function(coins, k) {
    // Sort coins so I can remove redundant larger denominations.
    coins.sort((a, b) => a - b);

    // Keep only coins that produce some multiples not already covered.
    const useful = [];

    for (const coin of coins) {
        let redundant = false;

        // If a kept coin divides this coin, every multiple of this coin
        // is already included among the multiples of the smaller coin.
        for (const prev of useful) {
            if (coin % prev === 0) {
                redundant = true;
                break;
            }
        }

        // Keep the coin only when it is useful.
        if (!redundant) {
            useful.push(coin);
        }
    }

    const m = useful.length;

    // The kth multiple of the smallest coin gives a safe upper bound.
    let low = 1;
    let high = useful[0] * k;

    const totalMasks = 1 << m;

    // Store the LCM and sign for every non-empty subset.
    const lcms = new Array(totalMasks).fill(1);
    const signs = new Array(totalMasks).fill(1);

    // Helper function for the greatest common divisor.
    const gcd = (a, b) => {
        while (b !== 0) {
            [a, b] = [b, a % b];
        }
        return a;
    };

    // Precompute information for every subset.
    for (let mask = 1; mask < totalMasks; mask++) {
        let currentLCM = 1;
        let bits = 0;

        for (let i = 0; i < m; i++) {
            // Include useful[i] when the corresponding bit is set.
            if ((mask & (1 << i)) !== 0) {
                const g = gcd(currentLCM, useful[i]);

                // Divide before multiplying to calculate the LCM safely.
                currentLCM /= g;

                // Cap values that are larger than the maximum binary search value.
                if (currentLCM > Math.floor(high / useful[i])) {
                    currentLCM = high + 1;
                    break;
                }

                currentLCM *= useful[i];
                bits++;
            }
        }

        // Save the subset LCM.
        lcms[mask] = currentLCM;

        // Odd subsets are added and even subsets are subtracted.
        signs[mask] = bits % 2 === 1 ? 1 : -1;
    }

    // Count how many valid amounts are less than or equal to x.
    const count = (x) => {
        let result = 0;

        for (let mask = 1; mask < totalMasks; mask++) {
            // Ignore subsets whose LCM cannot divide any value up to x.
            if (lcms[mask] <= x) {
                result += signs[mask] * Math.floor(x / lcms[mask]);
            }
        }

        return result;
    };

    // Binary search for the smallest value whose count is at least k.
    while (low < high) {
        const mid = Math.floor(low + (high - low) / 2);

        // Search left if mid already reaches the kth valid amount.
        if (count(mid) >= k) {
            high = mid;
        } else {
            // Otherwise, search larger values.
            low = mid + 1;
        }
    }

    // low is the kth smallest valid amount.
    return low;
};
```

### Python3

```python
from typing import List
from math import gcd

class Solution:

    def findKthSmallest(self, coins: List[int], k: int) -> int:
        # Sort coins so smaller denominations are checked first.
        coins.sort()

        # Keep only denominations that are not already covered.
        useful = []

        for coin in coins:
            redundant = False

            # If a smaller kept coin divides this coin, every multiple of
            # this coin is already produced by the smaller denomination.
            for prev in useful:
                if coin % prev == 0:
                    redundant = True
                    break

            # Keep the denomination only if it adds new multiples.
            if not redundant:
                useful.append(coin)

        # The kth multiple of the smallest coin gives a safe upper bound.
        low = 1
        high = useful[0] * k

        m = len(useful)
        total_masks = 1 << m

        # Store the LCM for every subset.
        lcms = [1] * total_masks

        # Store +1 for odd subsets and -1 for even subsets.
        signs = [1] * total_masks

        # Precompute LCM values and inclusion-exclusion signs.
        for mask in range(1, total_masks):
            current_lcm = 1
            bits = 0

            for i in range(m):
                # Include useful[i] when its bit is present in the subset.
                if mask & (1 << i):
                    # Divide first so the multiplication stays smaller.
                    current_lcm //= gcd(current_lcm, useful[i])

                    # If the LCM becomes larger than high, it can never
                    # contribute to any count during the binary search.
                    if current_lcm > high // useful[i]:
                        current_lcm = high + 1
                        break

                    current_lcm *= useful[i]
                    bits += 1

            # Save the final LCM for this subset.
            lcms[mask] = current_lcm

            # Inclusion-exclusion adds odd subsets and subtracts even ones.
            signs[mask] = 1 if bits % 2 == 1 else -1

        # Count valid amounts from 1 through x.
        def count(x: int) -> int:
            result = 0

            for mask in range(1, total_masks):
                # A larger LCM cannot divide any number up to x.
                if lcms[mask] <= x:
                    result += signs[mask] * (x // lcms[mask])

            return result

        # Find the smallest value containing at least k valid amounts.
        while low < high:
            mid = low + (high - low) // 2

            # Keep mid when it already reaches the kth position.
            if count(mid) >= k:
                high = mid
            else:
                # Otherwise, move to larger values.
                low = mid + 1

        # low is the kth smallest valid amount.
        return low
```

### Go

```go
func findKthSmallest(coins []int, k int) int64 {
 // Sort coins so smaller denominations are processed first.
 sort.Ints(coins)

 // Keep only denominations that are not already covered.
 useful := make([]int64, 0)

 for _, coin := range coins {
  redundant := false

  // Check whether a smaller kept coin already covers all multiples
  // of the current coin.
  for _, prev := range useful {
   if int64(coin)%prev == 0 {
    redundant = true
    break
   }
  }

  // Keep only coins that add new possible amounts.
  if !redundant {
   useful = append(useful, int64(coin))
  }
 }

 m := len(useful)

 // The kth multiple of the smallest coin is always a valid upper bound.
 low := int64(1)
 high := useful[0] * int64(k)

 totalMasks := 1 << m

 // lcms stores the LCM of every non-empty subset.
 lcms := make([]int64, totalMasks)

 // signs stores +1 for odd subsets and -1 for even subsets.
 signs := make([]int64, totalMasks)

 // Precompute subset LCMs and inclusion-exclusion signs.
 for mask := 1; mask < totalMasks; mask++ {
  currentLCM := int64(1)
  bits := 0

  for i := 0; i < m; i++ {
   // Include useful[i] when its bit is set in this subset.
   if mask&(1<<i) != 0 {
    g := gcd(currentLCM, useful[i])

    // Divide before multiplying to calculate the LCM safely.
    currentLCM /= g

    // Cap the LCM if it becomes too large to contribute.
    if currentLCM > high/useful[i] {
     currentLCM = high + 1
     break
    }

    currentLCM *= useful[i]
    bits++
   }
  }

  // Save the LCM for this subset.
  lcms[mask] = currentLCM

  // Odd subsets are added and even subsets are subtracted.
  if bits%2 == 1 {
   signs[mask] = 1
  } else {
   signs[mask] = -1
  }
 }

 // Count how many valid amounts are less than or equal to x.
 count := func(x int64) int64 {
  result := int64(0)

  for mask := 1; mask < totalMasks; mask++ {
   // Ignore subsets whose LCM is larger than x.
   if lcms[mask] <= x {
    result += signs[mask] * (x / lcms[mask])
   }
  }

  return result
 }

 // Binary search for the kth smallest valid amount.
 for low < high {
  mid := low + (high-low)/2

  // Search left when mid already contains at least k valid amounts.
  if count(mid) >= int64(k) {
   high = mid
  } else {
   // Otherwise, search larger values.
   low = mid + 1
  }
 }

 // low is the kth smallest valid amount.
 return low
}

// gcd returns the greatest common divisor using the Euclidean algorithm.
func gcd(a, b int64) int64 {
 for b != 0 {
  a, b = b, a%b
 }
 return a
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core algorithm is the same in C++, Java, JavaScript, Python3, and Go. The only differences are syntax, array handling, sorting methods, and GCD implementations.

### Step 1: Sort the Coin Denominations

I sort the denominations first.

This helps me process smaller values before larger values. It is useful because if a smaller denomination divides a larger one, the larger denomination is redundant.

For example:

`[2, 4, 6, 8]`

After keeping `2`, I can remove `4` and `8`.

Every multiple of `4` is already a multiple of `2`.

Every multiple of `8` is also already a multiple of `2`.

However, `6` still needs to remain because not every multiple of `6` is a multiple of `2` in the same way for inclusion-exclusion coverage? More precisely, multiples of `6` are covered by `2`, so if `2` is present, `6` is also redundant. Therefore, after processing this example, only `2` would remain.

Removing redundant coins is an optimization. The inclusion-exclusion logic can still work without this step, but fewer useful coins means fewer subsets and faster counting.

### Step 2: Build the Useful Coin List

For every sorted coin, I check the previously kept coins.

If any kept coin divides the current coin, I skip it.

The check is based on:

`currentCoin % previousCoin == 0`

This means every multiple of the current coin is already included in the multiples of the previous coin.

Otherwise, I keep the current denomination.

This step reduces unnecessary subset calculations before the main inclusion-exclusion process begins.

### Step 3: Set the Binary Search Range

The lower bound is `1`.

For the upper bound, I use:

`smallestCoin * k`

This is guaranteed to contain the answer.

For example, if the smallest coin is `3`, its first `k` multiples are:

`3, 6, 9, ..., 3 * k`

So by the time I reach `3 * k`, there are definitely at least `k` valid amounts.

Other denominations may create additional valid amounts between these multiples, which can only make the actual answer smaller.

### Step 4: Generate Every Non-empty Subset

If there are `m` useful coins, there are:

`2^m - 1`

non-empty subsets.

I represent every subset with a bitmask.

For three coins:

* `001` means the first coin
* `010` means the second coin
* `011` means the first and second coins
* `111` means all three coins

This is why bit manipulation is useful in this LeetCode solution.

Each subset represents one intersection used by the inclusion-exclusion principle.

### Step 5: Calculate the LCM of Each Subset

For every subset, I calculate the LCM of all selected coins.

The formula is:

`LCM(a, b) = a / GCD(a, b) * b`

I divide before multiplying.

This reduces the chance of overflow because the intermediate value stays smaller.

If the LCM becomes larger than the binary search upper bound, I do not need its exact value anymore.

Any LCM larger than every possible `x` will always contribute:

`floor(x / LCM) = 0`

So I can safely cap it at a value larger than the search range.

This is an important optimization for keeping the implementation safe with large values.

### Step 6: Apply Inclusion-Exclusion

Suppose the useful coins are:

`[2, 3]`

Up to `10`:

* Multiples of `2`: `5`
* Multiples of `3`: `3`

Adding them gives `8`, but `6` gets counted twice.

The overlap is divisible by:

`LCM(2, 3) = 6`

There is `1` multiple of `6` up to `10`.

So the correct count is:

`5 + 3 - 1 = 7`

This pattern continues for larger subsets.

* One selected coin: add
* Two selected coins: subtract
* Three selected coins: add
* Four selected coins: subtract

That is the inclusion-exclusion principle used by this optimized solution.

### Step 7: Count Valid Amounts Up to x

For a subset with LCM `lcm`, the number of values from `1` through `x` divisible by that LCM is:

`floor(x / lcm)`

I multiply this value by the subset's sign.

Then I add everything together.

The result is the number of distinct amounts less than or equal to `x` that can be created using exactly one denomination type.

This count function is monotonic.

As `x` increases, the number of valid amounts never decreases.

That monotonic behavior is exactly what makes binary search possible.

### Step 8: Binary Search for the Kth Amount

For each middle value `mid`, I calculate:

`count(mid)`

If:

`count(mid) >= k`

then the `k`th valid amount is at `mid` or somewhere before it.

So I move the right boundary:

`high = mid`

Otherwise, there are not enough valid amounts yet.

So I move the left boundary:

`low = mid + 1`

When the search ends, `low` and `high` point to the same smallest value where at least `k` valid amounts exist.

That value is exactly the `k`th smallest amount.

### Language-specific Implementation Notes

#### C++

I use `long long` because the answer and intermediate calculations can be larger than a regular `int`.

`std::gcd` can be used to calculate the GCD efficiently.

Vectors are used for the useful coins and precomputed subset information.

#### Java

I use `long` for LCM values, binary search bounds, and counting.

Java does not provide the same direct built-in GCD utility for this use, so I implement the Euclidean algorithm manually.

Arrays and `ArrayList` are enough for the required data.

#### JavaScript

JavaScript numbers are safe for integers up to `Number.MAX_SAFE_INTEGER`, which is enough for this problem's required answer range.

I use `Math.floor()` for integer division behavior.

The subset count is small because there are at most `15` original denominations.

#### Python3

Python integers automatically support arbitrary precision, so integer overflow is less of a concern.

I use `math.gcd()` for the GCD calculation.

Integer division is performed using `//`.

#### Go

I use `int64` for binary search values, LCM values, and counting.

The GCD function is implemented using the Euclidean algorithm.

Slices are used for storing useful coins and precomputed subset information.

## Examples

### Example 1

**Input:**

`coins = [3, 6, 9], k = 3`

**Expected Output:**

`9`

**Trace:**

The possible amounts are:

`3, 6, 9, 12, 15, ...`

Since `6` and `9` are both multiples of `3`, they are redundant denominations.

The useful set becomes:

`[3]`

The sequence starts:

`3, 6, 9, ...`

The 3rd smallest amount is:

`9`

### Example 2

**Input:**

`coins = [5, 2], k = 7`

**Expected Output:**

`12`

**Trace:**

Multiples of `2`:

`2, 4, 6, 8, 10, 12, ...`

Multiples of `5`:

`5, 10, 15, ...`

After removing duplicates and sorting all valid amounts:

`2, 4, 5, 6, 8, 10, 12, 14, 15, ...`

The 7th smallest amount is:

`12`

### Example 3

**Input:**

`coins = [3, 5, 7], k = 8`

**Expected Output:**

`14`

**Trace:**

The valid amounts in sorted order begin with:

`3, 5, 6, 7, 9, 10, 12, 14, ...`

The 8th smallest valid amount is:

`14`

The algorithm does not generate this entire sequence. Instead, it uses inclusion-exclusion to count how many valid amounts exist below each binary search value.

## How to Use / Run Locally

Create separate files for each language version and paste the corresponding solution code into them.

### C++

Save the code in a file such as:

`solution.cpp`

Compile it with:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it with:

```bash
./solution
```

On Windows, the executable may be:

```bash
solution.exe
```

### Java

Save the code in:

`Solution.java`

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

### JavaScript

Save the code in:

`solution.js`

Run it using Node.js:

```bash
node solution.js
```

### Python3

Save the code in:

`solution.py`

Run it with:

```bash
python3 solution.py
```

Depending on your system, this may also work:

```bash
python solution.py
```

### Go

Save the code in:

`solution.go`

Run it with:

```bash
go run solution.go
```

For LeetCode submissions, only the required `Solution` class or function needs to be submitted. A separate input/output driver is not required.

## Notes & Optimizations

* The biggest optimization is using binary search instead of generating valid amounts one by one.
* Inclusion-exclusion removes duplicate counting between different coin multiples.
* Removing redundant denominations can reduce the number of subsets significantly.
* Precomputing subset LCM values avoids repeating the same work during every binary search iteration.
* LCM values are calculated carefully to avoid unnecessary overflow.
* Any LCM larger than the binary search upper bound can be ignored because it contributes zero to every count.
* A priority queue approach could generate the sequence in increasing order, but it would not work efficiently when `k` is as large as `2 * 10^9`.
* The solution depends on the fact that there are at most `15` coin denominations, making subset enumeration practical.
* The binary search works because the number of valid amounts less than or equal to `x` never decreases as `x` increases.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
