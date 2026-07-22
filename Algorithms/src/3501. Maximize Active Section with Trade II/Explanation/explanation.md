# Maximize Active Section with Trade II

A beginner-friendly solution and explanation for **LeetCode 3501 - Maximize Active Section with Trade II**. This repository explains the idea behind the algorithm, walks through the reasoning step by step, and provides solutions in multiple programming languages. If you're preparing for coding interviews or improving your Data Structures and Algorithms (DSA) skills, this guide will help you understand both the intuition and the optimization used to solve the problem efficiently.

This solution uses **Run-Length Encoding (RLE)**, **Segment Compression**, **Sparse Table (Range Maximum Query)**, and **Binary String Processing** to answer each query in constant time after preprocessing. It is designed to handle the large input constraints efficiently while remaining easy to understand.

---

## Table of Contents

- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

You are given a binary string `s` where:

- `'1'` represents an active section.
- `'0'` represents an inactive section.

For every query, you are asked to work only on the specified substring `s[l...r]`. Before performing any operation, imagine that this substring is surrounded by two extra `'1'` characters. These extra characters only help determine valid trades and are **not counted** in the final answer.

You are allowed to perform **at most one trade**, which always consists of two consecutive steps:

1. Choose a contiguous block of `1`s that is completely surrounded by `0`s and convert the entire block into `0`s.
2. Then choose a contiguous block of `0`s that is completely surrounded by `1`s and convert the entire block into `1`s.

Your goal is to maximize the total number of active sections (`1`s) inside the queried substring after performing the best possible trade.

Since there can be up to `10^5` independent queries, checking every substring from scratch would be far too slow. The challenge is to preprocess the string so that each query can be answered extremely quickly.

This problem is a great example of combining **string algorithms**, **segment compression**, and **range query data structures** to achieve an efficient solution.

---

## Constraints

| Constraint | Value |
|------------|-------|
| `1 <= n == s.length <= 10^5` | Length of the binary string |
| `1 <= queries.length <= 10^5` | Number of queries |
| `s[i]` | Either `'0'` or `'1'` |
| `queries[i] = [l, r]` | Query range |
| `0 <= l <= r < n` | Valid substring indices |

### Why These Constraints Matter

- Both the string length and the number of queries can reach **100,000**.
- A solution that scans the substring for every query would be too slow.
- We need an algorithm that performs preprocessing once and answers each query in **constant time** (or very close to it).

---

## Intuition

At first, the trade operation looks complicated because it always happens in two separate steps. However, after working through a few examples, an interesting pattern appears.

Suppose we have three consecutive segments like this:

```text
000 111 000
```

The middle block of `1`s is surrounded by `0`s, so it can be selected for the first step.

After converting it into `0`s, the string becomes:

```text
000000000
```

Now all three zero segments have merged into one large contiguous block.

Since this merged zero block is surrounded by `1`s (either original ones or the augmented boundary `1`s mentioned in the problem), it immediately becomes a valid choice for the second step.

After converting it back into `1`s, the result is:

```text
111111111
```

Notice something interesting:

- We temporarily lose the middle `1` block.
- But we gain the left and right `0` blocks.

The size of the removed `1` block cancels out mathematically. The only thing that actually increases our answer is:

```text
Gain = Left Zero Block Length + Right Zero Block Length
```

This means we never need to simulate the trade directly.

Instead, for every valid `1` segment, we only need to know:

- How large is the zero block on its left?
- How large is the zero block on its right?

The best trade is simply the one that produces the largest gain.

Once this observation becomes clear, the entire problem becomes much easier to solve efficiently.

---

## Approach

Instead of treating every character individually, we first compress the binary string into contiguous segments of identical characters.

For example,

```text
100011100
```

becomes

```text
1 | 000 | 111 | 00
```

Each segment stores:

- whether it contains `0`s or `1`s,
- its starting index,
- its ending index,
- and its length.

Next, every position in the original string is mapped to its corresponding segment. This allows us to determine the segment containing any query boundary in constant time.

After building these segments, we compute the potential gain for every valid `1` segment.

For each such segment:

```text
gain = left zero length + right zero length
```

Since every query asks for the maximum gain inside a range, repeatedly scanning all segments would still be too expensive.

To solve this efficiently, we preprocess all gains using a **Sparse Table**, which supports **Range Maximum Queries (RMQ)**.

When processing a query:

1. Convert the query boundaries into segment indices.
2. Ignore ranges that are too small to contain a valid `0-1-0` pattern.
3. Check the middle segments using the Sparse Table.
4. Handle the first and last candidate segments separately because they may only be partially covered by the query.
5. Add the maximum possible gain to the original number of active sections.

This preprocessing allows every query to be answered in constant time while keeping the overall algorithm efficient enough for the given constraints.

---

## Data Structures Used

### 1. Arrays

Several arrays are used throughout the solution to store information efficiently.

Examples include:

- Segment type (`0` or `1`)
- Segment start index
- Segment end index
- Gain for each segment
- Logarithm table
- Position-to-segment mapping

Arrays provide constant-time access and keep the implementation simple.

---

### 2. Run-Length Encoded Segments

Instead of storing every character separately, consecutive identical characters are grouped into a single segment.

For example,

```text
111000011
```

becomes

```text
[111] [0000] [11]
```

This significantly reduces the number of elements that the algorithm works with.

---

### 3. Position-to-Segment Map

Every character index in the original string is mapped to its segment ID.

Example:

```text
Index : 0 1 2 3 4 5
Value : 1 1 0 0 1 1
SegID : 0 0 1 1 2 2
```

This allows query boundaries to be converted into segment indices in constant time.

---

### 4. Sparse Table

A Sparse Table is built over the gain array.

Its purpose is to answer:

> "What is the maximum gain inside this range of segments?"

