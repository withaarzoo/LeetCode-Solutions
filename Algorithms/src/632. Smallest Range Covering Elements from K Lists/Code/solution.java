import java.util.PriorityQueue;
import java.util.List;

class Solution {
    public int[] smallestRange(List<List<Integer>> nums) {
        int k = nums.size();
        PriorityQueue<int[]> minHeap = new PriorityQueue<>((a, b) -> a[0] - b[0]);
        int maxValue = Integer.MIN_VALUE;

        // Initialize heap with the first element of each list
        for (int i = 0; i < k; ++i) {
            minHeap.offer(new int[] { nums.get(i).get(0), i, 0 });
            maxValue = Math.max(maxValue, nums.get(i).get(0));
        }

        int rangeStart = 0, rangeEnd = Integer.MAX_VALUE;

        while (!minHeap.isEmpty()) {
            int[] minElement = minHeap.poll();
            int minValue = minElement[0], row = minElement[1], col = minElement[2];

            // Update the smallest range
            if (maxValue - minValue < rangeEnd - rangeStart) {
                rangeStart = minValue;
                rangeEnd = maxValue;
            }

            // Move to the next element in the current list
            if (col + 1 < nums.get(row).size()) {
                minHeap.offer(new int[] { nums.get(row).get(col + 1), row, col + 1 });
                maxValue = Math.max(maxValue, nums.get(row).get(col + 1));
            } else {
                break; // One list is exhausted
            }
        }

        return new int[] { rangeStart, rangeEnd };
    }
}