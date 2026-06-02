# 3633. Earliest Finish Time for Land and Water Rides I

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

Given two categories of theme park attractions:

* Land rides
* Water rides

A tourist must experience exactly one ride from each category.

The rides can be taken in any order:

* Land Ride → Water Ride
* Water Ride → Land Ride

Each ride has:

* An opening time
* A duration

A ride may start at its opening time or any time after that.

The goal is to find the earliest possible time at which the tourist can finish both rides.

This is a classic brute force simulation problem involving scheduling, waiting times, and finding the minimum completion time across all possible ride combinations.

Relevant DSA and competitive programming topics:

* Arrays
* Simulation
* Brute Force
* Greedy Observation
* Minimum Finish Time Calculation

---

## Constraints

| Constraint                | Value            |
| ------------------------- | ---------------- |
| Number of land rides (n)  | 1 ≤ n ≤ 100      |
| Number of water rides (m) | 1 ≤ m ≤ 100      |
| landStartTime.length      | n                |
| landDuration.length       | n                |
| waterStartTime.length     | m                |
| waterDuration.length      | m                |
| Ride times and durations  | 1 ≤ value ≤ 1000 |

---

## Intuition

My first observation was that the constraints are very small.

Since there are at most 100 land rides and 100 water rides, there can only be:

100 × 100 = 10,000 ride pairs.

For every selected pair, only two valid orders exist:

1. Land → Water
2. Water → Land

That means I can simply test both possibilities for every pair and keep track of the smallest finishing time.

There is no need for advanced data structures, sorting, binary search, or dynamic programming.

A direct simulation is enough.

---

## Approach

1. Loop through every land ride.
2. Loop through every water ride.
3. Calculate the finishing time if the land ride is taken first.
4. Calculate the finishing time if the water ride is taken first.
5. Take the minimum of the two schedules.
6. Update the global answer.
7. Return the smallest finishing time found.

The key idea is that whenever the second ride is not open yet, the tourist must wait.

That waiting time can be handled using the `max()` function.

---

## Data Structures Used

### Arrays

The input is already provided as arrays.

I only read values from these arrays while calculating finishing times.

### Integer Variables

A few integer variables are used to store:

* Ride finish times
* Ride start times
* Current minimum answer

No additional containers or data structures are required.

---

## Operations & Behavior Summary

The algorithm performs the following steps:

1. Select one land ride.

2. Select one water ride.

3. Simulate:

   * Land → Water

4. Simulate:

   * Water → Land

5. Compute completion time for both schedules.

6. Keep the smaller completion time.

7. Compare it against the global minimum answer.

8. Continue until every pair has been checked.

9. Return the best answer.

This guarantees that every valid ride schedule is evaluated.

---

## Complexity

| Metric           | Complexity | Explanation                                     |
| ---------------- | ---------- | ----------------------------------------------- |
| Time Complexity  | O(n × m)   | Every land ride is paired with every water ride |
| Space Complexity | O(1)       | Only a few variables are used                   |

Where:

