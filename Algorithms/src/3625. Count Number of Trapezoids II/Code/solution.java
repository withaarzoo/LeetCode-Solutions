import java.util.*;

class Solution {
    private static final int SHIFT = 3000; // for packing

    private int encodePair(int a, int b) {
        return ((a + SHIFT) << 13) ^ (b + SHIFT);
    }

    private int myGcd(int a, int b) {
        a = Math.abs(a);
        b = Math.abs(b);
        while (b != 0) {
            int t = a % b;
            a = b;
            b = t;
        }
        return a;
    }

    private long countPairs(Map<Integer, Map<Integer, Integer>> mp) {
        long res = 0;
        for (Map<Integer, Integer> inner : mp.values()) {
            long sum = 0;
            long sumSq = 0;
            for (int c : inner.values()) {
                long cc = c;
                sum += cc;
                sumSq += cc * cc;
            }
            res += (sum * sum - sumSq) / 2;
        }
        return res;
    }

    public int countTrapezoids(int[][] points) {
        int n = points.length;

        Map<Integer, Map<Integer, Integer>> bySlope = new HashMap<>();
        Map<Integer, Map<Integer, Integer>> byVector = new HashMap<>();

        for (int i = 0; i < n; ++i) {
            int x1 = points[i][0];
            int y1 = points[i][1];
            for (int j = i + 1; j < n; ++j) {
                int x2 = points[j][0];
                int y2 = points[j][1];

                int dx = x2 - x1;
                int dy = y2 - y1;

                if (dx < 0 || (dx == 0 && dy < 0)) {
                    dx = -dx;
                    dy = -dy;
                }

                int g = myGcd(dx, dy);
                int ux = dx / g;
                int uy = dy / g;

                int lineId = ux * y1 - uy * x1;

                int slopeKey = encodePair(ux, uy);
                int vectorKey = encodePair(dx, dy);

                bySlope
                        .computeIfAbsent(slopeKey, k -> new HashMap<>())
                        .merge(lineId, 1, Integer::sum);

                byVector
                        .computeIfAbsent(vectorKey, k -> new HashMap<>())
                        .merge(lineId, 1, Integer::sum);
            }
        }

        long withParallel = countPairs(bySlope);
        long parallelogramTwo = countPairs(byVector);
        long ans = withParallel - parallelogramTwo / 2;

        return (int) ans;
    }
}
