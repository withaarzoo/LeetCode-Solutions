# 3635. Earliest Finish Time for Land and Water Rides II

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

LeetCode 3635 - Earliest Finish Time for Land and Water Rides II is a binary search and sorting problem where a tourist must take exactly one land ride and one water ride.

Each ride has:

* An opening time
* A duration

The tourist can choose either order:

* Land ride → Water ride
* Water ride → Land ride

If the second ride is already open when the first ride finishes, it can be taken immediately. Otherwise, the tourist must wait until that ride opens.

The goal is to find the earliest possible time at which both rides can be completed.

This problem looks simple at first, but a brute-force solution would require checking every possible land-water pair, which is too slow for the given constraints. An optimized sorting, prefix minimum, suffix minimum, and binary search solution is required.

## Constraints

| Constraint                                       | Value                 |
| ------------------------------------------------ | --------------------- |
| 1 ≤ n ≤ 5 × 10⁴                                  | Number of land rides  |
| 1 ≤ m ≤ 5 × 10⁴                                  | Number of water rides |
| landStartTime.length = landDuration.length = n   | Guaranteed            |
| waterStartTime.length = waterDuration.length = m | Guaranteed            |
| 1 ≤ start time ≤ 10⁵                             | Valid range           |
| 1 ≤ duration ≤ 10⁵                               | Valid range           |

## Intuition

My first thought was to try every possible pair of rides.

For every land ride, I could check every water ride and compute the final finishing time.

The problem is that this creates an O(n × m) solution, which becomes far too slow when both arrays contain up to 50,000 rides.

After looking more carefully, I noticed that once the first ride finishes at a certain time, every second ride falls into one of two groups:

* Rides that are already open
* Rides that have not opened yet

For already-open rides, only the smallest duration matters.

For rides that open later, only the smallest value of start + duration matters.

That observation makes it possible to preprocess information and answer each query efficiently using binary search.

## Approach

1. Sort the rides of the second category by opening time.
2. Build a prefix minimum array storing the smallest duration seen so far.
3. Build a suffix minimum array storing the smallest finish value `(start + duration)` from the right.
4. For each ride chosen first:

   * Calculate when it finishes.
   * Use binary search to split rides into:

     * Start time ≤ finish time
     * Start time > finish time
5. Evaluate the best possible ride from both groups.
6. Keep track of the smallest finishing time.
7. Run the same logic twice:

   * Land first, water second
   * Water first, land second
8. Return the minimum answer.

## Data Structures Used

### Array

Used for storing:

* Start times
* Prefix minimum durations
* Suffix minimum finish values

Arrays provide O(1) access and are ideal for preprocessing.

### Pair / Tuple

Each ride is stored as:

```text
(startTime, duration)
```

This makes sorting by opening time straightforward.

### Binary Search

Used to quickly find the first ride whose opening time is greater than the finishing time of the first ride.

This reduces searching from O(n) to O(log n).

## Operations & Behavior Summary

1. Convert rides into sortable records.
2. Sort rides by opening time.
3. Build prefix minimum duration array.
4. Build suffix minimum finish array.
5. For every ride in the first category:

   * Compute finish time.
   * Binary search the sorted second category.
   * Check rides already available.
   * Check rides opening later.
6. Update the global answer.
7. Repeat for the reverse order.
8. Return the minimum finishing time found.

## Complexity

| Metric           | Complexity                               | Explanation                               |
| ---------------- | ---------------------------------------- | ----------------------------------------- |
| Time Complexity  | O(n log n + m log m + n log m + m log n) | Sorting plus binary search for every ride |
| Space Complexity | O(n + m)                                 | Additional arrays for preprocessing       |

Where:

