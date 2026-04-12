# 1320. Minimum Distance to Type a Word Using Two Fingers

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

We are given a word made up of uppercase English letters.

The keyboard layout is arranged in 6 columns:

```text
A B C D E F
G H I J K L
M N O P Q R
S T U V W X
Y Z
```

I need to type the given word using two fingers.

The movement cost between two letters is the Manhattan Distance:

```text
|x1 - x2| + |y1 - y2|
```

Initially, both fingers are free and can start from any letter without any cost.

The goal is to find the minimum total distance required to type the whole word.

---

## Constraints

```text
2 <= word.length <= 300
word consists of uppercase English letters.
```

---

## Intuition

I thought about keeping track of both fingers while typing.

At every character, I have only 2 choices:

1. Use finger 1 to type the current letter
2. Use finger 2 to type the current letter

So, for every position in the word, I need to know:

* Current index in the word
* Current position of finger 1
* Current position of finger 2

Since many states repeat again and again, Dynamic Programming with Memoization is the best choice.

---

## Approach

1. Convert every character into a number from `0` to `25`.
2. Each character has a coordinate:

   * `row = index / 6`
   * `col = index % 6`
3. Use recursion with memoization.
4. Define:

```text
dp(index, finger1, finger2)
```

This means:

```text
Minimum distance needed to type from current index onward
when finger1 is at one letter and finger2 is at another letter.
```

1. At every step:

   * Move finger 1 to current character
   * Move finger 2 to current character
2. Take the minimum result.
3. Use `26` as a special state meaning the finger has not been placed yet.

---

## Data Structures Used

* 3D DP Array
* HashMap / Map for memoization
* Recursive Function
* Helper Function for distance calculation

---

## Operations & Behavior Summary

| Operation            | Description                          |
| -------------------- | ------------------------------------ |
| Character to Index   | Convert `A-Z` into `0-25`            |
| Coordinate Mapping   | `row = index / 6`, `col = index % 6` |
| Distance Calculation | Manhattan Distance                   |
| DP State             | `(index, finger1, finger2)`          |
| Transition 1         | Move finger 1                        |
| Transition 2         | Move finger 2                        |
| Base Case            | If all letters are typed, return `0` |

---

## Complexity

* Time Complexity: `O(n * 27 * 27)`

  * `n` is the length of the word.
  * There are at most `27 * 27` finger combinations for every character.

* Space Complexity: `O(n * 27 * 27)`

  * Used for memoization.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int getDist(int a, int b) {
        if (a == 26 || b == 26) return 0;

        int row1 = a / 6, col1 = a % 6;
        int row2 = b / 6, col2 = b % 6;

        return abs(row1 - row2) + abs(col1 - col2);
    }

    int dp[301][27][27];

    int solve(int idx, int f1, int f2, string &word) {
        if (idx == word.size()) return 0;

        if (dp[idx][f1][f2] != -1) return dp[idx][f1][f2];

        int cur = word[idx] - 'A';

        int useFinger1 = getDist(f1, cur) + solve(idx + 1, cur, f2, word);
        int useFinger2 = getDist(f2, cur) + solve(idx + 1, f1, cur, word);

        return dp[idx][f1][f2] = min(useFinger1, useFinger2);
    }

    int minimumDistance(string word) {
        memset(dp, -1, sizeof(dp));
        return solve(0, 26, 26, word);
    }
};
```

### Java

```java
class Solution {
    private int[][][] dp;

    private int getDist(int a, int b) {
        if (a == 26 || b == 26) return 0;

        int row1 = a / 6, col1 = a % 6;
        int row2 = b / 6, col2 = b % 6;

        return Math.abs(row1 - row2) + Math.abs(col1 - col2);
    }

    private int solve(int idx, int f1, int f2, String word) {
        if (idx == word.length()) return 0;

        if (dp[idx][f1][f2] != -1) return dp[idx][f1][f2];

        int cur = word.charAt(idx) - 'A';

        int useFinger1 = getDist(f1, cur) + solve(idx + 1, cur, f2, word);
        int useFinger2 = getDist(f2, cur) + solve(idx + 1, f1, cur, word);

        return dp[idx][f1][f2] = Math.min(useFinger1, useFinger2);
    }

