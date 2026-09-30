# 1111. Maximum Nesting Depth of Two Valid Parentheses Strings – LeetCode Solution

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

You are given a valid parentheses string called `seq`. A valid parentheses string only contains the characters `(` and `)` and follows the classic rules of balanced brackets.

Your task is to split this string into two subsequences A and B. Both A and B must also be valid parentheses strings. The split does not have to keep characters next to each other; any character can go to either group as long as the relative order is preserved.

The nesting depth of a valid parentheses string is the deepest level of nested brackets it contains. You need to choose the split so that the maximum of the two depths (depth of A and depth of B) becomes as small as possible.

Finally return an array of the same length as `seq`. In this array, `0` means the character at that position belongs to group A and `1` means it belongs to group B. Any correct split that achieves the minimum possible maximum depth is accepted.

This is a classic LeetCode parentheses problem that asks you to minimize the maximum nesting depth after splitting a valid parentheses string into two groups.

## Constraints

- `1 <= seq.length <= 10000`
- `seq` consists only of the characters `(` and `)`
- `seq` is a valid parentheses string

## Intuition

When I first read the problem I noticed that the original string is already balanced. That means every opening bracket has a matching closing bracket and the nesting is well-defined.

The deepest nesting level in the whole string is what forces the answer higher. If I could give every other nesting level to a different group, neither group would ever get two consecutive levels. The maximum depth of either group would then be roughly half of the original maximum depth.

Tracking the current depth while walking the string and assigning each bracket to a group based on the parity of that depth felt like the simplest way to achieve this alternating assignment.

## Approach

I keep a single counter that stores the current nesting depth and an answer array of the same length as the input.

I walk through every character from left to right:

- When I see an opening parenthesis I first increase the depth counter. Then I look at whether the new depth is even or odd and store that parity (0 or 1) in the answer array. That decides which group receives the opening bracket.
- When I see a closing parenthesis I first store the current depth’s parity in the answer array so the closing bracket goes to the same group that received its matching open. After that I decrease the depth counter.

Because consecutive nesting levels always receive opposite groups, the maximum depth that appears in either group stays as small as possible. Both resulting groups remain valid parentheses strings because every open is matched by a close that was assigned the same group, and the original string was already balanced.

## Data Structures Used

- An integer array (or list) of the same length as the input string. This stores the final group assignment for every character.
- A single integer variable that tracks the current nesting depth while scanning the string.

No stacks, queues or other heavy structures are required because the input is already known to be a valid parentheses string.

## Operations & Behavior Summary

1. Create an answer array filled with zeros and set a depth counter to zero.
2. For each character in the string:
   - If the character is `(`, increase depth by one, then write `depth % 2` into the answer array.
   - If the character is `)`, write the current `depth % 2` into the answer array, then decrease depth by one.
3. After the whole string has been processed the answer array contains a valid split that minimizes the maximum nesting depth of the two groups.
4. Return the answer array.

## Complexity

| Type              | Value | Explanation                                      |
|-------------------|-------|--------------------------------------------------|
| Time Complexity   | O(n)  | We examine each of the n characters exactly once |
| Space Complexity  | O(n)  | We need an answer array of length n; only a few extra integer variables are used |

## Multi-language Solutions

### C++
```cpp
class Solution { 
public: 
    vector<int> maxDepthAfterSplit(string seq) { 
        vector<int> ans(seq.size());
        int depth = 0;
        for (int i = 0; i < seq.size(); ++i) {
            if (seq[i] == '(') {
                ++depth;
                ans[i] = depth % 2;
            } else {
                ans[i] = depth % 2;
                --depth;
            }
        }
        return ans;
    } 
};
```

### Java
```java
class Solution { 
    public int[] maxDepthAfterSplit(String seq) { 
        int n = seq.length();
        int[] ans = new int[n];
        int depth = 0;
        for (int i = 0; i < n; ++i) {
            if (seq.charAt(i) == '(') {
                ++depth;
                ans[i] = depth % 2;
            } else {
                ans[i] = depth % 2;
                --depth;
            }
        }
        return ans;
    } 
}
```

