class Solution
{
public:
    vector<int> arrayRankTransform(vector<int> &arr)
    {
        // Create a copy so the original order remains unchanged
        vector<int> sorted = arr;

        // Sort the copied array
        sort(sorted.begin(), sorted.end());

        // Store each unique value with its rank
        unordered_map<int, int> rank;
        int currentRank = 1;

        // Assign ranks only to unique values
        for (int num : sorted)
        {
            if (!rank.count(num))
            {
                rank[num] = currentRank++;
            }
        }

        // Replace every element with its assigned rank
        for (int &num : arr)
        {
            num = rank[num];
        }

        // Return the transformed array
        return arr;
    }
};