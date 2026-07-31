class Solution
{
public:
    int minimumPushes(string word)
    {
        // Store the frequency of every lowercase letter
        vector<int> freq(26, 0);

        // Count how many times each letter appears
        for (char ch : word)
        {
            freq[ch - 'a']++;
        }

        // Sort frequencies from largest to smallest
        sort(freq.begin(), freq.end(), greater<int>());

        int ans = 0;

        // Assign push cost according to the position
        for (int i = 0; i < 26; i++)
        {
            // Skip letters that do not appear
            if (freq[i] == 0)
                break;

            // Every 8 letters increase the push count by 1
            int pushes = (i / 8) + 1;

            // Add the total contribution of this letter
            ans += freq[i] * pushes;
        }

        return ans;
    }
};