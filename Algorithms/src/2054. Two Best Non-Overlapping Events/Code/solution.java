class Solution {
    public int maxTwoEvents(int[][] events) {
        Arrays.sort(events, (a, b) -> a[0] - b[0]);

        int[][] endSorted = events.clone();
        Arrays.sort(endSorted, (a, b) -> a[1] - b[1]);

        int n = events.length;
        int[] maxValueTill = new int[n];

        maxValueTill[0] = endSorted[0][2];
        for (int i = 1; i < n; i++) {
            maxValueTill[i] = Math.max(maxValueTill[i - 1], endSorted[i][2]);
        }

        int ans = 0;
        int j = 0;

        for (int i = 0; i < n; i++) {
            int start = events[i][0];
            int value = events[i][2];

            while (j < n && endSorted[j][1] < start) {
                j++;
            }

            ans = Math.max(ans, value);
            if (j > 0) {
                ans = Math.max(ans, value + maxValueTill[j - 1]);
            }
        }

        return ans;
    }
}
