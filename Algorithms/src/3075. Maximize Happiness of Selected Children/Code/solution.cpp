class Solution
{
public:
    long long maximumHappinessSum(vector<int> &happiness, int k)
    {
        // Sort happiness in descending order
        sort(happiness.begin(), happiness.end(), greater<int>());

        long long ans = 0;

        // Pick k children
        for (int i = 0; i < k; i++)
        {
            // Effective happiness after i decrements
            long long curr = happiness[i] - i;

            // Happiness cannot be negative
            if (curr > 0)
            {
                ans += curr;
            }
        }

        return ans;
    }
};
