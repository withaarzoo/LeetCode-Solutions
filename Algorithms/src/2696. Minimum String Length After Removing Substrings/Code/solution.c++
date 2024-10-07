class Solution
{
public:
    int minLength(string s)
    {
        stack<char> stk; // Stack to store characters

        // Traverse through each character in the string
        for (char ch : s)
        {
            // Check if the top of the stack forms "AB" or "CD" with current character
            if (!stk.empty() && ((stk.top() == 'A' && ch == 'B') || (stk.top() == 'C' && ch == 'D')))
            {
                stk.pop(); // Pop if pair is found and remove the substring
            }
            else
            {
                stk.push(ch); // Otherwise, push the character onto the stack
            }
        }

        return stk.size(); // The size of the stack is the minimum length
    }
};