import java.util.*;

class Solution {
    static class P {
        long pos, x, y;

        P(long pos, long x, long y) {
            this.pos = pos;
            this.x = x;
            this.y = y;
        }
    }

    private long getOffset(long side, long x, long y, long d) {
        if (x == 0) { // left
            if (d <= 2L * side - y)
                return d;
            if (d <= side + y)
                return 2L * side + d - 2L * y;
            return -1;
        } else if (y == side) { // top
            if (d <= 2L * side - x)
                return d;
            if (d <= side + x)
                return 2L * side + d - 2L * x;
            return -1;
        } else if (x == side) { // right
            if (d <= side + y)
                return d;
            if (d <= 2L * side - y)
                return d + 2L * y;
            return -1;
        } else { // bottom
            if (d <= side + x)
                return d;
            if (d <= 2L * side - x)
                return d + 2L * x;
            return -1;
        }
    }

    private int lowerBound(long[] arr, int l, int r, long target) {
        while (l < r) {
            int m = (l + r) >>> 1;
            if (arr[m] < target)
                l = m + 1;
            else
                r = m;
        }
        return l;
    }

    private boolean can(long side, P[] pts, int k, long d) {
        int n = pts.length;

        long[] pos3 = new long[3 * n];
        for (int i = 0; i < n; i++) {
            pos3[i] = pts[i].pos;
            pos3[i + n] = pts[i].pos + 4L * side;
            pos3[i + 2 * n] = pts[i].pos + 8L * side;
        }

        int[] nxt = new int[2 * n];
        Arrays.fill(nxt, -1);

        for (int i = 0; i < 2 * n; i++) {
            P p = pts[i % n];
            long off = getOffset(side, p.x, p.y, d);
            if (off < 0)
                continue;

            long target = pos3[i] + off;
            int hi = Math.min(i + n, 3 * n);
            int j = lowerBound(pos3, i + 1, hi, target);
            if (j < hi)
                nxt[i] = j;
        }

        for (int start = 0; start < n; start++) {
            int cur = start;
            int cnt = 1;

            while (cnt < k) {
                cur = nxt[cur];
                if (cur == -1 || cur >= start + n)
                    break;
                cnt++;
            }

            if (cnt >= k) {
                long dx = Math.abs(pts[start].x - pts[cur % n].x);
                long dy = Math.abs(pts[start].y - pts[cur % n].y);
                if (dx + dy >= d)
                    return true;
            }
        }

        return false;
    }

    public int maxDistance(int side, int[][] points, int k) {
        P[] pts = new P[points.length];

        for (int i = 0; i < points.length; i++) {
            long x = points[i][0];
            long y = points[i][1];
            long pos;
            if (x == 0)
                pos = y;
            else if (y == side)
                pos = 1L * side + x;
            else if (x == side)
                pos = 3L * side - y;
            else
                pos = 4L * side - x;
            pts[i] = new P(pos, x, y);
        }

        Arrays.sort(pts, Comparator.comparingLong(a -> a.pos));

        long lo = 0, hi = 2L * side;
        while (lo < hi) {
            long mid = (lo + hi + 1) >>> 1;
            if (can(side, pts, k, mid))
                lo = mid;
            else
                hi = mid - 1;
        }

        return (int) lo;
    }
}