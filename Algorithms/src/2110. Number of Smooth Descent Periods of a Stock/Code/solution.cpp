class Solution
{
public:
    long long getDescentPeriods(vector<int> &prices)
    {
        long long ans = 1; // first day is always a valid period
        long long len = 1; // length of current smooth descent

        for (int i = 1; i < prices.size(); i++)
        {
            // check if current price is exactly 1 less than previous
            if (prices[i] == prices[i - 1] - 1)
            {
                len++; // extend the descent
            }
            else
            {
                len = 1; // reset the descent
            }
            ans += len; // add all subarrays ending here
        }
        return ans;
    }
};
