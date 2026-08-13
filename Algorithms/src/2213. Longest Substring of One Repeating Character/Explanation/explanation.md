# 2213. Longest Substring of One Repeating Character

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

LeetCode 2213, **Longest Substring of One Repeating Character**, gives us a string `s` and a set of queries.

Each query changes one character of `s` at a given index. After every update, we have to find the length of the longest substring where all characters are the same.

For example, if the string becomes:

```text
bbbbcc
```

the longest repeating substring is `bbbb`, so the answer is `4`.

The input contains:

* `s`: the original string
* `queryCharacters`: the new character for each query
* `queryIndices`: the index changed by each query

The output is an array where each value represents the longest repeating substring length after that query.

This problem is a good example of using a **segment tree for point updates and range queries**.

## Constraints

* `1 <= s.length <= 10^5`
* `s` consists only of lowercase English letters
* `k == queryCharacters.length == queryIndices.length`
* `1 <= k <= 10^5`
* `queryCharacters` consists only of lowercase English letters
* `0 <= queryIndices[i] < s.length`

## Intuition

My first instinct was to update the character and scan the complete string after every query.

That works logically, but it becomes far too slow.

If the string has `10^5` characters and there are `10^5` queries, scanning the whole string every time can take around `10^10` operations.

So I needed a way to update one position without recalculating the entire string.

A segment tree fits this problem well because every query is a **point update**, but I still need the answer for the complete range.

The main idea is to store enough information at every segment tree node so that two neighboring segments can be combined quickly.

For every segment, I store:

* the first character
* the last character
* the segment length
* the longest repeating prefix
* the longest repeating suffix
* the longest repeating substring anywhere in that segment

The important detail is the segment length.

For example, in `aba`, the prefix, suffix, and best value are all `1`. That does not mean the entire segment contains one repeated character.

To know whether a complete segment is uniform, I check whether its prefix or suffix length is equal to the segment length.

That gives me a safe way to merge two nodes without looking at every character again.

## Approach

I build a segment tree over the string.

Each node represents a continuous part of the string and stores six values:

1. `leftChar` — first character of the segment
2. `rightChar` — last character of the segment
3. `len` — length of the segment
4. `prefix` — longest same-character substring starting from the left
5. `suffix` — longest same-character substring ending at the right
6. `best` — longest same-character substring anywhere inside the segment

For a single character, all three lengths are `1`.

When merging a left node and a right node, I first take the best answer from either side.

Then I check the boundary.

If:

```text
left.rightChar == right.leftChar
```

a repeating substring can cross the boundary.

Its length is:

```text
left.suffix + right.prefix
```

The prefix needs special handling.

If the entire left segment is one repeated character, then its prefix can continue into the right segment.

That happens when:

```text
left.prefix == left.len
```

and the boundary characters match.

The same idea applies to the suffix:

```text
right.suffix == right.len
```

Then I process every query as a point update.

Only the leaf representing the changed index is modified. After that, I rebuild the nodes on the path from that leaf to the root.

The root represents the complete string, so its `best` value is the answer after the current query.

## Data Structures Used

### Segment Tree

The main data structure is a segment tree.

I use it because:

* each query changes only one index
* updates can be done in `O(log n)`
* the root always stores the answer for the complete string
* each node stores only a constant amount of information

### Node

Each segment tree node stores:

* first character
* last character
* segment length
* repeating prefix length
* repeating suffix length
* best repeating substring length

This lets me merge two nodes in constant time.

### Answer Array

I store one result for every query in an answer array.

Its size is `k`, where `k` is the number of queries.

## Operations & Behavior Summary

The algorithm works in three main stages.

### 1. Build the Segment Tree

I split the string recursively until every node represents one character.

For every leaf:

```text
len = 1
prefix = 1
suffix = 1
best = 1
```

Then I merge child nodes while moving back toward the root.

### 2. Update One Character

For every query:

* read the index from `queryIndices`
* read the replacement character from `queryCharacters`
* go to the corresponding leaf
* replace that character
* rebuild all affected parent nodes

