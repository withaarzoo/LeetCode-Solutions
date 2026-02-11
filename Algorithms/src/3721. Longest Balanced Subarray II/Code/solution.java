
// Java (clear, commented)
import java.util.*;

class Solution {
    static class SegTree {
        int n;
        int[] mn, mx, lazy;

        SegTree(int n) {
            this.n = n;
            mn = new int[4 * n];
            mx = new int[4 * n];
            lazy = new int[4 * n];
            // arrays default 0
        }

        void apply(int idx, int v) {
            mn[idx] += v;
            mx[idx] += v;
            lazy[idx] += v;
        }

        void push(int idx) {
            int z = lazy[idx];
            if (z != 0) {
                apply(idx << 1, z);
                apply(idx << 1 | 1, z);
                lazy[idx] = 0;
            }
        }

        void pull(int idx) {
            mn[idx] = Math.min(mn[idx << 1], mn[idx << 1 | 1]);
            mx[idx] = Math.max(mx[idx << 1], mx[idx << 1 | 1]);
        }

        void addRange(int idx, int l, int r, int ql, int qr, int val) {
            if (ql > qr)
                return;
            if (ql <= l && r <= qr) {
                apply(idx, val);
                return;
            }
            push(idx);
            int mid = (l + r) >> 1;
            if (ql <= mid)
                addRange(idx << 1, l, mid, ql, Math.min(qr, mid), val);
            if (qr > mid)
                addRange(idx << 1 | 1, mid + 1, r, Math.max(ql, mid + 1), qr, val);
            pull(idx);
        }

        void addRange(int l, int r, int v) {
            if (l > r)
                return;
            addRange(1, 0, n - 1, l, r, v);
        }

        int findRightmostZero(int idx, int l, int r, int ql, int qr) {
            if (ql > qr || qr < l || ql > r)
                return -1;
            if (mn[idx] > 0 || mx[idx] < 0)
                return -1;
            if (l == r) {
                return mn[idx] == 0 ? l : -1;
            }
            push(idx);
            int mid = (l + r) >> 1;
            if (qr > mid) {
                int res = findRightmostZero(idx << 1 | 1, mid + 1, r, Math.max(ql, mid + 1), qr);
                if (res != -1)
                    return res;
            }
            if (ql <= mid) {
                return findRightmostZero(idx << 1, l, mid, ql, Math.min(qr, mid));
            }
            return -1;
        }

        int findRightmostZero(int l, int r) {
            if (l > r)
                return -1;
            return findRightmostZero(1, 0, n - 1, l, r);
        }
    }

    public int longestBalanced(int[] nums) {
        int n = nums.length;
        HashMap<Integer, ArrayList<Integer>> pos = new HashMap<>();
        for (int i = 0; i < n; ++i) {
            pos.computeIfAbsent(nums[i], k -> new ArrayList<>()).add(i);
        }
        SegTree st = new SegTree(n);
        for (Map.Entry<Integer, ArrayList<Integer>> e : pos.entrySet()) {
            int val = e.getKey();
            int sign = (val % 2 == 1) ? 1 : -1;
            int p = e.getValue().get(0);
            st.addRange(p, n - 1, sign);
        }
        HashMap<Integer, Integer> ptr = new HashMap<>();
        for (int k : pos.keySet())
            ptr.put(k, 0);

        int ans = 0;
        for (int l = 0; l < n; ++l) {
            int r = st.findRightmostZero(l, n - 1);
            if (r != -1)
                ans = Math.max(ans, r - l + 1);

            int x = nums[l];
            int pIndex = ptr.get(x);
            ptr.put(x, pIndex + 1);
            ArrayList<Integer> lst = pos.get(x);
            int nextPos = (pIndex + 1 < lst.size()) ? lst.get(pIndex + 1) : n;
            int sign = (x % 2 == 1) ? 1 : -1;
            int L = l, R = nextPos - 1;
            if (L <= R)
                st.addRange(L, R, -sign);
        }
        return ans;
    }
}