* n = number of land rides
* m = number of water rides

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    // Computes the best answer when rides A are taken first
    // and rides B are taken second.
    long long solve(vector<int>& startA, vector<int>& durA,
                    vector<int>& startB, vector<int>& durB) {

        int m = startB.size();

        // Store second-category rides as (start, duration)
        vector<pair<int, int>> rides(m);
        for (int i = 0; i < m; i++) {
            rides[i] = {startB[i], durB[i]};
        }

        // Sort by opening time
        sort(rides.begin(), rides.end());

        vector<int> starts(m);

        // prefixMinDur[i] = minimum duration in [0..i]
        vector<long long> prefixMinDur(m);

        // suffixMinFinish[i] = minimum (start + duration) in [i..m-1]
        vector<long long> suffixMinFinish(m);

        for (int i = 0; i < m; i++) {
            starts[i] = rides[i].first;

            if (i == 0)
                prefixMinDur[i] = rides[i].second;
            else
                prefixMinDur[i] = min(prefixMinDur[i - 1],
                                      (long long)rides[i].second);
        }

        for (int i = m - 1; i >= 0; i--) {
            long long finish = (long long)rides[i].first + rides[i].second;

            if (i == m - 1)
                suffixMinFinish[i] = finish;
            else
                suffixMinFinish[i] = min(suffixMinFinish[i + 1], finish);
        }

        long long ans = LLONG_MAX;

        for (int i = 0; i < (int)startA.size(); i++) {

            // Finish time after taking first ride
            long long finish1 = (long long)startA[i] + durA[i];

            // First index with start > finish1
            int pos = upper_bound(starts.begin(), starts.end(), finish1)
                      - starts.begin();

            // Rides already open
            if (pos > 0) {
                ans = min(ans, finish1 + prefixMinDur[pos - 1]);
            }

            // Rides opening later
            if (pos < m) {
                ans = min(ans, suffixMinFinish[pos]);
            }
        }

        return ans;
    }

    int earliestFinishTime(vector<int>& landStartTime, vector<int>& landDuration, vector<int>& waterStartTime, vector<int>& waterDuration) {

        long long ans1 = solve(
            landStartTime, landDuration,
            waterStartTime, waterDuration
        );

        long long ans2 = solve(
            waterStartTime, waterDuration,
            landStartTime, landDuration
        );

        return (int)min(ans1, ans2);
    }
};
```

### Java

```java
class Solution {

    // Computes the best answer when category A is taken first
    private long solve(int[] startA, int[] durA,
                       int[] startB, int[] durB) {

        int m = startB.length;

        int[][] rides = new int[m][2];

        // Store (start, duration)
        for (int i = 0; i < m; i++) {
            rides[i][0] = startB[i];
            rides[i][1] = durB[i];
        }

        // Sort by start time
        java.util.Arrays.sort(rides, (a, b) -> Integer.compare(a[0], b[0]));

        int[] starts = new int[m];
        long[] prefixMinDur = new long[m];
        long[] suffixMinFinish = new long[m];

        for (int i = 0; i < m; i++) {
            starts[i] = rides[i][0];

            if (i == 0)
                prefixMinDur[i] = rides[i][1];
            else
                prefixMinDur[i] = Math.min(prefixMinDur[i - 1], rides[i][1]);
        }

        for (int i = m - 1; i >= 0; i--) {
            long finish = (long) rides[i][0] + rides[i][1];

            if (i == m - 1)
                suffixMinFinish[i] = finish;
            else
                suffixMinFinish[i] = Math.min(suffixMinFinish[i + 1], finish);
        }

        long ans = Long.MAX_VALUE;

        for (int i = 0; i < startA.length; i++) {

            long finish1 = (long) startA[i] + durA[i];

            int pos = upperBound(starts, finish1);

            if (pos > 0) {
                ans = Math.min(ans, finish1 + prefixMinDur[pos - 1]);
            }

            if (pos < m) {
                ans = Math.min(ans, suffixMinFinish[pos]);
            }
        }

        return ans;
    }

    // First index with value > target
    private int upperBound(int[] arr, long target) {
        int left = 0;
        int right = arr.length;

        while (left < right) {
            int mid = left + (right - left) / 2;

            if (arr[mid] <= target)
                left = mid + 1;
            else
                right = mid;
        }

        return left;
    }

