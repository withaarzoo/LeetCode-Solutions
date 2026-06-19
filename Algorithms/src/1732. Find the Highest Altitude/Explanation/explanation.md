# Find the Highest Altitude - LeetCode 1732 Solution

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

LeetCode 1732, Find the Highest Altitude, is a simple array and prefix sum problem.

A biker starts a road trip at altitude `0`. You are given an integer array called `gain`, where each element represents the net altitude change between two consecutive points.

Your task is to determine the highest altitude reached during the entire trip.

### Input

* An integer array `gain`

### Output

* The highest altitude reached at any point during the trip

This problem is commonly categorized under:

* Array
* Prefix Sum
* Simulation
* LeetCode Easy
* Competitive Programming
* Data Structures and Algorithms (DSA)

## Constraints

| Constraint           | Value |
| -------------------- | ----- |
| gain.length          | n     |
| 1 ≤ n ≤ 100          |       |
| -100 ≤ gain[i] ≤ 100 |       |

## Intuition

The first thing I noticed was that the biker always starts at altitude `0`.

Every value in the `gain` array tells me how much the altitude changes when moving to the next point. If I keep adding these values one by one, I can calculate the current altitude at every step.

Instead of storing all altitudes, I only need to track:

1. The current altitude.
2. The highest altitude seen so far.

As soon as the current altitude becomes larger than the maximum altitude, I update the answer.

This makes the solution simple and efficient.

## Approach

I use the following strategy:

1. Start with altitude `0`.
2. Create a variable to store the highest altitude reached.
3. Traverse the `gain` array.
4. Add each gain value to the current altitude.
5. Compare the current altitude with the highest altitude.
6. Update the highest altitude whenever necessary.
7. Return the highest altitude after processing all elements.

This approach avoids storing unnecessary information and works in a single pass.

## Data Structures Used

### Integer Variables

I only use a few integer variables:

* `currentAltitude` → stores the current altitude after each move.
* `maxAltitude` → stores the highest altitude reached so far.

### Array

* `gain` stores the altitude changes provided in the input.

No extra arrays, stacks, queues, hash maps, or other data structures are required.

## Operations & Behavior Summary

The algorithm performs the following actions:

1. Initialize current altitude as `0`.
2. Initialize highest altitude as `0`.
3. Visit each value in the gain array.
4. Update current altitude using the gain value.
5. Check whether the new altitude is the highest seen so far.
6. Update the answer if needed.
7. Continue until all elements are processed.
8. Return the highest altitude.

In simple terms:

* Keep moving.
* Keep updating altitude.
* Remember the highest point reached.

## Complexity

| Metric           | Complexity | Explanation                                                          |
| ---------------- | ---------- | -------------------------------------------------------------------- |
| Time Complexity  | O(n)       | The gain array is traversed exactly once.                            |
| Space Complexity | O(1)       | Only a few variables are used. No extra data structures are created. |

Where:

* `n` = length of the `gain` array

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int largestAltitude(vector<int>& gain) {
        // Current altitude starts at 0
        int currentAltitude = 0;
        
        // Highest altitude seen so far
        int maxAltitude = 0;

        // Process every gain value
        for (int change : gain) {
            // Move to the next point by applying altitude change
            currentAltitude += change;

            // Update highest altitude if current altitude is greater
            maxAltitude = max(maxAltitude, currentAltitude);
        }

        // Return the highest altitude reached
        return maxAltitude;
    }
};
```

### Java

```java
class Solution {
    public int largestAltitude(int[] gain) {
        // Current altitude starts at 0
        int currentAltitude = 0;

        // Highest altitude seen so far
        int maxAltitude = 0;

        // Process every gain value
        for (int change : gain) {
            // Apply altitude change
            currentAltitude += change;

            // Update highest altitude if needed
            maxAltitude = Math.max(maxAltitude, currentAltitude);
        }

        // Return the answer
        return maxAltitude;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} gain
 * @return {number}
 */
var largestAltitude = function(gain) {
    // Current altitude starts at 0
    let currentAltitude = 0;

    // Highest altitude seen so far
    let maxAltitude = 0;

    // Process every gain value
    for (const change of gain) {
        // Apply altitude change
        currentAltitude += change;

        // Update highest altitude if current altitude is larger
        maxAltitude = Math.max(maxAltitude, currentAltitude);
    }

    // Return the highest altitude reached
    return maxAltitude;
};
```

### Python3

```python
class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        # Current altitude starts at 0
        currentAltitude = 0

