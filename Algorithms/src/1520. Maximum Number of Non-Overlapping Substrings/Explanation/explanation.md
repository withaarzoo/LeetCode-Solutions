# 1520. Maximum Number of Non-Overlapping Substrings

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

You are given a string made of lowercase letters. Your task is to find the largest possible set of non-empty substrings that follow two strict rules.

First, no two substrings are allowed to overlap.  
Second, if a substring contains a character, it must contain every occurrence of that character in the whole string.

Among all ways to pick the maximum number of such substrings, you must return the set whose total length is the smallest. The problem guarantees that this minimum-length answer is unique. You may return the substrings in any order.

This is a classic greedy + interval problem that appears in many competitive programming contests and LeetCode hard-level string questions.

## Constraints

- 1 <= s.length <= 10^5
- s contains only lowercase English letters

## Intuition

When I first read the problem I noticed that every character forces a fixed range: from its first appearance to its last appearance. Any valid substring that uses that character has to cover at least that whole range.

I also realized that two valid ranges can never partially overlap. One must sit completely inside the other or they stay completely separate. That observation immediately suggested a greedy strategy: find all the tightest possible valid intervals, sort them by ending position, and keep the ones that do not overlap.

## Approach

1. Walk through the string once and record the first and last index of every letter.  
2. For each letter that appears, start with the interval [first, last].  
3. Scan everything inside that interval. If any character inside needs an earlier start, discard the interval. Otherwise keep expanding the right end until the interval is closed.  
4. Collect all such closed intervals (at most 26 of them).  
5. Sort the intervals by their right endpoint.  
6. Greedily pick an interval only when its left end is strictly after the right end of the previously chosen interval.  
7. Extract the corresponding substrings and return them.

This produces the maximum number of non-overlapping valid substrings and, among all such maximum sets, the one with the smallest total length.

## Data Structures Used

- Two fixed-size arrays of length 26 to store the first and last occurrence of each letter. Constant space and extremely fast lookups.  
- A list (or vector) of pairs that holds the candidate intervals. Because there are only 26 letters the list stays tiny.  
- A simple integer variable that remembers the ending position of the last accepted interval for the greedy step.

## Operations & Behavior Summary

