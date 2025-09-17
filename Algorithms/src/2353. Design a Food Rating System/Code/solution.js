/**
 * @param {string[]} foods
 * @param {string[]} cuisines
 * @param {number[]} ratings
 */
var FoodRatings = function (foods, cuisines, ratings) {
  this.foodInfo = new Map();

  this.cuisineHeaps = new Map();

  for (let i = 0; i < foods.length; i++) {
    const food = foods[i];
    const cuisine = cuisines[i];
    const rating = ratings[i];

    this.foodInfo.set(food, { cuisine, rating });

    if (!this.cuisineHeaps.has(cuisine)) {
      this.cuisineHeaps.set(cuisine, []);
    }

    this.cuisineHeaps.get(cuisine).push({ food, rating });

    this._heapifyUp(
      this.cuisineHeaps.get(cuisine),
      this.cuisineHeaps.get(cuisine).length - 1
    );
  }
};

/**
 * @param {string} food
 * @param {number} newRating
 * @return {void}
 */
FoodRatings.prototype.changeRating = function (food, newRating) {
  const foodData = this.foodInfo.get(food);
  foodData.rating = newRating;

  const cuisine = foodData.cuisine;
  this.cuisineHeaps.get(cuisine).push({ food, rating: newRating });

  const heap = this.cuisineHeaps.get(cuisine);
  this._heapifyUp(heap, heap.length - 1);
};

/**
 * @param {string} cuisine
 * @return {string}
 */
FoodRatings.prototype.highestRated = function (cuisine) {
  const heap = this.cuisineHeaps.get(cuisine);

  while (heap.length > 0) {
    const topFood = heap[0];
    const currentRating = this.foodInfo.get(topFood.food).rating;

    if (topFood.rating === currentRating) {
      return topFood.food;
    }

    this._extractMax(heap);
  }

  return "";
};

FoodRatings.prototype._heapifyUp = function (heap, index) {
  while (index > 0) {
    const parentIndex = Math.floor((index - 1) / 2);

    if (this._shouldSwap(heap[index], heap[parentIndex])) {
      [heap[index], heap[parentIndex]] = [heap[parentIndex], heap[index]];
      index = parentIndex;
    } else {
      break;
    }
  }
};

FoodRatings.prototype._extractMax = function (heap) {
  if (heap.length === 0) return null;
  if (heap.length === 1) return heap.pop();

  const max = heap[0];
  heap[0] = heap.pop();
  this._heapifyDown(heap, 0);
  return max;
};

FoodRatings.prototype._heapifyDown = function (heap, index) {
  while (true) {
    let largest = index;
    const leftChild = 2 * index + 1;
    const rightChild = 2 * index + 2;

    if (
      leftChild < heap.length &&
      this._shouldSwap(heap[leftChild], heap[largest])
    ) {
      largest = leftChild;
    }

    if (
      rightChild < heap.length &&
      this._shouldSwap(heap[rightChild], heap[largest])
    ) {
      largest = rightChild;
    }

    if (largest !== index) {
      [heap[index], heap[largest]] = [heap[largest], heap[index]];
      index = largest;
    } else {
      break;
    }
  }
};

FoodRatings.prototype._shouldSwap = function (a, b) {
  if (a.rating !== b.rating) {
    return a.rating > b.rating;
  }

  return a.food < b.food;
};

/**
 * Your FoodRatings object will be instantiated and called as such:
 * var obj = new FoodRatings(foods, cuisines, ratings)
 * obj.changeRating(food,newRating)
 * var param_2 = obj.highestRated(cuisine)
 */
