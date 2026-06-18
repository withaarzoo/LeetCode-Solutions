# 1344. Angle Between Hands of a Clock

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

Given the current hour and minute of an analog clock, the goal is to find the smaller angle formed between the hour hand and the minute hand.

This is a classic mathematics and geometry problem frequently asked in coding interviews and competitive programming contests.

### Input

* `hour` → Current hour value
* `minutes` → Current minute value

### Output

* Return the smallest angle between the hour hand and minute hand.

The answer can be a decimal value because the hour hand continuously moves as minutes pass.

---

## Constraints

| Constraint   | Value                |
| ------------ | -------------------- |
| Hour Range   | `1 <= hour <= 12`    |
| Minute Range | `0 <= minutes <= 59` |

---

## Intuition

The first thing I noticed is that both clock hands move at a fixed speed.

The minute hand is straightforward because every minute advances it by the same amount.

The hour hand is slightly different. It does not stay fixed at an hour mark. As the minutes increase, the hour hand gradually moves toward the next number.

Once I know the exact position of both hands, finding the angle becomes simple. I just calculate their difference and return the smaller angle formed on the clock.

This makes the problem more about geometry than simulation.

---

## Approach

I solve the problem in four simple steps:

### Step 1: Find the minute hand angle

A clock contains 360 degrees and 60 minutes.

So each minute contributes:

```text
360 / 60 = 6 degrees
```

Minute hand angle:

```text
minutes × 6
```

### Step 2: Find the hour hand angle

A clock contains 12 hours.

Each hour contributes:

```text
360 / 12 = 30 degrees
```

The hour hand also moves continuously.

Every minute contributes:

```text
30 / 60 = 0.5 degrees
```

Hour hand angle:

```text
(hour % 12) × 30 + minutes × 0.5
```

### Step 3: Calculate the difference

```text
abs(hourAngle - minuteAngle)
```

### Step 4: Return the smaller angle

A circle always creates two possible angles.

```text
difference
360 - difference
```

Return:

```text
min(difference, 360 - difference)
```

---

## Data Structures Used

No data structures are required.

The solution only uses a few variables:

| Variable      | Purpose                     |
| ------------- | --------------------------- |
| `hourAngle`   | Stores hour hand position   |
| `minuteAngle` | Stores minute hand position |
| `difference`  | Stores angle difference     |

Since no arrays, maps, stacks, queues, or trees are used, memory usage stays constant.

---

## Operations & Behavior Summary

1. Convert hour 12 into 0 using modulo.
2. Calculate minute hand angle.
3. Calculate hour hand angle.
4. Find absolute difference.
5. Compute the opposite angle using `360 - difference`.
6. Return the smaller angle.
7. Finish.

The algorithm performs only a handful of arithmetic operations.

---

## Complexity

| Complexity       | Value  | Explanation                                   |
| ---------------- | ------ | --------------------------------------------- |
| Time Complexity  | `O(1)` | Only constant-time calculations are performed |
| Space Complexity | `O(1)` | No extra data structures are used             |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    double angleClock(int hour, int minutes) {
        // Convert 12 to 0 because both point to the same position
        hour %= 12;

        // Minute hand moves 6 degrees per minute
        double minuteAngle = minutes * 6.0;

        // Hour hand moves 30 degrees per hour
        // and 0.5 degrees per minute
        double hourAngle = hour * 30.0 + minutes * 0.5;

        // Find the angle difference between both hands
        double diff = abs(hourAngle - minuteAngle);

        // Return the smaller angle
        return min(diff, 360.0 - diff);
    }
};
```

### Java

```java
class Solution {
    public double angleClock(int hour, int minutes) {
        // Convert 12 to 0 because both point to the same position
        hour %= 12;

        // Minute hand moves 6 degrees per minute
        double minuteAngle = minutes * 6.0;

        // Hour hand moves 30 degrees per hour
        // and 0.5 degrees per minute
        double hourAngle = hour * 30.0 + minutes * 0.5;

        // Find the difference between the two angles
        double diff = Math.abs(hourAngle - minuteAngle);

        // Return the smaller angle
        return Math.min(diff, 360.0 - diff);
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} hour
 * @param {number} minutes
 * @return {number}
 */
var angleClock = function(hour, minutes) {
    // Convert 12 to 0 because both point to the same position
    hour %= 12;

    // Minute hand moves 6 degrees per minute
    const minuteAngle = minutes * 6;

    // Hour hand moves 30 degrees per hour
    // and 0.5 degrees per minute
    const hourAngle = hour * 30 + minutes * 0.5;

    // Find the difference between both angles
    const diff = Math.abs(hourAngle - minuteAngle);

    // Return the smaller angle
    return Math.min(diff, 360 - diff);
};
```

### Python3

```python
class Solution:
    def angleClock(self, hour: int, minutes: int) -> float:
        # Convert 12 to 0 because both point to the same position
        hour %= 12

