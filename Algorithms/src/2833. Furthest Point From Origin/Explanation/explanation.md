# Problem Title

**2833. Furthest Point From Origin**

---

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

You are given a string `moves` consisting of characters `'L'`, `'R'`, and `'_'`.

* `'L'` means move one step to the left (-1)
* `'R'` means move one step to the right (+1)
* `'_'` means you can choose either left or right

You start at position `0`.

Goal: Return the **maximum possible distance from origin** after performing all moves.

---

## Constraints

* `1 <= moves.length <= 50`
* `moves[i] ∈ {'L', 'R', '_'}`

---

## Intuition

I thought about how each character affects my position.

* `'L'` decreases position
* `'R'` increases position
* `'_'` gives flexibility

So instead of deciding `_` one by one, I realized:

To maximize distance, I should push all `_` in the direction that increases my final distance the most.

---

## Approach

1. Count:

   * number of `'L'`
   * number of `'R'`
   * number of `'_'`

2. Compute current position:

   ```
   position = right - left
   ```

3. Use all `_` moves in the same direction:

   ```
   result = abs(position) + blank
   ```

---

## Data Structures Used

* Integer counters (`left`, `right`, `blank`)

No extra data structures are required.

---

## Operations & Behavior Summary

* Single pass traversal
* Count characters
* Simple arithmetic
* Absolute value calculation

---

## Complexity

* **Time Complexity:** O(n)

  * Traverse the string once

* **Space Complexity:** O(1)

  * Constant variables only

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int furthestDistanceFromOrigin(string moves) {
        int left = 0, right = 0, blank = 0;

        for (char c : moves) {
            if (c == 'L') left++;
            else if (c == 'R') right++;
            else blank++;
        }

        int position = right - left;
        return abs(position) + blank;
    }
};
```

### Java

```java
class Solution {
    public int furthestDistanceFromOrigin(String moves) {
        int left = 0, right = 0, blank = 0;

        for (char c : moves.toCharArray()) {
            if (c == 'L') left++;
            else if (c == 'R') right++;
            else blank++;
        }

        int position = right - left;
        return Math.abs(position) + blank;
    }
}
```

### JavaScript

```javascript
var furthestDistanceFromOrigin = function(moves) {
    let left = 0, right = 0, blank = 0;

    for (let c of moves) {
        if (c === 'L') left++;
        else if (c === 'R') right++;
        else blank++;
    }

    let position = right - left;
    return Math.abs(position) + blank;
};
```

### Python3

```python
class Solution:
    def furthestDistanceFromOrigin(self, moves: str) -> int:
        left = moves.count('L')
        right = moves.count('R')
        blank = moves.count('_')

        position = right - left
        return abs(position) + blank
```

### Go

```go
func furthestDistanceFromOrigin(moves string) int {
    left, right, blank := 0, 0, 0

    for _, c := range moves {
        if c == 'L' {
            left++
        } else if c == 'R' {
            right++
        } else {
            blank++
        }
    }

    position := right - left
    if position < 0 {
        position = -position
    }

    return position + blank
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### Step 1: Count all characters

* Loop through the string
* Count how many `'L'`, `'R'`, and `'_'`

### Step 2: Calculate base position

```
position = right - left
```

### Step 3: Maximize distance using `_`

* `_` can go either direction
* To maximize distance, use all `_` in the same direction

```
answer = abs(position) + blank
```

---

## Examples

### Example 1

```
Input: moves = "L_RL__R"
Output: 3
```

### Example 2

```
Input: moves = "_R__LL_"
Output: 5
```

### Example 3

```
Input: moves = "_______"
Output: 7
```

---

## How to use / Run locally

### C++

```
g++ solution.cpp -o solution
./solution
```

### Java

```
javac Solution.java
java Solution
```

### JavaScript

```
node solution.js
```

### Python

```
python3 solution.py
```

### Go

```
go run solution.go
```

---

## Notes & Optimizations

* No need for simulation
* No need for backtracking
* Greedy approach works perfectly
* `_` should always be used to extend distance

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