    public int earliestFinishTime(int[] landStartTime, int[] landDuration, int[] waterStartTime, int[] waterDuration) {

        long ans1 = solve(
            landStartTime, landDuration,
            waterStartTime, waterDuration
        );

        long ans2 = solve(
            waterStartTime, waterDuration,
            landStartTime, landDuration
        );

        return (int) Math.min(ans1, ans2);
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} landStartTime
 * @param {number[]} landDuration
 * @param {number[]} waterStartTime
 * @param {number[]} waterDuration
 * @return {number}
 */
var earliestFinishTime = function(landStartTime, landDuration, waterStartTime, waterDuration) {

    // Computes the best answer when category A is taken first
    const solve = (startA, durA, startB, durB) => {

        const rides = [];

        // Store (start, duration)
        for (let i = 0; i < startB.length; i++) {
            rides.push([startB[i], durB[i]]);
        }

        // Sort by start time
        rides.sort((a, b) => a[0] - b[0]);

        const m = rides.length;

        const starts = new Array(m);
        const prefixMinDur = new Array(m);
        const suffixMinFinish = new Array(m);

        for (let i = 0; i < m; i++) {
            starts[i] = rides[i][0];

            if (i === 0)
                prefixMinDur[i] = rides[i][1];
            else
                prefixMinDur[i] = Math.min(prefixMinDur[i - 1], rides[i][1]);
        }

        for (let i = m - 1; i >= 0; i--) {
            const finish = rides[i][0] + rides[i][1];

            if (i === m - 1)
                suffixMinFinish[i] = finish;
            else
                suffixMinFinish[i] = Math.min(suffixMinFinish[i + 1], finish);
        }

        let ans = Number.MAX_SAFE_INTEGER;

        for (let i = 0; i < startA.length; i++) {

            const finish1 = startA[i] + durA[i];

            let left = 0;
            let right = m;

            // Upper bound: first start > finish1
            while (left < right) {
                const mid = Math.floor((left + right) / 2);

                if (starts[mid] <= finish1)
                    left = mid + 1;
                else
                    right = mid;
            }

            const pos = left;

            if (pos > 0) {
                ans = Math.min(ans, finish1 + prefixMinDur[pos - 1]);
            }

            if (pos < m) {
                ans = Math.min(ans, suffixMinFinish[pos]);
            }
        }

        return ans;
    };

    return Math.min(
        solve(landStartTime, landDuration, waterStartTime, waterDuration),
        solve(waterStartTime, waterDuration, landStartTime, landDuration)
    );
};
```

### Python3

```python
class Solution:
    def earliestFinishTime(self, landStartTime: List[int], landDuration: List[int], waterStartTime: List[int], waterDuration: List[int]) -> int:

        from bisect import bisect_right

        # Computes the best answer when category A is taken first
        def solve(startA, durA, startB, durB):

            # Store (start, duration)
            rides = sorted(zip(startB, durB))

            m = len(rides)

            starts = [0] * m
            prefix_min_dur = [0] * m
            suffix_min_finish = [0] * m

            for i in range(m):
                starts[i] = rides[i][0]

                if i == 0:
                    prefix_min_dur[i] = rides[i][1]
                else:
                    prefix_min_dur[i] = min(
                        prefix_min_dur[i - 1],
                        rides[i][1]
                    )

            for i in range(m - 1, -1, -1):
                finish = rides[i][0] + rides[i][1]

                if i == m - 1:
                    suffix_min_finish[i] = finish
                else:
                    suffix_min_finish[i] = min(
                        suffix_min_finish[i + 1],
                        finish
                    )

            ans = float("inf")

            for s, d in zip(startA, durA):

                # Finish time of first ride
                finish1 = s + d

                # First ride with start > finish1
                pos = bisect_right(starts, finish1)

                if pos > 0:
                    ans = min(
                        ans,
                        finish1 + prefix_min_dur[pos - 1]
                    )

                if pos < m:
                    ans = min(
                        ans,
                        suffix_min_finish[pos]
                    )

            return ans

