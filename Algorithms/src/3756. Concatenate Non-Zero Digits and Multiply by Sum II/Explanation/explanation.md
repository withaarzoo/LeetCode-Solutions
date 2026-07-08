# 3756. Concatenate Non-Zero Digits and Multiply by Sum II

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
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

In this problem, we are given a string `s` containing only digits and a list of range queries.

Each query contains two indices:

`[l, r]`

For every query, we take the substring from index `l` to index `r`.

Then we remove all zero digits from that substring while keeping the remaining digits in their original order.

The remaining non-zero digits are concatenated to form a new integer `x`.

After that, we calculate the sum of all digits in `x`.

The final answer for the query is:

`x × sum of digits in x`

Since the result can be very large, every answer must be returned modulo:

`10^9 + 7`

The main challenge is answering up to `10^5` substring queries efficiently. Processing every substring character by character would be too slow, so this solution uses prefix sums, compressed non-zero digits, modular arithmetic, and prefix concatenation.

## Constraints

| Constraint        | Value                         |
| ----------------- | ----------------------------- |
| Length of `s`     | `1 <= s.length <= 10^5`       |
| Characters in `s` | Digits from `0` to `9`        |
| Number of queries | `1 <= queries.length <= 10^5` |
| Query format      | `queries[i] = [l_i, r_i]`     |
| Left index        | `0 <= l_i`                    |
| Right index       | `l_i <= r_i < s.length`       |
| Required modulo   | `10^9 + 7`                    |

These constraints make a direct solution too slow. In the worst case, scanning every query range separately could take `O(n × q)` time.

An optimized `O(n + q)` solution is needed.

## Intuition

My first thought was to handle every query directly.

For each range, I could take the substring, skip every zero, build the new number, calculate the digit sum, and multiply both values.

That approach is simple, but it repeats the same work many times.

With up to `10^5` characters and `10^5` queries, a large substring could be scanned again and again. That would be far too slow.

The key observation is that zeroes never affect the final result.

They are removed before building `x`, and they add nothing to the digit sum.

So instead of repeatedly removing zeroes for every query, I can remove them conceptually during preprocessing.

For example:

```text
Original string:   10203004
Non-zero digits:   1234
```

Now the problem becomes easier.

I only need a fast way to map each original query range to the correct range inside the compressed non-zero digit sequence.

A prefix count of non-zero digits gives me that mapping.

Once I know the compressed range, I need to find two values quickly:

1. The number formed by the digits in that range.
2. The sum of those digits.

A normal prefix sum handles the digit sum.

A prefix concatenation value, together with powers of `10`, handles the number.

This allows every query to be answered in constant time after preprocessing.

## Approach

I solve the problem in four main stages.

### 1. Count non-zero digits with a prefix array

I create an array where:

`nonZeroCount[i]`

stores the number of non-zero digits in the first `i` characters of the original string.

For a query `[l, r]`:

```text
left  = nonZeroCount[l]
right = nonZeroCount[r + 1]
```

The required digits are now the compressed range:

```text
[left, right)
```

This maps the original substring to the sequence containing only non-zero digits.

### 2. Build the compressed digit sequence

I scan the original string and store only non-zero digits.

For example:

```text
s = "10203004"
```

becomes:

```text
digits = [1, 2, 3, 4]
```

The original order is preserved.

### 3. Build prefix information

I create three prefix arrays over the compressed digits.

The first array stores concatenated prefix values.

If the compressed digits are:

```text
1, 2, 3, 4
```

then the prefix values represent:

```text
0
1
12
123
1234
```

All values are stored modulo `10^9 + 7`.

The second array stores prefix digit sums:

```text
0
1
3
6
10
```

The third array stores powers of `10`:

```text
10^0, 10^1, 10^2, 10^3, ...
```

These powers are also stored modulo `10^9 + 7`.

### 4. Answer each query in constant time

Suppose the compressed query range is:

