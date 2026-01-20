class Solution
{
public:
    vector<int> minBitwiseArray(vector<int> &nums)
    {
        vector<int> ans;

        for (int p : nums)
        {
            int found = -1;

            // Try all possible x values
            for (int x = 0; x <= p; x++)
            {
                if ((x | (x + 1)) == p)
                {
                    found = x;
                    break; // smallest x found
                }
            }

            ans.push_back(found);
        }

        return ans;
    }
};
