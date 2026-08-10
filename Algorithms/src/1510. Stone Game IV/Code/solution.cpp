class Solution
{
public:
    bool winnerSquareGame(int n)
    {
        // dp[i] tells me whether the current player can win with i stones.
        vector<bool> dp(n + 1, false);

        // With 0 stones, the current player has no move and loses.
        dp[0] = false;

        // Calculate the result for every number of stones from 1 to n.
        for (int i = 1; i <= n; ++i)
        {
            // Try removing every perfect square that is at most i.
            for (int j = 1; j * j <= i; ++j)
            {
                // If removing j*j gives the opponent a losing state,
                // then I can make this move and win.
                if (!dp[i - j * j])
                {
                    dp[i] = true;

                    // One winning move is enough, so I stop checking squares.
                    break;
                }
            }
        }

        // Return whether Alice can force a win from n stones.
        return dp[n];
    }
};