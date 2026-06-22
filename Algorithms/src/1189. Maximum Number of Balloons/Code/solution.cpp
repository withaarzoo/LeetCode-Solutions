class Solution
{
public:
    int maxNumberOfBalloons(string text)
    {
        // Store frequency of all lowercase letters
        vector<int> freq(26, 0);

        // Count each character
        for (char ch : text)
        {
            freq[ch - 'a']++;
        }

        // Find the limiting character count
        return min({
            freq['b' - 'a'],     // Need 1 'b'
            freq['a' - 'a'],     // Need 1 'a'
            freq['l' - 'a'] / 2, // Need 2 'l'
            freq['o' - 'a'] / 2, // Need 2 'o'
            freq['n' - 'a']      // Need 1 'n'
        });
    }
};