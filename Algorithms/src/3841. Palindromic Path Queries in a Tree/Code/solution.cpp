class Solution
{
public:
    vector<bool> palindromePath(int n, vector<vector<int>> &edges, string s, vector<string> &queries)
    {
        vector<vector<int>> g(n);
        for (auto &e : edges)
        {
            g[e[0]].push_back(e[1]);
            g[e[1]].push_back(e[0]);
        }
        const int LOG = 16;
        vector<vector<int>> up(LOG, vector<int>(n, -1));
        vector<int> depth(n), in(n), out(n), parent(n, -1);
        vector<int> mask(n);
        for (int i = 0; i < n; i++)
            mask[i] = 1 << (s[i] - 'a');
        int timer = 1;
        function<void(int, int)> dfs = [&](int u, int p)
        {
            in[u] = timer++;
            parent[u] = p;
            up[0][u] = p;
            for (int k = 1; k < LOG; k++)
            {
                if (up[k - 1][u] != -1)
                    up[k][u] = up[k - 1][up[k - 1][u]];
            }
            for (int v : g[u])
                if (v != p)
                {
                    depth[v] = depth[u] + 1;
                    dfs(v, u);
                }
            out[u] = timer;
        };
        dfs(0, -1);
        vector<int> bit(n + 2, 0);
        auto upd = [&](int i, int val)
        {
            for (; i <= n; i += i & -i)
                bit[i] ^= val;
        };
        auto qry = [&](int i)
        {
            int res = 0;
            for (; i > 0; i -= i & -i)
                res ^= bit[i];
            return res;
        };
        auto rangeXor = [&](int l, int r, int val)
        {
            if (l > r)
                return;
            upd(l, val);
            upd(r + 1, val);
        };
        for (int i = 0; i < n; i++)
            rangeXor(in[i], out[i] - 1, mask[i]);
        auto getLca = [&](int u, int v)
        {
            if (depth[u] < depth[v])
                swap(u, v);
            int diff = depth[u] - depth[v];
            for (int k = 0; k < LOG; k++)
                if (diff >> k & 1)
                    u = up[k][u];
            if (u == v)
                return u;
            for (int k = LOG - 1; k >= 0; k--)
            {
                if (up[k][u] != up[k][v])
                {
                    u = up[k][u];
                    v = up[k][v];
                }
            }
            return parent[u];
        };
        auto pathXor = [&](int u, int v)
        {
            int l = getLca(u, v);
            return qry(in[u]) ^ qry(in[v]) ^ mask[l];
        };
        vector<bool> ans;
        for (auto &q : queries)
        {
            if (q[0] == 'u')
            {
                int u = 0, i = 7;
                while (q[i] != ' ')
                    u = u * 10 + (q[i++] - '0');
                char c = q.back();
                int nm = 1 << (c - 'a');
                int delta = mask[u] ^ nm;
                mask[u] = nm;
                rangeXor(in[u], out[u] - 1, delta);
            }
            else
            {
                int u = 0, v = 0, i = 6;
                while (q[i] != ' ')
                    u = u * 10 + (q[i++] - '0');
                i++;
                while (i < (int)q.size())
                    v = v * 10 + (q[i++] - '0');
                int m = pathXor(u, v);
                ans.push_back(__builtin_popcount(m) <= 1);
            }
        }
        return ans;
    }
};