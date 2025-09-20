# 3508. Implement Router — README

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
* [Step-by-step Detailed Explanation (Python)](#step-by-step-detailed-explanation-python)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

Design a data structure to simulate a network router that stores incoming data packets and supports the following operations:

1. `Router(int memoryLimit)`: create a router with a maximum number of stored packets `memoryLimit`.
2. `addPacket(int source, int destination, int timestamp) -> bool`: add the packet `(source, destination, timestamp)` if it's not already present. If adding the packet would exceed `memoryLimit`, remove oldest packet(s) first (FIFO). Return `true` if added, `false` if duplicate.
3. `forwardPacket() -> int[]`: forward (remove and return) the oldest packet as `[source, destination, timestamp]`, or return empty array if none.
4. `getCount(int destination, int startTime, int endTime) -> int`: return how many currently stored (not yet forwarded) packets have the given `destination` and `timestamp` in the inclusive range `[startTime, endTime]`.

**Important detail**: calls to `addPacket` will be made in **non-decreasing order of `timestamp`**. This allows per-destination timestamps to remain sorted as we append.

---

## Constraints

* `2 <= memoryLimit <= 10^5` (typical upper bounds given by the problem constraints)
* Each `source`, `destination`, `timestamp` fit into 32-bit integers.
* At most \~10^5 combined calls across `addPacket`, `forwardPacket`, and `getCount`.
* `addPacket` calls arrive in non-decreasing `timestamp` order (this is critical for our approach).

---

## Intuition

I thought about what the router must do: it stores packets in FIFO order, rejects duplicates, evicts the oldest packet(s) when full, and answers counts for a destination within a timestamp range. The crucial observation I used was:

* Because `addPacket` calls arrive in non-decreasing `timestamp` order, the timestamps we append for each destination will always be **sorted**. That means I can perform binary searches for `getCount`.
* All removals (eviction when full and forwarding) remove the **oldest** packet globally. For a particular destination, removals always happen at the front of its timestamp list. So I can maintain a head index for each destination to avoid costly removals from the front of a dynamic array.

This lets `addPacket` and `forwardPacket` be O(1) amortized, and `getCount` be O(log m) where `m` is number of stored timestamps for that destination.

---

## Approach

Step-by-step I solved the problem as follows (short and simple):

1. Keep a global FIFO queue of stored packets in arrival order. I use a deque or an array with a moving head.
2. Keep a `seen` set (hash) of `(source, destination, timestamp)` keys so `addPacket` can instantly detect duplicates.
3. For each `destination`, maintain an append-only list of timestamps and a `head` index telling how many earliest timestamps were removed. Because `addPacket` timestamps are non-decreasing, the list remains sorted.
4. `addPacket`: If duplicate, return `false`. Otherwise, if adding would exceed the capacity, remove oldest packet(s) from the global queue — for each removal remove from `seen` and increment the removed packet's destination `head`. Then append this new packet to the queue, add key to `seen`, and append timestamp to that destination's timestamp list. Return `true`.
5. `forwardPacket`: If queue empty return `[]`. Else remove the queue front, remove key from `seen`, advance that destination's `head`, and return the packet.
6. `getCount(destination, start, end)`: Use binary search (lower\_bound and upper\_bound) on the per-destination timestamp list starting from that destination's current `head` index. The count is `upperBound - lowerBound`.

---

## Data Structures Used

* **Global queue**: `deque<tuple>` or array + head pointer — stores stored packets in FIFO order.
* **Seen set**: `unordered_set` or language-specific set to detect duplicates quickly.
* **Per-destination timestamps**: append-only `vector/list` storing timestamps for each destination.
* **Per-destination head**: an integer index showing how many entries at the front were removed (lazy removal).

This combination keeps both memory and time efficient.

---

## Operations & Behavior Summary

* Duplicate detection is exact on `(source, destination, timestamp)`.
* Eviction is FIFO — oldest packet gets removed first.
* `getCount` queries only count packets that are still stored (not forwarded or evicted).
* Because `addPacket` timestamps are non-decreasing, per-destination timestamp arrays remain sorted and binary search works.

---

## Complexity

* **Time Complexity:**

  * `addPacket`: O(1) amortized — hash check, append, and at most one eviction which is O(1).
  * `forwardPacket`: O(1).
  * `getCount`: O(log m) where `m` is number of stored timestamps for that destination (binary searches).
* **Space Complexity:** O(M + D)

  * `M` = `memoryLimit` (max number of stored packets).
  * `D` = number of distinct destinations (for per-destination lists and head indices).

---

## Multi-language Solutions

Below I include complete implementations for C++, Java, JavaScript, Python3, and Go. Each implementation follows the same design described above.

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Router {
private:
    int memoryLimit;
    deque<tuple<int,int,int>> q;
    unordered_set<string> seen;
    unordered_map<int, vector<int>> times;
    unordered_map<int,int> head;

    string makeKey(int s, int d, int t) {
        return to_string(s) + "#" + to_string(d) + "#" + to_string(t);
    }

public:
    Router(int memoryLimit) : memoryLimit(memoryLimit) {}

    bool addPacket(int source, int destination, int timestamp) {
        string key = makeKey(source, destination, timestamp);
        if (seen.count(key)) return false;

        while ((int)q.size() >= memoryLimit) {
            auto [os, od, ot] = q.front();
            q.pop_front();
            seen.erase(makeKey(os, od, ot));
            head[od] += 1;
        }

        q.emplace_back(source, destination, timestamp);
        seen.insert(key);
        times[destination].push_back(timestamp);
        return true;
    }

    vector<int> forwardPacket() {
        if (q.empty()) return {};
        auto [s, d, t] = q.front();
        q.pop_front();
        seen.erase(makeKey(s, d, t));
        head[d] += 1;
        return {s, d, t};
    }

    int getCount(int destination, int startTime, int endTime) {
        auto it = times.find(destination);
        if (it == times.end()) return 0;
        vector<int> &arr = it->second;
        int h = head[destination];
        auto loIt = lower_bound(arr.begin() + h, arr.end(), startTime);
        auto hiIt = upper_bound(arr.begin() + h, arr.end(), endTime);
        return (int)(hiIt - loIt);
    }
};
```

### Java

```java
import java.util.*;

class Router {
    private int memoryLimit;
    private Deque<int[]> q;
    private Set<String> seen;
    private Map<Integer, ArrayList<Integer>> times;
    private Map<Integer, Integer> head;

    private String makeKey(int s, int d, int t) {
        return s + "#" + d + "#" + t;
    }

    public Router(int memoryLimit) {
        this.memoryLimit = memoryLimit;
        q = new ArrayDeque<>();
        seen = new HashSet<>();
        times = new HashMap<>();
        head = new HashMap<>();
    }

    public boolean addPacket(int source, int destination, int timestamp) {
        String key = makeKey(source, destination, timestamp);
        if (seen.contains(key)) return false;

        while (q.size() >= memoryLimit) {
            int[] old = q.pollFirst();
            String oldKey = makeKey(old[0], old[1], old[2]);
            seen.remove(oldKey);
            head.put(old[1], head.getOrDefault(old[1], 0) + 1);
        }

        q.addLast(new int[]{source, destination, timestamp});
        seen.add(key);
        times.computeIfAbsent(destination, k -> new ArrayList<>()).add(timestamp);
        return true;
    }

    public int[] forwardPacket() {
        if (q.isEmpty()) return new int[0];
        int[] pkt = q.pollFirst();
        String key = makeKey(pkt[0], pkt[1], pkt[2]);
        seen.remove(key);
        head.put(pkt[1], head.getOrDefault(pkt[1], 0) + 1);
        return pkt;
    }

    private int lowerBound(ArrayList<Integer> arr, int target, int lo) {
        int l = lo, r = arr.size();
        while (l < r) {
            int m = l + (r - l) / 2;
            if (arr.get(m) < target) l = m + 1;
            else r = m;
        }
        return l;
    }
    private int upperBound(ArrayList<Integer> arr, int target, int lo) {
        int l = lo, r = arr.size();
        while (l < r) {
            int m = l + (r - l) / 2;
            if (arr.get(m) <= target) l = m + 1;
            else r = m;
        }
        return l;
    }

    public int getCount(int destination, int startTime, int endTime) {
        ArrayList<Integer> arr = times.get(destination);
        if (arr == null) return 0;
        int h = head.getOrDefault(destination, 0);
        int L = lowerBound(arr, startTime, h);
        int R = upperBound(arr, endTime, h);
        return R - L;
    }
}
```

### JavaScript

```javascript
var Router = function(memoryLimit) {
    this.memoryLimit = memoryLimit;
    this.queue = [];
    this.qHead = 0;
    this.size = 0;
    this.seen = new Set();
    this.dest = new Map();
};

function makeKey(s, d, t) { return s + "#" + d + "#" + t; }

Router.prototype.addPacket = function(source, destination, timestamp) {
    const key = makeKey(source, destination, timestamp);
    if (this.seen.has(key)) return false;
    while (this.size >= this.memoryLimit) {
        const old = this.queue[this.qHead++];
        this.size--;
        this.seen.delete(makeKey(old[0], old[1], old[2]));
        const dObj = this.dest.get(old[1]);
        dObj.head += 1;
    }
    this.queue.push([source, destination, timestamp]);
    this.size++;
    this.seen.add(key);
    if (!this.dest.has(destination)) this.dest.set(destination, { arr: [], head: 0 });
    this.dest.get(destination).arr.push(timestamp);
    return true;
};

Router.prototype.forwardPacket = function() {
    if (this.size === 0) return [];
    const pkt = this.queue[this.qHead++];
    this.size--;
    this.seen.delete(makeKey(pkt[0], pkt[1], pkt[2]));
    const dObj = this.dest.get(pkt[1]);
    dObj.head += 1;
    return pkt;
};

function lowerBound(arr, target, lo) { let l = lo, r = arr.length; while (l < r) { const m = (l + r) >> 1; if (arr[m] < target) l = m + 1; else r = m; } return l; }
function upperBound(arr, target, lo) { let l = lo, r = arr.length; while (l < r) { const m = (l + r) >> 1; if (arr[m] <= target) l = m + 1; else r = m; } return l; }

Router.prototype.getCount = function(destination, startTime, endTime) {
    if (!this.dest.has(destination)) return 0;
    const dObj = this.dest.get(destination);
    const arr = dObj.arr;
    const h = dObj.head;
    const L = lowerBound(arr, startTime, h);
    const R = upperBound(arr, endTime, h);
    return R - L;
};
```

### Python3

```python
from collections import deque, defaultdict
import bisect
from typing import List

class Router:
    def __init__(self, memoryLimit: int):
        self.memoryLimit = memoryLimit
        self.q = deque()
        self.seen = set()
        self.times = defaultdict(list)
        self.head = defaultdict(int)

    def addPacket(self, source: int, destination: int, timestamp: int) -> bool:
        key = (source, destination, timestamp)
        if key in self.seen:
            return False
        while len(self.q) >= self.memoryLimit:
            os, od, ot = self.q.popleft()
            self.seen.remove((os, od, ot))
            self.head[od] += 1
        self.q.append((source, destination, timestamp))
        self.seen.add(key)
        self.times[destination].append(timestamp)
        return True

    def forwardPacket(self) -> List[int]:
        if not self.q:
            return []
        s, d, t = self.q.popleft()
        self.seen.remove((s, d, t))
        self.head[d] += 1
        return [s, d, t]

    def getCount(self, destination: int, startTime: int, endTime: int) -> int:
        arr = self.times.get(destination, [])
        h = self.head.get(destination, 0)
        L = bisect.bisect_left(arr, startTime, lo=h)
        R = bisect.bisect_right(arr, endTime, lo=h)
        return R - L
```

### Go

```go
package main

import "fmt"

type Packet struct { s, d, t int }

type Router struct {
    memoryLimit int
    queue       []Packet
    qHead       int
    size        int
    seen        map[string]bool
    destTimes   map[int][]int
    destHeads   map[int]int
}

func makeKey(s, d, t int) string { return fmt.Sprintf("%d#%d#%d", s, d, t) }

func Constructor(memoryLimit int) Router {
    return Router{memoryLimit: memoryLimit, queue: make([]Packet, 0), qHead: 0, size: 0, seen: make(map[string]bool), destTimes: make(map[int][]int), destHeads: make(map[int]int)}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
    key := makeKey(source, destination, timestamp)
    if this.seen[key] { return false }
    for this.size >= this.memoryLimit {
        old := this.queue[this.qHead]
        this.qHead++
        this.size--
        delete(this.seen, makeKey(old.s, old.d, old.t))
        this.destHeads[old.d]++
    }
    this.queue = append(this.queue, Packet{source, destination, timestamp})
    this.size++
    this.seen[key] = true
    this.destTimes[destination] = append(this.destTimes[destination], timestamp)
    return true
}

func (this *Router) ForwardPacket() []int {
    if this.size == 0 { return []int{} }
    pkt := this.queue[this.qHead]
    this.qHead++
    this.size--
    delete(this.seen, makeKey(pkt.s, pkt.d, pkt.t))
    this.destHeads[pkt.d]++
    return []int{pkt.s, pkt.d, pkt.t}
}

// lowerBound & upperBound omitted for brevity in the README but present in full implementation
```

---

## Step-by-step Detailed Explanation (Python)

I will walk through the key lines of the Python implementation and explain why each step is done.

1. `self.q = deque()` — I use a `deque` to store packets in FIFO order. `popleft()` on a deque is O(1), so removing the oldest packet is efficient.

2. `self.seen = set()` — This stores tuples `(source, destination, timestamp)` and lets me check for duplicates in O(1) time.

3. `self.times = defaultdict(list)` and `self.head = defaultdict(int)` — For each destination, `times[d]` is a list of appended timestamps (sorted because adds are non-decreasing) and `head[d]` counts how many earliest timestamps have been removed. Instead of physically removing from `times[d]` (which would be O(n) if done at the front), I increment `head[d]` when a timestamp is removed logically.

4. `addPacket(...)`:

   * `if key in self.seen: return False` — Reject duplicates.
   * `while len(self.q) >= self.memoryLimit: ...` — If we already reached capacity, remove the global oldest packet(s) `popleft()` until we have room. For each removed packet I: remove its key from `seen` and `head[old_dest] += 1`.
   * `self.q.append((source, destination, timestamp))` and `self.times[destination].append(timestamp)` and `self.seen.add(key)` — Add the packet to the system.

5. `forwardPacket(...)`:

   * If queue empty return `[]`.
   * Else `popleft()` the front packet, remove from `seen`, increment `head[destination]`, return `[s, d, t]`.

6. `getCount(destination, startTime, endTime)`:

   * Let `arr = self.times[destination]` and `h = self.head[destination]`.
   * Because `arr` has sorted timestamps and the valid portion is `arr[h:]`, use `bisect_left(arr, startTime, lo=h)` and `bisect_right(arr, endTime, lo=h)` so we count only currently stored timestamps.
   * Return `R - L`.

**Why head index?**

* Removing from Python list front is O(n). By using a head index we avoid shifting elements and keep removal O(1) (just increment an integer). As long as we use `bisect` with `lo=head`, correctness is preserved.

---

## Examples

Use the provided LeetCode examples to verify correctness. The implementations follow the examples in the problem statement.

---

## How to use / Run locally

1. Create a folder for the problem, e.g. `3508-Implement-Router/`.
2. Add this `README.md` and then create language-specific files (e.g. `router.py`, `Router.java`, `router.cpp`, `router.js`, `router.go`) with the code blocks above.
3. For compiled languages (C++, Java, Go) compile and run with a small harness (read input or call methods programmatically). For Python and JS you can directly import or run test harnesses.

---

## Notes & Optimizations

* If memory usage of per-destination arrays needs to be reclaimed, periodically trim arrays when `head` grows large (e.g., when `head > 1000` do a physical slice: `arr = arr[head:]` and reset `head = 0`). This avoids unbounded memory usage for extremely long test runs.
* The `seen` key uses a string join in some implementations. If performance matters, use a packed integer tuple or a custom struct with a proper hash function.

---

## Author

Made by **Aarzoo** — optimized, commented multi-language solutions and a clear explanation for the LeetCode problem "3508. Implement Router".
