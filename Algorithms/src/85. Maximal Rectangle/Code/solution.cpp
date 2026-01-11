class Solution
{
public:
    int largestRectangleArea(vector<int> &heights)
    {
        stack<int> st;
        int maxArea = 0;
        heights.push_back(0); // sentinel

        for (int i = 0; i < heights.size(); i++)
        {
            while (!st.empty() && heights[st.top()] > heights[i])
            {
                int h = heights[st.top()];
                st.pop();
                int w = st.empty() ? i : i - st.top() - 1;
                maxArea = max(maxArea, h * w);
            }
            st.push(i);
        }
        heights.pop_back();
        return maxArea;
    }

    int maximalRectangle(vector<vector<char>> &matrix)
    {
        if (matrix.empty())
            return 0;

        int cols = matrix[0].size();
        vector<int> heights(cols, 0);
        int ans = 0;

        for (auto &row : matrix)
        {
            for (int j = 0; j < cols; j++)
            {
                heights[j] = (row[j] == '1') ? heights[j] + 1 : 0;
            }
            ans = max(ans, largestRectangleArea(heights));
        }
        return ans;
    }
};
