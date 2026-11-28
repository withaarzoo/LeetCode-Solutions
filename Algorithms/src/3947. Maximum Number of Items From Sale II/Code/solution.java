class Solution {
    public int maximumSaleItems(int[][] items, int budget) {
        int n = items.length;
        int[] cnt = new int[n + 1];
        for (int[] it : items) cnt[it[0]]++;
        int[] mult = new int[n + 1];
        for (int f = 1; f <= n; f++) {
            for (int m = f; m <= n; m += f) mult[f] += cnt[m];
        }
        int minP = Integer.MAX_VALUE;
        List<int[]> cand = new ArrayList<>();
        for (int[] it : items) {
            int f = it[0], p = it[1];
            minP = Math.min(minP, p);
            int g = mult[f] - 1;
            if (g > 0) cand.add(new int[]{p, g});
        }
        cand.sort((a, b) -> Integer.compare(a[0], b[0]));
        long spent = 0, total = 0, B = budget;
        for (int[] c : cand) {
            int p = c[0], g = c[1];
            if (p > 2L * minP) break;
            long can = Math.min((long)g, (B - spent) / p);
            total += can * 2;
            spent += can * p;
        }
        total += (B - spent) / minP;
        return (int)total;
    }
}