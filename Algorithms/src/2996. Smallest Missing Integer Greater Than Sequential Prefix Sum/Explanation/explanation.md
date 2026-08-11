# 2996. Smallest Missing Integer Greater Than Sequential Prefix Sum

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

Given a 0-indexed integer array `nums`, I need to find the longest prefix that forms a sequential sequence.

A prefix is sequential when every element after the first one is exactly `1` greater than the element before it.

For example:

```text
[1, 2, 3]
```

is sequential because:

```text
2 = 1 + 1
3 = 2 + 1
```

But in:

```text
[1, 2, 3, 5, 6]
```

the sequential prefix is only:

```text
[1, 2, 3]
```

because `5` is not `3 + 1`.

After finding the longest sequential prefix, I calculate its sum.

Then I need to find the smallest integer that:

* is greater than or equal to this sum
* does not exist anywhere in the original array

### Input

An integer array `nums`.

### Output

The smallest missing integer greater than or equal to the sum of the longest sequential prefix.

This problem is commonly solved using array traversal, a hash set, and a simple linear-time approach.

## Constraints

| Constraint     | Value                    |
| -------------- | ------------------------ |
| `nums.length`  | `1 <= nums.length <= 50` |
| `nums[i]`      | `1 <= nums[i] <= 50`     |
| Array indexing | 0-indexed                |

## Intuition

I first looked at the sequential prefix because the problem gives me a very direct way to identify it.

I start from the first element and move from left to right. As long as the current number is exactly one greater than the previous number, I keep adding it to the prefix sum.

For example:

```text
nums = [1, 2, 3, 2, 5]
```

The sequence starts as:

```text
1 -> 2 -> 3
```

But after `3`, the next value is `2`, so the sequence stops.

The longest sequential prefix is therefore:

```text
[1, 2, 3]
```

Its sum is:

```text
1 + 2 + 3 = 6
```

Now I start from `6` and check whether it exists in the array.

If it exists, I try `7`, then `8`, and so on.

The first number that is not present is the answer.

## Approach

I solve the problem in two main stages.

### 1. Find the sequential prefix sum

I start with the first element as the initial sum.

Then I compare every element with the previous element.

If:

```text
nums[i] == nums[i - 1] + 1
```

the sequence continues, so I add `nums[i]` to the sum.

Otherwise, the sequential prefix has ended, and I stop.

### 2. Find the smallest missing integer

I store all values from `nums` in a hash-based set.

Then I start with the sequential prefix sum as my candidate answer.

While the candidate exists in the set, I increase it by `1`.

When I find a candidate that does not exist in the set, I return it.

This works because I start checking exactly from the smallest value allowed by the problem.

## Data Structures Used

### Hash Set

I use a hash set to store all values present in the input array.

The main reason is fast membership checking.

For example, I can quickly ask:

```text
Does 12 exist in nums?
```

instead of scanning the entire array every time.

The average lookup time of a hash set is `O(1)`.

The exact implementation depends on the language:

* C++: `unordered_set`
* Java: `HashSet`
* JavaScript: `Set`
* Python3: `set`
* Go: `map[int]bool`

## Operations & Behavior Summary

The algorithm can be viewed as the following simple process:

1. Take the first array element as the initial sequential prefix sum.
2. Start from the second element.
3. Compare the current element with the previous element.
4. If the current element is exactly one greater, add it to the sum.
5. If the sequence breaks, stop scanning the prefix.
6. Store every array value in a hash set.
7. Set the candidate answer equal to the sequential prefix sum.
8. Check whether the candidate exists in the set.
9. If it exists, increase the candidate by `1`.
10. Continue until a missing value is found.
11. Return that value.

## Complexity

| Complexity       |           Cost | Explanation                                                                                                    |
| ---------------- | -------------: | -------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n)` average | I scan the array to find the sequential prefix and use average `O(1)` hash-set operations to check candidates. |
| Space Complexity |         `O(n)` | The hash set can store up to `n` different values.                                                             |

Here, `n` represents the number of elements in `nums`.

Because the constraints are small, this approach is easily fast enough and also keeps the implementation simple.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int missingInteger(vector<int>& nums) {
        // I start with the first element because nums[0] is always
        // considered part of the sequential prefix.
        int sum = nums[0];

        // I scan the array from the second element onward.
        for (int i = 1; i < nums.size(); i++) {
            // If the current value is exactly one greater than
            // the previous value, the sequential prefix continues.
            if (nums[i] == nums[i - 1] + 1) {
                sum += nums[i];
            } else {
                // The sequence breaks here, so the prefix ends.
                break;
            }
        }

        // I store every value so I can quickly check whether
        // a candidate number exists anywhere in the array.
        unordered_set<int> seen(nums.begin(), nums.end());

        // I start checking from the sum of the longest sequential prefix.
        int answer = sum;

        // If the current number exists, it cannot be the answer,
        // so I keep moving to the next integer.
        while (seen.count(answer)) {
            answer++;
        }

        // The first missing value is the required answer.
        return answer;
    }
};
```

