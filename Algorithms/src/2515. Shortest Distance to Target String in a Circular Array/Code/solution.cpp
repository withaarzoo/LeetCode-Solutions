class Solution
{
public:
    int closestTarget(vector<string> &words, string target, int startIndex)
    {
        int n = words.size();
        int ans = INT_MAX;

        // Check every index in the array
        for (int i = 0; i < n; i++)
        {
            // If current word matches target
            if (words[i] == target)
            {
                // Normal distance between indices
                int diff = abs(i - startIndex);

                // In circular array, we can also go the other way
                int circularDist = n - diff;

                // Take the minimum possible distance
                ans = min(ans, min(diff, circularDist));
            }
        }

        // If target was never found
        return ans == INT_MAX ? -1 : ans;
    }
};