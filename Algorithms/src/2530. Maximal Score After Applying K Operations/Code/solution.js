/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxKelements = function (nums, k) {
  // Max heap using a priority queue simulation
  const maxHeap = new MaxPriorityQueue();

  // Insert all elements into the max-heap
  for (let num of nums) {
    maxHeap.enqueue(num);
  }

  let score = 0;

  // Perform k operations
  for (let i = 0; i < k; i++) {
    let maxVal = maxHeap.dequeue().element;

    // Add the largest value to the score
    score += maxVal;

    // Replace the number with ceil(maxVal / 3)
    maxHeap.enqueue(Math.ceil(maxVal / 3));
  }

  return score;
};
