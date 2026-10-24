# Count Days Spent Together - LeetCode 2409 Solution

## Table of Contents
- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
  - [C++](#c)
  - [Java](#java)
  - [JavaScript](#javascript)
  - [TypeScript](#typescript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

Alice and Bob are both traveling to Rome for separate business meetings. You are given four strings: arriveAlice, leaveAlice, arriveBob, and leaveBob. Each string is in the format "MM-DD" and represents a date in the same non-leap year.

Alice stays in Rome from arriveAlice to leaveAlice (both days inclusive). Bob stays from arriveBob to leaveBob (also inclusive). Your task is to count the total number of days both of them are present in Rome at the same time.

This is a classic date interval overlap problem that appears frequently in coding interviews and competitive programming contests.

## Constraints

- All dates are provided in the format "MM-DD"
- Alice and Bob’s arrival dates are earlier than or equal to their leaving dates
- The given dates are valid dates of a non-leap year
- Number of days per month: [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]

## Intuition

The key observation is that we do not need a full calendar library. Since everything happens in the same non-leap year, we can convert every "MM-DD" date into a single integer that represents the day of the year (1 to 365).  

Once both Alice and Bob have a start day number and an end day number, the problem reduces to finding the overlap between two integer intervals. The number of overlapping days is simply the length of the intersection of those two intervals.

## Approach

1. Create a fixed array that stores the number of days in each month.
2. Write a small helper that takes a date string, extracts the month and day, and converts it into a day-of-year number by adding the lengths of all previous months.
3. Convert the four given dates into day-of-year numbers.
4. The first day both are present is the maximum of the two start days.
5. The last day both are present is the minimum of the two end days.
6. If the start of the overlap is less than or equal to the end of the overlap, return (end - start + 1). Otherwise return 0.

This approach is clean, constant-time, and works for every valid test case.

## Data Structures Used

- Fixed integer array of size 12 to store the number of days in each month.  
  Chosen because the month lengths never change and looking them up is O(1).

No dynamic data structures (lists, maps, sets, etc.) are required.

## Operations & Behavior Summary

- Convert each of the four date strings into a day-of-year integer.
- Compute the latest arrival day (max of the two starts).
- Compute the earliest departure day (min of the two ends).
- If the latest arrival is still on or before the earliest departure, count the inclusive number of days between them.
- Otherwise the two stays do not overlap and the answer is zero.

## Complexity

| Type              | Complexity | Explanation                                      |
|-------------------|------------|--------------------------------------------------|
| Time Complexity   | O(1)       | Only a constant number of operations on fixed-length strings |
| Space Complexity  | O(1)       | Only a fixed-size array of 12 integers is used   |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int countDaysTogether(string arriveAlice, string leaveAlice, string arriveBob, string leaveBob) {
        int days[] = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
        auto toDay = [&](const string& s) {
            int m = (s[0] - '0') * 10 + (s[1] - '0');
            int d = (s[3] - '0') * 10 + (s[4] - '0');
            int res = d;
            for (int i = 0; i < m - 1; i++) res += days[i];
            return res;
        };
        int aStart = toDay(arriveAlice), aEnd = toDay(leaveAlice);
        int bStart = toDay(arriveBob), bEnd = toDay(leaveBob);
        int start = max(aStart, bStart);
        int end = min(aEnd, bEnd);
        return start <= end ? end - start + 1 : 0;
    }
};
```

### Java
```java
class Solution {
    public int countDaysTogether(String arriveAlice, String leaveAlice, String arriveBob, String leaveBob) {
        int[] days = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
        int aStart = toDay(arriveAlice, days);
        int aEnd = toDay(leaveAlice, days);
        int bStart = toDay(arriveBob, days);
        int bEnd = toDay(leaveBob, days);
        int start = Math.max(aStart, bStart);
        int end = Math.min(aEnd, bEnd);
        return start <= end ? end - start + 1 : 0;
    }
    private int toDay(String s, int[] days) {
        int m = (s.charAt(0) - '0') * 10 + (s.charAt(1) - '0');
        int d = (s.charAt(3) - '0') * 10 + (s.charAt(4) - '0');
        int res = d;
        for (int i = 0; i < m - 1; i++) res += days[i];
        return res;
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} arriveAlice
 * @param {string} leaveAlice
 * @param {string} arriveBob
 * @param {string} leaveBob
 * @return {number}
 */
var countDaysTogether = function(arriveAlice, leaveAlice, arriveBob, leaveBob) {
    const days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
    const toDay = (s) => {
        const m = (s.charCodeAt(0) - 48) * 10 + (s.charCodeAt(1) - 48);
        const d = (s.charCodeAt(3) - 48) * 10 + (s.charCodeAt(4) - 48);
        let res = d;
        for (let i = 0; i < m - 1; i++) res += days[i];
        return res;
    };
    const aStart = toDay(arriveAlice), aEnd = toDay(leaveAlice);
    const bStart = toDay(arriveBob), bEnd = toDay(leaveBob);
    const start = Math.max(aStart, bStart);
    const end = Math.min(aEnd, bEnd);
    return start <= end ? end - start + 1 : 0;
};
```

### TypeScript
```typescript
function countDaysTogether(arriveAlice: string, leaveAlice: string, arriveBob: string, leaveBob: string): number {
    const days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
    const toDay = (s: string): number => {
        const m = (s.charCodeAt(0) - 48) * 10 + (s.charCodeAt(1) - 48);
        const d = (s.charCodeAt(3) - 48) * 10 + (s.charCodeAt(4) - 48);
        let res = d;
        for (let i = 0; i < m - 1; i++) res += days[i];
        return res;
    };
    const aStart = toDay(arriveAlice), aEnd = toDay(leaveAlice);
    const bStart = toDay(arriveBob), bEnd = toDay(leaveBob);
    const start = Math.max(aStart, bStart);
    const end = Math.min(aEnd, bEnd);
    return start <= end ? end - start + 1 : 0;
};
```

### Python3
```python
class Solution:
    def countDaysTogether(self, arriveAlice: str, leaveAlice: str, arriveBob: str, leaveBob: str) -> int:
        days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
        def toDay(s):
            m = (ord(s[0]) - 48) * 10 + (ord(s[1]) - 48)
            d = (ord(s[3]) - 48) * 10 + (ord(s[4]) - 48)
            res = d
            for i in range(m - 1):
                res += days[i]
            return res
        aStart, aEnd = toDay(arriveAlice), toDay(leaveAlice)
        bStart, bEnd = toDay(arriveBob), toDay(leaveBob)
        start = max(aStart, bStart)
        end = min(aEnd, bEnd)
        return end - start + 1 if start <= end else 0
```

### Go
```go
func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
    days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    toDay := func(s string) int {
        m := int(s[0]-'0')*10 + int(s[1]-'0')
        d := int(s[3]-'0')*10 + int(s[4]-'0')
        res := d
        for i := 0; i < m-1; i++ {
            res += days[i]
        }
        return res
    }
    aStart, aEnd := toDay(arriveAlice), toDay(leaveAlice)
    bStart, bEnd := toDay(arriveBob), toDay(leaveBob)
    start := max(aStart, bStart)
    end := min(aEnd, bEnd)
    if start <= end {
        return end - start + 1
    }
    return 0
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

All six solutions follow the exact same logic. The only differences are language syntax.

First we declare a constant array that holds the length of every month. Because the year is never a leap year, February is hard-coded as 28.

Next we write a helper that converts a five-character string "MM-DD" into a day-of-year number. We read the characters directly by index (this avoids any off-by-one errors that can appear with substring or slice). Month is taken from positions 0 and 1, day from positions 3 and 4. We start with the day value and then add the lengths of all months that come before the current month.

We call this helper four times to obtain Alice’s start, Alice’s end, Bob’s start and Bob’s end.

The overlapping interval begins at the later of the two start days and ends at the earlier of the two end days. If that interval is valid (start ≤ end) we return its length; otherwise we return 0.

This conversion-plus-interval-overlap pattern works uniformly across C++, Java, JavaScript, TypeScript, Python3 and Go. Edge cases such as no overlap, complete containment of one interval inside the other, and same-day overlap are all handled correctly by the same max / min arithmetic.

## Examples

**Example 1**  
Input: arriveAlice = "08-15", leaveAlice = "08-18", arriveBob = "08-16", leaveBob = "08-19"  
Output: 3  

Trace:  
- Alice: day 227 to day 230  
- Bob: day 228 to day 231  
- Overlap starts at 228 and ends at 230 → 3 days

**Example 2**  
Input: arriveAlice = "10-01", leaveAlice = "10-31", arriveBob = "11-01", leaveBob = "12-31"  
Output: 0  

Trace:  
- Alice ends on day 304  
- Bob starts on day 305  
- No overlap exists, so the answer is 0

**Example 3**  
Input: arriveAlice = "01-01", leaveAlice = "12-31", arriveBob = "06-01", leaveBob = "06-30"  
Output: 30  

Trace: Bob’s entire stay is completely inside Alice’s stay, so the answer is simply the length of Bob’s interval.

## How to Use / Run Locally

**C++**  
```bash
g++ -std=c++17 solution.cpp -o solution
./solution
```

**Java**  
```bash
javac Solution.java
java Solution
```

**JavaScript**  
```bash
node solution.js
```

**TypeScript**  
```bash
tsc solution.ts
node solution.js
```

**Python3**  
```bash
python3 solution.py
```

**Go**  
```bash
go run solution.go
```

Replace the main function or the LeetCode driver code with your own test cases to experiment.

## Notes & Optimizations

- The solution is already optimal in both time and space. No further asymptotic improvement is possible.
- Using character arithmetic instead of library parsing functions avoids potential index errors and keeps the code consistent across languages.
- If the problem ever allowed leap years, only the February entry in the days array would need to change.
- An alternative approach of comparing the original strings lexicographically works only because the format is zero-padded; converting to day numbers is clearer and safer.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)