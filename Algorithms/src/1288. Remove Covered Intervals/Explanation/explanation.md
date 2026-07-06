# 1288. Remove Covered Intervals

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

In this problem, I am given an array called `intervals`. Each interval contains two values:

`[start, end]`

An interval `[a, b]` is covered by another interval `[c, d]` when:

`c <= a` and `b <= d`

In simple words, the second interval must start at the same point or earlier and end at the same point or later.

My task is to remove all covered intervals and return the number of intervals that remain.

The main challenge is finding covered intervals efficiently. A brute-force solution can compare every interval with every other interval, but that takes `O(n²)` time. A better solution uses sorting and a greedy scan to solve the problem in `O(n log n)` time.

This approach is a common pattern in interval problems, sorting problems, and greedy algorithm questions.

## Constraints

| Constraint                      | Description                                   |
| ------------------------------- | --------------------------------------------- |
| `1 <= intervals.length <= 1000` | The number of intervals is between 1 and 1000 |
| `intervals[i].length == 2`      | Every interval contains exactly two values    |
| `0 <= li < ri <= 10^5`          | The start is smaller than the end             |
| All intervals are unique        | No two intervals are exactly the same         |

## Intuition

My first thought was to compare every interval with all the other intervals.

For each interval, I could check whether another interval starts earlier and ends later. That would work, but it would require too many comparisons.

Then I noticed that sorting can make the coverage condition much easier to check.

I sort the intervals by their starting point in ascending order. This makes intervals with smaller starting points come first.

If two intervals have the same starting point, I put the interval with the larger ending point first.

For example:

```text
[1, 4]
[1, 8]
```

The correct order should be:

```text
[1, 8]
[1, 4]
```

This is important because `[1, 8]` covers `[1, 4]`.

After sorting, I only need to remember the largest ending point seen so far.

If the current interval ends at or before that point, it is covered.

If it ends after that point, it is not covered, so I count it and update the largest ending point.

## Approach

I solve the Remove Covered Intervals problem with sorting and a greedy scan.

The steps are:

1. Sort all intervals by starting point in ascending order.
2. If two intervals have the same starting point, sort the one with the larger ending point first.
3. Create a variable to store the largest ending point seen so far.
4. Create another variable to count the intervals that are not covered.
5. Go through the sorted intervals one by one.
6. If the current ending point is greater than the largest ending point, the interval is not covered.
7. Count that interval and update the largest ending point.
8. If the current ending point is smaller than or equal to the largest ending point, skip it because it is covered.
9. Return the final count.

The custom sorting order is the most important part of the solution.

Without sorting equal starting points by ending point in descending order, a smaller covered interval could be processed before the larger interval that covers it.

## Data Structures Used

### Array or List of Intervals

The input itself is an array or list of intervals.

Each interval stores two values:

```text
[start, end]
```

I sort this existing collection directly. I do not need another array to store the remaining intervals.

### Integer Variables

I use one variable to store the largest ending point seen so far.

I also use one counter to store how many intervals are not covered.

Because the algorithm only needs these few variables after sorting, no extra data structure that grows with the input size is required.

## Operations & Behavior Summary

The algorithm works in four main stages.

First, I sort the intervals.

The interval with the smaller starting point comes first. If two intervals start at the same point, the interval with the larger ending point comes first.

Second, I initialize the largest ending point with a value smaller than every valid interval end.

Third, I scan the sorted intervals from left to right.

For every interval:

```text
If current end > largest end seen so far:
    Count the interval
    Update the largest end
Otherwise:
    Skip the interval because it is covered
```

Finally, I return the number of intervals that were counted.

This works because every previously processed interval starts at the same point or before the current interval. Therefore, if a previous interval also ends at or after the current interval, the current interval is covered.

## Complexity

| Complexity       | Value              | Explanation                                                                                        |
| ---------------- | ------------------ | -------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n log n)`       | Sorting `n` intervals takes `O(n log n)`, and scanning them takes `O(n)`                           |
| Space Complexity | `O(1)` extra space | Only a few variables are used, ignoring the internal memory required by the sorting implementation |

Here, `n` is the total number of intervals.

The sorting step dominates the total running time, which makes this an efficient solution for LeetCode 1288.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int removeCoveredIntervals(vector<vector<int>>& intervals) {
        // Sort by start in ascending order.
        // For the same start, place the larger end first
        // so the covering interval is processed before the covered one.
        sort(intervals.begin(), intervals.end(),
             [](const vector<int>& a, const vector<int>& b) {
                 if (a[0] == b[0]) {
                     return a[1] > b[1];
                 }
                 return a[0] < b[0];
             });

        // Every valid end is greater than 0, so -1 is a safe initial value.
        int maxEnd = -1;

        // This stores how many intervals are not covered.
        int remaining = 0;

        // Check every interval in the sorted order.
        for (const auto& interval : intervals) {
            // A larger end means no earlier interval can cover this one.
            if (interval[1] > maxEnd) {
                remaining++;

                // Remember the farthest ending point seen so far.
                maxEnd = interval[1];
            }
            // If interval[1] <= maxEnd, an earlier interval covers it,
            // so I do not count it.
        }

        // Return the number of intervals left after removing covered ones.
        return remaining;
    }
}; 
```

