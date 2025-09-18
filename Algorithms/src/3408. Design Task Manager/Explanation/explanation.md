# Design Task Manager

This repository contains my solution to **LeetCode 3408 — Design Task Manager**.
I include a clear explanation of my approach, complexity analysis, a step-by-step walkthrough, and working implementations in **C++**, **Java**, **JavaScript**, **Python3**, and **Go**.

---

## Problem Summary (short)

We need to design a task manager that supports:

* initialization with a list of tasks `[userId, taskId, priority]`,
* `add(userId, taskId, priority)` — add a new task,
* `edit(taskId, newPriority)` — update an existing task's priority,
* `rmv(taskId)` — remove a task,
* `execTop()` — execute and remove the task with the highest priority across all users (if tie, choose the largest `taskId`), and return the corresponding `userId`. If no tasks exist, return `-1`.

Important: operations `edit` and `rmv` are given valid `taskId`s (task exists when called).

---

## Intuition

I thought about using a single priority queue so I can always get the task with the highest priority fast. But a heap alone cannot quickly update or delete an arbitrary `taskId`. So I kept a **hash map** as the authoritative record `taskId -> (userId, priority)` and continued to push candidates into the heap on every add/edit. When popping from the heap I verify with the map whether the popped entry is still valid — if not, I skip it. This is **lazy deletion** and keeps all operations fast.

---

## Approach

1. Maintain a map: `taskId -> (userId, priority)` that stores the current valid state of each task.
2. Maintain a max-heap (priority queue) of candidates with key `(priority, taskId)` so the heap top is highest `priority`, and if tie highest `taskId`.
3. On `add` or `edit`, update the map and push a new candidate into the heap. Do not try to remove the old candidate from heap immediately (because removing arbitrary elements from a binary heap is expensive).
4. On `rmv`, remove the `taskId` from the map — the heap entry becomes stale and will be ignored later.
5. On `execTop`, pop candidates from heap until we find one that matches the map entry exactly (same `taskId` and same `priority`). That is the valid top task; remove it from the map and return the `userId`. If none exists, return `-1`.

This lazy-checking approach ensures we never need to remove or update arbitrary heap nodes directly.

---

## Complexity

* **Time Complexity**:

  * `add` / `edit` — O(log M) for heap push (M = current number of pushed entries).
  * `rmv` — O(1) removing from map.
  * `execTop` — Amortized O(log M): each heap entry is pushed once and popped at most once; skipping stale entries costs only once per stale push.
* **Space Complexity**: O(M) for heap + O(N) for map, where N is the number of current tasks (M ≥ N because of stale entries).

---

## Example

Input operations:

```
["TaskManager","add","edit","execTop","rmv","add","execTop"]
[[[[1,101,10],[2,102,20],[3,103,15]]],[4,104,5],[102,8],[],[101],[5,105,15],[]]
```

Output:

```
[null, null, null, 3, null, null, 5]
```

Explanation overview:

* Initialized with tasks for users 1,2,3.
* Added task 104 for user 4.
* Edited 102's priority to 8 (was 20 originally in this example — note: example changed).
* `execTop()` returns the `userId` of the highest priority valid task.
* etc.

---

## Implementations

> Each implementation uses the same lazy-deletion idea:
>
> * authoritative `map` from `taskId` to `(userId, priority)`.
> * `heap` stores candidate entries `(priority, taskId)` ordered by highest priority and highest taskId on ties.
> * On `execTop` we pop until we find a valid (non-stale) entry.

---

### C++ (clean and efficient)

