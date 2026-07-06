class Solution {
    public int removeCoveredIntervals(int[][] intervals) {
        // Sort by start in ascending order.
        // For equal starts, place the larger end first
        // so the covering interval is processed before the covered one.
        Arrays.sort(intervals, (a, b) -> {
            if (a[0] == b[0]) {
                return Integer.compare(b[1], a[1]);
            }
            return Integer.compare(a[0], b[0]);
        });

        // Every valid end is greater than 0, so -1 is a safe initial value.
        int maxEnd = -1;

        // This stores how many intervals are not covered.
        int remaining = 0;

        // Check every interval in the sorted order.
        for (int[] interval : intervals) {
            // A larger end means no earlier interval can cover this one.
            if (interval[1] > maxEnd) {
                remaining++;

                // Remember the farthest ending point seen so far.
                maxEnd = interval[1];
            }
            // Otherwise, the current interval is covered, so I skip it.
        }

        // Return the number of intervals that remain.
        return remaining;
    }
}