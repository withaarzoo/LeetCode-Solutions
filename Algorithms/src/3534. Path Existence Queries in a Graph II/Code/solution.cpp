class Solution
{
public:
    vector<int> pathExistenceQueries(int n, vector<int> &nums, int maxDiff, vector<vector<int>> &queries)
    {
        // Store each value together with its original node index.
        vector<pair<int, int>> nodes;
        nodes.reserve(n);

        for (int i = 0; i < n; ++i)
        {
            nodes.push_back({nums[i], i});
        }

        // Sorting turns the graph problem into a left-to-right jumping problem.
        sort(nodes.begin(), nodes.end());

        // pos[node] gives the position of an original node in sorted order.
        vector<int> pos(n);

        // values stores only the sorted nums values for easier comparisons.
        vector<int> values(n);

        for (int i = 0; i < n; ++i)
        {
            values[i] = nodes[i].first;
            pos[nodes[i].second] = i;
        }

        // LOG is the number of levels needed for binary lifting.
        int LOG = 1;
        while ((1 << LOG) <= n)
        {
            ++LOG;
        }

        // jump[p][i] is the farthest position reachable after at most 2^p greedy jumps.
        vector<vector<int>> jump(LOG, vector<int>(n));

        // Use two pointers to find the farthest node reachable in one edge.
        int r = 0;

        for (int i = 0; i < n; ++i)
        {
            // r never moves backward, so this whole loop is O(n).
            if (r < i)
            {
                r = i;
            }

            // Extend r while the value difference still allows a direct edge.
            while (r + 1 < n && values[r + 1] - values[i] <= maxDiff)
            {
                ++r;
            }

            // This is the farthest position reachable from i in one jump.
            jump[0][i] = r;
        }

        // Build larger jumps from smaller jumps.
        for (int p = 1; p < LOG; ++p)
        {
            for (int i = 0; i < n; ++i)
            {
                // Two jumps of size 2^(p-1) make one jump of size 2^p.
                jump[p][i] = jump[p - 1][jump[p - 1][i]];
            }
        }

        // Store one answer for every query.
        vector<int> answer;
        answer.reserve(queries.size());

        for (const auto &query : queries)
        {
            // Convert original node indices into sorted positions.
            int left = pos[query[0]];
            int right = pos[query[1]];

            // The graph is undirected, so I always move from left to right.
            if (left > right)
            {
                swap(left, right);
            }

            // A node has distance 0 from itself.
            if (left == right)
            {
                answer.push_back(0);
                continue;
            }

            int current = left;
            int distance = 0;

            // Take the largest groups of jumps that still stop before the target.
            for (int p = LOG - 1; p >= 0; --p)
            {
                if (jump[p][current] < right)
                {
                    // Skip 2^p greedy jumps at once.
                    current = jump[p][current];
                    distance += (1 << p);
                }
            }

            // One final edge must now reach the target.
            if (jump[0][current] >= right)
            {
                answer.push_back(distance + 1);
            }
            else
            {
                // If even the farthest one-step jump cannot move forward enough,
                // the target lies in a different connected component.
                answer.push_back(-1);
            }
        }

        return answer;
    }
};