```c++
#include <bits/stdc++.h>
using namespace std;

/*
C++ TaskManager:
- unordered_map<int, pair<int,int>> mp; // taskId -> {userId, priority}
- priority_queue<pair<int,int>> pq;     // {priority, taskId}
Lazy deletion approach: verify top with mp.
*/
class TaskManager {
private:
    unordered_map<int, pair<int,int>> mp;
    priority_queue<pair<int,int>> pq; // top = highest priority, if tie highest taskId

public:
    TaskManager(vector<vector<int>>& tasks) {
        for (auto &t : tasks) {
            if (t.size() < 3) continue;
            int user = t[0], task = t[1], pr = t[2];
            mp[task] = {user, pr};
            pq.push({pr, task});
        }
    }
    
    void add(int userId, int taskId, int priority) {
        mp[taskId] = {userId, priority};
        pq.push({priority, taskId});
    }
    
    void edit(int taskId, int newPriority) {
        // guaranteed to exist
        mp[taskId].second = newPriority;
        pq.push({newPriority, taskId});
    }
    
    void rmv(int taskId) {
        // guaranteed to exist
        mp.erase(taskId);
    }
    
    int execTop() {
        while (!pq.empty()) {
            auto top = pq.top(); pq.pop();
            int pr = top.first, id = top.second;
            auto it = mp.find(id);
            if (it == mp.end()) continue;           // removed
            if (it->second.second != pr) continue;  // stale
            int userId = it->second.first;
            mp.erase(it);
            return userId;
        }
        return -1;
    }
};
```

---

### Java (clear, robust)

```java
import java.util.*;

class TaskManager {
    private static class Item {
        int priority, taskId;
        Item(int p, int t) { priority = p; taskId = t; }
    }
    
    private final Map<Integer, int[]> map; // taskId -> [userId, priority]
    private final PriorityQueue<Item> pq;
    
    public TaskManager(List<List<Integer>> tasks) {
        map = new HashMap<>();
        pq = new PriorityQueue<>((a, b) -> {
            if (a.priority != b.priority) return Integer.compare(b.priority, a.priority);
            return Integer.compare(b.taskId, a.taskId);
        });
        for (List<Integer> t : tasks) {
            if (t.size() < 3) continue;
            int user = t.get(0), task = t.get(1), pr = t.get(2);
            map.put(task, new int[]{user, pr});
            pq.offer(new Item(pr, task));
        }
    }
    
    public void add(int userId, int taskId, int priority) {
        map.put(taskId, new int[]{userId, priority});
        pq.offer(new Item(priority, taskId));
    }
    
    public void edit(int taskId, int newPriority) {
        int[] arr = map.get(taskId); // guaranteed to exist
        arr[1] = newPriority;
        pq.offer(new Item(newPriority, taskId));
    }
    
    public void rmv(int taskId) {
        map.remove(taskId);
    }
    
    public int execTop() {
        while (!pq.isEmpty()) {
            Item it = pq.poll();
            int id = it.taskId, pr = it.priority;
            int[] arr = map.get(id);
            if (arr == null) continue;
            if (arr[1] != pr) continue;
            int userId = arr[0];
            map.remove(id);
            return userId;
        }
        return -1;
    }
}
```

---

### JavaScript (safe for repeated runs on judge platforms)

