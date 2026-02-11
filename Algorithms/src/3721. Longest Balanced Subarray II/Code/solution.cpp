/* C++ (optimized, commented) */
#include <bits/stdc++.h>
using namespace std;

struct SegTree
{
    int n;
    vector<int> mn, mx, lazy;
    SegTree(int _n) : n(_n), mn(4 * n, 0), mx(4 * n, 0), lazy(4 * n, 0) {}
    void apply(int idx, int v)
    {
        mn[idx] += v;
        mx[idx] += v;
        lazy[idx] += v;
    }
    void push(int idx)
    {
        if (lazy[idx] != 0)
        {
            apply(idx << 1, lazy[idx]);
            apply(idx << 1 | 1, lazy[idx]);
            lazy[idx] = 0;
        }
    }
    void pull(int idx)
    {
        mn[idx] = min(mn[idx << 1], mn[idx << 1 | 1]);
        mx[idx] = max(mx[idx << 1], mx[idx << 1 | 1]);
    }
    void add_range(int idx, int l, int r, int ql, int qr, int val)
    {
        if (ql > qr)
            return;
        if (ql <= l && r <= qr)
        {
            apply(idx, val);
            return;
        }
        push(idx);
        int mid = (l + r) >> 1;
        if (ql <= mid)
            add_range(idx << 1, l, mid, ql, min(qr, mid), val);
        if (qr > mid)
            add_range(idx << 1 | 1, mid + 1, r, max(ql, mid + 1), qr, val);
        pull(idx);
    }
    // public wrapper
    void add_range(int l, int r, int val)
    {
        if (l > r)
            return;
        add_range(1, 0, n - 1, l, r, val);
    }
    // find rightmost index in [ql, qr] with value == 0, or -1 if none
    int find_rightmost_zero(int idx, int l, int r, int ql, int qr)
    {
        if (ql > qr || qr < l || ql > r)
            return -1;
        if (mn[idx] > 0 || mx[idx] < 0)
            return -1; // no zero inside
        if (l == r)
        {
            if (mn[idx] == 0)
                return l;
            return -1;
        }
        push(idx);
        int mid = (l + r) >> 1;
        // try right child first to get rightmost
        if (qr > mid)
        {
            int res = find_rightmost_zero(idx << 1 | 1, mid + 1, r, max(ql, mid + 1), qr);
            if (res != -1)
                return res;
        }
        if (ql <= mid)
        {
            return find_rightmost_zero(idx << 1, l, mid, ql, min(qr, mid));
        }
        return -1;
    }
    int find_rightmost_zero(int ql, int qr)
    {
        if (ql > qr)
            return -1;
        return find_rightmost_zero(1, 0, n - 1, ql, qr);
    }
};

class Solution
{
public:
    int longestBalanced(vector<int> &nums)
    {
        int n = nums.size();
        unordered_map<int, vector<int>> pos;
        pos.reserve(n * 2);
        for (int i = 0; i < n; ++i)
            pos[nums[i]].push_back(i);

        SegTree st(n);
        // initial: for each value, add sign to [firstPos, n-1]
        for (auto &kv : pos)
        {
            int val = kv.first;
            int sign = (val & 1) ? 1 : -1;
            int p = kv.second[0];
            st.add_range(p, n - 1, sign);
        }

        // pointers to current first occurrence for each value
        unordered_map<int, int> ptr;
        ptr.reserve(pos.size() * 2);
        for (auto &kv : pos)
            ptr[kv.first] = 0;

        int ans = 0;
        for (int l = 0; l < n; ++l)
        {
            int r = st.find_rightmost_zero(l, n - 1);
            if (r != -1)
                ans = max(ans, r - l + 1);

            int x = nums[l];
            int pIndex = ptr[x]; // should point to l
            // move pointer forward
            ptr[x] = pIndex + 1;
            int nextPos = (ptr[x] < (int)pos[x].size()) ? pos[x][ptr[x]] : n;
            int sign = (x & 1) ? 1 : -1;
            // net effect: apply -sign to range [l, nextPos-1]
            int L = l, R = nextPos - 1;
            if (L <= R)
                st.add_range(L, R, -sign);
        }
        return ans;
    }
};
