# 3130. Find All Possible Stable Binary Arrays II

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

# Problem Summary

We are given three integers:

* `zero` → number of 0s
* `one` → number of 1s
* `limit` → maximum number of consecutive equal values allowed

A binary array is called **stable** if:

1. It contains exactly `zero` number of `0`s
2. It contains exactly `one` number of `1`s
3. No subarray longer than `limit` contains only `0`s or only `1`s

In simple terms:

* We cannot have more than `limit` consecutive `0`s
* We cannot have more than `limit` consecutive `1`s

Our goal is to count how many such valid arrays exist.

Because the answer can be very large, we return the result **modulo 1e9 + 7**.

---

# Constraints

```
1 <= zero, one, limit <= 1000
```

The solution must therefore be efficient enough to handle up to **1000 × 1000 states**.

---

# Intuition

When I first read the problem, I noticed that the main restriction is about **consecutive numbers**.

So I started thinking about building the binary array step by step.

At each step I can place either:

* `0`
* `1`

But I must make sure I never create more than `limit` identical consecutive values.

This naturally suggests a **Dynamic Programming** solution.

I decided to track:

* how many zeros I used
* how many ones I used
* what the last placed value was

This allows me to enforce the consecutive limit rule.

---

# Approach

### Step 1 — Define DP State

Let:

```
dp[i][j][0] = number of arrays using i zeros and j ones ending with 0

dp[i][j][1] = number of arrays using i zeros and j ones ending with 1
```

Where:

* `i` = zeros used
* `j` = ones used

---

### Step 2 — Base Cases

If the array contains only zeros:

```
dp[i][0][0] = 1   if i <= limit
```

Similarly for ones:

```
dp[0][j][1] = 1   if j <= limit
```

Because longer sequences would violate the limit.

---

### Step 3 — Transition

If we want to end with `0`:

We append `0` to a previous array.

```
dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1]
```

However, we must subtract cases where we exceed the `limit` consecutive `0`s.

Those cases occur when the sequence started with `1` and then added `limit+1` zeros.

So we subtract:

```
dp[i-limit-1][j][1]
```

Similarly for `1`:

```
dp[i][j][1] = dp[i][j-1][0] + dp[i][j-1][1]
```

And subtract invalid sequences:

```
dp[i][j-limit-1][0]
```

---

### Step 4 — Final Result

The array may end with either value:

```
answer = dp[zero][one][0] + dp[zero][one][1]
```

---

# Data Structures Used

The solution uses a **3D Dynamic Programming array**:

```
dp[zero+1][one+1][2]
```

The third dimension represents the last element:

* `0` → array ends with zero
* `1` → array ends with one

---

# Operations & Behavior Summary

| Operation               | Purpose                                    |
| ----------------------- | ------------------------------------------ |
| Initialize base states  | Valid sequences of only zeros or only ones |
| DP transitions          | Extend previous sequences                  |
| Subtract invalid states | Prevent exceeding consecutive limit        |
| Modulo operation        | Prevent integer overflow                   |

---

# Complexity

### Time Complexity

```
O(zero × one)
```

We compute each DP state once.

### Space Complexity

```
O(zero × one)
```

For storing the DP table.

---

# Multi-language Solutions

## C++

```cpp
class Solution {
public:
    int numberOfStableArrays(int zero, int one, int limit) {

        const int MOD = 1e9 + 7;

        vector<vector<array<long long,2>>> dp(
            zero+1, vector<array<long long,2>>(one+1,{0,0}));

        for(int i=1;i<=min(zero,limit);i++)
            dp[i][0][0]=1;

        for(int j=1;j<=min(one,limit);j++)
            dp[0][j][1]=1;

        for(int i=1;i<=zero;i++){
            for(int j=1;j<=one;j++){

                long long over0=(i-limit-1>=0)?dp[i-limit-1][j][1]:0;
                long long over1=(j-limit-1>=0)?dp[i][j-limit-1][0]:0;

                dp[i][j][0]=(dp[i-1][j][0]+dp[i-1][j][1]-over0+MOD)%MOD;
                dp[i][j][1]=(dp[i][j-1][0]+dp[i][j-1][1]-over1+MOD)%MOD;
            }
        }

        return (dp[zero][one][0]+dp[zero][one][1])%MOD;
    }
};
```

---

## Java

```java
class Solution {
    public int numberOfStableArrays(int zero, int one, int limit) {

        int MOD = 1_000_000_007;

        long[][][] dp = new long[zero+1][one+1][2];

        for(int i=1;i<=Math.min(zero,limit);i++)
            dp[i][0][0]=1;

        for(int j=1;j<=Math.min(one,limit);j++)
            dp[0][j][1]=1;

        for(int i=1;i<=zero;i++){
            for(int j=1;j<=one;j++){

                long over0=(i-limit-1>=0)?dp[i-limit-1][j][1]:0;
                long over1=(j-limit-1>=0)?dp[i][j-limit-1][0]:0;

                dp[i][j][0]=(dp[i-1][j][0]+dp[i-1][j][1]-over0+MOD)%MOD;
                dp[i][j][1]=(dp[i][j-1][0]+dp[i][j-1][1]-over1+MOD)%MOD;
            }
        }

        return (int)((dp[zero][one][0]+dp[zero][one][1])%MOD);
    }
}
```

