class Solution {
    static class Node {
        int prod;
        int[] freq;
        Node(int k) {
            prod = 1;
            freq = new int[k];
        }
    }
    
    int k;
    Node[] tree;
    int[] nums;
    
    Node merge(Node L, Node R) {
        Node res = new Node(k);
        res.prod = (int)((1L * L.prod * R.prod) % k);
        System.arraycopy(L.freq, 0, res.freq, 0, k);
        for (int r = 0; r < k; r++) {
            if (R.freq[r] != 0) {
                int nr = (int)((1L * L.prod * r) % k);
                res.freq[nr] += R.freq[r];
            }
        }
        return res;
    }
    
    void build(int v, int tl, int tr) {
        if (tl == tr) {
            tree[v].prod = nums[tl] % k;
            tree[v].freq[tree[v].prod] = 1;
            return;
        }
        int tm = (tl + tr) / 2;
        build(v * 2, tl, tm);
        build(v * 2 + 1, tm + 1, tr);
        tree[v] = merge(tree[v * 2], tree[v * 2 + 1]);
    }
    
    void update(int v, int tl, int tr, int pos, int val) {
        if (tl == tr) {
            tree[v].prod = val % k;
            Arrays.fill(tree[v].freq, 0);
            tree[v].freq[tree[v].prod] = 1;
            return;
        }
        int tm = (tl + tr) / 2;
        if (pos <= tm) update(v * 2, tl, tm, pos, val);
        else update(v * 2 + 1, tm + 1, tr, pos, val);
        tree[v] = merge(tree[v * 2], tree[v * 2 + 1]);
    }
    
    Node query(int v, int tl, int tr, int l, int r) {
        if (l > r) return new Node(k);
        if (l == tl && r == tr) return tree[v];
        int tm = (tl + tr) / 2;
        return merge(query(v * 2, tl, tm, l, Math.min(r, tm)),
                     query(v * 2 + 1, tm + 1, tr, Math.max(l, tm + 1), r));
    }
    
    public int[] resultArray(int[] nums, int k, int[][] queries) {
        this.k = k;
        this.nums = nums;
        int n = nums.length;
        tree = new Node[4 * n];
        for (int i = 0; i < tree.length; i++) tree[i] = new Node(k);
        build(1, 0, n - 1);
        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int idx = queries[i][0], val = queries[i][1], start = queries[i][2], x = queries[i][3];
            update(1, 0, n - 1, idx, val);
            Node res = query(1, 0, n - 1, start, n - 1);
            ans[i] = res.freq[x];
        }
        return ans;
    }
}