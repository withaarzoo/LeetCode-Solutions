# Problem Title

**2147. Number of Ways to Divide a Long Corridor**

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

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given a corridor represented as a string consisting of:

* `S` → Seat
* `P` → Plant

The corridor already has a divider at the extreme left and right.

My task is to count how many **different ways** I can add extra dividers such that:

* Each resulting section contains **exactly two seats**
* Each section may contain **any number of plants**

If it is impossible, I must return `0`.
The answer can be very large, so I return it modulo `10^9 + 7`.

---

## Constraints

* `1 <= corridor.length <= 100000`
* `corridor[i]` is either `'S'` or `'P'`

---

## Intuition

When I looked at the problem carefully, I realized one very strict rule:

> Every section must contain **exactly two seats**.

So I stopped focusing on plants and started focusing only on **where the seats are placed**.

My thinking became simple:

* If the total number of seats is odd or zero → impossible
* Every pair of seats forms one section
* Between two seat-pairs, I must place exactly one divider

The only freedom I have is **where to place that divider between two sections**.

---

## Approach

1. Traverse the corridor and store the **indices of all seats (`S`)**
2. If total seats are `0` or odd → return `0`
3. Treat every **two seats as one section**
4. Between section `i` and section `i+1`, count how many positions exist where a divider can be placed
5. Multiply all such choices together
6. Take modulo `10^9 + 7`

---

## Data Structures Used

* Simple list / array to store **seat indices**
* Integer variables for counting and multiplication

No complex data structures are required.

---

## Operations & Behavior Summary

* Scan corridor once
* Track positions of `S`
* Multiply gaps between seat-pairs
* Ensure valid seat count

The solution is greedy and mathematical in nature.

---

## Complexity

* **Time Complexity:** `O(n)`

  * I scan the corridor only once

* **Space Complexity:** `O(1)` (ignoring output)

  * Only a few variables and seat positions are stored

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int numberOfWays(string corridor) {
        const int MOD = 1e9 + 7;
        vector<int> seats;

        for (int i = 0; i < corridor.size(); i++) {
            if (corridor[i] == 'S') {
                seats.push_back(i);
            }
        }

        if (seats.size() == 0 || seats.size() % 2 != 0) return 0;

        long long ways = 1;
        for (int i = 2; i < seats.size(); i += 2) {
            ways = (ways * (seats[i] - seats[i - 1])) % MOD;
        }

        return ways;
    }
};
```

---

### Java

```java
class Solution {
    public int numberOfWays(String corridor) {
        final int MOD = 1_000_000_007;
        List<Integer> seats = new ArrayList<>();

        for (int i = 0; i < corridor.length(); i++) {
            if (corridor.charAt(i) == 'S') {
                seats.add(i);
            }
        }

        if (seats.size() == 0 || seats.size() % 2 != 0) return 0;

        long ways = 1;
        for (int i = 2; i < seats.size(); i += 2) {
            ways = (ways * (seats.get(i) - seats.get(i - 1))) % MOD;
        }

        return (int) ways;
    }
}
```

---

### JavaScript

```javascript
var numberOfWays = function(corridor) {
    const MOD = 1_000_000_007;
    const seats = [];

    for (let i = 0; i < corridor.length; i++) {
        if (corridor[i] === 'S') seats.push(i);
    }

    if (seats.length === 0 || seats.length % 2 !== 0) return 0;

    let ways = 1;
    for (let i = 2; i < seats.length; i += 2) {
        ways = (ways * (seats[i] - seats[i - 1])) % MOD;
    }

    return ways;
};
```

---

### Python3

```python
class Solution:
    def numberOfWays(self, corridor: str) -> int:
        MOD = 10**9 + 7
        seats = []

        for i, c in enumerate(corridor):
            if c == 'S':
                seats.append(i)

        if len(seats) == 0 or len(seats) % 2 != 0:
            return 0

        ways = 1
        for i in range(2, len(seats), 2):
            ways = (ways * (seats[i] - seats[i - 1])) % MOD

        return ways
```

---

### Go

```go
func numberOfWays(corridor string) int {
    const mod = 1000000007
    seats := []int{}

    for i := 0; i < len(corridor); i++ {
        if corridor[i] == 'S' {
            seats = append(seats, i)
        }
    }

    if len(seats) == 0 || len(seats)%2 != 0 {
        return 0
    }

    ways := 1
    for i := 2; i < len(seats); i += 2 {
        ways = (ways * (seats[i] - seats[i-1])) % mod
    }

    return ways
}
```

---

## Step-by-step Detailed Explanation

1. I scan the corridor and note where each seat appears
2. I verify whether the number of seats allows valid sections
3. I group seats into pairs
4. Between every pair-group, I calculate valid divider positions
5. I multiply all choices to get the final answer

---

## Examples

**Input:** `"SSPPSPS"`

Seat positions: `[0, 1, 4, 6]`

Possible divider positions between index `1` and `4`:

* Positions = `4 - 1 = 3`

**Output:** `3`

---

## How to use / Run locally

1. Copy the code for your preferred language
2. Paste it into the LeetCode editor or your local IDE
3. Run with sample inputs
4. Verify outputs

---

## Notes & Optimizations

* Plants are never counted directly
* Only seat indices matter
* No nested loops → highly efficient
* Works perfectly for very large inputs

---

## Author

* **Md Aarzoo Islam**
  [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
