# 3129. Find All Possible Stable Binary Arrays I

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

We are given three integers:

* `zero` → number of zeros that must appear in the array
* `one` → number of ones that must appear in the array
* `limit` → maximum number of identical consecutive values allowed

A binary array is considered **stable** if:

1. The number of `0`s is exactly `zero`
2. The number of `1`s is exactly `one`
3. No subarray longer than `limit` contains only `0`s or only `1`s

In simpler terms, we **cannot place more than `limit` consecutive identical numbers**.

The task is to compute the **total number of valid stable arrays**.

Since the answer may be very large, we return it **modulo 1e9 + 7**.

---

## Constraints

```
1 <= zero, one, limit <= 200
```

The total array size can be up to:

```
zero + one <= 400
```

This allows combinatorial or dynamic programming solutions.

---

## Intuition

When I first read the problem, I realized the main restriction is:

> We cannot place more than `limit` identical values consecutively.

So the array naturally breaks into **groups (blocks)** of identical numbers.

Example:

```
0 0 | 1 | 0 | 1 1
```

Here we have blocks:

```
[00] [1] [0] [11]
```

Each block length must be between:

```
1 and limit
```

So instead of building the array element by element, I thought about:

1. Splitting zeros into blocks
2. Splitting ones into blocks
3. Interleaving those blocks

The key subproblem becomes:

"How many ways can we divide `N` items into `K` groups such that no group exceeds `limit`?"

This is solved using **combinatorics with inclusion–exclusion**.

---

## Approach

### Step 1: Precompute combinations

We precompute factorials and inverse factorials so we can calculate:

```
C(n, k)
```

quickly using modular arithmetic.

---

### Step 2: Count valid block partitions

Suppose we divide `N` items into `K` blocks.

Without restriction:

```
Ways = C(N-1, K-1)
```

But some blocks may exceed `limit`.

We remove invalid cases using **inclusion–exclusion**:

```
F(N,K) = Σ (-1)^j * C(K,j) * C(N - j*limit - 1, K - 1)
```

This gives the number of ways to split `N` items into `K` groups where every group size ≤ `limit`.

---

### Step 3: Combine zero blocks and one blocks

If zeros form `k` blocks, ones may form:

```
k-1, k, or k+1
```

This happens because the sequence may start or end with either digit.

So we sum all valid combinations.

---

### Step 4: Accumulate the answer

For each possible number of zero blocks:

```
answer += ways_zero(k) * ways_one(k-1,k,k+1)
```

All operations are done modulo `1e9+7`.

---

## Data Structures Used

* Arrays for factorials
* Arrays for inverse factorials
* Helper combination function
* Temporary arrays to store block counts

These allow fast combinatorial calculations.

---

## Operations & Behavior Summary

Main operations performed in the algorithm:

1. Precompute factorials
2. Precompute inverse factorials
3. Compute combinations using modular arithmetic
4. Apply inclusion–exclusion to limit block sizes
5. Iterate through possible block counts
6. Combine zero-block and one-block possibilities

---

## Complexity

### Time Complexity

```
O((zero + one)^2)
```

Because:

* factorial preprocessing takes O(n)
* inclusion–exclusion loops run multiple times

Where:

```
n = zero + one
```

---

### Space Complexity

```
O(zero + one)
```

Used for factorial and inverse factorial arrays.

---

# Multi-language Solutions

## C++

```cpp
class Solution {
public:
    const long long MOD = 1e9 + 7;

    long long modPow(long long a,long long b){
        long long res=1;
        while(b){
            if(b&1) res=res*a%MOD;
            a=a*a%MOD;
            b>>=1;
        }
        return res;
    }

    int numberOfStableArrays(int zero,int one,int limit){

        int n=zero+one;

        vector<long long> fact(n+1),invFact(n+1);

        fact[0]=1;
        for(int i=1;i<=n;i++)
            fact[i]=fact[i-1]*i%MOD;

        invFact[n]=modPow(fact[n],MOD-2);

        for(int i=n-1;i>=0;i--)
            invFact[i]=invFact[i+1]*(i+1)%MOD;

        auto C=[&](int n,int k)->long long{
            if(k<0||k>n) return 0;
            return fact[n]*invFact[k]%MOD*invFact[n-k]%MOD;
        };

        auto F=[&](int N,int K)->long long{

            if(K<=0||K>N) return 0;

            long long ans=0;

            int maxJ=(N-K)/limit;

            for(int j=0;j<=maxJ;j++){

                long long ways=C(K,j)*C(N-j*limit-1,K-1)%MOD;

                if(j%2)
                    ans=(ans-ways+MOD)%MOD;
                else
                    ans=(ans+ways)%MOD;
            }

            return ans;
        };

        int maxK=min(zero,one+1);

        vector<long long> oneWays(maxK+3);

        for(int k=1;k<=maxK+1;k++)
            oneWays[k]=F(one,k);

        long long ans=0;

        for(int k=1;k<=maxK;k++){

            long long zWays=F(zero,k);

            long long oWays=(oneWays[k-1]+2*oneWays[k]+oneWays[k+1])%MOD;

            ans=(ans+zWays*oWays)%MOD;
        }

        return ans;
    }
};
```

