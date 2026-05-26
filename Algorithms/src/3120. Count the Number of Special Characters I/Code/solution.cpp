class Solution
{
public:
    int numberOfSpecialChars(string word)
    {

        // Store all characters present in the string
        unordered_set<char> st(word.begin(), word.end());

        // Variable to store final answer
        int count = 0;

        // Check every lowercase letter from 'a' to 'z'
        for (char ch = 'a'; ch <= 'z'; ch++)
        {

            // If both lowercase and uppercase exist,
            // then this character is special
            if (st.count(ch) && st.count(ch - 'a' + 'A'))
            {
                count++;
            }
        }

        // Return total special characters
        return count;
    }
};