### Java

```java
class Solution {
    public int missingInteger(int[] nums) {
        // I start the sum with the first element because
        // nums[0] always belongs to the sequential prefix.
        int sum = nums[0];

        // I check each following element to find the longest sequential prefix.
        for (int i = 1; i < nums.length; i++) {
            // The sequence continues only when the current value
            // is exactly one greater than the previous value.
            if (nums[i] == nums[i - 1] + 1) {
                sum += nums[i];
            } else {
                // The sequence breaks, so I stop checking the prefix.
                break;
            }
        }

        // I put all array values into a HashSet for O(1) average lookup.
        java.util.HashSet<Integer> seen = new java.util.HashSet<>();

        // I add every value from the array to the set.
        for (int num : nums) {
            seen.add(num);
        }

        // I start looking for the missing number from the prefix sum.
        int answer = sum;

        // If the current candidate exists, I move to the next integer.
        while (seen.contains(answer)) {
            answer++;
        }

        // The first value not present in the array is the answer.
        return answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var missingInteger = function(nums) {
    // I start the sum with the first element because
    // the first element always belongs to the sequential prefix.
    let sum = nums[0];

    // I scan from the second element to find where the sequence breaks.
    for (let i = 1; i < nums.length; i++) {
        // The prefix continues only when the current value
        // is exactly one greater than the previous value.
        if (nums[i] === nums[i - 1] + 1) {
            sum += nums[i];
        } else {
            // The sequential prefix ends at the previous element.
            break;
        }
    }

    // I store all values in a Set so I can check membership quickly.
    const seen = new Set(nums);

    // I begin checking from the sum of the longest sequential prefix.
    let answer = sum;

    // If the candidate exists in the array, I keep increasing it.
    while (seen.has(answer)) {
        answer++;
    }

    // The first missing candidate is the required answer.
    return answer;
};
```

### Python3

```python
from typing import List

class Solution:
    def missingInteger(self, nums: List[int]) -> int:
        # I start with the first value because nums[0]
        # is always part of the sequential prefix.
        total = nums[0]

        # I scan the remaining values to find the longest sequential prefix.
        for i in range(1, len(nums)):
            # The sequence continues only if the current value
            # is exactly one greater than the previous value.
            if nums[i] == nums[i - 1] + 1:
                total += nums[i]
            else:
                # The sequence breaks here, so I stop the prefix scan.
                break

        # I create a set containing every array value for fast lookup.
        seen = set(nums)

        # I start searching from the sequential prefix sum.
        answer = total

        # If the current number exists in the array, it is not missing,
        # so I try the next number.
        while answer in seen:
            answer += 1

        # The first missing value is the required answer.
        return answer
```

### Go