in **O(1)** time.

This preprocessing is the key optimization that allows the solution to handle up to **100,000 queries** efficiently.

---

## Operations & Behavior Summary

The algorithm follows the same sequence for every test case.

### Step 1

Count the total number of active sections (`1`s) in the original string.

This becomes the base answer for every query.

---

### Step 2

Compress consecutive characters into alternating segments of `0`s and `1`s.

This reduces unnecessary work.

---

### Step 3

Store metadata for every segment:

- Segment type
- Start index
- End index
- Length

---

### Step 4

Create a lookup table that maps every character position to its corresponding segment.

This makes locating query boundaries extremely fast.

---

### Step 5

For every valid `1` segment, calculate:

```text
Gain = Left Zero Length + Right Zero Length
```

These values are stored for later use.

---

### Step 6

Build a Sparse Table over the gain array.

This allows maximum-gain queries to be answered instantly.

---

### Step 7

For each query:

- Find the starting and ending segments.
- Compute the maximum gain inside the range.
- Carefully handle edge segments that are only partially included.
- Add the gain to the original number of active sections.

---

### Step 8

Return the answer for every query in the order they were given.

---

## Complexity

| Operation | Time Complexity | Explanation |
|-----------|-----------------|-------------|
| Count total active sections | `O(n)` | Scan the binary string once. |
| Build compressed segments | `O(n)` | Every character belongs to exactly one segment. |
| Build position-to-segment mapping | `O(n)` | Each index is mapped once. |
| Compute gains | `O(n)` | Each segment is processed once. |
| Build Sparse Table | `O(n log n)` | Standard RMQ preprocessing. |
| Process one query | `O(1)` | Uses direct lookups and Range Maximum Query. |
| Overall Complexity | `O(n log n + q)` | Efficient for the maximum constraints. |

### Space Complexity

| Data Structure | Space |
|---------------|-------|
| Segment arrays | `O(n)` |
| Position mapping | `O(n)` |
| Gain array | `O(n)` |
| Sparse Table | `O(n log n)` |
| Overall Space | `O(n log n)` |

