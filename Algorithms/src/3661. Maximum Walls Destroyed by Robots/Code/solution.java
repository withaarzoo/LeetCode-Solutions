class Solution {
    public int maxWalls(int[] robots, int[] distance, int[] walls) {
        int n = robots.length;

        int[][] arr = new int[n + 1][2];

        for (int i = 0; i < n; i++) {
            arr[i][0] = robots[i];
            arr[i][1] = distance[i];
        }

        Arrays.sort(arr, 0, n, (a, b) -> Integer.compare(a[0], b[0]));
        Arrays.sort(walls);

        // Dummy robot
        arr[n][0] = (int) 1e9;
        arr[n][1] = 0;

        int[][] dp = new int[n][2];

        dp[0][0] = countWalls(walls, arr[0][0] - arr[0][1], arr[0][0]);

        int firstRightEnd = (n == 1)
                ? arr[0][0] + arr[0][1]
                : Math.min(arr[0][0] + arr[0][1], arr[1][0] - 1);

        dp[0][1] = countWalls(walls, arr[0][0], firstRightEnd);

        for (int i = 1; i < n; i++) {
            int pos = arr[i][0];
            int dist = arr[i][1];

            // Shoot right
            int rightEnd = Math.min(pos + dist, arr[i + 1][0] - 1);
            int rightWalls = countWalls(walls, pos, rightEnd);

            dp[i][1] = Math.max(dp[i - 1][0], dp[i - 1][1]) + rightWalls;

            // Shoot left
            int leftStart = Math.max(pos - dist, arr[i - 1][0] + 1);
            int leftWalls = countWalls(walls, leftStart, pos);

            dp[i][0] = dp[i - 1][0] + leftWalls;

            int prevRightEnd = Math.min(arr[i - 1][0] + arr[i - 1][1], pos - 1);

            int overlapStart = leftStart;
            int overlapEnd = Math.min(prevRightEnd, pos - 1);

            int overlapWalls = countWalls(walls, overlapStart, overlapEnd);

            dp[i][0] = Math.max(dp[i][0], dp[i - 1][1] + leftWalls - overlapWalls);
        }

        return Math.max(dp[n - 1][0], dp[n - 1][1]);
    }

    private int countWalls(int[] walls, int left, int right) {
        if (left > right)
            return 0;

        int l = lowerBound(walls, left);
        int r = upperBound(walls, right);

        return r - l;
    }

    private int lowerBound(int[] arr, int target) {
        int left = 0, right = arr.length;

        while (left < right) {
            int mid = left + (right - left) / 2;

            if (arr[mid] < target) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }

        return left;
    }

    private int upperBound(int[] arr, int target) {
        int left = 0, right = arr.length;

        while (left < right) {
            int mid = left + (right - left) / 2;

            if (arr[mid] <= target) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }

        return left;
    }
}