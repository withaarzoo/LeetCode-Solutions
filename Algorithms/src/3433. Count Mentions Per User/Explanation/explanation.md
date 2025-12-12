# 3433. Count Mentions Per User

## Table of Contents

* ## Problem Summary

* ## Constraints

* ## Intuition

* ## Approach

* ## Data Structures Used

* ## Operations & Behavior Summary

* ## Complexity

* ## Multi-language Solutions

  * ### C++

  * ### Java

  * ### JavaScript

  * ### Python3

  * ### Go

* ## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

* ## Examples

* ## How to use / Run locally

* ## Notes & Optimizations

* ## Author

  * [Md Aarzoo Islam](https://bento.me/withaarzoo)

---

## Problem Summary

I am given the total number of users `numberOfUsers` and a list of `events`. Each event is one of:

1. **MESSAGE**: `["MESSAGE", timestamp, "mentions_string"]`

   * `mentions_string` contains tokens separated by single spaces. Tokens can be:

     * `idX` (e.g., `id0`, `id5`) — mention that specific user (counts even if offline).
     * `ALL` — mention every user (including offline).
     * `HERE` — mention every user who is currently online.
   * A message can contain multiple tokens and duplicates — each token is counted separately.

2. **OFFLINE**: `["OFFLINE", timestamp, "id"]`

   * When a user `id` goes offline at timestamp `t`, they stay offline for exactly 60 time units and automatically become online again at time `t + 60`.
   * It's guaranteed that the user referenced by an `OFFLINE` event is online when the `OFFLINE` occurs.

All users are initially online. For events that share the same timestamp:

* First process automatic state changes (users whose offline period ended at or before that timestamp).
* Then apply `OFFLINE` events at that timestamp.
* Then process `MESSAGE` events at that timestamp.

Return an array `mentions` of length `numberOfUsers` where `mentions[i]` is the number of times user `i` is mentioned across all `MESSAGE` events.

---

## Constraints

* `1 <= numberOfUsers <= 100`
* `1 <= events.length <= 100`
* Each `events[i].length == 3`
* `events[i][0]` ∈ {`"MESSAGE"`, `"OFFLINE"`}
* `1 <= int(events[i][1]) <= 1e5` (timestamps)
* Number of `id<number>` mentions in a single `MESSAGE` ≤ 100
* `0 <= <number> <= numberOfUsers - 1`
* User referenced in an `OFFLINE` event is guaranteed to be online at event time.

---

## Intuition

I thought: Mentions only depend on who is online when a `MESSAGE` is processed. So time must be handled in order. Also rules say all state changes at a timestamp must happen before processing `MESSAGE` events at that timestamp.

So I group events by timestamp and then for each timestamp:

1. Bring users back online whose offline period ended at or before this time.
2. Apply any `OFFLINE` events (they take effect before messages at this timestamp).
3. Process `MESSAGE` events using current online/offline state.

For tokens:

* `ALL` → increment all users
* `HERE` → increment only users currently online
* `idX` → increment user X (regardless of online/offline)

I count duplicates (if token appears multiple times).

---

## Approach

1. Group events by timestamp (map `timestamp -> list of events`).
2. Maintain three arrays:

   * `mentions[numberOfUsers]` — final counts.
   * `isOnline[numberOfUsers]` — boolean, initially all `true`.
   * `offlineUntil[numberOfUsers]` — timestamp when user automatically returns online (0 if online).
3. Iterate timestamps ascending:

   * For each timestamp `t`:

     * For every user, if `!isOnline[i]` and `offlineUntil[i] <= t`, set `isOnline[i] = true` and `offlineUntil[i] = 0`.
     * Apply `OFFLINE` events at `t`: set `isOnline[id] = false` and `offlineUntil[id] = t + 60`.
     * Apply `MESSAGE` events at `t`: split `mentions_string` into tokens and handle `ALL`, `HERE`, and `idX` accordingly.
4. Return `mentions`.

This ensures correct handling of simultaneous events and the 60-time-unit offline windows.

---

## Data Structures Used

* Map/TreeMap / ordered map to group events by timestamp (so we can process timestamps in ascending order).
* Arrays / vectors for `mentions`, `isOnline`, and `offlineUntil`.
* For parsing tokens: simple string splitting (`split`, `stringstream`, `strings.Fields`, etc.).

---

## Operations & Behavior Summary

* Bring expired offline users back online (checked before any `OFFLINE` or `MESSAGE` at same timestamp).
* Apply `OFFLINE` events at a timestamp — users become offline immediately and set `offlineUntil = t + 60`.
* Process `MESSAGE` events at that timestamp using the current `isOnline` state.
* `ALL` counts everyone; `HERE` counts only online users; `idX` counts that user (even if offline).
* Duplicate tokens are treated as separate mentions.

---

## Complexity

* **Time Complexity:** `O(E * U)` in worst case, where:

  * `E` = number of events (≤ 100)
  * `U` = number of users (≤ 100)
    Explanation: For tokens `ALL` or `HERE`, we iterate all `U` users. All other operations are O(1) per token. With problem constraints this is fast. If data were much bigger, we'd optimize repeated `ALL`/`HERE` by maintaining counts of online users, but constraints make `O(E * U)` acceptable.
* **Space Complexity:** `O(U + E)` to store `mentions`, `isOnline`, `offlineUntil`, and grouped events.

---

## Multi-language Solutions

Below are complete and readable implementations for C++, Java, JavaScript, Python3, and Go. Each solution follows the approach above and is commented.

---

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

/*
 Solution for 3433. Count Mentions Per User

 Approach:
 - Group events by timestamp (map<int, vector<event>>).
 - For each timestamp in order:
     1) Bring users back online whose offlineUntil <= timestamp.
     2) Apply OFFLINE events at this timestamp.
     3) Process MESSAGE events at this timestamp.
 - Return mentions vector.
*/

class Solution {
public:
    vector<int> countMentions(int numberOfUsers, vector<vector<string>>& events) {
        // Group events by timestamp (sorted by key).
        map<int, vector<vector<string>>> byTime;
        for (auto &ev : events) {
            int t = stoi(ev[1]);
            byTime[t].push_back(ev);
        }

        vector<int> mentions(numberOfUsers, 0);
        vector<bool> isOnline(numberOfUsers, true);
        vector<int> offlineUntil(numberOfUsers, 0);

        for (auto &entry : byTime) {
            int t = entry.first;
            auto &evs = entry.second;

            // 1) Expirations: bring users back online if their offline period ended.
            for (int i = 0; i < numberOfUsers; ++i) {
                if (!isOnline[i] && offlineUntil[i] <= t) {
                    isOnline[i] = true;
                    offlineUntil[i] = 0;
                }
            }

            // 2) Apply OFFLINE events first at timestamp t.
            for (auto &ev : evs) {
                if (ev[0] == "OFFLINE") {
                    int id = stoi(ev[2]);
                    isOnline[id] = false;
                    offlineUntil[id] = t + 60;
                }
            }

            // 3) Process MESSAGE events.
            for (auto &ev : evs) {
                if (ev[0] != "MESSAGE") continue;
                string mentionsStr = ev[2];
                stringstream ss(mentionsStr);
                string token;
                while (ss >> token) {
                    if (token == "ALL") {
                        for (int i = 0; i < numberOfUsers; ++i) mentions[i]++;
                    } else if (token == "HERE") {
                        for (int i = 0; i < numberOfUsers; ++i)
                            if (isOnline[i]) mentions[i]++;
                    } else if (token.rfind("id", 0) == 0) {
                        int id = stoi(token.substr(2));
                        if (0 <= id && id < numberOfUsers) mentions[id]++;
                    }
                }
            }
        }

        return mentions;
    }
};
```

---

### Java

```java
import java.util.*;

/*
 Solution for 3433. Count Mentions Per User

 Explanation:
 - Use TreeMap to sort events by timestamp.
 - Maintain mentions[], isOnline[], offlineUntil[].
 - For each timestamp: expire offline users, apply OFFLINE events, then process MESSAGE events.
*/

class Solution {
    public int[] countMentions(int numberOfUsers, List<List<String>> events) {
        TreeMap<Integer, List<List<String>>> byTime = new TreeMap<>();
        for (List<String> ev : events) {
            int t = Integer.parseInt(ev.get(1));
            byTime.computeIfAbsent(t, k -> new ArrayList<>()).add(ev);
        }

        int[] mentions = new int[numberOfUsers];
        boolean[] isOnline = new boolean[numberOfUsers];
        int[] offlineUntil = new int[numberOfUsers];
        Arrays.fill(isOnline, true);

        for (Map.Entry<Integer, List<List<String>>> entry : byTime.entrySet()) {
            int t = entry.getKey();
            List<List<String>> evs = entry.getValue();

            // 1) expirations
            for (int i = 0; i < numberOfUsers; ++i) {
                if (!isOnline[i] && offlineUntil[i] <= t) {
                    isOnline[i] = true;
                    offlineUntil[i] = 0;
                }
            }

            // 2) OFFLINE events
            for (List<String> ev : evs) {
                if (ev.get(0).equals("OFFLINE")) {
                    int id = Integer.parseInt(ev.get(2));
                    isOnline[id] = false;
                    offlineUntil[id] = t + 60;
                }
            }

            // 3) MESSAGE events
            for (List<String> ev : evs) {
                if (!ev.get(0).equals("MESSAGE")) continue;
                String[] tokens = ev.get(2).split("\\s+");
                for (String token : tokens) {
                    if (token.equals("ALL")) {
                        for (int i = 0; i < numberOfUsers; ++i) mentions[i]++;
                    } else if (token.equals("HERE")) {
                        for (int i = 0; i < numberOfUsers; ++i)
                            if (isOnline[i]) mentions[i]++;
                    } else if (token.startsWith("id")) {
                        int id = Integer.parseInt(token.substring(2));
                        if (id >= 0 && id < numberOfUsers) mentions[id]++;
                    }
                }
            }
        }

        return mentions;
    }
}
```

---

### JavaScript

```javascript
/**
 * Solution for 3433. Count Mentions Per User
 *
 * @param {number} numberOfUsers
 * @param {string[][]} events
 * @return {number[]}
 */
var countMentions = function(numberOfUsers, events) {
    const byTime = new Map();
    for (const ev of events) {
        const t = parseInt(ev[1], 10);
        if (!byTime.has(t)) byTime.set(t, []);
        byTime.get(t).push(ev);
    }

    const timestamps = Array.from(byTime.keys()).sort((a,b) => a - b);
    const mentions = Array(numberOfUsers).fill(0);
    const isOnline = Array(numberOfUsers).fill(true);
    const offlineUntil = Array(numberOfUsers).fill(0);

    for (const t of timestamps) {
        const evs = byTime.get(t);

        // 1) expirations
        for (let i = 0; i < numberOfUsers; ++i) {
            if (!isOnline[i] && offlineUntil[i] <= t) {
                isOnline[i] = true;
                offlineUntil[i] = 0;
            }
        }

        // 2) OFFLINE events
        for (const ev of evs) {
            if (ev[0] === "OFFLINE") {
                const id = parseInt(ev[2], 10);
                isOnline[id] = false;
                offlineUntil[id] = t + 60;
            }
        }

        // 3) MESSAGE events
        for (const ev of evs) {
            if (ev[0] !== "MESSAGE") continue;
            const tokens = ev[2].trim().split(/\s+/);
            for (const token of tokens) {
                if (token === "ALL") {
                    for (let i = 0; i < numberOfUsers; ++i) mentions[i]++;
                } else if (token === "HERE") {
                    for (let i = 0; i < numberOfUsers; ++i) if (isOnline[i]) mentions[i]++;
                } else if (token.startsWith("id")) {
                    const id = parseInt(token.slice(2), 10);
                    if (id >= 0 && id < numberOfUsers) mentions[id]++;
                }
            }
        }
    }

    return mentions;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def countMentions(self, numberOfUsers: int, events: List[List[str]]) -> List[int]:
        # Group events by timestamp
        by_time = {}
        for ev in events:
            t = int(ev[1])
            by_time.setdefault(t, []).append(ev)

        mentions = [0] * numberOfUsers
        is_online = [True] * numberOfUsers
        offline_until = [0] * numberOfUsers

        for t in sorted(by_time.keys()):
            evs = by_time[t]

            # 1) Expirations: bring users back online if offline_until <= t
            for i in range(numberOfUsers):
                if not is_online[i] and offline_until[i] <= t:
                    is_online[i] = True
                    offline_until[i] = 0

            # 2) OFFLINE events at t
            for ev in evs:
                if ev[0] == "OFFLINE":
                    id_ = int(ev[2])
                    is_online[id_] = False
                    offline_until[id_] = t + 60

            # 3) MESSAGE events at t
            for ev in evs:
                if ev[0] != "MESSAGE":
                    continue
                tokens = ev[2].split()
                for token in tokens:
                    if token == "ALL":
                        for i in range(numberOfUsers):
                            mentions[i] += 1
                    elif token == "HERE":
                        for i in range(numberOfUsers):
                            if is_online[i]:
                                mentions[i] += 1
                    elif token.startswith("id"):
                        id_ = int(token[2:])
                        if 0 <= id_ < numberOfUsers:
                            mentions[id_] += 1

        return mentions
```

---

### Go

```go
package main

import (
 "sort"
 "strconv"
 "strings"
)

/*
 Solution for 3433. Count Mentions Per User

 - Group events by timestamp (map[int][][]string)
 - Process timestamps in ascending order
 - Expire offline users, apply OFFLINE events, process MESSAGE events
*/

func countMentions(numberOfUsers int, events [][]string) []int {
 byTime := make(map[int][][]string)
 for _, ev := range events {
  t, _ := strconv.Atoi(ev[1])
  byTime[t] = append(byTime[t], ev)
 }

 timestamps := make([]int, 0, len(byTime))
 for t := range byTime {
  timestamps = append(timestamps, t)
 }
 sort.Ints(timestamps)

 mentions := make([]int, numberOfUsers)
 isOnline := make([]bool, numberOfUsers)
 offlineUntil := make([]int, numberOfUsers)
 for i := 0; i < numberOfUsers; i++ {
  isOnline[i] = true
 }

 for _, t := range timestamps {
  evs := byTime[t]

  // 1) expirations
  for i := 0; i < numberOfUsers; i++ {
   if !isOnline[i] && offlineUntil[i] <= t {
    isOnline[i] = true
    offlineUntil[i] = 0
   }
  }

  // 2) OFFLINE events
  for _, ev := range evs {
   if ev[0] == "OFFLINE" {
    id, _ := strconv.Atoi(ev[2])
    isOnline[id] = false
    offlineUntil[id] = t + 60
   }
  }

  // 3) MESSAGE events
  for _, ev := range evs {
   if ev[0] != "MESSAGE" {
    continue
   }
   tokens := strings.Fields(ev[2])
   for _, token := range tokens {
    if token == "ALL" {
     for i := 0; i < numberOfUsers; i++ {
      mentions[i]++
     }
    } else if token == "HERE" {
     for i := 0; i < numberOfUsers; i++ {
      if isOnline[i] {
       mentions[i]++
      }
     }
    } else if strings.HasPrefix(token, "id") {
     id, _ := strconv.Atoi(token[2:])
     if id >= 0 && id < numberOfUsers {
      mentions[id]++
     }
    }
   }
  }
 }

 return mentions
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the logic using the Python solution style because it maps easily to the other languages. The same sequence applies in each implementation.

1. **Group events by timestamp**

   * I create a dictionary/map with key = timestamp and value = list of events at that timestamp.
   * This lets me process all state changes and messages at the same time-point correctly.

2. **Initialize arrays**

   * `mentions = [0]*numberOfUsers` — final counts.
   * `is_online = [True]*numberOfUsers` — all users initially online.
   * `offline_until = [0]*numberOfUsers` — when user comes back online; 0 indicates already online.

3. **Process timestamps in ascending order**

   * For each timestamp `t`:

     * **Expire offline windows**: for each user `i`, if `not is_online[i]` and `offline_until[i] <= t`, then set `is_online[i] = True` and `offline_until[i] = 0`. This is important to do before other events at `t`.
     * **Apply OFFLINE events**: for each `OFFLINE` event at `t`, set `is_online[id] = False` and `offline_until[id] = t + 60`. This models the offline duration of exactly 60 time units.
     * **Process MESSAGE events**: for each `MESSAGE` event at `t`, split the `mentions_string` into tokens. For each token:

       * `ALL` -> increment every `mentions[i]`.
       * `HERE` -> increment `mentions[i]` only if `is_online[i]` is `True`.
       * `idX` -> parse `X` and increment `mentions[X]` (no check on online status; offline users can be mentioned by `idX` and `ALL`).
   * Repeat until all timestamps are processed.

4. **Return mentions array**

Important small details:

* If a user goes offline at timestamp `t`, and a `MESSAGE` also occurs at `t`, the `OFFLINE` must be applied before the `MESSAGE` (so that the message at `t` does not count that user as online).
* If `offline_until[i] == t` (meaning the user's offline window ended at `t`), they are considered online at `t` before other actions.
* Duplicate mention tokens count separately (e.g., `id0 id0` increments `id0` twice).

---

## Examples

Example 1:

```
Input:
numberOfUsers = 2,
events = [
  ["MESSAGE","10","id1 id0"],
  ["OFFLINE","11","id0"],
  ["MESSAGE","71","HERE"]
]

Output: [2,2]
Explanation:
- Initially both online.
- t=10: id1 and id0 are mentioned -> mentions = [1,1]
- t=11: id0 goes offline -> id0 offline until 71
- t=71: offline period for id0 expired at t=71, so id0 is online before processing events at t=71.
        MESSAGE "HERE" counts online users -> both online -> mentions = [2,2]
```

Example 2:

```
Input:
numberOfUsers = 2,
events = [
  ["MESSAGE","10","id1 id0"],
  ["OFFLINE","11","id0"],
  ["MESSAGE","12","ALL"]
]

Output: [2,2]
Explanation:
- t=10: id1,id0 mentioned -> [1,1]
- t=11: id0 offline until 71
- t=12: ALL mentions everyone, including offline users -> both increment -> [2,2]
```

Example 3:

```
Input:
numberOfUsers = 2,
events = [
  ["OFFLINE","10","0"],
  ["MESSAGE","12","HERE"]
]

Output: [0,1]
Explanation:
- Initially both online.
- t=10: user 0 goes offline until 70
- t=12: "HERE" only counts online users. user 0 still offline, user 1 is online -> mentions = [0,1]
```

---

## How to use / Run locally

### General

* Use any of the provided language implementations.
* Replace the `events` input with the test events (matching the structure used in that language).
* Compile/run according to each language instructions below.

### C++

* Put the C++ class into a file (e.g., `solution.cpp`) and add a `main()` to construct `events` and call `countMentions`.
* Compile: `g++ -std=c++17 solution.cpp -O2 -o solution`
* Run: `./solution`

### Java

* Save class `Solution` in `Solution.java`.
* Add a `public static void main(String[] args)` method that builds sample events and calls `countMentions`.
* Compile: `javac Solution.java`
* Run: `java Solution`

### JavaScript

* Save as `solution.js` and create a wrapper to call `countMentions(numberOfUsers, events)` and `console.log` the result.
* Run: `node solution.js`

### Python3

* Save as `solution.py`.
* Instantiate `Solution()` and call `countMentions(numberOfUsers, events)` then `print()` the result.
* Run: `python3 solution.py`

### Go

* Put the `countMentions` function in `main.go` and add a `main()` to call it with sample data and print results.
* Build/run: `go run main.go`

---

## Notes & Optimizations

* With constraints (≤ 100 events and ≤ 100 users), the simple approach that loops users during `ALL`/`HERE` tokens is fine and very clear.
* If numberOfUsers and events were large (e.g., 1e5), optimize:

  * Maintain a global `onlineCount` for quick handling of `HERE` token; keep an array `personalCounts` to handle `idX`. For `ALL`, you could add a global accumulator and finalize per-user counts later (lazy addition).
  * Use difference arrays or segment trees if patterns require range updates.
* Always respect the order-of-operations at the same timestamp:

  1. automatic online expirations (offline windows ending)
  2. OFFLINE events
  3. MESSAGE events

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