Only one path in the segment tree is recalculated.

### 3. Read the Answer

The root covers the entire string.

So after every update:

```text
answer = root.best
```

I append that value to the result array.

## Complexity

| Operation          |       Complexity | Explanation                             |
| ------------------ | ---------------: | --------------------------------------- |
| Build Segment Tree |           `O(n)` | Every string position is processed once |
| One Point Update   |       `O(log n)` | Only one root-to-leaf path changes      |
| All Queries        |     `O(k log n)` | There are `k` updates                   |
| Total              | `O(n + k log n)` | Initial build plus all updates          |
| Space              |       `O(n + k)` | Segment tree plus result array          |

Here, `n` is the length of the string and `k` is the number of queries.

The extra information stored in each segment tree node is constant-sized, so the tree itself uses `O(n)` space.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    struct Node {
        char leftChar = 0;
        char rightChar = 0;
        int len = 0;
        int prefix = 0;
        int suffix = 0;
        int best = 0;
    };

    vector<Node> tree;

    Node merge(Node left, Node right) {
        Node res;

        res.leftChar = left.leftChar;
        res.rightChar = right.rightChar;
        res.len = left.len + right.len;

        res.best = max(left.best, right.best);

        res.prefix = left.prefix;

        if (left.prefix == left.len &&
            left.rightChar == right.leftChar) {
            res.prefix = left.len + right.prefix;
        }

        res.suffix = right.suffix;

        if (right.suffix == right.len &&
            left.rightChar == right.leftChar) {
            res.suffix = left.suffix + right.len;
        }

        if (left.rightChar == right.leftChar) {
            res.best = max(res.best, left.suffix + right.prefix);
        }

        return res;
    }

    void build(int node, int l, int r, const string& s) {
        if (l == r) {
            tree[node].leftChar = s[l];
            tree[node].rightChar = s[l];
            tree[node].len = 1;
            tree[node].prefix = 1;
            tree[node].suffix = 1;
            tree[node].best = 1;
            return;
        }

        int mid = l + (r - l) / 2;

        build(node * 2, l, mid, s);

        build(node * 2 + 1, mid + 1, r, s);

        tree[node] = merge(tree[node * 2], tree[node * 2 + 1]);
    }

    void update(int node, int l, int r, int idx, char c) {
        if (l == r) {
            tree[node].leftChar = c;
            tree[node].rightChar = c;
            tree[node].len = 1;
            tree[node].prefix = 1;
            tree[node].suffix = 1;
            tree[node].best = 1;
            return;
        }

        int mid = l + (r - l) / 2;

        if (idx <= mid) {
            update(node * 2, l, mid, idx, c);
        } else {
            update(node * 2 + 1, mid + 1, r, idx, c);
        }

        tree[node] = merge(tree[node * 2], tree[node * 2 + 1]);
    }

    vector<int> longestRepeating(string s, string queryCharacters, vector<int>& queryIndices) {
        int n = s.size();
        int k = queryCharacters.size();

        tree.resize(4 * n);

        build(1, 0, n - 1, s);

        vector<int> answer;
        answer.reserve(k);

        for (int i = 0; i < k; ++i) {
            update(1, 0, n - 1, queryIndices[i], queryCharacters[i]);

            answer.push_back(tree[1].best);
        }

        return answer;
    }
};
```

### Java

```java
class Solution {
    static class Node {
        char leftChar;
        char rightChar;
        int len;
        int prefix;
        int suffix;
        int best;

        Node() {
            leftChar = 0;
            rightChar = 0;
            len = 0;
            prefix = 0;
            suffix = 0;
            best = 0;
        }
    }

    private Node[] tree;

