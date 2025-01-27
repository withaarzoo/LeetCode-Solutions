class Solution
{
public:
    vector<bool> checkIfPrerequisite(int numCourses, vector<vector<int>> &prerequisites, vector<vector<int>> &queries)
    {
        // Initialize the graph
        vector<vector<bool>> graph(numCourses, vector<bool>(numCourses, false));

        // Build the direct edges from prerequisites
        for (const auto &edge : prerequisites)
        {
            graph[edge[0]][edge[1]] = true;
        }

        // Floyd-Warshall to compute transitive closure
        for (int k = 0; k < numCourses; ++k)
        {
            for (int i = 0; i < numCourses; ++i)
            {
                for (int j = 0; j < numCourses; ++j)
                {
                    if (graph[i][k] && graph[k][j])
                    {
                        graph[i][j] = true;
                    }
                }
            }
        }

        // Answer the queries
        vector<bool> result;
        for (const auto &query : queries)
        {
            result.push_back(graph[query[0]][query[1]]);
        }

        return result;
    }
};
