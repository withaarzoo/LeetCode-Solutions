class Solution {
    public int maxNumberOfFamilies(int n, int[][] reservedSeats) {
        // Store the reserved seats of each affected row as a bitmask.
        HashMap<Integer, Integer> rows = new HashMap<>();

        // Process every reserved seat once.
        for (int[] seat : reservedSeats) {
            int row = seat[0];
            int col = seat[1];

            // Only seats 2 through 9 can affect a four-person group.
            if (col >= 2 && col <= 9) {
                // Set the bit corresponding to this reserved seat.
                rows.put(row, rows.getOrDefault(row, 0) | (1 << col));
            }
        }

        // Rows not present in the map have no useful reserved seats.
        // Every such row can fit two groups.
        int answer = 2 * (n - rows.size());

        // Build masks for the three possible group positions.
        // left = seats 2-5, middle = seats 4-7, right = seats 6-9.
        int left = (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5);
        int middle = (1 << 4) | (1 << 5) | (1 << 6) | (1 << 7);
        int right = (1 << 6) | (1 << 7) | (1 << 8) | (1 << 9);

        // Check every row that has relevant reserved seats.
        for (int mask : rows.values()) {
            // Check whether every seat in each block is free.
            boolean canLeft = (mask & left) == 0;
            boolean canMiddle = (mask & middle) == 0;
            boolean canRight = (mask & right) == 0;

            // The left and right blocks are independent, so both can be used.
            if (canLeft && canRight) {
                answer += 2;
            }
            // Otherwise, any available block lets me place one group.
            else if (canLeft || canMiddle || canRight) {
                answer += 1;
            }
            // Otherwise, this row cannot accommodate a group.
        }

        // Return the maximum number of groups.
        return answer;
    }
}