    private Node merge(Node left, Node right) {
        Node res = new Node();

        res.leftChar = left.leftChar;
        res.rightChar = right.rightChar;
        res.len = left.len + right.len;

        res.best = Math.max(left.best, right.best);

        res.prefix = left.prefix;

        if (left.prefix == left.len &&
            left.rightChar == right.leftChar) {
            res.prefix = left.len + right.prefix;
        }

        res.suffix = right.suffix;

        if (right.suffix == right.len &&
            left.rightChar == right.leftChar) {
            res.suffix = left.suffix + right.len;
        }

        if (left.rightChar == right.leftChar) {
            res.best = Math.max(
                res.best,
                left.suffix + right.prefix
            );
        }

        return res;
    }

    private void build(int node, int l, int r, char[] s) {
        if (l == r) {
            tree[node].leftChar = s[l];
            tree[node].rightChar = s[l];
            tree[node].len = 1;
            tree[node].prefix = 1;
            tree[node].suffix = 1;
            tree[node].best = 1;
            return;
        }

        int mid = l + (r - l) / 2;

        build(node * 2, l, mid, s);
        build(node * 2 + 1, mid + 1, r, s);

        tree[node] = merge(tree[node * 2], tree[node * 2 + 1]);
    }

    private void update(int node, int l, int r, int idx, char c) {
        if (l == r) {
            tree[node].leftChar = c;
            tree[node].rightChar = c;
            tree[node].len = 1;
            tree[node].prefix = 1;
            tree[node].suffix = 1;
            tree[node].best = 1;
            return;
        }

        int mid = l + (r - l) / 2;

        if (idx <= mid) {
            update(node * 2, l, mid, idx, c);
        } else {
            update(node * 2 + 1, mid + 1, r, idx, c);
        }

        tree[node] = merge(tree[node * 2], tree[node * 2 + 1]);
    }

