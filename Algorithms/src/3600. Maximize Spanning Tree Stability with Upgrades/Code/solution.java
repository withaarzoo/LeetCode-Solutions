class Solution {

    class DSU {
        int[] parent, rank;

        DSU(int n) {
            parent = new int[n];
            rank = new int[n];

            for (int i = 0; i < n; i++)
                parent[i] = i;
        }

        int find(int x) {
            if (parent[x] == x)
                return x;
            return parent[x] = find(parent[x]);
        }

        boolean union(int a, int b) {
            a = find(a);
            b = find(b);

            if (a == b)
                return false;

            if (rank[a] < rank[b]) {
                int t = a;
                a = b;
                b = t;
            }

            parent[b] = a;

            if (rank[a] == rank[b])
                rank[a]++;

            return true;
        }
    }

    public int maxStability(int n, int[][] edges, int k) {

        DSU dsu = new DSU(n);

        int comp = n;
        int mandatoryMin = Integer.MAX_VALUE;

        List<int[]> optional = new ArrayList<>();

        for (int[] e : edges) {
            if (e[3] == 1) {
                if (!dsu.union(e[0], e[1]))
                    return -1;

                comp--;
                mandatoryMin = Math.min(mandatoryMin, e[2]);
            } else
                optional.add(e);
        }

        optional.sort((a, b) -> b[2] - a[2]);

        List<Integer> used = new ArrayList<>();

        for (int[] e : optional) {
            if (dsu.union(e[0], e[1])) {
                used.add(e[2]);
                comp--;
                if (comp == 1)
                    break;
            }
        }

        if (comp > 1)
            return -1;

        Collections.sort(used);

        int ans = mandatoryMin;

        for (int w : used) {
            int val = w;

            if (k > 0) {
                val *= 2;
                k--;
            }

            if (ans == Integer.MAX_VALUE)
                ans = val;
            else
                ans = Math.min(ans, val);
        }

        return ans;
    }
}