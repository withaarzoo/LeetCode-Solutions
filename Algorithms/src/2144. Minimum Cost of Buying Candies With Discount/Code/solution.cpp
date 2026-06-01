class Solution
{
public:
    int minimumCost(vector<int> &cost)
    {
        // Sort candies from highest cost to lowest cost
        sort(cost.begin(), cost.end(), greater<int>());

        int total = 0;

        // Add all candies except every third candy
        for (int i = 0; i < cost.size(); i++)
        {
            // Index 2,5,8,... are free candies
            if (i % 3 == 2)
                continue;

            total += cost[i];
        }

        return total;
    }
};