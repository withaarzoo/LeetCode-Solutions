class Solution
{
public:
    // Disjoint Set Union (Union-Find) with path compression + union by size
    struct DSU
    {
        vector<int> p, sz;
        DSU(int n = 0) : p(n + 1), sz(n + 1, 1)
        {
            iota(p.begin(), p.end(), 0);
        }
        int find(int x) { return p[x] == x ? x : p[x] = find(p[x]); }
        void unite(int a, int b)
        {
            a = find(a);
            b = find(b);
            if (a == b)
                return;
            if (sz[a] < sz[b])
                swap(a, b);
            p[b] = a;
            sz[a] += sz[b];
        }
    };

    vector<int> processQueries(int c, vector<vector<int>> &connections, vector<vector<int>> &queries)
    {
        DSU dsu(c);
        for (auto &e : connections)
            dsu.unite(e[0], e[1]);

        // Map root -> min-heap of all member ids
        unordered_map<int, priority_queue<int, vector<int>, greater<int>>> heap;
        heap.reserve(c * 2);

        for (int i = 1; i <= c; ++i)
        {
            int r = dsu.find(i);
            heap[r].push(i);
        }

        vector<char> offline(c + 1, false);
        vector<int> ans;
        ans.reserve(queries.size());

        for (auto &q : queries)
        {
            int t = q[0], x = q[1];
            if (t == 2)
            {
                offline[x] = true; // go offline
            }
            else
            { // t == 1
                if (!offline[x])
                {
                    ans.push_back(x); // it can resolve itself
                }
                else
                {
                    int r = dsu.find(x);
                    auto &pq = heap[r];
                    // Lazy delete: remove offline nodes from the top
                    while (!pq.empty() && offline[pq.top()])
                        pq.pop();
                    if (pq.empty())
                        ans.push_back(-1);
                    else
                        ans.push_back(pq.top());
                }
            }
        }
        return ans;
    }
};
