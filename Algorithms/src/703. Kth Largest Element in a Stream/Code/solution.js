// Define the KthLargest class to keep track of the k-th largest element in a stream of numbers
class KthLargest {
  // Constructor to initialize the class with 'k' and an array of numbers 'nums'
  constructor(k, nums) {
    // Store the value of 'k' (the position of the k-th largest element)
    this.k = k;
    // Initialize a minimum priority queue (min-heap) to efficiently keep track of the k largest elements
    this.minHeap = new MinPriorityQueue();

    // Iterate over each number in the 'nums' array and add it to the min-heap
    for (let num of nums) {
      this.add(num);
    }
  }

  // Method to add a new value to the data structure and return the k-th largest element
  add(val) {
    // Check if the size of the min-heap is less than 'k'
    // If true, directly add the value to the min-heap
    if (this.minHeap.size() < this.k) {
      this.minHeap.enqueue(val);
    }
    // If the size of the min-heap is equal to 'k', compare the new value with the smallest value in the heap
    // The smallest value in the min-heap is at the front (root of the heap)
    else if (val > this.minHeap.front().element) {
      // Remove the smallest value from the heap to maintain the size of 'k'
      this.minHeap.dequeue();
      // Add the new value to the heap as it is larger than the smallest value in the heap
      this.minHeap.enqueue(val);
    }

    // Return the smallest value in the heap, which is the k-th largest element in the stream
    return this.minHeap.front().element;
  }
}
