class Solution
{
public:
    vector<int> parent, rankArr;

    int find(int x)
    {
        if (parent[x] == x)
            return x;
        return parent[x] = find(parent[x]);
    }

    void unite(int a, int b)
    {
        int pa = find(a);
        int pb = find(b);

        if (pa == pb)
            return;

        if (rankArr[pa] < rankArr[pb])
        {
            parent[pa] = pb;
        }
        else if (rankArr[pb] < rankArr[pa])
        {
            parent[pb] = pa;
        }
        else
        {
            parent[pb] = pa;
            rankArr[pa]++;
        }
    }

    int minimumHammingDistance(vector<int> &source, vector<int> &target, vector<vector<int>> &allowedSwaps)
    {
        int n = source.size();

        parent.resize(n);
        rankArr.resize(n, 0);

        for (int i = 0; i < n; i++)
        {
            parent[i] = i;
        }

        // Build connected components using DSU
        for (auto &swap : allowedSwaps)
        {
            unite(swap[0], swap[1]);
        }

        // Group indices by component parent
        unordered_map<int, vector<int>> groups;
        for (int i = 0; i < n; i++)
        {
            groups[find(i)].push_back(i);
        }

        int answer = 0;

        // Process each connected component
        for (auto &entry : groups)
        {
            unordered_map<int, int> freq;

            // Count source values in this component
            for (int idx : entry.second)
            {
                freq[source[idx]]++;
            }

            // Match target values
            for (int idx : entry.second)
            {
                if (freq[target[idx]] > 0)
                {
                    freq[target[idx]]--;
                }
                else
                {
                    answer++;
                }
            }
        }

        return answer;
    }
};