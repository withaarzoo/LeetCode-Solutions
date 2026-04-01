class Solution
{
public:
    vector<int> survivedRobotsHealths(vector<int> &positions, vector<int> &healths, string directions)
    {
        int n = positions.size();

        // Store robot indices
        vector<int> indices(n);
        for (int i = 0; i < n; i++)
        {
            indices[i] = i;
        }

        // Sort indices based on robot positions
        sort(indices.begin(), indices.end(), [&](int a, int b)
             { return positions[a] < positions[b]; });

        // Stack to keep indices of robots moving right
        stack<int> st;

        for (int idx : indices)
        {
            // If robot moves right, push it into stack
            if (directions[idx] == 'R')
            {
                st.push(idx);
            }
            else
            {
                // Current robot is moving left
                while (!st.empty() && healths[idx] > 0)
                {
                    int topIdx = st.top();

                    // Right robot has smaller health
                    if (healths[topIdx] < healths[idx])
                    {
                        st.pop();
                        healths[idx]--;
                        healths[topIdx] = 0;
                    }
                    // Both have same health
                    else if (healths[topIdx] == healths[idx])
                    {
                        st.pop();
                        healths[topIdx] = 0;
                        healths[idx] = 0;
                    }
                    // Left robot has smaller health
                    else
                    {
                        healths[topIdx]--;
                        healths[idx] = 0;
                    }
                }
            }
        }

        // Collect surviving robots in original order
        vector<int> result;
        for (int i = 0; i < n; i++)
        {
            if (healths[i] > 0)
            {
                result.push_back(healths[i]);
            }
        }

        return result;
    }
};