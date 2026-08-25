class Solution
{
public:
    int missingMultiple(vector<int> &nums, int k)
    {
        // Store every number so I can check whether a multiple exists quickly.
        unordered_set<int> present(nums.begin(), nums.end());

        // Start from the smallest positive multiple of k.
        int multiple = k;

        // Keep checking multiples until I find one that is missing.
        while (present.count(multiple))
        {
            // Move to the next positive multiple of k.
            multiple += k;
        }

        // This is the smallest multiple of k that does not exist in nums.
        return multiple;
    }
};