        # Highest altitude seen so far
        maxAltitude = 0

        # Process every altitude change
        for change in gain:
            # Apply the altitude change
            currentAltitude += change

            # Update highest altitude if needed
            maxAltitude = max(maxAltitude, currentAltitude)

        # Return the answer
        return maxAltitude
```

### Go

```go
func largestAltitude(gain []int) int {
    // Current altitude starts at 0
    currentAltitude := 0

    // Highest altitude seen so far
    maxAltitude := 0

    // Process every gain value
    for _, change := range gain {
        // Apply altitude change
        currentAltitude += change

        // Update highest altitude if current altitude is greater
        if currentAltitude > maxAltitude {
            maxAltitude = currentAltitude
        }
    }

    // Return the highest altitude reached
    return maxAltitude
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same across all five programming languages.

### Step 1: Start at Altitude Zero

The problem states that the biker starts at altitude `0`.

This becomes our initial altitude.

### Step 2: Store the Highest Altitude

Initially, the highest altitude is also `0`.

This is important because all future altitudes could be negative.

Example:

```text
gain = [-4, -3, -2]
```

Altitudes:

```text
0 → -4 → -7 → -9
```

The highest altitude is still `0`.

### Step 3: Process Every Gain Value

Each gain value represents an altitude change.

For every value:

```text
current altitude += gain value
```

This simulates the biker moving to the next point.

### Step 4: Check for a New Maximum

After updating the altitude, compare it with the highest altitude seen so far.

If the current altitude is larger:

```text
highest altitude = current altitude
```

Otherwise, keep the previous maximum.

### Step 5: Continue Until the End

Repeat the process for every element in the array.

Since every gain value is processed exactly once, the algorithm remains efficient.

### Step 6: Return the Result

After processing the complete array, the stored maximum altitude becomes the final answer.

## Examples

### Example 1

**Input**

```text
gain = [-5,1,5,0,-7]
```

**Output**

```text
1
```

**Trace**

```text
Start = 0
0 + (-5) = -5
-5 + 1 = -4
-4 + 5 = 1
1 + 0 = 1
1 + (-7) = -6
```

Altitudes:

```text
[0, -5, -4, 1, 1, -6]
```

Highest altitude:

```text
1
```

---

### Example 2

**Input**

```text
gain = [-4,-3,-2,-1,4,3,2]
```

**Output**

```text
0
```

**Trace**

```text
0
-4
-7
-9
-10
-6
-3
-1
```

Highest altitude:

```text
0
```

---

### Example 3

**Input**

```text
gain = [2,3,-1,4]
```

**Output**

```text
8
```

**Trace**

```text
0
2
5
4
8
```

Highest altitude:

```text
8
```

## How to Use / Run Locally

### C++

Compile:

```bash
g++ main.cpp -o main
```

Run:

```bash
./main
```

### Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

### JavaScript

Run:

```bash
node solution.js
```

### Python3

Run:

```bash
python solution.py
```

or

```bash
python3 solution.py
```

### Go

Run:

```bash
go run main.go
```

Build:

```bash
go build
```

## Notes & Optimizations

### Edge Cases

#### All Negative Gains

```text
gain = [-5,-5,-5]
```

The answer remains:

```text
0
```

because the trip starts at altitude zero.

#### Single Element Array

```text
gain = [10]
```

Highest altitude:

```text
10
```

#### Large Positive Changes

```text
gain = [100,100,100]
```

The algorithm still works correctly.

### Why This Solution Is Optimal

* Single traversal of the array
* Constant extra memory
* No unnecessary storage
* Easy to implement
* Works within all constraints

### Alternative Approach

Another approach would be to build a complete altitude array and then find its maximum value.

However, that requires extra memory.

Tracking the maximum during traversal is more efficient and keeps the space complexity at O(1).

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
