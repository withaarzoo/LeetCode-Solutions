# Sum of Sortable Integers – LeetCode Hard Solution (Cyclic Rotation & Divisors)

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

You are given an integer array called `nums` of length `n`.  

An integer `k` is called **sortable** when two conditions are true:

1. `k` divides `n` evenly.
2. You can turn the whole array into non-decreasing order by cutting it into consecutive blocks of length `k` and then cyclically rotating each block any number of times (left or right).

The task is to find every possible sortable `k` and return their sum.

This is a classic LeetCode hard problem that mixes array partitioning, cyclic rotation checks, and efficient divisor enumeration.

## Constraints

- `1 <= n == nums.length <= 10^5`
- `1 <= nums[i] <= 10^5`

Because `n` can be as large as 100 000, any solution that is quadratic in `n` will time out. We need a clean linear-per-divisor approach.

## Intuition

The first thing I noticed is that `k` must be a divisor of `n`. There is no point checking any other value because the array is always split into equal-sized consecutive pieces.

Once I have the sorted version of the whole array, every block of size `k` has a clear target: the matching slice of the sorted array.  

A block can be rotated into that sorted target if and only if it contains at most one strict descent (including the wrap-around from the last element back to the first). When that single descent exists, I simply start reading the block right after the descent and compare the result with the target. This check is strictly linear and avoids any expensive string-search or extra memory.

## Approach

1. Make a sorted copy of the input array. This gives every target block for free.
2. Collect all positive divisors of `n`.
3. For each candidate `k`:
   - Walk through the array in steps of `k`.
   - Inside each block count the number of positions where an element is strictly greater than the next one (modular arithmetic handles the wrap-around).
   - Reject the whole `k` as soon as any block has more than one descent.
   - If the descent count is legal, compute the correct rotation offset and verify that the rotated block exactly matches the corresponding sorted slice.
4. Add every `k` that survives all blocks to the final answer and return the sum.

The same logic works unchanged in every language; only the syntax for arrays, loops and modulo differs.

## Data Structures Used

- A plain array / vector / list that holds the sorted version of `nums`.  
  Chosen because we need random-access slices of the fully sorted order in O(1) time.
- A dynamic list (or vector) that stores the divisors of `n`.  
  The number of divisors is tiny (usually under 64 for `n ≤ 10^5`), so a simple list is perfect and keeps the code easy to read.

No hash maps, trees or complicated containers are required.

## Operations & Behavior Summary

- Sort a copy of the input once.
- Generate all divisors with a classic loop up to √n.
- For every divisor `k`:
  - For every starting index `s = 0, k, 2k, …`:
    - Count strict descents inside the window `[s … s+k)`.
    - If more than one descent appears, mark `k` invalid and move to the next divisor.
    - Otherwise rotate the window so that the element after the (possible) descent becomes the new start and compare it element-by-element with the sorted slice.
- Accumulate the valid `k` values and return their sum.

## Complexity

| Complexity | Value                  | Explanation |
|------------|------------------------|-------------|
| Time       | O(n log n + n · d(n)) | Sorting costs O(n log n). For each of the d(n) divisors we examine every element a constant number of times, giving a total of O(n · d(n)). |
| Space      | O(n)                   | Only the sorted copy is stored. All other work uses a handful of integer variables. |

