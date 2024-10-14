import java.util.PriorityQueue;

class Solution {
    public long maxKelements(int[] nums, int k) {
        // Priority queue as max-heap
        PriorityQueue<Long> maxHeap = new PriorityQueue<>((a, b) -> Long.compare(b, a));
        
        // Insert all elements into the max-heap
        for (int num : nums) {
            maxHeap.add((long) num);
        }
        
        long score = 0;
        
        // Perform k operations
        for (int i = 0; i < k; i++) {
            long maxVal = maxHeap.poll();
            
            // Add the largest value to the score
            score += maxVal;
            
            // Replace the number with ceil(maxVal / 3)
            maxHeap.add((maxVal + 2) / 3);  // Using (maxVal + 2) / 3 to simulate ceil
        }
        
        return score;
    }
}