    public int minimumDistance(String word) {
        dp = new int[word.length()][27][27];

        for (int i = 0; i < word.length(); i++) {
            for (int j = 0; j < 27; j++) {
                for (int k = 0; k < 27; k++) {
                    dp[i][j][k] = -1;
                }
            }
        }

        return solve(0, 26, 26, word);
    }
}
```

### JavaScript

```javascript
var minimumDistance = function(word) {
    const memo = new Map();

    function getDist(a, b) {
        if (a === 26 || b === 26) return 0;

        const row1 = Math.floor(a / 6);
        const col1 = a % 6;
        const row2 = Math.floor(b / 6);
        const col2 = b % 6;

        return Math.abs(row1 - row2) + Math.abs(col1 - col2);
    }

    function solve(idx, f1, f2) {
        if (idx === word.length) return 0;

        const key = `${idx},${f1},${f2}`;

        if (memo.has(key)) return memo.get(key);

        const cur = word.charCodeAt(idx) - 65;

        const useFinger1 = getDist(f1, cur) + solve(idx + 1, cur, f2);
        const useFinger2 = getDist(f2, cur) + solve(idx + 1, f1, cur);

        const ans = Math.min(useFinger1, useFinger2);
        memo.set(key, ans);

        return ans;
    }

    return solve(0, 26, 26);
};
```

### Python3

```python
class Solution:
    def minimumDistance(self, word: str) -> int:
        from functools import lru_cache

        def get_dist(a, b):
            if a == 26 or b == 26:
                return 0

            row1, col1 = divmod(a, 6)
            row2, col2 = divmod(b, 6)

            return abs(row1 - row2) + abs(col1 - col2)

        @lru_cache(None)
        def solve(idx, f1, f2):
            if idx == len(word):
                return 0

            cur = ord(word[idx]) - ord('A')

            use_finger1 = get_dist(f1, cur) + solve(idx + 1, cur, f2)
            use_finger2 = get_dist(f2, cur) + solve(idx + 1, f1, cur)

            return min(use_finger1, use_finger2)

        return solve(0, 26, 26)
```

### Go

```go
func minimumDistance(word string) int {
    memo := make(map[[3]int]int)

    getDist := func(a, b int) int {
        if a == 26 || b == 26 {
            return 0
        }

        row1, col1 := a/6, a%6
        row2, col2 := b/6, b%6

        if row1 < row2 {
            row1, row2 = row2, row1
        }

        if col1 < col2 {
            col1, col2 = col2, col1
        }

        return (row1 - row2) + (col1 - col2)
    }

    var solve func(int, int, int) int

    solve = func(idx, f1, f2 int) int {
        if idx == len(word) {
            return 0
        }

        key := [3]int{idx, f1, f2}

        if val, exists := memo[key]; exists {
            return val
        }

        cur := int(word[idx] - 'A')

        useFinger1 := getDist(f1, cur) + solve(idx+1, cur, f2)
        useFinger2 := getDist(f2, cur) + solve(idx+1, f1, cur)

        ans := useFinger1
        if useFinger2 < ans {
            ans = useFinger2
        }

        memo[key] = ans
        return ans
    }

    return solve(0, 26, 26)
}
```

---

## Step-by-step Detailed Explanation

### C++

* `getDist(a, b)` calculates Manhattan Distance.
* If a finger is not placed yet (`26`), cost is `0`.
* `dp[idx][f1][f2]` stores the minimum answer.
* At every index:

  * Move finger 1
  * Move finger 2
* Store the smaller answer.

### Java

* `dp[idx][f1][f2]` memoizes repeated states.
* The recursive function explores both choices.
* Distance is calculated using row and column values.
* The minimum result is returned.

### JavaScript

* A `Map` is used for memoization.
* The key format is:

```text
index,finger1,finger2
```

* Recursion tries both finger movements.
* Minimum value is stored and reused.

### Python3

* `@lru_cache(None)` is used for memoization.
* `divmod(index, 6)` gives row and column.
* The recursive function checks both finger options.
* It returns the minimum total distance.

### Go

* Uses a map with a 3-integer array key.
* Recursive function checks both possibilities.
* Distance calculation is done separately.
* Memoization prevents repeated work.

---

## Examples

### Example 1

```text
Input: word = "CAKE"
Output: 3
```

Explanation:

```text
Finger 1 -> C -> A
Finger 2 -> K -> E
Total distance = 3
```

### Example 2

```text
Input: word = "HAPPY"
Output: 6
```

---

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

---

## Notes & Optimizations

* Using DP avoids recalculating repeated states.
* There are only 27 possible positions for each finger.
* `26` is used to represent an unused finger.
* Memoization makes the solution fast enough for length up to `300`.
* Recursive DP is easier to write and understand than iterative DP here.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
