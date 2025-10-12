# 3539. Find Sum of Array Product of Magical Sequences

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
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

I am given two integers `m` and `k`, and an integer array `nums` of length `n`. A sequence `seq` of length `m` (ordered picks of indices from `0..n-1`) is called **magical** if the integer
`S = 2^{seq[0]} + 2^{seq[1]} + ... + 2^{seq[m-1]}`
has exactly `k` set bits in its binary representation. The **array product** of sequence `seq` is `prod(seq) = nums[seq[0]] * nums[seq[1]] * ... * nums[seq[m-1]]`. I must return the **sum of array products** over all valid magical sequences modulo `10^9 + 7`.

Important points:

* Sequences are **ordered** (so permutations of the same indices count separately).
* When I pick index `i` `t_i` times, product contribution is `nums[i]^{t_i}`.
* Binary addition of `t_i * 2^i` across indices includes carries — these determine set bits.

---

## Constraints

* `1 <= k <= m <= 30`
* `1 <= nums.length <= 50`
* `1 <= nums[i] <= 10^8`
* Answer modulo `10^9 + 7`.

Note: `m ≤ 30` is small and lets me use DP states indexed by `m` dimensions.

---

## Intuition

I thought about the problem as distributing `m` ordered picks among `n` indices. When index `i` is chosen `t_i` times:

* The multiplicative product is `nums[i]^{t_i}`.
* The count of ordered sequences that realize counts `{t_i}` equals `m! / (∏ t_i!)`. I can produce this iteratively using binomial coefficients `C[r][t]` while reducing remaining picks `r`.
* The binary number formed by sum `S = Σ t_i * 2^i` is affected by carries between bit positions. So when processing index `i`, I add `t_i` to the incoming carry for that bit. The resulting bit is `(carry + t_i) & 1`; the next carry is `(carry + t_i) >> 1`.

So I use a DP that processes indices one-by-one and tracks:

* `r`: remaining picks (0..m),
* `carry`: carry into current bit (0..m),
* `ones`: number of ones produced in lower bits (0..m).

DP values are the sum of contributions (products × multiplicity) for that state.

---

## Approach

1. Precompute binomial coefficients `C[r][t]` for `0 <= r,t <= m`.
2. Precompute `powA[i][t] = nums[i]^t mod MOD` for `t = 0..m`.
3. Use DP array `dp[r][carry][ones]` (rolled by index `i`):

   * Initialize `dp[m][0][0] = 1` (no index processed, m picks remain).
   * For each index `i` and each state with value `val`, consider assigning `t` picks to `i` (0..r).
   * New state: `newr = r - t`, `s = carry + t`, `bit = s & 1`, `newones = ones + bit`, `newcarry = s >> 1`.
   * Multiply `val` by `C[r][t] * powA[i][t]` and add to `dp_next[newr][newcarry][newones]`.
4. After all indices processed, require `r == 0`. For each state add `popcount(carry)` to `ones` and if it equals `k`, accumulate that `dp` value to answer.
5. Return answer mod `10^9+7`.

Key correctness reasons:

* The DP enumerates count vectors `(t_0..t_{n-1})` and their multinomial multiplicity.
* It simulates binary addition bit-by-bit via `carry` propagation.
* `powA` accounts for product contributions.

---

## Data Structures Used

* 3D arrays (rolled by index) for DP: `dp[r][carry][ones]` (integers modulo `MOD`).
* 2D array for combinations `C` (Pascal triangle).
* 2D array for `powA`.
* Primitive integers / BigInt (JS) to keep modular arithmetic exact.

---

## Operations & Behavior Summary

* Pascal triangle computation: `O(m^2)`.
* Power table per `nums[i]`: `O(n * m)`.
* DP transitions: for each index `i`, loop over `r` (0..m), `carry` (0..m), `ones` (0..m) and `t` (0..r). This is bounded by `O(n * m^3)` worst-case (but `m ≤ 30`).
* Modular multiplication and addition are applied at each DP update.

---

## Complexity