        return min(
            solve(
                landStartTime,
                landDuration,
                waterStartTime,
                waterDuration
            ),
            solve(
                waterStartTime,
                waterDuration,
                landStartTime,
                landDuration
            )
        )
```

### Go

```go
func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {

 type Ride struct {
  start int
  dur   int
 }

 // Computes the best answer when category A is taken first
 var solve func([]int, []int, []int, []int) int64

 solve = func(startA []int, durA []int,
  startB []int, durB []int) int64 {

  m := len(startB)

  rides := make([]Ride, m)

  // Store (start, duration)
  for i := 0; i < m; i++ {
   rides[i] = Ride{startB[i], durB[i]}
  }

  // Sort by start time
  sort.Slice(rides, func(i, j int) bool {
   return rides[i].start < rides[j].start
  })

  starts := make([]int, m)
  prefixMinDur := make([]int64, m)
  suffixMinFinish := make([]int64, m)

  for i := 0; i < m; i++ {
   starts[i] = rides[i].start

   if i == 0 {
    prefixMinDur[i] = int64(rides[i].dur)
   } else {
    prefixMinDur[i] = min64(
     prefixMinDur[i-1],
     int64(rides[i].dur),
    )
   }
  }

  for i := m - 1; i >= 0; i-- {
   finish := int64(rides[i].start + rides[i].dur)

   if i == m-1 {
    suffixMinFinish[i] = finish
   } else {
    suffixMinFinish[i] = min64(
     suffixMinFinish[i+1],
     finish,
    )
   }
  }

  ans := int64(1 << 60)

  for i := 0; i < len(startA); i++ {

   // Finish time of first ride
   finish1 := int64(startA[i] + durA[i])

   // First index with start > finish1
   pos := sort.Search(m, func(j int) bool {
    return int64(starts[j]) > finish1
   })

   if pos > 0 {
    ans = min64(
     ans,
     finish1+prefixMinDur[pos-1],
    )
   }

   if pos < m {
    ans = min64(
     ans,
     suffixMinFinish[pos],
    )
   }
  }

  return ans
 }

 ans1 := solve(
  landStartTime,
  landDuration,
  waterStartTime,
  waterDuration,
 )

 ans2 := solve(
  waterStartTime,
  waterDuration,
  landStartTime,
  landDuration,
 )

 if ans1 < ans2 {
  return int(ans1)
 }
 return int(ans2)
}

func min64(a, b int64) int64 {
 if a < b {
  return a
 }
 return b
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical in all five implementations.

### Step 1: Sort the second ride category

The rides are sorted by opening time.

After sorting, all rides that open before a certain moment form one continuous section, and all later rides form another section.

This property is what allows binary search to work.

### Step 2: Build prefix minimum durations

For rides that are already open when the first ride finishes, the final completion time becomes:

```text
finishFirstRide + duration
```

Since finishFirstRide is fixed, only the smallest duration matters.

The prefix array stores that information.

### Step 3: Build suffix minimum finish values

For rides that are not open yet, the tourist must wait.

The completion time becomes:

```text
start + duration
```

The suffix array stores the smallest possible value of this expression for every position.

### Step 4: Process every first ride

For each ride:

```text
finishTime = start + duration
```

This is the earliest moment the second ride can begin.

### Step 5: Binary search

Find the first ride whose opening time is greater than finishTime.

This creates two groups:

Group 1:

```text
start <= finishTime
```

Group 2:

```text
start > finishTime
```

### Step 6: Evaluate both groups

For Group 1:

```text
finishTime + minimumDuration
```

For Group 2:

```text
minimum(start + duration)
```

Take the smaller result.

### Step 7: Repeat for both ride orders

The tourist may choose:

```text
Land → Water
```

or

```text
Water → Land
```

Both possibilities must be checked.

### Step 8: Return the best answer

The smallest completion time across all valid plans becomes the final answer.

## Examples

### Example 1

Input

```text
landStartTime = [2,8]
landDuration = [4,1]
waterStartTime = [6]
waterDuration = [3]
```

Output

```text
9
```

Explanation

```text
Land ride 0:
Finish = 2 + 4 = 6

Water ride opens at 6

Finish = 6 + 3 = 9
```

No other ordering produces a smaller answer.

---

### Example 2

Input

```text
landStartTime = [5]
landDuration = [3]
waterStartTime = [1]
waterDuration = [10]
```

Output

```text
14
```

Explanation

```text
Water ride:
1 -> 11

Land ride:
11 -> 14
```

Earliest completion time is 14.

---

### Example 3

Input

```text
landStartTime = [1]
landDuration = [2]
waterStartTime = [10]
waterDuration = [1]
```

Output

```text
11
```

Explanation

```text
Land ride:
1 -> 3

Wait until 10

Water ride:
10 -> 11
```

Final answer is 11.

## How to Use / Run Locally

### C++

Compile

```bash
g++ solution.cpp -o solution
```

Run

```bash
./solution
```

### Java

Compile

```bash
javac Solution.java
```

Run

```bash
java Solution
```

### JavaScript

Run

```bash
node solution.js
```

### Python3

Run

```bash
python solution.py
```

### Go

Run

```bash
go run solution.go
```

## Notes & Optimizations

* A brute-force solution requires checking every pair of rides and results in O(n × m) complexity.
* Sorting allows binary search to be used efficiently.
* Prefix minimum arrays remove the need to repeatedly search for the shortest duration.
* Suffix minimum arrays remove the need to repeatedly search for the smallest future finish time.
* The solution scales comfortably to the maximum constraint size.
* Long integer types should be used where appropriate to avoid overflow issues in languages that distinguish integer sizes.
* The same preprocessing strategy works regardless of whether land rides or water rides are chosen first.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
