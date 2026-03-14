# 1415. The k-th Lexicographical String of All Happy Strings of Length n

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

A happy string is defined as a string that:

1. Contains only characters from the set: `['a', 'b', 'c']`
2. No two adjacent characters are the same.

Given two integers `n` and `k`, the task is to return the **k-th lexicographical happy string of length n**.

If there are fewer than `k` happy strings possible, return an empty string.

The strings must be considered in **lexicographical (dictionary) order**.

---

## Constraints

* `1 <= n <= 10`
* `1 <= k <= 100`

---

## Intuition

When I first looked at the problem, I noticed that the characters allowed are only **a, b, c**, and we cannot place the same character next to itself.

This means:

* The first position has **3 possible choices**.
* Every next position has **2 possible choices** (anything except the previous character).

So the total number of happy strings is:

```
3 * 2^(n-1)
```

Since `n` is small (≤10), I realized I could **generate all valid strings using backtracking (DFS)**.

While generating them in lexicographical order, I simply count them. When the counter reaches `k`, that string becomes my answer.

---

## Approach

1. Use **Depth First Search (DFS)** to construct the string character by character.
2. Maintain a list of characters `['a','b','c']`.
3. For every step:

   * Skip the character if it is the same as the previous character.
4. Continue building until the string length becomes `n`.
5. Every valid string increases a counter.
6. When the counter becomes `k`, store the result and stop further recursion.

This ensures we generate strings in **lexicographical order automatically**.

---

## Data Structures Used

* Recursion stack
* String builder / string
* Counter variable

No complex data structures are required because the search space is small.

---

## Operations & Behavior Summary

| Operation          | Purpose                             |
| ------------------ | ----------------------------------- |
| DFS / Backtracking | Generate valid happy strings        |
| Character Check    | Avoid repeating adjacent characters |
| Counter Tracking   | Identify the k-th string            |
| Early Termination  | Stop recursion once answer is found |

---

## Complexity

### Time Complexity

```
O(2^n)
```

Explanation:

* First character has 3 choices
* Remaining positions have 2 choices

Total combinations:

```
3 * 2^(n-1)
```

Because `n <= 10`, this is very small.

### Space Complexity

```
O(n)
```

Used for recursion depth and the temporary string being built.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string result = "";
    int count = 0;

    void dfs(int n, int k, string &curr) {
        if (!result.empty()) return;

        if (curr.size() == n) {
            count++;
            if (count == k) result = curr;
            return;
        }

        for (char c : {'a','b','c'}) {
            if (!curr.empty() && curr.back() == c) continue;

            curr.push_back(c);
            dfs(n, k, curr);
            curr.pop_back();
        }
    }

    string getHappyString(int n, int k) {
        string curr = "";
        dfs(n, k, curr);
        return result;
    }
};
```

### Java

```java
class Solution {

    int count = 0;
    String result = "";

    private void dfs(int n, int k, StringBuilder curr) {

        if (!result.equals("")) return;

        if (curr.length() == n) {
            count++;
            if (count == k) result = curr.toString();
            return;
        }

        char[] chars = {'a','b','c'};

        for (char c : chars) {

            if (curr.length() > 0 && curr.charAt(curr.length()-1) == c)
                continue;

            curr.append(c);
            dfs(n, k, curr);
            curr.deleteCharAt(curr.length()-1);
        }
    }

    public String getHappyString(int n, int k) {
        dfs(n, k, new StringBuilder());
        return result;
    }
}
```

### JavaScript

```javascript
var getHappyString = function(n, k) {

    let result = "";
    let count = 0;

    function dfs(curr){

        if(result !== "") return;

        if(curr.length === n){
            count++;
            if(count === k) result = curr;
            return;
        }

        for(let c of ['a','b','c']){

            if(curr.length > 0 && curr[curr.length-1] === c)
                continue;

            dfs(curr + c);
        }
    }

    dfs("");

    return result;
};
```

### Python3

```python
class Solution:
    def getHappyString(self, n: int, k: int) -> str:
        
        self.count = 0
        self.result = ""

        def dfs(curr):

            if self.result:
                return

            if len(curr) == n:
                self.count += 1
                if self.count == k:
                    self.result = curr
                return

            for c in ['a','b','c']:

                if curr and curr[-1] == c:
                    continue

                dfs(curr + c)

        dfs("")
        return self.result
```

### Go

```go
func getHappyString(n int, k int) string {

    chars := []byte{'a','b','c'}
    count := 0
    result := ""

    var dfs func(curr string)

    dfs = func(curr string){

        if result != "" {
            return
        }

        if len(curr) == n {
            count++
            if count == k {
                result = curr
            }
            return
        }

        for _,c := range chars {

            if len(curr) > 0 && curr[len(curr)-1] == c {
                continue
            }

            dfs(curr + string(c))
        }
    }

    dfs("")
    return result
}
```

---

## Step-by-step Detailed Explanation

1. Start with an empty string.
2. Try adding characters `a`, `b`, and `c` one by one.
3. Before adding a character, check if it is the same as the previous character.
4. If it is the same, skip it.
5. Otherwise, append the character and continue recursion.
6. When the string length becomes `n`, it becomes a valid happy string.
7. Increase a counter for every valid string.
8. When the counter reaches `k`, store the result.
9. Stop recursion early once the result is found.

This ensures that we only explore the minimum required search space.

---

## Examples

### Example 1

Input

```
n = 1
k = 3
```

Output

```
c
```

Explanation

All happy strings:

```
a
b
c
```

The third string is `c`.

---

### Example 2

Input

```
n = 1
k = 4
```

Output

```
""
```

Explanation

Only 3 valid happy strings exist.

---

### Example 3

Input

```
n = 3
k = 9
```

Output

```
cab
```

---

## How to use / Run locally

Example (C++):

```
g++ solution.cpp
./a.out
```

Example (Python):

```
python solution.py
```

Example (JavaScript):

```
node solution.js
```

---

## Notes & Optimizations

* Early stopping avoids generating all combinations.
* Backtracking ensures lexicographical order naturally.
* The search space is small due to the constraints.
* No additional data structures are required.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