* **Time Complexity:** `O(n * m^3)` worst-case. Here `n` = `nums.length`, `m` is the sequence length (≤ 30). Because `m` is small, this is practical.
* **Space Complexity:** `O(m^3)` for the DP arrays. We roll only one index at a time so memory is manageable for `m ≤ 30`.

---

## Multi-language Solutions

Below are clean, working implementations in the requested languages. The JavaScript solution uses `BigInt` to avoid precision loss (important for intermediate modular arithmetic).

---

### C++

```c++
/* C++ implementation */
#include <bits/stdc++.h>
using namespace std;
using int64 = long long;
const int MOD = 1000000007;

class Solution {
public:
    int magicalSum(int m, int k, vector<int>& nums) {
        int n = nums.size();
        // Precompute combinations C up to m
        vector<vector<int64>> C(m+1, vector<int64>(m+1,0));
        for(int i=0;i<=m;i++){
            C[i][0] = C[i][i] = 1;
            for(int j=1;j<i;j++){
                C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD;
            }
        }
        // Precompute powA[i][t] = nums[i]^t mod MOD
        vector<vector<int64>> powA(n, vector<int64>(m+1,1));
        for(int i=0;i<n;i++){
            int64 a = nums[i] % MOD;
            for(int t=1;t<=m;t++){
                powA[i][t] = (powA[i][t-1] * a) % MOD;
            }
        }

        int M = m;
        // dp[r][carry][ones]
        vector<vector<vector<int64>>> cur(M+1, vector<vector<int64>>(M+1, vector<int64>(M+1,0)));
        cur[M][0][0] = 1;

        for(int i=0;i<n;i++){
            vector<vector<vector<int64>>> nxt(M+1, vector<vector<int64>>(M+1, vector<int64>(M+1,0)));
            for(int r=0;r<=M;r++){
                for(int carry=0; carry<=M; carry++){
                    for(int ones=0; ones<=M; ones++){
                        int64 val = cur[r][carry][ones];
                        if(val==0) continue;
                        for(int t=0;t<=r;t++){
                            int newr = r - t;
                            int sum = carry + t;
                            int bit = sum & 1;
                            int newones = ones + bit;
                            if(newones > M) continue;
                            int newcarry = sum >> 1;
                            int64 mult = (C[r][t] * powA[i][t]) % MOD;
                            int64 add = (val * mult) % MOD;
                            nxt[newr][newcarry][newones] += add;
                            if(nxt[newr][newcarry][newones] >= MOD) nxt[newr][newcarry][newones] -= MOD;
                        }
                    }
                }
            }
            cur.swap(nxt);
        }

        int64 ans = 0;
        for(int carry=0; carry<=M; carry++){
            for(int ones=0; ones<=M; ones++){
                int64 val = cur[0][carry][ones];
                if(val==0) continue;
                int extra = __builtin_popcount((unsigned)carry);
                if(ones + extra == k){
                    ans = (ans + val) % MOD;
                }
            }
        }
        return (int)ans;
    }
};
```

---

### Java

```java
// Java implementation
import java.util.*;

public class Solution {
    static final int MOD = 1000000007;
    public int magicalSum(int m, int k, int[] nums) {
        int n = nums.length;
        long[][] C = new long[m+1][m+1];
        for(int i=0;i<=m;i++){
            C[i][0] = C[i][i] = 1;
            for(int j=1;j<i;j++){
                C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD;
            }
        }
        long[][] powA = new long[n][m+1];
        for(int i=0;i<n;i++){
            powA[i][0] = 1;
            for(int t=1;t<=m;t++){
                powA[i][t] = (powA[i][t-1] * (nums[i] % MOD)) % MOD;
            }
        }

        int M = m;
        long[][][] cur = new long[M+1][M+1][M+1];
        cur[M][0][0] = 1;

        for(int i=0;i<n;i++){
            long[][][] nxt = new long[M+1][M+1][M+1];
            for(int r=0;r<=M;r++){
                for(int carry=0; carry<=M; carry++){
                    for(int ones=0; ones<=M; ones++){
                        long val = cur[r][carry][ones];
                        if(val==0) continue;
                        for(int t=0;t<=r;t++){
                            int newr = r - t;
                            int sum = carry + t;
                            int bit = (sum & 1);
                            int newones = ones + bit;
                            if(newones > M) continue;
                            int newcarry = sum >> 1;
                            long mult = (C[r][t] * powA[i][t]) % MOD;
                            long add = (val * mult) % MOD;
                            nxt[newr][newcarry][newones] += add;
                            if(nxt[newr][newcarry][newones] >= MOD) nxt[newr][newcarry][newones] -= MOD;
                        }
                    }
                }
            }
            cur = nxt;
        }

        long ans = 0;
        for(int carry=0; carry<=M; carry++){
            for(int ones=0; ones<=M; ones++){
                long val = cur[0][carry][ones];
                if(val==0) continue;
                int extra = Integer.bitCount(carry);
                if(ones + extra == k){
                    ans = (ans + val) % MOD;
                }
            }
        }
        return (int)ans;
    }
}
```

