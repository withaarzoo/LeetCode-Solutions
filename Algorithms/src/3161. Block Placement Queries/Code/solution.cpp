class Fenwick
{
    int n;
    vector<int> bit;

public:
    Fenwick(int n = 0) : n(n), bit(n + 1, 0) {}

    // Add "val" at index "idx" because I need point updates on occupied positions.
    void add(int idx, int val)
    {
        for (; idx <= n; idx += idx & -idx)
        {
            bit[idx] += val;
        }
    }

    // Return prefix sum [1..idx] because I use it to count how many obstacles are on the left.
    int sum(int idx) const
    {
        int res = 0;
        for (; idx > 0; idx -= idx & -idx)
        {
            res += bit[idx];
        }
        return res;
    }

    // Find the smallest index whose prefix sum is at least k.
    // I use this to jump to the k-th occupied position.
    int kth(int k) const
    {
        int idx = 0;
        int step = 1;
        while ((step << 1) <= n)
            step <<= 1;

        for (int d = step; d > 0; d >>= 1)
        {
            int next = idx + d;
            if (next <= n && bit[next] < k)
            {
                idx = next;
                k -= bit[next];
            }
        }
        return idx + 1;
    }
};

class SegTree
{
    int n;
    vector<int> tree;

    // Update one position because only one obstacle gap changes at a time.
    void update(int node, int l, int r, int pos, int val)
    {
        if (l == r)
        {
            tree[node] = val;
            return;
        }
        int mid = (l + r) >> 1;
        if (pos <= mid)
            update(node << 1, l, mid, pos, val);
        else
            update(node << 1 | 1, mid + 1, r, pos, val);
        tree[node] = max(tree[node << 1], tree[node << 1 | 1]);
    }

    // Query the maximum on a prefix because I only care about gaps ending at or before x.
    int query(int node, int l, int r, int ql, int qr) const
    {
        if (ql > r || qr < l)
            return 0;
        if (ql <= l && r <= qr)
            return tree[node];
        int mid = (l + r) >> 1;
        return max(query(node << 1, l, mid, ql, qr),
                   query(node << 1 | 1, mid + 1, r, ql, qr));
    }

public:
    SegTree(int n = 0) : n(n), tree(4 * max(1, n), 0) {}

    void setVal(int pos, int val)
    {
        if (n == 0)
            return;
        update(1, 0, n - 1, pos, val);
    }

    int getMax(int l, int r) const
    {
        if (n == 0 || l > r)
            return 0;
        return query(1, 0, n - 1, l, r);
    }
};

class Solution
{
public:
    vector<bool> getResults(vector<vector<int>> &queries)
    {
        int mx = 0;
        for (auto &q : queries)
        {
            mx = max(mx, q[1]);
        }

        // I shift positions by +1 in Fenwick so that position 0 can still be stored safely.
        int fenwickSize = mx + 2;
        Fenwick fw(fenwickSize);

        // I store one gap value for every coordinate from 0 to mx.
        SegTree st(mx + 1);

        // Sentinel obstacle at 0 makes predecessor logic simple.
        fw.add(1, 1);

        vector<bool> ans;
        ans.reserve(queries.size());

        for (auto &q : queries)
        {
            int type = q[0];
            int x = q[1];

            if (type == 1)
            {
                // Count occupied positions strictly smaller than x.
                int leftCount = fw.sum(x);
                int leftPos = fw.kth(leftCount) - 1;

                // Count occupied positions up to x, then jump to the next one if it exists.
                int occupiedUpToX = fw.sum(x + 1);
                int totalOccupied = fw.sum(fenwickSize);
                int rightPos = -1;
                if (occupiedUpToX < totalOccupied)
                {
                    rightPos = fw.kth(occupiedUpToX + 1) - 1;
                }

                // The new obstacle creates the gap ending at x.
                st.setVal(x, x - leftPos);

                // The next obstacle's gap shrinks because x is now between them.
                if (rightPos != -1)
                    st.setVal(rightPos, rightPos - x);

                // Mark x as occupied.
                fw.add(x + 1, 1);
            }
            else
            {
                int sz = q[2];

                // The last obstacle strictly before x gives the tail gap ending at x.
                int leftCount = fw.sum(x);
                int leftPos = fw.kth(leftCount) - 1;

                // Any gap ending at a position <= x is fully inside [0, x].
                int bestPrefix = st.getMax(0, x);

                // Either the tail gap is enough, or some earlier gap is enough.
                ans.push_back((x - leftPos >= sz) || (bestPrefix >= sz));
            }
        }

        return ans;
    }
};