- Record first and last indices of every character in one linear pass.  
- For each character, attempt to build a closed interval by expanding the right boundary while verifying that no character inside needs a smaller left boundary.  
- Discard any interval that fails the closed-range test.  
- Sort the surviving intervals by ending index.  
- Walk the sorted list and accept an interval only when it starts after the previous accepted interval ends.  
- Convert each accepted interval into its substring and collect the final answer.

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time       | O(n)  | One pass to build first/last arrays, another linear scan while expanding the 26 intervals, and sorting a constant number of intervals. |
| Space      | O(1)  | Only fixed-size arrays of length 26 and a tiny list of intervals are used. The output list is not counted as extra space. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    vector<string> maxNumOfSubstrings(string s) {
        int n = s.size();
        // first and last occurrence of each letter
        vector<int> left(26, -1), right(26, -1);
        for (int i = 0; i < n; ++i) {
            int c = s[i] - 'a';
            if (left[c] == -1) left[c] = i;   // record first time we see it
            right[c] = i;                     // always update last time
        }

        // collect every valid closed interval
        vector<pair<int,int>> intervals;
        for (int i = 0; i < 26; ++i) {
            if (left[i] == -1) continue;      // letter never appeared
            int L = left[i], R = right[i];
            bool valid = true;
            // expand right end while scanning the current range
            for (int j = L; j <= R; ++j) {
                int c = s[j] - 'a';
                if (left[c] < L) {            // needs to start earlier -> not closed
                    valid = false;
                    break;
                }
                R = max(R, right[c]);         // push right end if needed
            }
            if (valid) intervals.emplace_back(L, R);
        }

        // sort by ending position so greedy can pick earliest-ending first
        sort(intervals.begin(), intervals.end(),
             [](const pair<int,int>& a, const pair<int,int>& b) {
                 return a.second < b.second;
             });

        // greedy selection of non-overlapping intervals
        vector<string> ans;
        int lastEnd = -1;
        for (auto [L, R] : intervals) {
            if (L > lastEnd) {                // completely after previous piece
                ans.push_back(s.substr(L, R - L + 1));
                lastEnd = R;
            }
        }
        return ans;
    }
};
```

### Java
```java
class Solution {
    public List<String> maxNumOfSubstrings(String s) {
        int n = s.length();
        // first and last occurrence of each letter
        int[] left = new int[26];
        int[] right = new int[26];
        Arrays.fill(left, -1);
        Arrays.fill(right, -1);
        for (int i = 0; i < n; i++) {
            int c = s.charAt(i) - 'a';
            if (left[c] == -1) left[c] = i;   // record first time we see it
            right[c] = i;                     // always update last time
        }

        // collect every valid closed interval
        List<int[]> intervals = new ArrayList<>();
        for (int i = 0; i < 26; i++) {
            if (left[i] == -1) continue;      // letter never appeared
            int L = left[i], R = right[i];
            boolean valid = true;
            // expand right end while scanning the current range
            for (int j = L; j <= R; j++) {
                int c = s.charAt(j) - 'a';
                if (left[c] < L) {            // needs to start earlier -> not closed
                    valid = false;
                    break;
                }
                R = Math.max(R, right[c]);    // push right end if needed
            }
            if (valid) intervals.add(new int[]{L, R});
        }

        // sort by ending position so greedy can pick earliest-ending first
        intervals.sort((a, b) -> Integer.compare(a[1], b[1]));

        // greedy selection of non-overlapping intervals
        List<String> ans = new ArrayList<>();
        int lastEnd = -1;
        for (int[] iv : intervals) {
            int L = iv[0], R = iv[1];
            if (L > lastEnd) {                // completely after previous piece
                ans.add(s.substring(L, R + 1));
                lastEnd = R;
            }
        }
        return ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {string} s
 * @return {string[]}
 */
var maxNumOfSubstrings = function(s) {
    const n = s.length;
    // first and last occurrence of each letter
    const left = new Array(26).fill(-1);
    const right = new Array(26).fill(-1);
    for (let i = 0; i < n; i++) {
        const c = s.charCodeAt(i) - 97;
        if (left[c] === -1) left[c] = i;   // record first time we see it
        right[c] = i;                      // always update last time
    }

    // collect every valid closed interval
    const intervals = [];
    for (let i = 0; i < 26; i++) {
        if (left[i] === -1) continue;      // letter never appeared
        let L = left[i], R = right[i];
        let valid = true;
        // expand right end while scanning the current range
        for (let j = L; j <= R; j++) {
            const c = s.charCodeAt(j) - 97;
            if (left[c] < L) {             // needs to start earlier -> not closed
                valid = false;
                break;
            }
            R = Math.max(R, right[c]);     // push right end if needed
        }
        if (valid) intervals.push([L, R]);
    }

    // sort by ending position so greedy can pick earliest-ending first
    intervals.sort((a, b) => a[1] - b[1]);

    // greedy selection of non-overlapping intervals
    const ans = [];
    let lastEnd = -1;
    for (const [L, R] of intervals) {
        if (L > lastEnd) {                 // completely after previous piece
            ans.push(s.substring(L, R + 1));
            lastEnd = R;
        }
    }
    return ans;
};
```

### TypeScript
```typescript
function maxNumOfSubstrings(s: string): string[] {
    const n = s.length;
    // first and last occurrence of each letter
    const left: number[] = new Array(26).fill(-1);
    const right: number[] = new Array(26).fill(-1);
    for (let i = 0; i < n; i++) {
        const c = s.charCodeAt(i) - 97;
        if (left[c] === -1) left[c] = i;   // record first time we see it
        right[c] = i;                      // always update last time
    }

    // collect every valid closed interval
    const intervals: [number, number][] = [];
    for (let i = 0; i < 26; i++) {
        if (left[i] === -1) continue;      // letter never appeared
        let L = left[i], R = right[i];
        let valid = true;
        // expand right end while scanning the current range
        for (let j = L; j <= R; j++) {
            const c = s.charCodeAt(j) - 97;
            if (left[c] < L) {             // needs to start earlier -> not closed
                valid = false;
                break;
            }
            R = Math.max(R, right[c]);     // push right end if needed
        }
        if (valid) intervals.push([L, R]);
    }

    // sort by ending position so greedy can pick earliest-ending first
    intervals.sort((a, b) => a[1] - b[1]);

    // greedy selection of non-overlapping intervals
    const ans: string[] = [];
    let lastEnd = -1;
    for (const [L, R] of intervals) {
        if (L > lastEnd) {                 // completely after previous piece
            ans.push(s.substring(L, R + 1));
            lastEnd = R;
        }
    }
    return ans;
}
```

### Python3
```python
class Solution:
    def maxNumOfSubstrings(self, s: str) -> list[str]:
        n = len(s)
        # first and last occurrence of each letter
        left = [-1] * 26
        right = [-1] * 26
        for i, ch in enumerate(s):
            c = ord(ch) - 97
            if left[c] == -1:
                left[c] = i          # record first time we see it
            right[c] = i             # always update last time

        # collect every valid closed interval
        intervals = []
        for i in range(26):
            if left[i] == -1:
                continue             # letter never appeared
            L, R = left[i], right[i]
            valid = True
            # expand right end while scanning the current range
            j = L
            while j <= R:
                c = ord(s[j]) - 97
                if left[c] < L:      # needs to start earlier -> not closed
                    valid = False
                    break
                R = max(R, right[c]) # push right end if needed
                j += 1
            if valid:
                intervals.append((L, R))

        # sort by ending position so greedy can pick earliest-ending first
        intervals.sort(key=lambda x: x[1])

        # greedy selection of non-overlapping intervals
        ans = []
        last_end = -1
        for L, R in intervals:
            if L > last_end:         # completely after previous piece
                ans.append(s[L:R+1])
                last_end = R
        return ans
```

### Go
```go
func maxNumOfSubstrings(s string) []string {
    n := len(s)
    // first and last occurrence of each letter
    left := make([]int, 26)
    right := make([]int, 26)
    for i := range left {
        left[i] = -1
        right[i] = -1
    }
    for i := 0; i < n; i++ {
        c := s[i] - 'a'
        if left[c] == -1 {
            left[c] = i   // record first time we see it
        }
        right[c] = i      // always update last time
    }

    // collect every valid closed interval
    type interval struct{ L, R int }
    var intervals []interval
    for i := 0; i < 26; i++ {
        if left[i] == -1 {
            continue // letter never appeared
        }
        L, R := left[i], right[i]
        valid := true
        // expand right end while scanning the current range
        for j := L; j <= R; j++ {
            c := s[j] - 'a'
            if left[c] < L { // needs to start earlier -> not closed
                valid = false
                break
            }
            if right[c] > R {
                R = right[c] // push right end if needed
            }
        }
        if valid {
            intervals = append(intervals, interval{L, R})
        }
    }

    // sort by ending position so greedy can pick earliest-ending first
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].R < intervals[j].R
    })

    // greedy selection of non-overlapping intervals
    var ans []string
    lastEnd := -1
    for _, iv := range intervals {
        if iv.L > lastEnd { // completely after previous piece
            ans = append(ans, s[iv.L:iv.R+1])
            lastEnd = iv.R
        }
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

All six implementations follow the exact same logic; only the syntax differs.

First I allocate two arrays of size 26 and fill them while scanning the string. left[c] stores the earliest index of character c and right[c] stores the latest. This gives me the absolute span of every letter in constant space.

Next I try to turn each letter into a candidate interval. I set L to its first occurrence and R to its last occurrence. Then I examine every character that currently sits between L and R. If any of those characters has a first occurrence before L, the interval is not closed and I discard it. If a character has a later last occurrence, I push R farther. Because R can grow, the scan continues until the new right boundary is also examined. When the scan finishes without problems I keep the pair (L, R).

After processing all 26 letters I have a short list of valid closed intervals. I sort that list by the right endpoint so the intervals that finish earliest appear first. This ordering is exactly what the classic interval-scheduling greedy needs.

I then walk the sorted list while remembering the right end of the last interval I accepted. For each new interval I check whether its left end is strictly greater than that remembered value. If it is, the two pieces do not overlap, so I accept the interval, extract the substring, and update the remembered right end.

Because valid intervals only nest or stay completely separate, this greedy choice never misses a smaller inner piece that would allow more intervals later. The final list of substrings is the required answer.

Edge cases such as a string of length 1 or a string where every character appears only once are handled automatically: each single character forms its own valid interval and the greedy simply keeps all of them.

## Examples

**Example 1**  
Input: s = "adefaddaccc"  
Output: ["e","f","ccc"]

Trace:  
- e lives only at index 2 → interval [2,2]  
- f lives only at index 3 → interval [3,3]  
- c lives at indices 8-10 → interval [8,10]  
- a forces the larger interval [0,7] which is later discarded by the greedy step  
After sorting by end and picking non-overlapping intervals we obtain the three short pieces.

**Example 2**  
Input: s = "abbaccd"  
Output: ["d","bb","cc"]

Trace:  
- b produces the tight interval [1,2]  
- c produces [4,5]  
- d produces [6,6]  
- a produces the larger interval [0,3] which is skipped  
Greedy keeps the three small non-overlapping pieces.

**Example 3**  
Input: s = "abab"  
Output: ["abab"]

Trace:  
Both a and b force the whole string [0,3]. No smaller closed intervals exist, so the only answer is the entire string.

## How to Use / Run Locally

**C++**  
```bash
g++ -std=c++17 -O2 solution.cpp -o solution
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

Replace the empty function body with the code from the corresponding language section, then feed the input string through a simple main or test harness.

## Notes & Optimizations

- The algorithm is already optimal for the given constraints (n ≤ 10^5). Expanding the 26 intervals never exceeds linear work.  
- Because there are only 26 letters, sorting is effectively constant time.  
- An alternative approach could try dynamic programming over the possible cut points, but it would be slower and more complicated while giving the same answer.  
- Watch out for strings where one character appears only once at the very beginning or the very end; the expansion logic still works correctly.  
- The problem guarantees a unique minimum-length answer, so any correct greedy selection will produce the same set of substrings (order may differ).

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
