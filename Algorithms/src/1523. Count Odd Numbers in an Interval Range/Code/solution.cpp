class Solution
{
public:
    int countOdds(int low, int high)
    {
        // Helper lambda: count of odd numbers from 1 to x
        auto oddsUpTo = [](int x) -> int
        {
            // Integer division automatically takes floor
            return (x + 1) / 2;
        };

        // Odds in [low, high] = oddsUpTo(high) - oddsUpTo(low - 1)
        return oddsUpTo(high) - oddsUpTo(low - 1);
    }
};
