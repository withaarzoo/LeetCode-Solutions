import java.util.*;

class Solution {
    // Helper to make a string key from two ints
    private String key(int x, int y) {
        return x + "#" + y;
    }

    public int longestBalanced(String s) {
        int n = s.length();
        int a = 0, b = 0, c = 0;
        int ans = 0;

        // longest single-character run
        int run = 0;
        char prev = 0;
        for (int i = 0; i < n; ++i) {
            if (i == 0 || s.charAt(i) != prev)
                run = 1;
            else
                run++;
            prev = s.charAt(i);
            ans = Math.max(ans, run);
        }

        Map<String, Integer> map3 = new HashMap<>(); // (b-a, c-a)
        Map<String, Integer> map_ab_c = new HashMap<>(); // (b-a, c)
        Map<String, Integer> map_ac_b = new HashMap<>(); // (c-a, b)
        Map<String, Integer> map_bc_a = new HashMap<>(); // (c-b, a)

        map3.put(key(0, 0), 0);
        map_ab_c.put(key(0, 0), 0);
        map_ac_b.put(key(0, 0), 0);
        map_bc_a.put(key(0, 0), 0);

        for (int p = 1; p <= n; ++p) {
            char ch = s.charAt(p - 1);
            if (ch == 'a')
                a++;
            else if (ch == 'b')
                b++;
            else
                c++;

            String k3 = key(b - a, c - a);
            if (map3.containsKey(k3))
                ans = Math.max(ans, p - map3.get(k3));
            else
                map3.put(k3, p);

            String kabc = key(b - a, c);
            if (map_ab_c.containsKey(kabc))
                ans = Math.max(ans, p - map_ab_c.get(kabc));
            else
                map_ab_c.put(kabc, p);

            String kacb = key(c - a, b);
            if (map_ac_b.containsKey(kacb))
                ans = Math.max(ans, p - map_ac_b.get(kacb));
            else
                map_ac_b.put(kacb, p);

            String kbc = key(c - b, a);
            if (map_bc_a.containsKey(kbc))
                ans = Math.max(ans, p - map_bc_a.get(kbc));
            else
                map_bc_a.put(kbc, p);
        }

        return ans;
    }
}
