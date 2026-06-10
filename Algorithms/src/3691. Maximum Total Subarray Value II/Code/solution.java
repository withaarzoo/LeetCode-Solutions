class Solution {
    static class Node {
        long val;
        int l;
        int r;

        Node(long val, int l, int r) {
            this.val = val;
            this.l = l;
            this.r = r;
        }
    }

    public long maxTotalValue(int[] nums, int k) {
        int n = nums.length;

        int[] lg = new int[n + 1];
        for (int i = 2; i <= n; i++) {
            lg[i] = lg[i / 2] + 1;
        }

        int K = lg[n] + 1;

        int[][] mx = new int[K][n];
        int[][] mn = new int[K][n];

        for (int i = 0; i < n; i++) {
            mx[0][i] = nums[i];
            mn[0][i] = nums[i];
        }

        for (int j = 1; j < K; j++) {
            for (int i = 0; i + (1 << j) <= n; i++) {
                mx[j][i] = Math.max(
                        mx[j - 1][i],
                        mx[j - 1][i + (1 << (j - 1))]);

                mn[j][i] = Math.min(
                        mn[j - 1][i],
                        mn[j - 1][i + (1 << (j - 1))]);
            }
        }

        PriorityQueue<Node> pq = new PriorityQueue<>(
                (a, b) -> Long.compare(b.val, a.val));

        for (int l = 0; l < n; l++) {
            pq.offer(new Node(getValue(l, n - 1, mx, mn, lg), l, n - 1));
        }

        long ans = 0;

        while (k-- > 0) {
            Node cur = pq.poll();

            ans += cur.val;

            if (cur.r > cur.l) {
                pq.offer(new Node(
                        getValue(cur.l, cur.r - 1, mx, mn, lg),
                        cur.l,
                        cur.r - 1));
            }
        }

        return ans;
    }

    private long getValue(
            int l,
            int r,
            int[][] mx,
            int[][] mn,
            int[] lg) {
        int len = r - l + 1;
        int p = lg[len];

        int mxVal = Math.max(
                mx[p][l],
                mx[p][r - (1 << p) + 1]);

        int mnVal = Math.min(
                mn[p][l],
                mn[p][r - (1 << p) + 1]);

        return (long) mxVal - mnVal;
    }
}