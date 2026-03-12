class Solution
{
public:
    struct DSU
    {
        vector<int> parent, rank;

        DSU(int n)
        {
            parent.resize(n);
            rank.resize(n, 0);
            for (int i = 0; i < n; i++)
                parent[i] = i;
        }

        int find(int x)
        {
            if (parent[x] == x)
                return x;
            return parent[x] = find(parent[x]);
        }

        bool unite(int a, int b)
        {
            a = find(a);
            b = find(b);
            if (a == b)
                return false;

            if (rank[a] < rank[b])
                swap(a, b);

            parent[b] = a;
            if (rank[a] == rank[b])
                rank[a]++;

            return true;
        }
    };

    int maxStability(int n, vector<vector<int>> &edges, int k)
    {

        DSU dsu(n);

        int components = n;
        int mandatoryMin = INT_MAX;

        vector<vector<int>> optional;

        for (auto &e : edges)
        {
            int u = e[0], v = e[1], s = e[2], must = e[3];

            if (must)
            {
                if (!dsu.unite(u, v))
                    return -1;

                components--;
                mandatoryMin = min(mandatoryMin, s);
            }
            else
                optional.push_back(e);
        }

        sort(optional.begin(), optional.end(),
             [](auto &a, auto &b)
             {
                 return a[2] > b[2];
             });

        vector<int> used;

        for (auto &e : optional)
        {
            if (dsu.unite(e[0], e[1]))
            {
                used.push_back(e[2]);
                components--;
                if (components == 1)
                    break;
            }
        }

        if (components > 1)
            return -1;

        sort(used.begin(), used.end());

        int ans = mandatoryMin;

        for (int i = 0; i < used.size(); i++)
        {
            int val = used[i];

            if (k > 0)
            {
                val *= 2;
                k--;
            }

            if (ans == INT_MAX)
                ans = val;
            else
                ans = min(ans, val);
        }

        return ans;
    }
};