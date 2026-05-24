class Solution
{
public:
    vector<int> dp;

    // DFS function returns maximum jumps starting from index i
    int dfs(int i, vector<int> &arr, int d)
    {

        // If already computed, return stored answer
        if (dp[i] != -1)
            return dp[i];

        // Minimum answer is 1 because current index is counted
        int ans = 1;

        // Try jumping to the right
        for (int j = i + 1; j <= min((int)arr.size() - 1, i + d); j++)
        {

            // Cannot jump further if value is greater or equal
            if (arr[j] >= arr[i])
                break;

            // Take best possible path
            ans = max(ans, 1 + dfs(j, arr, d));
        }

        // Try jumping to the left
        for (int j = i - 1; j >= max(0, i - d); j--)
        {

            // Cannot jump further if value is greater or equal
            if (arr[j] >= arr[i])
                break;

            // Take best possible path
            ans = max(ans, 1 + dfs(j, arr, d));
        }

        // Store and return answer
        return dp[i] = ans;
    }

    int maxJumps(vector<int> &arr, int d)
    {

        int n = arr.size();

        // Initialize DP array with -1 meaning unvisited
        dp.assign(n, -1);

        int answer = 1;

        // Start DFS from every index
        for (int i = 0; i < n; i++)
        {
            answer = max(answer, dfs(i, arr, d));
        }

        return answer;
    }
};