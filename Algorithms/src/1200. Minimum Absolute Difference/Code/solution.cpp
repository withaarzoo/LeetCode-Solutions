class Solution
{
public:
    vector<vector<int>> minimumAbsDifference(vector<int> &arr)
    {
        // Step 1: Sort the array
        sort(arr.begin(), arr.end());

        int minDiff = INT_MAX;
        vector<vector<int>> result;

        // Step 2: Find the minimum difference
        for (int i = 1; i < arr.size(); i++)
        {
            minDiff = min(minDiff, arr[i] - arr[i - 1]);
        }

        // Step 3: Collect all pairs with minimum difference
        for (int i = 1; i < arr.size(); i++)
        {
            if (arr[i] - arr[i - 1] == minDiff)
            {
                result.push_back({arr[i - 1], arr[i]});
            }
        }

        return result;
    }
};