---

### JavaScript

```javascript
/**
 * JavaScript implementation (uses BigInt for correctness)
 * @param {number} m
 * @param {number} k
 * @param {number[]} nums
 * @return {number}
 */
var magicalSum = function(m, k, nums) {
    const MOD = BigInt(1000000007);
    const n = nums.length;
    // combinations as BigInt
    const C = Array.from({length: m+1}, () => Array(m+1).fill(0n));
    for(let i=0;i<=m;i++){
        C[i][0] = 1n; C[i][i] = 1n;
        for(let j=1;j<i;j++){
            C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD;
        }
    }
    // powers powA[i][t] = nums[i]^t (BigInt)
    const powA = Array.from({length: n}, () => Array(m+1).fill(1n));
    for(let i=0;i<n;i++){
        powA[i][0] = 1n;
        const a = BigInt(nums[i]) % MOD;
        for(let t=1;t<=m;t++){
            powA[i][t] = (powA[i][t-1] * a) % MOD;
        }
    }
    const M = m;
    // cur[r][carry][ones] : BigInt
    let cur = Array.from({length: M+1}, () =>
        Array.from({length: M+1}, () =>
            Array(M+1).fill(0n)
        )
    );
    cur[M][0][0] = 1n;

    for(let i=0;i<n;i++){
        let nxt = Array.from({length: M+1}, () =>
            Array.from({length: M+1}, () =>
                Array(M+1).fill(0n)
            )
        );
        for(let r=0;r<=M;r++){
            for(let carry=0; carry<=M; carry++){
                for(let ones=0; ones<=M; ones++){
                    let val = cur[r][carry][ones];
                    if(val === 0n) continue;
                    for(let t=0;t<=r;t++){
                        let newr = r - t;
                        let sum = carry + t;
                        let bit = sum & 1;
                        let newones = ones + bit;
                        if(newones > M) continue;
                        let newcarry = sum >>> 1;
                        let mult = (C[r][t] * powA[i][t]) % MOD;
                        let add = (val * mult) % MOD;
                        nxt[newr][newcarry][newones] = (nxt[newr][newcarry][newones] + add) % MOD;
                    }
                }
            }
        }
        cur = nxt;
    }

    let ans = 0n;
    for(let carry=0; carry<=M; carry++){
        for(let ones=0; ones<=M; ones++){
            let val = cur[0][carry][ones];
            if(val === 0n) continue;
            let extra = popcount(carry);
            if(ones + extra === k){
                ans = (ans + val) % MOD;
            }
        }
    }
    return Number(ans);

    function popcount(x){
        let c = 0;
        while(x > 0){
            c += x & 1;
            x >>>= 1;
        }
        return c;
    }
};
```

---

### Python3

