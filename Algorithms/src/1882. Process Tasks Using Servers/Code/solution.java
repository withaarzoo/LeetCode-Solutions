class Solution {
    public int[] assignTasks(int[] servers, int[] tasks) {
        PriorityQueue<int[]> free = new PriorityQueue<>((a, b) -> a[0] != b[0] ? a[0] - b[0] : a[1] - b[1]);
        for (int i = 0; i < servers.length; i++) {
            free.offer(new int[]{servers[i], i});
        }
        PriorityQueue<long[]> busy = new PriorityQueue<>((a, b) -> {
            if (a[0] != b[0]) return Long.compare(a[0], b[0]);
            if (a[1] != b[1]) return Long.compare(a[1], b[1]);
            return Long.compare(a[2], b[2]);
        });
        int m = tasks.length;
        int[] ans = new int[m];
        for (int j = 0; j < m; j++) {
            while (!busy.isEmpty() && busy.peek()[0] <= j) {
                long[] cur = busy.poll();
                free.offer(new int[]{(int)cur[1], (int)cur[2]});
            }
            if (!free.isEmpty()) {
                int[] cur = free.poll();
                ans[j] = cur[1];
                busy.offer(new long[]{j + (long)tasks[j], cur[0], cur[1]});
            } else {
                long[] cur = busy.poll();
                ans[j] = (int)cur[2];
                busy.offer(new long[]{cur[0] + tasks[j], cur[1], cur[2]});
            }
        }
        return ans;
    }
}