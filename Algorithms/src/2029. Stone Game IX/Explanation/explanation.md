# 2029. Stone Game IX - LeetCode Solution

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

LeetCode 2029, **Stone Game IX**, is a game theory and greedy problem involving two players, Alice and Bob.

We are given an integer array `stones`. Each number represents the value of one stone.

Alice always plays first. On every turn, the current player removes exactly one stone. The player who removes a stone loses immediately if the sum of all removed stones becomes divisible by `3`.

If all stones are removed without anyone losing because of this rule, Bob wins.

Both players play optimally, so the goal is to determine whether Alice can force a win.

The main observation is that we do not need the exact value of each stone. We only need its remainder when divided by `3`.

Every stone belongs to one of three groups:

* Remainder `0`
* Remainder `1`
* Remainder `2`

The final solution counts these three groups and uses their relationship to determine the winner.

## Constraints

* `1 <= stones.length <= 10^5`
* `1 <= stones[i] <= 10^4`

## Intuition

At first, this problem looks like I need to simulate Alice's and Bob's moves. That would be difficult because each player can choose many different stones.

I then noticed that the losing condition only cares whether the current sum is divisible by `3`.

Because of that, the exact stone values are not important. For example, `1`, `4`, `7`, and `10` all behave the same way because they all have remainder `1` when divided by `3`.

So I reduce every stone to one of these three values:

* `0` if `stone % 3 == 0`
* `1` if `stone % 3 == 1`
* `2` if `stone % 3 == 2`

Now the whole game depends only on three counts.

The interesting part is the number of remainder-`0` stones. These stones do not change the current sum modulo `3`, so their parity affects which player gets the advantage.

After analyzing the possible game states, I get two cases.

If the number of remainder-`0` stones is even, Alice needs at least one remainder-`1` stone and one remainder-`2` stone.

If the number of remainder-`0` stones is odd, Alice wins when the counts of remainder-`1` and remainder-`2` stones differ by more than `2`.

This gives a very small constant-size decision after one pass through the array.

## Approach

I use three counters to store the number of stones with each remainder modulo `3`.

First, I scan the complete `stones` array and increase the appropriate counter.

Let:

* `cnt[0]` be the number of stones divisible by `3`
* `cnt[1]` be the number of stones with remainder `1`
* `cnt[2]` be the number of stones with remainder `2`

Then I check the parity of `cnt[0]`.

If `cnt[0]` is even, Alice wins only when both `cnt[1]` and `cnt[2]` are greater than zero.

If `cnt[0]` is odd, I calculate the absolute difference between `cnt[1]` and `cnt[2]`.

Alice wins when:

`abs(cnt[1] - cnt[2]) > 2`

Otherwise, Bob wins.

For example, consider:

`stones = [2, 3, 2, 2]`

The remainders are:

`2, 0, 2, 2`

So:

* `cnt[0] = 1`
* `cnt[1] = 0`
* `cnt[2] = 3`

Since `cnt[0]` is odd, I use the second condition:

`abs(0 - 3) = 3`

Because `3 > 2`, Alice wins.

## Data Structures Used

### Three-element counter array

I use a fixed array of size `3` to count the three possible remainders.

* `cnt[0]` stores the number of remainder-`0` stones.
* `cnt[1]` stores the number of remainder-`1` stones.
* `cnt[2]` stores the number of remainder-`2` stones.

This is enough because there are only three possible values for `stone % 3`.

No dynamic programming table, graph, set, or other large data structure is required.

## Operations & Behavior Summary

1. Create three counters for remainders `0`, `1`, and `2`.
2. Traverse every stone in the input array.
3. Calculate `stone % 3`.
4. Increase the counter for that remainder.
5. Check whether the number of remainder-`0` stones is even or odd.
6. If it is even, Alice wins only when both remainder-`1` and remainder-`2` groups exist.
7. If it is odd, calculate the absolute difference between the remainder-`1` and remainder-`2` counts.
8. Alice wins if that difference is greater than `2`.
9. Return the resulting boolean value.

## Complexity

