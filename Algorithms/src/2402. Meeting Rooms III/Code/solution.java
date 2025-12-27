class Solution {
    public int mostBooked(int n, int[][] meetings) {
        Arrays.sort(meetings, (a, b) -> a[0] - b[0]);

        PriorityQueue<Integer> freeRooms = new PriorityQueue<>();
        for (int i = 0; i < n; i++)
            freeRooms.offer(i);

        PriorityQueue<long[]> busyRooms = new PriorityQueue<>(
                (a, b) -> a[0] == b[0] ? Long.compare(a[1], b[1]) : Long.compare(a[0], b[0]));

        long[] count = new long[n];

        for (int[] m : meetings) {
            long start = m[0], end = m[1], duration = end - start;

            while (!busyRooms.isEmpty() && busyRooms.peek()[0] <= start) {
                freeRooms.offer((int) busyRooms.poll()[1]);
            }

            if (!freeRooms.isEmpty()) {
                int room = freeRooms.poll();
                busyRooms.offer(new long[] { end, room });
                count[room]++;
            } else {
                long[] top = busyRooms.poll();
                busyRooms.offer(new long[] { top[0] + duration, top[1] });
                count[(int) top[1]]++;
            }
        }

        int ans = 0;
        for (int i = 1; i < n; i++) {
            if (count[i] > count[ans])
                ans = i;
        }
        return ans;
    }
}
