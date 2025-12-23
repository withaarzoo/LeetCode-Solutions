# 2054. Two Best Non-Overlapping Events

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

You are given a list of events.
Each event has:

* a start time
* an end time
* a value

You can attend **at most two events** such that:

* The events **do not overlap**
* Time is **inclusive**
* The total value is **maximum**

Your task is to return the **maximum possible value** you can get.

---

## Constraints

* `2 ≤ events.length ≤ 10^5`
* `events[i].length == 3`
* `1 ≤ startTime ≤ endTime ≤ 10^9`
* `1 ≤ value ≤ 10^6`

---

## Intuition

When I first saw this problem, I noticed two key points:

1. I can attend **only two events**, not more.
2. Events must **not overlap**, and time is inclusive.

So I thought:

* If I fix one event as the **second event**
* Then I just need to know the **best event that ended before it started**

This means:

* Sorting will help
* Binary search or prefix maximum will help

Brute force checking all pairs would be too slow, so I needed something faster.

---

## Approach

I solved the problem using these steps:

1. Sort all events by **start time**
2. Create another list sorted by **end time**
3. Build a prefix maximum array of event values (based on end time)
4. Traverse events by start time:

   * Try taking the event alone
   * Or combine it with the best previous non-overlapping event
5. Keep updating the maximum answer

This ensures efficiency and correctness.

---

## Data Structures Used

* Arrays / Lists (to store events)
* Prefix Maximum Array
* Sorting (based on start time and end time)

---

## Operations & Behavior Summary

* Sorting events by time
* Tracking maximum value of completed events
* Ensuring no overlapping using end time checks
* Choosing the best combination of at most two events

---

## Complexity

**Time Complexity:**
`O(n log n)`

* Sorting takes `O(n log n)`
* Traversal is linear

**Space Complexity:**
`O(n)`

* Extra arrays for sorted events and prefix max values

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxTwoEvents(vector<vector<int>>& events) {
        sort(events.begin(), events.end());

        vector<vector<int>> endSorted = events;
        sort(endSorted.begin(), endSorted.end(),
             [](auto &a, auto &b) {
                 return a[1] < b[1];
             });

        int n = events.size();
        vector<int> maxValueTill(n);
        maxValueTill[0] = endSorted[0][2];

        for (int i = 1; i < n; i++) {
            maxValueTill[i] = max(maxValueTill[i - 1], endSorted[i][2]);
        }

        int ans = 0, j = 0;

        for (int i = 0; i < n; i++) {
            int start = events[i][0];
            int value = events[i][2];

            while (j < n && endSorted[j][1] < start) {
                j++;
            }

            ans = max(ans, value);
            if (j > 0) {
                ans = max(ans, value + maxValueTill[j - 1]);
            }
        }

        return ans;
    }
};
```

---

### Java

```java
class Solution {
    public int maxTwoEvents(int[][] events) {
        Arrays.sort(events, (a, b) -> a[0] - b[0]);

        int[][] endSorted = events.clone();
        Arrays.sort(endSorted, (a, b) -> a[1] - b[1]);

        int n = events.length;
        int[] maxValueTill = new int[n];
        maxValueTill[0] = endSorted[0][2];

        for (int i = 1; i < n; i++) {
            maxValueTill[i] = Math.max(maxValueTill[i - 1], endSorted[i][2]);
        }

        int ans = 0, j = 0;

        for (int i = 0; i < n; i++) {
            int start = events[i][0];
            int value = events[i][2];

            while (j < n && endSorted[j][1] < start) {
                j++;
            }

            ans = Math.max(ans, value);
            if (j > 0) {
                ans = Math.max(ans, value + maxValueTill[j - 1]);
            }
        }

        return ans;
    }
}
```

---

### JavaScript

```javascript
var maxTwoEvents = function(events) {
    events.sort((a, b) => a[0] - b[0]);
    const endSorted = [...events].sort((a, b) => a[1] - b[1]);

    const n = events.length;
    const maxValueTill = new Array(n);
    maxValueTill[0] = endSorted[0][2];

    for (let i = 1; i < n; i++) {
        maxValueTill[i] = Math.max(maxValueTill[i - 1], endSorted[i][2]);
    }

    let ans = 0, j = 0;

    for (let i = 0; i < n; i++) {
        const [start, , value] = events[i];

        while (j < n && endSorted[j][1] < start) {
            j++;
        }

        ans = Math.max(ans, value);
        if (j > 0) {
            ans = Math.max(ans, value + maxValueTill[j - 1]);
        }
    }

    return ans;
};
```

---

### Python3

```python
class Solution:
    def maxTwoEvents(self, events):
        events.sort()
        end_sorted = sorted(events, key=lambda x: x[1])

        n = len(events)
        max_value_till = [0] * n
        max_value_till[0] = end_sorted[0][2]

        for i in range(1, n):
            max_value_till[i] = max(max_value_till[i - 1], end_sorted[i][2])

        ans = 0
        j = 0

        for start, end, value in events:
            while j < n and end_sorted[j][1] < start:
                j += 1

            ans = max(ans, value)
            if j > 0:
                ans = max(ans, value + max_value_till[j - 1])

        return ans
```

---

### Go

```go
func maxTwoEvents(events [][]int) int {
    sort.Slice(events, func(i, j int) bool {
        return events[i][0] < events[j][0]
    })

    endSorted := make([][]int, len(events))
    copy(endSorted, events)
    sort.Slice(endSorted, func(i, j int) bool {
        return endSorted[i][1] < endSorted[j][1]
    })

    n := len(events)
    maxValueTill := make([]int, n)
    maxValueTill[0] = endSorted[0][2]

    for i := 1; i < n; i++ {
        maxValueTill[i] = max(maxValueTill[i-1], endSorted[i][2])
    }

    ans, j := 0, 0

    for i := 0; i < n; i++ {
        start := events[i][0]
        value := events[i][2]

        for j < n && endSorted[j][1] < start {
            j++
        }

        ans = max(ans, value)
        if j > 0 {
            ans = max(ans, value+maxValueTill[j-1])
        }
    }

    return ans
}
```

---

## Step-by-step Detailed Explanation

1. Sort events by start time so I process them in order.
2. Sort another copy by end time to track finished events.
3. Build a prefix max array to know the best value till any index.
4. For each event:

   * Move pointer for events that ended earlier.
   * Try taking the event alone.
   * Try combining it with the best previous event.
5. Update the answer.

---

## Examples

**Input**

```
events = [[1,3,2],[4,5,2],[2,4,3]]
```

**Output**

```
4
```

**Explanation**
Choose events `[1,3,2]` and `[4,5,2]`.

---

## How to use / Run locally

1. Clone the repository
2. Open the solution file in your preferred language
3. Run using the standard compiler/interpreter:

   * `g++`, `javac`, `node`, `python`, `go run`

---

## Notes & Optimizations

* Avoid brute force (O(n²)) approaches
* Sorting + prefix max gives optimal performance
* Works efficiently even for large inputs

---

## Author

* **Md Aarzoo Islam**
  [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