| Complexity       | Result | Explanation                                                                 |
| ---------------- | ------ | --------------------------------------------------------------------------- |
| Time Complexity  | `O(n)` | I scan the `n` stones once and calculate each stone's remainder modulo `3`. |
| Space Complexity | `O(1)` | I only use three counters, whose size remains constant regardless of `n`.   |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool stoneGameIX(vector<int>& stones) {
        // cnt[r] stores how many stones have remainder r when divided by 3.
        int cnt[3] = {0, 0, 0};

        // Count stones in each remainder group.
        for (int stone : stones) {
            // Only the remainder matters for the game.
            cnt[stone % 3]++;
        }

        // When the number of remainder-0 stones is even,
        // Alice needs at least one remainder-1 and one remainder-2 stone.
        if (cnt[0] % 2 == 0) {
            return cnt[1] > 0 && cnt[2] > 0;
        }

        // When the number of remainder-0 stones is odd,
        // Alice wins if the two useful groups differ by more than 2.
        return abs(cnt[1] - cnt[2]) > 2;
    }
};
```

### Java

```java
class Solution {
    public boolean stoneGameIX(int[] stones) {
        // cnt[r] stores how many stones have remainder r when divided by 3.
        int[] cnt = new int[3];

        // Count the stones in each remainder group.
        for (int stone : stones) {
            // Only the remainder affects the game.
            cnt[stone % 3]++;
        }

        // With an even number of remainder-0 stones,
        // Alice needs both a remainder-1 and a remainder-2 stone.
        if (cnt[0] % 2 == 0) {
            return cnt[1] > 0 && cnt[2] > 0;
        }

        // With an odd number of remainder-0 stones,
        // Alice wins when the two useful groups differ by more than 2.
        return Math.abs(cnt[1] - cnt[2]) > 2;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} stones
 * @return {boolean}
 */
var stoneGameIX = function(stones) {
    // cnt[r] stores how many stones have remainder r modulo 3.
    const cnt = [0, 0, 0];

    // Count the stones in each remainder group.
    for (const stone of stones) {
        // Only the remainder matters for the game.
        cnt[stone % 3]++;
    }

    // With an even number of remainder-0 stones,
    // Alice needs at least one stone from both useful groups.
    if (cnt[0] % 2 === 0) {
        return cnt[1] > 0 && cnt[2] > 0;
    }

    // With an odd number of remainder-0 stones,
    // a difference greater than 2 lets Alice force a win.
    return Math.abs(cnt[1] - cnt[2]) > 2;
};
```

### Python3

```python
from typing import List

class Solution:
    def stoneGameIX(self, stones: List[int]) -> bool:
        # cnt[r] stores how many stones have remainder r modulo 3.
        cnt = [0, 0, 0]

        # Count the stones in each remainder group.
        for stone in stones:
            # Only the remainder matters for deciding the winner.
            cnt[stone % 3] += 1

        # With an even number of remainder-0 stones,
        # Alice needs both a remainder-1 and a remainder-2 stone.
        if cnt[0] % 2 == 0:
            return cnt[1] > 0 and cnt[2] > 0

