# 1569. Number of Ways to Reorder Array to Get Same BST

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

You are given an array called nums. It is a permutation of numbers from 1 to n.  

You build a binary search tree by inserting the numbers one by one in the order they appear in the array.  

The task is to count how many different ways you can rearrange the same numbers so that the binary search tree you get is exactly the same as the one built from the original array.  

Because the answer can be very large, you must return it modulo 10^9 + 7.  

In short, this is a classic combinatorics on binary search tree problem: count the valid reorderings that preserve the same BST structure.

## Constraints

- 1 <= nums.length <= 1000
- 1 <= nums[i] <= nums.length
- All values in nums are distinct

## Intuition

The first number in the array always becomes the root of the BST.  

Every number smaller than the root must go into the left subtree, and every larger number must go into the right subtree.  

The relative order of the numbers inside the left group and inside the right group must stay the same as in the original array. Otherwise the shape of the subtrees would change.  

So the problem reduces to choosing positions for the left-group numbers among the remaining spots after the root, then doing the same thing recursively for both subtrees.  

This is why a divide-and-conquer approach combined with binomial coefficients works so well for counting the number of ways to reorder an array to get the same BST.

## Approach

I start by precomputing factorials and their modular inverses up to the length of the array. This lets me calculate binomial coefficients quickly under the required modulo.  

Then I write a recursive function that takes a list of numbers that will form a subtree.  

The first element of that list is treated as the root.  

I split the remaining numbers into two lists: one for values smaller than the root (left) and one for values larger than the root (right). I keep the original relative order in both lists.  

The number of ways for the current subtree is the binomial coefficient C(leftSize + rightSize, leftSize) multiplied by the number of ways for the left subtree and the number of ways for the right subtree.  

If a list has two or fewer elements, there is only one possible ordering, so I return 1.  

At the very end I subtract 1 from the total because the original ordering itself is counted and the problem asks for the number of other valid reorderings.

## Data Structures Used

- Arrays / vectors / lists to store the current subtree numbers and to build the left and right groups.  
- Two extra arrays for precomputed factorials and modular inverses so binomial coefficients can be answered in constant time.  

These structures are simple and keep the implementation clean while still allowing the recursive divide-and-conquer logic to run efficiently for n up to 1000.

## Operations & Behavior Summary

1. Precompute fact[i] = i! mod 10^9+7 and the corresponding modular inverses.  
2. Call the recursive helper on the full input array.  
3. Inside the helper:  
   - If the list is tiny (size ≤ 2) return 1.  
   - Identify the root (first element).  
   - Partition the rest into left and right lists while preserving order.  
   - Compute C(total remaining positions, size of left) × ways(left) × ways(right).  
4. Subtract 1 from the final answer and take modulo 10^9+7.  

The recursion naturally mirrors the structure of the binary search tree, which is why the same approach works for every language version.

## Complexity

| Complexity       | Value   | Explanation |
|------------------|---------|-------------|
| Time Complexity  | O(n²)   | Each recursive call scans its list once (O(n) work). In the worst case the recursion tree has O(n) nodes, giving overall quadratic time. Factorial precomputation is only O(n). |
| Space Complexity | O(n²)   | New left and right lists are created at every level. In a skewed tree the total extra space across the call stack reaches O(n²). Factorial tables take only O(n). |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int numOfWays(vector<int>& nums) {
        const int MOD = 1e9 + 7;
        int n = nums.size();
        vector<long long> fact(n + 1, 1), inv(n + 1, 1);
        for (int i = 1; i <= n; i++) fact[i] = fact[i - 1] * i % MOD;
        inv[n] = modPow(fact[n], MOD - 2, MOD);
        for (int i = n - 1; i >= 0; i--) inv[i] = inv[i + 1] * (i + 1) % MOD;
        auto comb = [&](int a, int b) -> long long {
            if (b < 0 || b > a) return 0;
            return fact[a] * inv[b] % MOD * inv[a - b] % MOD;
        };
        function<long long(vector<int>&)> ways = [&](vector<int>& arr) -> long long {
            int m = arr.size();
            if (m <= 2) return 1;
            vector<int> left, right;
            for (int x : arr) {
                if (x < arr[0]) left.push_back(x);
                else if (x > arr[0]) right.push_back(x);
            }
            return comb(m - 1, left.size()) * ways(left) % MOD * ways(right) % MOD;
        };
        return (ways(nums) - 1 + MOD) % MOD;
    }
