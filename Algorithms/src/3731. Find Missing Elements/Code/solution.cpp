class Solution
{
public:
    vector<int> findMissingElements(vector<int> &nums)
    {
        // Find the smallest and largest values in the array
        int mn = *min_element(nums.begin(), nums.end());
        int mx = *max_element(nums.begin(), nums.end());

        // Store every number for O(1) average lookup
        unordered_set<int> seen(nums.begin(), nums.end());

        // Store all missing numbers
        vector<int> ans;

        // Check every value in the original range
        for (int num = mn; num <= mx; num++)
        {
            // If the number is not present, it is missing
            if (!seen.count(num))
            {
                ans.push_back(num);
            }
        }

        // Return the sorted missing numbers
        return ans;
    }
};