        # With an odd number of remainder-0 stones,
        # Alice wins when the two useful groups differ by more than 2.
        return abs(cnt[1] - cnt[2]) > 2
```

### Go

```go
func stoneGameIX(stones []int) bool {
 // cnt[r] stores how many stones have remainder r modulo 3.
 cnt := [3]int{}

 // Count the stones in each remainder group.
 for _, stone := range stones {
  // Only the remainder matters for the game.
  cnt[stone%3]++
 }

 // With an even number of remainder-0 stones,
 // Alice needs both a remainder-1 and a remainder-2 stone.
 if cnt[0]%2 == 0 {
  return cnt[1] > 0 && cnt[2] > 0
 }

 // Calculate the difference between the two useful groups.
 diff := cnt[1] - cnt[2]

 // Make the difference positive.
 if diff < 0 {
  diff = -diff
 }

 // With an odd number of zero-remainder stones,
 // a difference greater than 2 gives Alice a winning strategy.
 return diff > 2
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The implementation is almost identical in all five languages because the core algorithm is based only on counting remainders.

### Step 1: Create three counters

I start with a fixed-size array containing three counters.

The three positions represent the three possible remainders after division by `3`.

The meaning is:

* Position `0` represents numbers divisible by `3`.
* Position `1` represents numbers with remainder `1`.
* Position `2` represents numbers with remainder `2`.

I do not create a counter for every possible stone value because that would be unnecessary.

### Step 2: Count the remainders

I loop through every stone.

For each stone, I calculate its remainder using modulo `3`.

For example:

* `3 % 3 = 0`
* `4 % 3 = 1`
* `5 % 3 = 2`
* `6 % 3 = 0`

Then I increase the corresponding counter.

This reduces the original problem to only three numbers.

### Step 3: Check the number of remainder-0 stones

The next important value is `cnt[0]`.

A remainder-`0` stone does not change the current sum modulo `3`.

For example, adding `3` to a sum does not change its remainder modulo `3`.

Because of this, the parity of `cnt[0]` matters.

I check whether `cnt[0]` is even.

### Step 4: Handle an even number of remainder-0 stones

If `cnt[0]` is even, Alice needs both types of useful stones.

That means:

* At least one remainder-`1` stone must exist.
* At least one remainder-`2` stone must exist.

If either group is missing, Alice cannot force the required sequence of moves.

So the condition is:

`cnt[1] > 0 && cnt[2] > 0`

If both are available, Alice wins.

### Step 5: Handle an odd number of remainder-0 stones

If `cnt[0]` is odd, the condition changes.

I compare the number of remainder-`1` and remainder-`2` stones.

I calculate:

`abs(cnt[1] - cnt[2])`

If the difference is greater than `2`, Alice has enough imbalance to force a win.

So the condition becomes:

`abs(cnt[1] - cnt[2]) > 2`

### Step 6: Why the failing case works

Consider:

`stones = [2, 3, 2, 2]`

The remainder groups are:

* One remainder-`0` stone
* Zero remainder-`1` stones
* Three remainder-`2` stones

Therefore:

`cnt[0] = 1`

`cnt[1] = 0`

`cnt[2] = 3`

Since `cnt[0]` is odd, I do not require both `cnt[1]` and `cnt[2]` to exist.

Instead:

`abs(0 - 3) = 3`

Since `3 > 2`, the result is `true`.

This is an important edge case because a solution that always requires both remainder-`1` and remainder-`2` stones will fail here.

### Step 7: Language differences

The C++, Java, JavaScript, Python3, and Go implementations all follow the same mathematical logic.

The main syntax differences are only in:

* Declaring the counter array.
* Iterating over the input array.
* Calculating the absolute value.
* Returning the boolean result.

There is no language-specific algorithmic optimization needed.

## Examples

### Example 1

Input:

`stones = [2, 1]`

The remainders are:

`2, 1`

So:

* `cnt[0] = 0`
* `cnt[1] = 1`
* `cnt[2] = 1`

`cnt[0]` is even.

Both `cnt[1]` and `cnt[2]` are greater than zero, so Alice wins.

Expected Output:

`true`

### Example 2

Input:

`stones = [2]`

The counts are:

* `cnt[0] = 0`
* `cnt[1] = 0`
* `cnt[2] = 1`

`cnt[0]` is even, but `cnt[1]` is zero.

Therefore Alice cannot force a win.

Expected Output:

`false`

### Example 3

Input:

`stones = [2, 3, 2, 2]`

The remainders are:

`2, 0, 2, 2`

So:

* `cnt[0] = 1`
* `cnt[1] = 0`
* `cnt[2] = 3`

`cnt[0]` is odd, so I compare the two useful groups.

`abs(0 - 3) = 3`

Since `3 > 2`, Alice wins.

Expected Output:

`true`

## How to Use / Run Locally

### C++

Save the solution in a C++ source file such as `solution.cpp`.

Compile it with a C++17-compatible compiler using the command `g++ -std=c++17 solution.cpp -o solution`.

Run the compiled program using `./solution` on Linux or macOS. On Windows, run `solution.exe`.

For LeetCode, you do not need to write the input/output driver code. Paste the `Solution` class into the LeetCode editor and submit it.

### Java

Save the solution in `Solution.java`.

Compile it with `javac Solution.java`.

Run it with `java Solution`.

On LeetCode, the `Solution` class and required method can be submitted directly without creating a separate `main` method.

### JavaScript

Save the solution in a file such as `solution.js`.

Run it with Node.js using `node solution.js`.

For LeetCode, paste the provided function into the JavaScript editor. LeetCode handles the test cases automatically.

### Python3

Save the solution in a file such as `solution.py`.

Run it using `python3 solution.py`.

For LeetCode, paste the `Solution` class into the Python3 editor and submit it.

### Go

Save the solution in a Go source file such as `solution.go`.

Run it with `go run solution.go`.

For LeetCode, paste the required function into the Go editor according to LeetCode's expected function signature.

## Notes & Optimizations

The most important optimization is realizing that the actual stone values are unnecessary after calculating their remainder modulo `3`.

A brute-force game simulation would be far too expensive because there can be up to `10^5` stones.

A recursive minimax solution would also be impractical because the number of possible game states grows rapidly.

The modulo observation reduces the problem to only three counters.

One important edge case is when `cnt[1]` or `cnt[2]` is zero. It is not always correct to immediately return `false`.

The parity of `cnt[0]` must be checked first.

For an even `cnt[0]`, both remainder-`1` and remainder-`2` groups are required.

For an odd `cnt[0]`, the correct condition is based on the difference between the two groups:

`abs(cnt[1] - cnt[2]) > 2`

This distinction is essential for cases such as `[2, 3, 2, 2]`.

The solution uses constant extra space and performs only one pass through the input, making it efficient enough for the maximum constraint.

This approach is a useful example of how a game theory problem can sometimes be reduced to a small mathematical state instead of simulating every possible move.

## Author

Md Aarzoo Islam
