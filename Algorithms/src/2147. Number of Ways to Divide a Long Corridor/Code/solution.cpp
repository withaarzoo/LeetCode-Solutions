class Solution
{
public:
    int numberOfWays(string corridor)
    {
        const int MOD = 1e9 + 7;
        vector<int> seats;

        // Store indices of all seats
        for (int i = 0; i < corridor.size(); i++)
        {
            if (corridor[i] == 'S')
            {
                seats.push_back(i);
            }
        }

        int total = seats.size();
        if (total == 0 || total % 2 != 0)
            return 0;

        long long ways = 1;

        // Multiply gaps between sections
        for (int i = 2; i < total; i += 2)
        {
            ways = (ways * (seats[i] - seats[i - 1])) % MOD;
        }

        return ways;
    }
};
