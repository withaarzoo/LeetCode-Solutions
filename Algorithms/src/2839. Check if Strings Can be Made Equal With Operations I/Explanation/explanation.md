# Problem Title

1. Check if Strings Can be Made Equal With Operations I

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

## Problem Summary

We are given two strings `s1` and `s2` of length 4.

We can perform the following operation any number of times:

* Choose two indices `i` and `j` such that `j - i = 2`
* Swap the characters at those indices

We need to return `true` if we can make `s1` equal to `s2`, otherwise return `false`.

## Constraints

```text
s1.length == s2.length == 4
s1 and s2 contain only lowercase English letters
```

## Intuition

I noticed that the allowed swaps are very limited.

* Index `0` can only swap with index `2`
* Index `1` can only swap with index `3`

That means:

* Characters at even positions can only move among even positions
* Characters at odd positions can only move among odd positions

So instead of trying all possible swaps, I just compare:

* The even-indexed characters of both strings
* The odd-indexed characters of both strings

If both groups match after sorting, then I can make the strings equal.

## Approach

1. Extract characters at even indices from both strings.
2. Extract characters at odd indices from both strings.
3. Sort both even groups.
4. Sort both odd groups.
5. Compare the sorted even groups.
6. Compare the sorted odd groups.
7. If both match, return `true`.
8. Otherwise, return `false`.

## Data Structures Used

* Strings
* Character arrays / lists
* Sorting utilities

Since the input size is fixed to length 4, we only use a very small amount of extra memory.

## Operations & Behavior Summary

| Operation | Allowed Indices |
| --------- | --------------- |
| Swap 1    | 0 and 2         |
| Swap 2    | 1 and 3         |

Because of this:

| Position Type | Indices |
| ------------- | ------- |
| Even          | 0, 2    |
| Odd           | 1, 3    |

Characters never move from even positions to odd positions.

## Complexity

* Time Complexity: `O(1)`

  * The string length is always fixed at 4.
  * Sorting and comparisons happen on very small arrays.

* Space Complexity: `O(1)`

  * Only a few extra variables are used.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool canBeEqual(string s1, string s2) {
        string even1 = "", even2 = "";
        string odd1 = "", odd2 = "";

        for (int i = 0; i < 4; i++) {
            if (i % 2 == 0) {
                even1 += s1[i];
                even2 += s2[i];
            } else {
                odd1 += s1[i];
                odd2 += s2[i];
            }
        }

        sort(even1.begin(), even1.end());
        sort(even2.begin(), even2.end());
        sort(odd1.begin(), odd1.end());
        sort(odd2.begin(), odd2.end());

        return even1 == even2 && odd1 == odd2;
    }
};
```

### Java

```java
import java.util.Arrays;

class Solution {
    public boolean canBeEqual(String s1, String s2) {
        char[] even1 = {s1.charAt(0), s1.charAt(2)};
        char[] even2 = {s2.charAt(0), s2.charAt(2)};

        char[] odd1 = {s1.charAt(1), s1.charAt(3)};
        char[] odd2 = {s2.charAt(1), s2.charAt(3)};

        Arrays.sort(even1);
        Arrays.sort(even2);
        Arrays.sort(odd1);
        Arrays.sort(odd2);

        return Arrays.equals(even1, even2) && Arrays.equals(odd1, odd2);
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s1
 * @param {string} s2
 * @return {boolean}
 */
var canBeEqual = function(s1, s2) {
    let even1 = [s1[0], s1[2]].sort().join('');
    let even2 = [s2[0], s2[2]].sort().join('');

    let odd1 = [s1[1], s1[3]].sort().join('');
    let odd2 = [s2[1], s2[3]].sort().join('');

    return even1 === even2 && odd1 === odd2;
};
```

### Python3

```python
class Solution:
    def canBeEqual(self, s1: str, s2: str) -> bool:
        even1 = sorted([s1[0], s1[2]])
        even2 = sorted([s2[0], s2[2]])

        odd1 = sorted([s1[1], s1[3]])
        odd2 = sorted([s2[1], s2[3]])

        return even1 == even2 and odd1 == odd2
```

### Go

```go
import "sort"

func canBeEqual(s1 string, s2 string) bool {
    even1 := []byte{s1[0], s1[2]}
    even2 := []byte{s2[0], s2[2]}

    odd1 := []byte{s1[1], s1[3]}
    odd2 := []byte{s2[1], s2[3]}

    sort.Slice(even1, func(i, j int) bool {
        return even1[i] < even1[j]
    })

    sort.Slice(even2, func(i, j int) bool {
        return even2[i] < even2[j]
    })

    sort.Slice(odd1, func(i, j int) bool {
        return odd1[i] < odd1[j]
    })

    sort.Slice(odd2, func(i, j int) bool {
        return odd2[i] < odd2[j]
    })

    return even1[0] == even2[0] &&
           even1[1] == even2[1] &&
           odd1[0] == odd2[0] &&
           odd1[1] == odd2[1]
}
```

## Step-by-step Detailed Explanation

### C++

* I create separate strings for even and odd indexed characters.
* I loop through all 4 indices.
* If the index is even, I store that character in the even string.
* Otherwise, I store it in the odd string.
* Then I sort both groups.
* Finally, I compare them.

### Java

* I create character arrays for even and odd positions.
* I put indices `0` and `2` into the even array.
* I put indices `1` and `3` into the odd array.
* I sort all arrays.
* If both even arrays and odd arrays are equal, I return `true`.

### JavaScript

* I build small arrays using even and odd indices.
* I sort them and convert them back to strings.
* Then I compare the even parts and odd parts.
* If both are equal, I return `true`.

### Python3

* I use `sorted()` directly on even and odd indexed characters.
* This gives me sorted lists.
* Then I compare both even lists and odd lists.
* If both match, I return `True`.

### Go

* I create byte slices for even and odd positions.
* Then I sort them using `sort.Slice()`.
* Finally, I compare all characters one by one.

## Examples

### Example 1

```text
Input:
s1 = "abcd"
s2 = "cdab"

Output:
true
```

Explanation:

```text
s1 even indices = [a, c]
s2 even indices = [c, a]

After sorting:
[a, c] == [a, c]

s1 odd indices = [b, d]
s2 odd indices = [d, b]

After sorting:
[b, d] == [b, d]
```

So the answer is `true`.

### Example 2

```text
Input:
s1 = "abcd"
s2 = "dacb"

Output:
false
```

Explanation:

```text
s1 even indices = [a, c]
s2 even indices = [d, c]
```

These do not match, so the answer is `false`.

## How to use / Run locally

### C++

```bash
g++ solution.cpp -o solution
./solution
```

### Java

```bash
javac Solution.java
java Solution
```

### JavaScript

```bash
node solution.js
```

### Python3

```bash
python solution.py
```

### Go

```bash
go run solution.go
```

## Notes & Optimizations

* Since the string length is always 4, brute force would also work.
* But comparing even and odd groups is much cleaner.
* This solution is very easy to understand.
* It avoids unnecessary swapping simulations.
* The approach is optimal in both time and space.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
