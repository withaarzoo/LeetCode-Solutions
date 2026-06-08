class Solution
{
public:
    vector<int> pivotArray(vector<int> &nums, int pivot)
    {
        // Store elements smaller than pivot
        vector<int> smaller;

        // Store elements equal to pivot
        vector<int> equal;

        // Store elements greater than pivot
        vector<int> greater;

        // Classify every element into one of three groups
        for (int num : nums)
        {
            if (num < pivot)
            {
                smaller.push_back(num);
            }
            else if (num == pivot)
            {
                equal.push_back(num);
            }
            else
            {
                greater.push_back(num);
            }
        }

        // Append equal elements after smaller elements
        smaller.insert(smaller.end(), equal.begin(), equal.end());

        // Append greater elements at the end
        smaller.insert(smaller.end(), greater.begin(), greater.end());

        // Return the final partitioned array
        return smaller;
    }
};