The additional memory is mainly used by the Sparse Table, which enables extremely fast query processing.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> maxActiveSectionsAfterTrade(string s, vector<vector<int>>& queries) {
        int n = s.length(); // Get the size of the original string
        int total_ones = 0; // Track the total number of 1s initially present in the entire string
        
        // Count all the '1's in the string since our trades will only ADD to this base count
        for (char c : s) {
            if (c == '1') total_ones++;
        }
        
        // Arrays to represent compressed segments of identical characters
        vector<int> type;    // Stores 0 for a 0-segment, 1 for a 1-segment
        vector<int> start;   // Stores the starting index of the segment in the original string
        vector<int> end_idx; // Stores the ending index of the segment in the original string
        
        // Group consecutive identical characters into segments
        for (int i = 0; i < n; ) {
            int j = i;
            // Advance j as long as characters match the start of the current segment
            while (j < n && s[j] == s[i]) {
                j++;
            }
            // Record the segment metadata
            type.push_back(s[i] - '0');
            start.push_back(i);
            end_idx.push_back(j - 1);
            i = j; // Move to the start of the next segment
        }
        
        int N = type.size(); // Total number of compressed segments
        
        // Lookup array to map any string index to its corresponding segment ID in O(1) time
        vector<int> pos_to_seg(n);
        for (int i = 0; i < N; i++) {
            for (int j = start[i]; j <= end_idx[i]; j++) {
                pos_to_seg[j] = i; 
            }
        }
        
        // Precalculate the maximum potential gain for every segment
        vector<int> ans(N, 0);
        for (int i = 1; i < N - 1; i++) {
            // Only 1-segments surrounded by 0s yield a gain
            if (type[i] == 1) {
                // The gain is exactly the length of the left 0-segment + the length of the right 0-segment
                ans[i] = (end_idx[i - 1] - start[i - 1] + 1) + (end_idx[i + 1] - start[i + 1] + 1);
            }
        }
        
        // Precompute logarithms for the Sparse Table RMQ
        vector<int> log_table(N + 1, 0);
        for (int i = 2; i <= N; i++) {
            log_table[i] = log_table[i / 2] + 1;
        }
        
        int K = log_table[N] + 1; // Maximum power of 2 needed
        // Build the Sparse Table to answer range maximum queries in O(1)
        vector<vector<int>> st(K, vector<int>(N, 0));
        
        // Base case: intervals of length 2^0 = 1
        for (int i = 0; i < N; i++) {
            st[0][i] = ans[i];
        }
        
        // Dynamic programming to build larger intervals from smaller ones
        for (int j = 1; j < K; j++) {
            for (int i = 0; i + (1 << j) <= N; i++) {
                st[j][i] = max(st[j - 1][i], st[j - 1][i + (1 << (j - 1))]);
            }
        }
        
        // Helper lambda to query the maximum gain in a range of segments
        auto queryRMQ = [&](int L_idx, int R_idx) {
            if (L_idx > R_idx) return 0; // Invalid range protection
            int j = log_table[R_idx - L_idx + 1];
            // Combine two overlapping intervals of size 2^j to cover the whole query range
            return max(st[j][L_idx], st[j][R_idx - (1 << j) + 1]);
        };
        
        // Helper lambda to manually evaluate edge segments that might be partially chopped by query boundaries
        auto eval = [&](int i, int L, int R, int segL, int segR) {
            // The 1-segment must be strictly inside the query bounds to have 0s on both sides
            if (i <= segL || i >= segR) return 0;
            // Ignore if it's not a 1-segment
            if (type[i] == 0) return 0;
            
            int left_len = 0;
            // If the left 0-segment crosses the query's left bound, truncate it
            if (i - 1 == segL) left_len = max(0, end_idx[i - 1] - L + 1);
            else left_len = end_idx[i - 1] - start[i - 1] + 1; // Otherwise take full length
            
            int right_len = 0;
            // If the right 0-segment crosses the query's right bound, truncate it
            if (i + 1 == segR) right_len = max(0, R - start[i + 1] + 1);
            else right_len = end_idx[i + 1] - start[i + 1] + 1; // Otherwise take full length
            
            return left_len + right_len;
        };
        
        vector<int> res; // Array to hold answers for all queries
        
        for (const auto& q : queries) {
            int L = q[0]; // Query left bound
            int R = q[1]; // Query right bound
            
            // Find which compressed segments L and R fall into
            int segL = pos_to_seg[L];
            int segR = pos_to_seg[R];
            
            // If the query spans less than 3 segments, no 0-1-0 trade is possible
            if (segR - segL < 2) {
                res.push_back(total_ones);
                continue;
            }
            
            int max_gain = 0;
            // Check the 1-segments closest to the query boundaries (they might be partially truncated)
            max_gain = max(max_gain, eval(segL + 1, L, R, segL, segR));
            max_gain = max(max_gain, eval(segR - 1, L, R, segL, segR));
            
            // For all segments safely trapped in the middle, their 0s are fully intact. Fast query them!
            if (segL + 2 <= segR - 2) {
                max_gain = max(max_gain, queryRMQ(segL + 2, segR - 2));
            }
            
            // Final answer is the baseline 1s plus the maximum extra 1s we squeezed out
            res.push_back(total_ones + max_gain);
        }
        
        return res;
    }
};
```

### Java

```java
class Solution {
    public List<Integer> maxActiveSectionsAfterTrade(String s, int[][] queries) {
        int n = s.length(); // Length of string
        int totalOnes = 0; // Base count of 1s in the entire string
        
        // Count base ones 
        for (int i = 0; i < n; i++) {
            if (s.charAt(i) == '1') {
                totalOnes++;
            }
        }
        
        // Use dynamic lists to store compressed segment data
        List<Integer> typeList = new ArrayList<>();
        List<Integer> startList = new ArrayList<>();
        List<Integer> endList = new ArrayList<>();
        
        // Group consecutive identical characters
        for (int i = 0; i < n; ) {
            int j = i;
            // Advance as long as the character matches the segment's start
            while (j < n && s.charAt(j) == s.charAt(i)) {
                j++;
            }
            // Save segment info
            typeList.add(s.charAt(i) - '0');
            startList.add(i);
            endList.add(j - 1);
            i = j; // Move to next segment
        }
        
        int N = typeList.size(); // Total segment count
        // Arrays are faster in Java, so cast our lists to primitive arrays
        int[] type = new int[N];
        int[] start = new int[N];
        int[] endIdx = new int[N];
        for (int i = 0; i < N; i++) {
            type[i] = typeList.get(i);
            start[i] = startList.get(i);
            endIdx[i] = endList.get(i);
        }
        
        // Create an index mapping for O(1) segment lookups
        int[] posToSeg = new int[n];
        for (int i = 0; i < N; i++) {
            for (int j = start[i]; j <= endIdx[i]; j++) {
                posToSeg[j] = i;
            }
        }
        
        // Precalculate maximum possible gain if we swap around a given 1-segment
        int[] ans = new int[N];
        for (int i = 1; i < N - 1; i++) {
            if (type[i] == 1) {
                // Formula: length of left 0-segment + length of right 0-segment
                ans[i] = (endIdx[i - 1] - start[i - 1] + 1) + (endIdx[i + 1] - start[i + 1] + 1);
            }
        }
        
        // Prepare log array for Sparse Table
        int[] logTable = new int[N + 1];
        for (int i = 2; i <= N; i++) {
            logTable[i] = logTable[i / 2] + 1;
        }
        
        int K = logTable[N] + 1; // Max power of 2
        // Build Sparse Table for O(1) Range Maximum Queries
        int[][] st = new int[K][N];
        for (int i = 0; i < N; i++) {
            st[0][i] = ans[i];
        }
        
        for (int j = 1; j < K; j++) {
            for (int i = 0; i + (1 << j) <= N; i++) {
                st[j][i] = Math.max(st[j - 1][i], st[j - 1][i + (1 << (j - 1))]);
            }
        }
        
        List<Integer> res = new ArrayList<>(); // Store results
        
        // Process each query independently
        for (int[] q : queries) {
            int L = q[0]; // Query left bound
            int R = q[1]; // Query right bound
            
            // Map character indices to segment IDs
            int segL = posToSeg[L];
            int segR = posToSeg[R];
            
            // Need at least a 0-1-0 segment structure
            if (segR - segL < 2) {
                res.add(totalOnes);
                continue;
            }
            
            int maxGain = 0;
            // Check boundary-adjacent 1-segments (their 0s might be chopped off)
            maxGain = Math.max(maxGain, evaluateEdge(segL + 1, L, R, segL, segR, type, start, endIdx));
            maxGain = Math.max(maxGain, evaluateEdge(segR - 1, L, R, segL, segR, type, start, endIdx));
            
            // Check fully enclosed 1-segments using Sparse Table
            if (segL + 2 <= segR - 2) {
                int L_idx = segL + 2;
                int R_idx = segR - 2;
                int j = logTable[R_idx - L_idx + 1];
                int rmqVal = Math.max(st[j][L_idx], st[j][R_idx - (1 << j) + 1]);
                maxGain = Math.max(maxGain, rmqVal);
            }
            
            // Add max possible extra 1s to the initial total
            res.add(totalOnes + maxGain);
        }
        
        return res;
    }
    
