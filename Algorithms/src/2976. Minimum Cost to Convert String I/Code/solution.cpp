class Solution
{
public:
    long long minimumCost(string source, string target,
                          vector<char> &original,
                          vector<char> &changed,
                          vector<int> &cost)
    {

        const long long INF = 1e18;
        vector<vector<long long>> dist(26, vector<long long>(26, INF));

        // Cost to convert a character to itself is 0
        for (int i = 0; i < 26; i++)
        {
            dist[i][i] = 0;
        }

        // Build graph with minimum edge cost
        for (int i = 0; i < original.size(); i++)
        {
            int u = original[i] - 'a';
            int v = changed[i] - 'a';
            dist[u][v] = min(dist[u][v], (long long)cost[i]);
        }

        // Floyd-Warshall on 26 characters
        for (int k = 0; k < 26; k++)
        {
            for (int i = 0; i < 26; i++)
            {
                for (int j = 0; j < 26; j++)
                {
                    if (dist[i][k] + dist[k][j] < dist[i][j])
                    {
                        dist[i][j] = dist[i][k] + dist[k][j];
                    }
                }
            }
        }

        // Convert the string
        long long answer = 0;
        for (int i = 0; i < source.size(); i++)
        {
            int s = source[i] - 'a';
            int t = target[i] - 'a';

            if (dist[s][t] == INF)
                return -1;
            answer += dist[s][t];
        }

        return answer;
    }
};