```javascript
var TaskManager = function(tasks) {
    // Define MaxHeap inside so it doesn't clash between runs
    class MaxHeap {
        constructor() {
            this.heap = [];
        }
        size() { return this.heap.length; }
        _cmp(a, b) {
            if (a.priority !== b.priority) return a.priority > b.priority;
            return a.taskId > b.taskId;
        }
        _swap(i, j) {
            [this.heap[i], this.heap[j]] = [this.heap[j], this.heap[i]];
        }
        push(item) {
            this.heap.push(item);
            this._siftUp(this.heap.length - 1);
        }
        pop() {
            if (this.heap.length === 0) return null;
            const top = this.heap[0];
            const last = this.heap.pop();
            if (this.heap.length > 0) {
                this.heap[0] = last;
                this._siftDown(0);
            }
            return top;
        }
        _siftUp(idx) {
            while (idx > 0) {
                const parent = Math.floor((idx - 1) / 2);
                if (this._cmp(this.heap[idx], this.heap[parent])) {
                    this._swap(idx, parent);
                    idx = parent;
                } else break;
            }
        }
        _siftDown(idx) {
            const n = this.heap.length;
            while (true) {
                let largest = idx;
                const l = 2 * idx + 1, r = 2 * idx + 2;
                if (l < n && this._cmp(this.heap[l], this.heap[largest])) largest = l;
                if (r < n && this._cmp(this.heap[r], this.heap[largest])) largest = r;
                if (largest === idx) break;
                this._swap(idx, largest);
                idx = largest;
            }
        }
    }

    this.map = new Map(); // taskId -> {userId, priority}
    this.heap = new MaxHeap();

    for (const t of tasks) {
        if (t.length < 3) continue;
        const [user, task, pr] = t;
        this.map.set(task, { userId: user, priority: pr });
        this.heap.push({ priority: pr, taskId: task });
    }
};

TaskManager.prototype.add = function(userId, taskId, priority) {
    this.map.set(taskId, { userId, priority });
    this.heap.push({ priority, taskId });
};

TaskManager.prototype.edit = function(taskId, newPriority) {
    const cur = this.map.get(taskId);
    cur.priority = newPriority;
    this.heap.push({ priority: newPriority, taskId });
};

TaskManager.prototype.rmv = function(taskId) {
    this.map.delete(taskId);
};

TaskManager.prototype.execTop = function() {
    while (this.heap.size() > 0) {
        const top = this.heap.pop();
        const rec = this.map.get(top.taskId);
        if (!rec) continue; // removed
        if (rec.priority !== top.priority) continue; // stale
        this.map.delete(top.taskId);
        return rec.userId;
    }
    return -1;
};

/** 
 * Your TaskManager object will be instantiated and called as such:
 * var obj = new TaskManager(tasks)
 * obj.add(userId,taskId,priority)
 * obj.edit(taskId,newPriority)
 * obj.rmv(taskId)
 * var param_4 = obj.execTop()
 */
```

---

### Python3 (concise & readable)

```python
import heapq
from typing import List

class TaskManager:

    def __init__(self, tasks: List[List[int]]):
        self.mp = {}           # taskId -> (userId, priority)
        self.heap = []         # (-priority, -taskId, taskId)
        for t in tasks:
            if len(t) < 3: continue
            user, task, pr = t[0], t[1], t[2]
            self.mp[task] = (user, pr)
            heapq.heappush(self.heap, (-pr, -task, task))

    def add(self, userId: int, taskId: int, priority: int) -> None:
        self.mp[taskId] = (userId, priority)
        heapq.heappush(self.heap, (-priority, -taskId, taskId))

    def edit(self, taskId: int, newPriority: int) -> None:
        user = self.mp[taskId][0]  # guaranteed to exist
        self.mp[taskId] = (user, newPriority)
        heapq.heappush(self.heap, (-newPriority, -taskId, taskId))

    def rmv(self, taskId: int) -> None:
        if taskId in self.mp:
            del self.mp[taskId]

    def execTop(self) -> int:
        while self.heap:
            neg_pr, neg_tid, tid = heapq.heappop(self.heap)
            cur = self.mp.get(tid)
            if cur is None:
                continue
            if cur[1] != -neg_pr:
                continue
            userId = cur[0]
            del self.mp[tid]
            return userId
        return -1
```

---

### Go (optional — idiomatic)

