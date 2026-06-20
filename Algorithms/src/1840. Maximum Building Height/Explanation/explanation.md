# 1840. Maximum Building Height

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

LeetCode 1840 - Maximum Building Height is a challenging greedy and sorting problem.

You are given `n` buildings placed in a straight line and a list of building height restrictions. Every building must satisfy the following rules:

* Building `1` must have height `0`
* Heights cannot be negative
* The height difference between adjacent buildings cannot exceed `1`
* Some buildings have additional maximum height restrictions

The goal is to find the maximum possible height of the tallest building while still satisfying every restriction.

This problem combines sorting, constraint propagation, greedy reasoning, and interval optimization.

---

## Constraints

| Constraint                                     | Value                             |
| ---------------------------------------------- | --------------------------------- |
| `2 <= n <= 10^9`                               | Large building count              |
| `0 <= restrictions.length <= min(n - 1, 10^5)` | Up to 100,000 restrictions        |
| `2 <= idi <= n`                                | Building index                    |
| `idi` is unique                                | No duplicate restricted buildings |
| `0 <= maxHeighti <= 10^9`                      | Height limit                      |

---

## Intuition

The first thing I noticed was that checking every building is impossible because `n` can be as large as `10^9`.

Instead of thinking about every building, I focused only on the restricted buildings.

The adjacent difference rule tells us that height can only increase or decrease by one per position. Because of that, restrictions influence each other.

A restriction that looks valid by itself may become impossible once nearby restrictions are considered.

So before finding the tallest building, I need to make all restrictions consistent with one another.

Once all restrictions become valid, I can examine each interval between consecutive restricted buildings and calculate the highest peak possible inside that interval.

---

## Approach

1. Add building `1` with height `0`.
2. Add building `n` with maximum possible height `n - 1`.
3. Sort all restrictions by building position.
4. Perform a left-to-right pass.

   * Ensure each restriction is reachable from the previous restriction.
5. Perform a right-to-left pass.

   * Ensure each restriction is reachable from the next restriction.
6. After both passes, all restrictions become feasible.
7. For every adjacent pair of restrictions:

   * Calculate the maximum peak achievable between them.
8. Keep track of the largest peak found.
9. Return the answer.

This greedy approach avoids processing all buildings and works efficiently even for the largest constraints.

---

## Data Structures Used

### Array / List

Used to store all restrictions after adding the two extra boundary restrictions.

### Sorting

Sorting restrictions by building index allows us to process neighboring restrictions efficiently.

### Integer Variables

Used for distance calculations, height limits, and peak computation.

No advanced data structures such as heaps, segment trees, or graphs are required.

---

## Operations & Behavior Summary

1. Insert boundary restrictions.
2. Sort restrictions by building index.
3. Traverse from left to right and tighten invalid height limits.
4. Traverse from right to left and tighten invalid height limits again.
5. Every restriction now satisfies neighboring constraints.
6. Process each interval independently.
7. Compute the tallest achievable peak within that interval.
8. Update the global maximum height.
9. Return the largest valid height.

---

## Complexity

| Metric           | Complexity   | Explanation                                                   |
| ---------------- | ------------ | ------------------------------------------------------------- |
| Time Complexity  | `O(k log k)` | Sorting dominates the runtime                                 |
| Space Complexity | `O(k)`       | Stores all restrictions including added boundary restrictions |

Where:

