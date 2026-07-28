class Solution
{
public:
    string smallestPalindrome(string s)
    {
        // Store frequency of every lowercase letter
        vector<int> freq(26, 0);

        // Count character frequencies
        for (char c : s)
            freq[c - 'a']++;

        string left = "";
        string middle = "";

        // Build the left half and find the middle character
        for (int i = 0; i < 26; i++)
        {
            // Add half of the occurrences to the left side
            left.append(freq[i] / 2, char('a' + i));

            // If frequency is odd, this character stays in the center
            if (freq[i] % 2)
                middle = char('a' + i);
        }

        // Right half is simply the reverse of the left half
        string right = left;
        reverse(right.begin(), right.end());

        // Combine all three parts
        return left + middle + right;
    }
};