### JavaScript
```javascript
/** 
 * @param {string} seq 
 * @return {number[]} 
 */ 
var maxDepthAfterSplit = function(seq) { 
    const ans = new Array(seq.length);
    let depth = 0;
    for (let i = 0; i < seq.length; ++i) {
        if (seq[i] === '(') {
            ++depth;
            ans[i] = depth % 2;
        } else {
            ans[i] = depth % 2;
            --depth;
        }
    }
    return ans;
};
```

### TypeScript
```typescript
function maxDepthAfterSplit(seq: string): number[] {
    const ans: number[] = new Array(seq.length);
    let depth = 0;
    for (let i = 0; i < seq.length; ++i) {
        if (seq[i] === '(') {
            ++depth;
            ans[i] = depth % 2;
        } else {
            ans[i] = depth % 2;
            --depth;
        }
    }
    return ans;
};
```

### Python3
```python
class Solution:
    def maxDepthAfterSplit(self, seq: str) -> list[int]:
        ans = [0] * len(seq)
        depth = 0
        for i, ch in enumerate(seq):
            if ch == '(':
                depth += 1
                ans[i] = depth % 2
            else:
                ans[i] = depth % 2
                depth -= 1
        return ans
```

### Go
```go
func maxDepthAfterSplit(seq string) []int {
    ans := make([]int, len(seq))
    depth := 0
    for i, ch := range seq {
        if ch == '(' {
            depth++
            ans[i] = depth % 2
        } else {
            ans[i] = depth % 2
            depth--
        }
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages; only the syntax changes.

I start by allocating an answer array whose size equals the length of the input string. I also declare an integer depth and set it to zero. This depth variable will always tell me how deeply nested the current position is.

I then loop over every index of the string. At each index I look at the character.

When the character is an opening parenthesis I first add one to depth. The new value of depth is exactly the nesting level of this opening bracket. I store the result of depth modulo 2 into the answer array. That single bit decides whether the bracket goes into group 0 or group 1.

When the character is a closing parenthesis I do the opposite order: I first store the current depth modulo 2 into the answer array, then I subtract one from depth. Writing the parity before decreasing the counter guarantees that the closing bracket receives the same group number that its matching opening bracket received earlier.

Because every consecutive nesting level has opposite parity, no group ever receives two consecutive levels of nesting. That is why the maximum depth of either group is minimized. The depth counter never goes negative and always returns to zero at the end of a valid string, so both resulting subsequences stay balanced.

Edge cases such as a string of length 1 or a completely flat string of alternating pairs are handled automatically by the same parity rule.

## Examples

**Example 1**

Input: `seq = "(()())"`

Output: `[0,1,1,1,1,0]`

Trace:
- Position 0: `(` → depth becomes 1 → assign 1 % 2 = 1
- Position 1: `(` → depth becomes 2 → assign 2 % 2 = 0
- Position 2: `)` → assign 2 % 2 = 0, then depth becomes 1
- Position 3: `(` → depth becomes 2 → assign 2 % 2 = 0
- Position 4: `)` → assign 2 % 2 = 0, then depth becomes 1
- Position 5: `)` → assign 1 % 2 = 1, then depth becomes 0

(The actual published sample uses a different but equally valid assignment; any correct minimum-depth split is accepted.)

**Example 2**

Input: `seq = "()(())()"`

Output: `[0,0,0,1,1,0,1,1]`

Trace follows the same depth-parity rule. Each opening bracket receives a group based on the depth after it is opened, and each closing bracket receives the group of the depth before it is closed. The resulting two groups both have maximum depth 1, which is optimal.

**Example 3**

Input: `seq = "()"`

Output: `[0,0]`

The single pair is assigned to the same group. Depth never exceeds 1, so the maximum of the two groups stays 1 (the empty group has depth 0).

## How to Use / Run Locally

**C++**
```bash
g++ -std=c++17 solution.cpp -o solution
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

Replace the empty function body with the code you want to test, then feed the function a sample string and print the returned array.

## Notes & Optimizations

- The input is guaranteed to be a valid parentheses string, so we never need an explicit stack to verify balance.
- Using the parity of the depth is optimal; any other assignment that puts two consecutive nesting levels into the same group would produce a larger maximum depth.
- The algorithm runs in linear time and uses only a constant amount of extra memory beyond the output array, which is the best possible asymptotic complexity for this problem.
- If the problem ever allowed invalid strings, a stack-based validation step would be required first; that is unnecessary here.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)