```text
[left, right)
```

Its length is:

```text
length = right - left
```

The number formed by this range is:

```text
x = prefixValue[right]
    - prefixValue[left] × 10^length
```

The digit sum is:

```text
digitSum = prefixSum[right] - prefixSum[left]
```

The final answer is:

```text
x × digitSum mod (10^9 + 7)
```

If the query contains only zeroes, then `left == right`.

In that case, both `x` and the digit sum become `0`, so the same formula works without a separate special case.

## Data Structures Used

| Data Structure              | Purpose                                                      |
| --------------------------- | ------------------------------------------------------------ |
| Non-zero prefix count array | Maps original string positions to compressed digit positions |
| Compressed digit array      | Stores only non-zero digits in their original order          |
| Prefix value array          | Stores concatenated prefix numbers modulo `10^9 + 7`         |
| Prefix digit sum array      | Finds the sum of digits in any compressed range              |
| Power of 10 array           | Removes an unwanted prefix from a concatenated number        |
| Answer array                | Stores the result of every query                             |

The solution only uses arrays and simple arithmetic.

No tree, hash map, binary search, or advanced data structure is needed.

## Operations & Behavior Summary

The full algorithm works like this:

1. Start with the original digit string.
2. Scan it from left to right.
3. Count how many non-zero digits appear in every prefix.
4. Store every non-zero digit in a compressed sequence.
5. Build prefix concatenation values for the compressed digits.
6. Build prefix sums for the digit values.
7. Precompute powers of `10` modulo `10^9 + 7`.
8. For each query, convert the original range into a compressed range.
9. Use prefix concatenation to get the required number.
10. Use prefix sums to get the required digit sum.
11. Multiply both values.
12. Apply modulo `10^9 + 7`.
13. Store the result.

The expensive work is done only once during preprocessing.

After that, each query needs only a few array lookups and arithmetic operations.

## Complexity

| Type             | Complexity | Explanation                                                                          |
| ---------------- | ---------- | ------------------------------------------------------------------------------------ |
| Time Complexity  | `O(n + q)` | The string is processed once, and each of the `q` queries is answered in `O(1)` time |
| Space Complexity | `O(n)`     | Prefix arrays and the compressed digit sequence store at most `n` elements           |

Here:

* `n` is the length of the input string `s`.
* `q` is the number of queries.

