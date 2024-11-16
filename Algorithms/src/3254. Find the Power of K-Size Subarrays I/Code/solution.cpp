class Solution
{
public:
    vector<int> resultsArray(vector<int> &nums, int k)
    {
        int n = nums.size();
        vector<int> result;

        for (int i = 0; i <= n - k; i++)
        {
            vector<int> subarray(nums.begin() + i, nums.begin() + i + k);

            // Sort the subarray
            vector<int> sortedSubarray = subarray;
            sort(sortedSubarray.begin(), sortedSubarray.end());

            // Check if elements are consecutive
            bool isConsecutive = true;
            for (int j = 1; j < k; j++)
            {
                if (sortedSubarray[j] - sortedSubarray[j - 1] != 1)
                {
                    isConsecutive = false;
                    break;
                }
            }

            // Add the result based on conditions
            if (isConsecutive && subarray == sortedSubarray)
            {
                result.push_back(sortedSubarray.back()); // Max element
            }
            else
            {
                result.push_back(-1);
            }
        }

        return result;
    }
};