```go
func missingInteger(nums []int) int {
    // I start the sum with the first element because
    // nums[0] is always part of the sequential prefix.
    sum := nums[0]

    // I scan from the second element to find the end of the sequential prefix.
    for i := 1; i < len(nums); i++ {
        // The sequence continues only when the current value
        // is exactly one greater than the previous value.
        if nums[i] == nums[i-1]+1 {
            sum += nums[i]
        } else {
            // The sequence breaks here, so I stop the prefix scan.
            break
        }
    }

    // I use a map as a set to store every value from the array.
    seen := make(map[int]bool)

    // I mark every array value as present.
    for _, num := range nums {
        seen[num] = true
    }

    // I start checking candidates from the sequential prefix sum.
    answer := sum

    // If the candidate exists, I keep increasing it until
    // I find a number that is missing from the array.
    for seen[answer] {
        answer++
    }

    // The first missing candidate is the required answer.
    return answer
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core logic is the same in all five languages. Only the syntax and the hash-set implementation change.

### C++

I start by taking `nums[0]` as the initial sequential prefix sum.

I then use a loop beginning at index `1`.

For every element, I check whether it is exactly one greater than the previous element.

If the condition is true, I add the current number to the sum.

If the condition is false, I use `break` because the longest sequential prefix cannot continue after that point.

After finding the prefix sum, I create an `unordered_set` containing all values in the array.

I set the candidate answer to the prefix sum.

Then I repeatedly check `unordered_set` membership. If the candidate is present, I increment it.

The first candidate that is not present is returned.

### Java

The Java solution follows the same logic.

I use `HashSet<Integer>` to store the values from the array.

The first loop identifies the sequential prefix and calculates its sum.

The second part checks whether the current candidate exists in the `HashSet`.

If it does, I increment the candidate until I find a missing value.

Java's `HashSet` gives average constant-time membership checks, so the overall solution remains linear on average.

### JavaScript

In JavaScript, I use the built-in `Set`.

I first calculate the sequential prefix sum by comparing neighboring values.

Then I create a `Set` from the entire array.

The `has()` method tells me whether a candidate value exists.

If the candidate exists, I keep increasing it.

When `has()` returns `false`, I have found the smallest missing integer.

### Python3

Python makes the set-based part especially short because `set(nums)` directly creates a set containing all unique array values.

I first calculate the sequential prefix sum.

Then I check:

```text
candidate in seen
```

If the candidate exists, I increase it.

Otherwise, I return it.

The important thing is that the set is used only for membership checking. I do not need to modify the original array.

### Go

Go does not have a dedicated built-in `Set` type.

I use:

```text
map[int]bool
```

as a set.

For every value in `nums`, I mark:

```text
seen[num] = true
```

Then I can check whether a candidate exists with:

```text
seen[candidate]
```

The rest of the algorithm is identical to the other four languages.

## Examples

### Example 1

#### Input

```text
nums = [1, 2, 3, 2, 5]
```

#### Find the sequential prefix

The values start as:

```text
1 -> 2 -> 3
```

The next value is `2`, but the expected value is `4`.

So the longest sequential prefix is:

```text
[1, 2, 3]
```

Its sum is:

```text
1 + 2 + 3 = 6
```

The values present in the array are:

```text
1, 2, 3, 5
```

`6` is missing, so the answer is:

```text
6
```

#### Output

```text
6
```

### Example 2

#### Input

```text
nums = [3, 4, 5, 1, 12, 14, 13]
```

The sequential prefix is:

```text
[3, 4, 5]
```

because:

```text
4 = 3 + 1
5 = 4 + 1
```

The next value is `1`, so the prefix stops.

The prefix sum is:

```text
3 + 4 + 5 = 12
```

Now I check the candidates:

```text
12 -> exists
13 -> exists
14 -> exists
15 -> missing
```

Therefore:

```text
15
```

#### Output

```text
15
```

### Example 3

#### Input

```text
nums = [5]
```

There is only one element.

A prefix containing only `nums[0]` is always sequential.

So the prefix sum is:

```text
5
```

The value `5` exists in the array, so I try `6`.

`6` is missing.

#### Output

```text
6
```

## How to Use / Run Locally

The code is written in the format expected by LeetCode, where the platform provides the input and calls the `missingInteger` function.

If you want to test the solution locally, you can create a small driver program that creates an input array and calls the solution method.

### C++

Save the solution as:

```text
solution.cpp
```

Compile it with:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it with:

```bash
./solution
```

On Windows, the generated executable can be run with:

```bash
solution.exe
```

### Java

Save the solution in:

```text
Solution.java
```

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

The class name should remain `Solution` when following the LeetCode format.

### JavaScript

Save the solution as:

```text
solution.js
```

Make sure Node.js is installed.

Run:

```bash
node solution.js
```

For local testing, add your own test input and call the `missingInteger` function.

### Python3

Save the solution as:

```text
solution.py
```

Run it with:

```bash
python3 solution.py
```

You can add a simple test case at the bottom of the file to verify the result locally.

### Go

Save the solution as:

```text
solution.go
```

Run it directly with:

```bash
go run solution.go
```

You can also compile it first:

```bash
go build solution.go
```

Then run the generated executable.

## Notes & Optimizations

The most important detail is that the sequential prefix depends on the original order of the array.

I should not sort `nums` before finding the prefix. Sorting would destroy the original ordering and could produce the wrong prefix.

For example:

```text
[3, 4, 5, 1]
```

has a sequential prefix of:

```text
[3, 4, 5]
```

If I sorted the array, I would get:

```text
[1, 3, 4, 5]
```

which is a completely different sequence.

Another important point is that I need to search for the missing integer in the entire array, not only in the sequential prefix.

The prefix is used only to calculate the starting value of the search.

For example:

```text
[3, 4, 5, 1, 12, 13, 14]
```

has a prefix sum of `12`, but the values after the prefix can still affect the final answer because `12`, `13`, and `14` are present in the complete array.

The hash set keeps those membership checks fast.

Because `nums.length` is at most `50`, even a straightforward solution would be fast enough. Still, using a hash set gives a clean `O(n)` average-time solution and is a useful pattern to remember for similar DSA problems involving missing values.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
