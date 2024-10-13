/**
 * @param {number[][]} nums
 * @return {number[]}
 */
var smallestRange = function (nums) {
  const k = nums.length;
  const minHeap = new MinPriorityQueue({ priority: (x) => x[0] });
  let maxValue = -Infinity;

  // Initialize heap with first element from each list
  for (let i = 0; i < k; i++) {
    minHeap.enqueue([nums[i][0], i, 0]);
    maxValue = Math.max(maxValue, nums[i][0]);
  }

  let rangeStart = 0,
    rangeEnd = Infinity;

  while (!minHeap.isEmpty()) {
    const [minValue, row, col] = minHeap.dequeue().element;

    // Update the smallest range
    if (maxValue - minValue < rangeEnd - rangeStart) {
      rangeStart = minValue;
      rangeEnd = maxValue;
    }

    // Move to the next element in the current list
    if (col + 1 < nums[row].length) {
      minHeap.enqueue([nums[row][col + 1], row, col + 1]);
      maxValue = Math.max(maxValue, nums[row][col + 1]);
    } else {
      break; // One list is exhausted
    }
  }

  return [rangeStart, rangeEnd];
};
