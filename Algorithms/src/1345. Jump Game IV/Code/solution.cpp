class Solution
{
public:
    int minJumps(vector<int> &arr)
    {
        int n = arr.size();

        // If array has only one element, no jump is needed
        if (n == 1)
            return 0;

        // Map each value to all its indices
        unordered_map<int, vector<int>> mp;

        for (int i = 0; i < n; i++)
        {
            mp[arr[i]].push_back(i);
        }

        // Queue for BFS
        queue<int> q;

        // Visited array to avoid revisiting indices
        vector<bool> visited(n, false);

        q.push(0);
        visited[0] = true;

        int steps = 0;

        while (!q.empty())
        {

            int size = q.size();

            // Process one BFS level at a time
            while (size--)
            {

                int idx = q.front();
                q.pop();

                // If last index reached, return answer
                if (idx == n - 1)
                {
                    return steps;
                }

                // Move to index - 1
                if (idx - 1 >= 0 && !visited[idx - 1])
                {
                    visited[idx - 1] = true;
                    q.push(idx - 1);
                }

                // Move to index + 1
                if (idx + 1 < n && !visited[idx + 1])
                {
                    visited[idx + 1] = true;
                    q.push(idx + 1);
                }

                // Move to all same-value indices
                for (int nextIdx : mp[arr[idx]])
                {

                    if (!visited[nextIdx])
                    {
                        visited[nextIdx] = true;
                        q.push(nextIdx);
                    }
                }

                // Clear the list so we never process
                // same-value indices again
                mp[arr[idx]].clear();
            }

            // One BFS level completed
            steps++;
        }

        return -1;
    }
};