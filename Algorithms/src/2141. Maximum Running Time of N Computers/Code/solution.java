class Solution {
    public long maxRunTime(int n, int[] batteries) {
        long total = 0L;
        for (int b : batteries)
            total += b;

        long low = 0L;
        long high = total / n; // Upper bound on answer

        // Binary search on time
        while (low < high) {
            long mid = low + (high - low + 1) / 2; // upper mid

            long usable = 0L;
            for (int b : batteries) {
                // Each battery contributes at most mid minutes
                usable += Math.min((long) b, mid);
                if (usable >= mid * n)
                    break;
            }

            if (usable >= mid * n) {
                // mid minutes is possible
                low = mid;
            } else {
                // mid minutes is too large
                high = mid - 1;
            }
        }

        return low;
    }
}