* `k` = total number of restrictions after adding buildings `1` and `n`

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxBuilding(int n, vector<vector<int>>& restrictions) {
        // Building 1 must have height 0
        restrictions.push_back({1, 0});
        
        // Building n can never exceed n - 1
        restrictions.push_back({n, n - 1});

        // Sort restrictions by building index
        sort(restrictions.begin(), restrictions.end());

        int m = restrictions.size();

        // Left to right pass
        // Make sure every restriction is reachable from the left
        for (int i = 1; i < m; i++) {
            int dist = restrictions[i][0] - restrictions[i - 1][0];

            restrictions[i][1] = min(
                restrictions[i][1],
                restrictions[i - 1][1] + dist
            );
        }

        // Right to left pass
        // Make sure every restriction is reachable from the right
        for (int i = m - 2; i >= 0; i--) {
            int dist = restrictions[i + 1][0] - restrictions[i][0];

            restrictions[i][1] = min(
                restrictions[i][1],
                restrictions[i + 1][1] + dist
            );
        }

        long long ans = 0;

        // Compute highest peak inside every interval
        for (int i = 1; i < m; i++) {
            long long x1 = restrictions[i - 1][0];
            long long h1 = restrictions[i - 1][1];

            long long x2 = restrictions[i][0];
            long long h2 = restrictions[i][1];

            long long dist = x2 - x1;

            // Highest achievable height in this segment
            long long peak =
                max(h1, h2) +
                (dist - llabs(h1 - h2)) / 2;

            ans = max(ans, peak);
        }

        return (int)ans;
    }
};
```

### Java

```java
class Solution {
    public int maxBuilding(int n, int[][] restrictions) {
        int m = restrictions.length;

        // Create new array including building 1 and building n
        int[][] arr = new int[m + 2][2];

        for (int i = 0; i < m; i++) {
            arr[i] = restrictions[i];
        }

        // Building 1 has fixed height 0
        arr[m] = new int[]{1, 0};

        // Building n can be at most n - 1
        arr[m + 1] = new int[]{n, n - 1};

        // Sort by building index
        Arrays.sort(arr, (a, b) -> a[0] - b[0]);

        int size = arr.length;

        // Left to right pass
        for (int i = 1; i < size; i++) {
            int dist = arr[i][0] - arr[i - 1][0];

            arr[i][1] = Math.min(
                arr[i][1],
                arr[i - 1][1] + dist
            );
        }

        // Right to left pass
        for (int i = size - 2; i >= 0; i--) {
            int dist = arr[i + 1][0] - arr[i][0];

            arr[i][1] = Math.min(
                arr[i][1],
                arr[i + 1][1] + dist
            );
        }

        long ans = 0;

        // Calculate peak for every interval
        for (int i = 1; i < size; i++) {
            long x1 = arr[i - 1][0];
            long h1 = arr[i - 1][1];

            long x2 = arr[i][0];
            long h2 = arr[i][1];

            long dist = x2 - x1;

            long peak =
                Math.max(h1, h2) +
                (dist - Math.abs(h1 - h2)) / 2;

            ans = Math.max(ans, peak);
        }

        return (int) ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @param {number[][]} restrictions
 * @return {number}
 */
var maxBuilding = function(n, restrictions) {
    // Building 1 must have height 0
    restrictions.push([1, 0]);

    // Building n can be at most n - 1
    restrictions.push([n, n - 1]);

    // Sort by building position
    restrictions.sort((a, b) => a[0] - b[0]);

    const m = restrictions.length;

    // Left to right pass
    for (let i = 1; i < m; i++) {
        const dist = restrictions[i][0] - restrictions[i - 1][0];

        restrictions[i][1] = Math.min(
            restrictions[i][1],
            restrictions[i - 1][1] + dist
        );
    }

    // Right to left pass
    for (let i = m - 2; i >= 0; i--) {
        const dist = restrictions[i + 1][0] - restrictions[i][0];

        restrictions[i][1] = Math.min(
            restrictions[i][1],
            restrictions[i + 1][1] + dist
        );
    }

    let ans = 0;

    // Find tallest peak in each interval
    for (let i = 1; i < m; i++) {
        const x1 = restrictions[i - 1][0];
        const h1 = restrictions[i - 1][1];

        const x2 = restrictions[i][0];
        const h2 = restrictions[i][1];

        const dist = x2 - x1;

        const peak =
            Math.max(h1, h2) +
            Math.floor((dist - Math.abs(h1 - h2)) / 2);

        ans = Math.max(ans, peak);
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def maxBuilding(self, n: int, restrictions: List[List[int]]) -> int:
        # Building 1 must have height 0
        restrictions.append([1, 0])

        # Building n can be at most n - 1
        restrictions.append([n, n - 1])

        # Sort by building index
        restrictions.sort()

        m = len(restrictions)

        # Left to right pass
        for i in range(1, m):
            dist = restrictions[i][0] - restrictions[i - 1][0]

            restrictions[i][1] = min(
                restrictions[i][1],
                restrictions[i - 1][1] + dist
            )

        # Right to left pass
        for i in range(m - 2, -1, -1):
            dist = restrictions[i + 1][0] - restrictions[i][0]

            restrictions[i][1] = min(
                restrictions[i][1],
                restrictions[i + 1][1] + dist
            )

        ans = 0

        # Compute peak for every interval
        for i in range(1, m):
            x1, h1 = restrictions[i - 1]
            x2, h2 = restrictions[i]

            dist = x2 - x1

            peak = max(h1, h2) + (
                dist - abs(h1 - h2)
            ) // 2

            ans = max(ans, peak)

        return ans
```

### Go

```go
func maxBuilding(n int, restrictions [][]int) int {
 // Building 1 must have height 0
 restrictions = append(restrictions, []int{1, 0})

 // Building n can be at most n - 1
 restrictions = append(restrictions, []int{n, n - 1})

 // Sort restrictions by building index
 sort.Slice(restrictions, func(i, j int) bool {
  return restrictions[i][0] < restrictions[j][0]
 })

 m := len(restrictions)

 // Left to right pass
 for i := 1; i < m; i++ {
  dist := restrictions[i][0] - restrictions[i-1][0]

  restrictions[i][1] = min(
   restrictions[i][1],
   restrictions[i-1][1]+dist,
  )
 }

 // Right to left pass
 for i := m - 2; i >= 0; i-- {
  dist := restrictions[i+1][0] - restrictions[i][0]

  restrictions[i][1] = min(
   restrictions[i][1],
   restrictions[i+1][1]+dist,
  )
 }

 ans := 0

 // Calculate peak in every interval
 for i := 1; i < m; i++ {
  h1 := restrictions[i-1][1]
  h2 := restrictions[i][1]

  dist := restrictions[i][0] - restrictions[i-1][0]

  peak := max(h1, h2) +
   (dist-abs(h1-h2))/2

  ans = max(ans, peak)
 }

 return ans
}

func min(a, b int) int {
 if a < b {
  return a
 }
 return b
}

func max(a, b int) int {
 if a > b {
  return a
 }
 return b
}

func abs(x int) int {
 if x < 0 {
  return -x
 }
 return x
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical across all five languages.

### Step 1: Add Boundary Restrictions

The problem already guarantees:

* Building `1` must have height `0`

So we explicitly add:

```text
(1, 0)
```

We also add:

```text
(n, n - 1)
```

Even if building `n` is not restricted, it can never exceed `n - 1` because height increases by at most one per building.

These two restrictions make the entire range easier to process.

---

### Step 2: Sort Restrictions

After inserting the boundary restrictions, all restrictions are sorted according to building index.

This ensures neighboring restrictions become adjacent in the array.

Without sorting, distance calculations would be incorrect.

---

### Step 3: Left-to-Right Validation

For every restriction:

```text
currentHeight <= previousHeight + distance
```

If a restriction violates this condition, it must be reduced.

This pass guarantees every restriction is reachable from the left side.

---

### Step 4: Right-to-Left Validation

Now process restrictions from right to left.

For every restriction:

```text
currentHeight <= nextHeight + distance
```

If the current height is too large, reduce it.

This pass guarantees every restriction is reachable from the right side.

---

### Step 5: Restrictions Become Consistent

After both passes:

* Every restriction satisfies neighboring constraints
* No contradiction remains
* Any valid building arrangement can now be constructed

This is the key observation behind the solution.

---

### Step 6: Process Each Interval

Consider two neighboring restrictions:

```text
(x1, h1)
(x2, h2)
```

Distance:

```text
dist = x2 - x1
```

The tallest building inside this interval occurs when heights rise toward a peak and then decrease toward the second restriction.

---

### Step 7: Calculate Maximum Peak

Some distance is already consumed matching the difference:

```text
abs(h1 - h2)
```

The remaining distance can be used to climb higher.

Peak formula:

```text
max(h1, h2) + (dist - abs(h1 - h2)) / 2
```

This gives the tallest achievable height in that interval.

---

### Step 8: Track Global Maximum

For every interval:

1. Compute peak
2. Compare with answer
3. Keep the largest value

The final maximum is returned.

---

## Examples

### Example 1

Input

```text
n = 5
restrictions = [[2,1],[4,1]]
```

Output

```text
2
```

Trace

```text
Valid heights:
0 1 2 1 2
```

Tallest building:

```text
2
```

---

### Example 2

Input

```text
n = 6
restrictions = []
```

Output

```text
5
```

Trace

```text
0 1 2 3 4 5
```

No restrictions limit the growth.

Tallest building:

```text
5
```

---

### Example 3

Input

```text
n = 10
restrictions = [[5,3],[2,5],[7,4],[10,3]]
```

Output

```text
5
```

Trace

```text
0 1 2 3 3 4 4 5 4 3
```

Tallest building:

```text
5
```

---

## How to Use / Run Locally

### C++

Compile

```bash
g++ main.cpp -o main
```

Run

```bash
./main
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
go run main.go
```

Build

```bash
go build main.go
```

---

## Notes & Optimizations

* The main trick is realizing that only restricted buildings matter.
* Processing all buildings is impossible because `n` can reach `10^9`.
* The two-pass restriction adjustment is essential.
* Sorting allows neighboring restrictions to be processed efficiently.
* The peak formula removes the need to simulate heights one building at a time.
* This is the optimal solution for the given constraints.
* Any brute-force or simulation approach will exceed time limits.
* The solution combines greedy reasoning, interval analysis, and sorting.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
