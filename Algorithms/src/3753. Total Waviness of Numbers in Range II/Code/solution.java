class Solution {

    static class Node {
        long cnt;
        long wav;

        Node(long cnt, long wav) {
            this.cnt = cnt;
            this.wav = wav;
        }
    }

    String s;
    Node[][][][] dp = new Node[20][2][11][11];
    boolean[][][][] vis = new boolean[20][2][11][11];

    private Node dfs(int pos, int started, int last,
            int secondLast, boolean tight) {

        if (pos == s.length()) {
            return new Node(1, 0);
        }

        if (!tight && vis[pos][started][last][secondLast]) {
            return dp[pos][started][last][secondLast];
        }

        int limit = tight ? s.charAt(pos) - '0' : 9;

        long totalCnt = 0;
        long totalWav = 0;

        for (int d = 0; d <= limit; d++) {
            boolean ntight = tight && (d == limit);

            if (started == 0 && d == 0) {
                Node nxt = dfs(pos + 1, 0, 10, 10, ntight);

                totalCnt += nxt.cnt;
                totalWav += nxt.wav;
            } else {
                long add = 0;

                if (started == 1 && secondLast != 10) {
                    if ((last > secondLast && last > d) ||
                            (last < secondLast && last < d)) {
                        add = 1;
                    }
                }

                int nSecondLast = (started == 1) ? last : 10;

                Node nxt = dfs(pos + 1, 1, d, nSecondLast, ntight);

                totalCnt += nxt.cnt;
                totalWav += nxt.wav + add * nxt.cnt;
            }
        }

        Node res = new Node(totalCnt, totalWav);

        if (!tight) {
            vis[pos][started][last][secondLast] = true;
            dp[pos][started][last][secondLast] = res;
        }

        return res;
    }

    private long solve(long n) {
        if (n < 0)
            return 0;

        s = Long.toString(n);

        vis = new boolean[20][2][11][11];
        dp = new Node[20][2][11][11];

        return dfs(0, 0, 10, 10, true).wav;
    }

    public long totalWaviness(long num1, long num2) {
        return solve(num2) - solve(num1 - 1);
    }
}