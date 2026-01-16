import java.util.*;

class Solution {
    public int maximizeSquareArea(int m, int n, int[] hFences, int[] vFences) {
        long MOD = 1000000007L;

        List<Integer> h = new ArrayList<>();
        List<Integer> v = new ArrayList<>();

        for (int x : hFences)
            h.add(x);
        for (int x : vFences)
            v.add(x);

        h.add(1);
        h.add(m);
        v.add(1);
        v.add(n);

        Collections.sort(h);
        Collections.sort(v);

        Set<Long> horizontal = new HashSet<>();
        Set<Long> vertical = new HashSet<>();

        for (int i = 0; i < h.size(); i++) {
            for (int j = i + 1; j < h.size(); j++) {
                horizontal.add((long) h.get(j) - h.get(i));
            }
        }

        for (int i = 0; i < v.size(); i++) {
            for (int j = i + 1; j < v.size(); j++) {
                vertical.add((long) v.get(j) - v.get(i));
            }
        }

        long maxSide = 0;

        for (long d : horizontal) {
            if (vertical.contains(d)) {
                maxSide = Math.max(maxSide, d);
            }
        }

        if (maxSide == 0)
            return -1;

        return (int) ((maxSide * maxSide) % MOD);
    }
}
