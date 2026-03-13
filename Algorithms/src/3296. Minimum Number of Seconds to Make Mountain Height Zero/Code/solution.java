class Solution {

    private boolean can(long time, int mountainHeight, int[] workerTimes) {

        long totalHeight = 0;

        for (int t : workerTimes) {

            long left = 0, right = mountainHeight;

            while (left <= right) {

                long mid = (left + right) / 2;

                long required = (long) t * (mid * (mid + 1) / 2);

                if (required <= time) {
                    left = mid + 1;
                } else {
                    right = mid - 1;
                }
            }

            totalHeight += right;

            if (totalHeight >= mountainHeight)
                return true;
        }

        return false;
    }

    public long minNumberOfSeconds(int mountainHeight, int[] workerTimes) {

        long left = 1, right = (long) 1e18;
        long ans = right;

        while (left <= right) {

            long mid = (left + right) / 2;

            if (can(mid, mountainHeight, workerTimes)) {
                ans = mid;
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }

        return ans;
    }
}