### Java

```java
class Solution {
    public int removeCoveredIntervals(int[][] intervals) {
        // Sort by start in ascending order.
        // For equal starts, place the larger end first
        // so the covering interval is processed before the covered one.
        Arrays.sort(intervals, (a, b) -> {
            if (a[0] == b[0]) {
                return Integer.compare(b[1], a[1]);
            }
            return Integer.compare(a[0], b[0]);
        });

        // Every valid end is greater than 0, so -1 is a safe initial value.
        int maxEnd = -1;

        // This stores how many intervals are not covered.
        int remaining = 0;

        // Check every interval in the sorted order.
        for (int[] interval : intervals) {
            // A larger end means no earlier interval can cover this one.
            if (interval[1] > maxEnd) {
                remaining++;

                // Remember the farthest ending point seen so far.
                maxEnd = interval[1];
            }
            // Otherwise, the current interval is covered, so I skip it.
        }

        // Return the number of intervals that remain.
        return remaining;
    }
} 
```

### JavaScript

```javascript
/**
 * @param {number[][]} intervals
 * @return {number}
 */
var removeCoveredIntervals = function(intervals) {
    // Sort by start in ascending order.
    // For equal starts, place the larger end first
    // so the covering interval is processed before the covered one.
    intervals.sort((a, b) => {
        if (a[0] === b[0]) {
            return b[1] - a[1];
        }
        return a[0] - b[0];
    });

    // Every valid end is greater than 0, so -1 is a safe initial value.
    let maxEnd = -1;

    // This stores how many intervals are not covered.
    let remaining = 0;

    // Check every interval in the sorted order.
    for (const interval of intervals) {
        // A larger end means no earlier interval can cover this one.
        if (interval[1] > maxEnd) {
            remaining++;

            // Remember the farthest ending point seen so far.
            maxEnd = interval[1];
        }
        // Otherwise, the interval is covered, so I skip it.
    }

    // Return the number of intervals that remain.
    return remaining;
}; 
```

### Python3

```python
class Solution:
    def removeCoveredIntervals(self, intervals: List[List[int]]) -> int:
        # Sort by start in ascending order.
        # For equal starts, -end puts the larger end first,
        # so the covering interval is processed before the covered one.
        intervals.sort(key=lambda interval: (interval[0], -interval[1]))

        # Every valid end is greater than 0, so -1 is a safe initial value.
        max_end = -1

        # This stores how many intervals are not covered.
        remaining = 0

        # Check every interval in the sorted order.
        for start, end in intervals:
            # A larger end means no earlier interval can cover this one.
            if end > max_end:
                remaining += 1

                # Remember the farthest ending point seen so far.
                max_end = end

            # Otherwise, end <= max_end, so this interval is covered.

        # Return the number of intervals that remain.
        return remaining
```

### Go

```go
func removeCoveredIntervals(intervals [][]int) int {
    // Sort by start in ascending order.
    // For equal starts, place the larger end first
    // so the covering interval is processed before the covered one.
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] > intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    // Every valid end is greater than 0, so -1 is a safe initial value.
    maxEnd := -1

    // This stores how many intervals are not covered.
    remaining := 0

    // Check every interval in the sorted order.
    for _, interval := range intervals {
        // A larger end means no earlier interval can cover this one.
        if interval[1] > maxEnd {
            remaining++

            // Remember the farthest ending point seen so far.
            maxEnd = interval[1]
        }
        // Otherwise, the interval is covered, so I skip it.
    }

    // Return the number of intervals that remain.
    return remaining
} 
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The main algorithm stays the same in C++, Java, JavaScript, Python3, and Go. The only real difference is the syntax used for sorting and looping through the intervals.

### 1. Sort the intervals

The first step is to sort all intervals using two rules.

The primary rule is:

```text
Smaller start comes first
```

The secondary rule is:

```text
If starts are equal, larger end comes first
```

For example:

```text
Input:
[[1, 4], [1, 8], [2, 6]]