* n = number of land rides
* m = number of water rides

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int earliestFinishTime(vector<int>& landStartTime, vector<int>& landDuration,
                           vector<int>& waterStartTime, vector<int>& waterDuration) {

        // Store the minimum finishing time found so far
        int ans = INT_MAX;

        int n = landStartTime.size();
        int m = waterStartTime.size();

        // Try every land ride with every water ride
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {

                // ----------------------------
                // Option 1: Land -> Water
                // ----------------------------

                // Time when land ride finishes
                int landFinish = landStartTime[i] + landDuration[i];

                // Water ride can start only after both:
                // 1. land ride is finished
                // 2. water ride is open
                int waterStart = max(landFinish, waterStartTime[j]);

                // Final finishing time for this order
                int finish1 = waterStart + waterDuration[j];

                // ----------------------------
                // Option 2: Water -> Land
                // ----------------------------

                // Time when water ride finishes
                int waterFinish = waterStartTime[j] + waterDuration[j];

                // Land ride can start only after both:
                // 1. water ride is finished
                // 2. land ride is open
                int landStart = max(waterFinish, landStartTime[i]);

                // Final finishing time for this order
                int finish2 = landStart + landDuration[i];

                // Update answer with the better option
                ans = min(ans, min(finish1, finish2));
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int earliestFinishTime(int[] landStartTime, int[] landDuration,
                                  int[] waterStartTime, int[] waterDuration) {

        // Store the minimum finishing time found
        int ans = Integer.MAX_VALUE;

        int n = landStartTime.length;
        int m = waterStartTime.length;

        // Try every pair of rides
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {

                // Land -> Water
                int landFinish = landStartTime[i] + landDuration[i];
                int waterStart = Math.max(landFinish, waterStartTime[j]);
                int finish1 = waterStart + waterDuration[j];

                // Water -> Land
                int waterFinish = waterStartTime[j] + waterDuration[j];
                int landStart = Math.max(waterFinish, landStartTime[i]);
                int finish2 = landStart + landDuration[i];

                // Keep smallest finish time
                ans = Math.min(ans, Math.min(finish1, finish2));
            }
        }

        return ans;
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

    // Store minimum answer
    let ans = Number.MAX_SAFE_INTEGER;

    const n = landStartTime.length;
    const m = waterStartTime.length;

    // Try every land-water pair
    for (let i = 0; i < n; i++) {
        for (let j = 0; j < m; j++) {

            // Land -> Water
            const landFinish = landStartTime[i] + landDuration[i];
            const waterStart = Math.max(landFinish, waterStartTime[j]);
            const finish1 = waterStart + waterDuration[j];

            // Water -> Land
            const waterFinish = waterStartTime[j] + waterDuration[j];
            const landStart = Math.max(waterFinish, landStartTime[i]);
            const finish2 = landStart + landDuration[i];

            // Update minimum answer
            ans = Math.min(ans, finish1, finish2);
        }
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def earliestFinishTime(self, landStartTime: List[int], landDuration: List[int],
                           waterStartTime: List[int], waterDuration: List[int]) -> int:

        # Store minimum finishing time
        ans = float('inf')

        n = len(landStartTime)
        m = len(waterStartTime)

        # Try every land ride with every water ride
        for i in range(n):
            for j in range(m):

                # Land -> Water
                land_finish = landStartTime[i] + landDuration[i]

                # Water starts when both conditions are satisfied:
                # land ride finished and water ride opened
                water_start = max(land_finish, waterStartTime[j])

                finish1 = water_start + waterDuration[j]

                # Water -> Land
                water_finish = waterStartTime[j] + waterDuration[j]

                # Land starts when both conditions are satisfied:
                # water ride finished and land ride opened
                land_start = max(water_finish, landStartTime[i])

                finish2 = land_start + landDuration[i]

                # Keep the best answer
                ans = min(ans, finish1, finish2)

        return ans
```

### Go

```go
func earliestFinishTime(landStartTime []int, landDuration []int,
 waterStartTime []int, waterDuration []int) int {

 // Large initial value
 ans := int(1e9)

 n := len(landStartTime)
 m := len(waterStartTime)

 // Try every land-water pair
 for i := 0; i < n; i++ {
  for j := 0; j < m; j++ {

   // Land -> Water

   // Finish time of land ride
   landFinish := landStartTime[i] + landDuration[i]

   // Water ride starts after land ride finishes
   // and after it becomes available
   waterStart := max(landFinish, waterStartTime[j])

   // Final finish time
   finish1 := waterStart + waterDuration[j]

   // Water -> Land

   // Finish time of water ride
   waterFinish := waterStartTime[j] + waterDuration[j]

   // Land ride starts after water ride finishes
   // and after it becomes available
   landStart := max(waterFinish, landStartTime[i])

   // Final finish time
   finish2 := landStart + landDuration[i]

   // Keep minimum answer
   ans = min(ans, min(finish1, finish2))
  }
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
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical across all five implementations.

### Step 1: Initialize the Answer

I start with a very large value because I want to minimize the finishing time.

As I discover better schedules, I keep replacing the answer.

---

### Step 2: Try Every Ride Pair

I iterate through:

* Every land ride
* Every water ride

This ensures that no possible combination is missed.

---

### Step 3: Simulate Land → Water

First, I calculate when the land ride finishes.

Then I determine when the water ride can start.

The water ride can only begin after:

* The land ride has finished
* The water ride is already open

Therefore:

Start Water = max(Land Finish, Water Opening Time)

After that, I add the water duration to get the final completion time.

---

### Step 4: Simulate Water → Land

I repeat the same process in reverse.

First calculate when the water ride finishes.

Then determine when the land ride can begin.

The land ride can only start after:

* The water ride has finished
* The land ride is open

Therefore:

Start Land = max(Water Finish, Land Opening Time)

Then I add the land duration.

---

### Step 5: Update the Minimum Answer

For the current ride pair, I now have:

* Completion time for Land → Water
* Completion time for Water → Land

I take the smaller value.

Then I compare it against the overall answer.

---

### Step 6: Return the Result

After all ride pairs have been processed, the answer contains the earliest possible finishing time.

That value is returned.

---

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

Land Ride 0:

* Start = 2
* Finish = 6

Water Ride 0:

* Opens at 6
* Start = 6
* Finish = 9

Final answer = 9

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

Water first:

* Start = 1
* Finish = 11

Land:

* Already open
* Start = 11
* Finish = 14

Final answer = 14

---

### Example 3

Input

```text
landStartTime = [3]
landDuration = [2]
waterStartTime = [10]
waterDuration = [1]
```

Output

```text
11
```

Explanation

Land:

* Start = 3
* Finish = 5

Wait until water opens.

Water:

* Start = 10
* Finish = 11

Final answer = 11

---

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

Build

```bash
go build solution.go
```

---

## Notes & Optimizations

* The constraints are small enough for brute force.
* Every possible ride combination must be checked.
* Since only two ride categories exist, there are only two valid ride orders.
* No sorting is required.
* No priority queue is required.
* No dynamic programming is required.
* No binary search is required.
* The O(n × m) solution is already optimal for the given constraints.
* The implementation remains simple, readable, and easy to debug.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