    // Helper to evaluate partial segments manually
    private int evaluateEdge(int i, int L, int R, int segL, int segR, int[] type, int[] start, int[] endIdx) {
        // Exclude segments that touch edges because they lack surrounding 0s
        if (i <= segL || i >= segR) return 0;
        // Must be a 1-segment
        if (type[i] == 0) return 0;
        
        int leftLen = 0;
        // Truncate left 0-segment if it crosses L
        if (i - 1 == segL) leftLen = Math.max(0, endIdx[i - 1] - L + 1);
        else leftLen = endIdx[i - 1] - start[i - 1] + 1;
        
        int rightLen = 0;
        // Truncate right 0-segment if it crosses R
        if (i + 1 == segR) rightLen = Math.max(0, R - start[i + 1] + 1);
        else rightLen = endIdx[i + 1] - start[i + 1] + 1;
        
        return leftLen + rightLen; // Total gain
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
var maxActiveSectionsAfterTrade = function(s, queries) {
    let n = s.length; // Length of the original string
    let total_ones = 0; // Number of active sections before any trade
    
    // Count up the baseline 1s
    for (let i = 0; i < n; i++) {
        if (s[i] === '1') total_ones++;
    }
    
    // Arrays for RLE (Run-Length Encoding) of the string
    let type = [];
    let start = [];
    let end_idx = [];
    
    // Compress string into alternating blocks of 0s and 1s
    let i = 0;
    while (i < n) {
        let j = i;
        // Find the extent of the current block
        while (j < n && s[j] === s[i]) {
            j++;
        }
        // Save the metadata
        type.push(parseInt(s[i]));
        start.push(i);
        end_idx.push(j - 1);
        i = j; // Step forward
    }
    
    let N = type.length; // Total blocks
    
    // Map indices to block IDs to answer boundary queries instantly
    let pos_to_seg = new Int32Array(n);
    for (let i = 0; i < N; i++) {
        for (let j = start[i]; j <= end_idx[i]; j++) {
            pos_to_seg[j] = i;
        }
    }
    
    // Calculate full potential gains for each block
    let ans = new Int32Array(N);
    for (let i = 1; i < N - 1; i++) {
        if (type[i] === 1) {
            // Gain is the combined size of adjacent 0-blocks
            ans[i] = (end_idx[i - 1] - start[i - 1] + 1) + (end_idx[i + 1] - start[i + 1] + 1);
        }
    }
    
    // Precompute logs for Sparse Table
    let log_table = new Int32Array(N + 1);
    for (let i = 2; i <= N; i++) {
        log_table[i] = log_table[Math.floor(i / 2)] + 1;
    }
    
    let K = log_table[N] + 1;
    // Build the Sparse Table matrix
    let st = Array.from({ length: K }, () => new Int32Array(N));
    
    for (let i = 0; i < N; i++) {
        st[0][i] = ans[i];
    }
    
    for (let j = 1; j < K; j++) {
        for (let i = 0; i + (1 << j) <= N; i++) {
            st[j][i] = Math.max(st[j - 1][i], st[j - 1][i + (1 << (j - 1))]);
        }
    }
    
    // Function to retrieve max from Sparse Table in O(1)
    const queryRMQ = (L_q, R_q) => {
        if (L_q > R_q) return 0;
        let j = log_table[R_q - L_q + 1];
        return Math.max(st[j][L_q], st[j][R_q - (1 << j) + 1]);
    };
    
    // Function to evaluate segments hit by boundaries L or R
    const evalSeg = (idx, L, R, segL, segR) => {
        if (idx <= segL || idx >= segR) return 0; // Invalid because it lacks boundaries
        if (type[idx] === 0) return 0; // Only sacrifice 1-segments
        
        let left_len = 0;
        // Clamp the left 0-segment to query bounds
        if (idx - 1 === segL) left_len = Math.max(0, end_idx[idx - 1] - L + 1);
        else left_len = end_idx[idx - 1] - start[idx - 1] + 1;
        
        let right_len = 0;
        // Clamp the right 0-segment to query bounds
        if (idx + 1 === segR) right_len = Math.max(0, R - start[idx + 1] + 1);
        else right_len = end_idx[idx + 1] - start[idx + 1] + 1;
        
        return left_len + right_len;
    };
    
    let res = []; // Array to store outputs
    
    for (let q of queries) {
        let L = q[0];
        let R = q[1];
        
        let segL = pos_to_seg[L];
        let segR = pos_to_seg[R];
        
        // Ensure a trade pattern is physically possible
        if (segR - segL < 2) {
            res.push(total_ones);
            continue;
        }
        
        let max_gain = 0;
        // Test edge 1-segments
        max_gain = Math.max(max_gain, evalSeg(segL + 1, L, R, segL, segR));
        max_gain = Math.max(max_gain, evalSeg(segR - 1, L, R, segL, segR));
        
        // Test middle 1-segments instantly
        if (segL + 2 <= segR - 2) {
            max_gain = Math.max(max_gain, queryRMQ(segL + 2, segR - 2));
        }
        
        // The answer is the initial 1s count plus the highest net addition
        res.push(total_ones + max_gain);
    }
    
    return res;
};
```

### Python3

```python
class Solution:
    def maxActiveSectionsAfterTrade(self, s: str, queries: List[List[int]]) -> List[int]:
        n = len(s) # Size of the string
        total_ones = s.count('1') # Base count of 1s without trades
        
        # Parallel lists to hold run-length encoded string segments
        type_arr = []
        start = []
        end_idx = []
        
        # Build segments of contiguous blocks of same character
        i = 0
        while i < n:
            j = i
            while j < n and s[j] == s[i]:
                j += 1
            # Append 0 or 1 identifier and its bounds
            type_arr.append(int(s[i]))
            start.append(i)
            end_idx.append(j - 1)
            i = j
            
        N = len(type_arr) # Count of blocks
        
        # Maps absolute string index to block ID so we don't have to binary search later
        pos_to_seg = [0] * n
        for idx in range(N):
            for j in range(start[idx], end_idx[idx] + 1):
                pos_to_seg[j] = idx
                
        # Calculate optimal full gain for converting any 1-block to 0s
        ans = [0] * N
        for idx in range(1, N - 1):
            if type_arr[idx] == 1:
                # The net gain is literally the sum of the adjacent 0-blocks' lengths
                ans[idx] = (end_idx[idx - 1] - start[idx - 1] + 1) + (end_idx[idx + 1] - start[idx + 1] + 1)
                
        # Setup table for Sparse Table initialization
        log_table = [0] * (N + 1)
        for idx in range(2, N + 1):
            log_table[idx] = log_table[idx // 2] + 1
            
        K = log_table[N] + 1
        st = [[0] * N for _ in range(K)] # Sparse Table grid
        
        # Load base layer
        for idx in range(N):
            st[0][idx] = ans[idx]
            
        # Compute RMQ tree layers
        for j in range(1, K):
            for idx in range(N - (1 << j) + 1):
                st[j][idx] = max(st[j - 1][idx], st[j - 1][idx + (1 << (j - 1))])
                
        # O(1) query function using max of overlapping intervals
        def query_rmq(L_q, R_q):
            if L_q > R_q:
                return 0
            j = log_table[R_q - L_q + 1]
            return max(st[j][L_q], st[j][R_q - (1 << j) + 1])
            
        # Function to test 1-segments abutting query edges where truncation occurs
        def eval_seg(idx, L, R, segL, segR):
            # Abort if not strictly nested or not a 1-block
            if idx <= segL or idx >= segR: return 0
            if type_arr[idx] == 0: return 0
            
            # Left 0-block might spill past the query start L, clip it if necessary
            if idx - 1 == segL:
                left_len = max(0, end_idx[idx - 1] - L + 1)
            else:
                left_len = end_idx[idx - 1] - start[idx - 1] + 1
                
            # Right 0-block might spill past the query end R, clip it if necessary
            if idx + 1 == segR:
                right_len = max(0, R - start[idx + 1] + 1)
            else:
                right_len = end_idx[idx + 1] - start[idx + 1] + 1
                
            return left_len + right_len

        res = []
        for L, R in queries:
            # Pinpoint the blocks L and R land in
            segL = pos_to_seg[L]
            segR = pos_to_seg[R]
            
            # If the span is small, there's no way to sandwich a 1 between two 0s
            if segR - segL < 2:
                res.append(total_ones)
                continue
                
            max_gain = 0
            # Test immediate boundaries manually since they're vulnerable to truncation
            max_gain = max(max_gain, eval_seg(segL + 1, L, R, segL, segR))
            max_gain = max(max_gain, eval_seg(segR - 1, L, R, segL, segR))
            
            # Fetch highest yield from safe middle blocks instantly
            if segL + 2 <= segR - 2:
                max_gain = max(max_gain, query_rmq(segL + 2, segR - 2))
                
            # Aggregate and append score
            res.append(total_ones + max_gain)
            
        return res
```

### Go

```go
func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
    n := len(s) // Original string length
    totalOnes := 0 // Counter for default number of '1's in the string
    
    // Establishing the baseline score
    for i := 0; i < n; i++ {
        if s[i] == '1' {
            totalOnes++
        }
    }
    
    // Arrays tracking block contents
    typeArr := []int{}
    start := []int{}
    endIdx := []int{}
    
    // Condense repeating string characters into metadata blocks
    for i := 0; i < n; {
        j := i
        // Keep marching j while characters are uniform
        for j < n && s[j] == s[i] {
            j++
        }
        val := 0
        if s[i] == '1' {
            val = 1
        }
        typeArr = append(typeArr, val)
        start = append(start, i)
        endIdx = append(endIdx, j-1)
        i = j // Snap i to the end
    }
    
    N := len(typeArr) // Length of compressed structure
    
    // Creates a lookup table translating raw character index into block ID
    posToSeg := make([]int, n)
    for i := 0; i < N; i++ {
        for j := start[i]; j <= endIdx[i]; j++ {
            posToSeg[j] = i
        }
    }
    
    // Precompute full theoretical yield for trading any given 1-segment
    ans := make([]int, N)
    for i := 1; i < N-1; i++ {
        if typeArr[i] == 1 {
            // Net addition is the sum of enclosing 0-segments
            ans[i] = (endIdx[i-1] - start[i-1] + 1) + (endIdx[i+1] - start[i+1] + 1)
        }
    }
    
    // Populate logarithmic exponents for RMQ table sizing
    logTable := make([]int, N+1)
    for i := 2; i <= N; i++ {
        logTable[i] = logTable[i/2] + 1
    }
    
    K := logTable[N] + 1
    st := make([][]int, K)
    for i := 0; i < K; i++ {
        st[i] = make([]int, N)
    }
    
    // Fill the 0th layer of the Sparse Table
    for i := 0; i < N; i++ {
        st[0][i] = ans[i]
    }
    
    // Fill remaining layers by combining overlapping powers of 2
    for j := 1; j < K; j++ {
        for i := 0; i+(1<<j) <= N; i++ {
            a := st[j-1][i]
            b := st[j-1][i+(1<<(j-1))]
            if b > a {
                st[j][i] = b
            } else {
                st[j][i] = a
            }
        }
    }
    
    // Function returning range maximum in O(1)
    queryRMQ := func(L, R int) int {
        if L > R {
            return 0
        }
        j := logTable[R-L+1]
        a := st[j][L]
        b := st[j][R-(1<<j)+1]
        if b > a {
            return b
        }
        return a
    }
    
    maxFn := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    
    // Custom evaluation logic for blocks that are cut off by L or R edges
    eval := func(i, L, R, segL, segR int) int {
        // Drop invalid or out-of-bounds queries
        if i <= segL || i >= segR {
            return 0
        }
        if typeArr[i] == 0 {
            return 0
        }
        
        leftLen := 0
        // Snip the left 0-segment exactly at L if it crosses over
        if i-1 == segL {
            leftLen = maxFn(0, endIdx[i-1]-L+1)
        } else {
            leftLen = endIdx[i-1] - start[i-1] + 1
        }
        
        rightLen := 0
        // Snip the right 0-segment exactly at R if it crosses over
        if i+1 == segR {
            rightLen = maxFn(0, R-start[i+1]+1)
        } else {
            rightLen = endIdx[i+1] - start[i+1] + 1
        }
        
        return leftLen + rightLen
    }
    
    res := make([]int, len(queries))
    // Loop over queries
    for idx, q := range queries {
        L := q[0]
        R := q[1]
        
        segL := posToSeg[L]
        segR := posToSeg[R]
        
        // If there isn't enough room to pull off a trade, record baseline score
        if segR-segL < 2 {
            res[idx] = totalOnes
            continue
        }
        
        maxGain := 0
        // Only evaluate blocks immediately touching bounds manually
        maxGain = maxFn(maxGain, eval(segL+1, L, R, segL, segR))
        maxGain = maxFn(maxGain, eval(segR-1, L, R, segL, segR))
        
        // Let the Sparse Table figure out the highest gain among fully encased segments
        if segL+2 <= segR-2 {
            maxGain = maxFn(maxGain, queryRMQ(segL+2, segR-2))
        }
        
        // Store total calculated yield
        res[idx] = totalOnes + maxGain
    }
    
    return res
}
```

---

# Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Although the syntax changes from one language to another, the algorithm remains exactly the same in every implementation. Each solution follows the same sequence of operations and produces identical results. The only differences are language-specific syntax and the data structures provided by each language.

Below is a detailed walkthrough of the reasoning behind the implementation.

---

## Common Algorithm Explanation

### Step 1: Count the Initial Active Sections

The very first thing we do is count the number of `'1'` characters in the original string.

```text
s = "100101"
```

Number of active sections:

```text
4
```

This becomes our baseline answer.

Every trade simply adds some extra active sections on top of this count.

---

### Step 2: Compress the String into Segments

Instead of processing every character individually, consecutive identical characters are grouped together.

For example,

```text
1110001100
```

becomes

```text
[111]
[000]
[11]
[00]
```

Each segment stores:

- Segment type (`0` or `1`)
- Starting index
- Ending index
- Length

Working with segments instead of characters greatly reduces unnecessary processing.

---

### Step 3: Map Every Position to Its Segment

Each character index is assigned the ID of the segment it belongs to.

Example:

```text
Index : 0 1 2 3 4 5 6
Value : 1 1 0 0 0 1 1
SegID : 0 0 1 1 1 2 2
```

Now if a query starts at index `5`, we instantly know it belongs to segment `2`.

This lookup happens in constant time.

---

### Step 4: Compute the Gain of Every Valid 1 Segment

This is the most important observation in the entire problem.

Suppose we have

```text
00011100
```

The middle block of `1`s is surrounded by two zero blocks.

Its gain is simply

```text
Left Zero Length + Right Zero Length
```

Example

```text
0001110000
```

Left zero length:

```text
3
```

Right zero length:

```text
4
```

Gain:

```text
7
```

Notice that the size of the middle `1` block is irrelevant.

Only the neighboring zero blocks matter.

Every valid `1` segment receives one precomputed gain value.

---

### Step 5: Build a Sparse Table

Many queries ask the same question:

> Which segment inside this range has the maximum gain?

Scanning every segment for every query would be far too slow.

Instead, we preprocess all gains into a Sparse Table.

Once built, the maximum gain inside any range can be found in constant time.

The preprocessing costs more time initially but makes every query extremely fast.

---

### Step 6: Process Every Query

For each query:

1. Convert both query boundaries into segment IDs.
2. Check whether a valid `0-1-0` pattern even exists.
3. Use the Sparse Table to find the best middle segment.
4. Manually evaluate the boundary segments because they may only be partially covered.
5. Choose the largest gain.
6. Add it to the original number of active sections.

The computed value becomes the final answer for that query.

---

## C++ Explanation

The C++ implementation closely follows the algorithm described above while taking advantage of STL containers for speed and flexibility.

### Counting Active Sections

The solution first scans the string once and counts every `'1'`.

This value never changes and acts as the base score.

---

### Building Segments

Three vectors are created.

```text
type
start
end
```

Each index represents one compressed segment.

For example,

```text
type[i]
```

stores whether the segment contains `0`s or `1`s.

```text
start[i]
```

stores the first index of that segment.

```text
end[i]
```

stores the last index.

---

### Creating Position Lookup

Another vector stores the segment ID for every character position.

Instead of searching for the segment every time a query arrives, we simply perform

```text
segment = positionToSegment[index]
```

This takes constant time.

---

### Computing Gains

Only `1` segments that have both neighboring zero segments are considered.

For each one,

```text
gain =
length(left zero)
+
length(right zero)
```

The value is stored for future queries.

---

### Sparse Table Construction

The gain array is used to build the Sparse Table.

Each level stores maximum values for increasingly larger intervals.

Later, any range maximum query can be answered using only two overlapping intervals.

This keeps query processing extremely fast.

---

### Handling Query Boundaries

Middle segments are fully contained inside the query and can safely use the Sparse Table.

However, the first and last candidate segments may only be partially included.

Their neighboring zero blocks might extend outside the query.

A helper function computes the actual visible lengths of those zero blocks before calculating the gain.

---

### Producing the Final Answer

For every query, the algorithm computes

```text
Answer =
Initial Active Sections
+
Maximum Gain
```

This value is appended to the result vector.

After all queries have been processed, the vector is returned.

---

## Java Explanation

The Java implementation follows exactly the same logic as the C++ solution.

The main differences come from Java's standard library and object-oriented syntax.

### Counting Initial Active Sections

The algorithm scans the string once using `charAt()`.

Every `'1'` contributes to the baseline answer.

---

### Compressing the String

Instead of processing characters individually, consecutive identical characters are grouped together.

Initially, the implementation stores segment information inside `ArrayList<Integer>` objects.

Once compression is complete, these lists are converted into primitive arrays for better performance.

---

### Position Mapping

An integer array maps every character position to its segment.

Whenever a query arrives, the starting and ending segments are found immediately without performing any searches.

---

### Precomputing Gains

Each valid `1` segment calculates

```text
Left Zero Length
+
Right Zero Length
```

The computed gain is stored inside an array.

This preprocessing removes the need to recompute gains for every query.

---

### Building the Sparse Table

Java stores the Sparse Table as a two-dimensional integer array.

Each level combines results from the previous level until every interval size has been processed.

Once finished, range maximum queries become constant-time operations.

---

### Processing Queries

Every query follows the same sequence:

- Convert boundaries into segment IDs.
- Check if a valid trade is possible.
- Evaluate edge segments separately.
- Query the Sparse Table for middle segments.
- Choose the largest gain.

---

### Returning Results

Each answer is added to an `ArrayList<Integer>`.

After all queries are processed, the completed list is returned as the final output.

The Java version produces exactly the same answers as the C++ implementation while using Java's collections and arrays.

## JavaScript Explanation

The JavaScript implementation follows the same algorithm as the C++ and Java versions. The biggest difference is that JavaScript uses dynamic arrays and typed arrays instead of vectors or primitive arrays.

Even though JavaScript is an interpreted language, the algorithm remains highly efficient because every expensive computation is completed during preprocessing.

### Counting Active Sections

The solution starts by scanning the binary string once.

Every time it encounters a `'1'`, it increases the counter.

For example,

```text
1011010
```

The initial active section count is

```text
4
```

This value becomes the starting point for every query.

---

### Compressing the String

Instead of treating every character separately, the algorithm groups consecutive identical characters into segments.

Example:

```text
111000011
```

becomes

```text
[111]
[0000]
[11]
```

Each segment stores:

- Type (`0` or `1`)
- Starting position
- Ending position

This compressed representation dramatically reduces the number of elements that later computations must process.

---

### Building the Position Lookup

JavaScript creates a lookup array that maps every character position to its segment ID.

Example:

```text
Index : 0 1 2 3 4 5
SegID : 0 0 1 1 2 2
```

Whenever a query begins or ends, the corresponding segment can be found instantly.

No searching is required.

---

### Calculating Segment Gains

Every valid `1` segment calculates one value:

```text
Gain =
Left Zero Length
+
Right Zero Length
```

This value represents the maximum increase in active sections if that segment is selected during the trade.

Each gain is stored for later use.

---

### Sparse Table Construction

The gain array never changes after preprocessing.

Because of that, a Sparse Table is the perfect data structure.

Once constructed, finding the maximum gain inside any interval becomes an O(1) operation.

This is the reason the solution remains fast even when handling up to one hundred thousand queries.

---

### Query Processing

Each query performs the following operations:

- Find the corresponding starting and ending segments.
- Verify that a valid trade is possible.
- Evaluate edge segments manually.
- Use the Sparse Table for fully covered middle segments.
- Select the maximum gain.
- Add the gain to the original number of active sections.

The computed value is pushed into the result array.

---

### Returning the Answers

After every query has been processed, the result array contains the answer for each query in the same order they were received.

The function simply returns this array.

---

## Python3 Explanation

The Python implementation follows exactly the same logic while taking advantage of Python's simple syntax and built-in list operations.

Although Python is dynamically typed, the algorithm is still efficient because almost all expensive work happens during preprocessing.

### Counting Active Sections

Python makes this step very concise.

The solution counts the total number of `'1'` characters before processing any queries.

This value acts as the baseline score throughout the algorithm.

---

### Building Compressed Segments

The string is scanned from left to right.

Whenever the current character changes, a new segment begins.

For each segment, the algorithm stores:

- Segment type
- Starting index
- Ending index

For example,

```text
100011100
```

becomes

```text
1
000
111
00
```

Instead of working with nine characters, the algorithm now works with only four segments.

---

### Creating Position Mapping

Every index of the original string is assigned to exactly one segment.

This allows the algorithm to determine the segment containing any query boundary immediately.

Without this mapping, additional searching would be required for every query.

---

### Computing Gains

The solution loops through every `1` segment that has neighboring zero segments.

Its gain is computed as

```text
Left Zero Length
+
Right Zero Length
```

Each value is stored inside a list.

Later, these values become the input for the Sparse Table.

---

### Sparse Table

Python stores the Sparse Table as a two-dimensional list.

Each row represents intervals whose lengths are powers of two.

This preprocessing enables constant-time range maximum queries throughout the remaining execution.

---

### Processing Queries

For each query, the algorithm:

1. Finds the segment containing the left boundary.
2. Finds the segment containing the right boundary.
3. Evaluates partially covered edge segments.
4. Uses the Sparse Table for fully covered middle segments.
5. Selects the maximum gain.
6. Adds that gain to the baseline active section count.

This sequence guarantees the correct answer while remaining extremely efficient.

---

### Returning the Result

Each computed answer is appended to a list.

Once every query has been processed, the completed list is returned.

---

## Go Explanation

The Go implementation uses slices to represent arrays while following exactly the same algorithm as the other language versions.

Because Go provides fast memory access and efficient slices, it performs particularly well for this type of preprocessing problem.

### Counting Active Sections

The algorithm begins by scanning the string once.

Every `'1'` increases the baseline count.

This value never changes and becomes the starting score for every query.

---

### Segment Compression

The binary string is divided into consecutive blocks of identical characters.

Example:

```text
11000011100
```

becomes

```text
11
0000
111
00
```

For every block, the implementation stores:

- Segment type
- Starting index
- Ending index

This compressed representation keeps the algorithm compact and efficient.

---

### Position Lookup

A slice maps every character position to its segment.

Whenever a query arrives, both boundaries can immediately be converted into segment IDs.

This avoids repeated searches and keeps query processing fast.

---

### Gain Calculation

Each valid `1` segment computes

```text
Gain =
Left Zero Length
+
Right Zero Length
```

These gains are calculated only once during preprocessing.

After that, they are reused for every query.

---

### Sparse Table

The gain array is converted into a Sparse Table.

Each level stores the maximum value for intervals whose lengths are powers of two.

Later, finding the maximum gain inside any interval requires only two table lookups.

---

### Processing Each Query

Every query follows the same workflow.

First, the algorithm identifies the starting and ending segments.

Next, it evaluates edge segments separately because they may not be fully contained inside the query.

Then, it retrieves the maximum gain for all fully covered middle segments using the Sparse Table.

Finally, it combines the maximum gain with the original active section count.

---

### Returning the Output

The answer for every query is stored inside a result slice.

After all queries have been processed, the slice is returned.

---

## Why All Five Implementations Behave the Same

Although the syntax changes from language to language, every implementation follows the same sequence of operations.

1. Count the initial number of active sections.
2. Compress the binary string into alternating segments.
3. Map every character position to its corresponding segment.
4. Compute the gain for every valid `1` segment.
5. Build a Sparse Table for fast range maximum queries.
6. Process every query using the same logic.
7. Return the maximum possible number of active sections.

Because the underlying algorithm never changes, every implementation has identical performance characteristics.

| Language | Time Complexity | Space Complexity |
|----------|-----------------|------------------|
| C++ | `O(n log n + q)` | `O(n log n)` |
| Java | `O(n log n + q)` | `O(n log n)` |
| JavaScript | `O(n log n + q)` | `O(n log n)` |
| Python3 | `O(n log n + q)` | `O(n log n)` |
| Go | `O(n log n + q)` | `O(n log n)` |

No matter which programming language you choose, the solution remains efficient enough to handle the maximum input constraints and is well suited for competitive programming and technical interviews.

## Examples

### Example 1

**Input**

```text
s = "1001001"

queries = [[0,6]]
```

**Output**

```text
[6]
```

**Explanation**

The original string contains four active sections.

```text
1001001
```

Compressing the string gives:

```text
1 | 00 | 1 | 00 | 1
```

If we sacrifice the middle `1` segment, both neighboring zero segments merge into one larger block. Turning that merged block into `1`s gives the highest possible increase.

The best gain is:

```text
2 + 2 = 4
```

After applying the optimal trade, the maximum number of active sections becomes:

```text
6
```

---

### Example 2

**Input**

```text
s = "111111"

queries = [[0,5]]
```

**Output**

```text
[6]
```

**Explanation**

The string contains only one segment.

```text
111111
```

There is no `0-1-0` pattern.

Since no valid trade exists, the answer remains unchanged.

---

### Example 3

**Input**

```text
s = "001110011100"

queries = [
    [0,11],
    [2,9]
]
```

**Output**

```text
[10,8]
```

**Explanation**

For the first query, several valid trades are possible.

The algorithm evaluates every valid `1` segment and selects the one that provides the largest gain.

For the second query, only the segments completely inside the selected range are considered.

Because every gain has already been precomputed and stored inside the Sparse Table, each query is answered in constant time.

---

## How to Use / Run Locally

Clone the repository first.

```bash
git clone https://github.com/your-username/your-repository.git
```

Move into the project directory.

```bash
cd your-repository
```

---

### C++

Compile the source file.

```bash
g++ solution.cpp -o solution
```

Run the executable.

```bash
./solution
```

---

### Java

Compile the Java source.

```bash
javac Solution.java
```

Run the program.

```bash
java Solution
```

---

### JavaScript

Run the solution using Node.js.

```bash
node solution.js
```

---

### Python3

Run the Python solution.

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Compile and execute the Go solution.

```bash
go run solution.go
```

or build an executable.

```bash
go build solution.go
```

Then run it.

```bash
./solution
```

---

## Notes & Optimizations

- The binary string is compressed into contiguous segments before answering any queries.
- Segment compression significantly reduces unnecessary work by treating consecutive identical characters as a single unit.
- Every character position is mapped to its corresponding segment, allowing query boundaries to be located instantly.
- The gain for every valid `1` segment is calculated only once during preprocessing.
- A Sparse Table is used to answer Range Maximum Queries (RMQ) in constant time.
- Only the first and last candidate segments of each query require special handling because they may be partially covered.
- Fully contained middle segments are answered directly using the Sparse Table.
- This preprocessing strategy allows the solution to efficiently handle up to **100,000** queries.
- The algorithm avoids repeatedly scanning substrings, making it suitable for competitive programming and coding interviews.
- Overall complexity is **O(n log n + q)**, which satisfies the problem constraints comfortably.
- This approach is significantly faster than checking every possible trade for each query independently.

---

## Author

**Md Aarzoo Islam**

Software Engineer | Competitive Programming Enthusiast | Open Source Contributor

Instagram: <https://www.instagram.com/code.with.aarzoo/>

If you found this repository helpful, consider giving it a **Star**. It helps others discover the project and motivates future contributions.
