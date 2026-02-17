# 401. Binary Watch

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

A binary watch has:

* 4 LEDs to represent hours (0–11)
* 6 LEDs to represent minutes (0–59)

Each LED represents a binary bit (0 or 1).

Given an integer `turnedOn`, which represents how many LEDs are currently ON, return all possible times the watch could represent.

Rules:

* Hour must NOT contain a leading zero.
* Minute must always contain two digits.
* Return result in any order.

---

## Constraints

```bash
0 <= turnedOn <= 10
```

---

## Intuition

When I saw this problem, I did not try to simulate LED combinations directly.

Instead, I thought very simply.

There are only:

* 12 possible hours
* 60 possible minutes

Total combinations = 12 × 60 = 720

720 is very small.

So I thought, why not just try all possible times and count how many 1s are in their binary representation?

If total number of 1 bits equals `turnedOn`, then that time is valid.

Simple and clean.

---

## Approach

Step 1: Loop hour from 0 to 11.

Step 2: Loop minute from 0 to 59.

Step 3: Count number of set bits in hour.

Step 4: Count number of set bits in minute.

Step 5: If total bits equals `turnedOn`, format the time and store it.

Formatting rules:

* Hour printed normally.
* Minute must be two digits. Add leading zero if needed.

Return the result list.

---

## Data Structures Used

* Vector / List / Array to store valid times
* Basic loops
* Built-in bit counting functions (or manual bit counting)

No extra complex data structures are required.

---

## Operations & Behavior Summary

For every possible time:

1. Convert hour to binary.
2. Convert minute to binary.
3. Count total number of 1 bits.
4. Compare with `turnedOn`.
5. If equal → add formatted time to result.

---

## Complexity

### Time Complexity: O(1)

We check at most 720 combinations.

Since 720 is constant, time complexity is constant.

### Space Complexity: O(1)

We only store valid results.
Maximum possible results are bounded.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<string> readBinaryWatch(int turnedOn) {
        vector<string> result;
        
        for (int hour = 0; hour < 12; hour++) {
            for (int minute = 0; minute < 60; minute++) {
                if (__builtin_popcount(hour) + __builtin_popcount(minute) == turnedOn) {
                    string time = to_string(hour) + ":";
                    if (minute < 10) time += "0";
                    time += to_string(minute);
                    result.push_back(time);
                }
            }
        }
        
        return result;
    }
};
```

---

### Java

```java
class Solution {
    public List<String> readBinaryWatch(int turnedOn) {
        List<String> result = new ArrayList<>();
        
        for (int hour = 0; hour < 12; hour++) {
            for (int minute = 0; minute < 60; minute++) {
                if (Integer.bitCount(hour) + Integer.bitCount(minute) == turnedOn) {
                    String time = hour + ":" + (minute < 10 ? "0" + minute : minute);
                    result.add(time);
                }
            }
        }
        
        return result;
    }
}
```

---

### JavaScript

```javascript
var readBinaryWatch = function(turnedOn) {
    const result = [];

    for (let hour = 0; hour < 12; hour++) {
        for (let minute = 0; minute < 60; minute++) {
            const countBits = (n) => n.toString(2).split('0').join('').length;
            
            if (countBits(hour) + countBits(minute) === turnedOn) {
                result.push(hour + ":" + (minute < 10 ? "0" + minute : minute));
            }
        }
    }

    return result;
};
```

---

### Python3

```python
class Solution:
    def readBinaryWatch(self, turnedOn: int) -> List[str]:
        result = []
        
        for hour in range(12):
            for minute in range(60):
                if bin(hour).count('1') + bin(minute).count('1') == turnedOn:
                    result.append(f"{hour}:{minute:02d}")
        
        return result
```

---

### Go

```go
import "strconv"

func readBinaryWatch(turnedOn int) []string {
    var result []string

    for hour := 0; hour < 12; hour++ {
        for minute := 0; minute < 60; minute++ {
            if bitCount(hour) + bitCount(minute) == turnedOn {
                time := strconv.Itoa(hour) + ":"
                if minute < 10 {
                    time += "0"
                }
                time += strconv.Itoa(minute)
                result = append(result, time)
            }
        }
    }

    return result
}

func bitCount(n int) int {
    count := 0
    for n > 0 {
        count += n & 1
        n >>= 1
    }
    return count
}
```

---

## Step-by-step Detailed Explanation

1. Start looping hour from 0 to 11.
2. For each hour, loop minute from 0 to 59.
3. Convert hour into binary and count 1s.
4. Convert minute into binary and count 1s.
5. Add both counts.
6. If total equals `turnedOn`, create formatted string.
7. Push string into result list.
8. After loops finish, return result.

The logic is same in all languages.
Only syntax changes.

---

## Examples

Example 1:

```bash
Input: turnedOn = 1
Output: ["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]
```

Example 2:

```bash
Input: turnedOn = 9
Output: []
```

---

## How to use / Run locally

C++:

```bash
g++ file.cpp
./a.out
```

Java:

```bash
javac Solution.java
java Solution
```

Python:

```bash
python3 file.py
```

Go:

```bash
go run file.go
```

---

## Notes & Optimizations

* Total combinations are very small.
* Brute force is completely acceptable.
* Using built-in bit counting makes code cleaner.
* No need for recursion or backtracking.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
