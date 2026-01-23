# Approach

I solve this problem by **simulating the process exactly as the problem forces me to do**, but in an optimized way.

First, I notice that:

* The array is **not sorted** only at places where
  `nums[i] > nums[i+1]`
* I call these places **bad positions**.
* If there are **no bad positions**, the array is already non-decreasing.

So instead of rebuilding the array every time, I just **track how many bad positions exist**.

---

### Step 1: Simulate a linked list

Since elements are removed during merging, I simulate a **doubly linked list** using arrays:

* `prev[i]` → index of left neighbor
* `next[i]` → index of right neighbor
* `removed[i]` → whether this index is already merged and removed

This lets me delete elements in **O(1)** time.

---

### Step 2: Always pick the minimum sum pair

The problem forces me to always merge the **adjacent pair with the minimum sum**.

To do this efficiently, I use a **min heap** that stores:

```
(sum of nums[i] + nums[i+1], index i)
```

If multiple pairs have the same sum, the heap automatically picks the **leftmost one** because the index is stored.

---

### Step 3: Count bad positions

I count how many times:

```
a[i] > a[i+1]
```

This number (`bad`) tells me whether the array is still unsorted.

---

### Step 4: Process merges until array becomes sorted

While `bad > 0`, I do the following:

1. Take the smallest pair from the heap
2. Skip it if:

   * the index is already removed
   * the right neighbor no longer exists
   * the sum is outdated (stale heap entry)
3. Identify affected neighbors:

   * left of `i`
   * pair `(i, j)`
   * right of `j`
4. **Remove old violations** caused by these pairs
5. Merge `j` into `i`
6. Update linked list pointers
7. **Add new violations** created after the merge
8. Push newly formed adjacent pairs back into the heap
9. Increase operation count

Each merge only affects **local neighbors**, so the updates are fast.

---

### Step 5: Stop when sorted

The moment `bad` becomes `0`, the array is non-decreasing, and I return the total number of operations.