This is the best practical complexity for the given constraints because every input character and every query must be processed at least once.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> sumAndMultiply(string s, vector<vector<int>>& queries) {
        // I use long long so multiplication does not overflow before taking modulo.
        const long long MOD = 1000000007LL;
        int n = s.size();

        // nonZeroCount[i] = number of non-zero digits in s[0..i-1].
        vector<int> nonZeroCount(n + 1, 0);

        // I store only non-zero digits because zeroes are removed in every query.
        vector<int> digits;

        // I build the prefix count and compressed digit list in one pass.
        for (int i = 0; i < n; i++) {
            // I carry the previous count forward first.
            nonZeroCount[i + 1] = nonZeroCount[i];

            // Only non-zero digits belong to the compressed sequence.
            if (s[i] != '0') {
                nonZeroCount[i + 1]++;
                digits.push_back(s[i] - '0');
            }
        }

        int k = digits.size();

        // prefixValue[i] stores the first i compressed digits as a number modulo MOD.
        vector<long long> prefixValue(k + 1, 0);

        // prefixSum[i] stores the sum of the first i compressed digits.
        vector<long long> prefixSum(k + 1, 0);

        // power10[i] stores 10^i modulo MOD for removing an earlier prefix.
        vector<long long> power10(k + 1, 1);

        // I build all prefix information over the compressed digits.
        for (int i = 0; i < k; i++) {
            // Appending a digit means multiplying the old number by 10, then adding it.
            prefixValue[i + 1] = (prefixValue[i] * 10 + digits[i]) % MOD;

            // The digit sum is a normal prefix sum.
            prefixSum[i + 1] = prefixSum[i] + digits[i];

            // I precompute the next power of 10 for constant-time range extraction.
            power10[i + 1] = (power10[i] * 10) % MOD;
        }

        // I reserve the final size so the vector does not need repeated reallocations.
        vector<int> answer;
        answer.reserve(queries.size());

        // Each query is now answered in constant time.
        for (const auto& query : queries) {
            int l = query[0];
            int r = query[1];

            // left is the number of non-zero digits strictly before l.
            int left = nonZeroCount[l];

            // right is the number of non-zero digits up to and including r.
            int right = nonZeroCount[r + 1];

            // This is the number of digits in the compressed query range.
            int len = right - left;

            // I remove the earlier compressed prefix after shifting it by len places.
            long long x = (
                prefixValue[right]
                - (prefixValue[left] * power10[len]) % MOD
                + MOD
            ) % MOD;

            // I get the digit sum using a normal prefix-sum subtraction.
            long long sum = prefixSum[right] - prefixSum[left];

            // I multiply the compressed number by its digit sum under modulo.
            answer.push_back((int)((x * sum) % MOD));
        }

        return answer;
    }
};
```

### Java

```java
class Solution {
    public int[] sumAndMultiply(String s, int[][] queries) {
        // I use long values because multiplication can exceed the int range.
        final long MOD = 1_000_000_007L;
        int n = s.length();

        // nonZeroCount[i] = number of non-zero digits in s[0..i-1].
        int[] nonZeroCount = new int[n + 1];

        // I first count all non-zero digits so I can allocate the exact array size.
        for (int i = 0; i < n; i++) {
            // I extend the previous prefix count by one only for a non-zero digit.
            nonZeroCount[i + 1] =
                nonZeroCount[i] + (s.charAt(i) != '0' ? 1 : 0);
        }

        int k = nonZeroCount[n];

        // digits stores the string after removing every zero.
        int[] digits = new int[k];
        int index = 0;

        // I fill the compressed digit array in the original order.
        for (int i = 0; i < n; i++) {
            if (s.charAt(i) != '0') {
                digits[index++] = s.charAt(i) - '0';
            }
        }

        // prefixValue[i] stores the first i compressed digits as a number modulo MOD.
        long[] prefixValue = new long[k + 1];

        // prefixSum[i] stores the sum of the first i compressed digits.
        long[] prefixSum = new long[k + 1];

        // power10[i] stores 10^i modulo MOD.
        long[] power10 = new long[k + 1];
        power10[0] = 1;

        // I build all prefix arrays in one pass over the compressed digits.
        for (int i = 0; i < k; i++) {
            // I append the current digit to the previous prefix number.
            prefixValue[i + 1] =
                (prefixValue[i] * 10 + digits[i]) % MOD;

            // I add the current digit to the previous digit sum.
            prefixSum[i + 1] =
                prefixSum[i] + digits[i];

            // I compute the next power of 10 for later range extraction.
            power10[i + 1] =
                (power10[i] * 10) % MOD;
        }

        // I create one result slot for every query.
        int[] answer = new int[queries.length];

        // I answer each query using only prefix-array lookups.
        for (int i = 0; i < queries.length; i++) {
            int l = queries[i][0];
            int r = queries[i][1];

            // left counts non-zero digits before the query.
            int left = nonZeroCount[l];

            // right counts non-zero digits through the end of the query.
            int right = nonZeroCount[r + 1];

            // len is the number of digits in the compressed query number.
            int len = right - left;

            // I subtract the shifted earlier prefix to isolate the required number.
            long x = (
                prefixValue[right]
                - (prefixValue[left] * power10[len]) % MOD
                + MOD
            ) % MOD;

            // I subtract digit-sum prefixes to get the query digit sum.
            long sum =
                prefixSum[right] - prefixSum[left];

            // I store x multiplied by its digit sum under modulo.
            answer[i] = (int)((x * sum) % MOD);
        }

        return answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @param {number[][]} queries
 * @return {number[]}
 */
var sumAndMultiply = function(s, queries) {
    // I use BigInt because JavaScript Number cannot safely multiply values near MOD.
    const MOD = 1000000007n;
    const n = s.length;

    // nonZeroCount[i] = number of non-zero digits in s[0..i-1].
    const nonZeroCount = new Array(n + 1).fill(0);

    // I store only digits that can actually appear in a query result.
    const digits = [];

    // I build the original-position mapping and compressed digits together.
    for (let i = 0; i < n; i++) {
        // I start with the count from the previous prefix.
        nonZeroCount[i + 1] = nonZeroCount[i];

        // A non-zero digit increases the count and enters the compressed list.
        if (s[i] !== '0') {
            nonZeroCount[i + 1]++;
            digits.push(Number(s[i]));
        }
    }

    const k = digits.length;

    // prefixValue[i] stores the first i compressed digits modulo MOD.
    const prefixValue = new Array(k + 1).fill(0n);

    // prefixSum[i] stores the sum of the first i compressed digits.
    const prefixSum = new Array(k + 1).fill(0);

    // power10[i] stores 10^i modulo MOD.
    const power10 = new Array(k + 1).fill(1n);

    // I build all prefix information over the compressed sequence.
    for (let i = 0; i < k; i++) {
        // I append the current digit to the previous prefix number.
        prefixValue[i + 1] =
            (prefixValue[i] * 10n + BigInt(digits[i])) % MOD;

        // I extend the normal digit-sum prefix.
        prefixSum[i + 1] =
            prefixSum[i] + digits[i];

        // I compute the next power of 10.
        power10[i + 1] =
            (power10[i] * 10n) % MOD;
    }

    // I create the result array by processing every query independently.
    return queries.map(([l, r]) => {
        // left counts non-zero digits before l.
        const left = nonZeroCount[l];

        // right counts non-zero digits through r.
        const right = nonZeroCount[r + 1];

        // len is the compressed range length.
        const len = right - left;

        // I remove the earlier prefix after shifting it by len decimal places.
        const x = (
            prefixValue[right]
            - (prefixValue[left] * power10[len]) % MOD
            + MOD
        ) % MOD;

        // I get the digit sum with prefix subtraction.
        const sum =
            prefixSum[right] - prefixSum[left];

        // The final value is below MOD, so converting it back to Number is safe.
        return Number((x * BigInt(sum)) % MOD);
    });
};
```

### Python3

```python
class Solution:
    def sumAndMultiply(self, s: str, queries: List[List[int]]) -> List[int]:
        # I use the required modulo for all concatenated number calculations.
        MOD = 10**9 + 7
        n = len(s)

        # non_zero_count[i] = number of non-zero digits in s[0:i].
        non_zero_count = [0] * (n + 1)

        # I keep only non-zero digits because zeroes never enter x or its digit sum.
        digits = []

        # I build the position mapping and compressed sequence together.
        for i, ch in enumerate(s):
            # I copy the previous count first.
            non_zero_count[i + 1] = non_zero_count[i]

            # Only a non-zero digit increases the compressed sequence length.
            if ch != '0':
                non_zero_count[i + 1] += 1
                digits.append(int(ch))

        k = len(digits)

        # prefix_value[i] stores the first i compressed digits modulo MOD.
        prefix_value = [0] * (k + 1)

        # prefix_sum[i] stores the sum of the first i compressed digits.
        prefix_sum = [0] * (k + 1)

        # power10[i] stores 10^i modulo MOD.
        power10 = [1] * (k + 1)

        # I build all prefix arrays over the compressed digits.
        for i, digit in enumerate(digits):
            # I append the current digit to the previous prefix number.
            prefix_value[i + 1] = (
                prefix_value[i] * 10 + digit
            ) % MOD

            # I extend the digit-sum prefix.
            prefix_sum[i + 1] = (
                prefix_sum[i] + digit
            )

            # I compute the next power of 10 for range extraction.
            power10[i + 1] = (
                power10[i] * 10
            ) % MOD

        # I collect one answer for every query.
        answer = []

        # Every query now takes constant time.
        for l, r in queries:
            # left counts non-zero digits before l.
            left = non_zero_count[l]

            # right counts non-zero digits through r.
            right = non_zero_count[r + 1]

            # length is the number of digits in the compressed query range.
            length = right - left

            # I remove the earlier prefix after shifting it by length places.
            x = (
                prefix_value[right]
                - prefix_value[left] * power10[length]
            ) % MOD

            # I get the digit sum by subtracting two prefix sums.
            digit_sum = (
                prefix_sum[right]
                - prefix_sum[left]
            )

            # I multiply the compressed number by its digit sum.
            answer.append((x * digit_sum) % MOD)

        return answer
```

### Go

```go
func sumAndMultiply(s string, queries [][]int) []int {
    // I use int64 because intermediate multiplication can exceed 32-bit integer limits.
    const MOD int64 = 1000000007
    n := len(s)

    // nonZeroCount[i] = number of non-zero digits in s[0..i-1].
    nonZeroCount := make([]int, n+1)

    // I store only digits that remain after removing zeroes.
    digits := make([]int64, 0, n)

    // I build the original-position mapping and compressed digits together.
    for i := 0; i < n; i++ {
        // I carry the previous prefix count forward.
        nonZeroCount[i+1] = nonZeroCount[i]

        // Only non-zero digits enter the compressed sequence.
        if s[i] != '0' {
            nonZeroCount[i+1]++
            digits = append(digits, int64(s[i]-'0'))
        }
    }

    k := len(digits)

    // prefixValue[i] stores the first i compressed digits modulo MOD.
    prefixValue := make([]int64, k+1)

    // prefixSum[i] stores the sum of the first i compressed digits.
    prefixSum := make([]int64, k+1)

    // power10[i] stores 10^i modulo MOD.
    power10 := make([]int64, k+1)
    power10[0] = 1

    // I build all prefix information in one pass.
    for i := 0; i < k; i++ {
        // I append the current digit to the previous prefix number.
        prefixValue[i+1] = (prefixValue[i]*10 + digits[i]) % MOD

        // I extend the digit-sum prefix.
        prefixSum[i+1] = prefixSum[i] + digits[i]

        // I compute the next power of 10.
        power10[i+1] = (power10[i] * 10) % MOD
    }

    // I allocate exactly one result slot for each query.
    answer := make([]int, len(queries))

    // I answer every query using constant-time prefix calculations.
    for i, query := range queries {
        l := query[0]
        r := query[1]

        // left counts non-zero digits before l.
        left := nonZeroCount[l]

        // right counts non-zero digits through r.
        right := nonZeroCount[r+1]

        // length is the number of compressed digits inside the query.
        length := right - left

        // I subtract the shifted earlier prefix to isolate the query number.
        x := (prefixValue[right] - (prefixValue[left]*power10[length])%MOD + MOD) % MOD

        // I get the digit sum with prefix subtraction.
        sum := prefixSum[right] - prefixSum[left]

        // I multiply x by its digit sum and store the result.
        answer[i] = int((x * sum) % MOD)
    }

    return answer
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core algorithm is the same in all five languages.

The main differences come from integer types, modulo multiplication, array handling, and language syntax.

### Step 1: Build the non-zero prefix count

I start by creating an array of size `n + 1`.

The first value is `0` because an empty prefix contains no non-zero digits.

For every character in the string:

* I copy the previous prefix count.
* If the current character is not `'0'`, I increase the count by one.

Suppose:

```text
s = "10203004"
```

The prefix count becomes:

```text
[0, 1, 1, 2, 2, 3, 3, 3, 4]
```

This array is important because it connects original string indices with compressed digit indices.

For example, consider:

```text
query = [1, 3]
```

The substring is:

```text
"020"
```

I calculate:

```text
left  = nonZeroCount[1] = 1
right = nonZeroCount[4] = 2
```

So the required compressed range is:

```text
[1, 2)
```

That range contains only the digit `2`.

### Step 2: Store only non-zero digits

During preprocessing, I create a sequence containing only useful digits.

For:

```text
"10203004"
```

I store:

```text
[1, 2, 3, 4]
```

This removes the need to skip zeroes during every query.

The order of the remaining digits never changes.

### Step 3: Build prefix concatenation values

I process the compressed digits from left to right.

To append a digit to an existing number, I use:

```text
newValue = oldValue × 10 + digit
```

For digits:

```text
1, 2, 3, 4
```

the values grow like this:

```text
0
1
12
123
1234
```

Since the number may contain up to `10^5` digits, it cannot fit inside a normal integer type.

So I store every prefix value modulo:

```text
1,000,000,007
```

This is valid because the final answer is also required modulo the same value.

### Step 4: Build prefix digit sums

I also keep a normal prefix sum of the compressed digits.

For:

```text
1, 2, 3, 4
```

the prefix sums are:

```text
0, 1, 3, 6, 10
```

Now the sum of any compressed range `[left, right)` is:

```text
prefixSum[right] - prefixSum[left]
```

This takes constant time.

### Step 5: Precompute powers of 10

To extract a number from the middle of a concatenated prefix, I need powers of `10`.

I precompute:

```text
10^0
10^1
10^2
10^3
...
```

All values are stored modulo `10^9 + 7`.

This prevents repeated exponent calculations for every query.

### Step 6: Extract the number for a query

Suppose the compressed digits are:

```text
1, 2, 3, 4
```

and I want the range:

```text
[1, 4)
```

The required number is:

```text
234
```

The full prefix is:

```text
1234
```

The unwanted prefix is:

```text
1
```

I cannot calculate:

```text
1234 - 1
```

because that gives:

```text
1233
```

The unwanted prefix must first be shifted three decimal places:

```text
1 × 10^3 = 1000
```

Now:

```text
1234 - 1000 = 234
```

That gives the range formula:

```text
prefixValue[right]
- prefixValue[left] × power10[length]
```

The same formula works under modular arithmetic.

### Step 7: Calculate the final result

Once I have:

```text
x
```

and:

```text
digitSum
```

I calculate:

```text
x × digitSum
```

Then I take the result modulo:

```text
10^9 + 7
```

This gives the answer for the current query.

### C++ behavior

In C++, `long long` is used for multiplication before applying modulo.

Using only `int` would be risky because two values near `10^9` can produce a result much larger than the 32-bit integer range.

Vectors are used for all prefix arrays and results.

### Java behavior

In Java, `long` is used for prefix values and modular multiplication.

The final result can be converted back to `int` because every answer is smaller than `10^9 + 7`.

Arrays work well here because all required sizes are known during preprocessing.

### JavaScript behavior

JavaScript needs extra care.

The normal `Number` type cannot safely represent every large integer multiplication used in modular arithmetic.

Using `BigInt` for prefix values, powers of `10`, and modular multiplication avoids precision errors.

After the final modulo operation, the result is small enough to convert back to `Number`.

### Python3 behavior

Python integers can grow automatically, so there is no fixed integer overflow problem.

Modulo is still applied while building prefix values because keeping numbers small makes the calculations faster.

Python also handles negative modulo results naturally, which makes the range extraction formula simple.

### Go behavior

In Go, `int64` is used for modular arithmetic.

This prevents overflow during multiplication before the modulo operation.

Slices are used for compressed digits, prefix arrays, queries, and answers.

## Examples

### Example 1

Input:

```text
s = "10203004"
queries = [[0,7], [1,3], [4,6]]
```

Expected output:

```text
[12340, 4, 9]
```

For query `[0, 7]`:

```text
Substring = "10203004"
Non-zero digits = "1234"
x = 1234
Digit sum = 1 + 2 + 3 + 4 = 10
Answer = 1234 × 10 = 12340
```

For query `[1, 3]`:

```text
Substring = "020"
Non-zero digits = "2"
x = 2
Digit sum = 2
Answer = 2 × 2 = 4
```

For query `[4, 6]`:

```text
Substring = "300"
Non-zero digits = "3"
x = 3
Digit sum = 3
Answer = 3 × 3 = 9
```

### Example 2

Input:

```text
s = "1000"
queries = [[0,3], [1,1]]
```

Expected output:

```text
[1, 0]
```

For query `[0, 3]`:

```text
Substring = "1000"
Non-zero digits = "1"
x = 1
Digit sum = 1
Answer = 1
```

For query `[1, 1]`:

```text
Substring = "0"
Non-zero digits = ""
x = 0
Digit sum = 0
Answer = 0
```

### Example 3

Input:

```text
s = "9876543210"
queries = [[0,9]]
```

Expected output:

```text
[444444137]
```

The zero is removed:

```text
x = 987654321
```

The digit sum is:

```text
9 + 8 + 7 + 6 + 5 + 4 + 3 + 2 + 1 = 45
```

The multiplication gives:

```text
987654321 × 45 = 44444444445
```

After applying modulo `10^9 + 7`:

```text
444444137
```

## How to Use / Run Locally

The solution code follows the LeetCode function format, so it does not include a complete local input/output program by default.

To test it locally, add a small `main` function or driver code for the language you are using.

### C++

Save the file as:

```text
solution.cpp
```

Compile it:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it:

```bash
./solution
```

On Windows:

```bash
solution.exe
```

### Java

Save the file as:

```text
Solution.java
```

Compile it:

```bash
javac Solution.java
```

Run it:

```bash
java Solution
```

Make sure the local test version contains a `main` method.

### JavaScript

Save the file as:

```text
solution.js
```

Run it with Node.js:

```bash
node solution.js
```

Make sure Node.js is installed on your system.

### Python3

Save the file as:

```text
solution.py
```

Run it:

```bash
python3 solution.py
```

On some systems, the command may be:

```bash
python solution.py
```

Add a small test section at the bottom of the file to create the input and print the result.

### Go

Save the file as:

```text
main.go
```

Run it directly:

```bash
go run main.go
```

Or compile it first:

```bash
go build main.go
```

Then run the generated executable.

A local Go version needs a `main` function for testing.

## Notes & Optimizations

The most important optimization is avoiding repeated substring processing.

A direct solution may look simple, but it can become `O(n × q)` in the worst case. That is too slow for `10^5` queries.

Compressing the string to non-zero digits removes unnecessary work.

The prefix non-zero count is what makes this compression useful. Without it, finding the correct compressed range for every original query would require extra searching.

The prefix concatenation technique is similar to substring hashing, but here it is used to recover the numeric value of a digit range.

Precomputing powers of `10` is also important. Calculating `10^length` separately for every query would add unnecessary work.

Queries containing only zeroes do not need special handling. The general formula naturally produces:

```text
x = 0
digitSum = 0
answer = 0
```

A binary search approach is also possible if I store the original positions of non-zero digits. For each query, I could find the first and last valid compressed positions using binary search.

However, that would make each query take `O(log n)` time.

The prefix count mapping is better because it gives the compressed range in `O(1)` time.

The final solution is an optimized prefix sum and modular arithmetic approach with:

```text
O(n + q) time
O(n) extra space
```

This makes it suitable for the maximum constraints of LeetCode 3756, Concatenate Non-Zero Digits and Multiply by Sum II.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