```go
package main

import (
 "container/heap"
)

// Entry represents a task in the priority queue.
type Entry struct {
 priority int
 taskId   int
}

// MaxHeap implements a max-heap for Entries.
// Higher priority first; if tie, larger taskId first.
type MaxHeap []Entry

func (h MaxHeap) Len() int      { return len(h) }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool {
 if h[i].priority != h[j].priority {
  return h[i].priority > h[j].priority
 }
 return h[i].taskId > h[j].taskId
}
func (h *MaxHeap) Push(x interface{}) {
 *h = append(*h, x.(Entry))
}
func (h *MaxHeap) Pop() interface{} {
 old := *h
 n := len(old)
 item := old[n-1]
 *h = old[:n-1]
 return item
}

// TaskManager structure
type TaskManager struct {
 mp   map[int][2]int // taskId -> [userId, priority]
 heap MaxHeap
}

// Constructor initializes TaskManager with given tasks
func Constructor(tasks [][]int) TaskManager {
 tm := TaskManager{
  mp:   make(map[int][2]int),
  heap: MaxHeap{},
 }
 for _, t := range tasks {
  if len(t) < 3 {
   continue
  }
  userId, taskId, pr := t[0], t[1], t[2]
  tm.mp[taskId] = [2]int{userId, pr}
  heap.Push(&tm.heap, Entry{priority: pr, taskId: taskId})
 }
 return tm
}

// Add a new task
func (this *TaskManager) Add(userId int, taskId int, priority int) {
 this.mp[taskId] = [2]int{userId, priority}
 heap.Push(&this.heap, Entry{priority: priority, taskId: taskId})
}

// Edit modifies the priority of an existing task
func (this *TaskManager) Edit(taskId int, newPriority int) {
 rec := this.mp[taskId]
 this.mp[taskId] = [2]int{rec[0], newPriority}
 heap.Push(&this.heap, Entry{priority: newPriority, taskId: taskId})
}

// Rmv removes a task
func (this *TaskManager) Rmv(taskId int) {
 delete(this.mp, taskId)
}

// ExecTop executes the highest priority task and returns userId
func (this *TaskManager) ExecTop() int {
 for this.heap.Len() > 0 {
  top := heap.Pop(&this.heap).(Entry)
  rec, ok := this.mp[top.taskId]
  if !ok {
   continue // task removed
  }
  if rec[1] != top.priority {
   continue // stale entry
  }
  delete(this.mp, top.taskId)
  return rec[0]
 }
 return -1
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */

```

---

## Step-by-Step Detailed Explanation (walkthrough)

1. **Why not update heap entries directly?**
   A binary heap does not support removing an arbitrary item or updating an element in O(log n) unless you maintain an external index-to-position map and update positions on swaps. That is more complex and error-prone. Instead I use lazy deletion.

2. **What is lazy deletion?**

   * When I change or remove a task, I update the authoritative `map`.
   * I also push a new candidate into the heap on `add`/`edit`.
   * Old candidates in the heap become "stale". They remain in the heap but will be ignored when popped because they don't match the `map` state.

3. **execTop logic (most important)**:

   * Pop the heap top candidate (the one with highest priority and highest taskId for tie).
   * Check `map` for that `taskId`.

     * If not present → it was removed earlier → THIS CANDIDATE IS STALE → skip.
     * If present but priority != popped priority → stale candidate (priority outdated) → skip.
     * Otherwise candidate is valid → remove from `map` and return its `userId`.
   * Repeat until either a valid candidate is found or heap empties. If empty → return `-1`.

4. **Correctness**:

   * The map always reflects the most recent state (add/edit/rmv).
   * The heap always provides the next candidate by priority/taskId.
   * Because each push is popped at most once, the total number of heap operations is linear in the number of pushes (amortized efficient).

5. **Edge cases**:

   * Multiple edits for same task produce multiple entries in heap — we always check map to validate.
   * Removing a task (`rmv`) deletes map entry; heap may still contain stale entries for that task but they will be ignored later.
   * If `execTop` finds nothing valid, return `-1`.

---

## How to test locally

* **Python**: create a small script that constructs `TaskManager` and calls methods; run with `python3`.
* **C++**: compile with `g++ -std=c++17` and run with a main() wrapper to test methods.
* **Java**: create a `main` that calls the class methods.
* **JavaScript**: run on Node with a small harness or paste into LeetCode editor. (Note: the provided JS implementation defines the helper heap inside the function to avoid "identifier already declared" runtime errors on some judges.)
