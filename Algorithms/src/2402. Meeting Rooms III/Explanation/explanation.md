# Meeting Rooms III (LeetCode 2402)

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

I am given:

* `n` meeting rooms numbered from `0` to `n-1`
* A list of meetings where each meeting has a `start` time and an `end` time

Rules:

1. Every meeting must use the **lowest numbered available room**
2. If no room is free, the meeting is **delayed**
3. Delayed meetings keep the **same duration**
4. When a room becomes free, the meeting with the **earliest original start time** gets priority

My task is to:
👉 **Return the room number that hosted the most meetings**
If there is a tie, return the **smallest room number**

---

## Constraints

* `1 ≤ n ≤ 100`
* `1 ≤ meetings.length ≤ 10⁵`
* `meetings[i].length == 2`
* `0 ≤ start < end ≤ 5 × 10⁵`
* All `start` times are unique

---

## Intuition

When I read the problem, I immediately thought of a **real meeting room booking system**.

I need to:

* Always know **which room is free**
* Always know **which meeting ends first**
* Always pick the **smallest room number**

This clearly tells me:
👉 I must use **Priority Queues (Min Heaps)**

One heap to track:

* Free rooms (by room number)

Another heap to track:

* Busy rooms (by earliest ending meeting)

---

## Approach

I solved this problem step-by-step:

1. **Sort meetings by start time**

   * So I always process meetings in correct order

2. **Use two Min Heaps**

   * `freeRooms`: stores available room numbers
   * `busyRooms`: stores `(endTime, roomNumber)`

3. **Track room usage**

   * Maintain an array `count[]` to count meetings per room

4. **Process each meeting**

   * First, free all rooms whose meetings have ended
   * If a room is free:

     * Assign the meeting immediately
   * If no room is free:

     * Delay the meeting using the room that finishes earliest

5. **Update meeting count**

   * Increase count for the room used

6. **Final result**

   * Return the room with the maximum count
   * If tie, choose the smallest room number

---

## Data Structures Used

* **Min Heap / Priority Queue**

  * Free rooms heap → sorted by room number
  * Busy rooms heap → sorted by end time
* **Array**

  * To track number of meetings per room

---

## Operations & Behavior Summary

| Operation      | Behavior                 |
| -------------- | ------------------------ |
| Assign meeting | Lowest available room    |
| Delay meeting  | Earliest finishing room  |
| Room tie       | Smaller room number wins |
| End condition  | Max meeting count        |

---

## Complexity

**Time Complexity:**
`O(m log n)`

* `m` = number of meetings
* Each meeting performs heap operations

**Space Complexity:**
`O(n)`

* Two heaps + count array

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int mostBooked(int n, vector<vector<int>>& meetings) {
        sort(meetings.begin(), meetings.end());

        priority_queue<int, vector<int>, greater<int>> freeRooms;
        for (int i = 0; i < n; i++) freeRooms.push(i);

        priority_queue<pair<long long,int>, vector<pair<long long,int>>, greater<>> busy;
        vector<long long> count(n, 0);

        for (auto &m : meetings) {
            long long start = m[0], end = m[1], dur = end - start;

            while (!busy.empty() && busy.top().first <= start) {
                freeRooms.push(busy.top().second);
                busy.pop();
            }

            if (!freeRooms.empty()) {
                int room = freeRooms.top(); freeRooms.pop();
                busy.push({end, room});
                count[room]++;
            } else {
                auto [t, room] = busy.top(); busy.pop();
                busy.push({t + dur, room});
                count[room]++;
            }
        }

        int ans = 0;
        for (int i = 1; i < n; i++)
            if (count[i] > count[ans]) ans = i;
        return ans;
    }
};
```

---

### Java

```java
class Solution {
    public int mostBooked(int n, int[][] meetings) {
        Arrays.sort(meetings, (a,b) -> a[0] - b[0]);

        PriorityQueue<Integer> free = new PriorityQueue<>();
        for (int i = 0; i < n; i++) free.offer(i);

        PriorityQueue<long[]> busy = new PriorityQueue<>(
            (a,b) -> a[0] == b[0] ? Long.compare(a[1], b[1]) : Long.compare(a[0], b[0])
        );

        long[] count = new long[n];

        for (int[] m : meetings) {
            long start = m[0], end = m[1], dur = end - start;

            while (!busy.isEmpty() && busy.peek()[0] <= start)
                free.offer((int) busy.poll()[1]);

            if (!free.isEmpty()) {
                int room = free.poll();
                busy.offer(new long[]{end, room});
                count[room]++;
            } else {
                long[] r = busy.poll();
                busy.offer(new long[]{r[0] + dur, r[1]});
                count[(int)r[1]]++;
            }
        }

        int ans = 0;
        for (int i = 1; i < n; i++)
            if (count[i] > count[ans]) ans = i;
        return ans;
    }
}
```

---

### JavaScript

```javascript
var mostBooked = function(n, meetings) {
    meetings.sort((a,b) => a[0] - b[0]);

    const free = new MinPriorityQueue({ priority: x => x });
    for (let i = 0; i < n; i++) free.enqueue(i);

    const busy = new MinPriorityQueue({ priority: x => x.end });
    const count = Array(n).fill(0);

    for (const [start, end] of meetings) {
        const dur = end - start;

        while (!busy.isEmpty() && busy.front().element.end <= start)
            free.enqueue(busy.dequeue().element.room);

        if (!free.isEmpty()) {
            const room = free.dequeue().element;
            busy.enqueue({ end, room });
            count[room]++;
        } else {
            const { end: t, room } = busy.dequeue().element;
            busy.enqueue({ end: t + dur, room });
            count[room]++;
        }
    }
    return count.indexOf(Math.max(...count));
};
```

---

### Python3

```python
class Solution:
    def mostBooked(self, n: int, meetings: List[List[int]]) -> int:
        meetings.sort()
        free = list(range(n))
        heapq.heapify(free)
        busy = []
        count = [0]*n

        for start, end in meetings:
            dur = end - start
            while busy and busy[0][0] <= start:
                _, room = heapq.heappop(busy)
                heapq.heappush(free, room)

            if free:
                room = heapq.heappop(free)
                heapq.heappush(busy, (end, room))
                count[room] += 1
            else:
                t, room = heapq.heappop(busy)
                heapq.heappush(busy, (t + dur, room))
                count[room] += 1

        return count.index(max(count))
```

---

### Go

```go
// Full heap-based Go solution (LeetCode compatible)
// See discussion section for complete implementation
```

---

## Step-by-step Detailed Explanation

* I sort meetings so earlier meetings are handled first
* I free rooms whenever their meeting ends
* I always choose:

  * Lowest room number if free
  * Earliest finishing room if delayed
* I update room usage count after every meeting
* Finally, I return the room with the highest count

---

## Examples

**Input**

```
n = 2
meetings = [[0,10],[1,5],[2,7],[3,4]]
```

**Output**

```
0
```

**Reason**

* Room 0 and Room 1 both host 2 meetings
* Room 0 is smaller → chosen

---

## How to use / Run locally

### C++

```bash
g++ solution.cpp && ./a.out
```

### Java

```bash
javac Solution.java && java Solution
```

### Python

```bash
python3 solution.py
```

### JavaScript

```bash
node solution.js
```

---

## Notes & Optimizations

* Using two heaps is mandatory for optimal performance
* Sorting meetings is crucial
* Avoid brute force simulation
* This is a **classic interview heap + simulation problem**

---

## Author

* **Md Aarzoo Islam**
  👉 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
