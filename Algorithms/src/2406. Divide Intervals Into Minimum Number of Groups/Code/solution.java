import java.util.*;

class Solution {
    public int minGroups(int[][] intervals) {
        // Sort intervals by start time
        Arrays.sort(intervals, (a, b) -> a[0] - b[0]);

        // Min-heap to track end times of active groups
        PriorityQueue<Integer> pq = new PriorityQueue<>();

        // Traverse through all intervals
        for (int[] interval : intervals) {
            int start = interval[0], end = interval[1];

            // If the top of the heap (earliest end time) is less than the current start,
            // we can reuse that group
            if (!pq.isEmpty() && pq.peek() < start) {
                pq.poll();
            }

            // Add the current interval's end time to the heap
            pq.add(end);
        }

        // The size of the heap represents the number of groups
        return pq.size();
    }
}