---

## JavaScript

```javascript
var numberOfStableArrays = function(zero, one, limit) {

    const MOD = 1000000007;

    const dp = Array.from({length: zero+1},()=>
        Array.from({length: one+1},()=>[0,0])
    );

    for(let i=1;i<=Math.min(zero,limit);i++)
        dp[i][0][0]=1;

    for(let j=1;j<=Math.min(one,limit);j++)
        dp[0][j][1]=1;

    for(let i=1;i<=zero;i++){
        for(let j=1;j<=one;j++){

            let over0=(i-limit-1>=0)?dp[i-limit-1][j][1]:0;
            let over1=(j-limit-1>=0)?dp[i][j-limit-1][0]:0;

            dp[i][j][0]=(dp[i-1][j][0]+dp[i-1][j][1]-over0+MOD)%MOD;
            dp[i][j][1]=(dp[i][j-1][0]+dp[i][j-1][1]-over1+MOD)%MOD;
        }
    }

    return (dp[zero][one][0]+dp[zero][one][1])%MOD;
};
```

---

## Python3

```python
class Solution:
    def numberOfStableArrays(self, zero: int, one: int, limit: int) -> int:

        MOD = 10**9 + 7

        dp = [[[0,0] for _ in range(one+1)] for _ in range(zero+1)]

        for i in range(1,min(zero,limit)+1):
            dp[i][0][0]=1

        for j in range(1,min(one,limit)+1):
            dp[0][j][1]=1

        for i in range(1,zero+1):
            for j in range(1,one+1):

                over0 = dp[i-limit-1][j][1] if i-limit-1>=0 else 0
                over1 = dp[i][j-limit-1][0] if j-limit-1>=0 else 0

                dp[i][j][0]=(dp[i-1][j][0]+dp[i-1][j][1]-over0)%MOD
                dp[i][j][1]=(dp[i][j-1][0]+dp[i][j-1][1]-over1)%MOD

        return (dp[zero][one][0]+dp[zero][one][1])%MOD
```

---

## Go

```go
func numberOfStableArrays(zero int, one int, limit int) int {

    const MOD int = 1e9 + 7

    dp := make([][][]int, zero+1)

    for i:=range dp{
        dp[i]=make([][]int,one+1)
        for j:=range dp[i]{
            dp[i][j]=make([]int,2)
        }
    }

    for i:=1;i<=min(zero,limit);i++{
        dp[i][0][0]=1
    }

    for j:=1;j<=min(one,limit);j++{
        dp[0][j][1]=1
    }

    for i:=1;i<=zero;i++{
        for j:=1;j<=one;j++{

            over0:=0
            if i-limit-1>=0{
                over0=dp[i-limit-1][j][1]
            }

            over1:=0
            if j-limit-1>=0{
                over1=dp[i][j-limit-1][0]
            }

            dp[i][j][0]=(dp[i-1][j][0]+dp[i-1][j][1]-over0+MOD)%MOD
            dp[i][j][1]=(dp[i][j-1][0]+dp[i][j-1][1]-over1+MOD)%MOD
        }
    }

    return (dp[zero][one][0]+dp[zero][one][1])%MOD
}

func min(a,b int) int{
    if a<b{return a}
    return b
}
```

---

# Step-by-step Detailed Explanation

1. Create a DP table that tracks zeros used, ones used, and last element.

2. Initialize base states for sequences containing only zeros or only ones.

3. Iterate through all `(i, j)` states.

4. For each state:

   * Extend sequences ending with zero
   * Extend sequences ending with one

5. Subtract invalid sequences that exceed the `limit` consecutive rule.

6. Apply modulo to prevent overflow.

7. Return the sum of sequences ending with `0` and `1`.

---

# Examples

### Example 1

Input

```
zero = 1
one = 1
limit = 2
```

Output

```
2
```

Valid arrays

```
[1,0]
[0,1]
```

---

### Example 2

Input

```
zero = 1
one = 2
limit = 1
```

Output

```
1
```

Valid array

```
[1,0,1]
```

---

# How to use / Run locally

Example for C++:

```
g++ solution.cpp -o solution
./solution
```

Example for Python:

```
python solution.py
```

---

# Notes & Optimizations

* Dynamic Programming avoids exponential recursion.
* Subtracting invalid sequences ensures the limit constraint is satisfied.
* Using modulo prevents integer overflow.

Possible improvement:

* Space optimization could reduce memory usage.

---

# Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
