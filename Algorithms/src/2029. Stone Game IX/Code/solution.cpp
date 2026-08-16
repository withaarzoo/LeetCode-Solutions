class Solution
{
public:
    bool stoneGameIX(vector<int> &stones)
    {
        // cnt[r] stores how many stones have remainder r when divided by 3.
        int cnt[3] = {0, 0, 0};

        // Count stones in each remainder group.
        for (int stone : stones)
        {
            // Only the remainder matters for the game.
            cnt[stone % 3]++;
        }

        // When the number of remainder-0 stones is even,
        // Alice needs at least one remainder-1 and one remainder-2 stone.
        if (cnt[0] % 2 == 0)
        {
            return cnt[1] > 0 && cnt[2] > 0;
        }

        // When the number of remainder-0 stones is odd,
        // Alice wins if the two useful groups differ by more than 2.
        return abs(cnt[1] - cnt[2]) > 2;
    }
};