## Java

```java
class Solution {

    static final long MOD = 1000000007;

    long modPow(long a,long b){
        long res=1;
        while(b>0){
            if((b&1)==1) res=res*a%MOD;
            a=a*a%MOD;
            b>>=1;
        }
        return res;
    }

    public int numberOfStableArrays(int zero,int one,int limit){

        int n=zero+one;

        long[] fact=new long[n+1];
        long[] invFact=new long[n+1];

        fact[0]=1;

        for(int i=1;i<=n;i++)
            fact[i]=fact[i-1]*i%MOD;

        invFact[n]=modPow(fact[n],MOD-2);

        for(int i=n-1;i>=0;i--)
            invFact[i]=invFact[i+1]*(i+1)%MOD;

        java.util.function.BiFunction<Integer,Integer,Long> C=(nn,kk)->{
            if(kk<0||kk>nn) return 0L;
            return fact[nn]*invFact[kk]%MOD*invFact[nn-kk]%MOD;
        };

        java.util.function.BiFunction<Integer,Integer,Long> F=(N,K)->{

            if(K<=0||K>N) return 0L;

            long ans=0;

            int maxJ=(N-K)/limit;

            for(int j=0;j<=maxJ;j++){

                long term=C.apply(K,j)*C.apply(N-j*limit-1,K-1)%MOD;

                if(j%2==1)
                    ans=(ans-term+MOD)%MOD;
                else
                    ans=(ans+term)%MOD;
            }

            return ans;
        };

        int maxK=Math.min(zero,one+1);

        long[] oneWays=new long[maxK+3];

        for(int k=1;k<=maxK+1;k++)
            oneWays[k]=F.apply(one,k);

        long ans=0;

        for(int k=1;k<=maxK;k++){

            long z=F.apply(zero,k);

            long o=(oneWays[k-1]+2*oneWays[k]+oneWays[k+1])%MOD;

            ans=(ans+z*o)%MOD;
        }

        return (int)ans;
    }
}
```

## JavaScript

```javascript
var numberOfStableArrays = function(zero, one, limit) {

    const MOD = 1e9 + 7
    const n = zero + one

    const fact = new Array(n+1).fill(1)
    const invFact = new Array(n+1).fill(1)

    const modPow=(a,b)=>{
        let res=1n
        let base=BigInt(a)
        let exp=BigInt(b)

        while(exp>0n){
            if(exp&1n) res=res*base%BigInt(MOD)
            base=base*base%BigInt(MOD)
            exp>>=1n
        }
        return Number(res)
    }

    for(let i=1;i<=n;i++)
        fact[i]=fact[i-1]*i%MOD

    invFact[n]=modPow(fact[n],MOD-2)

    for(let i=n-1;i>=0;i--)
        invFact[i]=invFact[i+1]*(i+1)%MOD

    const C=(n,k)=>{
        if(k<0||k>n) return 0
        return fact[n]*invFact[k]%MOD*invFact[n-k]%MOD
    }

    const F=(N,K)=>{
        if(K<=0||K>N) return 0

        let ans=0
        const maxJ=Math.floor((N-K)/limit)

        for(let j=0;j<=maxJ;j++){

            let term=C(K,j)*C(N-j*limit-1,K-1)%MOD

            if(j%2) ans=(ans-term+MOD)%MOD
            else ans=(ans+term)%MOD
        }

        return ans
    }

    const maxK=Math.min(zero,one+1)

    const oneWays=new Array(maxK+3).fill(0)

    for(let k=1;k<=maxK+1;k++)
        oneWays[k]=F(one,k)

    let ans=0

    for(let k=1;k<=maxK;k++){

        let z=F(zero,k)

        let o=(oneWays[k-1]+2*oneWays[k]+oneWays[k+1])%MOD

        ans=(ans+z*o)%MOD
    }

    return ans
};
```