    public int[] longestRepeating(String s, String queryCharacters, int[] queryIndices) {
        int n = s.length();
        int k = queryCharacters.length();

        tree = new Node[4 * n];

        for (int i = 0; i < tree.length; i++) {
            tree[i] = new Node();
        }

        char[] chars = s.toCharArray();

        build(1, 0, n - 1, chars);

        int[] answer = new int[k];

        for (int i = 0; i < k; i++) {
            update(
                1,
                0,
                n - 1,
                queryIndices[i],
                queryCharacters.charAt(i)
            );

            answer[i] = tree[1].best;
        }

        return answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @param {string} queryCharacters
 * @param {number[]} queryIndices
 * @return {number[]}
 */
var longestRepeating = function(s, queryCharacters, queryIndices) {
    const n = s.length;

    const tree = new Array(4 * n);

    const createNode = () => ({
        leftChar: '',
        rightChar: '',
        len: 0,
        prefix: 0,
        suffix: 0,
        best: 0
    });

    const merge = (left, right) => {
        const res = createNode();

        res.leftChar = left.leftChar;
        res.rightChar = right.rightChar;
        res.len = left.len + right.len;

        res.best = Math.max(left.best, right.best);

        res.prefix = left.prefix;

        if (
            left.prefix === left.len &&
            left.rightChar === right.leftChar
        ) {
            res.prefix = left.len + right.prefix;
        }

        res.suffix = right.suffix;

        if (
            right.suffix === right.len &&
            left.rightChar === right.leftChar
        ) {
            res.suffix = left.suffix + right.len;
        }

        if (left.rightChar === right.leftChar) {
            res.best = Math.max(
                res.best,
                left.suffix + right.prefix
            );
        }

        return res;
    };

    const build = (node, l, r) => {
        if (l === r) {
            tree[node] = {
                leftChar: s[l],
                rightChar: s[l],
                len: 1,
                prefix: 1,
                suffix: 1,
                best: 1
            };
            return;
        }

        const mid = Math.floor((l + r) / 2);

        build(node * 2, l, mid);
        build(node * 2 + 1, mid + 1, r);

        tree[node] = merge(tree[node * 2], tree[node * 2 + 1]);
    };

    const update = (node, l, r, idx, char) => {
        if (l === r) {
            tree[node] = {
                leftChar: char,
                rightChar: char,
                len: 1,
                prefix: 1,
                suffix: 1,
                best: 1
            };
            return;
        }

        const mid = Math.floor((l + r) / 2);

        if (idx <= mid) {
            update(node * 2, l, mid, idx, char);
        } else {
            update(node * 2 + 1, mid + 1, r, idx, char);
        }

        tree[node] = merge(tree[node * 2], tree[node * 2 + 1]);
    };

    build(1, 0, n - 1);

    const answer = [];

    for (let i = 0; i < queryCharacters.length; i++) {
        update(
            1,
            0,
            n - 1,
            queryIndices[i],
            queryCharacters[i]
        );

        answer.push(tree[1].best);
    }

    return answer;
};
```

### Python3

```python
from typing import List

class Solution:
    def longestRepeating(self, s: str, queryCharacters: str, queryIndices: List[int]) -> List[int]:
        tree = [None] * (4 * len(s))

        def merge(left, right):
            left_char = left[0]
            right_char = right[1]

            length = left[2] + right[2]

            best = max(left[5], right[5])

            prefix = left[3]

            if left[3] == left[2] and left[1] == right[0]:
                prefix = left[2] + right[3]

            suffix = right[4]

            if right[4] == right[2] and left[1] == right[0]:
                suffix = left[4] + right[2]

            if left[1] == right[0]:
                best = max(best, left[4] + right[3])

            return (left_char, right_char, length, prefix, suffix, best)

        def build(node, l, r):
            if l == r:
                tree[node] = (s[l], s[l], 1, 1, 1, 1)
                return

            mid = (l + r) // 2

            build(node * 2, l, mid)
            build(node * 2 + 1, mid + 1, r)

            tree[node] = merge(tree[node * 2], tree[node * 2 + 1])

        def update(node, l, r, idx, char):
            if l == r:
                tree[node] = (char, char, 1, 1, 1, 1)
                return

            mid = (l + r) // 2

            if idx <= mid:
                update(node * 2, l, mid, idx, char)
            else:
                update(node * 2 + 1, mid + 1, r, idx, char)

            tree[node] = merge(tree[node * 2], tree[node * 2 + 1])

        build(1, 0, len(s) - 1)

        answer = []

        for i in range(len(queryCharacters)):
            update(
                1,
                0,
                len(s) - 1,
                queryIndices[i],
                queryCharacters[i]
            )

            answer.append(tree[1][5])

        return answer
```

### Go

```go
func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
 type Node struct {
  leftChar  byte
  rightChar byte
  len       int
  prefix    int
  suffix    int
  best      int
 }

 n := len(s)

 tree := make([]Node, 4*n)

 merge := func(left, right Node) Node {
  var res Node

  res.leftChar = left.leftChar
  res.rightChar = right.rightChar
  res.len = left.len + right.len

  res.best = left.best
  if right.best > res.best {
   res.best = right.best
  }

  res.prefix = left.prefix

  if left.prefix == left.len &&
   left.rightChar == right.leftChar {
   res.prefix = left.len + right.prefix
  }

  res.suffix = right.suffix

  if right.suffix == right.len &&
   left.rightChar == right.leftChar {
   res.suffix = left.suffix + right.len
  }

  if left.rightChar == right.leftChar {
   cross := left.suffix + right.prefix
   if cross > res.best {
    res.best = cross
   }
  }

  return res
 }

 var build func(int, int, int)

 build = func(node, l, r int) {
  if l == r {
   tree[node] = Node{
    leftChar:  s[l],
    rightChar: s[l],
    len:       1,
    prefix:    1,
    suffix:    1,
    best:      1,
   }
   return
  }

  mid := (l + r) / 2

  build(node*2, l, mid)
  build(node*2+1, mid+1, r)

  tree[node] = merge(tree[node*2], tree[node*2+1])
 }

 var update func(int, int, int, int, byte)

 update = func(node, l, r, idx int, c byte) {
  if l == r {
   tree[node] = Node{
    leftChar:  c,
    rightChar: c,
    len:       1,
    prefix:    1,
    suffix:    1,
    best:      1,
   }
   return
  }

  mid := (l + r) / 2

  if idx <= mid {
   update(node*2, l, mid, idx, c)
  } else {
   update(node*2+1, mid+1, r, idx, c)
  }

  tree[node] = merge(tree[node*2], tree[node*2+1])
 }

 build(1, 0, n-1)

 answer := make([]int, len(queryCharacters))

 for i := 0; i < len(queryCharacters); i++ {
  update(
   1,
   0,
   n-1,
   queryIndices[i],
   queryCharacters[i],
  )

  answer[i] = tree[1].best
 }

 return answer
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five languages. Only the syntax and data representation change.

### 1. Representing a Segment Tree Node

The first thing I need is a structure that represents a segment.

I store:

```text
leftChar
rightChar
len
prefix
suffix
best
```

The reason for storing both boundary characters is simple: when I merge two segments, I need to know whether the characters touching each other are equal.

The length is equally important.

Without `len`, I cannot reliably determine whether a complete segment consists of one repeated character.

For example:

```text
aba
```

has:

```text
prefix = 1
suffix = 1
best = 1
```

but:

```text
len = 3
```

So it is clearly not uniform.

### 2. Building a Leaf

When the segment contains only one character, every value is straightforward.

For:

```text
a
```

I have:

```text
leftChar = a
rightChar = a
len = 1
prefix = 1
suffix = 1
best = 1
```

There is no larger substring to consider because the segment contains exactly one position.

### 3. Merging Two Segments

Suppose the left segment and right segment are already solved.

I create a new node representing their combination.

The new first character is:

```text
left.leftChar
```

The new last character is:

```text
right.rightChar
```

The new length is:

```text
left.len + right.len
```

For the best value, I first consider the best result from each child:

```text
max(left.best, right.best)
```

Then I check the boundary.

If the characters touching the boundary are equal:

```text
left.rightChar == right.leftChar
```

then the suffix from the left and the prefix from the right form one continuous repeating substring.

So I also consider:

```text
left.suffix + right.prefix
```

### 4. Calculating the Prefix

Normally, the prefix belongs to the left child.

So I start with:

```text
prefix = left.prefix
```

But there is one special case.

Suppose the left segment is:

```text
aaaa
```

Its prefix is `4`, which is also its complete length.

Therefore:

```text
left.prefix == left.len
```

means every character in the left segment is the same.

If the right segment starts with the same character, I can extend the prefix:

```text
left.len + right.prefix
```

If the left segment is:

```text
aaba
```

then its prefix is only `2` while its length is `4`.

So I cannot extend the prefix across the entire left segment.

### 5. Calculating the Suffix

The suffix is symmetric.

Normally:

```text
suffix = right.suffix
```

But if the complete right segment contains one repeated character:

```text
right.suffix == right.len
```

and the boundary characters match, the suffix extends left:

```text
left.suffix + right.len
```

This gives the correct suffix for the merged segment.

### 6. Building the Entire Tree

I recursively split the original string into smaller segments.

Each internal node represents the combination of its two children.

The tree eventually has one root representing the complete string.

The root's `best` value is always the answer for the current version of the string.

### 7. Processing a Query

Each query changes exactly one character.

For example:

```text
s = babacc
index = 1
new character = b
```

The string becomes:

```text
bbbacc
```

I do not scan the entire string.

Instead, I go directly to index `1` in the segment tree and update its leaf.

Then I recalculate its parents.

Only `O(log n)` nodes are affected.

### 8. Why the Root Gives the Answer

Every node is responsible for a continuous part of the string.

The root is responsible for the entire string.

Therefore:

```text
root.best
```

is exactly the longest substring containing one repeating character in the complete current string.

This is why I can answer every query immediately after the update.

### 9. Language-specific Notes

#### C++

I use a `struct` for the segment tree node and a `vector<Node>` for the tree.

C++ is a natural fit here because the node structure is compact and recursive segment tree functions are straightforward.

#### Java

I use a `Node` class and a `Node[]` array.

Each tree position is initialized with a `Node` object before building the tree.

#### JavaScript

I represent each node as a JavaScript object containing the six required fields.

The segment tree itself is stored in a normal array.

#### Python3

I represent each node as a tuple:

```text
(leftChar, rightChar, len, prefix, suffix, best)
```

This keeps the implementation compact while still storing all required information.

#### Go

I use a `struct` for the node and a slice for the segment tree.

Go's value-based structs work well here because each node contains only a small fixed amount of data.

## Examples

### Example 1

Input:

```text
s = "babacc"
queryCharacters = "bcb"
queryIndices = [1, 3, 3]
```

Processing the queries:

After query 1:

```text
babacc
  |
  b
```

The string becomes:

```text
bbbacc
```

The longest repeating substring is:

```text
bbb
```

Answer:

```text
3
```

After query 2:

```text
bbbacc
   |
   c
```

The string becomes:

```text
bbbccc
```

The two largest repeating groups are:

```text
bbb
ccc
```

Both have length `3`.

Answer:

```text
3
```

After query 3:

```text
bbbccc
   |
   b
```

The string becomes:

```text
bbbbcc
```

Now the longest repeating substring is:

```text
bbbb
```

Answer:

```text
4
```

Final output:

```text
[3, 3, 4]
```

### Example 2

Input:

```text
s = "abyzz"
queryCharacters = "aa"
queryIndices = [2, 1]
```

After the first query:

```text
aba zz
```

The string becomes:

```text
abazz
```

The longest repeating substring is:

```text
zz
```

So the answer is:

```text
2
```

After the second query:

```text
aaazz
```

Now:

```text
aaa
```

is the longest repeating substring.

So the answer is:

```text
3
```

Final output:

```text
[2, 3]
```

### Example 3: Why Segment Length Matters

Consider a segment:

```text
aba
```

Its values are:

```text
prefix = 1
suffix = 1
best = 1
len = 3
```

Even though:

```text
prefix == suffix == best
```

the whole segment is not made of one repeated character.

The correct check is:

```text
prefix == len
```

or:

```text
suffix == len
```

depending on which side I am extending.

This detail prevents incorrect merging results.

## How to Use / Run Locally

LeetCode provides the class and function signature automatically, so the solution can be submitted directly there.

For local testing, create a source file for the language you want to test and add a small driver program around the solution class or function.

### C++

Save the file as:

```text
main.cpp
```

Compile it with:

```bash
g++ -std=c++17 -O2 -Wall main.cpp -o main
```

Run it with:

```bash
./main
```

### Java

Save the file as:

```text
Main.java
```

Compile it with:

```bash
javac Main.java
```

Run it with:

```bash
java Main
```

### JavaScript

Save the file as:

```text
main.js
```

Run it with:

```bash
node main.js
```

### Python3

Save the file as:

```text
main.py
```

Run it with:

```bash
python3 main.py
```

### Go

Save the file as:

```text
main.go
```

Run it with:

```bash
go run main.go
```

For local testing, I recommend using the examples from the problem statement first. They make it easy to verify whether the point updates and segment tree merge logic are working correctly.

## Notes & Optimizations

The biggest performance issue with this problem is repeatedly scanning the complete string.

That approach is simple but too slow for the maximum constraints.

A segment tree reduces every character update to `O(log n)`.

The most important implementation detail is checking whether an entire segment is uniform.

Do not assume that:

```text
prefix == suffix == best
```

means the segment contains one repeated character.

A string such as `aba` is a counterexample.

Instead, compare the repeating prefix or suffix against the segment's actual length.

Another useful point is that I do not need to store the entire substring inside every node. Doing that would make the tree unnecessarily large and expensive.

Only the boundary characters and repeating lengths are enough to merge two segments.

There are also alternative solutions using ordered sets or other interval-based structures, but the segment tree gives a clean and reliable `O(n + k log n)` solution.

The method also handles updates where the replacement character is the same as the current character. The leaf values stay effectively unchanged, and the result remains correct.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)

---
