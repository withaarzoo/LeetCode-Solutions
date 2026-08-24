class Solution
{
public:
    int stoneGameVIII(vector<int> &stones)
    {
        // Store the prefix sum directly in stones to avoid using extra space.
        for (int i = 1; i < stones.size(); ++i)
        {
            stones[i] += stones[i - 1];
        }

        // If Alice removes all stones, she gets the total sum.
        int best = stones.back();

        // Check every possible smaller prefix from right to left.
        for (int i = stones.size() - 2; i >= 1; --i)
        {
            // Either keep the current best move or take this prefix
            // and subtract the best score difference the opponent can force.
            best = max(best, stones[i] - best);
        }

        // Return the maximum score difference Alice can achieve.
        return best;
    }
};