```python
# Python3 implementation
from typing import List
MOD = 10**9 + 7

class Solution:
    def magicalSum(self, m: int, k: int, nums: List[int]) -> int:
        n = len(nums)
        # combinations
        C = [[0]*(m+1) for _ in range(m+1)]
        for i in range(m+1):
            C[i][0] = C[i][i] = 1
            for j in range(1,i):
                C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD
        # powers
        powA = [[1]*(m+1) for _ in range(n)]
        for i in range(n):
            a = nums[i] % MOD
            for t in range(1, m+1):
                powA[i][t] = (powA[i][t-1] * a) % MOD
        M = m
        cur = [[[0]*(M+1) for _ in range(M+1)] for __ in range(M+1)]
        cur[M][0][0] = 1

        for i in range(n):
            nxt = [[[0]*(M+1) for _ in range(M+1)] for __ in range(M+1)]
            for r in range(M+1):
                for carry in range(M+1):
                    for ones in range(M+1):
                        val = cur[r][carry][ones]
                        if val == 0:
                            continue
                        for t in range(r+1):
                            newr = r - t
                            s = carry + t
                            bit = s & 1
                            newones = ones + bit
                            if newones > M:
                                continue
                            newcarry = s >> 1
                            mult = (C[r][t] * powA[i][t]) % MOD
                            add = (val * mult) % MOD
                            nxt[newr][newcarry][newones] = (nxt[newr][newcarry][newones] + add) % MOD
            cur = nxt

        ans = 0
        for carry in range(M+1):
            for ones in range(M+1):
                val = cur[0][carry][ones]
                if val == 0:
                    continue
                extra = bin(carry).count("1")
                if ones + extra == k:
                    ans = (ans + val) % MOD
        return ans
```

---

### Go

