class Solution
{
public:
    string smallestSubsequence(string s)
    {
        // Store how many times each character still appears
        vector<int> freq(26, 0);
        for (char c : s)
            freq[c - 'a']++;

        // Track whether a character is already inside the stack
        vector<bool> inStack(26, false);

        // This string works as our stack
        string st;

        for (char c : s)
        {
            // One occurrence of this character has now been processed
            freq[c - 'a']--;

            // Skip duplicate characters already included
            if (inStack[c - 'a'])
                continue;

            // Remove larger characters if they appear again later
            while (!st.empty() &&
                   st.back() > c &&
                   freq[st.back() - 'a'] > 0)
            {

                // Mark removed character as no longer inside the stack
                inStack[st.back() - 'a'] = false;
                st.pop_back();
            }

            // Add current character to the stack
            st.push_back(c);
            inStack[c - 'a'] = true;
        }

        // The stack already contains the final answer
        return st;
    }
};