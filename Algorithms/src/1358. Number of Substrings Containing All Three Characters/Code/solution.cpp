class Solution
{
public:
    int numberOfSubstrings(string s)
    {
        // Store the frequency of 'a', 'b', and 'c' inside the current window
        vector<int> freq(3, 0);

        int left = 0;
        int ans = 0;
        int n = s.size();

        // Expand the window one character at a time
        for (int right = 0; right < n; right++)
        {

            // Add the current character into the window
            freq[s[right] - 'a']++;

            // Keep shrinking while the window contains all three characters
            while (freq[0] > 0 && freq[1] > 0 && freq[2] > 0)
            {

                // Every substring ending from 'right' to the last index is valid
                ans += (n - right);

                // Remove the leftmost character before shrinking the window
                freq[s[left] - 'a']--;

                // Move the left pointer forward
                left++;
            }
        }

        // Return the total number of valid substrings
        return ans;
    }
};