Sorted:
[[1, 8], [1, 4], [2, 6]]
```

This order guarantees that a larger interval is processed before a smaller interval with the same starting point.

C++, Java, JavaScript, and Go use a custom comparison function for this sorting rule.

Python3 can express the same rule using a sorting key based on:

```text
(start, -end)
```

The negative ending point makes larger ending values come first when the starting points are equal.

### 2. Track the largest ending point

After sorting, I create a variable that stores the farthest ending point seen so far.

I can initialize it with `-1` because the constraints guarantee that all valid interval values are non-negative.

At the beginning:

```text
maxEnd = -1
```

This guarantees that the first interval will be counted.

### 3. Scan every interval

I now go through the sorted intervals from left to right.

For each interval, I only need its ending point.

If:

```text
currentEnd > maxEnd
```

the current interval reaches farther than every interval processed before it.

That means it cannot be covered.

So I:

```text
Increase the remaining count
Update maxEnd
```

If:

```text
currentEnd <= maxEnd
```

the current interval is covered.

I know this because an earlier interval already starts at the same point or before the current interval, and it also ends at the same point or after it.

So I skip the current interval.

### 4. Why strict greater than matters

The comparison must use:

```text
currentEnd > maxEnd
```

It should not use:

```text
currentEnd >= maxEnd
```

Consider:

```text
[1, 8]
[3, 8]
```

The second interval ends at the same point as the first one.

Since `[1, 8]` starts earlier and ends at the same point, it covers `[3, 8]`.

After processing `[1, 8]`:

```text
maxEnd = 8
```

For `[3, 8]`:

```text
currentEnd = 8
```

Because `8` is not greater than `8`, the interval is correctly skipped.

### 5. Return the result

After checking all intervals, the counter contains the number of intervals that are not covered.

That value is the final answer.

The logic does not change between C++, Java, JavaScript, Python3, and Go. Each language only handles sorting and iteration with different syntax.

## Examples

### Example 1

Input:

```text
intervals = [[1,4], [3,6], [2,8]]
```

Output:

```text
2
```

After sorting:

```text
[[1,4], [2,8], [3,6]]
```

The process is:

```text
[1,4] -> Not covered, count it, maxEnd = 4
[2,8] -> Not covered, count it, maxEnd = 8
[3,6] -> Covered because 6 <= 8
```

The remaining intervals are `[1,4]` and `[2,8]`.

Final answer:

```text
2
```

### Example 2

Input:

```text
intervals = [[1,4], [2,3]]
```

Output:

```text
1
```

After sorting:

```text
[[1,4], [2,3]]
```

The process is:

```text
[1,4] -> Not covered, count it, maxEnd = 4
[2,3] -> Covered because 3 <= 4
```

The interval `[1,4]` completely covers `[2,3]`.

Final answer:

```text
1
```

### Example 3

Input:

```text
intervals = [[1,8], [1,4], [3,8], [4,10]]
```

Output:

```text
2
```

After sorting:

```text
[[1,8], [1,4], [3,8], [4,10]]
```

The process is:

```text
[1,8]  -> Not covered, count it, maxEnd = 8
[1,4]  -> Covered because 4 <= 8
[3,8]  -> Covered because 8 <= 8
[4,10] -> Not covered, count it, maxEnd = 10
```

Final answer:

```text
2
```

## How to Use / Run Locally

The code blocks above are intentionally empty so the solution code can be added later.

After adding the code, save each language version in its own file.

### C++

Save the solution as:

```text
solution.cpp
```

Compile it:

```bash
g++ solution.cpp -o solution
```

Run it:

```bash
./solution
```

On Windows, run:

```bash
solution.exe
```

### Java

Save the solution as:

```text
Solution.java
```

Compile it:

```bash
javac Solution.java
```

Run it:

```bash
java Solution
```

### JavaScript

Save the solution as:

```text
solution.js
```

Make sure Node.js is installed, then run:

```bash
node solution.js
```

### Python3

Save the solution as:

```text
solution.py
```

Run it:

```bash
python3 solution.py
```

On some systems, the command may be:

```bash
python solution.py
```

### Go

Save the solution as:

```text
main.go
```

Run it directly:

```bash
go run main.go
```

Or build it first:

```bash
go build main.go
```

Then run the generated executable.

## Notes & Optimizations

The biggest edge case is when multiple intervals have the same starting point.

The interval with the larger ending point must come first. Otherwise, a covered interval may be counted before the algorithm sees the interval that covers it.

Another important case is when two intervals have the same ending point.

For example:

```text
[1, 5]
[2, 5]
```

The second interval is covered because the first interval starts earlier and ends at the same point.

The greedy solution only needs the largest ending point seen so far. There is no need to store all previous intervals or create a separate list of covered intervals.

A brute-force approach can solve the problem by comparing every pair of intervals, but its time complexity is `O(n²)`. The sorting and greedy approach improves this to `O(n log n)`.

The input intervals are unique, so there is no need to handle exact duplicate intervals.

This same sorting pattern is useful in many DSA interval problems, especially when one interval can contain, overlap, or cover another interval.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
