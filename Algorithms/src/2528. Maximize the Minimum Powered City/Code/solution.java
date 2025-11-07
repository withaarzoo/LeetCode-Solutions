class Solution {
    public long maxPower(int[] stations, int r, int k) {
        int n = stations.length;

        // 1) Base power via difference array.
        long[] diff = new long[n + 1];
        for (int i = 0; i < n; i++) {
            int L = Math.max(0, i - r);
            int R = Math.min(n, i + r + 1);
            diff[L] += stations[i];
            diff[R] -= stations[i];
        }
        long[] base = new long[n];
        long run = 0;
        for (int i = 0; i < n; i++) {
            run += diff[i];
            base[i] = run;
        }

        long lo = 0, hi = 0, ans = 0;
        for (int v : stations)
            hi += v;
        hi += k;

        // 2) Feasibility check with greedy & difference.
        java.util.function.LongPredicate feasible = T -> {
            long[] add = new long[n + 1];
            long extra = 0, used = 0;
            for (int i = 0; i < n; i++) {
                extra += add[i];
                long curr = base[i] + extra;
                if (curr < T) {
                    long need = T - curr;
                    used += need;
                    if (used > k)
                        return false;
                    extra += need;
                    int end = Math.min(n, i + 2 * r + 1);
                    add[end] -= need;
                }
            }
            return true;
        };

        while (lo <= hi) {
            long mid = lo + ((hi - lo) >> 1);
            if (feasible.test(mid)) {
                ans = mid;
                lo = mid + 1;
            } else {
                hi = mid - 1;
            }
        }
        return ans;
    }
}