## Python3

```python
class Solution:
    def numberOfStableArrays(self, zero: int, one: int, limit: int) -> int:

        MOD = 10**9 + 7
        n = zero + one

        fact = [1]*(n+1)
        invFact = [1]*(n+1)

        for i in range(1,n+1):
            fact[i] = fact[i-1]*i % MOD

        invFact[n] = pow(fact[n], MOD-2, MOD)

        for i in range(n-1,-1,-1):
            invFact[i] = invFact[i+1]*(i+1) % MOD

        def C(n,k):
            if k<0 or k>n:
                return 0
            return fact[n]*invFact[k]%MOD*invFact[n-k]%MOD

        def F(N,K):
            if K<=0 or K>N:
                return 0

            ans=0
            maxJ=(N-K)//limit

            for j in range(maxJ+1):

                term=C(K,j)*C(N-j*limit-1,K-1)%MOD

                if j%2:
                    ans=(ans-term)%MOD
                else:
                    ans=(ans+term)%MOD

            return ans

        maxK=min(zero,one+1)

        oneWays=[0]*(maxK+3)

        for k in range(1,maxK+2):
            oneWays[k]=F(one,k)

        ans=0

        for k in range(1,maxK+1):

            z=F(zero,k)

            o=(oneWays[k-1]+2*oneWays[k]+oneWays[k+1])%MOD

            ans=(ans+z*o)%MOD

        return ans
```

## Go

```go
func numberOfStableArrays(zero int, one int, limit int) int {

    const MOD int = 1e9 + 7
    n := zero + one

    fact := make([]int, n+1)
    invFact := make([]int, n+1)

    fact[0] = 1

    for i := 1; i <= n; i++ {
        fact[i] = fact[i-1] * i % MOD
    }

    modPow := func(a, b int) int {
        res := 1
        base := a

        for b > 0 {
            if b&1 == 1 {
                res = res * base % MOD
            }
            base = base * base % MOD
            b >>= 1
        }
        return res
    }

    invFact[n] = modPow(fact[n], MOD-2)

    for i := n - 1; i >= 0; i-- {
        invFact[i] = invFact[i+1] * (i + 1) % MOD
    }

    C := func(n, k int) int {
        if k < 0 || k > n {
            return 0
        }
        return fact[n] * invFact[k] % MOD * invFact[n-k] % MOD
    }

    F := func(N, K int) int {

        if K <= 0 || K > N {
            return 0
        }

        ans := 0

        maxJ := (N - K) / limit

        for j := 0; j <= maxJ; j++ {

            term := C(K, j) * C(N-j*limit-1, K-1) % MOD

            if j%2 == 1 {
                ans = (ans - term + MOD) % MOD
            } else {
                ans = (ans + term) % MOD
            }
        }

        return ans
    }

    maxK := zero
    if one+1 < maxK {
        maxK = one + 1
    }

    oneWays := make([]int, maxK+3)

    for k := 1; k <= maxK+1; k++ {
        oneWays[k] = F(one, k)
    }

    ans := 0

    for k := 1; k <= maxK; k++ {

        z := F(zero, k)

        o := (oneWays[k-1] + 2*oneWays[k] + oneWays[k+1]) % MOD

        ans = (ans + z*o%MOD) % MOD
    }

    return ans
}
```

---

## Step-by-step Detailed Explanation

1. Precompute factorials up to `zero + one`
2. Compute inverse factorials using Fermat's Little Theorem
3. Create a function to compute combinations `C(n,k)`
4. Implement function `F(N,K)` using inclusion–exclusion
5. Loop through possible block counts
6. Calculate valid combinations for zeros
7. Calculate valid combinations for ones
8. Combine both counts to accumulate the final answer

---

## Examples

### Example 1

```
Input:
zero = 1
one = 1
limit = 2

Output:
2

Possible arrays
[0,1]
[1,0]
```

---

### Example 2

```
Input:
zero = 1
one = 2
limit = 1

Output:
1

Valid array
[1,0,1]
```

---

## How to use / Run locally

Clone repository:

```
git clone https://github.com/yourusername/repository
```

Compile and run C++:

```
g++ solution.cpp
./a.out
```

Run Python:

```
python solution.py
```

---

## Notes & Optimizations

Key optimizations used:

* Precomputed factorials
* Modular inverse using fast exponentiation
* Inclusion–exclusion principle
* Avoided brute force generation of arrays

This reduces the complexity drastically compared to generating all permutations.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