```go
// Go implementation
package main
import (
    "fmt"
)

const MOD int64 = 1000000007

func magicalSum(m int, k int, nums []int) int {
    n := len(nums)
    C := make([][]int64, m+1)
    for i:=0;i<=m;i++ { C[i] = make([]int64, m+1) }
    for i:=0;i<=m;i++ {
        C[i][0] = 1
        C[i][i] = 1
        for j:=1;j<i;j++ {
            C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD
        }
    }
    powA := make([][]int64, n)
    for i:=0;i<n;i++ {
        powA[i] = make([]int64, m+1)
        powA[i][0] = 1
        a := int64(nums[i]) % MOD
        for t:=1;t<=m;t++ {
            powA[i][t] = (powA[i][t-1] * a) % MOD
        }
    }

    M := m
    cur := make([][][]int64, M+1)
    for r:=0;r<=M;r++ {
        cur[r] = make([][]int64, M+1)
        for c:=0;c<=M;c++ {
            cur[r][c] = make([]int64, M+1)
        }
    }
    cur[M][0][0] = 1

    for i:=0;i<n;i++ {
        nxt := make([][][]int64, M+1)
        for r:=0;r<=M;r++ {
            nxt[r] = make([][]int64, M+1)
            for c:=0;c<=M;c++ { nxt[r][c] = make([]int64, M+1) }
        }
        for r:=0;r<=M;r++ {
            for carry:=0; carry<=M; carry++ {
                for ones:=0; ones<=M; ones++ {
                    val := cur[r][carry][ones]
                    if val == 0 { continue }
                    for t:=0; t<=r; t++ {
                        newr := r - t
                        s := carry + t
                        bit := s & 1
                        newones := ones + bit
                        if newones > M { continue }
                        newcarry := s >> 1
                        mult := (C[r][t] * powA[i][t]) % MOD
                        add := (val * mult) % MOD
                        nxt[newr][newcarry][newones] = (nxt[newr][newcarry][newones] + add) % MOD
                    }
                }
            }
        }
        cur = nxt
    }

    var ans int64 = 0
    for carry:=0; carry<=M; carry++ {
        for ones:=0; ones<=M; ones++ {
            val := cur[0][carry][ones]
            if val == 0 { continue }
            extra := popcount(carry)
            if ones + extra == k {
                ans = (ans + val) % MOD
            }
        }
    }
    return int(ans)
}

func popcount(x int) int {
    cnt := 0
    for x>0 {
        if x&1 == 1 { cnt++ }
        x >>= 1
    }
    return cnt
}

func main() {
    fmt.Println(magicalSum(2,2, []int{5,4,3,2,1})) // quick test
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I will explain the core pieces of the code and how they work; the logic is identical across languages.

1. **Combinations (Pascal triangle)**

   * We compute `C[i][j]` for `0 <= i, j <= m`:

     * `C[i][0] = C[i][i] = 1`.
     * `C[i][j] = C[i-1][j-1] + C[i-1][j]`.
   * Reason: When `r` picks remain and I choose `t` of them to assign to the current index, there are `C[r][t]` ways.

2. **Powers `powA[i][t] = nums[i]^t (mod MOD)`**

   * Iterate `t` from `1..m` and multiply progressively:

     * `powA[i][t] = powA[i][t-1] * nums[i] (mod MOD)`.
   * Reason: If I select index `i` `t` times, its multiplicative contribution is `nums[i]^t`.

3. **DP State**

   * `dp[r][carry][ones]` holds the sum of contributions for the partially processed prefix of indices such that:

     * `r` picks still remain,
     * the incoming binary `carry` to current bit is `carry`,
     * the number of `1` bits already placed in processed (lower) positions is `ones`.
   * Initialize with all picks unassigned: `dp[m][0][0] = 1`.

4. **Transition**

   * For each index `i` and each reachable `(r, carry, ones)` with value `val`:

     * For `t` from `0` to `r`:

       * `newr = r - t`
       * `s = carry + t`
       * `bit = s & 1` → whether bit `i` becomes 1
       * `newones = ones + bit`
       * `newcarry = s >> 1`
       * multiplicative factor = `C[r][t] * powA[i][t]`
       * `dp_next[newr][newcarry][newones] += val * factor (mod MOD)`
   * This covers all ways to allocate the `t` picks to index `i` and accounts for ordered multiplicity.

5. **Finalization**

   * After processing all indices we only accept states with `r == 0` (all picks used).
   * Any leftover `carry` contributes `popcount(carry)` to the number of set bits.
   * If `ones + popcount(carry) == k`, add that `dp` value to the answer.

6. **BigInt note for JavaScript**

   * Intermediate multiplications (combinations × pow × counts) can exceed `2^53`.
   * Using `BigInt` in JS prevents precision loss. Convert final answer back to `Number` since `ans mod 1e9+7` fits in Number.

---

## Examples

1. Example from prompt:

```
Input: m = 5, k = 5, nums = [1,10,100,1000,10000]
Output: 991600007
Explanation: Every permutation of indices [0,1,2,3,4] is a magical sequence in this case and each product equals 1*10*100*1000*10000.
```

2. Another:

```
Input: m = 2, k = 2, nums = [5,4,3,2,1]
Output: 170
```

3. Edge:

```
Input: m = 1, k = 1, nums = [28]
Output: 28
```

---

## How to use / Run locally

### C++

* Create `solution.cpp`, paste the C++ class and a main driver if desired.
* Compile and run:

```bash
g++ -std=c++17 solution.cpp -O2 -o sol
./sol
```

### Java

* Put `Solution` in `Solution.java` and add a `main` method that reads input or tests.
* Compile and run:

```bash
javac Solution.java
java Solution
```

### JavaScript (Node)

* Save the JS function in `solution.js` and call `magicalSum` from a small driver or REPL.
* Run:

```bash
node solution.js
```

Note: Node supports `BigInt` (Node >= 10.4). Make sure to use `BigInt` literals if testing manually.

### Python3

* Save `Solution` class in `solution.py` and call `Solution().magicalSum(...)` from `if __name__ == '__main__'`.
* Run:

```bash
python3 solution.py
```

### Go

* Save in `main.go`, implement `main()` to test.
* Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* The DP state dimension limits derive from `m ≤ 30`. If `m` were much larger, this approach becomes infeasible.
* Memory optimization: I already roll by index `i` (only `cur` and `nxt` exist). Further micro-optimization can compress `ones` dimension if needed by storing only states with non-zero values (sparse hash map) — but for `m ≤ 30` dense arrays are simpler and fast.
* JavaScript must use `BigInt` due to intermediate overflow concerns. All modular arithmetic steps are computed with `BigInt` and reduced by `MOD`.
* The algorithm correctly accounts for ordered sequences via incremental binomial choices, ultimately yielding the multinomial factor `m! / ∏ t_i!`.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
