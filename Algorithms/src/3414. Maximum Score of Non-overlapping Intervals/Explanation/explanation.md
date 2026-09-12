# 3414. Maximum Score of Non-overlapping Intervals

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
  * [TypeScript](#typescript)
  * [Python3](#python3)
  * [Go](#go)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

LeetCode 3414, **Maximum Score of Non-overlapping Intervals**, gives an array of intervals.

Each interval looks like:

```text
[left, right, weight]
```

The `left` value tells where the interval starts, `right` tells where it ends, and `weight` is the score I get by choosing that interval.

I can choose at most 4 intervals, but the selected intervals must not overlap.

Two intervals are considered overlapping even when one ends exactly where another starts. So for two intervals to be valid together, I need:

```text
previousRight < currentLeft
```

The goal is to return the original indices of the selected intervals that give the **maximum total weight**.

There is one more condition. If multiple selections have the same maximum score, I must return the **lexicographically smallest array of indices**.

For example:

```text
intervals = [[1,3,2], [4,5,2], [1,5,5], [6,9,3], [6,7,1], [8,9,1]]
```

The best choice is:

```text
[2, 3]
```

because interval `2` gives weight `5` and interval `3` gives weight `3`, so the total score is `8`.

## Constraints

| Constraint                 | Value                               |
| -------------------------- | ----------------------------------- |
| Number of intervals        | `1 <= intervals.length <= 5 * 10^4` |
| Elements in each interval  | `intervals[i].length == 3`          |
| Interval format            | `intervals[i] = [li, ri, weighti]`  |
| Left endpoint              | `1 <= li`                           |
| Right endpoint             | `li <= ri <= 10^9`                  |
| Weight                     | `1 <= weighti <= 10^9`              |
| Maximum selected intervals | `4`                                 |

## Intuition

My first thought was to try different combinations of intervals and keep the one with the highest score. But with up to `5 * 10^4` intervals, checking combinations is far too expensive.

The important observation is that I only need at most 4 intervals.

I first sort the intervals by their right endpoint. Once they are sorted, for every interval I can find the last interval that ends before its starting point.

For example:

```text
Previous intervals:

[1, 3]   [2, 5]   [4, 6]   [7, 9]
   |        |        |         |
   3        5        6         9

Current interval starts at 7

Need:
previousRight < 7

Compatible:
[1,3]
[2,5]
[4,6]

Not considered:
[7,9]
```

Because the intervals are sorted by right endpoint, I can find this compatible position using binary search.

After that, I use dynamic programming.

I keep track of how many intervals I have selected so far:

```text
0 intervals
1 interval
2 intervals
3 intervals
4 intervals
```

For every interval, I have two choices:

* Skip it.
* Take it and add its weight to the best compatible state.

This turns the problem into a weighted interval scheduling problem with a small DP dimension.

The lexicographical condition is handled at the same time. When two choices have the same score, I compare their original index arrays and keep the smaller one.

## Approach

I use the following steps.

### 1. Store the original index

Before sorting, I attach the original index to every interval:

```text
[left, right, weight, originalIndex]
```

This is necessary because the answer must contain the indices from the original input order.

### 2. Sort by right endpoint

I sort all intervals using their right endpoint.

After sorting, compatible intervals always appear before the current interval.

### 3. Find the previous compatible interval

For the current interval `[left, right, weight]`, I need an earlier interval whose right endpoint is strictly smaller than `left`.

So I find the first endpoint satisfying:

```text
right >= left
```

using binary search.

Everything before that position satisfies:

```text
right < left
```

and is therefore compatible.

### 4. Build the dynamic programming table

I define:

```text
dp[k][i]
```

as the best result when I choose exactly `k` intervals from the first `i` sorted intervals.

Since the problem allows at most 4 intervals, I only need:

```text
k = 0, 1, 2, 3, 4
```

For every interval, I check two possibilities.

If I skip it:

```text
dp[k][i] = dp[k][i - 1]
```

If I take it:

```text
dp[k][i] =
    best compatible result with k - 1 intervals
    + current weight
```

The compatible result comes from the prefix found through binary search.

### 5. Handle lexicographical order

If two candidates have the same score, I compare their original index arrays.

For example:

```text
Candidate A: [1, 4]
Candidate B: [2, 3]

Same score
```

I choose:

```text
[1, 4]
```

because `1 < 2`.

### 6. Check all possible final sizes

The answer may contain 1, 2, 3, or 4 intervals.

So I compare:

```text
dp[1][n]
dp[2][n]
dp[3][n]
dp[4][n]
```

and return the best one.

## Data Structures Used

| Data Structure    | Purpose                                                  |
| ----------------- | -------------------------------------------------------- |
| Array / Vector    | Store intervals along with their original indices        |
| Sorted array      | Keep intervals ordered by right endpoint                 |
| DP table          | Store the best result for selecting `0` to `4` intervals |
| Small index array | Store the original indices selected by a DP state        |
| Binary search     | Find the last compatible interval efficiently            |

The selected index list contains at most 4 values, so comparing and sorting it is very small and does not become a bottleneck.

## Operations & Behavior Summary

The overall logic can be summarized like this:

```text
Read intervals
      |
      v
Attach original indices
      |
      v
Sort by right endpoint
      |
      v
For every interval
      |
      +----> Binary search for first right >= current left
      |
      +----> Skip current interval
      |
      +----> Take current interval
                  |
                  v
          Use best compatible
          state with k - 1 intervals
      |
      v
Compare score
      |
      +----> Higher score wins
      |
      +----> Same score:
             lexicographically smaller indices win
      |
      v
Check answers for 1..4 intervals
      |
      v
Return best original indices
```

The most important condition is:

```text
previousRight < currentLeft
```

Using `<=` would be incorrect because intervals touching at the boundary are considered overlapping in this problem.

## Complexity

| Type  | Complexity   | Explanation                                                                                      |
| ----- | ------------ | ------------------------------------------------------------------------------------------------ |
| Time  | `O(n log n)` | Sorting takes `O(n log n)` and every interval performs binary search for each of the 4 DP levels |
| Space | `O(n)`       | The DP stores `5 * (n + 1)` states, and every selected index list contains at most 4 indices     |

Here, `n` is the number of intervals.

Since the number of selected intervals is fixed at `4`, the factor of `4` is treated as a constant:

```text
O(4 * n log n) = O(n log n)
```

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    // Each DP state stores the maximum score and the sorted original indices.
    struct Node {
        long long score;      // Total weight of the selected intervals.
        vector<int> ids;      // Original indices, kept in sorted order.
        bool valid;           // Tells me whether this state can be formed.

        Node() : score(0), valid(false) {}
        Node(long long s, vector<int> v) : score(s), ids(std::move(v)), valid(true) {}
    };

    // Returns true when 'a' is better than 'b'.
    bool better(const Node& a, const Node& b) {
        // A valid candidate is always better than an invalid state.
        if (!a.valid) return false;
        if (!b.valid) return true;

        // A larger score is the primary requirement.
        if (a.score != b.score) return a.score > b.score;

        // If scores tie, I need the lexicographically smaller index array.
        return a.ids < b.ids;
    }

    vector<int> maximumWeight(vector<vector<int>>& intervals) {
        int n = intervals.size(); // Number of intervals.
        const int K = 4;          // I can choose at most four intervals.

        // I store [left, right, weight, originalIndex].
        vector<array<long long, 4>> a(n);

        // Attach each interval's original index before sorting.
        for (int i = 0; i < n; ++i) {
            a[i] = {intervals[i][0], intervals[i][1], intervals[i][2], i};
        }

        // Sorting by right endpoint makes compatible intervals form a prefix.
        sort(a.begin(), a.end(), [](const auto& x, const auto& y) {
            return x[1] < y[1];
        });

        // Store all right endpoints so binary search becomes easy.
        vector<long long> ends(n);
        for (int i = 0; i < n; ++i) {
            ends[i] = a[i][1];
        }

        // dp[k][i] = best result using exactly k intervals
        // from the first i sorted intervals.
        vector<vector<Node>> dp(K + 1, vector<Node>(n + 1));

        // Choosing zero intervals always gives score 0 and an empty index list.
        for (int i = 0; i <= n; ++i) {
            dp[0][i] = Node(0, {});
        }

        // Process the sorted intervals one by one.
        for (int i = 1; i <= n; ++i) {
            long long l = a[i - 1][0]; // Current interval's left endpoint.
            long long w = a[i - 1][2]; // Current interval's weight.
            int idx = (int)a[i - 1][3]; // Current interval's original index.

            // Find the first ending point >= l.
            // Every interval before it has end < l and is compatible.
            int p = lower_bound(ends.begin(), ends.begin() + (i - 1), l) - ends.begin();

            // For every possible number of selected intervals...
            for (int k = 1; k <= K; ++k) {
                // Option 1: skip the current interval.
                dp[k][i] = dp[k][i - 1];

                // Option 2 is possible only when dp[k - 1][p] exists.
                if (dp[k - 1][p].valid) {
                    // Copy the previous selected indices.
                    vector<int> ids = dp[k - 1][p].ids;

                    // Add the current original index.
                    ids.push_back(idx);

                    // Keep indices sorted because the final answer
                    // has to be compared lexicographically.
                    sort(ids.begin(), ids.end());

                    // Build the candidate obtained by taking this interval.
                    Node take(dp[k - 1][p].score + w, std::move(ids));

                    // Keep whichever of skip/take is better.
                    if (better(take, dp[k][i])) {
                        dp[k][i] = std::move(take);
                    }
                }
            }
        }

        // The answer can contain 1, 2, 3, or 4 intervals.
        Node ans;

        // Compare all valid final states.
        for (int k = 1; k <= K; ++k) {
            if (better(dp[k][n], ans)) {
                ans = dp[k][n];
            }
        }

        // Return the original indices of the best set.
        return ans.ids;
    }
};
```

### Java

```java
import java.util.*;

class Solution {

    // Stores the best score and the original indices of selected intervals.
    static class Node {
        long score;
        int[] ids;

        Node(long score, int[] ids) {
            this.score = score;
            this.ids = ids;
        }
    }

    // Returns true if candidate 'a' is better than candidate 'b'.
    private boolean better(Node a, Node b) {
        // If a does not exist, it cannot be better.
        if (a == null) {
            return false;
        }

        // If b does not exist, a is automatically better.
        if (b == null) {
            return true;
        }

        // Higher score is always better.
        if (a.score != b.score) {
            return a.score > b.score;
        }

        // Same score -> lexicographically smaller index array is better.
        int len = Math.min(a.ids.length, b.ids.length);

        for (int i = 0; i < len; i++) {
            if (a.ids[i] != b.ids[i]) {
                return a.ids[i] < b.ids[i];
            }
        }

        // If one array is a prefix of the other,
        // the shorter array is lexicographically smaller.
        return a.ids.length < b.ids.length;
    }

    // Inserts one index into an already sorted array.
    private int[] addSorted(int[] ids, int value) {
        // Create a new array with one extra position.
        int[] result = Arrays.copyOf(ids, ids.length + 1);

        // Add the new original index.
        result[ids.length] = value;

        // There are at most 4 indices, so sorting is effectively constant time.
        Arrays.sort(result);

        return result;
    }

    // Finds the first position where ends[pos] >= target.
    // All positions before it have end < target and are compatible.
    private int lowerBound(long[] ends, int length, long target) {
        int left = 0;
        int right = length;

        while (left < right) {
            int mid = left + (right - left) / 2;

            // Current interval ends too late or exactly at target,
            // so it overlaps and cannot be used.
            if (ends[mid] >= target) {
                right = mid;
            } else {
                // This interval is compatible, so search further right.
                left = mid + 1;
            }
        }

        // Number of previous intervals with end < target.
        return left;
    }

    public int[] maximumWeight(List<List<Integer>> intervals) {
        int n = intervals.size();

        // I can choose at most 4 intervals.
        final int K = 4;

        // Store:
        // [left, right, weight, originalIndex]
        long[][] arr = new long[n][4];

        // Save every interval together with its original index.
        for (int i = 0; i < n; i++) {
            arr[i][0] = intervals.get(i).get(0);
            arr[i][1] = intervals.get(i).get(1);
            arr[i][2] = intervals.get(i).get(2);
            arr[i][3] = i;
        }

        // Sort intervals by right endpoint.
        Arrays.sort(arr, (a, b) -> Long.compare(a[1], b[1]));

        // Store all right endpoints for binary search.
        long[] ends = new long[n];

        for (int i = 0; i < n; i++) {
            ends[i] = arr[i][1];
        }

        /*
         * dp[k][i] =
         * best result using exactly k intervals
         * from the first i sorted intervals.
         *
         * null means this state is currently impossible.
         */
        Node[][] dp = new Node[K + 1][n + 1];

        // Choosing zero intervals is always possible.
        for (int i = 0; i <= n; i++) {
            dp[0][i] = new Node(0, new int[0]);
        }

        // Process each interval.
        for (int i = 1; i <= n; i++) {
            long left = arr[i - 1][0];
            long weight = arr[i - 1][2];
            int originalIndex = (int) arr[i - 1][3];

            /*
             * Find the first previous interval whose end >= left.
             *
             * Therefore, all intervals before p satisfy:
             *
             * previousEnd < currentStart
             *
             * and are valid predecessors.
             */
            int p = lowerBound(ends, i - 1, left);

            // Try selecting exactly k intervals.
            for (int k = 1; k <= K; k++) {

                // Option 1: skip the current interval.
                dp[k][i] = dp[k][i - 1];

                /*
                 * Option 2: take the current interval.
                 *
                 * Then I need k - 1 compatible intervals before it.
                 */
                if (dp[k - 1][p] != null) {

                    // Add the current original index.
                    int[] ids = addSorted(
                        dp[k - 1][p].ids,
                        originalIndex
                    );

                    // Calculate the score after taking this interval.
                    Node take = new Node(
                        dp[k - 1][p].score + weight,
                        ids
                    );

                    // Keep the better of skip and take.
                    if (better(take, dp[k][i])) {
                        dp[k][i] = take;
                    }
                }
            }
        }

        // The final answer can use 1, 2, 3, or 4 intervals.
        Node answer = null;

        // Compare all possible numbers of selected intervals.
        for (int k = 1; k <= K; k++) {
            if (better(dp[k][n], answer)) {
                answer = dp[k][n];
            }
        }

        // Return the original indices.
        return answer.ids;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} intervals
 * @return {number[]}
 */
var maximumWeight = function(intervals) {
    const n = intervals.length; // Number of intervals.
    const K = 4;                // At most four intervals can be selected.

    // Add the original index because sorting would otherwise lose it.
    const a = intervals.map((x, i) => [x[0], x[1], x[2], i]);

    // Sort by right endpoint so compatible intervals become a prefix.
    a.sort((x, y) => x[1] - y[1]);

    // Store all right endpoints for binary search.
    const ends = a.map(x => x[1]);

    // dp[k][i] stores the best result using exactly k intervals
    // from the first i sorted intervals.
    const dp = Array.from({ length: K + 1 }, () => Array(n + 1).fill(null));

    // Zero intervals always give score 0 and an empty selection.
    for (let i = 0; i <= n; i++) {
        dp[0][i] = { score: 0, ids: [] };
    }

    // Compare two DP candidates.
    const better = (a, b) => {
        // A real candidate beats an empty/unreachable state.
        if (a === null) return false;
        if (b === null) return true;

        // Larger score is always better.
        if (a.score !== b.score) return a.score > b.score;

        // For equal scores, compare original indices lexicographically.
        const len = Math.min(a.ids.length, b.ids.length);

        for (let i = 0; i < len; i++) {
            if (a.ids[i] !== b.ids[i]) {
                return a.ids[i] < b.ids[i];
            }
        }

        // If one array is a prefix of the other, shorter is smaller.
        return a.ids.length < b.ids.length;
    };

    // Find the first end >= target among the first 'length' intervals.
    const lowerBound = (length, target) => {
        let lo = 0;
        let hi = length;

        // Binary search for the first incompatible endpoint.
        while (lo < hi) {
            const mid = lo + Math.floor((hi - lo) / 2);

            // End >= target cannot be used, so search left.
            if (ends[mid] >= target) {
                hi = mid;
            } else {
                // End < target is compatible, so search right.
                lo = mid + 1;
            }
        }

        // This is the number of compatible previous intervals.
        return lo;
    };

    // Process intervals in sorted order.
    for (let i = 1; i <= n; i++) {
        const [left, , weight, originalIndex] = a[i - 1];

        // Find how many earlier intervals are compatible.
        const p = lowerBound(i - 1, left);

        // Try every possible number of selected intervals.
        for (let k = 1; k <= K; k++) {
            // Option 1: skip the current interval.
            dp[k][i] = dp[k][i - 1];

            // Option 2: take it, if k - 1 intervals can be chosen before it.
            if (dp[k - 1][p] !== null) {
                // Copy the old indices because DP states must not be mutated.
                const ids = [...dp[k - 1][p].ids, originalIndex];

                // The final comparison must use sorted original indices.
                ids.sort((x, y) => x - y);

                // Build the candidate created by taking this interval.
                const take = {
                    score: dp[k - 1][p].score + weight,
                    ids
                };

                // Keep the better of skipping and taking.
                if (better(take, dp[k][i])) {
                    dp[k][i] = take;
                }
            }
        }
    }

    // The answer may contain 1 through 4 intervals.
    let answer = null;

    // Compare every possible count.
    for (let k = 1; k <= K; k++) {
        if (better(dp[k][n], answer)) {
            answer = dp[k][n];
        }
    }

    // Return only the original indices.
    return answer.ids;
};
```

### TypeScript

```typescript
function maximumWeight(intervals: number[][]): number[] {
    const n = intervals.length; // Number of intervals.
    const K = 4;                // At most four intervals may be selected.

    // Each item stores left, right, weight, and original index.
    const a: number[][] = intervals.map((x, i) => [x[0], x[1], x[2], i]);

    // Sort by right endpoint to make predecessor searching possible.
    a.sort((x, y) => x[1] - y[1]);

    // Save the right endpoints for binary search.
    const ends: number[] = a.map(x => x[1]);

    // A DP state contains the maximum score and its original indices.
    type State = {
        score: number;
        ids: number[];
    };

    // dp[k][i] = best result using exactly k intervals from first i intervals.
    const dp: Array<Array<State | null>> = Array.from(
        { length: K + 1 },
        () => Array<State | null>(n + 1).fill(null)
    );

    // Choosing zero intervals always gives score 0.
    for (let i = 0; i <= n; i++) {
        dp[0][i] = { score: 0, ids: [] };
    }

    // Checks whether candidate a is better than candidate b.
    const better = (a: State | null, b: State | null): boolean => {
        // A valid state beats an unreachable state.
        if (a === null) return false;
        if (b === null) return true;

        // Higher score has priority.
        if (a.score !== b.score) return a.score > b.score;

        // Equal scores are resolved lexicographically.
        const len = Math.min(a.ids.length, b.ids.length);

        for (let i = 0; i < len; i++) {
            if (a.ids[i] !== b.ids[i]) {
                return a.ids[i] < b.ids[i];
            }
        }

        // A shorter prefix is lexicographically smaller.
        return a.ids.length < b.ids.length;
    };

    // Finds the first right endpoint >= target.
    const lowerBound = (length: number, target: number): number => {
        let lo = 0;
        let hi = length;

        // Standard binary search for the first incompatible interval.
        while (lo < hi) {
            const mid = lo + Math.floor((hi - lo) / 2);

            // End >= target is not allowed because touching overlaps.
            if (ends[mid] >= target) {
                hi = mid;
            } else {
                // End < target is compatible.
                lo = mid + 1;
            }
        }

        // This position is also the count of compatible intervals.
        return lo;
    };

    // Process every interval after sorting.
    for (let i = 1; i <= n; i++) {
        const left = a[i - 1][0];   // Current interval's left endpoint.
        const weight = a[i - 1][2]; // Current interval's weight.
        const index = a[i - 1][3];  // Original index before sorting.

        // Find the compatible prefix before the current interval.
        const p = lowerBound(i - 1, left);

        // Try selecting exactly k intervals.
        for (let k = 1; k <= K; k++) {
            // Option 1: skip the current interval.
            dp[k][i] = dp[k][i - 1];

            // Option 2: take the current interval.
            const prev = dp[k - 1][p];

            if (prev !== null) {
                // Copy the old indices so other DP states stay unchanged.
                const ids = [...prev.ids, index];

                // Keep original indices sorted for lexicographical comparison.
                ids.sort((x, y) => x - y);

                // Build the candidate formed by taking this interval.
                const take: State = {
                    score: prev.score + weight,
                    ids
                };

                // Keep the better candidate.
                if (better(take, dp[k][i])) {
                    dp[k][i] = take;
                }
            }
        }
    }

    // The best result can contain from 1 to 4 intervals.
    let answer: State | null = null;

    // Compare every possible number of chosen intervals.
    for (let k = 1; k <= K; k++) {
        if (better(dp[k][n], answer)) {
            answer = dp[k][n];
        }
    }

    // 'answer' must exist because n >= 1.
    return answer!.ids;
}
```

### Python3

```python
from typing import List


class Solution:
    def maximumWeight(self, intervals: List[List[int]]) -> List[int]:
        n = len(intervals)  # Number of intervals.
        K = 4               # I can choose at most four intervals.

        # Add the original index because sorting changes the order.
        intervals = [
            [l, r, w, i]
            for i, (l, r, w) in enumerate(intervals)
        ]

        # Sort by right endpoint so compatible intervals form a prefix.
        intervals.sort(key=lambda x: x[1])

        # Store right endpoints for binary search.
        ends = [x[1] for x in intervals]

        # A DP state is [score, sorted original indices].
        # None means that the state cannot be formed.
        dp = [[None] * (n + 1) for _ in range(K + 1)]

        # Selecting zero intervals always gives score 0.
        for i in range(n + 1):
            dp[0][i] = (0, [])

        # Compares two states and returns the better one.
        def better(a, b):
            # A valid state beats an unreachable state.
            if a is None:
                return False
            if b is None:
                return True

            # Maximum score is the primary condition.
            if a[0] != b[0]:
                return a[0] > b[0]

            # Equal scores require the lexicographically smallest indices.
            return a[1] < b[1]

        # Finds the first endpoint >= target.
        def lower_bound(length, target):
            lo = 0
            hi = length

            # Standard binary search.
            while lo < hi:
                mid = lo + (hi - lo) // 2

                # end >= target is incompatible because boundaries overlap.
                if ends[mid] >= target:
                    hi = mid
                else:
                    # end < target is compatible.
                    lo = mid + 1

            # This is the number of compatible previous intervals.
            return lo

        # Process intervals in sorted order.
        for i in range(1, n + 1):
            l, _, w, original_index = intervals[i - 1]

            # Find the compatible prefix before the current interval.
            p = lower_bound(i - 1, l)

            # Try choosing exactly k intervals.
            for k in range(1, K + 1):
                # Option 1: skip the current interval.
                dp[k][i] = dp[k][i - 1]

                # Option 2: take the current interval.
                prev = dp[k - 1][p]

                if prev is not None:
                    # Copy previous indices so this DP state stays independent.
                    ids = prev[1] + [original_index]

                    # Sort indices because the answer is compared lexicographically.
                    ids.sort()

                    # Build the candidate produced by taking this interval.
                    take = (prev[0] + w, ids)

                    # Keep whichever option is better.
                    if better(take, dp[k][i]):
                        dp[k][i] = take

        # The answer can contain 1, 2, 3, or 4 intervals.
        answer = None

        # Compare every possible number of selected intervals.
        for k in range(1, K + 1):
            if better(dp[k][n], answer):
                answer = dp[k][n]

        # Return only the selected original indices.
        return answer[1]
```

### Go

```go
func maximumWeight(intervals [][]int) []int {
 type State struct {
  score  int64   // Total score of the selected intervals.
  ids    []int   // Original indices kept in sorted order.
  valid  bool    // True when this DP state is reachable.
 }

 n := len(intervals) // Number of intervals.
 const K = 4         // At most four intervals can be selected.

 // Store left, right, weight, and original index.
 a := make([][4]int64, n)

 // Save the original index before sorting.
 for i := 0; i < n; i++ {
  a[i] = [4]int64{
   int64(intervals[i][0]),
   int64(intervals[i][1]),
   int64(intervals[i][2]),
   int64(i),
  }
 }

 // Sort by right endpoint so binary search can find compatible prefixes.
 sort.Slice(a, func(i, j int) bool {
  return a[i][1] < a[j][1]
 })

 // Store every right endpoint for predecessor searches.
 ends := make([]int64, n)
 for i := 0; i < n; i++ {
  ends[i] = a[i][1]
 }

 // dp[k][i] is the best result using exactly k intervals
 // from the first i sorted intervals.
 dp := make([][]State, K+1)
 for k := 0; k <= K; k++ {
  dp[k] = make([]State, n+1)
 }

 // Choosing zero intervals is always possible with score zero.
 for i := 0; i <= n; i++ {
  dp[0][i] = State{
   score: 0,
   ids:   []int{},
   valid: true,
  }
 }

 // Returns true when a is better than b.
 better := func(a State, b State) bool {
  // Any valid state beats an invalid state.
  if !a.valid {
   return false
  }
  if !b.valid {
   return true
  }

  // Higher score always wins.
  if a.score != b.score {
   return a.score > b.score
  }

  // Equal scores are compared lexicographically.
  for i := 0; i < len(a.ids) && i < len(b.ids); i++ {
   if a.ids[i] != b.ids[i] {
    return a.ids[i] < b.ids[i]
   }
  }

  // If one is a prefix of the other, shorter is smaller.
  return len(a.ids) < len(b.ids)
 }

 // Finds the first endpoint >= target.
 // Everything before that position has endpoint < target.
 lowerBound := func(length int, target int64) int {
  lo := 0
  hi := length

  // Standard lower-bound binary search.
  for lo < hi {
   mid := lo + (hi-lo)/2

   // End >= target is incompatible because touching overlaps.
   if ends[mid] >= target {
    hi = mid
   } else {
    // End < target is compatible, so move right.
    lo = mid + 1
   }
  }

  // This position is the number of compatible previous intervals.
  return lo
 }

 // Process intervals in sorted order.
 for i := 1; i <= n; i++ {
  left := a[i-1][0]         // Current interval's left endpoint.
  weight := a[i-1][2]       // Current interval's weight.
  originalIndex := int(a[i-1][3]) // Original index before sorting.

  // Find the compatible prefix before the current interval.
  p := lowerBound(i-1, left)

  // Try selecting exactly k intervals.
  for k := 1; k <= K; k++ {
   // Option 1: skip the current interval.
   dp[k][i] = dp[k][i-1]

   // Option 2: take the current interval.
   prev := dp[k-1][p]

   if prev.valid {
    // Copy the previous indices so other states are not modified.
    ids := append([]int{}, prev.ids...)

    // Add the current original index.
    ids = append(ids, originalIndex)

    // Sort the tiny index list for lexicographical comparison.
    sort.Ints(ids)

    // Build the candidate created by taking this interval.
    take := State{
     score: prev.score + weight,
     ids:   ids,
     valid: true,
    }

    // Keep the better of skip and take.
    if better(take, dp[k][i]) {
     dp[k][i] = take
    }
   }
  }
 }

 // Compare all possible final states containing at most four intervals.
 var answer State

 // Check every possible number of selected intervals.
 for k := 1; k <= K; k++ {
  if better(dp[k][n], answer) {
   answer = dp[k][n]
  }
 }

 // Return the original indices of the best choice.
 return answer.ids
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core algorithm is the same in all six languages. The syntax changes, but the reasoning stays the same.

### Step 1: Keep the original index

Sorting the intervals changes their positions, but the final answer must contain the indices from the original input.

So every interval is internally treated as:

```text
[left, right, weight, originalIndex]
```

For example:

```text
Original:

Index 0 -> [5, 8, 1]
Index 1 -> [6, 7, 7]
Index 2 -> [4, 7, 3]

After sorting by right endpoint:

[6,7,7,1]
[4,7,3,2]
[5,8,1,0]
```

The last value still tells me where the interval came from.

### Step 2: Sort by right endpoint

I sort by the `right` value.

This is what makes binary search possible.

Suppose the right endpoints are:

```text
2  4  6  8  10
```

If the current interval starts at `7`, I want:

```text
right < 7
```

So the compatible endpoints are:

```text
2  4  6
```

and the first endpoint that is not compatible is:

```text
8
```

Binary search finds that boundary in `O(log n)` time.

### Step 3: Use strict non-overlap

This problem does not allow intervals that only touch.

For example:

```text
[1, 5]
     [5, 9]
```

These overlap at point `5`.

So the valid condition is:

```text
5 < 5
```

which is false.

That is why the binary search looks for the first:

```text
right >= left
```

instead of the first:

```text
right > left
```

This is a small detail, but it is one of the most important parts of the solution.

### Step 4: Define the DP state

I use:

```text
dp[k][i]
```

to mean:

> The best answer I can create by using exactly `k` intervals from the first `i` sorted intervals.

The DP has only five rows:

```text
k = 0
k = 1
k = 2
k = 3
k = 4
```

This is enough because selecting more than 4 intervals is not allowed.

### Step 5: Base case

Selecting zero intervals is always possible.

So:

```text
dp[0][i] = score 0
```

with an empty index list.

This lets the algorithm create a one-interval solution by adding the first chosen interval to the empty state.

### Step 6: Skip the current interval

For every interval, one possibility is to ignore it.

So:

```text
dp[k][i] = dp[k][i - 1]
```

This means the best result using the first `i` intervals may be exactly the same as the best result using the first `i - 1`.

### Step 7: Take the current interval

The second possibility is to choose the current interval.

Suppose the current interval is:

```text
[left, right, weight]
```

and binary search tells me that the first incompatible previous interval is at position `p`.

Then all intervals before `p` are compatible.

To choose exactly `k` intervals including the current one, I need:

```text
dp[k - 1][p]
```

Then:

```text
newScore = dp[k - 1][p].score + weight
```

The current interval is added to the selected index list.

### Step 8: Compare the two choices

Now I have:

```text
skip
take
```

I first compare their scores.

Higher score always wins.

For equal scores, I compare the original indices.

Example:

```text
skip -> score 15, indices [0, 4]
take -> score 15, indices [1, 2]
```

The first values differ:

```text
0 < 1
```

so `[0, 4]` is lexicographically smaller.

The DP keeps that candidate.

### Step 9: Why I store the actual indices

Another possible design is to store only the score and reconstruct the answer later using parent pointers.

That can work, but this problem has a special lexicographical tie-breaking rule.

Since I select at most 4 intervals, storing the selected original indices directly keeps the implementation easier to understand.

Each state stores only a tiny list:

```text
[]
[2]
[1,4]
[0,2,5]
[0,1,3,6]
```

The maximum length is always 4.

### Step 10: Final comparison

At the end, I do not have to choose exactly 4 intervals.

The problem says **at most 4**.

So I compare:

```text
1 interval
2 intervals
3 intervals
4 intervals
```

The candidate with the highest score is the answer.

If several have the same score, I again use lexicographical order.

### C++

C++ uses `vector` for the selected indices and `lower_bound` for binary search.

The STL makes the predecessor search straightforward:

```text
lower_bound(endpoints, currentLeft)
```

The C++ vector comparison can also handle lexicographical comparison naturally.

### Java

In Java, the main difference is that the DP states are objects.

An unreachable state can simply be represented by `null`.

This is important because trying to access a property such as `state.score` when the state is `null` causes a `NullPointerException`.

So the comparison helper must always check:

```text
a == null
b == null
```

before accessing their values.

Java's `Arrays.sort` and a custom binary search handle the interval processing.

### JavaScript

JavaScript can represent each DP state using a small object:

```text
{
    score,
    ids
}
```

`null` is used for unreachable states.

The binary search is implemented manually because this version does not rely on a built-in lower-bound function.

Since there are at most four selected indices, copying and sorting the index array is small.

### TypeScript

TypeScript follows the same JavaScript logic but defines the DP state shape explicitly.

The state type makes it easier to understand that every candidate contains:

```text
score
ids
```

and that an unreachable state is represented by `null`.

### Python3

Python stores each state as a tuple:

```text
(score, ids)
```

or `None` when unreachable.

Python's list comparison already follows lexicographical order, which makes tie-breaking very clean.

The interval list is sorted by the right endpoint, and binary search finds the compatible prefix.

### Go

Go uses a small `struct` to represent a DP state.

A boolean field indicates whether the state is valid because Go does not use `null` for ordinary structs.

The selected indices are stored in a slice, and `sort.Ints` keeps them lexicographically comparable.

The sorting and binary search logic are still exactly the same.

## Examples

### Example 1

Input:

```text
intervals = [[1,3,2],[4,5,2],[1,5,5],[6,9,3],[6,7,1],[8,9,1]]
```

The important intervals are:

```text
Index 2 -> [1,5,5]
Index 3 -> [6,9,3]
```

They do not overlap because:

```text
5 < 6
```

Their total score is:

```text
5 + 3 = 8
```

So the output is:

```text
[2,3]
```

### Example 2

Input:

```text
intervals = [[5,8,1],[6,7,7],[4,7,3],[9,10,6],[7,8,2],[11,14,3],[3,5,5]]
```

A maximum-score selection is:

```text
Index 1 -> weight 7
Index 3 -> weight 6
Index 5 -> weight 3
Index 6 -> weight 5
```

Total score:

```text
7 + 6 + 3 + 5 = 21
```

The output is:

```text
[1,3,5,6]
```

The DP finds these intervals while making sure every selected interval is compatible with the previous selection.

### Example 3

Consider two intervals:

```text
Index 0 -> [1,5,10]
Index 1 -> [5,9,10]
```

They cannot both be selected.

Even though:

```text
5 = 5
```

the problem says sharing a boundary means overlapping.

So the maximum score is `10`, and between:

```text
[0]
[1]
```

the lexicographically smaller answer is:

```text
[0]
```

Therefore the correct output is:

```text
[0]
```

## How to Use / Run Locally

The code is written in the standard LeetCode `Solution` format, so it can be pasted directly into the corresponding LeetCode language editor.

For local testing, I can also wrap the solution inside a small driver program.

### C++

Save the solution as:

```text
solution.cpp
```

Compile:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run:

```bash
./solution
```

### Java

Save the file as:

```text
Solution.java
```

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

### JavaScript

Save the file as:

```text
solution.js
```

Run:

```bash
node solution.js
```

### TypeScript

Save the file as:

```text
solution.ts
```

Install TypeScript if needed:

```bash
npm install -g typescript
```

Compile:

```bash
tsc solution.ts
```

Then run:

```bash
node solution.js
```

### Python3

Save the file as:

```text
solution.py
```

Run:

```bash
python3 solution.py
```

### Go

Save the file as:

```text
solution.go
```

Run:

```bash
go run solution.go
```

For local testing, I need to add a small `main` function in C++, Go, Java, JavaScript, TypeScript, or Python3 that creates the input and prints the returned indices.

## Notes & Optimizations

The most important detail is the strict non-overlap condition:

```text
previousRight < currentLeft
```

Do not replace it with `<=`.

Another important detail is that the answer must use **original indices**, not indices after sorting. That is why every interval keeps its original position.

The problem only allows 4 selected intervals, which makes the DP dimension very small:

```text
O(4n)
```

If the problem allowed an arbitrary number `k` of selected intervals, the same dynamic programming idea would become more expensive because the DP would depend on `k`.

I also use a `64-bit` integer type for scores in C++, Java, TypeScript where needed by the implementation, and Go. Four weights can add up to:

```text
4 * 10^9 = 4 * 10^9
```

which does not safely fit in a signed 32-bit integer.

The key optimization is sorting once and using binary search instead of scanning backward for every interval. A backward scan could lead to `O(n^2)` time, while binary search keeps the solution at `O(n log n)`.

This solution is a useful example of combining **sorting, binary search, weighted interval scheduling, dynamic programming, and lexicographical tie-breaking** in one problem.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
