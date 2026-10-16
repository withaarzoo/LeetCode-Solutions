class Solution {
public:
    static const int MOD = 1e9 + 7;
    struct Node {
        long long sum = 0, sumsq = 0, lazy = 0;
    };
    vector<Node> tree;
    int n;
    void push(int v, int tl, int tr) {
        if (tree[v].lazy && tl != tr) {
            int mid = (tl + tr) >> 1;
            int llen = mid - tl + 1, rlen = tr - mid;
            long long add = tree[v].lazy;
            tree[v << 1].sumsq = (tree[v << 1].sumsq + 2 * add * tree[v << 1].sum + add * add % MOD * llen) % MOD;
            tree[v << 1].sum = (tree[v << 1].sum + add * llen) % MOD;
            tree[v << 1].lazy = (tree[v << 1].lazy + add) % MOD;
            tree[v << 1 | 1].sumsq = (tree[v << 1 | 1].sumsq + 2 * add * tree[v << 1 | 1].sum + add * add % MOD * rlen) % MOD;
            tree[v << 1 | 1].sum = (tree[v << 1 | 1].sum + add * rlen) % MOD;
            tree[v << 1 | 1].lazy = (tree[v << 1 | 1].lazy + add) % MOD;
            tree[v].lazy = 0;
        }
    }
    void update(int v, int tl, int tr, int l, int r, long long add) {
        if (l > r) return;
        if (l == tl && r == tr) {
            int len = tr - tl + 1;
            tree[v].sumsq = (tree[v].sumsq + 2 * add * tree[v].sum + add * add % MOD * len) % MOD;
            tree[v].sum = (tree[v].sum + add * len) % MOD;
            tree[v].lazy = (tree[v].lazy + add) % MOD;
            return;
        }
        push(v, tl, tr);
        int mid = (tl + tr) >> 1;
        update(v << 1, tl, mid, l, min(r, mid), add);
        update(v << 1 | 1, mid + 1, tr, max(l, mid + 1), r, add);
        tree[v].sum = (tree[v << 1].sum + tree[v << 1 | 1].sum) % MOD;
        tree[v].sumsq = (tree[v << 1].sumsq + tree[v << 1 | 1].sumsq) % MOD;
    }
    long long query(int v, int tl, int tr, int l, int r) {
        if (l > r) return 0;
        if (l == tl && r == tr) return tree[v].sumsq;
        push(v, tl, tr);
        int mid = (tl + tr) >> 1;
        return (query(v << 1, tl, mid, l, min(r, mid)) + query(v << 1 | 1, mid + 1, tr, max(l, mid + 1), r)) % MOD;
    }
    int sumCounts(vector<int>& nums) {
        n = nums.size();
        tree.assign(4 * n, Node());
        vector<int> last(100001, -1);
        long long ans = 0;
        for (int r = 0; r < n; ++r) {
            int prv = last[nums[r]];
            if (prv + 1 <= r - 1) update(1, 0, n - 1, prv + 1, r - 1, 1);
            update(1, 0, n - 1, r, r, 1);
            ans = (ans + query(1, 0, n - 1, 0, r)) % MOD;
            last[nums[r]] = r;
        }
        return (int)ans;
    }
};