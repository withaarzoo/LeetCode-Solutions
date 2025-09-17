# Food Ratings System

A data structure problem where we manage foods, cuisines, and ratings efficiently.  
We need to:

- Track foods and their cuisines
- Update ratings dynamically
- Quickly fetch the highest-rated food for any cuisine  

This is solved using **hash maps + heaps (priority queues)** with **lazy deletion** for efficiency.

---

## Intuition

Think of it like managing a **restaurant database**:

- Each food belongs to a cuisine and has a rating.
- We must frequently update ratings.
- We must quickly query the "best food" in a cuisine.

Hash maps let us store and update food details in **O(1)**, while heaps let us retrieve the maximum efficiently in **O(log n)**.

---

## Approach

1. **Store food info**  
   Use a `Map` (`foodInfo`) to keep each food’s cuisine and current rating for fast lookups.

2. **Group foods by cuisine**  
   Use another `Map` (`cuisineHeaps`) where each cuisine has a **max-heap** of its foods, ordered by:
   - Higher rating first  
   - Lexicographically smaller name in case of tie  

3. **Handle rating updates**  
   When a food’s rating changes:
   - Update `foodInfo` directly  
   - Push a new entry into the heap (don’t remove the old one yet)  

4. **Get highest-rated food**  
   - Look at the heap’s top element  
   - If it matches the current rating in `foodInfo`, return it  
   - If not, discard it and continue until a valid one is found (lazy deletion)  

---

## Complexity

- **Time Complexity:**  
  - `changeRating`: **O(log n)** (heap insertion)  
  - `highestRated`: **O(log n)** (may pop stale entries)  

- **Space Complexity:** **O(n)** for food info + heaps.  
  Some duplicates exist in heaps due to updates, but still linear overall.

---

## Code (JavaScript)

```javascript
/**
 * @param {string[]} foods
 * @param {string[]} cuisines
 * @param {number[]} ratings
 */
var FoodRatings = function(foods, cuisines, ratings) {
    this.foodInfo = new Map();       // food -> { cuisine, rating }
    this.cuisineHeaps = new Map();   // cuisine -> max-heap of foods

    for (let i = 0; i < foods.length; i++) {
        const food = foods[i];
        const cuisine = cuisines[i];
        const rating = ratings[i];
        
        this.foodInfo.set(food, { cuisine, rating });
        
        if (!this.cuisineHeaps.has(cuisine)) {
            this.cuisineHeaps.set(cuisine, []);
        }
        this.cuisineHeaps.get(cuisine).push({ food, rating });
        this._heapifyUp(this.cuisineHeaps.get(cuisine), this.cuisineHeaps.get(cuisine).length - 1);
    }
};

/** Update a food’s rating */
FoodRatings.prototype.changeRating = function(food, newRating) {
    const foodData = this.foodInfo.get(food);
    foodData.rating = newRating;

    const cuisine = foodData.cuisine;
    this.cuisineHeaps.get(cuisine).push({ food, rating: newRating });

    const heap = this.cuisineHeaps.get(cuisine);
    this._heapifyUp(heap, heap.length - 1);
};

/** Get highest-rated food in a cuisine */
FoodRatings.prototype.highestRated = function(cuisine) {
    const heap = this.cuisineHeaps.get(cuisine);

    while (heap.length > 0) {
        const topFood = heap[0];
        const currentRating = this.foodInfo.get(topFood.food).rating;

        if (topFood.rating === currentRating) {
            return topFood.food;
        }
        this._extractMax(heap); // remove stale entry
    }
    return "";
};

/* -------- Heap Utilities -------- */
FoodRatings.prototype._heapifyUp = function(heap, index) {
    while (index > 0) {
        const parent = Math.floor((index - 1) / 2);
        if (this._shouldSwap(heap[index], heap[parent])) {
            [heap[index], heap[parent]] = [heap[parent], heap[index]];
            index = parent;
        } else break;
    }
};

FoodRatings.prototype._extractMax = function(heap) {
    if (heap.length <= 1) return heap.pop();
    const max = heap[0];
    heap[0] = heap.pop();
    this._heapifyDown(heap, 0);
    return max;
};

FoodRatings.prototype._heapifyDown = function(heap, index) {
    while (true) {
        let largest = index;
        const left = 2 * index + 1, right = 2 * index + 2;

        if (left < heap.length && this._shouldSwap(heap[left], heap[largest])) largest = left;
        if (right < heap.length && this._shouldSwap(heap[right], heap[largest])) largest = right;

        if (largest !== index) {
            [heap[index], heap[largest]] = [heap[largest], heap[index]];
            index = largest;
        } else break;
    }
};

FoodRatings.prototype._shouldSwap = function(a, b) {
    if (a.rating !== b.rating) return a.rating > b.rating;
    return a.food < b.food; // lexicographic tie-breaker
};
````

---

## Step-by-Step Explanation

### Constructor (`FoodRatings`)

- Builds two maps:

  - `foodInfo` for quick access to a food’s cuisine & rating
  - `cuisineHeaps` for cuisine-specific heaps
- Inserts foods into their cuisine heap while keeping max-heap order.

### `changeRating(food, newRating)`

- Updates `foodInfo` with the new rating.
- Pushes a fresh entry into the cuisine’s heap (lazy handling of stale entries).

### `highestRated(cuisine)`

- Checks the top of the cuisine’s heap.
- If the rating matches `foodInfo`, it’s valid → return it.
- If not, discard it and recheck until valid (lazy deletion).

### Heap Helpers

- `_heapifyUp` and `_heapifyDown` maintain max-heap order.
- `_shouldSwap` enforces ordering by rating, then lexicographically.
- `_extractMax` pops the maximum while restoring heap structure.

---

✅ **Key Insight:**
Instead of updating heap entries directly (expensive), we allow duplicates and clean up only when necessary. This *lazy deletion* keeps everything efficient.