`d(n)` (number of divisors) stays under 100 for the given constraints, so the algorithm comfortably fits inside the time limits of every major language.

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int sortableIntegers(vector<int>& nums) {
        int n = nums.size();
        vector<int> sorted = nums;
        sort(sorted.begin(), sorted.end());
        vector<int> divs;
        for (int i = 1; i * i <= n; ++i) {
            if (n % i == 0) {
                divs.push_back(i);
                if (i != n / i) divs.push_back(n / i);
            }
        }
        int ans = 0;
        for (int k : divs) {
            bool ok = true;
            for (int s = 0; s < n && ok; s += k) {
                int desc = 0, pos = -1;
                for (int j = 0; j < k; ++j) {
                    int a = nums[s + j];
                    int b = nums[s + (j + 1) % k];
                    if (a > b) {
                        ++desc;
                        pos = j;
                        if (desc > 1) break;
                    }
                }
                if (desc > 1) {
                    ok = false;
                    break;
                }
                int r = (desc == 0) ? 0 : (pos + 1) % k;
                for (int j = 0; j < k; ++j) {
                    if (nums[s + (r + j) % k] != sorted[s + j]) {
                        ok = false;
                        break;
                    }
                }
            }
            if (ok) ans += k;
        }
        return ans;
    }
};
```

### Java
```java
class Solution {
    public int sortableIntegers(int[] nums) {
        int n = nums.length;
        int[] sorted = nums.clone();
        Arrays.sort(sorted);
        List<Integer> divs = new ArrayList<>();
        for (int i = 1; i * i <= n; i++) {
            if (n % i == 0) {
                divs.add(i);
                if (i != n / i) divs.add(n / i);
            }
        }
        int ans = 0;
        for (int k : divs) {
            boolean ok = true;
            for (int s = 0; s < n && ok; s += k) {
                int desc = 0, pos = -1;
                for (int j = 0; j < k; j++) {
                    int a = nums[s + j];
                    int b = nums[s + (j + 1) % k];
                    if (a > b) {
                        desc++;
                        pos = j;
                        if (desc > 1) break;
                    }
                }
                if (desc > 1) {
                    ok = false;
                    break;
                }
                int r = (desc == 0) ? 0 : (pos + 1) % k;
                for (int j = 0; j < k; j++) {
                    if (nums[s + (r + j) % k] != sorted[s + j]) {
                        ok = false;
                        break;
                    }
                }
            }
            if (ok) ans += k;
        }
        return ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var sortableIntegers = function(nums) {
    const n = nums.length;
    const sorted = [...nums].sort((a, b) => a - b);
    const divs = [];
    for (let i = 1; i * i <= n; i++) {
        if (n % i === 0) {
            divs.push(i);
            if (i !== n / i) divs.push(n / i);
        }
    }
    let ans = 0;
    for (const k of divs) {
        let ok = true;
        for (let s = 0; s < n && ok; s += k) {
            let desc = 0, pos = -1;
            for (let j = 0; j < k; j++) {
                const a = nums[s + j];
                const b = nums[s + (j + 1) % k];
                if (a > b) {
                    desc++;
                    pos = j;
                    if (desc > 1) break;
                }
            }
            if (desc > 1) {
                ok = false;
                break;
            }
            const r = desc === 0 ? 0 : (pos + 1) % k;
            for (let j = 0; j < k; j++) {
                if (nums[s + (r + j) % k] !== sorted[s + j]) {
                    ok = false;
                    break;
                }
            }
        }
        if (ok) ans += k;
    }
    return ans;
};
```

### TypeScript
```typescript
function sortableIntegers(nums: number[]): number {
    const n = nums.length;
    const sorted = [...nums].sort((a, b) => a - b);
    const divs: number[] = [];
    for (let i = 1; i * i <= n; i++) {
        if (n % i === 0) {
            divs.push(i);
            if (i !== n / i) divs.push(n / i);
        }
    }
    let ans = 0;
    for (const k of divs) {
        let ok = true;
        for (let s = 0; s < n && ok; s += k) {
            let desc = 0, pos = -1;
            for (let j = 0; j < k; j++) {
                const a = nums[s + j];
                const b = nums[s + (j + 1) % k];
                if (a > b) {
                    desc++;
                    pos = j;
                    if (desc > 1) break;
                }
            }
            if (desc > 1) {
                ok = false;
                break;
            }
            const r = desc === 0 ? 0 : (pos + 1) % k;
            for (let j = 0; j < k; j++) {
                if (nums[s + (r + j) % k] !== sorted[s + j]) {
                    ok = false;
                    break;
                }
            }
        }
        if (ok) ans += k;
    }
    return ans;
}
```

### Python3
```python
class Solution:
    def sortableIntegers(self, nums: list[int]) -> int:
        n = len(nums)
        sorted_nums = sorted(nums)
        divs = []
        i = 1
        while i * i <= n:
            if n % i == 0:
                divs.append(i)
                if i != n // i:
                    divs.append(n // i)
            i += 1
        ans = 0
        for k in divs:
            ok = True
            for s in range(0, n, k):
                desc = 0
                pos = -1
                for j in range(k):
                    a = nums[s + j]
                    b = nums[s + (j + 1) % k]
                    if a > b:
                        desc += 1
                        pos = j
                        if desc > 1:
                            break
                if desc > 1:
                    ok = False
                    break
                r = 0 if desc == 0 else (pos + 1) % k
                for j in range(k):
                    if nums[s + (r + j) % k] != sorted_nums[s + j]:
                        ok = False
                        break
                if not ok:
                    break
            if ok:
                ans += k
        return ans
```

### Go
```go
func sortableIntegers(nums []int) int {
    n := len(nums)
    sorted := make([]int, n)
    copy(sorted, nums)
    sort.Ints(sorted)
    divs := []int{}
    for i := 1; i*i <= n; i++ {
        if n%i == 0 {
            divs = append(divs, i)
            if i != n/i {
                divs = append(divs, n/i)
            }
        }
    }
    ans := 0
    for _, k := range divs {
        ok := true
        for s := 0; s < n && ok; s += k {
            desc, pos := 0, -1
            for j := 0; j < k; j++ {
                a := nums[s+j]
                b := nums[s+(j+1)%k]
                if a > b {
                    desc++
                    pos = j
                    if desc > 1 {
                        break
                    }
                }
            }
            if desc > 1 {
                ok = false
                break
            }
            r := 0
            if desc == 1 {
                r = (pos + 1) % k
            }
            for j := 0; j < k; j++ {
                if nums[s+(r+j)%k] != sorted[s+j] {
                    ok = false
                    break
                }
            }
        }
        if ok {
            ans += k
        }
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core reasoning is identical across all six languages; only the way we write loops, modulo and array access changes.

1. **Create the sorted target**  
   A single sort gives us the unique non-decreasing order the whole array must reach. Every later comparison simply reads a slice of this sorted array.

2. **Enumerate divisors**  
   We loop only up to √n. Whenever `i` divides `n` we record both `i` and `n/i`. This produces the complete candidate list in negligible time.

3. **Process one candidate k**  
   We assume the candidate is valid and try to prove it wrong. The outer loop steps by `k`, so each element is visited exactly once per candidate.

4. **Count descents inside a block**  
   For every position `j` we compare `nums[s+j]` with `nums[s+(j+1)%k]`.  
   A strict greater-than is a descent.  
   - Zero descents → the block is already sorted.  
   - One descent → the sorted sequence must start right after that position.  
   - Two or more → impossible for a rotation of a sorted sequence, so we reject `k` immediately.

5. **Verify the rotation**  
   After computing the correct start offset `r` we walk the block once more:  
   `nums[s + (r + j) % k]` must equal `sorted[s + j]` for every `j`.  
   Any mismatch rejects the candidate.

6. **Accumulate the answer**  
   Only when every block of the current `k` passes both checks do we add `k` to the running sum.

Edge cases handled naturally:
- `k = 1` reduces to “is the array already sorted?”
- `k = n` reduces to “does the whole array become sorted after one cyclic rotation?”
- Arrays that are already sorted accept every divisor.
- Arrays with many equal elements still work because we only count strict descents.

## Examples

**Example 1**  
Input: `nums = [3,1,2]`  
Sorted: `[1,2,3]`  
Divisors of 3: 1 and 3.  

- `k = 1`: each singleton must equal the corresponding sorted value → fails.  
- `k = 3`: one descent (3 > 1). Rotate after that descent → `[1,2,3]` matches the target.  
Answer: 3

**Example 2**  
Input: `nums = [7,6,5]`  
Sorted: `[5,6,7]`  
- `k = 1` fails (array not sorted).  
- `k = 3` has two descents → rejected.  
Answer: 0

**Example 3**  
Input: `nums = [5,8]`  
Already sorted. Both divisors 1 and 2 succeed.  
Answer: 3

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

Replace the empty function bodies with the code you want to test, then feed the sample inputs through standard input or hard-code them inside `main`.

## Notes & Optimizations

- The descent-count technique is the key optimization. Earlier versions that built a doubled array and searched for the target were quadratic and timed out on large tests.
- Because the number of divisors is tiny, even a slightly slower constant factor is acceptable; the current version stays under a few million operations.
- Alternative ideas (KMP on the doubled block, rolling hashes, sorting each block) all work but are heavier and unnecessary once the single-descent observation is used.
- Watch out for the modulo when the block size is 1 – the wrap-around comparison becomes an element compared with itself, which correctly never counts as a descent.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)