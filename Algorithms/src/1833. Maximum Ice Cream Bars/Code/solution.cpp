class Solution
{
public:
    int maxIceCream(vector<int> &costs, int coins)
    {
        // Maximum possible cost according to constraints
        const int MAX_COST = 100000;

        // Frequency array to count ice cream bars of each cost
        vector<int> freq(MAX_COST + 1, 0);

        // Count occurrences of every cost
        for (int cost : costs)
        {
            freq[cost]++;
        }

        // Stores total ice cream bars purchased
        int answer = 0;

        // Try buying from cheapest cost to most expensive
        for (int cost = 1; cost <= MAX_COST; cost++)
        {

            // Skip if no ice cream bar has this cost
            if (freq[cost] == 0)
                continue;

            // Maximum bars of this cost that can be afforded
            int canBuy = min(freq[cost], coins / cost);

            // Add purchased bars to answer
            answer += canBuy;

            // Deduct spent coins
            coins -= canBuy * cost;

            // No need to continue if coins are exhausted
            if (coins < cost)
                continue;
        }

        return answer;
    }
};