        # Minute hand moves 6 degrees per minute
        minute_angle = minutes * 6.0

        # Hour hand moves 30 degrees per hour
        # and 0.5 degrees per minute
        hour_angle = hour * 30.0 + minutes * 0.5

        # Find the difference between both angles
        diff = abs(hour_angle - minute_angle)

        # Return the smaller angle
        return min(diff, 360.0 - diff)
```

### Go

```go
import "math"

func angleClock(hour int, minutes int) float64 {
    // Convert 12 to 0 because both point to the same position
    hour %= 12

    // Minute hand moves 6 degrees per minute
    minuteAngle := float64(minutes) * 6.0

    // Hour hand moves 30 degrees per hour
    // and 0.5 degrees per minute
    hourAngle := float64(hour)*30.0 + float64(minutes)*0.5

    // Find the difference between both angles
    diff := math.Abs(hourAngle - minuteAngle)

    // Return the smaller angle
    return math.Min(diff, 360.0-diff)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains identical across all five languages.

### Calculating the Minute Hand

The minute hand completes one full rotation every 60 minutes.

Since a full circle contains 360 degrees:

```text
360 ÷ 60 = 6
```

Each minute contributes 6 degrees.

Examples:

```text
15 minutes → 90°
30 minutes → 180°
45 minutes → 270°
```

---

### Calculating the Hour Hand

The hour hand completes one full rotation every 12 hours.

Since a full circle contains 360 degrees:

```text
360 ÷ 12 = 30
```

Each hour contributes 30 degrees.

However, the hour hand also moves between hour marks.

Every minute contributes:

```text
30 ÷ 60 = 0.5°
```

Therefore:

```text
hourAngle =
(hour × 30)
+
(minutes × 0.5)
```

This gives the exact position of the hour hand.

---

### Why Use `hour % 12`?

On a clock:

```text
12 = 0
```

Both point to the same location.

Examples:

```text
12:00 → 0°
1:00 → 30°
2:00 → 60°
```

Using modulo keeps calculations consistent.

---

### Finding the Angle Difference

Once both hand positions are known:

```text
difference =
abs(hourAngle - minuteAngle)
```

This gives one angle between the hands.

---

### Returning the Smaller Angle

Every pair of clock hands forms two angles.

Example:

```text
75°
285°
```

The problem specifically asks for the smaller angle.

Therefore:

```text
min(difference, 360 - difference)
```

always gives the correct answer.

---

### Language Notes

#### C++

Uses:

* `abs()`
* `min()`

#### Java

Uses:

* `Math.abs()`
* `Math.min()`

#### JavaScript

Uses:

* `Math.abs()`
* `Math.min()`

#### Python3

Uses:

* `abs()`
* `min()`

#### Go

Uses:

* `math.Abs()`
* `math.Min()`

The underlying algorithm remains exactly the same.

---

## Examples

### Example 1

Input

```text
hour = 12
minutes = 30
```

Calculation

```text
Minute Angle = 30 × 6 = 180°

Hour Angle =
0 × 30 + 30 × 0.5
= 15°
```

Difference

```text
|180 - 15| = 165°
```

Output

```text
165
```

---

### Example 2

Input

```text
hour = 3
minutes = 30
```

Calculation

```text
Minute Angle = 180°

Hour Angle =
3 × 30 + 30 × 0.5
= 105°
```

Difference

```text
75°
```

Output

```text
75
```

---

### Example 3

Input

```text
hour = 3
minutes = 15
```

Calculation

```text
Minute Angle = 90°

Hour Angle =
3 × 30 + 15 × 0.5
= 97.5°
```

Difference

```text
7.5°
```

Output

```text
7.5
```

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

* This is already the optimal solution.
* No iteration is required.
* No extra memory is required.
* Floating-point calculations are necessary because angles can contain decimal values.
* Remember that the hour hand moves continuously and does not stay fixed at hour marks.
* Returning the smaller angle is important because two angles always exist between clock hands.
* This problem is commonly categorized under:

  * Clock Angle Problem
  * Geometry
  * Mathematics
  * LeetCode Math Problems
  * Interview Preparation
  * Competitive Programming

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