private:
    long long modPow(long long base, long long exp, long long mod) {
        long long res = 1;
        while (exp) {
            if (exp & 1) res = res * base % mod;
            base = base * base % mod;
            exp >>= 1;
        }
        return res;
    }
};
```

### Java
```java
class Solution {
    public int numOfWays(int[] nums) {
        final int MOD = 1000000007;
        int n = nums.length;
        long[] fact = new long[n + 1];
        long[] inv = new long[n + 1];
        fact[0] = 1;
        for (int i = 1; i <= n; i++) fact[i] = fact[i - 1] * i % MOD;
        inv[n] = modPow(fact[n], MOD - 2, MOD);
        for (int i = n - 1; i >= 0; i--) inv[i] = inv[i + 1] * (i + 1) % MOD;
        return (int)((ways(nums, fact, inv, MOD) - 1 + MOD) % MOD);
    }
    private long ways(int[] arr, long[] fact, long[] inv, int MOD) {
        int m = arr.length;
        if (m <= 2) return 1;
        List<Integer> left = new ArrayList<>();
        List<Integer> right = new ArrayList<>();
        for (int x : arr) {
            if (x < arr[0]) left.add(x);
            else if (x > arr[0]) right.add(x);
        }
        int[] l = left.stream().mapToInt(i -> i).toArray();
        int[] r = right.stream().mapToInt(i -> i).toArray();
        return comb(m - 1, l.length, fact, inv, MOD) * ways(l, fact, inv, MOD) % MOD * ways(r, fact, inv, MOD) % MOD;
    }
    private long comb(int a, int b, long[] fact, long[] inv, int MOD) {
        if (b < 0 || b > a) return 0;
        return fact[a] * inv[b] % MOD * inv[a - b] % MOD;
    }
    private long modPow(long base, long exp, long mod) {
        long res = 1;
        while (exp > 0) {
            if ((exp & 1) == 1) res = res * base % mod;
            base = base * base % mod;
            exp >>= 1;
        }
        return res;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var numOfWays = function(nums) {
    const MOD = 1000000007n;
    const n = nums.length;
    const fact = new Array(n + 1).fill(1n);
    const inv = new Array(n + 1).fill(1n);
    for (let i = 1; i <= n; i++) fact[i] = fact[i - 1] * BigInt(i) % MOD;
    inv[n] = modPow(fact[n], MOD - 2n, MOD);
    for (let i = n - 1; i >= 0; i--) inv[i] = inv[i + 1] * BigInt(i + 1) % MOD;
    function comb(a, b) {
        if (b < 0 || b > a) return 0n;
        return fact[a] * inv[b] % MOD * inv[a - b] % MOD;
    }
    function ways(arr) {
        const m = arr.length;
        if (m <= 2) return 1n;
        const left = [], right = [];
        for (const x of arr) {
            if (x < arr[0]) left.push(x);
            else if (x > arr[0]) right.push(x);
        }
        return comb(m - 1, left.length) * ways(left) % MOD * ways(right) % MOD;
    }
    return Number((ways(nums) - 1n + MOD) % MOD);
};
function modPow(base, exp, mod) {
    let res = 1n;
    while (exp > 0n) {
        if (exp & 1n) res = res * base % mod;
        base = base * base % mod;
        exp >>= 1n;
    }
    return res;
}
```

### TypeScript
```typescript
function numOfWays(nums: number[]): number {
    const MOD = 1000000007n;
    const n = nums.length;
    const fact: bigint[] = new Array(n + 1).fill(1n);
    const inv: bigint[] = new Array(n + 1).fill(1n);
    for (let i = 1; i <= n; i++) fact[i] = fact[i - 1] * BigInt(i) % MOD;
    inv[n] = modPow(fact[n], MOD - 2n, MOD);
    for (let i = n - 1; i >= 0; i--) inv[i] = inv[i + 1] * BigInt(i + 1) % MOD;
    function comb(a: number, b: number): bigint {
        if (b < 0 || b > a) return 0n;
        return fact[a] * inv[b] % MOD * inv[a - b] % MOD;
    }
    function ways(arr: number[]): bigint {
        const m = arr.length;
        if (m <= 2) return 1n;
        const left: number[] = [], right: number[] = [];
        for (const x of arr) {
            if (x < arr[0]) left.push(x);
            else if (x > arr[0]) right.push(x);
        }
        return comb(m - 1, left.length) * ways(left) % MOD * ways(right) % MOD;
    }
    return Number((ways(nums) - 1n + MOD) % MOD);
}
function modPow(base: bigint, exp: bigint, mod: bigint): bigint {
    let res = 1n;
    while (exp > 0n) {
        if (exp & 1n) res = res * base % mod;
        base = base * base % mod;
        exp >>= 1n;
    }
    return res;
}
```

### Python3
```python
class Solution:
    def numOfWays(self, nums: list[int]) -> int:
        MOD = 10**9 + 7
        n = len(nums)
        fact = [1] * (n + 1)
        for i in range(1, n + 1):
            fact[i] = fact[i - 1] * i % MOD
        inv = [1] * (n + 1)
        inv[n] = pow(fact[n], MOD - 2, MOD)
        for i in range(n - 1, -1, -1):
            inv[i] = inv[i + 1] * (i + 1) % MOD
        def comb(a, b):
            if b < 0 or b > a:
                return 0
            return fact[a] * inv[b] % MOD * inv[a - b] % MOD
        def ways(arr):
            m = len(arr)
            if m <= 2:
                return 1
            root = arr[0]
            left = [x for x in arr if x < root]
            right = [x for x in arr if x > root]
            return comb(m - 1, len(left)) * ways(left) % MOD * ways(right) % MOD
        return (ways(nums) - 1) % MOD
```

### Go
```go
func numOfWays(nums []int) int {
    const MOD = 1000000007
    n := len(nums)
    fact := make([]int64, n+1)
    inv := make([]int64, n+1)
    fact[0] = 1
    for i := 1; i <= n; i++ {
        fact[i] = fact[i-1] * int64(i) % MOD
    }
    inv[n] = modPow(fact[n], MOD-2, MOD)
    for i := n - 1; i >= 0; i-- {
        inv[i] = inv[i+1] * int64(i+1) % MOD
    }
    var ways func([]int) int64
    ways = func(arr []int) int64 {
        m := len(arr)
        if m <= 2 {
            return 1
        }
        left, right := []int{}, []int{}
        for _, x := range arr {
            if x < arr[0] {
                left = append(left, x)
            } else if x > arr[0] {
                right = append(right, x)
            }
        }
        return comb(m-1, len(left), fact, inv, MOD) * ways(left) % MOD * ways(right) % MOD
    }
    return int((ways(nums) - 1 + MOD) % MOD)
}
func comb(a, b int, fact, inv []int64, MOD int64) int64 {
    if b < 0 || b > a {
        return 0
    }
    return fact[a] * inv[b] % MOD * inv[a-b] % MOD
}
func modPow(base, exp, mod int64) int64 {
    res := int64(1)
    for exp > 0 {
        if exp&1 == 1 {
            res = res * base % mod
        }
        base = base * base % mod
        exp >>= 1
    }
    return res
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages, so the reasoning stays the same.

I first prepare the factorial and inverse-factorial tables. This is pure modular arithmetic. Once those tables exist, any binomial coefficient C(a, b) is just three array lookups and two multiplications under the modulo.

The recursive function always receives a sequence that already respects the relative order needed for its subtree.  

- The first number is forced to be the root.  
- Everything smaller must stay on the left, everything larger on the right.  
- Because the original relative order is kept when I build the two new lists, the sub-problems are exactly the same kind of problem on smaller arrays.  

The binomial coefficient tells me how many ways I can interleave the left-group numbers with the right-group numbers in the remaining positions after the root. Multiplying by the answers of the two recursive calls finishes the count for the current node.

When the list becomes size 0, 1 or 2 there is nothing left to choose, so the base case returns 1.  

Finally I subtract one. The original sequence is always a valid way to build the tree, but the problem statement asks only for the other reorderings. Taking the result modulo 10^9+7 keeps the answer inside the required range and non-negative.

Edge cases such as a completely sorted array (only one possible BST) or an array that produces a balanced tree are handled automatically by the same recursion and the final “minus one”.

## Examples

**Example 1**  
Input: nums = [2,1,3]  
Output: 1  

The original order builds a tree with 2 as root, 1 on the left and 3 on the right.  
The only other valid order is [2,3,1].  
Total ways including the original = 2, so the answer after subtracting 1 is 1.

**Example 2**  
Input: nums = [3,4,5,1,2]  
Output: 5  

Root is 3. Left group in order: [1,2]. Right group in order: [4,5].  
There are C(4,2) = 6 ways to interleave them, and each subgroup has only one internal ordering.  
After subtracting the original sequence we obtain 5.

**Example 3**  
Input: nums = [1,2,3]  
Output: 0  

This produces a right spine. The relative order is completely fixed, so only one possible sequence exists. After subtracting 1 the answer becomes 0.

## How to Use / Run Locally

**C++**  
1. Copy the C++ solution into a file named main.cpp.  
2. Compile with: g++ -std=c++17 main.cpp -o main  
3. Run: ./main  
(You will need to add a small main function that reads the array and prints the result if you want to test interactively.)

**Java**  
1. Place the Solution class inside a file named Solution.java.  
2. Compile: javac Solution.java  
3. Run: java Solution  
(Again add a main method for local testing.)

**JavaScript**  
1. Save the code in a file named solution.js.  
2. Run with Node: node solution.js  
(Add a few console.log calls with sample inputs to see the output.)

**TypeScript**  
1. Save as solution.ts.  
2. Compile: tsc solution.ts  
3. Run the generated JavaScript: node solution.js  

**Python3**  
1. Save as solution.py.  
2. Run: python3 solution.py  
(Add a few print statements with the sample arrays.)

**Go**  
1. Save as main.go.  
2. Run: go run main.go  
(Add a main function that calls numOfWays with the example arrays.)

## Notes & Optimizations

- The length constraint (n ≤ 1000) makes the O(n²) solution perfectly acceptable.  
- Using modular inverses instead of computing binomial coefficients from scratch each time keeps every call fast.  
- An alternative bottom-up dynamic-programming version is possible but more awkward to write because the sub-problems are defined by the actual values that appear, not by simple index ranges.  
- Watch out for the final “minus one”. Forgetting it produces the wrong answer on every sample that has only one valid ordering.  
- In languages that lack built-in big integers (plain JavaScript numbers, for example) you must be careful with intermediate multiplications; using BigInt or a modular multiplication helper avoids overflow.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)