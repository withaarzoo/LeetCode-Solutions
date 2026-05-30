import java.util.*;

class Fenwick {
    int n;
    int[] bit;

    Fenwick(int n) {
        this.n = n;
        this.bit = new int[n + 1];
    }

    // Point add because each obstacle insertion changes one count.
    void add(int idx, int val) {
        for (; idx <= n; idx += idx & -idx) {
            bit[idx] += val;
        }
    }

    // Prefix sum because I need counts of occupied positions on the left.
    int sum(int idx) {
        int res = 0;
        for (; idx > 0; idx -= idx & -idx) {
            res += bit[idx];
        }
        return res;
    }

    // Find the smallest index whose prefix sum is at least k.
    int kth(int k) {
        int idx = 0;
        int step = 1;
        while ((step << 1) <= n)
            step <<= 1;

        for (int d = step; d > 0; d >>= 1) {
            int next = idx + d;
            if (next <= n && bit[next] < k) {
                idx = next;
                k -= bit[next];
            }
        }
        return idx + 1;
    }
}

class SegTree {
    int n;
    int[] tree;

    SegTree(int n) {
        this.n = n;
        this.tree = new int[Math.max(4, 4 * n)];
    }

    // Update one coordinate because only one gap value changes at a time.
    void update(int node, int l, int r, int pos, int val) {
        if (l == r) {
            tree[node] = val;
            return;
        }
        int mid = (l + r) >>> 1;
        if (pos <= mid)
            update(node << 1, l, mid, pos, val);
        else
            update(node << 1 | 1, mid + 1, r, pos, val);
        tree[node] = Math.max(tree[node << 1], tree[node << 1 | 1]);
    }

    // Query a prefix maximum because all valid stored gaps end inside [0, x].
    int query(int node, int l, int r, int ql, int qr) {
        if (ql > r || qr < l)
            return 0;
        if (ql <= l && r <= qr)
            return tree[node];
        int mid = (l + r) >>> 1;
        return Math.max(query(node << 1, l, mid, ql, qr),
                query(node << 1 | 1, mid + 1, r, ql, qr));
    }
}

class Solution {
    public List<Boolean> getResults(int[][] queries) {
        int mx = 0;
        for (int[] q : queries) {
            mx = Math.max(mx, q[1]);
        }

        int fenwickSize = mx + 2;
        Fenwick fw = new Fenwick(fenwickSize);
        SegTree st = new SegTree(mx + 1);

        // Sentinel obstacle at 0 makes predecessor search easier.
        fw.add(1, 1);

        List<Boolean> ans = new ArrayList<>();

        for (int[] q : queries) {
            int type = q[0];
            int x = q[1];

            if (type == 1) {
                // Count obstacles strictly smaller than x.
                int leftCount = fw.sum(x);
                int leftPos = fw.kth(leftCount) - 1;

                // Find the next obstacle after x if it exists.
                int occupiedUpToX = fw.sum(x + 1);
                int totalOccupied = fw.sum(fenwickSize);
                int rightPos = -1;
                if (occupiedUpToX < totalOccupied) {
                    rightPos = fw.kth(occupiedUpToX + 1) - 1;
                }

                // New obstacle x becomes the end of a new gap.
                st.update(1, 0, mx, x, x - leftPos);

                // The following obstacle now sees a shorter gap.
                if (rightPos != -1) {
                    st.update(1, 0, mx, rightPos, rightPos - x);
                }

                // Mark x as occupied.
                fw.add(x + 1, 1);
            } else {
                int sz = q[2];

                // The last obstacle strictly before x gives the free tail up to x.
                int leftCount = fw.sum(x);
                int leftPos = fw.kth(leftCount) - 1;

                // Prefix max over gaps that end at or before x.
                int bestPrefix = st.query(1, 0, mx, 0, x);

                ans.add((x - leftPos >= sz) || (bestPrefix >= sz));
            }
        }

        return ans;
    }
}