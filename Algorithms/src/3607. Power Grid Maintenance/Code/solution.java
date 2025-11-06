import java.util.*;

class Solution {
    static class DSU {
        int[] p, sz;

        DSU(int n) {
            p = new int[n + 1];
            sz = new int[n + 1];
            for (int i = 0; i <= n; i++) {
                p[i] = i;
                sz[i] = 1;
            }
        }

        int find(int x) {
            return p[x] == x ? x : (p[x] = find(p[x]));
        }

        void unite(int a, int b) {
            a = find(a);
            b = find(b);
            if (a == b)
                return;
            if (sz[a] < sz[b]) {
                int tmp = a;
                a = b;
                b = tmp;
            }
            p[b] = a;
            sz[a] += sz[b];
        }
    }

    public int[] processQueries(int c, int[][] connections, int[][] queries) {
        DSU dsu = new DSU(c);
        for (int[] e : connections)
            dsu.unite(e[0], e[1]);

        Map<Integer, PriorityQueue<Integer>> heap = new HashMap<>();
        for (int i = 1; i <= c; i++) {
            int r = dsu.find(i);
            heap.computeIfAbsent(r, k -> new PriorityQueue<>()).offer(i);
        }

        boolean[] offline = new boolean[c + 1];
        List<Integer> res = new ArrayList<>(queries.length);

        for (int[] q : queries) {
            int t = q[0], x = q[1];
            if (t == 2) {
                offline[x] = true;
            } else {
                if (!offline[x]) {
                    res.add(x);
                } else {
                    int r = dsu.find(x);
                    PriorityQueue<Integer> pq = heap.get(r);
                    if (pq == null) {
                        res.add(-1);
                        continue;
                    }
                    while (!pq.isEmpty() && offline[pq.peek()])
                        pq.poll();
                    res.add(pq.isEmpty() ? -1 : pq.peek());
                }
            }
        }

        int[] ans = new int[res.size()];
        for (int i = 0; i < res.size(); i++)
            ans[i] = res.get(i);
        return ans;
    }
}
